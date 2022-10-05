// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package any_2_evm_subscription_offramp

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

var EVM2EVMSubscriptionOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SubscriptionNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlobVerifier\",\"outputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"}],\"name\":\"setBlobVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005fd938038062005fd983398101604081905262000034916200063c565b6000805460ff191681556001908a908a908a908a908a908a908a908a908a90829082908690869089903390819081620000b45760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ee57620000ee8162000385565b5050506001600160a01b0381166200011957604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200015b5760405162d8548360e71b815260040160405180910390fd5b81516200017090600490602085019062000436565b5060005b82518110156200023b57600082828151811062000195576200019562000796565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001df57620001df62000796565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b0292909116919091179055506200023381620007ac565b905062000174565b5050600580546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600855600993909355600a55600b91909155875116620002b9576040516342bcdf7f60e11b815260040160405180910390fd5b5050506080958652505060a0929092528051600e8054602084015160408501516001600160401b03908116600160c01b026001600160c01b0363ffffffff909316600160a01b026001600160c01b03199094166001600160a01b03968716179390931791909116919091179091556060830151600f80549490960151821668010000000000000000026001600160801b031990941691161791909117909255600d8054919092166001600160a01b031991909116179055151560c05250620007d4975050505050505050565b336001600160a01b03821603620003df5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ab565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200048e579160200282015b828111156200048e57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000457565b506200049c929150620004a0565b5090565b5b808211156200049c5760008155600101620004a1565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715620004f257620004f2620004b7565b60405290565b6001600160a01b03811681146200050e57600080fd5b50565b80516200051e81620004f8565b919050565b80516001600160401b03811681146200051e57600080fd5b600082601f8301126200054d57600080fd5b815160206001600160401b03808311156200056c576200056c620004b7565b8260051b604051601f19603f83011681018181108482111715620005945762000594620004b7565b604052938452858101830193838101925087851115620005b357600080fd5b83870191505b84821015620005df578151620005cf81620004f8565b83529183019190830190620005b9565b979650505050505050565b600060408284031215620005fd57600080fd5b604080519081016001600160401b0381118282101715620006225762000622620004b7565b604052825181526020928301519281019290925250919050565b6000806000806000806000806000898b036101c08112156200065d57600080fd5b8a51995060208b0151985060a0603f19820112156200067b57600080fd5b5062000686620004cd565b60408b01516200069681620004f8565b815260608b015163ffffffff81168114620006b057600080fd5b6020820152620006c360808c0162000523565b6040820152620006d660a08c0162000523565b6060820152620006e960c08c0162000523565b60808201529650620006fe60e08b0162000511565b95506200070f6101008b0162000511565b6101208b01519095506001600160401b03808211156200072e57600080fd5b6200073c8d838e016200053b565b95506101408c01519150808211156200075457600080fd5b50620007638c828d016200053b565b935050620007768b6101608c01620005ea565b9150620007876101a08b0162000511565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b600060018201620007cd57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c0516157c06200081960003960006115b80152600081816102db0152611e120152600081816102b801528181611ded01526138dd01526157c06000f3fe608060405234801561001057600080fd5b50600436106102ad5760003560e01c8063814118341161017b578063b1dc65a4116100d8578063c3f909d41161008c578063e3d0e71211610071578063e3d0e7121461078a578063eb511dd41461079d578063f2fde38b146107b057600080fd5b8063c3f909d41461065f578063c9029f6a1461077757600080fd5b8063bbe4f6db116100bd578063bbe4f6db1461060d578063be9b03f114610639578063c0d786551461064c57600080fd5b8063b1dc65a4146105e7578063b4069b31146105fa57600080fd5b806390c2339b1161012f578063a639d1c011610114578063a639d1c0146105a3578063afcb95d7146105b6578063b0f479a1146105d657600080fd5b806390c2339b14610555578063918725431461059057600080fd5b80638456cb59116101605780638456cb591461052f57806389c06568146105375780638da5cb5b1461053f57600080fd5b806381411834146104f757806381ff7048146104ff57600080fd5b80633f4ba83a116102295780636133dc24116101dd578063694ec2b1116101c2578063694ec2b1146104c9578063744b92e2146104dc57806379ba5097146104ef57600080fd5b80636133dc24146104a3578063681fba16146104b457600080fd5b80634741062e1161020e5780634741062e14610467578063599f6431146104875780635c975abb1461049857600080fd5b80633f4ba83a1461044c5780634352fa9f1461045457600080fd5b8063181f5a77116102805780632d0335ab116102655780632d0335ab146103df578063351f0faf1461042b57806339aa92641461043957600080fd5b8063181f5a77146103715780632222dd42146103ba57600080fd5b8063087ae6df146102b2578063108ee5fc1461030a578063142a98fc1461031f578063147809b314610359575b600080fd5b604080517f000000000000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006020820152015b60405180910390f35b61031d610318366004614083565b6107c3565b005b61034c61032d3660046140c1565b67ffffffffffffffff1660009081526010602052604090205460ff1690565b604051610301919061410d565b610361610879565b6040519015158152602001610301565b6103ad6040518060400160405280602081526020017f45564d3245564d537562736372697074696f6e4f666652616d7020312e302e3081525081565b60405161030191906141a6565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610301565b6104126103ed366004614083565b6001600160a01b031660009081526017602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff9091168152602001610301565b61031d6102ad3660046141c4565b61031d610447366004614083565b610906565b61031d61093d565b61031d6104623660046143a1565b61094f565b61047a610475366004614405565b610ba4565b604051610301919061447d565b6005546001600160a01b03166103c7565b60005460ff16610361565b600d546001600160a01b03166103c7565b6104bc610c6c565b60405161030191906144c9565b61031d6104d7366004614555565b610d31565b61031d6104ea36600461467b565b610daa565b61031d6110d0565b6104bc6111b8565b6013546011546040805163ffffffff80851682526401000000009094049093166020840152820152606001610301565b61031d61121a565b6104bc61122a565b60005461010090046001600160a01b03166103c7565b61055d61128a565b60405161030191908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b61031d61059e3660046146b4565b61132b565b61031d6105b1366004614083565b61145e565b604080516001815260006020820181905291810191909152606001610301565b600c546001600160a01b03166103c7565b61031d6105f536600461474f565b611495565b6103c7610608366004614083565b611a68565b6103c761061b366004614083565b6001600160a01b039081166000908152600360205260409020541690565b61031d610647366004614931565b611b56565b61031d61065a366004614083565b612471565b61071b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a081018252600e546001600160a01b038116825263ffffffff74010000000000000000000000000000000000000000820416602083015267ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811692820192909252600f548083166060830152680100000000000000009004909116608082015290565b6040805182516001600160a01b0316815260208084015163ffffffff16908201528282015167ffffffffffffffff908116928201929092526060808401518316908201526080928301519091169181019190915260a001610301565b61031d610785366004614a7a565b6124d0565b61031d610798366004614b2c565b61266f565b61031d6107ab36600461467b565b612e8c565b61031d6107be366004614083565b613065565b6107cb613076565b6001600160a01b03811661080b576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28910160405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa1580156108dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109009190614bf9565b15905090565b61090e613076565b6005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b610945613076565b61094d6130d5565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561098c57506005546001600160a01b03163314155b156109c3576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151815181146109ff576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610a59576006600060078381548110610a2457610a24614c16565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610a5281614c74565b9050610a05565b5060005b82811015610b89576000858281518110610a7957610a79614c16565b6020026020010151905060006001600160a01b0316816001600160a01b031603610acf576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610ae157610ae1614c16565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610b4657610b46614c16565b6020026020010151604051610b709291906001600160a01b03929092168252602082015260400190565b60405180910390a150610b8281614c74565b9050610a5d565b508351610b9d906007906020870190613fcc565b5050505050565b80516060908067ffffffffffffffff811115610bc257610bc2614200565b604051908082528060200260200182016040528015610beb578160200160208202803683370190505b50915060005b81811015610c655760066000858381518110610c0f57610c0f614c16565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610c4a57610c4a614c16565b6020908102919091010152610c5e81614c74565b9050610bf1565b5050919050565b60045460609067ffffffffffffffff811115610c8a57610c8a614200565b604051908082528060200260200182016040528015610cb3578160200160208202803683370190505b50905060005b600454811015610d2d57610cf360048281548110610cd957610cd9614c16565b6000918252602090912001546001600160a01b0316611a68565b828281518110610d0557610d05614c16565b6001600160a01b0390921660209283029190910190910152610d2681614c74565b9050610cb9565b5090565b333014610d6a576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60808101515115610d9e57610d8781608001518260c00151613171565b610d9e8160a001518260c001518360400151613304565b610da78161339e565b50565b610db2613076565b6004546000819003610df0576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610e7e576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614610ecd576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004610edc600185614c8e565b81548110610eec57610eec614c16565b9060005260206000200160009054906101000a90046001600160a01b03169050600482602001516bffffffffffffffffffffffff1681548110610f3157610f31614c16565b6000918252602090912001546001600160a01b03166004610f53600186614c8e565b81548110610f6357610f63614c16565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600483602001516bffffffffffffffffffffffff1681548110610fb757610fb7614c16565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff909216740100000000000000000000000000000000000000000291909216179055600480548061104157611041614ca5565b600082815260208082206000199084018101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b03878116808452600383526040808520949094558351908152908716918101919091527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b0316331461112f5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060601680548060200260200160405190810160405280929190818152602001828054801561121057602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116111f2575b5050505050905090565b611222613076565b61094d613476565b60606004805480602002602001604051908101604052809291908181526020018280548015611210576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116111f2575050505050905090565b6112b56040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b546060820181905242906000906112ef9083614c8e565b6020840151845191925061131b916113079084614cd4565b85604001516113169190614cf3565b6134fe565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561136857506005546001600160a01b03163314155b1561139f576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff116113f3576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113fd6008613514565b602081015160098190558151600855600a5461141991906134fe565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b611466613076565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916114eb91849163ffffffff851691908e908e90819084018382808284376000920191909152506135c192505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260115480825260125460ff808216602085015261010090910416928201929092529083146115a65760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401611126565b6115b48b8b8b8b8b8b61363a565b60007f000000000000000000000000000000000000000000000000000000000000000015611611576002826020015183604001516115f29190614d0b565b6115fc9190614d5f565b611607906001614d0b565b60ff169050611627565b6020820151611621906001614d0b565b60ff1690505b8881146116765760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401611126565b8887146116c55760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401611126565b3360009081526014602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115611708576117086140de565b6002811115611719576117196140de565b9052509050600281602001516002811115611736576117366140de565b14801561177057506016816000015160ff168154811061175857611758614c16565b6000918252602090912001546001600160a01b031633145b6117bc5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401611126565b5050505050600088886040516117d3929190614d81565b6040519081900381206117ea918c90602001614d91565b60405160208183030381529060405280519060200120905061180a61403a565b604080518082019091526000808252602082015260005b88811015611a4657600060018588846020811061184057611840614c16565b61184d91901a601b614d0b565b8d8d8681811061185f5761185f614c16565b905060200201358c8c8781811061187857611878614c16565b90506020020135604051600081526020016040526040516118b5949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156118d7573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526014602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561192c5761192c6140de565b600281111561193d5761193d6140de565b905250925060018360200151600281111561195a5761195a6140de565b146119a75760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401611126565b8251849060ff16601f81106119be576119be614c16565b602002015115611a105760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401611126565b600184846000015160ff16601f8110611a2b57611a2b614c16565b9115156020909202015250611a3f81614c74565b9050611821565b5050505063ffffffff8110611a5d57611a5d614dad565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611abc576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611b2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b4f9190614ddc565b9392505050565b60005460ff1615611ba95760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611126565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611bfc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c209190614bf9565b15611c56576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c546001600160a01b031680611c99576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060830151516000819003611cda576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115611cf557611cf5614200565b604051908082528060200260200182016040528015611d6b57816020015b60408051610120810182526000808252602080830182905292820181905260608083018290526080830182905260a0830181905260c0830181905260e08301526101008201528252600019909201910181611d135790505b50905060008267ffffffffffffffff811115611d8957611d89614200565b604051908082528060200260200182016040528015611db2578160200160208202803683370190505b5090506000611e637f3997e2cfd3ccacf768662bd35c3dbf323724407d75aae3019c04f4aa59b1193f600e54604080516020808201949094527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201526001600160a01b039092166080808401919091528151808403909101815260a09092019052805191012090565b905060005b84811015611efd57600088606001518281518110611e8857611e88614c16565b6020026020010151806020019051810190611ea39190614f13565b9050611eaf81846136ce565b848381518110611ec157611ec1614c16565b60200260200101818152505080858381518110611ee057611ee0614c16565b60200260200101819052505080611ef690614c74565b9050611e68565b50600080611f1e848a608001518b60a001518c60c001518d60e001516137e1565b915091506000855182611f319190615024565b9050888015611f685750600e5474010000000000000000000000000000000000000000900463ffffffff16611f668442614c8e565b105b15611f9f576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008a60400151600081518110611fb857611fb8614c16565b6020026020010151905060005b888110156124635760005a90506000898381518110611fe657611fe6614c16565b602002602001015190506000612019826020015167ffffffffffffffff1660009081526010602052604090205460ff1690565b9050600281600381111561202f5761202f6140de565b036120785760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611126565b60608201516040517f0cbebc240000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526000918f1690630cbebc2490602401600060405180830381865afa1580156120df573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526121079190810190615038565b60208101519091506001600160a01b031661215f5760608301516040517f8515736a0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401611126565b608083015160608401516001600160a01b0316600090815260176020526040812054909167ffffffffffffffff9081169161219c911660016150f4565b67ffffffffffffffff1614905080806121d3575081604001511580156121d3575060038360038111156121d1576121d16140de565b145b61221b5760808401516040517fb0241f4a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611126565b612224846138d9565b60208481015167ffffffffffffffff166000908152601090915260408120805460ff1916600117905561225e61225986613a20565b613c9f565b60208087015167ffffffffffffffff1660009081526010909152604090208054919250829160ff1916600183600381111561229b5761229b6140de565b0217905550846020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf826040516122de919061410d565b60405180910390a281801561231057506002816003811115612302576123026140de565b148061231057508260400151155b1561236f5760608501516001600160a01b03166000908152601760205260408120805467ffffffffffffffff169161234783615120565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b50508e61244e578d6001600160a01b031663294d266184606001518560400151670de0b6b3a76400008a3a8d5a6123a6908d614c8e565b6123b09190614cf3565b6123ba9190614cd4565b6123c49190614cd4565b6123ce9190615024565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1681526001600160a01b0393841660048201529290911660248301526044820152606401600060405180830381600087803b15801561243557600080fd5b505af1158015612449573d6000803e3d6000fd5b505050505b505050508061245c90614c74565b9050611fc5565b505050505050505050505050565b612479613076565b600c805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b6124d8613076565b80516001600160a01b0316612519576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600e805460208085018051604080880180516001600160a01b039098167fffffffffffffffff00000000000000000000000000000000000000000000000090961686177401000000000000000000000000000000000000000063ffffffff948516021777ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff998a16021790965560608089018051600f80546080808e018051948e167fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169290921768010000000000000000948e1694909402939093179091558451988952955190941695870195909552955187169085015291518516938301939093529151909216908201527fd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa09060a001611453565b855185518560ff16601f8311156126c85760405162461bcd60e51b815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401611126565b600081116127185760405162461bcd60e51b815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401611126565b81831461278c5760405162461bcd60e51b8152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401611126565b612797816003614cd4565b83116127e55760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401611126565b6127ed613076565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b601554156129745760155460009061284590600190614c8e565b905060006015828154811061285c5761285c614c16565b6000918252602082200154601680546001600160a01b039092169350908490811061288957612889614c16565b60009182526020808320909101546001600160a01b0385811684526014909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000908116909155929091168084529220805490911690556015805491925090806128fc576128fc614ca5565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff19169055019055601680548061293c5761293c614ca5565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff191690550190555061282b915050565b60005b815151811015612d05576000601460008460000151848151811061299d5761299d614c16565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff1660028111156129da576129da6140de565b14612a275760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401611126565b6040805180820190915260ff82168152600160208201528251805160149160009185908110612a5857612a58614c16565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612ace57612ace6140de565b021790555060009150612ade9050565b6014600084602001518481518110612af857612af8614c16565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115612b3557612b356140de565b14612b825760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401611126565b6040805180820190915260ff821681526020810160028152506014600084602001518481518110612bb557612bb5614c16565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612c2b57612c2b6140de565b021790555050825180516015925083908110612c4957612c49614c16565b6020908102919091018101518254600181018455600093845292829020909201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039093169290921790915582015180516016919083908110612cad57612cad614c16565b602090810291909101810151825460018101845560009384529190922001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909216919091179055612cfe81614c74565b9050612977565b5060408101516012805460ff191660ff909216919091179055601380547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092612d79928692908216911617615147565b92506101000a81548163ffffffff021916908363ffffffff160217905550612dd84630601360009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613dd3565b6011819055825180516012805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560135460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612e77988b98919763ffffffff909216969095919491939192615166565b60405180910390a15050505050505050505050565b612e94613076565b6001600160a01b0382161580612eb157506001600160a01b038116155b15612ee8576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612f77576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201805473ffffffffffffffffffffffffffffffffffffffff1916821790558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c91015b60405180910390a1505050565b61306d613076565b610da781613e60565b60005461010090046001600160a01b0316331461094d5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401611126565b60005460ff166131275760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401611126565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6000805b83518110156132685760006006600086848151811061319657613196614c16565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600003613225578482815181106131db576131db614c16565b60200260200101516040517f9a655f7b00000000000000000000000000000000000000000000000000000000815260040161112691906001600160a01b0391909116815260200190565b83828151811061323757613237614c16565b60200260200101518161324a9190614cd4565b6132549084614cf3565b9250508061326190614c74565b9050613175565b5080156132ff576132796008613514565b600a548111156132b5576040517f3bfa6f3800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600860020160008282546132ca9190614c8e565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba37890602001613058565b505050565b815183511461333f576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b83518110156133985761338884828151811061336057613360614c16565b602002602001015184838151811061337a5761337a614c16565b602002602001015184613f1c565b61339181614c74565b9050613342565b50505050565b60408101516001600160a01b03163b6133b45750565b600c546040517f5c9fa11e0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690635c9fa11e906133fd9084906004016151fc565b6020604051808303816000875af115801561341c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134409190614bf9565b610da7576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005460ff16156134c95760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611126565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586131543390565b600081831061350d5781611b4f565b5090919050565b600181015460028201544291148061352f5750808260030154145b15613538575050565b81600101548260020154111561357a576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082600301548261358c9190614c8e565b600184015484549192506135b3916135a49084614cd4565b85600201546113169190614cf3565b600284015550600390910155565b306001600160a01b031663be9b03f1828060200190518101906135e491906153b1565b60006040518363ffffffff1660e01b8152600401613603929190615579565b600060405180830381600087803b15801561361d57600080fd5b505af1158015613631573d6000803e3d6000fd5b50505050505050565b6000613647826020614cd4565b613652856020614cd4565b61365e88610144614cf3565b6136689190614cf3565b6136729190614cf3565b61367d906000614cf3565b90503681146136315760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401611126565b60008060001b828460200151856040015186606001518760a00151805190602001208860c0015160405160200161370591906144c9565b604051602081830303815290604052805190602001208960e00151604051602001613730919061447d565b604051602081830303815290604052805190602001208a61010001518b608001516040516020016137c39a99989796959493929190998a5260208a019890985267ffffffffffffffff96871660408a01526001600160a01b0395861660608a015293909416608088015260a087019190915260c086015260e0850191909152610100840152166101208201526101400190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600d546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce9061383e908c908c908c908c908c90600401615670565b6020604051808303816000875af115801561385d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061388191906156c2565b9050600081116138bd576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a6138c99084614c8e565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146139395780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401611126565b600f5460c0820151516801000000000000000090910467ffffffffffffffff16108061396f57508060e00151518160c001515114155b156139b85760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611126565b600f5460a08201515167ffffffffffffffff9091161015610da757600f5460a0820151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301526024820152604401611126565b613a71604051806101000160405280600081526020016060815260200160006001600160a01b0316815260200160608152602001606081526020016060815260200160608152602001600081525090565b60c08201515160008167ffffffffffffffff811115613a9257613a92614200565b604051908082528060200260200182016040528015613abb578160200160208202803683370190505b50905060008267ffffffffffffffff811115613ad957613ad9614200565b604051908082528060200260200182016040528015613b02578160200160208202803683370190505b50905060005b83811015613c0a576000613b388760c001518381518110613b2b57613b2b614c16565b6020026020010151613f6a565b905080838381518110613b4d57613b4d614c16565b60200260200101906001600160a01b031690816001600160a01b031681525050806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613bab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613bcf9190614ddc565b848381518110613be157613be1614c16565b6001600160a01b039092166020928302919091019091015250613c0381614c74565b9050613b08565b50604051806101000160405280866000015181526020018660400151604051602001613c4591906001600160a01b0391909116815260200190565b604051602081830303815290604052815260200186606001516001600160a01b031681526020018660a0015181526020018381526020018281526020018660e0015181526020018661010001518152509350505050919050565b6040517f694ec2b1000000000000000000000000000000000000000000000000000000008152600090309063694ec2b190613cde9085906004016151fc565b600060405180830381600087803b158015613cf857600080fd5b505af1925050508015613d09575060015b613dcb573d808015613d37576040519150601f19603f3d011682016040523d82523d6000602084013e613d3c565b606091505b50613d46816156db565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da80000000000000000000000000000000000000000000000000000000003613d975750600392915050565b806040517fcf19edfd00000000000000000000000000000000000000000000000000000000815260040161112691906141a6565b506002919050565b6000808a8a8a8a8a8a8a8a8a604051602001613df79998979695949392919061572b565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b03821603613eb85760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401611126565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401613603565b6001600160a01b038181166000908152600360205260409020541680613fc7576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401611126565b919050565b82805482825590600052602060002090810192821561402e579160200282015b8281111561402e578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190613fec565b50610d2d929150614059565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610d2d576000815560010161405a565b6001600160a01b0381168114610da757600080fd5b60006020828403121561409557600080fd5b8135611b4f8161406e565b67ffffffffffffffff81168114610da757600080fd5b8035613fc7816140a0565b6000602082840312156140d357600080fd5b8135611b4f816140a0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310614148577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b83811015614169578181015183820152602001614151565b838111156133985750506000910152565b6000815180845261419281602086016020860161414e565b601f01601f19169290920160200192915050565b602081526000611b4f602083018461417a565b8035613fc78161406e565b6000602082840312156141d657600080fd5b813567ffffffffffffffff8111156141ed57600080fd5b82016101008185031215611b4f57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff8111828210171561425357614253614200565b60405290565b604051610120810167ffffffffffffffff8111828210171561425357614253614200565b604051601f8201601f1916810167ffffffffffffffff811182821017156142a6576142a6614200565b604052919050565b600067ffffffffffffffff8211156142c8576142c8614200565b5060051b60200190565b600082601f8301126142e357600080fd5b813560206142f86142f3836142ae565b61427d565b82815260059290921b8401810191818101908684111561431757600080fd5b8286015b8481101561433b57803561432e8161406e565b835291830191830161431b565b509695505050505050565b600082601f83011261435757600080fd5b813560206143676142f3836142ae565b82815260059290921b8401810191818101908684111561438657600080fd5b8286015b8481101561433b578035835291830191830161438a565b600080604083850312156143b457600080fd5b823567ffffffffffffffff808211156143cc57600080fd5b6143d8868387016142d2565b935060208501359150808211156143ee57600080fd5b506143fb85828601614346565b9150509250929050565b60006020828403121561441757600080fd5b813567ffffffffffffffff81111561442e57600080fd5b61443a848285016142d2565b949350505050565b600081518084526020808501945080840160005b8381101561447257815187529582019590820190600101614456565b509495945050505050565b602081526000611b4f6020830184614442565b600081518084526020808501945080840160005b838110156144725781516001600160a01b0316875295820195908201906001016144a4565b602081526000611b4f6020830184614490565b600067ffffffffffffffff8211156144f6576144f6614200565b50601f01601f191660200190565b600082601f83011261451557600080fd5b81356145236142f3826144dc565b81815284602083860101111561453857600080fd5b816020850160208301376000918101602001919091529392505050565b60006020828403121561456757600080fd5b813567ffffffffffffffff8082111561457f57600080fd5b90830190610100828603121561459457600080fd5b61459c61422f565b823581526020830135828111156145b257600080fd5b6145be87828601614504565b6020830152506145d0604084016141b9565b60408201526060830135828111156145e757600080fd5b6145f387828601614504565b60608301525060808301358281111561460b57600080fd5b614617878286016142d2565b60808301525060a08301358281111561462f57600080fd5b61463b878286016142d2565b60a08301525060c08301358281111561465357600080fd5b61465f87828601614346565b60c08301525060e083013560e082015280935050505092915050565b6000806040838503121561468e57600080fd5b82356146998161406e565b915060208301356146a98161406e565b809150509250929050565b6000604082840312156146c657600080fd5b6040516040810181811067ffffffffffffffff821117156146e9576146e9614200565b604052823581526020928301359281019290925250919050565b60008083601f84011261471557600080fd5b50813567ffffffffffffffff81111561472d57600080fd5b6020830191508360208260051b850101111561474857600080fd5b9250929050565b60008060008060008060008060e0898b03121561476b57600080fd5b606089018a81111561477c57600080fd5b8998503567ffffffffffffffff8082111561479657600080fd5b818b0191508b601f8301126147aa57600080fd5b8135818111156147b957600080fd5b8c60208285010111156147cb57600080fd5b6020830199508098505060808b01359150808211156147e957600080fd5b6147f58c838d01614703565b909750955060a08b013591508082111561480e57600080fd5b5061481b8b828c01614703565b999c989b50969995989497949560c00135949350505050565b600082601f83011261484557600080fd5b813560206148556142f3836142ae565b82815260059290921b8401810191818101908684111561487457600080fd5b8286015b8481101561433b57803561488b816140a0565b8352918301918301614878565b600082601f8301126148a957600080fd5b813560206148b96142f3836142ae565b82815260059290921b840181019181810190868411156148d857600080fd5b8286015b8481101561433b57803567ffffffffffffffff8111156148fc5760008081fd5b61490a8986838b0101614504565b8452509183019183016148dc565b8015158114610da757600080fd5b8035613fc781614918565b6000806040838503121561494457600080fd5b823567ffffffffffffffff8082111561495c57600080fd5b90840190610100828703121561497157600080fd5b61497961422f565b82358281111561498857600080fd5b61499488828601614834565b8252506020830135828111156149a957600080fd5b6149b5888286016142d2565b6020830152506040830135828111156149cd57600080fd5b6149d988828601614346565b6040830152506060830135828111156149f157600080fd5b6149fd88828601614898565b606083015250608083013582811115614a1557600080fd5b614a2188828601614346565b60808301525060a083013560a082015260c083013582811115614a4357600080fd5b614a4f88828601614346565b60c08301525060e083013560e0820152809450505050614a7160208401614926565b90509250929050565b600060a08284031215614a8c57600080fd5b60405160a0810181811067ffffffffffffffff82111715614aaf57614aaf614200565b6040528235614abd8161406e565b8152602083013563ffffffff81168114614ad657600080fd5b60208201526040830135614ae9816140a0565b60408201526060830135614afc816140a0565b60608201526080830135614b0f816140a0565b60808201529392505050565b803560ff81168114613fc757600080fd5b60008060008060008060c08789031215614b4557600080fd5b863567ffffffffffffffff80821115614b5d57600080fd5b614b698a838b016142d2565b97506020890135915080821115614b7f57600080fd5b614b8b8a838b016142d2565b9650614b9960408a01614b1b565b95506060890135915080821115614baf57600080fd5b614bbb8a838b01614504565b9450614bc960808a016140b6565b935060a0890135915080821115614bdf57600080fd5b50614bec89828a01614504565b9150509295509295509295565b600060208284031215614c0b57600080fd5b8151611b4f81614918565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006000198203614c8757614c87614c45565b5060010190565b600082821015614ca057614ca0614c45565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000816000190483118215151615614cee57614cee614c45565b500290565b60008219821115614d0657614d06614c45565b500190565b600060ff821660ff84168060ff03821115614d2857614d28614c45565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff831680614d7257614d72614d30565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600060208284031215614dee57600080fd5b8151611b4f8161406e565b8051613fc7816140a0565b8051613fc78161406e565b600082601f830112614e2057600080fd5b8151614e2e6142f3826144dc565b818152846020838601011115614e4357600080fd5b61443a82602083016020870161414e565b600082601f830112614e6557600080fd5b81516020614e756142f3836142ae565b82815260059290921b84018101918181019086841115614e9457600080fd5b8286015b8481101561433b578051614eab8161406e565b8352918301918301614e98565b600082601f830112614ec957600080fd5b81516020614ed96142f3836142ae565b82815260059290921b84018101918181019086841115614ef857600080fd5b8286015b8481101561433b5780518352918301918301614efc565b600060208284031215614f2557600080fd5b815167ffffffffffffffff80821115614f3d57600080fd5b908301906101208286031215614f5257600080fd5b614f5a614259565b82518152614f6a60208401614df9565b6020820152614f7b60408401614e04565b6040820152614f8c60608401614e04565b6060820152614f9d60808401614df9565b608082015260a083015182811115614fb457600080fd5b614fc087828601614e0f565b60a08301525060c083015182811115614fd857600080fd5b614fe487828601614e54565b60c08301525060e083015182811115614ffc57600080fd5b61500887828601614eb8565b60e0830152506101009283015192810192909252509392505050565b60008261503357615033614d30565b500490565b60006020828403121561504a57600080fd5b815167ffffffffffffffff8082111561506257600080fd5b908301906080828603121561507657600080fd5b60405160808101818110838211171561509157615091614200565b6040528251828111156150a357600080fd5b6150af87828601614e54565b825250602083015191506150c28261406e565b816020820152604083015191506150d882614918565b8160408201526060830151606082015280935050505092915050565b600067ffffffffffffffff80831681851680830382111561511757615117614c45565b01949350505050565b600067ffffffffffffffff80831681810361513d5761513d614c45565b6001019392505050565b600063ffffffff80831681851680830382111561511757615117614c45565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526151968184018a614490565b905082810360808401526151aa8189614490565b905060ff871660a084015282810360c08401526151c7818761417a565b905067ffffffffffffffff851660e08401528281036101008401526151ec818561417a565b9c9b505050505050505050505050565b60208152815160208201526000602083015161010080604085015261522561012085018361417a565b9150604085015161524160608601826001600160a01b03169052565b506060850151601f198086850301608087015261525e848361417a565b935060808701519150808685030160a087015261527b8483614490565b935060a08701519150808685030160c08701526152988483614490565b935060c08701519150808685030160e0870152506152b68382614442565b92505060e085015181850152508091505092915050565b600082601f8301126152de57600080fd5b815160206152ee6142f3836142ae565b82815260059290921b8401810191818101908684111561530d57600080fd5b8286015b8481101561433b578051615324816140a0565b8352918301918301615311565b600082601f83011261534257600080fd5b815160206153526142f3836142ae565b82815260059290921b8401810191818101908684111561537157600080fd5b8286015b8481101561433b57805167ffffffffffffffff8111156153955760008081fd5b6153a38986838b0101614e0f565b845250918301918301615375565b6000602082840312156153c357600080fd5b815167ffffffffffffffff808211156153db57600080fd5b9083019061010082860312156153f057600080fd5b6153f861422f565b82518281111561540757600080fd5b615413878286016152cd565b82525060208301518281111561542857600080fd5b61543487828601614e54565b60208301525060408301518281111561544c57600080fd5b61545887828601614eb8565b60408301525060608301518281111561547057600080fd5b61547c87828601615331565b60608301525060808301518281111561549457600080fd5b6154a087828601614eb8565b60808301525060a083015160a082015260c0830151828111156154c257600080fd5b6154ce87828601614eb8565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b8381101561447257815167ffffffffffffffff16875295820195908201906001016154fe565b600081518084526020808501808196508360051b8101915082860160005b8581101561556c57828403895261555a84835161417a565b98850198935090840190600101615542565b5091979650505050505050565b60408152600083516101008060408501526155986101408501836154ea565b915060208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808685030160608701526155d48483614490565b935060408801519150808685030160808701526155f18483614442565b935060608801519150808685030160a087015261560e8483615524565b935060808801519150808685030160c087015261562b8483614442565b935060a088015160e087015260c08801519150808685030183870152506156528382614442565b60e088015161012087015286151560208701529350611b4f92505050565b60a08152600061568360a0830188614442565b82810360208401526156958188614442565b905085604084015282810360608401526156af8186614442565b9150508260808301529695505050505050565b6000602082840312156156d457600080fd5b5051919050565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156157235780818460040360031b1b83161693505b505050919050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526157658285018b614490565b91508382036080850152615779828a614490565b915060ff881660a085015283820360c0850152615796828861417a565b90861660e085015283810361010085015290506151ec818561417a56fea164736f6c634300080f000a",
}

