// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package any_2_evm_free_offramp

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

var EVM2EVMFreeOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlobVerifier\",\"outputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"}],\"name\":\"setBlobVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005dc238038062005dc283398101604081905262000034916200063c565b6000805460ff191681556001908a908a908a908a908a908a908a908a908a90829082908690869089903390819081620000b45760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ee57620000ee8162000385565b5050506001600160a01b0381166200011957604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200015b5760405162d8548360e71b815260040160405180910390fd5b81516200017090600490602085019062000436565b5060005b82518110156200023b57600082828151811062000195576200019562000796565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001df57620001df62000796565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b0292909116919091179055506200023381620007ac565b905062000174565b5050600580546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600855600993909355600a55600b91909155875116620002b9576040516342bcdf7f60e11b815260040160405180910390fd5b5050506080958652505060a0929092528051600e8054602084015160408501516001600160401b03908116600160c01b026001600160c01b0363ffffffff909316600160a01b026001600160c01b03199094166001600160a01b03968716179390931791909116919091179091556060830151600f80549490960151821668010000000000000000026001600160801b031990941691161791909117909255600d8054919092166001600160a01b031991909116179055151560c05250620007d4975050505050505050565b336001600160a01b03821603620003df5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ab565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200048e579160200282015b828111156200048e57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000457565b506200049c929150620004a0565b5090565b5b808211156200049c5760008155600101620004a1565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715620004f257620004f2620004b7565b60405290565b6001600160a01b03811681146200050e57600080fd5b50565b80516200051e81620004f8565b919050565b80516001600160401b03811681146200051e57600080fd5b600082601f8301126200054d57600080fd5b815160206001600160401b03808311156200056c576200056c620004b7565b8260051b604051601f19603f83011681018181108482111715620005945762000594620004b7565b604052938452858101830193838101925087851115620005b357600080fd5b83870191505b84821015620005df578151620005cf81620004f8565b83529183019190830190620005b9565b979650505050505050565b600060408284031215620005fd57600080fd5b604080519081016001600160401b0381118282101715620006225762000622620004b7565b604052825181526020928301519281019290925250919050565b6000806000806000806000806000898b036101c08112156200065d57600080fd5b8a51995060208b0151985060a0603f19820112156200067b57600080fd5b5062000686620004cd565b60408b01516200069681620004f8565b815260608b015163ffffffff81168114620006b057600080fd5b6020820152620006c360808c0162000523565b6040820152620006d660a08c0162000523565b6060820152620006e960c08c0162000523565b60808201529650620006fe60e08b0162000511565b95506200070f6101008b0162000511565b6101208b01519095506001600160401b03808211156200072e57600080fd5b6200073c8d838e016200053b565b95506101408c01519150808211156200075457600080fd5b50620007638c828d016200053b565b935050620007768b6101608c01620005ea565b9150620007876101a08b0162000511565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b600060018201620007cd57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c0516155a96200081960003960006115c50152600081816102db0152611e1e0152600081816102b801528181611df901526136ed01526155a96000f3fe608060405234801561001057600080fd5b50600436106102ad5760003560e01c8063814118341161017b578063b1dc65a4116100d8578063c3f909d41161008c578063e3d0e71211610071578063e3d0e71214610797578063eb511dd4146107aa578063f2fde38b146107bd57600080fd5b8063c3f909d41461066c578063c9029f6a1461078457600080fd5b8063bbe4f6db116100bd578063bbe4f6db1461061a578063be9b03f114610646578063c0d786551461065957600080fd5b8063b1dc65a4146105f4578063b4069b311461060757600080fd5b806390c2339b1161012f578063a639d1c011610114578063a639d1c0146105b0578063afcb95d7146105c3578063b0f479a1146105e357600080fd5b806390c2339b14610562578063918725431461059d57600080fd5b80638456cb59116101605780638456cb591461053c57806389c06568146105445780638da5cb5b1461054c57600080fd5b806381411834146104f757806381ff70481461050c57600080fd5b80633f4ba83a116102295780636133dc24116101dd578063694ec2b1116101c2578063694ec2b1146104c9578063744b92e2146104dc57806379ba5097146104ef57600080fd5b80636133dc24146104a3578063681fba16146104b457600080fd5b80634741062e1161020e5780634741062e14610467578063599f6431146104875780635c975abb1461049857600080fd5b80633f4ba83a1461044c5780634352fa9f1461045457600080fd5b8063181f5a77116102805780632d0335ab116102655780632d0335ab146103df578063351f0faf1461042b57806339aa92641461043957600080fd5b8063181f5a77146103715780632222dd42146103ba57600080fd5b8063087ae6df146102b2578063108ee5fc1461030a578063142a98fc1461031f578063147809b314610359575b600080fd5b604080517f000000000000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006020820152015b60405180910390f35b61031d610318366004613e93565b6107d0565b005b61034c61032d366004613ed1565b67ffffffffffffffff1660009081526010602052604090205460ff1690565b6040516103019190613f1d565b610361610886565b6040519015158152602001610301565b6103ad6040518060400160405280602081526020017f45564d3245564d537562736372697074696f6e4f666652616d7020312e302e3081525081565b6040516103019190613fb6565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610301565b6104126103ed366004613e93565b6001600160a01b031660009081526017602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff9091168152602001610301565b61031d6102ad366004613fd4565b61031d610447366004613e93565b610913565b61031d61094a565b61031d6104623660046141b1565b61095c565b61047a610475366004614215565b610bb1565b604051610301919061428d565b6005546001600160a01b03166103c7565b60005460ff16610361565b600d546001600160a01b03166103c7565b6104bc610c79565b60405161030191906142a0565b61031d6104d73660046143ca565b610d3e565b61031d6104ea3660046144f0565b610db7565b61031d6110dd565b6104ff6111c5565b6040516103019190614562565b6013546011546040805163ffffffff80851682526401000000009094049093166020840152820152606001610301565b61031d611227565b6104bc611237565b60005461010090046001600160a01b03166103c7565b61056a611297565b60405161030191908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b61031d6105ab366004614575565b611338565b61031d6105be366004613e93565b61146b565b604080516001815260006020820181905291810191909152606001610301565b600c546001600160a01b03166103c7565b61031d610602366004614610565b6114a2565b6103c7610615366004613e93565b611a75565b6103c7610628366004613e93565b6001600160a01b039081166000908152600360205260409020541690565b61031d6106543660046147f2565b611b63565b61031d610667366004613e93565b6121f2565b6107286040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a081018252600e546001600160a01b038116825263ffffffff74010000000000000000000000000000000000000000820416602083015267ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811692820192909252600f548083166060830152680100000000000000009004909116608082015290565b6040805182516001600160a01b0316815260208084015163ffffffff16908201528282015167ffffffffffffffff908116928201929092526060808401518316908201526080928301519091169181019190915260a001610301565b61031d61079236600461493b565b612251565b61031d6107a53660046149ed565b6123f0565b61031d6107b83660046144f0565b612ca9565b61031d6107cb366004613e93565b612e82565b6107d8612e93565b6001600160a01b038116610818576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28910160405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa1580156108e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061090d9190614aba565b15905090565b61091b612e93565b6005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b610952612e93565b61095a612ef2565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561099957506005546001600160a01b03163314155b156109d0576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610a0c576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610a66576006600060078381548110610a3157610a31614ad7565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610a5f81614b35565b9050610a12565b5060005b82811015610b96576000858281518110610a8657610a86614ad7565b6020026020010151905060006001600160a01b0316816001600160a01b031603610adc576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610aee57610aee614ad7565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610b5357610b53614ad7565b6020026020010151604051610b7d9291906001600160a01b03929092168252602082015260400190565b60405180910390a150610b8f81614b35565b9050610a6a565b508351610baa906007906020870190613ddc565b5050505050565b80516060908067ffffffffffffffff811115610bcf57610bcf614010565b604051908082528060200260200182016040528015610bf8578160200160208202803683370190505b50915060005b81811015610c725760066000858381518110610c1c57610c1c614ad7565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610c5757610c57614ad7565b6020908102919091010152610c6b81614b35565b9050610bfe565b5050919050565b60045460609067ffffffffffffffff811115610c9757610c97614010565b604051908082528060200260200182016040528015610cc0578160200160208202803683370190505b50905060005b600454811015610d3a57610d0060048281548110610ce657610ce6614ad7565b6000918252602090912001546001600160a01b0316611a75565b828281518110610d1257610d12614ad7565b6001600160a01b0390921660209283029190910190910152610d3381614b35565b9050610cc6565b5090565b333014610d77576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60808101515115610dab57610d9481608001518260c00151612f8e565b610dab8160a001518260c001518360400151613121565b610db4816131bb565b50565b610dbf612e93565b6004546000819003610dfd576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610e8b576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614610eda576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004610ee9600185614b4f565b81548110610ef957610ef9614ad7565b9060005260206000200160009054906101000a90046001600160a01b03169050600482602001516bffffffffffffffffffffffff1681548110610f3e57610f3e614ad7565b6000918252602090912001546001600160a01b03166004610f60600186614b4f565b81548110610f7057610f70614ad7565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600483602001516bffffffffffffffffffffffff1681548110610fc457610fc4614ad7565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff909216740100000000000000000000000000000000000000000291909216179055600480548061104e5761104e614b66565b600082815260208082206000199084018101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b03878116808452600383526040808520949094558351908152908716918101919091527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b0316331461113c5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060601680548060200260200160405190810160405280929190818152602001828054801561121d57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116111ff575b5050505050905090565b61122f612e93565b61095a613293565b6060600480548060200260200160405190810160405280929190818152602001828054801561121d576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116111ff575050505050905090565b6112c26040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b546060820181905242906000906112fc9083614b4f565b60208401518451919250611328916113149084614b95565b85604001516113239190614bb4565b61331b565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561137557506005546001600160a01b03163314155b156113ac576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff11611400576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61140a6008613331565b602081015160098190558151600855600a54611426919061331b565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b611473612e93565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916114f891849163ffffffff851691908e908e90819084018382808284376000920191909152506133de92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260115480825260125460ff808216602085015261010090910416928201929092529083146115b35760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401611133565b6115c18b8b8b8b8b8b613457565b60007f00000000000000000000000000000000000000000000000000000000000000001561161e576002826020015183604001516115ff9190614bcc565b6116099190614bf1565b611614906001614bcc565b60ff169050611634565b602082015161162e906001614bcc565b60ff1690505b8881146116835760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401611133565b8887146116d25760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401611133565b3360009081526014602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561171557611715613eee565b600281111561172657611726613eee565b905250905060028160200151600281111561174357611743613eee565b14801561177d57506016816000015160ff168154811061176557611765614ad7565b6000918252602090912001546001600160a01b031633145b6117c95760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401611133565b5050505050600088886040516117e0929190614c3a565b6040519081900381206117f7918c90602001614c4a565b604051602081830303815290604052805190602001209050611817613e4a565b604080518082019091526000808252602082015260005b88811015611a5357600060018588846020811061184d5761184d614ad7565b61185a91901a601b614bcc565b8d8d8681811061186c5761186c614ad7565b905060200201358c8c8781811061188557611885614ad7565b90506020020135604051600081526020016040526040516118c2949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156118e4573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526014602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561193957611939613eee565b600281111561194a5761194a613eee565b905250925060018360200151600281111561196757611967613eee565b146119b45760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401611133565b8251849060ff16601f81106119cb576119cb614ad7565b602002015115611a1d5760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401611133565b600184846000015160ff16601f8110611a3857611a38614ad7565b9115156020909202015250611a4c81614b35565b905061182e565b5050505063ffffffff8110611a6a57611a6a614c66565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611ac9576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611b38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b5c9190614c95565b9392505050565b60005460ff1615611bb65760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611133565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c2d9190614aba565b15611c63576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c546001600160a01b0316611ca5576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060820151516000819003611ce6576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115611d0157611d01614010565b604051908082528060200260200182016040528015611d7757816020015b60408051610120810182526000808252602080830182905292820181905260608083018290526080830182905260a0830181905260c0830181905260e08301526101008201528252600019909201910181611d1f5790505b50905060008267ffffffffffffffff811115611d9557611d95614010565b604051908082528060200260200182016040528015611dbe578160200160208202803683370190505b5090506000611e6f7f3997e2cfd3ccacf768662bd35c3dbf323724407d75aae3019c04f4aa59b1193f600e54604080516020808201949094527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201526001600160a01b039092166080808401919091528151808403909101815260a09092019052805191012090565b905060005b84811015611f0957600087606001518281518110611e9457611e94614ad7565b6020026020010151806020019051810190611eaf9190614dcc565b9050611ebb81846134de565b848381518110611ecd57611ecd614ad7565b60200260200101818152505080858381518110611eec57611eec614ad7565b60200260200101819052505080611f0290614b35565b9050611e74565b50600080611f2a8489608001518a60a001518b60c001518c60e001516135f1565b91509150868015611f635750600e5474010000000000000000000000000000000000000000900463ffffffff16611f618342614b4f565b105b15611f9a576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b86811015611a6a576000868281518110611fb957611fb9614ad7565b602002602001015190506000611fec826020015167ffffffffffffffff1660009081526010602052604090205460ff1690565b9050600281600381111561200257612002613eee565b0361204b5760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611133565b612054826136e9565b60208281015167ffffffffffffffff166000908152601090915260408120805460ff1916600117905561208e61208984613830565b613aaf565b60208085015167ffffffffffffffff1660009081526010909152604090208054919250829160ff191660018360038111156120cb576120cb613eee565b0217905550826020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf8260405161210e9190613f1d565b60405180910390a2608083015160608401516001600160a01b0316600090815260176020526040812054909167ffffffffffffffff9081169161215391166001614edd565b67ffffffffffffffff1614905080801561217e5750600282600381111561217c5761217c613eee565b145b156121dd5760608401516001600160a01b03166000908152601760205260408120805467ffffffffffffffff16916121b583614f09565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b50505050806121eb90614b35565b9050611f9d565b6121fa612e93565b600c805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b612259612e93565b80516001600160a01b031661229a576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600e805460208085018051604080880180516001600160a01b039098167fffffffffffffffff00000000000000000000000000000000000000000000000090961686177401000000000000000000000000000000000000000063ffffffff948516021777ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff998a16021790965560608089018051600f80546080808e018051948e167fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169290921768010000000000000000948e1694909402939093179091558451988952955190941695870195909552955187169085015291518516938301939093529151909216908201527fd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa09060a001611460565b855185518560ff16601f831115612463576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401611133565b806000036124cd576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401611133565b81831461255b576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401611133565b612566816003614b95565b83116125ce576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401611133565b6125d6612e93565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b6015541561275d5760155460009061262e90600190614b4f565b905060006015828154811061264557612645614ad7565b6000918252602082200154601680546001600160a01b039092169350908490811061267257612672614ad7565b60009182526020808320909101546001600160a01b0385811684526014909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000908116909155929091168084529220805490911690556015805491925090806126e5576126e5614b66565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff19169055019055601680548061272557612725614b66565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff1916905501905550612614915050565b60005b815151811015612b22576000601460008460000151848151811061278657612786614ad7565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff1660028111156127c3576127c3613eee565b1461282a576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401611133565b6040805180820190915260ff8216815260016020820152825180516014916000918590811061285b5761285b614ad7565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156128d1576128d1613eee565b0217905550600091506128e19050565b60146000846020015184815181106128fb576128fb614ad7565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff16600281111561293857612938613eee565b1461299f576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401611133565b6040805180820190915260ff8216815260208101600281525060146000846020015184815181106129d2576129d2614ad7565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612a4857612a48613eee565b021790555050825180516015925083908110612a6657612a66614ad7565b6020908102919091018101518254600181018455600093845292829020909201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039093169290921790915582015180516016919083908110612aca57612aca614ad7565b602090810291909101810151825460018101845560009384529190922001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909216919091179055612b1b81614b35565b9050612760565b5060408101516012805460ff191660ff909216919091179055601380547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092612b96928692908216911617614f30565b92506101000a81548163ffffffff021916908363ffffffff160217905550612bf54630601360009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613be3565b6011819055825180516012805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560135460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612c94988b98919763ffffffff909216969095919491939192614f4f565b60405180910390a15050505050505050505050565b612cb1612e93565b6001600160a01b0382161580612cce57506001600160a01b038116155b15612d05576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612d94576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201805473ffffffffffffffffffffffffffffffffffffffff1916821790558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c91015b60405180910390a1505050565b612e8a612e93565b610db481613c70565b60005461010090046001600160a01b0316331461095a5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401611133565b60005460ff16612f445760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401611133565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6000805b835181101561308557600060066000868481518110612fb357612fb3614ad7565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000205490508060000361304257848281518110612ff857612ff8614ad7565b60200260200101516040517f9a655f7b00000000000000000000000000000000000000000000000000000000815260040161113391906001600160a01b0391909116815260200190565b83828151811061305457613054614ad7565b6020026020010151816130679190614b95565b6130719084614bb4565b9250508061307e90614b35565b9050612f92565b50801561311c576130966008613331565b600a548111156130d2576040517f3bfa6f3800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600860020160008282546130e79190614b4f565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba37890602001612e75565b505050565b815183511461315c576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b83518110156131b5576131a584828151811061317d5761317d614ad7565b602002602001015184838151811061319757613197614ad7565b602002602001015184613d2c565b6131ae81614b35565b905061315f565b50505050565b60408101516001600160a01b03163b6131d15750565b600c546040517f5c9fa11e0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690635c9fa11e9061321a908490600401614fe5565b6020604051808303816000875af1158015613239573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061325d9190614aba565b610db4576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005460ff16156132e65760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611133565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612f713390565b600081831061332a5781611b5c565b5090919050565b600181015460028201544291148061334c5750808260030154145b15613355575050565b816001015482600201541115613397576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260030154826133a99190614b4f565b600184015484549192506133d0916133c19084614b95565b85600201546113239190614bb4565b600284015550600390910155565b306001600160a01b031663be9b03f182806020019051810190613401919061519a565b60006040518363ffffffff1660e01b8152600401613420929190615362565b600060405180830381600087803b15801561343a57600080fd5b505af115801561344e573d6000803e3d6000fd5b50505050505050565b6000613464826020614b95565b61346f856020614b95565b61347b88610144614bb4565b6134859190614bb4565b61348f9190614bb4565b61349a906000614bb4565b905036811461344e576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401611133565b60008060001b828460200151856040015186606001518760a00151805190602001208860c001516040516020016135159190614562565b604051602081830303815290604052805190602001208960e00151604051602001613540919061428d565b604051602081830303815290604052805190602001208a61010001518b608001516040516020016135d39a99989796959493929190998a5260208a019890985267ffffffffffffffff96871660408a01526001600160a01b0395861660608a015293909416608088015260a087019190915260c086015260e0850191909152610100840152166101208201526101400190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600d546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce9061364e908c908c908c908c908c90600401615459565b6020604051808303816000875af115801561366d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061369191906154ab565b9050600081116136cd576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a6136d99084614b4f565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146137495780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401611133565b600f5460c0820151516801000000000000000090910467ffffffffffffffff16108061377f57508060e00151518160c001515114155b156137c85760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611133565b600f5460a08201515167ffffffffffffffff9091161015610db457600f5460a0820151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301526024820152604401611133565b613881604051806101000160405280600081526020016060815260200160006001600160a01b0316815260200160608152602001606081526020016060815260200160608152602001600081525090565b60c08201515160008167ffffffffffffffff8111156138a2576138a2614010565b6040519080825280602002602001820160405280156138cb578160200160208202803683370190505b50905060008267ffffffffffffffff8111156138e9576138e9614010565b604051908082528060200260200182016040528015613912578160200160208202803683370190505b50905060005b83811015613a1a5760006139488760c00151838151811061393b5761393b614ad7565b6020026020010151613d7a565b90508083838151811061395d5761395d614ad7565b60200260200101906001600160a01b031690816001600160a01b031681525050806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156139bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139df9190614c95565b8483815181106139f1576139f1614ad7565b6001600160a01b039092166020928302919091019091015250613a1381614b35565b9050613918565b50604051806101000160405280866000015181526020018660400151604051602001613a5591906001600160a01b0391909116815260200190565b604051602081830303815290604052815260200186606001516001600160a01b031681526020018660a0015181526020018381526020018281526020018660e0015181526020018661010001518152509350505050919050565b6040517f694ec2b1000000000000000000000000000000000000000000000000000000008152600090309063694ec2b190613aee908590600401614fe5565b600060405180830381600087803b158015613b0857600080fd5b505af1925050508015613b19575060015b613bdb573d808015613b47576040519150601f19603f3d011682016040523d82523d6000602084013e613b4c565b606091505b50613b56816154c4565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da80000000000000000000000000000000000000000000000000000000003613ba75750600392915050565b806040517fcf19edfd0000000000000000000000000000000000000000000000000000000081526004016111339190613fb6565b506002919050565b6000808a8a8a8a8a8a8a8a8a604051602001613c0799989796959493929190615514565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b03821603613cc85760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401611133565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401613420565b6001600160a01b038181166000908152600360205260409020541680613dd7576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401611133565b919050565b828054828255906000526020600020908101928215613e3e579160200282015b82811115613e3e578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190613dfc565b50610d3a929150613e69565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610d3a5760008155600101613e6a565b6001600160a01b0381168114610db457600080fd5b600060208284031215613ea557600080fd5b8135611b5c81613e7e565b67ffffffffffffffff81168114610db457600080fd5b8035613dd781613eb0565b600060208284031215613ee357600080fd5b8135611b5c81613eb0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310613f58577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b83811015613f79578181015183820152602001613f61565b838111156131b55750506000910152565b60008151808452613fa2816020860160208601613f5e565b601f01601f19169290920160200192915050565b602081526000611b5c6020830184613f8a565b8035613dd781613e7e565b600060208284031215613fe657600080fd5b813567ffffffffffffffff811115613ffd57600080fd5b82016101008185031215611b5c57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff8111828210171561406357614063614010565b60405290565b604051610120810167ffffffffffffffff8111828210171561406357614063614010565b604051601f8201601f1916810167ffffffffffffffff811182821017156140b6576140b6614010565b604052919050565b600067ffffffffffffffff8211156140d8576140d8614010565b5060051b60200190565b600082601f8301126140f357600080fd5b81356020614108614103836140be565b61408d565b82815260059290921b8401810191818101908684111561412757600080fd5b8286015b8481101561414b57803561413e81613e7e565b835291830191830161412b565b509695505050505050565b600082601f83011261416757600080fd5b81356020614177614103836140be565b82815260059290921b8401810191818101908684111561419657600080fd5b8286015b8481101561414b578035835291830191830161419a565b600080604083850312156141c457600080fd5b823567ffffffffffffffff808211156141dc57600080fd5b6141e8868387016140e2565b935060208501359150808211156141fe57600080fd5b5061420b85828601614156565b9150509250929050565b60006020828403121561422757600080fd5b813567ffffffffffffffff81111561423e57600080fd5b61424a848285016140e2565b949350505050565b600081518084526020808501945080840160005b8381101561428257815187529582019590820190600101614266565b509495945050505050565b602081526000611b5c6020830184614252565b6020808252825182820181905260009190848201906040850190845b818110156142e15783516001600160a01b0316835292840192918401916001016142bc565b50909695505050505050565b600067ffffffffffffffff82111561430757614307614010565b50601f01601f191660200190565b600082601f83011261432657600080fd5b8135614334614103826142ed565b81815284602083860101111561434957600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261437757600080fd5b81356020614387614103836140be565b82815260059290921b840181019181810190868411156143a657600080fd5b8286015b8481101561414b5780356143bd81613e7e565b83529183019183016143aa565b6000602082840312156143dc57600080fd5b813567ffffffffffffffff808211156143f457600080fd5b90830190610100828603121561440957600080fd5b61441161403f565b8235815260208301358281111561442757600080fd5b61443387828601614315565b60208301525061444560408401613fc9565b604082015260608301358281111561445c57600080fd5b61446887828601614315565b60608301525060808301358281111561448057600080fd5b61448c87828601614366565b60808301525060a0830135828111156144a457600080fd5b6144b087828601614366565b60a08301525060c0830135828111156144c857600080fd5b6144d487828601614156565b60c08301525060e083013560e082015280935050505092915050565b6000806040838503121561450357600080fd5b823561450e81613e7e565b9150602083013561451e81613e7e565b809150509250929050565b600081518084526020808501945080840160005b838110156142825781516001600160a01b03168752958201959082019060010161453d565b602081526000611b5c6020830184614529565b60006040828403121561458757600080fd5b6040516040810181811067ffffffffffffffff821117156145aa576145aa614010565b604052823581526020928301359281019290925250919050565b60008083601f8401126145d657600080fd5b50813567ffffffffffffffff8111156145ee57600080fd5b6020830191508360208260051b850101111561460957600080fd5b9250929050565b60008060008060008060008060e0898b03121561462c57600080fd5b606089018a81111561463d57600080fd5b8998503567ffffffffffffffff8082111561465757600080fd5b818b0191508b601f83011261466b57600080fd5b81358181111561467a57600080fd5b8c602082850101111561468c57600080fd5b6020830199508098505060808b01359150808211156146aa57600080fd5b6146b68c838d016145c4565b909750955060a08b01359150808211156146cf57600080fd5b506146dc8b828c016145c4565b999c989b50969995989497949560c00135949350505050565b600082601f83011261470657600080fd5b81356020614716614103836140be565b82815260059290921b8401810191818101908684111561473557600080fd5b8286015b8481101561414b57803561474c81613eb0565b8352918301918301614739565b600082601f83011261476a57600080fd5b8135602061477a614103836140be565b82815260059290921b8401810191818101908684111561479957600080fd5b8286015b8481101561414b57803567ffffffffffffffff8111156147bd5760008081fd5b6147cb8986838b0101614315565b84525091830191830161479d565b8015158114610db457600080fd5b8035613dd7816147d9565b6000806040838503121561480557600080fd5b823567ffffffffffffffff8082111561481d57600080fd5b90840190610100828703121561483257600080fd5b61483a61403f565b82358281111561484957600080fd5b614855888286016146f5565b82525060208301358281111561486a57600080fd5b61487688828601614366565b60208301525060408301358281111561488e57600080fd5b61489a88828601614156565b6040830152506060830135828111156148b257600080fd5b6148be88828601614759565b6060830152506080830135828111156148d657600080fd5b6148e288828601614156565b60808301525060a083013560a082015260c08301358281111561490457600080fd5b61491088828601614156565b60c08301525060e083013560e0820152809450505050614932602084016147e7565b90509250929050565b600060a0828403121561494d57600080fd5b60405160a0810181811067ffffffffffffffff8211171561497057614970614010565b604052823561497e81613e7e565b8152602083013563ffffffff8116811461499757600080fd5b602082015260408301356149aa81613eb0565b604082015260608301356149bd81613eb0565b606082015260808301356149d081613eb0565b60808201529392505050565b803560ff81168114613dd757600080fd5b60008060008060008060c08789031215614a0657600080fd5b863567ffffffffffffffff80821115614a1e57600080fd5b614a2a8a838b01614366565b97506020890135915080821115614a4057600080fd5b614a4c8a838b01614366565b9650614a5a60408a016149dc565b95506060890135915080821115614a7057600080fd5b614a7c8a838b01614315565b9450614a8a60808a01613ec6565b935060a0890135915080821115614aa057600080fd5b50614aad89828a01614315565b9150509295509295509295565b600060208284031215614acc57600080fd5b8151611b5c816147d9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006000198203614b4857614b48614b06565b5060010190565b600082821015614b6157614b61614b06565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000816000190483118215151615614baf57614baf614b06565b500290565b60008219821115614bc757614bc7614b06565b500190565b600060ff821660ff84168060ff03821115614be957614be9614b06565b019392505050565b600060ff831680614c2b577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600060208284031215614ca757600080fd5b8151611b5c81613e7e565b8051613dd781613eb0565b8051613dd781613e7e565b600082601f830112614cd957600080fd5b8151614ce7614103826142ed565b818152846020838601011115614cfc57600080fd5b61424a826020830160208701613f5e565b600082601f830112614d1e57600080fd5b81516020614d2e614103836140be565b82815260059290921b84018101918181019086841115614d4d57600080fd5b8286015b8481101561414b578051614d6481613e7e565b8352918301918301614d51565b600082601f830112614d8257600080fd5b81516020614d92614103836140be565b82815260059290921b84018101918181019086841115614db157600080fd5b8286015b8481101561414b5780518352918301918301614db5565b600060208284031215614dde57600080fd5b815167ffffffffffffffff80821115614df657600080fd5b908301906101208286031215614e0b57600080fd5b614e13614069565b82518152614e2360208401614cb2565b6020820152614e3460408401614cbd565b6040820152614e4560608401614cbd565b6060820152614e5660808401614cb2565b608082015260a083015182811115614e6d57600080fd5b614e7987828601614cc8565b60a08301525060c083015182811115614e9157600080fd5b614e9d87828601614d0d565b60c08301525060e083015182811115614eb557600080fd5b614ec187828601614d71565b60e0830152506101009283015192810192909252509392505050565b600067ffffffffffffffff808316818516808303821115614f0057614f00614b06565b01949350505050565b600067ffffffffffffffff808316818103614f2657614f26614b06565b6001019392505050565b600063ffffffff808316818516808303821115614f0057614f00614b06565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614f7f8184018a614529565b90508281036080840152614f938189614529565b905060ff871660a084015282810360c0840152614fb08187613f8a565b905067ffffffffffffffff851660e0840152828103610100840152614fd58185613f8a565b9c9b505050505050505050505050565b60208152815160208201526000602083015161010080604085015261500e610120850183613f8a565b9150604085015161502a60608601826001600160a01b03169052565b506060850151601f19808685030160808701526150478483613f8a565b935060808701519150808685030160a08701526150648483614529565b935060a08701519150808685030160c08701526150818483614529565b935060c08701519150808685030160e08701525061509f8382614252565b92505060e085015181850152508091505092915050565b600082601f8301126150c757600080fd5b815160206150d7614103836140be565b82815260059290921b840181019181810190868411156150f657600080fd5b8286015b8481101561414b57805161510d81613eb0565b83529183019183016150fa565b600082601f83011261512b57600080fd5b8151602061513b614103836140be565b82815260059290921b8401810191818101908684111561515a57600080fd5b8286015b8481101561414b57805167ffffffffffffffff81111561517e5760008081fd5b61518c8986838b0101614cc8565b84525091830191830161515e565b6000602082840312156151ac57600080fd5b815167ffffffffffffffff808211156151c457600080fd5b9083019061010082860312156151d957600080fd5b6151e161403f565b8251828111156151f057600080fd5b6151fc878286016150b6565b82525060208301518281111561521157600080fd5b61521d87828601614d0d565b60208301525060408301518281111561523557600080fd5b61524187828601614d71565b60408301525060608301518281111561525957600080fd5b6152658782860161511a565b60608301525060808301518281111561527d57600080fd5b61528987828601614d71565b60808301525060a083015160a082015260c0830151828111156152ab57600080fd5b6152b787828601614d71565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b8381101561428257815167ffffffffffffffff16875295820195908201906001016152e7565b600081518084526020808501808196508360051b8101915082860160005b85811015615355578284038952615343848351613f8a565b9885019893509084019060010161532b565b5091979650505050505050565b60408152600083516101008060408501526153816101408501836152d3565b915060208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808685030160608701526153bd8483614529565b935060408801519150808685030160808701526153da8483614252565b935060608801519150808685030160a08701526153f7848361530d565b935060808801519150808685030160c08701526154148483614252565b935060a088015160e087015260c088015191508086850301838701525061543b8382614252565b60e088015161012087015286151560208701529350611b5c92505050565b60a08152600061546c60a0830188614252565b828103602084015261547e8188614252565b905085604084015282810360608401526154988186614252565b9150508260808301529695505050505050565b6000602082840312156154bd57600080fd5b5051919050565b6000815160208301517fffffffff000000000000000000000000000000000000000000000000000000008082169350600483101561550c5780818460040360031b1b83161693505b505050919050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b16604085015281606085015261554e8285018b614529565b91508382036080850152615562828a614529565b915060ff881660a085015283820360c085015261557f8288613f8a565b90861660e08501528381036101008501529050614fd58185613f8a56fea164736f6c634300080f000a",
}

