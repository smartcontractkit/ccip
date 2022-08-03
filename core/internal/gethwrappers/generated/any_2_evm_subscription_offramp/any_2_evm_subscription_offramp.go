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
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated"
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

type BaseOffRampInterfaceOffRampConfig struct {
	ExecutionDelaySeconds                   uint64
	MaxDataSize                             uint64
	MaxTokensLength                         uint64
	PermissionLessExecutionThresholdSeconds uint32
}

type CCIPAny2EVMMessage struct {
	SourceChainId  *big.Int
	SequenceNumber uint64
	Sender         []byte
	Receiver       common.Address
	Data           []byte
	Tokens         []common.Address
	Amounts        []*big.Int
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

var EVM2EVMSubscriptionOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SubscriptionNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlobVerifier\",\"outputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"s_receiverToNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_router\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"}],\"name\":\"setBlobVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005f2838038062005f2883398101604081905262000034916200061f565b6000805460ff191681556001908a908a908a908a908a908a908a908a908a9083908390869084903390819081620000b25760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ec57620000ec8162000325565b5050506001600160a01b038216158062000104575080155b156200012357604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001695760405162d8548360e71b815260040160405180910390fd5b81516200017e906005906020850190620003d6565b5060005b825181101562000260576000828281518110620001a357620001a362000759565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001ed57620001ed62000759565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff1916600117905562000258816200076f565b905062000182565b50505060809890985250505060a0939093528151600980546020850151604086015160609096015163ffffffff16600160c01b0263ffffffff60c01b196001600160401b03978816600160801b0216600160801b600160e01b031992881668010000000000000000026001600160801b0319909416979095169690961791909117169190911792909217909155600880546001600160a01b039092166001600160a01b03199092169190911790555050151560c0525062000797975050505050505050565b336001600160a01b038216036200037f5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a9565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200042e579160200282015b828111156200042e57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003f7565b506200043c92915062000440565b5090565b5b808211156200043c576000815560010162000441565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b038111828210171562000492576200049262000457565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620004c357620004c362000457565b604052919050565b80516001600160401b0381168114620004e357600080fd5b919050565b6001600160a01b0381168114620004fe57600080fd5b50565b8051620004e381620004e8565b60006001600160401b038211156200052a576200052a62000457565b5060051b60200190565b600082601f8301126200054657600080fd5b815160206200055f62000559836200050e565b62000498565b82815260059290921b840181019181810190868411156200057f57600080fd5b8286015b84811015620005a75780516200059981620004e8565b835291830191830162000583565b509695505050505050565b600082601f830112620005c457600080fd5b81516020620005d762000559836200050e565b82815260059290921b84018101918181019086841115620005f757600080fd5b8286015b84811015620005a75780516200061181620004e8565b8352918301918301620005fb565b6000806000806000806000806000898b036101808112156200064057600080fd5b8a51995060208b015198506080603f19820112156200065e57600080fd5b50620006696200046d565b6200067760408c01620004cb565b81526200068760608c01620004cb565b60208201526200069a60808c01620004cb565b604082015260a08b015163ffffffff81168114620006b757600080fd5b60608201529650620006cc60c08b0162000501565b9550620006dc60e08b0162000501565b9450620006ed6101008b0162000501565b6101208b01519094506001600160401b03808211156200070c57600080fd5b6200071a8d838e0162000534565b94506101408c01519150808211156200073257600080fd5b50620007418c828d01620005b2565b9250506101608a015190509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b6000600182016200079057634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05161575a620007ce60003960006112bb015260006104fe0152600081816104310152613886015261575a6000f3fe608060405234801561001057600080fd5b50600436106102925760003560e01c80638b721c7111610160578063bbe4f6db116100d8578063e3d0e7121161008c578063eb511dd411610071578063eb511dd414610774578063f2fde38b14610787578063fb777fc71461079a57600080fd5b8063e3d0e71214610753578063e402bc541461076657600080fd5b8063c0d78655116100bd578063c0d786551461063c578063c3f909d41461064f578063e16e632c1461073357600080fd5b8063bbe4f6db146105f0578063be9b03f11461062957600080fd5b8063b034909c1161012f578063b1dc65a411610114578063b1dc65a4146105b7578063b4069b31146105ca578063b6608c3b146105dd57600080fd5b8063b034909c14610591578063b0f479a11461059957600080fd5b80638b721c71146105285780638da5cb5b1461053b578063a639d1c01461055e578063afcb95d71461057157600080fd5b8063681fba161161020e57806381411834116101c25780638456cb59116101a75780638456cb59146104f157806385e1f4d0146104f957806389c065681461052057600080fd5b806381411834146104ac57806381ff7048146104c157600080fd5b806374be2150116101f357806374be21501461042c57806379ba5097146104615780637c34718c1461046957600080fd5b8063681fba1614610404578063744b92e21461041957600080fd5b80633f4ba83a116102655780635b16ebb71161024a5780635b16ebb7146103a25780635c975abb146103db5780636133dc24146103e657600080fd5b80633f4ba83a14610377578063567c814b1461037f57600080fd5b8063108ee5fc14610297578063142a98fc146102ac578063181f5a77146102ef5780632222dd4214610338575b600080fd5b6102aa6102a5366004613f0b565b6107ad565b005b6102d96102ba366004613f49565b67ffffffffffffffff166000908152600a602052604090205460ff1690565b6040516102e69190613f95565b60405180910390f35b61032b6040518060400160405280602081526020017f45564d3245564d537562736372697074696f6e4f666652616d7020312e302e3081525081565b6040516102e6919061404c565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102e6565b6102aa610889565b61039261038d36600461405f565b61089b565b60405190151581526020016102e6565b6103926103b0366004613f0b565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff16610392565b60085473ffffffffffffffffffffffffffffffffffffffff16610352565b61040c6109e2565b6040516102e69190614083565b6102aa6104273660046140dd565b610ac1565b6104537f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016102e6565b6102aa610ec0565b610493610477366004613f0b565b60116020526000908152604090205467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016102e6565b6104b4610fe7565b6040516102e69190614167565b600d54600b546040805163ffffffff808516825264010000000090940490931660208401528201526060016102e6565b6102aa611056565b6104537f000000000000000000000000000000000000000000000000000000000000000081565b61040c611066565b6102aa6105363660046143d0565b6110d3565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610352565b6102aa61056c366004613f0b565b61112f565b6040805160018152600060208201819052918101919091526060016102e6565b600354610453565b60075473ffffffffffffffffffffffffffffffffffffffff16610352565b6102aa6105c536600461452f565b61117e565b6103526105d8366004613f0b565b611825565b6102aa6105eb36600461405f565b61192a565b6103526105fe366004613f0b565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6102aa610637366004614711565b6119aa565b6102aa61064a366004613f0b565b61248b565b6106ec604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260095467ffffffffffffffff80821683526801000000000000000082048116602084015270010000000000000000000000000000000082041692820192909252780100000000000000000000000000000000000000000000000090910463ffffffff16606082015290565b60408051825167ffffffffffffffff908116825260208085015182169083015283830151169181019190915260609182015163ffffffff16918101919091526080016102e6565b6007546103529073ffffffffffffffffffffffffffffffffffffffff1681565b6102aa61076136600461486b565b612502565b6102aa610292366004614938565b6102aa6107823660046140dd565b612ee5565b6102aa610795366004613f0b565b61312d565b6102aa6107a8366004614974565b61313e565b6107b5613283565b73ffffffffffffffffffffffffffffffffffffffff8116610802576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b610891613283565b610899613309565b565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa15801561090b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061092f9190614a02565b1580156109dc5750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa1580156109a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109cb9190614a1f565b602001516109d99084614aaa565b11155b92915050565b60055460609067ffffffffffffffff811115610a0057610a0061417a565b604051908082528060200260200182016040528015610a29578160200160208202803683370190505b50905060005b600554811015610abd57610a7660058281548110610a4f57610a4f614ac1565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16611825565b828281518110610a8857610a88614ac1565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152610ab681614af0565b9050610a2f565b5090565b610ac9613283565b6005546000819003610b07576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610ba2576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610c0b576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006005610c1a600185614aaa565b81548110610c2a57610c2a614ac1565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff1681548110610c7c57610c7c614ac1565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166005610cab600186614aaa565b81548110610cbb57610cbb614ac1565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110610d2957610d29614ac1565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480610dcb57610dcb614b28565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610f46576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060601080548060200260200160405190810160405280929190818152602001828054801561104c57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611021575b5050505050905090565b61105e613283565b6108996133ea565b6060600580548060200260200160405190810160405280929190818152602001828054801561104c5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611021575050505050905090565b33301461110c576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111238160a001518260c0015183606001516134aa565b61112c81613544565b50565b611137613283565b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916111d491849163ffffffff851691908e908e908190840183828082843760009201919091525061363692505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600b54808252600c5460ff808216602085015261010090910416928201929092529083146112a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610f3d565b6112b78b8b8b8b8b8b6136c6565b60007f000000000000000000000000000000000000000000000000000000000000000015611314576002826020015183604001516112f59190614b57565b6112ff9190614bab565b61130a906001614b57565b60ff16905061132a565b6020820151611324906001614b57565b60ff1690505b888114611393576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610f3d565b8887146113fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610f3d565b336000908152600e602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561143f5761143f613f66565b600281111561145057611450613f66565b905250905060028160200151600281111561146d5761146d613f66565b1480156114b457506010816000015160ff168154811061148f5761148f614ac1565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b61151a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610f3d565b505050505060008888604051611531929190614bcd565b604051908190038120611548918c90602001614bdd565b604051602081830303815290604052805190602001209050611568613eca565b604080518082019091526000808252602082015260005b8881101561180357600060018588846020811061159e5761159e614ac1565b6115ab91901a601b614b57565b8d8d868181106115bd576115bd614ac1565b905060200201358c8c878181106115d6576115d6614ac1565b9050602002013560405160008152602001604052604051611613949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611635573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600e602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156116b5576116b5613f66565b60028111156116c6576116c6613f66565b90525092506001836020015160028111156116e3576116e3613f66565b1461174a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610f3d565b8251849060ff16601f811061176157611761614ac1565b6020020151156117cd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610f3d565b600184846000015160ff16601f81106117e8576117e8614ac1565b91151560209092020152506117fc81614af0565b905061157f565b5050505063ffffffff811061181a5761181a614bf9565b505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff80821660009081526004602052604081205490911680611886576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8084166000908152600460208181526040928390205483517f21df0da700000000000000000000000000000000000000000000000000000000815293519416936321df0da79380840193908290030181865afa1580156118ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119239190614c28565b9392505050565b611932613283565b8060000361196c576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251910161087d565b60005460ff1615611a17576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610f3d565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aa89190614a02565b15611ade576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa158015611b4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b729190614a1f565b9050600354816020015142611b879190614aaa565b1115611bbf576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075473ffffffffffffffffffffffffffffffffffffffff1680611c0f576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060840151516000819003611c50576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115611c6b57611c6b61417a565b604051908082528060200260200182016040528015611c94578160200160208202803683370190505b50905060008267ffffffffffffffff811115611cb257611cb261417a565b604051908082528060200260200182016040528015611d4657816020015b60408051610120810182526000808252602080830182905292820181905260608083018290526080830182905260a0830181905260c0830181905260e083015261010082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff909201910181611cd05790505b50905060005b83811015611e195787606001518181518110611d6a57611d6a614ac1565b6020026020010151806020019051810190611d859190614d67565b828281518110611d9757611d97614ac1565b6020026020010181905250600088606001518281518110611dba57611dba614ac1565b6020026020010151604051602001611dd29190614e78565b60405160208183030381529060405290508080519060200120848381518110611dfd57611dfd614ac1565b602090810291909101015250611e1281614af0565b9050611d4c565b50600080611e3a848a608001518b60a001518c60c001518d60e0015161377d565b915091506000835182611e4d9190614e9e565b9050888015611e8857506009547801000000000000000000000000000000000000000000000000900463ffffffff16611e868442614aaa565b105b15611ebf576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408a015160005b8781101561247d5760005a90506000878381518110611ee857611ee8614ac1565b602002602001015190506000611f1b826020015167ffffffffffffffff166000908152600a602052604090205460ff1690565b90506002816003811115611f3157611f31613f66565b03611f7a5760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610f3d565b60608201516040517f0cbebc2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526000918e1690630cbebc2490602401600060405180830381865afa158015611fee573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526120349190810190614f16565b602081015190915073ffffffffffffffffffffffffffffffffffffffff166120a65760608301516040517f8515736a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610f3d565b6080830151606084015173ffffffffffffffffffffffffffffffffffffffff16600090815260116020526040812054909167ffffffffffffffff908116916120f091166001614fd2565b67ffffffffffffffff161490508080612127575081604001511580156121275750600383600381111561212557612125613f66565b145b61216f5760808401516040517fb0241f4a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610f3d565b61217884613882565b60005b8460c00151518110156121c1576121ae8560c0015182815181106121a1576121a1614ac1565b60200260200101516139e7565b50806121b981614af0565b91505061217b565b5060208481015167ffffffffffffffff166000908152600a9091526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561221a61221586613a63565b613b88565b60208087015167ffffffffffffffff166000908152600a90915260409020805491925082917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600183600381111561227557612275613f66565b0217905550846020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf826040516122b89190613f95565b60405180910390a28180156122ea575060028160038111156122dc576122dc613f66565b14806122ea57508260400151155b1561235657606085015173ffffffffffffffffffffffffffffffffffffffff166000908152601160205260408120805467ffffffffffffffff169161232e83614ffe565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b508f612467578d73ffffffffffffffffffffffffffffffffffffffff1663294d266185606001518660400151670de0b6b3a76400008b8b8151811061239d5761239d614ac1565b60200260200101513a8e5a6123b2908e614aaa565b6123bc9190615025565b6123c6919061503d565b6123d0919061503d565b6123da9190614e9e565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815273ffffffffffffffffffffffffffffffffffffffff93841660048201529290911660248301526044820152606401600060405180830381600087803b15801561244e57600080fd5b505af1158015612462573d6000803e3d6000fd5b505050505b50505050508061247690614af0565b9050611ec7565b505050505050505050505050565b612493613283565b600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b855185518560ff16601f831115612575576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610f3d565b600081116125df576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610f3d565b81831461266d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610f3d565b61267881600361503d565b83116126e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610f3d565b6126e8613283565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600f54156128db57600f5460009061274090600190614aaa565b90506000600f828154811061275757612757614ac1565b60009182526020822001546010805473ffffffffffffffffffffffffffffffffffffffff9092169350908490811061279157612791614ac1565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600e909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600f8054919250908061281157612811614b28565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055601080548061287a5761287a614b28565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550612726915050565b60005b815151811015612d40576000600e60008460000151848151811061290457612904614ac1565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561294e5761294e613f66565b146129b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610f3d565b6040805180820190915260ff821681526001602082015282518051600e91600091859081106129e6576129e6614ac1565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612a8757612a87613f66565b021790555060009150612a979050565b600e600084602001518481518110612ab157612ab1614ac1565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff166002811115612afb57612afb613f66565b14612b62576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610f3d565b6040805180820190915260ff82168152602081016002815250600e600084602001518481518110612b9557612b95614ac1565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612c3657612c36613f66565b02179055505082518051600f925083908110612c5457612c54614ac1565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9093169290921790915582015180516010919083908110612cd057612cd0614ac1565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055612d3981614af0565b90506128de565b506040810151600c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600d80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092612dd292869290821691161761507a565b92506101000a81548163ffffffff021916908363ffffffff160217905550612e314630600d60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613cba565b600b81905582518051600c805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600d5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612ed0988b98919763ffffffff909216969095919491939192615099565b60405180910390a15050505050505050505050565b612eed613283565b73ffffffffffffffffffffffffffffffffffffffff82161580612f24575073ffffffffffffffffffffffffffffffffffffffff8116155b15612f5b576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612ff7576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b613135613283565b61112c81613d65565b613146613283565b80516009805460208085018051604080880180516060808b01805163ffffffff9081167801000000000000000000000000000000000000000000000000027fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff95861670010000000000000000000000000000000002167fffffffff000000000000000000000000ffffffffffffffffffffffffffffffff98861668010000000000000000027fffffffffffffffffffffffffffffffff00000000000000000000000000000000909c169d86169d8e179b909b17979097169990991795909517909855825198895293518416948801949094529251909116918501919091525116908201527f187b05a44b331bbf26fcef5672820acbcd72e33ec0214a0b89042db7b73b46a29060800160405180910390a150565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610899576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610f3d565b60005460ff16613375576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610f3d565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60005460ff1615613457576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610f3d565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586133c03390565b81518351146134e5576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b835181101561353e5761352e84828151811061350657613506614ac1565b602002602001015184838151811061352057613520614ac1565b602002602001015184613e60565b61353781614af0565b90506134e8565b50505050565b606081015173ffffffffffffffffffffffffffffffffffffffff163b6135675750565b6007546040517fc39a285b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063c39a285b906135bd90849060040161515f565b6020604051808303816000875af11580156135dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136009190614a02565b61112c576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008180602001905181019061364c919061533e565b6040517fbe9b03f1000000000000000000000000000000000000000000000000000000008152909150309063be9b03f19061368e908490600190600401615506565b600060405180830381600087803b1580156136a857600080fd5b505af11580156136bc573d6000803e3d6000fd5b5050505050505050565b60006136d382602061503d565b6136de85602061503d565b6136ea88610144615025565b6136f49190615025565b6136fe9190615025565b613709906000615025565b9050368114613774576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610f3d565b50505050505050565b60008060005a6008546040517fe71e65ce00000000000000000000000000000000000000000000000000000000815291925060009173ffffffffffffffffffffffffffffffffffffffff9091169063e71e65ce906137e7908c908c908c908c908c906004016155fd565b6020604051808303816000875af1158015613806573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061382a919061564f565b905060008111613866576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a6138729084614aaa565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146138e25780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610f3d565b60095460c08201515170010000000000000000000000000000000090910467ffffffffffffffff16108061392057508060e00151518160c001515114155b156139695760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610f3d565b60095460a0820151516801000000000000000090910467ffffffffffffffff16101561112c5760095460a0820151516040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090920467ffffffffffffffff1660048301526024820152604401610f3d565b73ffffffffffffffffffffffffffffffffffffffff8181166000908152600460205260409020541680613a5e576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610f3d565b919050565b613acb60405180610100016040528060008152602001600067ffffffffffffffff16815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001600081525090565b60405180610100016040528083600001518152602001836020015167ffffffffffffffff1681526020018360400151604051602001613b26919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b6040516020818303038152906040528152602001836060015173ffffffffffffffffffffffffffffffffffffffff1681526020018360a0015181526020018360c0015181526020018360e0015181526020018361010001518152509050919050565b6040517f8b721c710000000000000000000000000000000000000000000000000000000081526000903090638b721c7190613bc790859060040161515f565b600060405180830381600087803b158015613be157600080fd5b505af1925050508015613bf2575060015b613cb2573d808015613c20576040519150601f19603f3d011682016040523d82523d6000602084013e613c25565b606091505b50613c2f81615668565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da80000000000000000000000000000000000000000000000000000000003613c805750600392915050565b6040517f2532cf4500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506002919050565b6000808a8a8a8a8a8a8a8a8a604051602001613cde999897969594939291906156b8565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603613de4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610f3d565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000613e6b846139e7565b6040517fea6192a200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690529192509082169063ea6192a29060440161368e565b604051806103e00160405280601f906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff8116811461112c57600080fd5b600060208284031215613f1d57600080fd5b813561192381613ee9565b67ffffffffffffffff8116811461112c57600080fd5b8035613a5e81613f28565b600060208284031215613f5b57600080fd5b813561192381613f28565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310613fd0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b83811015613ff1578181015183820152602001613fd9565b8381111561353e5750506000910152565b6000815180845261401a816020860160208601613fd6565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006119236020830184614002565b60006020828403121561407157600080fd5b5035919050565b8035613a5e81613ee9565b6020808252825182820181905260009190848201906040850190845b818110156140d157835173ffffffffffffffffffffffffffffffffffffffff168352928401929184019160010161409f565b50909695505050505050565b600080604083850312156140f057600080fd5b82356140fb81613ee9565b9150602083013561410b81613ee9565b809150509250929050565b600081518084526020808501945080840160005b8381101561415c57815173ffffffffffffffffffffffffffffffffffffffff168752958201959082019060010161412a565b509495945050505050565b6020815260006119236020830184614116565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff811182821017156141cd576141cd61417a565b60405290565b604051610120810167ffffffffffffffff811182821017156141cd576141cd61417a565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561423e5761423e61417a565b604052919050565b600067ffffffffffffffff8211156142605761426061417a565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261429d57600080fd5b81356142b06142ab82614246565b6141f7565b8181528460208386010111156142c557600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff8211156142fc576142fc61417a565b5060051b60200190565b600082601f83011261431757600080fd5b813560206143276142ab836142e2565b82815260059290921b8401810191818101908684111561434657600080fd5b8286015b8481101561436a57803561435d81613ee9565b835291830191830161434a565b509695505050505050565b600082601f83011261438657600080fd5b813560206143966142ab836142e2565b82815260059290921b840181019181810190868411156143b557600080fd5b8286015b8481101561436a57803583529183019183016143b9565b6000602082840312156143e257600080fd5b813567ffffffffffffffff808211156143fa57600080fd5b90830190610100828603121561440f57600080fd5b6144176141a9565b8235815261442760208401613f3e565b602082015260408301358281111561443e57600080fd5b61444a8782860161428c565b60408301525061445c60608401614078565b606082015260808301358281111561447357600080fd5b61447f8782860161428c565b60808301525060a08301358281111561449757600080fd5b6144a387828601614306565b60a08301525060c0830135828111156144bb57600080fd5b6144c787828601614375565b60c08301525060e083013560e082015280935050505092915050565b60008083601f8401126144f557600080fd5b50813567ffffffffffffffff81111561450d57600080fd5b6020830191508360208260051b850101111561452857600080fd5b9250929050565b60008060008060008060008060e0898b03121561454b57600080fd5b606089018a81111561455c57600080fd5b8998503567ffffffffffffffff8082111561457657600080fd5b818b0191508b601f83011261458a57600080fd5b81358181111561459957600080fd5b8c60208285010111156145ab57600080fd5b6020830199508098505060808b01359150808211156145c957600080fd5b6145d58c838d016144e3565b909750955060a08b01359150808211156145ee57600080fd5b506145fb8b828c016144e3565b999c989b50969995989497949560c00135949350505050565b600082601f83011261462557600080fd5b813560206146356142ab836142e2565b82815260059290921b8401810191818101908684111561465457600080fd5b8286015b8481101561436a57803561466b81613f28565b8352918301918301614658565b600082601f83011261468957600080fd5b813560206146996142ab836142e2565b82815260059290921b840181019181810190868411156146b857600080fd5b8286015b8481101561436a57803567ffffffffffffffff8111156146dc5760008081fd5b6146ea8986838b010161428c565b8452509183019183016146bc565b801515811461112c57600080fd5b8035613a5e816146f8565b6000806040838503121561472457600080fd5b823567ffffffffffffffff8082111561473c57600080fd5b90840190610100828703121561475157600080fd5b6147596141a9565b82358281111561476857600080fd5b61477488828601614614565b82525060208301358281111561478957600080fd5b61479588828601614306565b6020830152506040830135828111156147ad57600080fd5b6147b988828601614375565b6040830152506060830135828111156147d157600080fd5b6147dd88828601614678565b6060830152506080830135828111156147f557600080fd5b61480188828601614375565b60808301525060a083013560a082015260c08301358281111561482357600080fd5b61482f88828601614375565b60c08301525060e083013560e082015280945050505061485160208401614706565b90509250929050565b803560ff81168114613a5e57600080fd5b60008060008060008060c0878903121561488457600080fd5b863567ffffffffffffffff8082111561489c57600080fd5b6148a88a838b01614306565b975060208901359150808211156148be57600080fd5b6148ca8a838b01614306565b96506148d860408a0161485a565b955060608901359150808211156148ee57600080fd5b6148fa8a838b0161428c565b945061490860808a01613f3e565b935060a089013591508082111561491e57600080fd5b5061492b89828a0161428c565b9150509295509295509295565b60006020828403121561494a57600080fd5b813567ffffffffffffffff81111561496157600080fd5b8201610100818503121561192357600080fd5b60006080828403121561498657600080fd5b6040516080810181811067ffffffffffffffff821117156149a9576149a961417a565b60405282356149b781613f28565b815260208301356149c781613f28565b602082015260408301356149da81613f28565b6040820152606083013563ffffffff811681146149f657600080fd5b60608201529392505050565b600060208284031215614a1457600080fd5b8151611923816146f8565b600060608284031215614a3157600080fd5b6040516060810181811067ffffffffffffffff82111715614a5457614a5461417a565b80604052508251815260208301516020820152604083015160408201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015614abc57614abc614a7b565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614b2157614b21614a7b565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff03821115614b7457614b74614a7b565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff831680614bbe57614bbe614b7c565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600060208284031215614c3a57600080fd5b815161192381613ee9565b8051613a5e81613f28565b8051613a5e81613ee9565b600082601f830112614c6c57600080fd5b8151614c7a6142ab82614246565b818152846020838601011115614c8f57600080fd5b614ca0826020830160208701613fd6565b949350505050565b600082601f830112614cb957600080fd5b81516020614cc96142ab836142e2565b82815260059290921b84018101918181019086841115614ce857600080fd5b8286015b8481101561436a578051614cff81613ee9565b8352918301918301614cec565b600082601f830112614d1d57600080fd5b81516020614d2d6142ab836142e2565b82815260059290921b84018101918181019086841115614d4c57600080fd5b8286015b8481101561436a5780518352918301918301614d50565b600060208284031215614d7957600080fd5b815167ffffffffffffffff80821115614d9157600080fd5b908301906101208286031215614da657600080fd5b614dae6141d3565b82518152614dbe60208401614c45565b6020820152614dcf60408401614c50565b6040820152614de060608401614c50565b6060820152614df160808401614c45565b608082015260a083015182811115614e0857600080fd5b614e1487828601614c5b565b60a08301525060c083015182811115614e2c57600080fd5b614e3887828601614ca8565b60c08301525060e083015182811115614e5057600080fd5b614e5c87828601614d0c565b60e0830152506101009283015192810192909252509392505050565b6000815260008251614e91816001850160208701613fd6565b9190910160010192915050565b600082614ead57614ead614b7c565b500490565b600082601f830112614ec357600080fd5b81516020614ed36142ab836142e2565b82815260059290921b84018101918181019086841115614ef257600080fd5b8286015b8481101561436a578051614f0981613ee9565b8352918301918301614ef6565b600060208284031215614f2857600080fd5b815167ffffffffffffffff80821115614f4057600080fd5b9083019060808286031215614f5457600080fd5b604051608081018181108382111715614f6f57614f6f61417a565b604052825182811115614f8157600080fd5b614f8d87828601614eb2565b82525060208301519150614fa082613ee9565b81602082015260408301519150614fb6826146f8565b8160408201526060830151606082015280935050505092915050565b600067ffffffffffffffff808316818516808303821115614ff557614ff5614a7b565b01949350505050565b600067ffffffffffffffff80831681810361501b5761501b614a7b565b6001019392505050565b6000821982111561503857615038614a7b565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561507557615075614a7b565b500290565b600063ffffffff808316818516808303821115614ff557614ff5614a7b565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526150c98184018a614116565b905082810360808401526150dd8189614116565b905060ff871660a084015282810360c08401526150fa8187614002565b905067ffffffffffffffff851660e084015282810361010084015261511f8185614002565b9c9b505050505050505050505050565b600081518084526020808501945080840160005b8381101561415c57815187529582019590820190600101615143565b602081528151602082015260006020830151615187604084018267ffffffffffffffff169052565b5060408301516101008060608501526151a4610120850183614002565b915060608501516151cd608086018273ffffffffffffffffffffffffffffffffffffffff169052565b5060808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808685030160a08701526152088483614002565b935060a08701519150808685030160c08701526152258483614116565b935060c08701519150808685030160e087015250615243838261512f565b92505060e085015181850152508091505092915050565b600082601f83011261526b57600080fd5b8151602061527b6142ab836142e2565b82815260059290921b8401810191818101908684111561529a57600080fd5b8286015b8481101561436a5780516152b181613f28565b835291830191830161529e565b600082601f8301126152cf57600080fd5b815160206152df6142ab836142e2565b82815260059290921b840181019181810190868411156152fe57600080fd5b8286015b8481101561436a57805167ffffffffffffffff8111156153225760008081fd5b6153308986838b0101614c5b565b845250918301918301615302565b60006020828403121561535057600080fd5b815167ffffffffffffffff8082111561536857600080fd5b90830190610100828603121561537d57600080fd5b6153856141a9565b82518281111561539457600080fd5b6153a08782860161525a565b8252506020830151828111156153b557600080fd5b6153c187828601614eb2565b6020830152506040830151828111156153d957600080fd5b6153e587828601614d0c565b6040830152506060830151828111156153fd57600080fd5b615409878286016152be565b60608301525060808301518281111561542157600080fd5b61542d87828601614d0c565b60808301525060a083015160a082015260c08301518281111561544f57600080fd5b61545b87828601614d0c565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b8381101561415c57815167ffffffffffffffff168752958201959082019060010161548b565b600081518084526020808501808196508360051b8101915082860160005b858110156154f95782840389526154e7848351614002565b988501989350908401906001016154cf565b5091979650505050505050565b6040815260008351610100806040850152615525610140850183615477565b915060208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808685030160608701526155618483614116565b9350604088015191508086850301608087015261557e848361512f565b935060608801519150808685030160a087015261559b84836154b1565b935060808801519150808685030160c08701526155b8848361512f565b935060a088015160e087015260c08801519150808685030183870152506155df838261512f565b60e08801516101208701528615156020870152935061192392505050565b60a08152600061561060a083018861512f565b8281036020840152615622818861512f565b9050856040840152828103606084015261563c818661512f565b9150508260808301529695505050505050565b60006020828403121561566157600080fd5b5051919050565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156156b05780818460040360031b1b83161693505b505050919050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526156ff8285018b614116565b91508382036080850152615713828a614116565b915060ff881660a085015283820360c08501526157308288614002565b90861660e0850152838103610100850152905061511f818561400256fea164736f6c634300080f000a",
}

