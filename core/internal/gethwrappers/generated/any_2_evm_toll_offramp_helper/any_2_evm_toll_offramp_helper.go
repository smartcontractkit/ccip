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

var EVM2EVMTollOffRampHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlobVerifier\",\"outputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_router\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"}],\"name\":\"setBlobVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005d9e38038062005d9e833981016040819052620000349162000695565b88888888888888888860018989898989898989898282858333806000806000806101000a81548160ff02191690831515021790555060006001600160a01b0316826001600160a01b031603620000d15760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b0319909216919091179091558116156200010b576200010b816200039b565b5050506001600160a01b038216158062000123575080155b156200014257604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001885760405162d8548360e71b815260040160405180910390fd5b81516200019d9060059060208501906200044c565b5060005b82518110156200027f576000828281518110620001c257620001c2620007cf565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b0316815250600460008685815181106200020c576200020c620007cf565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff191660011790556200027781620007e5565b9050620001a1565b50505088608081815250508760a0818152505086600960008201518160000160006101000a8154816001600160401b0302191690836001600160401b0316021790555060208201518160000160086101000a8154816001600160401b0302191690836001600160401b0316021790555060408201518160000160106101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160000160186101000a81548163ffffffff021916908363ffffffff16021790555090505085600860006101000a8154816001600160a01b0302191690836001600160a01b0316021790555050505050505050505080151560c081151581525050505050505050505050505050505050505050506200080d565b336001600160a01b03821603620003f55760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000c8565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620004a4579160200282015b82811115620004a457825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200046d565b50620004b2929150620004b6565b5090565b5b80821115620004b25760008155600101620004b7565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715620005085762000508620004cd565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620005395762000539620004cd565b604052919050565b80516001600160401b03811681146200055957600080fd5b919050565b6001600160a01b03811681146200057457600080fd5b50565b805162000559816200055e565b60006001600160401b03821115620005a057620005a0620004cd565b5060051b60200190565b600082601f830112620005bc57600080fd5b81516020620005d5620005cf8362000584565b6200050e565b82815260059290921b84018101918181019086841115620005f557600080fd5b8286015b848110156200061d5780516200060f816200055e565b8352918301918301620005f9565b509695505050505050565b600082601f8301126200063a57600080fd5b815160206200064d620005cf8362000584565b82815260059290921b840181019181810190868411156200066d57600080fd5b8286015b848110156200061d57805162000687816200055e565b835291830191830162000671565b6000806000806000806000806000898b03610180811215620006b657600080fd5b8a51995060208b015198506080603f1982011215620006d457600080fd5b50620006df620004e3565b620006ed60408c0162000541565b8152620006fd60608c0162000541565b60208201526200071060808c0162000541565b604082015260a08b015163ffffffff811681146200072d57600080fd5b606082015296506200074260c08b0162000577565b95506200075260e08b0162000577565b9450620007636101008b0162000577565b6101208b01519094506001600160401b03808211156200078257600080fd5b620007908d838e01620005aa565b94506101408c0151915080821115620007a857600080fd5b50620007b78c828d0162000628565b9250506101608a015190509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b6000600182016200080657634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05161555a62000844600039600061128b015260006104bb0152600081816104310152613719015261555a6000f3fe608060405234801561001057600080fd5b50600436106102925760003560e01c80638da5cb5b11610160578063bbe4f6db116100d8578063e3d0e7121161008c578063eb511dd411610071578063eb511dd414610744578063f2fde38b14610757578063fb777fc71461076a57600080fd5b8063e3d0e71214610723578063e402bc541461073657600080fd5b8063c0d78655116100bd578063c0d786551461060c578063c3f909d41461061f578063e16e632c1461070357600080fd5b8063bbe4f6db146105c0578063be9b03f1146105f957600080fd5b8063b0f479a11161012f578063b4069b3111610114578063b4069b3114610587578063b57671661461059a578063b6608c3b146105ad57600080fd5b8063b0f479a114610556578063b1dc65a41461057457600080fd5b80638da5cb5b146104f8578063a639d1c01461051b578063afcb95d71461052e578063b034909c1461054e57600080fd5b8063681fba161161020e57806381ff7048116101c257806385e1f4d0116101a757806385e1f4d0146104b657806389c06568146104dd5780638b721c71146104e557600080fd5b806381ff70481461047e5780638456cb59146104ae57600080fd5b806374be2150116101f357806374be21501461042c57806379ba509714610461578063814118341461046957600080fd5b8063681fba1614610404578063744b92e21461041957600080fd5b80633f4ba83a116102655780635b16ebb71161024a5780635b16ebb7146103a25780635c975abb146103db5780636133dc24146103e657600080fd5b80633f4ba83a14610377578063567c814b1461037f57600080fd5b8063108ee5fc14610297578063142a98fc146102ac578063181f5a77146102ef5780632222dd4214610338575b600080fd5b6102aa6102a5366004613dd1565b61077d565b005b6102d96102ba366004613e0f565b67ffffffffffffffff166000908152600a602052604090205460ff1690565b6040516102e69190613e5b565b60405180910390f35b61032b6040518060400160405280601881526020017f45564d3245564d546f6c6c4f666652616d7020312e302e30000000000000000081525081565b6040516102e69190613f12565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102e6565b6102aa610859565b61039261038d366004613f25565b61086b565b60405190151581526020016102e6565b6103926103b0366004613dd1565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff16610392565b60085473ffffffffffffffffffffffffffffffffffffffff16610352565b61040c6109b2565b6040516102e69190613f49565b6102aa610427366004613fa3565b610a91565b6104537f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016102e6565b6102aa610e90565b610471610fb7565b6040516102e6919061402d565b600d54600b546040805163ffffffff808516825264010000000090940490931660208401528201526060016102e6565b6102aa611026565b6104537f000000000000000000000000000000000000000000000000000000000000000081565b61040c611036565b6102aa6104f3366004614296565b6110a3565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610352565b6102aa610529366004613dd1565b6110ff565b6040805160018152600060208201819052918101919091526060016102e6565b600354610453565b60075473ffffffffffffffffffffffffffffffffffffffff16610352565b6102aa6105823660046143f5565b61114e565b610352610595366004613dd1565b6117f5565b6102aa6105a83660046144da565b6118fa565b6102aa6105bb366004613f25565b611906565b6103526105ce366004613dd1565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6102aa610607366004614614565b611986565b6102aa61061a366004613dd1565b612331565b6106bc604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260095467ffffffffffffffff80821683526801000000000000000082048116602084015270010000000000000000000000000000000082041692820192909252780100000000000000000000000000000000000000000000000090910463ffffffff16606082015290565b60408051825167ffffffffffffffff908116825260208085015182169083015283830151169181019190915260609182015163ffffffff16918101919091526080016102e6565b6007546103529073ffffffffffffffffffffffffffffffffffffffff1681565b6102aa61073136600461476e565b6123a8565b6102aa61029236600461483b565b6102aa610752366004613fa3565b612d8b565b6102aa610765366004613dd1565b612fd3565b6102aa610778366004614877565b612fe4565b610785613129565b73ffffffffffffffffffffffffffffffffffffffff81166107d2576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b610861613129565b6108696131af565b565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa1580156108db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ff9190614905565b1580156109ac5750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610977573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099b9190614922565b602001516109a990846149ad565b11155b92915050565b60055460609067ffffffffffffffff8111156109d0576109d0614040565b6040519080825280602002602001820160405280156109f9578160200160208202803683370190505b50905060005b600554811015610a8d57610a4660058281548110610a1f57610a1f6149c4565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166117f5565b828281518110610a5857610a586149c4565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152610a86816149f3565b90506109ff565b5090565b610a99613129565b6005546000819003610ad7576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610b72576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610bdb576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006005610bea6001856149ad565b81548110610bfa57610bfa6149c4565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff1681548110610c4c57610c4c6149c4565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166005610c7b6001866149ad565b81548110610c8b57610c8b6149c4565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110610cf957610cf96149c4565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480610d9b57610d9b614a2b565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610f16576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060601080548060200260200160405190810160405280929190818152602001828054801561101c57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610ff1575b5050505050905090565b61102e613129565b610869613290565b6060600580548060200260200160405190810160405280929190818152602001828054801561101c5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610ff1575050505050905090565b3330146110dc576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110f38160a001518260c001518360600151613350565b6110fc816133ea565b50565b611107613129565b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916111a491849163ffffffff851691908e908e90819084018382808284376000920191909152506134dc92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600b54808252600c5460ff80821660208501526101009091041692820192909252908314611279576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610f0d565b6112878b8b8b8b8b8b613562565b60007f0000000000000000000000000000000000000000000000000000000000000000156112e4576002826020015183604001516112c59190614a5a565b6112cf9190614aae565b6112da906001614a5a565b60ff1690506112fa565b60208201516112f4906001614a5a565b60ff1690505b888114611363576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610f0d565b8887146113cc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610f0d565b336000908152600e602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561140f5761140f613e2c565b600281111561142057611420613e2c565b905250905060028160200151600281111561143d5761143d613e2c565b14801561148457506010816000015160ff168154811061145f5761145f6149c4565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b6114ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610f0d565b505050505060008888604051611501929190614ad0565b604051908190038120611518918c90602001614ae0565b604051602081830303815290604052805190602001209050611538613d90565b604080518082019091526000808252602082015260005b888110156117d357600060018588846020811061156e5761156e6149c4565b61157b91901a601b614a5a565b8d8d8681811061158d5761158d6149c4565b905060200201358c8c878181106115a6576115a66149c4565b90506020020135604051600081526020016040526040516115e3949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611605573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600e602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561168557611685613e2c565b600281111561169657611696613e2c565b90525092506001836020015160028111156116b3576116b3613e2c565b1461171a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610f0d565b8251849060ff16601f8110611731576117316149c4565b60200201511561179d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610f0d565b600184846000015160ff16601f81106117b8576117b86149c4565b91151560209092020152506117cc816149f3565b905061154f565b5050505063ffffffff81106117ea576117ea614afc565b505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff80821660009081526004602052604081205490911680611856576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8084166000908152600460208181526040928390205483517f21df0da700000000000000000000000000000000000000000000000000000000815293519416936321df0da79380840193908290030181865afa1580156118cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f39190614b36565b9392505050565b6110fc600080836134dc565b61190e613129565b80600003611948576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251910161084d565b60005460ff16156119f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610f0d565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a849190614905565b15611aba576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa158015611b2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b4e9190614922565b9050600354816020015142611b6391906149ad565b1115611b9b576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075473ffffffffffffffffffffffffffffffffffffffff16611bea576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060830151516000819003611c2b576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115611c4657611c46614040565b604051908082528060200260200182016040528015611c6f578160200160208202803683370190505b50905060008267ffffffffffffffff811115611c8d57611c8d614040565b604051908082528060200260200182016040528015611d6057816020015b611d4d60405180610140016040528060008152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b815260200190600190039081611cab5790505b50905060005b83811015611e2d5786606001518181518110611d8457611d846149c4565b6020026020010151806020019051810190611d9f9190614c62565b828281518110611db157611db16149c4565b602002602001018190525086606001518181518110611dd257611dd26149c4565b6020026020010151604051602001611dea9190614d7e565b60405160208183030381529060405280519060200120838281518110611e1257611e126149c4565b6020908102919091010152611e26816149f3565b9050611d66565b50600080611e4e8489608001518a60a001518b60c001518c60e00151613610565b915091506000835182611e619190614da4565b9050878015611e9c57506009547801000000000000000000000000000000000000000000000000900463ffffffff16611e9a84426149ad565b105b15611ed3576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b86811015612325576000858281518110611ef257611ef26149c4565b602002602001015190506000611f25826020015167ffffffffffffffff166000908152600a602052604090205460ff1690565b90506002816003811115611f3b57611f3b613e2c565b03611f845760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610f0d565b611f8d82613715565b60005b8260a0015151811015611fd457611fc38360a001518281518110611fb657611fb66149c4565b602002602001015161387a565b50611fcd816149f3565b9050611f90565b506000816003811115611fe957611fe9613e2c565b148015611ff457508a155b15612213576000806120098460e0015161387a565b73ffffffffffffffffffffffffffffffffffffffff166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612053573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120779190614b36565b905060005b8e602001515181101561210a578173ffffffffffffffffffffffffffffffffffffffff168f6020015182815181106120b6576120b66149c4565b602002602001015173ffffffffffffffffffffffffffffffffffffffff16036120fa578e6040015181815181106120ef576120ef6149c4565b602002602001015192505b612103816149f3565b905061207c565b508161215a576040517fce480bcc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610f0d565b6000670de0b6b3a7640000833a8761012001518a6121789190614db8565b6121829190614dd0565b61218c9190614dd0565b6121969190614da4565b90508461010001518111156121fb5760208501516101008601516040517f394a2c2700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9092166004830152602482018390526044820152606401610f0d565b61220f8560e00151866101000151306138f6565b5050505b60208281015167ffffffffffffffff166000908152600a9091526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561226b61226684613993565b613ab8565b60208085015167ffffffffffffffff166000908152600a90915260409020805491925082917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660018360038111156122c6576122c6613e2c565b0217905550826020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf826040516123099190613e5b565b60405180910390a25050508061231e906149f3565b9050611ed6565b50505050505050505050565b612339613129565b600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b855185518560ff16601f83111561241b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610f0d565b60008111612485576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610f0d565b818314612513576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610f0d565b61251e816003614dd0565b8311612586576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610f0d565b61258e613129565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600f541561278157600f546000906125e6906001906149ad565b90506000600f82815481106125fd576125fd6149c4565b60009182526020822001546010805473ffffffffffffffffffffffffffffffffffffffff90921693509084908110612637576126376149c4565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600e909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600f805491925090806126b7576126b7614a2b565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055601080548061272057612720614a2b565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055506125cc915050565b60005b815151811015612be6576000600e6000846000015184815181106127aa576127aa6149c4565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff1660028111156127f4576127f4613e2c565b1461285b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610f0d565b6040805180820190915260ff821681526001602082015282518051600e916000918590811061288c5761288c6149c4565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561292d5761292d613e2c565b02179055506000915061293d9050565b600e600084602001518481518110612957576129576149c4565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff1660028111156129a1576129a1613e2c565b14612a08576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610f0d565b6040805180820190915260ff82168152602081016002815250600e600084602001518481518110612a3b57612a3b6149c4565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612adc57612adc613e2c565b02179055505082518051600f925083908110612afa57612afa6149c4565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9093169290921790915582015180516010919083908110612b7657612b766149c4565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055612bdf816149f3565b9050612784565b506040810151600c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600d80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092612c78928692908216911617614e0d565b92506101000a81548163ffffffff021916908363ffffffff160217905550612cd74630600d60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613bea565b600b81905582518051600c805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600d5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612d76988b98919763ffffffff909216969095919491939192614e35565b60405180910390a15050505050505050505050565b612d93613129565b73ffffffffffffffffffffffffffffffffffffffff82161580612dca575073ffffffffffffffffffffffffffffffffffffffff8116155b15612e01576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612e9d576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b612fdb613129565b6110fc81613c95565b612fec613129565b80516009805460208085018051604080880180516060808b01805163ffffffff9081167801000000000000000000000000000000000000000000000000027fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff95861670010000000000000000000000000000000002167fffffffff000000000000000000000000ffffffffffffffffffffffffffffffff98861668010000000000000000027fffffffffffffffffffffffffffffffff00000000000000000000000000000000909c169d86169d8e179b909b17979097169990991795909517909855825198895293518416948801949094529251909116918501919091525116908201527f187b05a44b331bbf26fcef5672820acbcd72e33ec0214a0b89042db7b73b46a29060800160405180910390a150565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610869576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610f0d565b60005460ff1661321b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610f0d565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60005460ff16156132fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610f0d565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586132663390565b815183511461338b576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b83518110156133e4576133d48482815181106133ac576133ac6149c4565b60200260200101518483815181106133c6576133c66149c4565b6020026020010151846138f6565b6133dd816149f3565b905061338e565b50505050565b606081015173ffffffffffffffffffffffffffffffffffffffff163b61340d5750565b6007546040517fc39a285b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063c39a285b90613463908490600401614efb565b6020604051808303816000875af1158015613482573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134a69190614905565b6110fc576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff1663be9b03f18280602001905181019061350c919061513e565b60006040518363ffffffff1660e01b815260040161352b929190615306565b600060405180830381600087803b15801561354557600080fd5b505af1158015613559573d6000803e3d6000fd5b50505050505050565b600061356f826020614dd0565b61357a856020614dd0565b61358688610144614db8565b6135909190614db8565b61359a9190614db8565b6135a5906000614db8565b9050368114613559576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610f0d565b60008060005a6008546040517fe71e65ce00000000000000000000000000000000000000000000000000000000815291925060009173ffffffffffffffffffffffffffffffffffffffff9091169063e71e65ce9061367a908c908c908c908c908c906004016153fd565b6020604051808303816000875af1158015613699573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136bd919061544f565b9050600081116136f9576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a61370590846149ad565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146137755780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610f0d565b60095460a08201515170010000000000000000000000000000000090910467ffffffffffffffff1610806137b357508060c00151518160a001515114155b156137fc5760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610f0d565b6009546080820151516801000000000000000090910467ffffffffffffffff1610156110fc576009546080820151516040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090920467ffffffffffffffff1660048301526024820152604401610f0d565b73ffffffffffffffffffffffffffffffffffffffff81811660009081526004602052604090205416806138f1576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610f0d565b919050565b60006139018461387a565b6040517fea6192a200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690529192509082169063ea6192a290604401600060405180830381600087803b15801561397557600080fd5b505af1158015613989573d6000803e3d6000fd5b5050505050505050565b6139fb60405180610100016040528060008152602001600067ffffffffffffffff16815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001600081525090565b60405180610100016040528083600001518152602001836020015167ffffffffffffffff1681526020018360400151604051602001613a56919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b6040516020818303038152906040528152602001836060015173ffffffffffffffffffffffffffffffffffffffff168152602001836080015181526020018360a0015181526020018360c0015181526020018361012001518152509050919050565b6040517f8b721c710000000000000000000000000000000000000000000000000000000081526000903090638b721c7190613af7908590600401614efb565b600060405180830381600087803b158015613b1157600080fd5b505af1925050508015613b22575060015b613be2573d808015613b50576040519150601f19603f3d011682016040523d82523d6000602084013e613b55565b606091505b50613b5f81615468565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da80000000000000000000000000000000000000000000000000000000003613bb05750600392915050565b6040517f2532cf4500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506002919050565b6000808a8a8a8a8a8a8a8a8a604051602001613c0e999897969594939291906154b8565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603613d14576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610f0d565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604051806103e00160405280601f906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff811681146110fc57600080fd5b600060208284031215613de357600080fd5b81356118f381613daf565b67ffffffffffffffff811681146110fc57600080fd5b80356138f181613dee565b600060208284031215613e2157600080fd5b81356118f381613dee565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310613e96577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b83811015613eb7578181015183820152602001613e9f565b838111156133e45750506000910152565b60008151808452613ee0816020860160208601613e9c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006118f36020830184613ec8565b600060208284031215613f3757600080fd5b5035919050565b80356138f181613daf565b6020808252825182820181905260009190848201906040850190845b81811015613f9757835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101613f65565b50909695505050505050565b60008060408385031215613fb657600080fd5b8235613fc181613daf565b91506020830135613fd181613daf565b809150509250929050565b600081518084526020808501945080840160005b8381101561402257815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101613ff0565b509495945050505050565b6020815260006118f36020830184613fdc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff8111828210171561409357614093614040565b60405290565b604051610140810167ffffffffffffffff8111828210171561409357614093614040565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561410457614104614040565b604052919050565b600067ffffffffffffffff82111561412657614126614040565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261416357600080fd5b81356141766141718261410c565b6140bd565b81815284602083860101111561418b57600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff8211156141c2576141c2614040565b5060051b60200190565b600082601f8301126141dd57600080fd5b813560206141ed614171836141a8565b82815260059290921b8401810191818101908684111561420c57600080fd5b8286015b8481101561423057803561422381613daf565b8352918301918301614210565b509695505050505050565b600082601f83011261424c57600080fd5b8135602061425c614171836141a8565b82815260059290921b8401810191818101908684111561427b57600080fd5b8286015b84811015614230578035835291830191830161427f565b6000602082840312156142a857600080fd5b813567ffffffffffffffff808211156142c057600080fd5b9083019061010082860312156142d557600080fd5b6142dd61406f565b823581526142ed60208401613e04565b602082015260408301358281111561430457600080fd5b61431087828601614152565b60408301525061432260608401613f3e565b606082015260808301358281111561433957600080fd5b61434587828601614152565b60808301525060a08301358281111561435d57600080fd5b614369878286016141cc565b60a08301525060c08301358281111561438157600080fd5b61438d8782860161423b565b60c08301525060e083013560e082015280935050505092915050565b60008083601f8401126143bb57600080fd5b50813567ffffffffffffffff8111156143d357600080fd5b6020830191508360208260051b85010111156143ee57600080fd5b9250929050565b60008060008060008060008060e0898b03121561441157600080fd5b606089018a81111561442257600080fd5b8998503567ffffffffffffffff8082111561443c57600080fd5b818b0191508b601f83011261445057600080fd5b81358181111561445f57600080fd5b8c602082850101111561447157600080fd5b6020830199508098505060808b013591508082111561448f57600080fd5b61449b8c838d016143a9565b909750955060a08b01359150808211156144b457600080fd5b506144c18b828c016143a9565b999c989b50969995989497949560c00135949350505050565b6000602082840312156144ec57600080fd5b813567ffffffffffffffff81111561450357600080fd5b61450f84828501614152565b949350505050565b600082601f83011261452857600080fd5b81356020614538614171836141a8565b82815260059290921b8401810191818101908684111561455757600080fd5b8286015b8481101561423057803561456e81613dee565b835291830191830161455b565b600082601f83011261458c57600080fd5b8135602061459c614171836141a8565b82815260059290921b840181019181810190868411156145bb57600080fd5b8286015b8481101561423057803567ffffffffffffffff8111156145df5760008081fd5b6145ed8986838b0101614152565b8452509183019183016145bf565b80151581146110fc57600080fd5b80356138f1816145fb565b6000806040838503121561462757600080fd5b823567ffffffffffffffff8082111561463f57600080fd5b90840190610100828703121561465457600080fd5b61465c61406f565b82358281111561466b57600080fd5b61467788828601614517565b82525060208301358281111561468c57600080fd5b614698888286016141cc565b6020830152506040830135828111156146b057600080fd5b6146bc8882860161423b565b6040830152506060830135828111156146d457600080fd5b6146e08882860161457b565b6060830152506080830135828111156146f857600080fd5b6147048882860161423b565b60808301525060a083013560a082015260c08301358281111561472657600080fd5b6147328882860161423b565b60c08301525060e083013560e082015280945050505061475460208401614609565b90509250929050565b803560ff811681146138f157600080fd5b60008060008060008060c0878903121561478757600080fd5b863567ffffffffffffffff8082111561479f57600080fd5b6147ab8a838b016141cc565b975060208901359150808211156147c157600080fd5b6147cd8a838b016141cc565b96506147db60408a0161475d565b955060608901359150808211156147f157600080fd5b6147fd8a838b01614152565b945061480b60808a01613e04565b935060a089013591508082111561482157600080fd5b5061482e89828a01614152565b9150509295509295509295565b60006020828403121561484d57600080fd5b813567ffffffffffffffff81111561486457600080fd5b820161010081850312156118f357600080fd5b60006080828403121561488957600080fd5b6040516080810181811067ffffffffffffffff821117156148ac576148ac614040565b60405282356148ba81613dee565b815260208301356148ca81613dee565b602082015260408301356148dd81613dee565b6040820152606083013563ffffffff811681146148f957600080fd5b60608201529392505050565b60006020828403121561491757600080fd5b81516118f3816145fb565b60006060828403121561493457600080fd5b6040516060810181811067ffffffffffffffff8211171561495757614957614040565b80604052508251815260208301516020820152604083015160408201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156149bf576149bf61497e565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614a2457614a2461497e565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff03821115614a7757614a7761497e565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff831680614ac157614ac1614a7f565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b80516138f181613daf565b600060208284031215614b4857600080fd5b81516118f381613daf565b80516138f181613dee565b600082601f830112614b6f57600080fd5b8151614b7d6141718261410c565b818152846020838601011115614b9257600080fd5b61450f826020830160208701613e9c565b600082601f830112614bb457600080fd5b81516020614bc4614171836141a8565b82815260059290921b84018101918181019086841115614be357600080fd5b8286015b84811015614230578051614bfa81613daf565b8352918301918301614be7565b600082601f830112614c1857600080fd5b81516020614c28614171836141a8565b82815260059290921b84018101918181019086841115614c4757600080fd5b8286015b848110156142305780518352918301918301614c4b565b600060208284031215614c7457600080fd5b815167ffffffffffffffff80821115614c8c57600080fd5b908301906101408286031215614ca157600080fd5b614ca9614099565b82518152614cb960208401614b53565b6020820152614cca60408401614b2b565b6040820152614cdb60608401614b2b565b6060820152608083015182811115614cf257600080fd5b614cfe87828601614b5e565b60808301525060a083015182811115614d1657600080fd5b614d2287828601614ba3565b60a08301525060c083015182811115614d3a57600080fd5b614d4687828601614c07565b60c083015250614d5860e08401614b2b565b60e082015261010083810151908201526101209283015192810192909252509392505050565b6000815260008251614d97816001850160208701613e9c565b9190910160010192915050565b600082614db357614db3614a7f565b500490565b60008219821115614dcb57614dcb61497e565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614e0857614e0861497e565b500290565b600063ffffffff808316818516808303821115614e2c57614e2c61497e565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614e658184018a613fdc565b90508281036080840152614e798189613fdc565b905060ff871660a084015282810360c0840152614e968187613ec8565b905067ffffffffffffffff851660e0840152828103610100840152614ebb8185613ec8565b9c9b505050505050505050505050565b600081518084526020808501945080840160005b8381101561402257815187529582019590820190600101614edf565b602081528151602082015260006020830151614f23604084018267ffffffffffffffff169052565b506040830151610100806060850152614f40610120850183613ec8565b91506060850151614f69608086018273ffffffffffffffffffffffffffffffffffffffff169052565b5060808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808685030160a0870152614fa48483613ec8565b935060a08701519150808685030160c0870152614fc18483613fdc565b935060c08701519150808685030160e087015250614fdf8382614ecb565b92505060e085015181850152508091505092915050565b600082601f83011261500757600080fd5b81516020615017614171836141a8565b82815260059290921b8401810191818101908684111561503657600080fd5b8286015b8481101561423057805161504d81613dee565b835291830191830161503a565b600082601f83011261506b57600080fd5b8151602061507b614171836141a8565b82815260059290921b8401810191818101908684111561509a57600080fd5b8286015b848110156142305780516150b181613daf565b835291830191830161509e565b600082601f8301126150cf57600080fd5b815160206150df614171836141a8565b82815260059290921b840181019181810190868411156150fe57600080fd5b8286015b8481101561423057805167ffffffffffffffff8111156151225760008081fd5b6151308986838b0101614b5e565b845250918301918301615102565b60006020828403121561515057600080fd5b815167ffffffffffffffff8082111561516857600080fd5b90830190610100828603121561517d57600080fd5b61518561406f565b82518281111561519457600080fd5b6151a087828601614ff6565b8252506020830151828111156151b557600080fd5b6151c18782860161505a565b6020830152506040830151828111156151d957600080fd5b6151e587828601614c07565b6040830152506060830151828111156151fd57600080fd5b615209878286016150be565b60608301525060808301518281111561522157600080fd5b61522d87828601614c07565b60808301525060a083015160a082015260c08301518281111561524f57600080fd5b61525b87828601614c07565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b8381101561402257815167ffffffffffffffff168752958201959082019060010161528b565b600081518084526020808501808196508360051b8101915082860160005b858110156152f95782840389526152e7848351613ec8565b988501989350908401906001016152cf565b5091979650505050505050565b6040815260008351610100806040850152615325610140850183615277565b915060208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808685030160608701526153618483613fdc565b9350604088015191508086850301608087015261537e8483614ecb565b935060608801519150808685030160a087015261539b84836152b1565b935060808801519150808685030160c08701526153b88483614ecb565b935060a088015160e087015260c08801519150808685030183870152506153df8382614ecb565b60e0880151610120870152861515602087015293506118f392505050565b60a08152600061541060a0830188614ecb565b82810360208401526154228188614ecb565b9050856040840152828103606084015261543c8186614ecb565b9150508260808301529695505050505050565b60006020828403121561546157600080fd5b5051919050565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156154b05780818460040360031b1b83161693505b505050919050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526154ff8285018b613fdc565b91508382036080850152615513828a613fdc565b915060ff881660a085015283820360c08501526155308288613ec8565b90861660e08501528381036101008501529050614ebb8185613ec856fea164736f6c634300080f000a",
}