var EVM2EVMSubscriptionOffRampABI = EVM2EVMSubscriptionOffRampMetaData.ABI

var EVM2EVMSubscriptionOffRampBin = EVM2EVMSubscriptionOffRampMetaData.Bin

func DeployEVM2EVMSubscriptionOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, blobVerifier common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig AggregateRateLimiterInterfaceRateLimiterConfig, tokenLimitsAdmin common.Address) (common.Address, *types.Transaction, *EVM2EVMSubscriptionOffRamp, error) {
	parsed, err := EVM2EVMSubscriptionOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMSubscriptionOffRampBin), backend, sourceChainId, chainId, offRampConfig, blobVerifier, afn, sourceTokens, pools, rateLimiterConfig, tokenLimitsAdmin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMSubscriptionOffRamp{EVM2EVMSubscriptionOffRampCaller: EVM2EVMSubscriptionOffRampCaller{contract: contract}, EVM2EVMSubscriptionOffRampTransactor: EVM2EVMSubscriptionOffRampTransactor{contract: contract}, EVM2EVMSubscriptionOffRampFilterer: EVM2EVMSubscriptionOffRampFilterer{contract: contract}}, nil
}

type EVM2EVMSubscriptionOffRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMSubscriptionOffRampCaller
	EVM2EVMSubscriptionOffRampTransactor
	EVM2EVMSubscriptionOffRampFilterer
}

