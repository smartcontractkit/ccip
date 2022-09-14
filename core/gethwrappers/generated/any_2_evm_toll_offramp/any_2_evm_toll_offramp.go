// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package any_2_evm_toll_offramp

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

var EVM2EVMTollOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeTaken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlobVerifier\",\"outputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"merkleGasShare\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVM2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"overheadGasToll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"}],\"name\":\"setBlobVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200609b3803806200609b83398101604081905262000034916200063c565b6000805460ff191681556001908a908a908a908a908a908a908a908a908a90829082908690869089903390819081620000b45760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ee57620000ee8162000385565b5050506001600160a01b0381166200011957604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200015b5760405162d8548360e71b815260040160405180910390fd5b81516200017090600490602085019062000436565b5060005b82518110156200023b57600082828151811062000195576200019562000796565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001df57620001df62000796565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b0292909116919091179055506200023381620007ac565b905062000174565b5050600580546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600855600993909355600a55600b91909155875116620002b9576040516342bcdf7f60e11b815260040160405180910390fd5b5050506080958652505060a0929092528051600e8054602084015160408501516001600160401b03908116600160c01b026001600160c01b0363ffffffff909316600160a01b026001600160c01b03199094166001600160a01b03968716179390931791909116919091179091556060830151600f80549490960151821668010000000000000000026001600160801b031990941691161791909117909255600d8054919092166001600160a01b031991909116179055151560c05250620007d4975050505050505050565b336001600160a01b03821603620003df5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ab565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200048e579160200282015b828111156200048e57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000457565b506200049c929150620004a0565b5090565b5b808211156200049c5760008155600101620004a1565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715620004f257620004f2620004b7565b60405290565b6001600160a01b03811681146200050e57600080fd5b50565b80516200051e81620004f8565b919050565b80516001600160401b03811681146200051e57600080fd5b600082601f8301126200054d57600080fd5b815160206001600160401b03808311156200056c576200056c620004b7565b8260051b604051601f19603f83011681018181108482111715620005945762000594620004b7565b604052938452858101830193838101925087851115620005b357600080fd5b83870191505b84821015620005df578151620005cf81620004f8565b83529183019190830190620005b9565b979650505050505050565b600060408284031215620005fd57600080fd5b604080519081016001600160401b0381118282101715620006225762000622620004b7565b604052825181526020928301519281019290925250919050565b6000806000806000806000806000898b036101c08112156200065d57600080fd5b8a51995060208b0151985060a0603f19820112156200067b57600080fd5b5062000686620004cd565b60408b01516200069681620004f8565b815260608b015163ffffffff81168114620006b057600080fd5b6020820152620006c360808c0162000523565b6040820152620006d660a08c0162000523565b6060820152620006e960c08c0162000523565b60808201529650620006fe60e08b0162000511565b95506200070f6101008b0162000511565b6101208b01519095506001600160401b03808211156200072e57600080fd5b6200073c8d838e016200053b565b95506101408c01519150808211156200075457600080fd5b50620007638c828d016200053b565b935050620007768b6101608c01620005ea565b9150620007876101a08b0162000511565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b600060018201620007cd57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c0516158826200081960003960006115960152600081816102f60152611e0c0152600081816102d301528181611de7015261357c01526158826000f3fe608060405234801561001057600080fd5b50600436106102c85760003560e01c806381ff70481161017b578063b1dc65a4116100d8578063c3f909d41161008c578063e3d0e71211610071578063e3d0e71214610774578063eb511dd414610787578063f2fde38b1461079a57600080fd5b8063c3f909d41461066f578063c9029f6a1461076157600080fd5b8063bbe4f6db116100bd578063bbe4f6db1461061d578063be9b03f114610649578063c0d786551461065c57600080fd5b8063b1dc65a4146105f7578063b4069b311461060a57600080fd5b8063918725431161012f578063a8b640c111610114578063a8b640c1146105a6578063afcb95d7146105c6578063b0f479a1146105e657600080fd5b80639187254314610580578063a639d1c01461059357600080fd5b806389c065681161016057806389c06568146105275780638da5cb5b1461052f57806390c2339b1461054557600080fd5b806381ff7048146104ef5780638456cb591461051f57600080fd5b80634741062e11610229578063681fba16116101dd578063744b92e2116101c2578063744b92e2146104cc57806379ba5097146104df57806381411834146104e757600080fd5b8063681fba16146104a4578063694ec2b1146104b957600080fd5b8063599f64311161020e578063599f6431146104775780635c975abb146104885780636133dc241461049357600080fd5b80634741062e146104365780634dbcb0ef1461045657600080fd5b80632222dd421161028057806339aa92641161026557806339aa9264146104085780633f4ba83a1461041b5780634352fa9f1461042357600080fd5b80632222dd42146103d5578063351f0faf146103fa57600080fd5b8063142a98fc116102b1578063142a98fc1461033a578063147809b314610374578063181f5a771461038c57600080fd5b8063087ae6df146102cd578063108ee5fc14610325575b600080fd5b604080517f000000000000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006020820152015b60405180910390f35b61033861033336600461411e565b6107ad565b005b61036761034836600461415c565b67ffffffffffffffff1660009081526010602052604090205460ff1690565b60405161031c919061418f565b61037c610856565b604051901515815260200161031c565b6103c86040518060400160405280601881526020017f45564d3245564d546f6c6c4f666652616d7020312e302e30000000000000000081525081565b60405161031c919061420f565b6002546001600160a01b03165b6040516001600160a01b03909116815260200161031c565b6103386102c8366004614222565b61033861041636600461411e565b6108e3565b61033861090d565b6103386104313660046143f1565b61091f565b610449610444366004614455565b610b5b565b60405161031c91906144cd565b610469610464366004614559565b610c23565b60405190815260200161031c565b6005546001600160a01b03166103e2565b60005460ff1661037c565b600d546001600160a01b03166103e2565b6104ac610cd2565b60405161031c91906146ba565b6103386104c73660046146cd565b610d97565b6103386104da3660046147f3565b610e10565b6103386110e1565b6104ac6111bc565b6013546011546040805163ffffffff8085168252640100000000909404909316602084015282015260600161031c565b61033861121e565b6104ac61122e565b60005461010090046001600160a01b03166103e2565b61054d61128e565b60405161031c91908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b61033861058e36600461482c565b61132f565b6103386105a136600461411e565b611449565b6104696105b436600461487b565b60176020526000908152604090205481565b60408051600181526000602082018190529181019190915260600161031c565b600c546001600160a01b03166103e2565b6103386106053660046148e0565b611473565b6103e261061836600461411e565b611a46565b6103e261062b36600461411e565b6001600160a01b039081166000908152600360205260409020541690565b610338610657366004614b26565b611b1b565b61033861066a36600461411e565b61224f565b6107056040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a081018252600e546001600160a01b038116825263ffffffff600160a01b820416602083015267ffffffffffffffff600160c01b909104811692820192909252600f548083166060830152680100000000000000009004909116608082015290565b6040805182516001600160a01b0316815260208084015163ffffffff16908201528282015167ffffffffffffffff908116928201929092526060808401518316908201526080928301519091169181019190915260a00161031c565b61033861076f366004614c6f565b6122a1565b610338610782366004614d21565b61240b565b6103386107953660046147f3565b612b69565b6103386107a836600461411e565b612d13565b6107b5612d24565b6001600160a01b0381166107f5576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b038381166001600160a01b031983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28910160405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa1580156108b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108dd9190614dee565b15905090565b6108eb612d24565b600580546001600160a01b0319166001600160a01b0392909216919091179055565b610915612d24565b61091d612d83565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561095c57506005546001600160a01b03163314155b1561097a576040516307b66ab160e51b815260040160405180910390fd5b8151815181146109b6576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610a105760066000600783815481106109db576109db614e0b565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610a0981614e37565b90506109bc565b5060005b82811015610b40576000858281518110610a3057610a30614e0b565b6020026020010151905060006001600160a01b0316816001600160a01b031603610a86576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610a9857610a98614e0b565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610afd57610afd614e0b565b6020026020010151604051610b279291906001600160a01b03929092168252602082015260400190565b60405180910390a150610b3981614e37565b9050610a14565b508351610b54906007906020870190614074565b5050505050565b80516060908067ffffffffffffffff811115610b7957610b79614269565b604051908082528060200260200182016040528015610ba2578160200160208202803683370190505b50915060005b81811015610c1c5760066000858381518110610bc657610bc6614e0b565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610c0157610c01614e0b565b6020908102919091010152610c1581614e37565b9050610ba8565b5050919050565b6000808260800151518360a001515160206014610c409190614e50565b610c4a9190614e68565b610c55906086614e50565b610c5f9190614e50565b90506000610c6e601083614e68565b9050610a28611bbc8560a00151516001610c889190614e50565b610c9490618aac614e68565b6156b8610ca18986614e50565b610cab9190614e50565b610cb59190614e50565b610cbf9190614e50565b610cc99190614e50565b95945050505050565b60045460609067ffffffffffffffff811115610cf057610cf0614269565b604051908082528060200260200182016040528015610d19578160200160208202803683370190505b50905060005b600454811015610d9357610d5960048281548110610d3f57610d3f614e0b565b6000918252602090912001546001600160a01b0316611a46565b828281518110610d6b57610d6b614e0b565b6001600160a01b0390921660209283029190910190910152610d8c81614e37565b9050610d1f565b5090565b333014610dd0576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60808101515115610e0457610ded81608001518260c00151612e1f565b610e048160a001518260c001518360400151612fb2565b610e0d8161304c565b50565b610e18612d24565b6004546000819003610e56576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352600160a01b9093046bffffffffffffffffffffffff169082015290610eba576040516302721e1f60e61b815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614610f09576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004610f18600185614e87565b81548110610f2857610f28614e0b565b9060005260206000200160009054906101000a90046001600160a01b03169050600482602001516bffffffffffffffffffffffff1681548110610f6d57610f6d614e0b565b6000918252602090912001546001600160a01b03166004610f8f600186614e87565b81548110610f9f57610f9f614e0b565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600483602001516bffffffffffffffffffffffff1681548110610ff357610ff3614e0b565b6000918252602080832090910180546001600160a01b039485166001600160a01b03199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff909216600160a01b0291909216179055600480548061105f5761105f614e9e565b60008281526020808220600019908401810180546001600160a01b03191690559092019092556001600160a01b03878116808452600383526040808520949094558351908152908716918101919091527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b031633146111405760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180546001600160a01b03191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060601680548060200260200160405190810160405280929190818152602001828054801561121457602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116111f6575b5050505050905090565b611226612d24565b61091d61310b565b60606004805480602002602001604051908101604052809291908181526020018280548015611214576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116111f6575050505050905090565b6112b96040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b546060820181905242906000906112f39083614e87565b6020840151845191925061131f9161130b9084614e68565b856040015161131a9190614e50565b613193565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561136c57506005546001600160a01b03163314155b1561138a576040516307b66ab160e51b815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff116113de576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113e860086131a9565b602081015160098190558151600855600a546114049190613193565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b611451612d24565b600d80546001600160a01b0319166001600160a01b0392909216919091179055565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916114c991849163ffffffff851691908e908e908190840183828082843760009201919091525061325692505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260115480825260125460ff808216602085015261010090910416928201929092529083146115845760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401611137565b6115928b8b8b8b8b8b6132cf565b60007f0000000000000000000000000000000000000000000000000000000000000000156115ef576002826020015183604001516115d09190614eb4565b6115da9190614eef565b6115e5906001614eb4565b60ff169050611605565b60208201516115ff906001614eb4565b60ff1690505b8881146116545760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401611137565b8887146116a35760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401611137565b3360009081526014602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156116e6576116e6614179565b60028111156116f7576116f7614179565b905250905060028160200151600281111561171457611714614179565b14801561174e57506016816000015160ff168154811061173657611736614e0b565b6000918252602090912001546001600160a01b031633145b61179a5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401611137565b5050505050600088886040516117b1929190614f11565b6040519081900381206117c8918c90602001614f21565b6040516020818303038152906040528051906020012090506117e86140d5565b604080518082019091526000808252602082015260005b88811015611a2457600060018588846020811061181e5761181e614e0b565b61182b91901a601b614eb4565b8d8d8681811061183d5761183d614e0b565b905060200201358c8c8781811061185657611856614e0b565b9050602002013560405160008152602001604052604051611893949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156118b5573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526014602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561190a5761190a614179565b600281111561191b5761191b614179565b905250925060018360200151600281111561193857611938614179565b146119855760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401611137565b8251849060ff16601f811061199c5761199c614e0b565b6020020151156119ee5760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401611137565b600184846000015160ff16601f8110611a0957611a09614e0b565b9115156020909202015250611a1d81614e37565b90506117ff565b5050505063ffffffff8110611a3b57611a3b614f3d565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611a81576040516302721e1f60e61b815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611af0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b149190614f5e565b9392505050565b60005460ff1615611b6e5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611137565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611bc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611be59190614dee565b15611c1b576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c546001600160a01b0316611c5d576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060820151516000819003611c9e576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115611cb957611cb9614269565b604051908082528060200260200182016040528015611d6557816020015b611d5260405180610140016040528060008152602001600067ffffffffffffffff16815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160608152602001606081526020016060815260200160006001600160a01b0316815260200160008152602001600081525090565b815260200190600190039081611cd75790505b50905060008267ffffffffffffffff811115611d8357611d83614269565b604051908082528060200260200182016040528015611dac578160200160208202803683370190505b5090506000611e5d7f31a97e998befaf21ee85c8e3e1879003d67d46abb0b6b552882133e4c782986d600e54604080516020808201949094527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201526001600160a01b039092166080808401919091528151808403909101815260a09092019052805191012090565b905060005b84811015611ef757600087606001518281518110611e8257611e82614e0b565b6020026020010151806020019051810190611e9d919061508a565b9050611ea98184613363565b848381518110611ebb57611ebb614e0b565b60200260200101818152505080858381518110611eda57611eda614e0b565b60200260200101819052505080611ef090614e37565b9050611e62565b50600080611f188489608001518a60a001518b60c001518c60e00151613480565b600e549193509150600090600160a01b900463ffffffff16611f3a8442614e87565b11905060005b87811015612243576000878281518110611f5c57611f5c614e0b565b602002602001015190506000611f8f826020015167ffffffffffffffff1660009081526010602052604090205460ff1690565b90506002816003811115611fa557611fa5614179565b03611fee5760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611137565b8a1580611ff85750835b806120145750600381600381111561201257612012614179565b145b61204a576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61205382613578565b60008082600381111561206857612068614179565b14801561207357508b155b156120c55761208f8a518761208891906151a6565b8e856136bf565b905060006120a08460e001516138a2565b90506120ad818330613904565b8184610100018181516120c09190614e87565b905250505b60008260038111156120d9576120d9614179565b146121135760208084015167ffffffffffffffff166000908152601790915260409020546101008401805161210f908390614e87565b9052505b60208381015167ffffffffffffffff166000908152601090915260408120805460ff1916600117905561214d61214885613952565b613bec565b60208086015167ffffffffffffffff1660009081526010909152604090208054919250829160ff1916600183600381111561218a5761218a614179565b021790555060008360038111156121a3576121a3614179565b1480156121c1575060038160038111156121bf576121bf614179565b145b156121e85760208085015167ffffffffffffffff1660009081526017909152604090208290555b836020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf82604051612226919061418f565b60405180910390a2505050508061223c90614e37565b9050611f40565b50505050505050505050565b612257612d24565b600c80546001600160a01b0319166001600160a01b0383169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b6122a9612d24565b80516001600160a01b03166122ea576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600e805460208085018051604080880180516001600160a01b039098167fffffffffffffffff0000000000000000000000000000000000000000000000009096168617600160a01b63ffffffff948516021777ffffffffffffffffffffffffffffffffffffffffffffffff16600160c01b67ffffffffffffffff998a16021790965560608089018051600f80546080808e018051948e166fffffffffffffffffffffffffffffffff199093169290921768010000000000000000948e1694909402939093179091558451988952955190941695870195909552955187169085015291518516938301939093529151909216908201527fd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa09060a00161143e565b855185518560ff16601f8311156124645760405162461bcd60e51b815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401611137565b600081116124b45760405162461bcd60e51b815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401611137565b8183146125285760405162461bcd60e51b8152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401611137565b612533816003614e68565b83116125815760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401611137565b612589612d24565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b601554156126d9576015546000906125e190600190614e87565b90506000601582815481106125f8576125f8614e0b565b6000918252602082200154601680546001600160a01b039092169350908490811061262557612625614e0b565b60009182526020808320909101546001600160a01b03858116845260149092526040808420805461ffff199081169091559290911680845292208054909116905560158054919250908061267b5761267b614e9e565b600082815260209020810160001990810180546001600160a01b031916905501905560168054806126ae576126ae614e9e565b600082815260209020810160001990810180546001600160a01b0319169055019055506125c7915050565b60005b815151811015612a16576000601460008460000151848151811061270257612702614e0b565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff16600281111561273f5761273f614179565b1461278c5760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401611137565b6040805180820190915260ff821681526001602082015282518051601491600091859081106127bd576127bd614e0b565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff19161761010083600281111561281657612816614179565b0217905550600091506128269050565b601460008460200151848151811061284057612840614e0b565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff16600281111561287d5761287d614179565b146128ca5760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401611137565b6040805180820190915260ff8216815260208101600281525060146000846020015184815181106128fd576128fd614e0b565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff19161761010083600281111561295657612956614179565b02179055505082518051601592508390811061297457612974614e0b565b602090810291909101810151825460018101845560009384529282902090920180546001600160a01b0319166001600160a01b0390931692909217909155820151805160169190839081106129cb576129cb614e0b565b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b03909216919091179055612a0f81614e37565b90506126dc565b5060408101516012805460ff191660ff9092169190911790556013805467ffffffff0000000019811664010000000063ffffffff438116820292831785559083048116936001939092600092612a739286929082169116176151ba565b92506101000a81548163ffffffff021916908363ffffffff160217905550612ad24630601360009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613d05565b6011819055825180516012805460ff9092166101000261ff001990921691909117905560135460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612b54988b98919763ffffffff9092169690959194919391926151e2565b60405180910390a15050505050505050505050565b612b71612d24565b6001600160a01b0382161580612b8e57506001600160a01b038116155b15612bc5576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352600160a01b9093046bffffffffffffffffffffffff16908201529015612c43576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a519551909616600160a01b0294909816939093179093558354600181018555939091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b90920180546001600160a01b031916821790558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c91015b60405180910390a1505050565b612d1b612d24565b610e0d81613d92565b60005461010090046001600160a01b0316331461091d5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401611137565b60005460ff16612dd55760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401611137565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6000805b8351811015612f1657600060066000868481518110612e4457612e44614e0b565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600003612ed357848281518110612e8957612e89614e0b565b60200260200101516040517f9a655f7b00000000000000000000000000000000000000000000000000000000815260040161113791906001600160a01b0391909116815260200190565b838281518110612ee557612ee5614e0b565b602002602001015181612ef89190614e68565b612f029084614e50565b92505080612f0f90614e37565b9050612e23565b508015612fad57612f2760086131a9565b600a54811115612f63576040517f3bfa6f3800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060086002016000828254612f789190614e87565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba37890602001612d06565b505050565b8151835114612fed576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b83518110156130465761303684828151811061300e5761300e614e0b565b602002602001015184838151811061302857613028614e0b565b602002602001015184613904565b61303f81614e37565b9050612ff0565b50505050565b60408101516001600160a01b03163b6130625750565b600c546040517f5c9fa11e0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690635c9fa11e906130ab908490600401615278565b6020604051808303816000875af11580156130ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130ee9190614dee565b610e0d57604051631dc9e9b560e31b815260040160405180910390fd5b60005460ff161561315e5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611137565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612e023390565b60008183106131a25781611b14565b5090919050565b60018101546002820154429114806131c45750808260030154145b156131cd575050565b81600101548260020154111561320f576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260030154826132219190614e87565b60018401548454919250613248916132399084614e68565b856002015461131a9190614e50565b600284015550600390910155565b306001600160a01b031663be9b03f1828060200190518101906132799190615491565b60006040518363ffffffff1660e01b8152600401613298929190615659565b600060405180830381600087803b1580156132b257600080fd5b505af11580156132c6573d6000803e3d6000fd5b50505050505050565b60006132dc826020614e68565b6132e7856020614e68565b6132f388610144614e50565b6132fd9190614e50565b6133079190614e50565b613312906000614e50565b90503681146132c65760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401611137565b60008060001b828460200151856040015186606001518760800151805190602001208860a0015160405160200161339a91906146ba565b604051602081830303815290604052805190602001208960c001516040516020016133c591906144cd565b604051602081830303815290604052805190602001208a61012001518b60e001518c61010001516040516020016134629b9a999897969594939291909a8b5260208b019990995267ffffffffffffffff9790971660408a01526001600160a01b0395861660608a0152938516608089015260a088019290925260c087015260e0860152610100850152166101208301526101408201526101600190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600d546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce906134dd908c908c908c908c908c90600401615732565b6020604051808303816000875af11580156134fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135209190615784565b90506000811161355c576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a6135689084614e87565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146135d85780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401611137565b600f5460a0820151516801000000000000000090910467ffffffffffffffff16108061360e57508060c00151518160a001515114155b156136575760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611137565b600f5460808201515167ffffffffffffffff9091161015610e0d57600f546080820151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301526024820152604401611137565b60008060006136d18460e001516138a2565b6001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561370e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137329190614f5e565b905060005b8560200151518110156137ab57816001600160a01b03168660200151828151811061376457613764614e0b565b60200260200101516001600160a01b03160361379b578560400151818151811061379057613790614e0b565b602002602001015192505b6137a481614e37565b9050613737565b50816137ee576040517fce480bcc0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401611137565b6000670de0b6b3a7640000833a87610120015161380b8b8a610c23565b6138159190614e50565b61381f9190614e68565b6138299190614e68565b61383391906151a6565b90508461010001518111156138985760208501516101008601516040517f394a2c2700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9092166004830152602482018390526044820152606401611137565b9695505050505050565b6001600160a01b0381811660009081526003602052604090205416806138ff576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401611137565b919050565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401613298565b6139a3604051806101000160405280600081526020016060815260200160006001600160a01b0316815260200160608152602001606081526020016060815260200160608152602001600081525090565b6000806139c38460a001518560c001518660e00151876101000151613e41565b8151919350915060008167ffffffffffffffff8111156139e5576139e5614269565b604051908082528060200260200182016040528015613a0e578160200160208202803683370190505b50905060008267ffffffffffffffff811115613a2c57613a2c614269565b604051908082528060200260200182016040528015613a55578160200160208202803683370190505b50905060005b83811015613b59576000613a87878381518110613a7a57613a7a614e0b565b60200260200101516138a2565b905080838381518110613a9c57613a9c614e0b565b60200260200101906001600160a01b031690816001600160a01b031681525050806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613afa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b1e9190614f5e565b848381518110613b3057613b30614e0b565b6001600160a01b039092166020928302919091019091015250613b5281614e37565b9050613a5b565b50604051806101000160405280886000015181526020018860400151604051602001613b9491906001600160a01b0391909116815260200190565b604051602081830303815290604052815260200188606001516001600160a01b031681526020018860800151815260200183815260200182815260200185815260200188610120015181525095505050505050919050565b6040517f694ec2b1000000000000000000000000000000000000000000000000000000008152600090309063694ec2b190613c2b908590600401615278565b600060405180830381600087803b158015613c4557600080fd5b505af1925050508015613c56575060015b613cfd573d808015613c84576040519150601f19603f3d011682016040523d82523d6000602084013e613c89565b606091505b50613c938161579d565b7fffffffff0000000000000000000000000000000000000000000000000000000016631dc9e9b560e31b03613ccb5750600392915050565b6040517f2532cf4500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506002919050565b6000808a8a8a8a8a8a8a8a8a604051602001613d29999897969594939291906157ed565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b03821603613dea5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401611137565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60608060005b8651811015613ec957846001600160a01b0316878281518110613e6c57613e6c614e0b565b60200260200101516001600160a01b031603613eb95783868281518110613e9557613e95614e0b565b60200260200101818151613ea99190614e50565b90525086925085915061406b9050565b613ec281614e37565b9050613e47565b50600086516001613eda9190614e50565b67ffffffffffffffff811115613ef257613ef2614269565b604051908082528060200260200182016040528015613f1b578160200160208202803683370190505b509050600086516001613f2e9190614e50565b67ffffffffffffffff811115613f4657613f46614269565b604051908082528060200260200182016040528015613f6f578160200160208202803683370190505b50905060005b885181101561401157888181518110613f9057613f90614e0b565b6020026020010151838281518110613faa57613faa614e0b565b60200260200101906001600160a01b031690816001600160a01b031681525050878181518110613fdc57613fdc614e0b565b6020026020010151828281518110613ff657613ff6614e0b565b602090810291909101015261400a81614e37565b9050613f75565b50858289518151811061402657614026614e0b565b60200260200101906001600160a01b031690816001600160a01b031681525050848188518151811061405a5761405a614e0b565b602090810291909101015290925090505b94509492505050565b8280548282559060005260206000209081019282156140c9579160200282015b828111156140c957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614094565b50610d939291506140f4565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610d9357600081556001016140f5565b6001600160a01b0381168114610e0d57600080fd5b60006020828403121561413057600080fd5b8135611b1481614109565b67ffffffffffffffff81168114610e0d57600080fd5b80356138ff8161413b565b60006020828403121561416e57600080fd5b8135611b148161413b565b634e487b7160e01b600052602160045260246000fd5b60208101600483106141b157634e487b7160e01b600052602160045260246000fd5b91905290565b60005b838110156141d25781810151838201526020016141ba565b838111156130465750506000910152565b600081518084526141fb8160208601602086016141b7565b601f01601f19169290920160200192915050565b602081526000611b1460208301846141e3565b60006020828403121561423457600080fd5b813567ffffffffffffffff81111561424b57600080fd5b82016101008185031215611b1457600080fd5b80356138ff81614109565b634e487b7160e01b600052604160045260246000fd5b604051610140810167ffffffffffffffff811182821017156142a3576142a3614269565b60405290565b604051610100810167ffffffffffffffff811182821017156142a3576142a3614269565b604051601f8201601f1916810167ffffffffffffffff811182821017156142f6576142f6614269565b604052919050565b600067ffffffffffffffff82111561431857614318614269565b5060051b60200190565b600082601f83011261433357600080fd5b81356020614348614343836142fe565b6142cd565b82815260059290921b8401810191818101908684111561436757600080fd5b8286015b8481101561438b57803561437e81614109565b835291830191830161436b565b509695505050505050565b600082601f8301126143a757600080fd5b813560206143b7614343836142fe565b82815260059290921b840181019181810190868411156143d657600080fd5b8286015b8481101561438b57803583529183019183016143da565b6000806040838503121561440457600080fd5b823567ffffffffffffffff8082111561441c57600080fd5b61442886838701614322565b9350602085013591508082111561443e57600080fd5b5061444b85828601614396565b9150509250929050565b60006020828403121561446757600080fd5b813567ffffffffffffffff81111561447e57600080fd5b61448a84828501614322565b949350505050565b600081518084526020808501945080840160005b838110156144c2578151875295820195908201906001016144a6565b509495945050505050565b602081526000611b146020830184614492565b600067ffffffffffffffff8211156144fa576144fa614269565b50601f01601f191660200190565b600082601f83011261451957600080fd5b8135614527614343826144e0565b81815284602083860101111561453c57600080fd5b816020850160208301376000918101602001919091529392505050565b6000806040838503121561456c57600080fd5b82359150602083013567ffffffffffffffff8082111561458b57600080fd5b9084019061014082870312156145a057600080fd5b6145a861427f565b823581526145b860208401614151565b60208201526145c96040840161425e565b60408201526145da6060840161425e565b60608201526080830135828111156145f157600080fd5b6145fd88828601614508565b60808301525060a08301358281111561461557600080fd5b61462188828601614322565b60a08301525060c08301358281111561463957600080fd5b61464588828601614396565b60c08301525061465760e0840161425e565b60e08201526101008381013590820152610120928301359281019290925250919491935090915050565b600081518084526020808501945080840160005b838110156144c25781516001600160a01b031687529582019590820190600101614695565b602081526000611b146020830184614681565b6000602082840312156146df57600080fd5b813567ffffffffffffffff808211156146f757600080fd5b90830190610100828603121561470c57600080fd5b6147146142a9565b8235815260208301358281111561472a57600080fd5b61473687828601614508565b6020830152506147486040840161425e565b604082015260608301358281111561475f57600080fd5b61476b87828601614508565b60608301525060808301358281111561478357600080fd5b61478f87828601614322565b60808301525060a0830135828111156147a757600080fd5b6147b387828601614322565b60a08301525060c0830135828111156147cb57600080fd5b6147d787828601614396565b60c08301525060e083013560e082015280935050505092915050565b6000806040838503121561480657600080fd5b823561481181614109565b9150602083013561482181614109565b809150509250929050565b60006040828403121561483e57600080fd5b6040516040810181811067ffffffffffffffff8211171561486157614861614269565b604052823581526020928301359281019290925250919050565b60006020828403121561488d57600080fd5b5035919050565b60008083601f8401126148a657600080fd5b50813567ffffffffffffffff8111156148be57600080fd5b6020830191508360208260051b85010111156148d957600080fd5b9250929050565b60008060008060008060008060e0898b0312156148fc57600080fd5b606089018a81111561490d57600080fd5b8998503567ffffffffffffffff8082111561492757600080fd5b818b0191508b601f83011261493b57600080fd5b81358181111561494a57600080fd5b8c602082850101111561495c57600080fd5b6020830199508098505060808b013591508082111561497a57600080fd5b6149868c838d01614894565b909750955060a08b013591508082111561499f57600080fd5b506149ac8b828c01614894565b999c989b50969995989497949560c00135949350505050565b600082601f8301126149d657600080fd5b813560206149e6614343836142fe565b82815260059290921b84018101918181019086841115614a0557600080fd5b8286015b8481101561438b578035614a1c8161413b565b8352918301918301614a09565b600082601f830112614a3a57600080fd5b81356020614a4a614343836142fe565b82815260059290921b84018101918181019086841115614a6957600080fd5b8286015b8481101561438b578035614a8081614109565b8352918301918301614a6d565b600082601f830112614a9e57600080fd5b81356020614aae614343836142fe565b82815260059290921b84018101918181019086841115614acd57600080fd5b8286015b8481101561438b57803567ffffffffffffffff811115614af15760008081fd5b614aff8986838b0101614508565b845250918301918301614ad1565b8015158114610e0d57600080fd5b80356138ff81614b0d565b60008060408385031215614b3957600080fd5b823567ffffffffffffffff80821115614b5157600080fd5b908401906101008287031215614b6657600080fd5b614b6e6142a9565b823582811115614b7d57600080fd5b614b89888286016149c5565b825250602083013582811115614b9e57600080fd5b614baa88828601614a29565b602083015250604083013582811115614bc257600080fd5b614bce88828601614396565b604083015250606083013582811115614be657600080fd5b614bf288828601614a8d565b606083015250608083013582811115614c0a57600080fd5b614c1688828601614396565b60808301525060a083013560a082015260c083013582811115614c3857600080fd5b614c4488828601614396565b60c08301525060e083013560e0820152809450505050614c6660208401614b1b565b90509250929050565b600060a08284031215614c8157600080fd5b60405160a0810181811067ffffffffffffffff82111715614ca457614ca4614269565b6040528235614cb281614109565b8152602083013563ffffffff81168114614ccb57600080fd5b60208201526040830135614cde8161413b565b60408201526060830135614cf18161413b565b60608201526080830135614d048161413b565b60808201529392505050565b803560ff811681146138ff57600080fd5b60008060008060008060c08789031215614d3a57600080fd5b863567ffffffffffffffff80821115614d5257600080fd5b614d5e8a838b01614a29565b97506020890135915080821115614d7457600080fd5b614d808a838b01614a29565b9650614d8e60408a01614d10565b95506060890135915080821115614da457600080fd5b614db08a838b01614508565b9450614dbe60808a01614151565b935060a0890135915080821115614dd457600080fd5b50614de189828a01614508565b9150509295509295509295565b600060208284031215614e0057600080fd5b8151611b1481614b0d565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201614e4957614e49614e21565b5060010190565b60008219821115614e6357614e63614e21565b500190565b6000816000190483118215151615614e8257614e82614e21565b500290565b600082821015614e9957614e99614e21565b500390565b634e487b7160e01b600052603160045260246000fd5b600060ff821660ff84168060ff03821115614ed157614ed1614e21565b019392505050565b634e487b7160e01b600052601260045260246000fd5b600060ff831680614f0257614f02614ed9565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b634e487b7160e01b600052600160045260246000fd5b80516138ff81614109565b600060208284031215614f7057600080fd5b8151611b1481614109565b80516138ff8161413b565b600082601f830112614f9757600080fd5b8151614fa5614343826144e0565b818152846020838601011115614fba57600080fd5b61448a8260208301602087016141b7565b600082601f830112614fdc57600080fd5b81516020614fec614343836142fe565b82815260059290921b8401810191818101908684111561500b57600080fd5b8286015b8481101561438b57805161502281614109565b835291830191830161500f565b600082601f83011261504057600080fd5b81516020615050614343836142fe565b82815260059290921b8401810191818101908684111561506f57600080fd5b8286015b8481101561438b5780518352918301918301615073565b60006020828403121561509c57600080fd5b815167ffffffffffffffff808211156150b457600080fd5b9083019061014082860312156150c957600080fd5b6150d161427f565b825181526150e160208401614f7b565b60208201526150f260408401614f53565b604082015261510360608401614f53565b606082015260808301518281111561511a57600080fd5b61512687828601614f86565b60808301525060a08301518281111561513e57600080fd5b61514a87828601614fcb565b60a08301525060c08301518281111561516257600080fd5b61516e8782860161502f565b60c08301525061518060e08401614f53565b60e082015261010083810151908201526101209283015192810192909252509392505050565b6000826151b5576151b5614ed9565b500490565b600063ffffffff8083168185168083038211156151d9576151d9614e21565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526152128184018a614681565b905082810360808401526152268189614681565b905060ff871660a084015282810360c084015261524381876141e3565b905067ffffffffffffffff851660e084015282810361010084015261526881856141e3565b9c9b505050505050505050505050565b6020815281516020820152600060208301516101008060408501526152a16101208501836141e3565b915060408501516152bd60608601826001600160a01b03169052565b506060850151601f19808685030160808701526152da84836141e3565b935060808701519150808685030160a08701526152f78483614681565b935060a08701519150808685030160c08701526153148483614681565b935060c08701519150808685030160e0870152506153328382614492565b92505060e085015181850152508091505092915050565b600082601f83011261535a57600080fd5b8151602061536a614343836142fe565b82815260059290921b8401810191818101908684111561538957600080fd5b8286015b8481101561438b5780516153a08161413b565b835291830191830161538d565b600082601f8301126153be57600080fd5b815160206153ce614343836142fe565b82815260059290921b840181019181810190868411156153ed57600080fd5b8286015b8481101561438b57805161540481614109565b83529183019183016153f1565b600082601f83011261542257600080fd5b81516020615432614343836142fe565b82815260059290921b8401810191818101908684111561545157600080fd5b8286015b8481101561438b57805167ffffffffffffffff8111156154755760008081fd5b6154838986838b0101614f86565b845250918301918301615455565b6000602082840312156154a357600080fd5b815167ffffffffffffffff808211156154bb57600080fd5b9083019061010082860312156154d057600080fd5b6154d86142a9565b8251828111156154e757600080fd5b6154f387828601615349565b82525060208301518281111561550857600080fd5b615514878286016153ad565b60208301525060408301518281111561552c57600080fd5b6155388782860161502f565b60408301525060608301518281111561555057600080fd5b61555c87828601615411565b60608301525060808301518281111561557457600080fd5b6155808782860161502f565b60808301525060a083015160a082015260c0830151828111156155a257600080fd5b6155ae8782860161502f565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b838110156144c257815167ffffffffffffffff16875295820195908201906001016155de565b600081518084526020808501808196508360051b8101915082860160005b8581101561564c57828403895261563a8483516141e3565b98850198935090840190600101615622565b5091979650505050505050565b60408152600083516101008060408501526156786101408501836155ca565b91506020860151603f19808685030160608701526156968483614681565b935060408801519150808685030160808701526156b38483614492565b935060608801519150808685030160a08701526156d08483615604565b935060808801519150808685030160c08701526156ed8483614492565b935060a088015160e087015260c08801519150808685030183870152506157148382614492565b60e088015161012087015286151560208701529350611b1492505050565b60a08152600061574560a0830188614492565b82810360208401526157578188614492565b905085604084015282810360608401526157718186614492565b9150508260808301529695505050505050565b60006020828403121561579657600080fd5b5051919050565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156157e55780818460040360031b1b83161693505b505050919050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526158278285018b614681565b9150838203608085015261583b828a614681565b915060ff881660a085015283820360c085015261585882886141e3565b90861660e0850152838103610100850152905061526881856141e356fea164736f6c634300080f000a",
}