var EVM2EVMFreeOffRampABI = EVM2EVMFreeOffRampMetaData.ABI

var EVM2EVMFreeOffRampBin = EVM2EVMFreeOffRampMetaData.Bin

func DeployEVM2EVMFreeOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, blobVerifier common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig AggregateRateLimiterInterfaceRateLimiterConfig, tokenLimitsAdmin common.Address) (common.Address, *types.Transaction, *EVM2EVMFreeOffRamp, error) {
	parsed, err := EVM2EVMFreeOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMFreeOffRampBin), backend, sourceChainId, chainId, offRampConfig, blobVerifier, afn, sourceTokens, pools, rateLimiterConfig, tokenLimitsAdmin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMFreeOffRamp{EVM2EVMFreeOffRampCaller: EVM2EVMFreeOffRampCaller{contract: contract}, EVM2EVMFreeOffRampTransactor: EVM2EVMFreeOffRampTransactor{contract: contract}, EVM2EVMFreeOffRampFilterer: EVM2EVMFreeOffRampFilterer{contract: contract}}, nil
}

type EVM2EVMFreeOffRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMFreeOffRampCaller
	EVM2EVMFreeOffRampTransactor
	EVM2EVMFreeOffRampFilterer
}

type EVM2EVMFreeOffRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMFreeOffRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMFreeOffRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMFreeOffRampSession struct {
	Contract     *EVM2EVMFreeOffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMFreeOffRampCallerSession struct {
	Contract *EVM2EVMFreeOffRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMFreeOffRampTransactorSession struct {
	Contract     *EVM2EVMFreeOffRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMFreeOffRampRaw struct {
	Contract *EVM2EVMFreeOffRamp
}

type EVM2EVMFreeOffRampCallerRaw struct {
	Contract *EVM2EVMFreeOffRampCaller
}

type EVM2EVMFreeOffRampTransactorRaw struct {
	Contract *EVM2EVMFreeOffRampTransactor
}

func NewEVM2EVMFreeOffRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMFreeOffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMFreeOffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMFreeOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRamp{address: address, abi: abi, EVM2EVMFreeOffRampCaller: EVM2EVMFreeOffRampCaller{contract: contract}, EVM2EVMFreeOffRampTransactor: EVM2EVMFreeOffRampTransactor{contract: contract}, EVM2EVMFreeOffRampFilterer: EVM2EVMFreeOffRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMFreeOffRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMFreeOffRampCaller, error) {
	contract, err := bindEVM2EVMFreeOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampCaller{contract: contract}, nil
}

func NewEVM2EVMFreeOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMFreeOffRampTransactor, error) {
	contract, err := bindEVM2EVMFreeOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampTransactor{contract: contract}, nil
}

func NewEVM2EVMFreeOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMFreeOffRampFilterer, error) {
	contract, err := bindEVM2EVMFreeOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampFilterer{contract: contract}, nil
}

func bindEVM2EVMFreeOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMFreeOffRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMFreeOffRamp.Contract.EVM2EVMFreeOffRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.EVM2EVMFreeOffRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.EVM2EVMFreeOffRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMFreeOffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(AggregateRateLimiterInterfaceTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(AggregateRateLimiterInterfaceTokenBucket)).(*AggregateRateLimiterInterfaceTokenBucket)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMFreeOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMFreeOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMFreeOffRamp.Contract.CcipReceive(&_EVM2EVMFreeOffRamp.CallOpts, arg0)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMFreeOffRamp.Contract.CcipReceive(&_EVM2EVMFreeOffRamp.CallOpts, arg0)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetAFN() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetAFN(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetAFN(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetBlobVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getBlobVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetBlobVerifier() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetBlobVerifier(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetBlobVerifier() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetBlobVerifier(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

	error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getChainIDs")

	outstruct := new(GetChainIDs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMFreeOffRamp.Contract.GetChainIDs(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMFreeOffRamp.Contract.GetChainIDs(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOffRampInterfaceOffRampConfig)).(*BaseOffRampInterfaceOffRampConfig)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetConfig(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetConfig(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetDestinationToken(&_EVM2EVMFreeOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetDestinationToken(&_EVM2EVMFreeOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetDestinationTokens(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetDestinationTokens(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetExecutionState(&_EVM2EVMFreeOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetExecutionState(&_EVM2EVMFreeOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetNonce(opts *bind.CallOpts, receiver common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getNonce", receiver)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetNonce(receiver common.Address) (uint64, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetNonce(&_EVM2EVMFreeOffRamp.CallOpts, receiver)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetNonce(receiver common.Address) (uint64, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetNonce(&_EVM2EVMFreeOffRamp.CallOpts, receiver)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetPool(&_EVM2EVMFreeOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetPool(&_EVM2EVMFreeOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetPoolTokens(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetPoolTokens(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetPricesForTokens(&_EVM2EVMFreeOffRamp.CallOpts, tokens)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetPricesForTokens(&_EVM2EVMFreeOffRamp.CallOpts, tokens)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetRouter() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetRouter(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetRouter(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMFreeOffRamp.Contract.IsAFNHealthy(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMFreeOffRamp.Contract.IsAFNHealthy(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMFreeOffRamp.Contract.LatestConfigDetails(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMFreeOffRamp.Contract.LatestConfigDetails(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMFreeOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMFreeOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) Owner() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.Owner(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.Owner(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) Paused() (bool, error) {
	return _EVM2EVMFreeOffRamp.Contract.Paused(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMFreeOffRamp.Contract.Paused(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.Transmitters(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.Transmitters(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMFreeOffRamp.Contract.TypeAndVersion(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMFreeOffRamp.Contract.TypeAndVersion(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.AcceptOwnership(&_EVM2EVMFreeOffRamp.TransactOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.AcceptOwnership(&_EVM2EVMFreeOffRamp.TransactOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.AddPool(&_EVM2EVMFreeOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.AddPool(&_EVM2EVMFreeOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) Execute(opts *bind.TransactOpts, report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "execute", report, manualExecution)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.Execute(&_EVM2EVMFreeOffRamp.TransactOpts, report, manualExecution)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.Execute(&_EVM2EVMFreeOffRamp.TransactOpts, report, manualExecution)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "executeSingleMessage", message)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMFreeOffRamp.TransactOpts, message)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMFreeOffRamp.TransactOpts, message)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.Pause(&_EVM2EVMFreeOffRamp.TransactOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.Pause(&_EVM2EVMFreeOffRamp.TransactOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.RemovePool(&_EVM2EVMFreeOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.RemovePool(&_EVM2EVMFreeOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetAFN(&_EVM2EVMFreeOffRamp.TransactOpts, afn)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetAFN(&_EVM2EVMFreeOffRamp.TransactOpts, afn)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetBlobVerifier(opts *bind.TransactOpts, blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setBlobVerifier", blobVerifier)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetBlobVerifier(&_EVM2EVMFreeOffRamp.TransactOpts, blobVerifier)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetBlobVerifier(&_EVM2EVMFreeOffRamp.TransactOpts, blobVerifier)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetConfig(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setConfig", config)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetConfig(&_EVM2EVMFreeOffRamp.TransactOpts, config)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetConfig(&_EVM2EVMFreeOffRamp.TransactOpts, config)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetConfig0(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setConfig0", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetConfig0(&_EVM2EVMFreeOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetConfig0(&_EVM2EVMFreeOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetPrices(&_EVM2EVMFreeOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetPrices(&_EVM2EVMFreeOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMFreeOffRamp.TransactOpts, config)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMFreeOffRamp.TransactOpts, config)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetRouter(&_EVM2EVMFreeOffRamp.TransactOpts, router)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetRouter(&_EVM2EVMFreeOffRamp.TransactOpts, router)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMFreeOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMFreeOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.TransferOwnership(&_EVM2EVMFreeOffRamp.TransactOpts, to)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.TransferOwnership(&_EVM2EVMFreeOffRamp.TransactOpts, to)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.Transmit(&_EVM2EVMFreeOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.Transmit(&_EVM2EVMFreeOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.Unpause(&_EVM2EVMFreeOffRamp.TransactOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.Unpause(&_EVM2EVMFreeOffRamp.TransactOpts)
}

type EVM2EVMFreeOffRampAFNSetIterator struct {
	Event *EVM2EVMFreeOffRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampAFNSet)
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
		it.Event = new(EVM2EVMFreeOffRampAFNSet)
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

func (it *EVM2EVMFreeOffRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampAFNSetIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampAFNSet)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseAFNSet(log types.Log) (*EVM2EVMFreeOffRampAFNSet, error) {
	event := new(EVM2EVMFreeOffRampAFNSet)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampConfigChangedIterator struct {
	Event *EVM2EVMFreeOffRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampConfigChanged)
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
		it.Event = new(EVM2EVMFreeOffRampConfigChanged)
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

func (it *EVM2EVMFreeOffRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampConfigChangedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampConfigChanged)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMFreeOffRampConfigChanged, error) {
	event := new(EVM2EVMFreeOffRampConfigChanged)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampConfigSetIterator struct {
	Event *EVM2EVMFreeOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampConfigSet)
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
		it.Event = new(EVM2EVMFreeOffRampConfigSet)
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

func (it *EVM2EVMFreeOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampConfigSet struct {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampConfigSetIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampConfigSet)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseConfigSet(log types.Log) (*EVM2EVMFreeOffRampConfigSet, error) {
	event := new(EVM2EVMFreeOffRampConfigSet)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampExecutionStateChangedIterator struct {
	Event *EVM2EVMFreeOffRampExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampExecutionStateChanged)
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
		it.Event = new(EVM2EVMFreeOffRampExecutionStateChanged)
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

func (it *EVM2EVMFreeOffRampExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampExecutionStateChanged struct {
	SequenceNumber uint64
	State          uint8
	Raw            types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMFreeOffRampExecutionStateChangedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampExecutionStateChangedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampExecutionStateChanged)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseExecutionStateChanged(log types.Log) (*EVM2EVMFreeOffRampExecutionStateChanged, error) {
	event := new(EVM2EVMFreeOffRampExecutionStateChanged)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampOffRampConfigSetIterator struct {
	Event *EVM2EVMFreeOffRampOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampOffRampConfigSet)
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
		it.Event = new(EVM2EVMFreeOffRampOffRampConfigSet)
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

func (it *EVM2EVMFreeOffRampOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampOffRampConfigSet struct {
	Config BaseOffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampOffRampConfigSetIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampOffRampConfigSet)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseOffRampConfigSet(log types.Log) (*EVM2EVMFreeOffRampOffRampConfigSet, error) {
	event := new(EVM2EVMFreeOffRampOffRampConfigSet)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampOffRampRouterSetIterator struct {
	Event *EVM2EVMFreeOffRampOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampOffRampRouterSet)
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
		it.Event = new(EVM2EVMFreeOffRampOffRampRouterSet)
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

func (it *EVM2EVMFreeOffRampOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampOffRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMFreeOffRampOffRampRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampOffRampRouterSetIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampOffRampRouterSet)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseOffRampRouterSet(log types.Log) (*EVM2EVMFreeOffRampOffRampRouterSet, error) {
	event := new(EVM2EVMFreeOffRampOffRampRouterSet)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMFreeOffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMFreeOffRampOwnershipTransferRequested)
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

func (it *EVM2EVMFreeOffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMFreeOffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampOwnershipTransferRequestedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampOwnershipTransferRequested)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMFreeOffRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMFreeOffRampOwnershipTransferRequested)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampOwnershipTransferredIterator struct {
	Event *EVM2EVMFreeOffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMFreeOffRampOwnershipTransferred)
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

func (it *EVM2EVMFreeOffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMFreeOffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampOwnershipTransferredIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampOwnershipTransferred)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMFreeOffRampOwnershipTransferred, error) {
	event := new(EVM2EVMFreeOffRampOwnershipTransferred)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampPausedIterator struct {
	Event *EVM2EVMFreeOffRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampPaused)
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
		it.Event = new(EVM2EVMFreeOffRampPaused)
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

func (it *EVM2EVMFreeOffRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampPausedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampPaused)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParsePaused(log types.Log) (*EVM2EVMFreeOffRampPaused, error) {
	event := new(EVM2EVMFreeOffRampPaused)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampPoolAddedIterator struct {
	Event *EVM2EVMFreeOffRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampPoolAdded)
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
		it.Event = new(EVM2EVMFreeOffRampPoolAdded)
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

func (it *EVM2EVMFreeOffRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampPoolAddedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampPoolAdded)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMFreeOffRampPoolAdded, error) {
	event := new(EVM2EVMFreeOffRampPoolAdded)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampPoolRemovedIterator struct {
	Event *EVM2EVMFreeOffRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampPoolRemoved)
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
		it.Event = new(EVM2EVMFreeOffRampPoolRemoved)
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

func (it *EVM2EVMFreeOffRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampPoolRemovedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampPoolRemoved)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMFreeOffRampPoolRemoved, error) {
	event := new(EVM2EVMFreeOffRampPoolRemoved)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampTokenPriceChangedIterator struct {
	Event *EVM2EVMFreeOffRampTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampTokenPriceChanged)
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
		it.Event = new(EVM2EVMFreeOffRampTokenPriceChanged)
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

func (it *EVM2EVMFreeOffRampTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampTokenPriceChangedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampTokenPriceChanged)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMFreeOffRampTokenPriceChanged, error) {
	event := new(EVM2EVMFreeOffRampTokenPriceChanged)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMFreeOffRampTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampTokensRemovedFromBucket)
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
		it.Event = new(EVM2EVMFreeOffRampTokensRemovedFromBucket)
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

func (it *EVM2EVMFreeOffRampTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampTokensRemovedFromBucketIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampTokensRemovedFromBucket)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMFreeOffRampTokensRemovedFromBucket, error) {
	event := new(EVM2EVMFreeOffRampTokensRemovedFromBucket)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampTransmittedIterator struct {
	Event *EVM2EVMFreeOffRampTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampTransmitted)
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
		it.Event = new(EVM2EVMFreeOffRampTransmitted)
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

func (it *EVM2EVMFreeOffRampTransmittedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampTransmittedIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampTransmittedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampTransmitted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampTransmitted)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseTransmitted(log types.Log) (*EVM2EVMFreeOffRampTransmitted, error) {
	event := new(EVM2EVMFreeOffRampTransmitted)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampUnpausedIterator struct {
	Event *EVM2EVMFreeOffRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampUnpaused)
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
		it.Event = new(EVM2EVMFreeOffRampUnpaused)
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

func (it *EVM2EVMFreeOffRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampUnpausedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampUnpaused)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMFreeOffRampUnpaused, error) {
	event := new(EVM2EVMFreeOffRampUnpaused)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMFreeOffRamp.abi.Events["AFNSet"].ID:
		return _EVM2EVMFreeOffRamp.ParseAFNSet(log)
	case _EVM2EVMFreeOffRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMFreeOffRamp.ParseConfigChanged(log)
	case _EVM2EVMFreeOffRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMFreeOffRamp.ParseConfigSet(log)
	case _EVM2EVMFreeOffRamp.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMFreeOffRamp.ParseExecutionStateChanged(log)
	case _EVM2EVMFreeOffRamp.abi.Events["OffRampConfigSet"].ID:
		return _EVM2EVMFreeOffRamp.ParseOffRampConfigSet(log)
	case _EVM2EVMFreeOffRamp.abi.Events["OffRampRouterSet"].ID:
		return _EVM2EVMFreeOffRamp.ParseOffRampRouterSet(log)
	case _EVM2EVMFreeOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMFreeOffRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMFreeOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMFreeOffRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMFreeOffRamp.abi.Events["Paused"].ID:
		return _EVM2EVMFreeOffRamp.ParsePaused(log)
	case _EVM2EVMFreeOffRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMFreeOffRamp.ParsePoolAdded(log)
	case _EVM2EVMFreeOffRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMFreeOffRamp.ParsePoolRemoved(log)
	case _EVM2EVMFreeOffRamp.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMFreeOffRamp.ParseTokenPriceChanged(log)
	case _EVM2EVMFreeOffRamp.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMFreeOffRamp.ParseTokensRemovedFromBucket(log)
	case _EVM2EVMFreeOffRamp.abi.Events["Transmitted"].ID:
		return _EVM2EVMFreeOffRamp.ParseTransmitted(log)
	case _EVM2EVMFreeOffRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMFreeOffRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMFreeOffRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMFreeOffRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMFreeOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMFreeOffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf")
}

func (EVM2EVMFreeOffRampOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa0")
}

func (EVM2EVMFreeOffRampOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4")
}

func (EVM2EVMFreeOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMFreeOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMFreeOffRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMFreeOffRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMFreeOffRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMFreeOffRampTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMFreeOffRampTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMFreeOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (EVM2EVMFreeOffRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRamp) Address() common.Address {
	return _EVM2EVMFreeOffRamp.address
}

type EVM2EVMFreeOffRampInterface interface {
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error)

	CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetBlobVerifier(opts *bind.CallOpts) (common.Address, error)

	GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

		error)

	GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetNonce(opts *bind.CallOpts, receiver common.Address) (uint64, error)

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

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMFreeOffRampAFNSet, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMFreeOffRampConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMFreeOffRampConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMFreeOffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMFreeOffRampExecutionStateChanged, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*EVM2EVMFreeOffRampOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMFreeOffRampOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*EVM2EVMFreeOffRampOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMFreeOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMFreeOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMFreeOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMFreeOffRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMFreeOffRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMFreeOffRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMFreeOffRampPoolRemoved, error)

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMFreeOffRampTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMFreeOffRampTokensRemovedFromBucket, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMFreeOffRampTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMFreeOffRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