var EVM2EVMSubscriptionOffRampABI = EVM2EVMSubscriptionOffRampMetaData.ABI

var EVM2EVMSubscriptionOffRampBin = EVM2EVMSubscriptionOffRampMetaData.Bin

func DeployEVM2EVMSubscriptionOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, blobVerifier common.Address, onRampAddress common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, maxTimeWithoutAFNSignal *big.Int) (common.Address, *types.Transaction, *EVM2EVMSubscriptionOffRamp, error) {
	parsed, err := EVM2EVMSubscriptionOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMSubscriptionOffRampBin), backend, sourceChainId, chainId, offRampConfig, blobVerifier, onRampAddress, afn, sourceTokens, pools, maxTimeWithoutAFNSignal)
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) CHAINID() (*big.Int, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.CHAINID(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) CHAINID() (*big.Int, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.CHAINID(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "SOURCE_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) SOURCECHAINID() (*big.Int, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SOURCECHAINID(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) SOURCECHAINID() (*big.Int, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SOURCECHAINID(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessage) error {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) CcipReceive(arg0 CCIPAny2EVMMessage) error {
	return _EVM2EVMSubscriptionOffRamp.Contract.CcipReceive(&_EVM2EVMSubscriptionOffRamp.CallOpts, arg0)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) CcipReceive(arg0 CCIPAny2EVMMessage) error {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getMaxSecondsWithoutAFNHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMSubscriptionOffRamp.CallOpts)
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "isHealthy", timeNow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.IsHealthy(&_EVM2EVMSubscriptionOffRamp.CallOpts, timeNow)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.IsHealthy(&_EVM2EVMSubscriptionOffRamp.CallOpts, timeNow)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) IsPool(addr common.Address) (bool, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.IsPool(&_EVM2EVMSubscriptionOffRamp.CallOpts, addr)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) IsPool(addr common.Address) (bool, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.IsPool(&_EVM2EVMSubscriptionOffRamp.CallOpts, addr)
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) SReceiverToNonce(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "s_receiverToNonce", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) SReceiverToNonce(arg0 common.Address) (uint64, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SReceiverToNonce(&_EVM2EVMSubscriptionOffRamp.CallOpts, arg0)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) SReceiverToNonce(arg0 common.Address) (uint64, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SReceiverToNonce(&_EVM2EVMSubscriptionOffRamp.CallOpts, arg0)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) SRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "s_router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) SRouter() (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SRouter(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) SRouter() (common.Address, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SRouter(&_EVM2EVMSubscriptionOffRamp.CallOpts)
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessage) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "executeSingleMessage", message)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) ExecuteSingleMessage(message CCIPAny2EVMMessage) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMSubscriptionOffRamp.TransactOpts, message)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMMessage) (*types.Transaction, error) {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetConfig(&_EVM2EVMSubscriptionOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetConfig(&_EVM2EVMSubscriptionOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) SetConfig0(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "setConfig0", config)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) SetConfig0(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetConfig0(&_EVM2EVMSubscriptionOffRamp.TransactOpts, config)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) SetConfig0(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetConfig0(&_EVM2EVMSubscriptionOffRamp.TransactOpts, config)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMSubscriptionOffRamp.TransactOpts, newTime)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMSubscriptionOffRamp.TransactOpts, newTime)
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

type EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSetIterator struct {
	Event *EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet)
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
		it.Event = new(EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet)
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

func (it *EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSetIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSetIterator{contract: _EVM2EVMSubscriptionOffRamp.contract, event: "AFNMaxHeartbeatTimeSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet)
				if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampFilterer) ParseAFNMaxHeartbeatTimeSet(log types.Log) (*EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet, error) {
	event := new(EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet)
	if err := _EVM2EVMSubscriptionOffRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	case _EVM2EVMSubscriptionOffRamp.abi.Events["AFNMaxHeartbeatTimeSet"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseAFNMaxHeartbeatTimeSet(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["AFNSet"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseAFNSet(log)
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
	case _EVM2EVMSubscriptionOffRamp.abi.Events["Transmitted"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseTransmitted(log)
	case _EVM2EVMSubscriptionOffRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMSubscriptionOffRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet) Topic() common.Hash {
	return common.HexToHash("0x72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251")
}

func (EVM2EVMSubscriptionOffRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMSubscriptionOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMSubscriptionOffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf")
}

func (EVM2EVMSubscriptionOffRampOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x187b05a44b331bbf26fcef5672820acbcd72e33ec0214a0b89042db7b73b46a2")
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
	CHAINID(opts *bind.CallOpts) (*big.Int, error)

	SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error)

	CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessage) error

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetBlobVerifier(opts *bind.CallOpts) (common.Address, error)

	GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error)

	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error)

	IsPool(opts *bind.CallOpts, addr common.Address) (bool, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	SReceiverToNonce(opts *bind.CallOpts, arg0 common.Address) (uint64, error)

	SRouter(opts *bind.CallOpts) (common.Address, error)

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessage) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetBlobVerifier(opts *bind.TransactOpts, blobVerifier common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	SetConfig0(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error)

	SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSetIterator, error)

	WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet) (event.Subscription, error)

	ParseAFNMaxHeartbeatTimeSet(log types.Log) (*EVM2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet, error)

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMSubscriptionOffRampAFNSet, error)

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

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMSubscriptionOffRampTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOffRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOffRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMSubscriptionOffRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