var EVM2EVMTollOffRampABI = EVM2EVMTollOffRampMetaData.ABI

var EVM2EVMTollOffRampBin = EVM2EVMTollOffRampMetaData.Bin

func DeployEVM2EVMTollOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, blobVerifier common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig AggregateRateLimiterInterfaceRateLimiterConfig, tokenLimitsAdmin common.Address) (common.Address, *types.Transaction, *EVM2EVMTollOffRamp, error) {
	parsed, err := EVM2EVMTollOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMTollOffRampBin), backend, sourceChainId, chainId, offRampConfig, blobVerifier, afn, sourceTokens, pools, rateLimiterConfig, tokenLimitsAdmin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMTollOffRamp{EVM2EVMTollOffRampCaller: EVM2EVMTollOffRampCaller{contract: contract}, EVM2EVMTollOffRampTransactor: EVM2EVMTollOffRampTransactor{contract: contract}, EVM2EVMTollOffRampFilterer: EVM2EVMTollOffRampFilterer{contract: contract}}, nil
}

type EVM2EVMTollOffRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMTollOffRampCaller
	EVM2EVMTollOffRampTransactor
	EVM2EVMTollOffRampFilterer
}

type EVM2EVMTollOffRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampSession struct {
	Contract     *EVM2EVMTollOffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMTollOffRampCallerSession struct {
	Contract *EVM2EVMTollOffRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMTollOffRampTransactorSession struct {
	Contract     *EVM2EVMTollOffRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMTollOffRampRaw struct {
	Contract *EVM2EVMTollOffRamp
}

type EVM2EVMTollOffRampCallerRaw struct {
	Contract *EVM2EVMTollOffRampCaller
}

type EVM2EVMTollOffRampTransactorRaw struct {
	Contract *EVM2EVMTollOffRampTransactor
}

func NewEVM2EVMTollOffRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMTollOffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMTollOffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMTollOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRamp{address: address, abi: abi, EVM2EVMTollOffRampCaller: EVM2EVMTollOffRampCaller{contract: contract}, EVM2EVMTollOffRampTransactor: EVM2EVMTollOffRampTransactor{contract: contract}, EVM2EVMTollOffRampFilterer: EVM2EVMTollOffRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMTollOffRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMTollOffRampCaller, error) {
	contract, err := bindEVM2EVMTollOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampCaller{contract: contract}, nil
}

func NewEVM2EVMTollOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMTollOffRampTransactor, error) {
	contract, err := bindEVM2EVMTollOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampTransactor{contract: contract}, nil
}

func NewEVM2EVMTollOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMTollOffRampFilterer, error) {
	contract, err := bindEVM2EVMTollOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampFilterer{contract: contract}, nil
}

func bindEVM2EVMTollOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMTollOffRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMTollOffRamp.Contract.EVM2EVMTollOffRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.EVM2EVMTollOffRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.EVM2EVMTollOffRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMTollOffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(AggregateRateLimiterInterfaceTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(AggregateRateLimiterInterfaceTokenBucket)).(*AggregateRateLimiterInterfaceTokenBucket)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMTollOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMTollOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMTollOffRamp.Contract.CcipReceive(&_EVM2EVMTollOffRamp.CallOpts, arg0)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMTollOffRamp.Contract.CcipReceive(&_EVM2EVMTollOffRamp.CallOpts, arg0)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) FeeTaken(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "feeTaken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) FeeTaken(arg0 *big.Int) (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.FeeTaken(&_EVM2EVMTollOffRamp.CallOpts, arg0)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) FeeTaken(arg0 *big.Int) (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.FeeTaken(&_EVM2EVMTollOffRamp.CallOpts, arg0)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetAFN() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetAFN(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetAFN(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetBlobVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getBlobVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetBlobVerifier() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetBlobVerifier(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetBlobVerifier() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetBlobVerifier(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getChainIDs")

	outstruct := new(GetChainIDs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMTollOffRamp.Contract.GetChainIDs(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMTollOffRamp.Contract.GetChainIDs(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOffRampInterfaceOffRampConfig)).(*BaseOffRampInterfaceOffRampConfig)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMTollOffRamp.Contract.GetConfig(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMTollOffRamp.Contract.GetConfig(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetDestinationToken(&_EVM2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetDestinationToken(&_EVM2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetDestinationTokens(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetDestinationTokens(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMTollOffRamp.Contract.GetExecutionState(&_EVM2EVMTollOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMTollOffRamp.Contract.GetExecutionState(&_EVM2EVMTollOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPool(&_EVM2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPool(&_EVM2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPoolTokens(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPoolTokens(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPricesForTokens(&_EVM2EVMTollOffRamp.CallOpts, tokens)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPricesForTokens(&_EVM2EVMTollOffRamp.CallOpts, tokens)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetRouter() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetRouter(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetRouter(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMTollOffRamp.Contract.IsAFNHealthy(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMTollOffRamp.Contract.IsAFNHealthy(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMTollOffRamp.Contract.LatestConfigDetails(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMTollOffRamp.Contract.LatestConfigDetails(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMTollOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMTollOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) OverheadGasToll(opts *bind.CallOpts, merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "overheadGasToll", merkleGasShare, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) OverheadGasToll(merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.OverheadGasToll(&_EVM2EVMTollOffRamp.CallOpts, merkleGasShare, message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) OverheadGasToll(merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.OverheadGasToll(&_EVM2EVMTollOffRamp.CallOpts, merkleGasShare, message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Owner() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.Owner(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.Owner(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Paused() (bool, error) {
	return _EVM2EVMTollOffRamp.Contract.Paused(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMTollOffRamp.Contract.Paused(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.Transmitters(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.Transmitters(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMTollOffRamp.Contract.TypeAndVersion(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMTollOffRamp.Contract.TypeAndVersion(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.AcceptOwnership(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.AcceptOwnership(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.AddPool(&_EVM2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.AddPool(&_EVM2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) Execute(opts *bind.TransactOpts, report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "execute", report, manualExecution)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Execute(&_EVM2EVMTollOffRamp.TransactOpts, report, manualExecution)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Execute(&_EVM2EVMTollOffRamp.TransactOpts, report, manualExecution)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "executeSingleMessage", message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRamp.TransactOpts, message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRamp.TransactOpts, message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Pause(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Pause(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.RemovePool(&_EVM2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.RemovePool(&_EVM2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetAFN(&_EVM2EVMTollOffRamp.TransactOpts, afn)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetAFN(&_EVM2EVMTollOffRamp.TransactOpts, afn)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetBlobVerifier(opts *bind.TransactOpts, blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setBlobVerifier", blobVerifier)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetBlobVerifier(&_EVM2EVMTollOffRamp.TransactOpts, blobVerifier)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetBlobVerifier(&_EVM2EVMTollOffRamp.TransactOpts, blobVerifier)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetConfig(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setConfig", config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetConfig0(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setConfig0", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig0(&_EVM2EVMTollOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig0(&_EVM2EVMTollOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetPrices(&_EVM2EVMTollOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetPrices(&_EVM2EVMTollOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetRouter(&_EVM2EVMTollOffRamp.TransactOpts, router)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetRouter(&_EVM2EVMTollOffRamp.TransactOpts, router)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMTollOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMTollOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.TransferOwnership(&_EVM2EVMTollOffRamp.TransactOpts, to)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.TransferOwnership(&_EVM2EVMTollOffRamp.TransactOpts, to)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Transmit(&_EVM2EVMTollOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Transmit(&_EVM2EVMTollOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Unpause(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Unpause(&_EVM2EVMTollOffRamp.TransactOpts)
}

type EVM2EVMTollOffRampAFNSetIterator struct {
	Event *EVM2EVMTollOffRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampAFNSet)
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
		it.Event = new(EVM2EVMTollOffRampAFNSet)
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

func (it *EVM2EVMTollOffRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampAFNSetIterator{contract: _EVM2EVMTollOffRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampAFNSet)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseAFNSet(log types.Log) (*EVM2EVMTollOffRampAFNSet, error) {
	event := new(EVM2EVMTollOffRampAFNSet)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampConfigChangedIterator struct {
	Event *EVM2EVMTollOffRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampConfigChanged)
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
		it.Event = new(EVM2EVMTollOffRampConfigChanged)
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

func (it *EVM2EVMTollOffRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampConfigChangedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampConfigChanged)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMTollOffRampConfigChanged, error) {
	event := new(EVM2EVMTollOffRampConfigChanged)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampConfigSetIterator struct {
	Event *EVM2EVMTollOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampConfigSet)
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
		it.Event = new(EVM2EVMTollOffRampConfigSet)
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

func (it *EVM2EVMTollOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampConfigSet struct {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampConfigSetIterator{contract: _EVM2EVMTollOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampConfigSet)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseConfigSet(log types.Log) (*EVM2EVMTollOffRampConfigSet, error) {
	event := new(EVM2EVMTollOffRampConfigSet)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampExecutionStateChangedIterator struct {
	Event *EVM2EVMTollOffRampExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampExecutionStateChanged)
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
		it.Event = new(EVM2EVMTollOffRampExecutionStateChanged)
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

func (it *EVM2EVMTollOffRampExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampExecutionStateChanged struct {
	SequenceNumber uint64
	State          uint8
	Raw            types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMTollOffRampExecutionStateChangedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampExecutionStateChangedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampExecutionStateChanged)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseExecutionStateChanged(log types.Log) (*EVM2EVMTollOffRampExecutionStateChanged, error) {
	event := new(EVM2EVMTollOffRampExecutionStateChanged)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampOffRampConfigSetIterator struct {
	Event *EVM2EVMTollOffRampOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampOffRampConfigSet)
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
		it.Event = new(EVM2EVMTollOffRampOffRampConfigSet)
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

func (it *EVM2EVMTollOffRampOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampOffRampConfigSet struct {
	Config BaseOffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampOffRampConfigSetIterator{contract: _EVM2EVMTollOffRamp.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampOffRampConfigSet)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseOffRampConfigSet(log types.Log) (*EVM2EVMTollOffRampOffRampConfigSet, error) {
	event := new(EVM2EVMTollOffRampOffRampConfigSet)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampOffRampRouterSetIterator struct {
	Event *EVM2EVMTollOffRampOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampOffRampRouterSet)
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
		it.Event = new(EVM2EVMTollOffRampOffRampRouterSet)
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

func (it *EVM2EVMTollOffRampOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampOffRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMTollOffRampOffRampRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampOffRampRouterSetIterator{contract: _EVM2EVMTollOffRamp.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampOffRampRouterSet)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseOffRampRouterSet(log types.Log) (*EVM2EVMTollOffRampOffRampRouterSet, error) {
	event := new(EVM2EVMTollOffRampOffRampRouterSet)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMTollOffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMTollOffRampOwnershipTransferRequested)
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

func (it *EVM2EVMTollOffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampOwnershipTransferRequestedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampOwnershipTransferRequested)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMTollOffRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMTollOffRampOwnershipTransferRequested)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampOwnershipTransferredIterator struct {
	Event *EVM2EVMTollOffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMTollOffRampOwnershipTransferred)
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

func (it *EVM2EVMTollOffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampOwnershipTransferredIterator{contract: _EVM2EVMTollOffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampOwnershipTransferred)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMTollOffRampOwnershipTransferred, error) {
	event := new(EVM2EVMTollOffRampOwnershipTransferred)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampPausedIterator struct {
	Event *EVM2EVMTollOffRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampPaused)
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
		it.Event = new(EVM2EVMTollOffRampPaused)
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

func (it *EVM2EVMTollOffRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampPausedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampPaused)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParsePaused(log types.Log) (*EVM2EVMTollOffRampPaused, error) {
	event := new(EVM2EVMTollOffRampPaused)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampPoolAddedIterator struct {
	Event *EVM2EVMTollOffRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampPoolAdded)
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
		it.Event = new(EVM2EVMTollOffRampPoolAdded)
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

func (it *EVM2EVMTollOffRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampPoolAddedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampPoolAdded)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMTollOffRampPoolAdded, error) {
	event := new(EVM2EVMTollOffRampPoolAdded)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampPoolRemovedIterator struct {
	Event *EVM2EVMTollOffRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampPoolRemoved)
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
		it.Event = new(EVM2EVMTollOffRampPoolRemoved)
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

func (it *EVM2EVMTollOffRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampPoolRemovedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampPoolRemoved)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMTollOffRampPoolRemoved, error) {
	event := new(EVM2EVMTollOffRampPoolRemoved)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampTokenPriceChangedIterator struct {
	Event *EVM2EVMTollOffRampTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampTokenPriceChanged)
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
		it.Event = new(EVM2EVMTollOffRampTokenPriceChanged)
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

func (it *EVM2EVMTollOffRampTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampTokenPriceChangedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampTokenPriceChanged)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMTollOffRampTokenPriceChanged, error) {
	event := new(EVM2EVMTollOffRampTokenPriceChanged)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMTollOffRampTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampTokensRemovedFromBucket)
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
		it.Event = new(EVM2EVMTollOffRampTokensRemovedFromBucket)
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

func (it *EVM2EVMTollOffRampTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampTokensRemovedFromBucketIterator{contract: _EVM2EVMTollOffRamp.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampTokensRemovedFromBucket)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMTollOffRampTokensRemovedFromBucket, error) {
	event := new(EVM2EVMTollOffRampTokensRemovedFromBucket)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampTransmittedIterator struct {
	Event *EVM2EVMTollOffRampTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampTransmitted)
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
		it.Event = new(EVM2EVMTollOffRampTransmitted)
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

func (it *EVM2EVMTollOffRampTransmittedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTransmittedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampTransmittedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTransmitted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampTransmitted)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseTransmitted(log types.Log) (*EVM2EVMTollOffRampTransmitted, error) {
	event := new(EVM2EVMTollOffRampTransmitted)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampUnpausedIterator struct {
	Event *EVM2EVMTollOffRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampUnpaused)
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
		it.Event = new(EVM2EVMTollOffRampUnpaused)
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

func (it *EVM2EVMTollOffRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampUnpausedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampUnpaused)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMTollOffRampUnpaused, error) {
	event := new(EVM2EVMTollOffRampUnpaused)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMTollOffRamp.abi.Events["AFNSet"].ID:
		return _EVM2EVMTollOffRamp.ParseAFNSet(log)
	case _EVM2EVMTollOffRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMTollOffRamp.ParseConfigChanged(log)
	case _EVM2EVMTollOffRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMTollOffRamp.ParseConfigSet(log)
	case _EVM2EVMTollOffRamp.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMTollOffRamp.ParseExecutionStateChanged(log)
	case _EVM2EVMTollOffRamp.abi.Events["OffRampConfigSet"].ID:
		return _EVM2EVMTollOffRamp.ParseOffRampConfigSet(log)
	case _EVM2EVMTollOffRamp.abi.Events["OffRampRouterSet"].ID:
		return _EVM2EVMTollOffRamp.ParseOffRampRouterSet(log)
	case _EVM2EVMTollOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMTollOffRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMTollOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMTollOffRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMTollOffRamp.abi.Events["Paused"].ID:
		return _EVM2EVMTollOffRamp.ParsePaused(log)
	case _EVM2EVMTollOffRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMTollOffRamp.ParsePoolAdded(log)
	case _EVM2EVMTollOffRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMTollOffRamp.ParsePoolRemoved(log)
	case _EVM2EVMTollOffRamp.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMTollOffRamp.ParseTokenPriceChanged(log)
	case _EVM2EVMTollOffRamp.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMTollOffRamp.ParseTokensRemovedFromBucket(log)
	case _EVM2EVMTollOffRamp.abi.Events["Transmitted"].ID:
		return _EVM2EVMTollOffRamp.ParseTransmitted(log)
	case _EVM2EVMTollOffRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMTollOffRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMTollOffRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMTollOffRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMTollOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMTollOffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf")
}

func (EVM2EVMTollOffRampOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa0")
}

func (EVM2EVMTollOffRampOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4")
}

func (EVM2EVMTollOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMTollOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMTollOffRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMTollOffRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMTollOffRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMTollOffRampTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMTollOffRampTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMTollOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (EVM2EVMTollOffRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRamp) Address() common.Address {
	return _EVM2EVMTollOffRamp.address
}

type EVM2EVMTollOffRampInterface interface {
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

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMTollOffRampAFNSet, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMTollOffRampConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMTollOffRampConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMTollOffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMTollOffRampExecutionStateChanged, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*EVM2EVMTollOffRampOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMTollOffRampOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*EVM2EVMTollOffRampOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMTollOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMTollOffRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMTollOffRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMTollOffRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMTollOffRampPoolRemoved, error)

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMTollOffRampTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMTollOffRampTokensRemovedFromBucket, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMTollOffRampTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMTollOffRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