var EVM2EVMTollOffRampHelperABI = EVM2EVMTollOffRampHelperMetaData.ABI

var EVM2EVMTollOffRampHelperBin = EVM2EVMTollOffRampHelperMetaData.Bin

func DeployEVM2EVMTollOffRampHelper(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, blobVerifier common.Address, onRampAddress common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, maxTimeWithoutAFNSignal *big.Int) (common.Address, *types.Transaction, *EVM2EVMTollOffRampHelper, error) {
	parsed, err := EVM2EVMTollOffRampHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMTollOffRampHelperBin), backend, sourceChainId, chainId, offRampConfig, blobVerifier, onRampAddress, afn, sourceTokens, pools, maxTimeWithoutAFNSignal)
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) CHAINID() (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.CHAINID(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) CHAINID() (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.CHAINID(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "SOURCE_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SOURCECHAINID() (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SOURCECHAINID(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) SOURCECHAINID() (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SOURCECHAINID(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessage) error {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) CcipReceive(arg0 CCIPAny2EVMMessage) error {
	return _EVM2EVMTollOffRampHelper.Contract.CcipReceive(&_EVM2EVMTollOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) CcipReceive(arg0 CCIPAny2EVMMessage) error {
	return _EVM2EVMTollOffRampHelper.Contract.CcipReceive(&_EVM2EVMTollOffRampHelper.CallOpts, arg0)
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getMaxSecondsWithoutAFNHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMTollOffRampHelper.CallOpts)
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "isHealthy", timeNow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _EVM2EVMTollOffRampHelper.Contract.IsHealthy(&_EVM2EVMTollOffRampHelper.CallOpts, timeNow)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _EVM2EVMTollOffRampHelper.Contract.IsHealthy(&_EVM2EVMTollOffRampHelper.CallOpts, timeNow)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) IsPool(addr common.Address) (bool, error) {
	return _EVM2EVMTollOffRampHelper.Contract.IsPool(&_EVM2EVMTollOffRampHelper.CallOpts, addr)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) IsPool(addr common.Address) (bool, error) {
	return _EVM2EVMTollOffRampHelper.Contract.IsPool(&_EVM2EVMTollOffRampHelper.CallOpts, addr)
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) SRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "s_router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SRouter() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SRouter(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) SRouter() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SRouter(&_EVM2EVMTollOffRampHelper.CallOpts)
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessage) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "executeSingleMessage", message)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) ExecuteSingleMessage(message CCIPAny2EVMMessage) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRampHelper.TransactOpts, message)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMMessage) (*types.Transaction, error) {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetConfig(&_EVM2EVMTollOffRampHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetConfig(&_EVM2EVMTollOffRampHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetConfig0(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setConfig0", config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetConfig0(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetConfig0(&_EVM2EVMTollOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetConfig0(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetConfig0(&_EVM2EVMTollOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMTollOffRampHelper.TransactOpts, newTime)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMTollOffRampHelper.TransactOpts, newTime)
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

type EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSetIterator struct {
	Event *EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet)
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
		it.Event = new(EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet)
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

func (it *EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSetIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "AFNMaxHeartbeatTimeSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseAFNMaxHeartbeatTimeSet(log types.Log) (*EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet, error) {
	event := new(EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	case _EVM2EVMTollOffRampHelper.abi.Events["AFNMaxHeartbeatTimeSet"].ID:
		return _EVM2EVMTollOffRampHelper.ParseAFNMaxHeartbeatTimeSet(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["AFNSet"].ID:
		return _EVM2EVMTollOffRampHelper.ParseAFNSet(log)
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
	case _EVM2EVMTollOffRampHelper.abi.Events["Transmitted"].ID:
		return _EVM2EVMTollOffRampHelper.ParseTransmitted(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["Unpaused"].ID:
		return _EVM2EVMTollOffRampHelper.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet) Topic() common.Hash {
	return common.HexToHash("0x72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251")
}

func (EVM2EVMTollOffRampHelperAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMTollOffRampHelperConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMTollOffRampHelperExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf")
}

func (EVM2EVMTollOffRampHelperOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x187b05a44b331bbf26fcef5672820acbcd72e33ec0214a0b89042db7b73b46a2")
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

	SRouter(opts *bind.CallOpts) (common.Address, error)

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessage) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetBlobVerifier(opts *bind.TransactOpts, blobVerifier common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	SetConfig0(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error)

	SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSetIterator, error)

	WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet) (event.Subscription, error)

	ParseAFNMaxHeartbeatTimeSet(log types.Log) (*EVM2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet, error)

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMTollOffRampHelperAFNSet, error)

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

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMTollOffRampHelperTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMTollOffRampHelperUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