type EVM2EVMSubscriptionOffRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMSubscriptionOffRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMSubscriptionOffRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMSubscriptionOffRampSession struct {
	Contract     *EVM2EVMSubscriptionOffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMSubscriptionOffRampCallerSession struct {
	Contract *EVM2EVMSubscriptionOffRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMSubscriptionOffRampTransactorSession struct {
	Contract     *EVM2EVMSubscriptionOffRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMSubscriptionOffRampRaw struct {
	Contract *EVM2EVMSubscriptionOffRamp
}

type EVM2EVMSubscriptionOffRampCallerRaw struct {
	Contract *EVM2EVMSubscriptionOffRampCaller
}

type EVM2EVMSubscriptionOffRampTransactorRaw struct {
	Contract *EVM2EVMSubscriptionOffRampTransactor
}

func NewEVM2EVMSubscriptionOffRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMSubscriptionOffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMSubscriptionOffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMSubscriptionOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRamp{address: address, abi: abi, EVM2EVMSubscriptionOffRampCaller: EVM2EVMSubscriptionOffRampCaller{contract: contract}, EVM2EVMSubscriptionOffRampTransactor: EVM2EVMSubscriptionOffRampTransactor{contract: contract}, EVM2EVMSubscriptionOffRampFilterer: EVM2EVMSubscriptionOffRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMSubscriptionOffRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMSubscriptionOffRampCaller, error) {
	contract, err := bindEVM2EVMSubscriptionOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampCaller{contract: contract}, nil
}

func NewEVM2EVMSubscriptionOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMSubscriptionOffRampTransactor, error) {
	contract, err := bindEVM2EVMSubscriptionOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampTransactor{contract: contract}, nil
}

func NewEVM2EVMSubscriptionOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMSubscriptionOffRampFilterer, error) {
	contract, err := bindEVM2EVMSubscriptionOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampFilterer{contract: contract}, nil
}

func bindEVM2EVMSubscriptionOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMSubscriptionOffRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMSubscriptionOffRamp.Contract.EVM2EVMSubscriptionOffRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.EVM2EVMSubscriptionOffRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.EVM2EVMSubscriptionOffRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMSubscriptionOffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(AggregateRateLimiterInterfaceTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(AggregateRateLimiterInterfaceTokenBucket)).(*AggregateRateLimiterInterfaceTokenBucket)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMSubscriptionOffRamp.Contract.CcipReceive(&_EVM2EVMSubscriptionOffRamp.CallOpts, arg0)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMSubscriptionOffRamp.Contract.CcipReceive(&_EVM2EVMSubscriptionOffRamp.CallOpts, arg0)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) GetAFN() (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetAFN(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetAFN(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) GetBlobVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getBlobVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) GetBlobVerifier() (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetBlobVerifier(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) GetBlobVerifier() (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetBlobVerifier(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

	error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getChainIDs")

	outstruct := new(GetChainIDs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetChainIDs(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetChainIDs(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOffRampInterfaceOffRampConfig)).(*BaseOffRampInterfaceOffRampConfig)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetConfig(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetConfig(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetDestinationToken(&_EVM2EVMSubscriptionOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetDestinationToken(&_EVM2EVMSubscriptionOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetDestinationTokens(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetDestinationTokens(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetExecutionState(&_EVM2EVMSubscriptionOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetExecutionState(&_EVM2EVMSubscriptionOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) GetNonce(opts *bind.CallOpts, receiver common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getNonce", receiver)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) GetNonce(receiver common.Address) (uint64, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetNonce(&_EVM2EVMSubscriptionOffRamp.CallOpts, receiver)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) GetNonce(receiver common.Address) (uint64, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetNonce(&_EVM2EVMSubscriptionOffRamp.CallOpts, receiver)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetPool(&_EVM2EVMSubscriptionOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetPool(&_EVM2EVMSubscriptionOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetPoolTokens(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetPoolTokens(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetPricesForTokens(&_EVM2EVMSubscriptionOffRamp.CallOpts, tokens)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetPricesForTokens(&_EVM2EVMSubscriptionOffRamp.CallOpts, tokens)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) GetRouter() (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetRouter(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetRouter(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.IsAFNHealthy(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.IsAFNHealthy(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.LatestConfigDetails(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.LatestConfigDetails(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) Owner() (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.Owner(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.Owner(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) Paused() (bool, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.Paused(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.Paused(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.Transmitters(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.Transmitters(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.TypeAndVersion(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.TypeAndVersion(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.AcceptOwnership(&_EVM2EVMSubscriptionOffRamp.TransactOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.AcceptOwnership(&_EVM2EVMSubscriptionOffRamp.TransactOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.AddPool(&_EVM2EVMSubscriptionOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.AddPool(&_EVM2EVMSubscriptionOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) Execute(opts *bind.TransactOpts, report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "execute", report, manualExecution)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.Execute(&_EVM2EVMSubscriptionOffRamp.TransactOpts, report, manualExecution)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.Execute(&_EVM2EVMSubscriptionOffRamp.TransactOpts, report, manualExecution)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "executeSingleMessage", message)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMSubscriptionOffRamp.TransactOpts, message)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMSubscriptionOffRamp.TransactOpts, message)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.Pause(&_EVM2EVMSubscriptionOffRamp.TransactOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.Pause(&_EVM2EVMSubscriptionOffRamp.TransactOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.RemovePool(&_EVM2EVMSubscriptionOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.RemovePool(&_EVM2EVMSubscriptionOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetAFN(&_EVM2EVMSubscriptionOffRamp.TransactOpts, afn)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetAFN(&_EVM2EVMSubscriptionOffRamp.TransactOpts, afn)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) SetBlobVerifier(opts *bind.TransactOpts, blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "setBlobVerifier", blobVerifier)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetBlobVerifier(&_EVM2EVMSubscriptionOffRamp.TransactOpts, blobVerifier)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetBlobVerifier(&_EVM2EVMSubscriptionOffRamp.TransactOpts, blobVerifier)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) SetConfig(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "setConfig", config)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetConfig(&_EVM2EVMSubscriptionOffRamp.TransactOpts, config)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetConfig(&_EVM2EVMSubscriptionOffRamp.TransactOpts, config)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) SetConfig0(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "setConfig0", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetConfig0(&_EVM2EVMSubscriptionOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetConfig0(&_EVM2EVMSubscriptionOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetPrices(&_EVM2EVMSubscriptionOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetPrices(&_EVM2EVMSubscriptionOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMSubscriptionOffRamp.TransactOpts, config)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMSubscriptionOffRamp.TransactOpts, config)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetRouter(&_EVM2EVMSubscriptionOffRamp.TransactOpts, router)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetRouter(&_EVM2EVMSubscriptionOffRamp.TransactOpts, router)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMSubscriptionOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMSubscriptionOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.TransferOwnership(&_EVM2EVMSubscriptionOffRamp.TransactOpts, to)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.TransferOwnership(&_EVM2EVMSubscriptionOffRamp.TransactOpts, to)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.Transmit(&_EVM2EVMSubscriptionOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.Transmit(&_EVM2EVMSubscriptionOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.Unpause(&_EVM2EVMSubscriptionOffRamp.TransactOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.Unpause(&_EVM2EVMSubscriptionOffRamp.TransactOpts)
}

type EVM2EVMSubscriptionOffRampAFNSetIterator struct {
	Event *EVM2EVMSubscriptionOffRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampAFNSet)
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
		it.Event = new(EVM2EVMSubscriptionOffRampAFNSet)
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

func (it *EVM2EVMSubscriptionOffRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampAFNSetIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampAFNSet)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParseAFNSet(log types.Log) (*EVM2EVMSubscriptionOffRampAFNSet, error) {
	event := new(EVM2EVMSubscriptionOffRampAFNSet)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOffRampConfigChangedIterator struct {
	Event *EVM2EVMSubscriptionOffRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampConfigChanged)
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
		it.Event = new(EVM2EVMSubscriptionOffRampConfigChanged)
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

func (it *EVM2EVMSubscriptionOffRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampConfigChangedIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampConfigChanged)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMSubscriptionOffRampConfigChanged, error) {
	event := new(EVM2EVMSubscriptionOffRampConfigChanged)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOffRampConfigSetIterator struct {
	Event *EVM2EVMSubscriptionOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampConfigSet)
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
		it.Event = new(EVM2EVMSubscriptionOffRampConfigSet)
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

func (it *EVM2EVMSubscriptionOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampConfigSet struct {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampConfigSetIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampConfigSet)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParseConfigSet(log types.Log) (*EVM2EVMSubscriptionOffRampConfigSet, error) {
	event := new(EVM2EVMSubscriptionOffRampConfigSet)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOffRampExecutionStateChangedIterator struct {
	Event *EVM2EVMSubscriptionOffRampExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampExecutionStateChanged)
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
		it.Event = new(EVM2EVMSubscriptionOffRampExecutionStateChanged)
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

func (it *EVM2EVMSubscriptionOffRampExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampExecutionStateChanged struct {
	SequenceNumber uint64
	State          uint8
	Raw            types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMSubscriptionOffRampExecutionStateChangedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampExecutionStateChangedIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampExecutionStateChanged)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParseExecutionStateChanged(log types.Log) (*EVM2EVMSubscriptionOffRampExecutionStateChanged, error) {
	event := new(EVM2EVMSubscriptionOffRampExecutionStateChanged)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOffRampOffRampConfigSetIterator struct {
	Event *EVM2EVMSubscriptionOffRampOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampOffRampConfigSet)
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
		it.Event = new(EVM2EVMSubscriptionOffRampOffRampConfigSet)
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

func (it *EVM2EVMSubscriptionOffRampOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampOffRampConfigSet struct {
	Config BaseOffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampOffRampConfigSetIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampOffRampConfigSet)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParseOffRampConfigSet(log types.Log) (*EVM2EVMSubscriptionOffRampOffRampConfigSet, error) {
	event := new(EVM2EVMSubscriptionOffRampOffRampConfigSet)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOffRampOffRampRouterSetIterator struct {
	Event *EVM2EVMSubscriptionOffRampOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampOffRampRouterSet)
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
		it.Event = new(EVM2EVMSubscriptionOffRampOffRampRouterSet)
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

func (it *EVM2EVMSubscriptionOffRampOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampOffRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMSubscriptionOffRampOffRampRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampOffRampRouterSetIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampOffRampRouterSet)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParseOffRampRouterSet(log types.Log) (*EVM2EVMSubscriptionOffRampOffRampRouterSet, error) {
	event := new(EVM2EVMSubscriptionOffRampOffRampRouterSet)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOffRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMSubscriptionOffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMSubscriptionOffRampOwnershipTransferRequested)
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

func (it *EVM2EVMSubscriptionOffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMSubscriptionOffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampOwnershipTransferRequestedIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampOwnershipTransferRequested)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMSubscriptionOffRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMSubscriptionOffRampOwnershipTransferRequested)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOffRampOwnershipTransferredIterator struct {
	Event *EVM2EVMSubscriptionOffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMSubscriptionOffRampOwnershipTransferred)
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

func (it *EVM2EVMSubscriptionOffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMSubscriptionOffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampOwnershipTransferredIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampOwnershipTransferred)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMSubscriptionOffRampOwnershipTransferred, error) {
	event := new(EVM2EVMSubscriptionOffRampOwnershipTransferred)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOffRampPausedIterator struct {
	Event *EVM2EVMSubscriptionOffRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampPaused)
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
		it.Event = new(EVM2EVMSubscriptionOffRampPaused)
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

func (it *EVM2EVMSubscriptionOffRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampPausedIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampPaused)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParsePaused(log types.Log) (*EVM2EVMSubscriptionOffRampPaused, error) {
	event := new(EVM2EVMSubscriptionOffRampPaused)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOffRampPoolAddedIterator struct {
	Event *EVM2EVMSubscriptionOffRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampPoolAdded)
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
		it.Event = new(EVM2EVMSubscriptionOffRampPoolAdded)
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

func (it *EVM2EVMSubscriptionOffRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampPoolAddedIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampPoolAdded)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMSubscriptionOffRampPoolAdded, error) {
	event := new(EVM2EVMSubscriptionOffRampPoolAdded)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOffRampPoolRemovedIterator struct {
	Event *EVM2EVMSubscriptionOffRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampPoolRemoved)
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
		it.Event = new(EVM2EVMSubscriptionOffRampPoolRemoved)
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

func (it *EVM2EVMSubscriptionOffRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampPoolRemovedIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampPoolRemoved)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMSubscriptionOffRampPoolRemoved, error) {
	event := new(EVM2EVMSubscriptionOffRampPoolRemoved)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOffRampTokenPriceChangedIterator struct {
	Event *EVM2EVMSubscriptionOffRampTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampTokenPriceChanged)
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
		it.Event = new(EVM2EVMSubscriptionOffRampTokenPriceChanged)
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

func (it *EVM2EVMSubscriptionOffRampTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampTokenPriceChangedIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampTokenPriceChanged)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMSubscriptionOffRampTokenPriceChanged, error) {
	event := new(EVM2EVMSubscriptionOffRampTokenPriceChanged)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOffRampTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMSubscriptionOffRampTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampTokensRemovedFromBucket)
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
		it.Event = new(EVM2EVMSubscriptionOffRampTokensRemovedFromBucket)
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

func (it *EVM2EVMSubscriptionOffRampTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampTokensRemovedFromBucketIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampTokensRemovedFromBucket)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMSubscriptionOffRampTokensRemovedFromBucket, error) {
	event := new(EVM2EVMSubscriptionOffRampTokensRemovedFromBucket)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOffRampTransmittedIterator struct {
	Event *EVM2EVMSubscriptionOffRampTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampTransmitted)
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
		it.Event = new(EVM2EVMSubscriptionOffRampTransmitted)
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

func (it *EVM2EVMSubscriptionOffRampTransmittedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampTransmittedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampTransmittedIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampTransmitted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampTransmitted)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParseTransmitted(log types.Log) (*EVM2EVMSubscriptionOffRampTransmitted, error) {
	event := new(EVM2EVMSubscriptionOffRampTransmitted)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOffRampUnpausedIterator struct {
	Event *EVM2EVMSubscriptionOffRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampUnpaused)
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
		it.Event = new(EVM2EVMSubscriptionOffRampUnpaused)
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

func (it *EVM2EVMSubscriptionOffRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampUnpausedIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampUnpaused)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMSubscriptionOffRampUnpaused, error) {
	event := new(EVM2EVMSubscriptionOffRampUnpaused)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMSubscriptionOffRamp.abi.Events["AFNSet"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseAFNSet(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseConfigChanged(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseConfigSet(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseExecutionStateChanged(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["OffRampConfigSet"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseOffRampConfigSet(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["OffRampRouterSet"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseOffRampRouterSet(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["Paused"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParsePaused(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParsePoolAdded(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParsePoolRemoved(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseTokenPriceChanged(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseTokensRemovedFromBucket(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["Transmitted"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseTransmitted(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMSubscriptionOffRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMSubscriptionOffRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMSubscriptionOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMSubscriptionOffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf")
}

func (EVM2EVMSubscriptionOffRampOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa0")
}

func (EVM2EVMSubscriptionOffRampOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4")
}

func (EVM2EVMSubscriptionOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMSubscriptionOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMSubscriptionOffRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMSubscriptionOffRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMSubscriptionOffRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMSubscriptionOffRampTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMSubscriptionOffRampTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMSubscriptionOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (EVM2EVMSubscriptionOffRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRamp) Address() common.Address {
	return _EVM2EVMSubscriptionOffRamp.address
}

type EVM2EVMSubscriptionOffRampInterface interface {
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

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMSubscriptionOffRampAFNSet, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMSubscriptionOffRampConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMSubscriptionOffRampConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMSubscriptionOffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMSubscriptionOffRampExecutionStateChanged, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*EVM2EVMSubscriptionOffRampOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMSubscriptionOffRampOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*EVM2EVMSubscriptionOffRampOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMSubscriptionOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMSubscriptionOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMSubscriptionOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMSubscriptionOffRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMSubscriptionOffRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMSubscriptionOffRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMSubscriptionOffRampPoolRemoved, error)

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMSubscriptionOffRampTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMSubscriptionOffRampTokensRemovedFromBucket, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMSubscriptionOffRampTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMSubscriptionOffRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
