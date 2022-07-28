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
	SourceChainId                           *big.Int
	ExecutionDelaySeconds                   uint64
	MaxDataSize                             uint64
	MaxTokensLength                         uint64
	PermissionLessExecutionThresholdSeconds uint32
}

type CCIPAny2EVMSubscriptionMessage struct {
	SourceChainId  *big.Int
	SequenceNumber uint64
	Sender         common.Address
	Receiver       common.Address
	Nonce          uint64
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

var Any2EVMSubscriptionOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SubscriptionNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMSubscriptionMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMSubscriptionMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlobVerifier\",\"outputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"s_receiverToNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_router\",\"outputs\":[{\"internalType\":\"contractAny2EVMSubscriptionOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"}],\"name\":\"setBlobVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMSubscriptionOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005a9338038062005a9383398101604081905262000034916200058c565b6000805460ff191681556001908990899089908990899089908990899083908390869084903390819081620000b05760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ea57620000ea8162000327565b5050506001600160a01b038216158062000102575080155b156200012157604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001675760405162d8548360e71b815260040160405180910390fd5b81516200017c906005906020850190620003d8565b5060005b82518110156200025e576000828281518110620001a157620001a1620006c7565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001eb57620001eb620006c7565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff191660011790556200025681620006dd565b905062000180565b50508751608090815260a0999099525050855160085550505060208301516009805460408601516060870151979096015163ffffffff16600160c01b0263ffffffff60c01b196001600160401b03988916600160801b0216600160801b600160e01b031997891668010000000000000000026001600160801b0319909316989094169790971717949094161793909317909155600780546001600160a01b039092166001600160a01b031990921691909117905550151560c05250620007059650505050505050565b336001600160a01b03821603620003815760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a7565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b82805482825590600052602060002090810192821562000430579160200282015b828111156200043057825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003f9565b506200043e92915062000442565b5090565b5b808211156200043e576000815560010162000443565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b038111828210171562000494576200049462000459565b60405290565b80516001600160401b0381168114620004b257600080fd5b919050565b6001600160a01b0381168114620004cd57600080fd5b50565b8051620004b281620004b7565b600082601f830112620004ef57600080fd5b815160206001600160401b03808311156200050e576200050e62000459565b8260051b604051601f19603f8301168101818110848211171562000536576200053662000459565b6040529384528581018301938381019250878511156200055557600080fd5b83870191505b84821015620005815781516200057181620004b7565b835291830191908301906200055b565b979650505050505050565b600080600080600080600080888a03610180811215620005ab57600080fd5b8951985060a0601f1982011215620005c257600080fd5b50620005cd6200046f565b60208a01518152620005e260408b016200049a565b6020820152620005f560608b016200049a565b60408201526200060860808b016200049a565b606082015260a08a015163ffffffff811681146200062557600080fd5b608082015296506200063a60c08a01620004d0565b95506200064a60e08a01620004d0565b94506200065b6101008a01620004d0565b6101208a01519094506001600160401b03808211156200067a57600080fd5b620006888c838d01620004dd565b94506101408b0151915080821115620006a057600080fd5b50620006af8b828c01620004dd565b92505061016089015190509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b600060018201620006fe57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c051615350620007436000396000611169015260006104ab0152600081816103de01528181612ef701526136e501526153506000f3fe608060405234801561001057600080fd5b50600436106102415760003560e01c806385e1f4d011610145578063bbe4f6db116100bd578063e16e632c1161008c578063eb511dd411610071578063eb511dd414610701578063efaf6eab14610714578063f2fde38b1461072757600080fd5b8063e16e632c146106ce578063e3d0e712146106ee57600080fd5b8063bbe4f6db14610567578063be9b03f1146105a0578063c0d78655146105b3578063c3f909d4146105c657600080fd5b8063a988980811610114578063b034909c116100f9578063b034909c14610539578063b1dc65a414610541578063b6608c3b1461055457600080fd5b8063a98898081461050b578063afcb95d71461051957600080fd5b806385e1f4d0146104a657806389c06568146104cd5780638da5cb5b146104d5578063a639d1c0146104f857600080fd5b80636133dc24116101d857806379ba5097116101a7578063814118341161018c578063814118341461045957806381ff70481461046e5780638456cb591461049e57600080fd5b806379ba50971461040e5780637c34718c1461041657600080fd5b80636133dc241461039557806372b9d105146103b3578063744b92e2146103c657806374be2150146103d957600080fd5b80633f4ba83a116102145780633f4ba83a14610326578063567c814b1461032e5780635b16ebb7146103515780635c975abb1461038a57600080fd5b8063108ee5fc14610246578063142a98fc1461025b578063181f5a771461029e5780632222dd42146102e7575b600080fd5b610259610254366004613bc2565b61073a565b005b610288610269366004613c07565b67ffffffffffffffff166000908152600a602052604090205460ff1690565b6040516102959190613c53565b60405180910390f35b6102da6040518060400160405280602081526020017f416e793245564d537562736372697074696f6e4f666652616d7020312e302e3081525081565b6040516102959190613d0a565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610295565b610259610816565b61034161033c366004613d1d565b610828565b6040519015158152602001610295565b61034161035f366004613bc2565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff16610341565b60075473ffffffffffffffffffffffffffffffffffffffff16610301565b6102596103c1366004613f97565b61096f565b6102596103d43660046140a8565b6109cb565b6104007f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610295565b610259610dca565b610440610424366004613bc2565b60126020526000908152604090205467ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610295565b610461610ef1565b6040516102959190614132565b600d54600b546040805163ffffffff80851682526401000000009094049093166020840152820152606001610295565b610259610f60565b6104007f000000000000000000000000000000000000000000000000000000000000000081565b610461610f70565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610301565b610259610506366004613bc2565b610fdd565b610259610241366004614145565b604080516001815260006020820181905291810191909152606001610295565b600354610400565b61025961054f3660046141cd565b61102c565b610259610562366004613d1d565b6116d3565b610301610575366004613bc2565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6102596105ae3660046143af565b611753565b6102596105c1366004613bc2565b612251565b6106746040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a081018252600854815260095467ffffffffffffffff808216602084015268010000000000000000820481169383019390935270010000000000000000000000000000000081049092166060820152780100000000000000000000000000000000000000000000000090910463ffffffff16608082015290565b6040516102959190600060a08201905082518252602083015167ffffffffffffffff8082166020850152806040860151166040850152806060860151166060850152505063ffffffff608084015116608083015292915050565b6011546103019073ffffffffffffffffffffffffffffffffffffffff1681565b6102596106fc366004614509565b6122c8565b61025961070f3660046140a8565b612ca3565b6102596107223660046145d6565b612eeb565b610259610735366004613bc2565b6130c0565b6107426130d1565b73ffffffffffffffffffffffffffffffffffffffff811661078f576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b61081e6130d1565b610826613157565b565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa158015610898573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bc919061466e565b1580156109695750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610934573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610958919061468b565b602001516109669084614716565b11155b92915050565b3330146109a8576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109bf8160c001518260e001518360600151613238565b6109c8816132d2565b50565b6109d36130d1565b6005546000819003610a11576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610aac576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610b15576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006005610b24600185614716565b81548110610b3457610b3461472d565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff1681548110610b8657610b8661472d565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166005610bb5600186614716565b81548110610bc557610bc561472d565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110610c3357610c3361472d565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480610cd557610cd561475c565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610e50576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60606010805480602002602001604051908101604052809291908181526020018280548015610f5657602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610f2b575b5050505050905090565b610f686130d1565b6108266133d5565b60606005805480602002602001604051908101604052809291908181526020018280548015610f565760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610f2b575050505050905090565b610fe56130d1565b600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161108291849163ffffffff851691908e908e908190840183828082843760009201919091525061349592505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600b54808252600c5460ff80821660208501526101009091041692820192909252908314611157576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610e47565b6111658b8b8b8b8b8b613525565b60007f0000000000000000000000000000000000000000000000000000000000000000156111c2576002826020015183604001516111a3919061478b565b6111ad91906147df565b6111b890600161478b565b60ff1690506111d8565b60208201516111d290600161478b565b60ff1690505b888114611241576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610e47565b8887146112aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610e47565b336000908152600e602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156112ed576112ed613c24565b60028111156112fe576112fe613c24565b905250905060028160200151600281111561131b5761131b613c24565b14801561136257506010816000015160ff168154811061133d5761133d61472d565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b6113c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610e47565b5050505050600088886040516113df929190614801565b6040519081900381206113f6918c90602001614811565b604051602081830303815290604052805190602001209050611416613b81565b604080518082019091526000808252602082015260005b888110156116b157600060018588846020811061144c5761144c61472d565b61145991901a601b61478b565b8d8d8681811061146b5761146b61472d565b905060200201358c8c878181106114845761148461472d565b90506020020135604051600081526020016040526040516114c1949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156114e3573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600e602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561156357611563613c24565b600281111561157457611574613c24565b905250925060018360200151600281111561159157611591613c24565b146115f8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610e47565b8251849060ff16601f811061160f5761160f61472d565b60200201511561167b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610e47565b600184846000015160ff16601f81106116965761169661472d565b91151560209092020152506116aa8161482d565b905061142d565b5050505063ffffffff81106116c8576116c8614865565b505050505050505050565b6116db6130d1565b80600003611715576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251910161080a565b60005460ff16156117c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610e47565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561182d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611851919061466e565b15611887576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa1580156118f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061191b919061468b565b90506003548160200151426119309190614716565b1115611968576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60115473ffffffffffffffffffffffffffffffffffffffff166119b7576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608301515160008190036119f8576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115611a1357611a13613d41565b604051908082528060200260200182016040528015611a3c578160200160208202803683370190505b50905060008267ffffffffffffffff811115611a5a57611a5a613d41565b604051908082528060200260200182016040528015611aee57816020015b60408051610120810182526000808252602080830182905292820181905260608083018290526080830182905260a0830181905260c0830181905260e083015261010082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff909201910181611a785790505b50905060005b83811015611bc15786606001518181518110611b1257611b1261472d565b6020026020010151806020019051810190611b2d91906149b6565b828281518110611b3f57611b3f61472d565b6020026020010181905250600087606001518281518110611b6257611b6261472d565b6020026020010151604051602001611b7a9190614ac7565b60405160208183030381529060405290508080519060200120848381518110611ba557611ba561472d565b602090810291909101015250611bba8161482d565b9050611af4565b50600080611be28489608001518a60a001518b60c001518c60e001516135dc565b915091506000835182611bf59190614aed565b9050878015611c3057506009547801000000000000000000000000000000000000000000000000900463ffffffff16611c2e8442614716565b105b15611c67576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604089015160005b878110156122445760005a90506000878381518110611c9057611c9061472d565b602002602001015190506000611cc3826020015167ffffffffffffffff166000908152600a602052604090205460ff1690565b90506002816003811115611cd957611cd9613c24565b03611d225760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610e47565b60115460608301516040517f0cbebc2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526000929190911690630cbebc2490602401600060405180830381865afa158015611d9b573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052611de19190810190614b01565b602081015190915073ffffffffffffffffffffffffffffffffffffffff16611e535760608301516040517f8515736a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610e47565b6080830151606084015173ffffffffffffffffffffffffffffffffffffffff16600090815260126020526040812054909167ffffffffffffffff90811691611e9d91166001614bbd565b67ffffffffffffffff161490508080611ed457508160400151158015611ed457506003836003811115611ed257611ed2613c24565b145b611f1c5760808401516040517fb0241f4a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610e47565b611f25846136e1565b60005b8460c0015151811015611f6e57611f5b8560c001518281518110611f4e57611f4e61472d565b6020026020010151613846565b5080611f668161482d565b915050611f28565b5060208481015167ffffffffffffffff166000908152600a9091526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611fbf856138c2565b60208087015167ffffffffffffffff166000908152600a90915260409020805491925082917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600183600381111561201a5761201a613c24565b02179055508180156120495750600281600381111561203b5761203b613c24565b148061204957508260400151155b156120b557606085015173ffffffffffffffffffffffffffffffffffffffff166000908152601260205260408120805467ffffffffffffffff169161208d83614be9565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b8f6121e757601160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663294d266186606001518760400151670de0b6b3a76400008c8c8151811061211d5761211d61472d565b60200260200101513a8f5a612132908f614716565b61213c9190614c10565b6121469190614c28565b6121509190614c28565b61215a9190614aed565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815273ffffffffffffffffffffffffffffffffffffffff93841660048201529290911660248301526044820152606401600060405180830381600087803b1580156121ce57600080fd5b505af11580156121e2573d6000803e3d6000fd5b505050505b846020015167ffffffffffffffff167fbca6416e78a437ab47530846568a4d78457e41bc2adc0d91a826090e2d853d1c826040516122259190613c53565b60405180910390a25050505050508061223d9061482d565b9050611c6f565b5050505050505050505050565b6122596130d1565b601180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b855185518560ff16601f83111561233b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610e47565b600081116123a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610e47565b818314612433576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610e47565b61243e816003614c28565b83116124a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610e47565b6124ae6130d1565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600f54156126a157600f5460009061250690600190614716565b90506000600f828154811061251d5761251d61472d565b60009182526020822001546010805473ffffffffffffffffffffffffffffffffffffffff909216935090849081106125575761255761472d565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600e909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600f805491925090806125d7576125d761475c565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905560108054806126405761264061475c565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055506124ec915050565b60005b815151811015612b06576000600e6000846000015184815181106126ca576126ca61472d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561271457612714613c24565b1461277b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610e47565b6040805180820190915260ff821681526001602082015282518051600e91600091859081106127ac576127ac61472d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561284d5761284d613c24565b02179055506000915061285d9050565b600e6000846020015184815181106128775761287761472d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff1660028111156128c1576128c1613c24565b14612928576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610e47565b6040805180820190915260ff82168152602081016002815250600e60008460200151848151811061295b5761295b61472d565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156129fc576129fc613c24565b02179055505082518051600f925083908110612a1a57612a1a61472d565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9093169290921790915582015180516010919083908110612a9657612a9661472d565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055612aff8161482d565b90506126a4565b506040810151600c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600d80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092612b98928692908216911617614c65565b92506101000a81548163ffffffff021916908363ffffffff160217905550612bf74630600d60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613971565b600b81905582518051600c805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600d5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612c96988b98919763ffffffff909216969095919491939192614c84565b60405180910390a1612244565b612cab6130d1565b73ffffffffffffffffffffffffffffffffffffffff82161580612ce2575073ffffffffffffffffffffffffffffffffffffffff8116155b15612d19576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612db5576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b612ef36130d1565b80517f000000000000000000000000000000000000000000000000000000000000000014612f535780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610e47565b80516008556020810151600980546040808501516060860151608087015163ffffffff167801000000000000000000000000000000000000000000000000027fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff92831670010000000000000000000000000000000002167fffffffff000000000000000000000000ffffffffffffffffffffffffffffffff93831668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009096169290971691909117939093171693909317179055517fedc1b12e6a2ea72b9768b6c0e185d04d9c656f3a270976aa68badc3c1d550902906130b5908390600060a08201905082518252602083015167ffffffffffffffff8082166020850152806040860151166040850152806060860151166060850152505063ffffffff608084015116608083015292915050565b60405180910390a150565b6130c86130d1565b6109c881613a1c565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610826576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610e47565b60005460ff166131c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610e47565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b8151835114613273576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b83518110156132cc576132bc8482815181106132945761329461472d565b60200260200101518483815181106132ae576132ae61472d565b602002602001015184613b17565b6132c58161482d565b9050613276565b50505050565b606081015173ffffffffffffffffffffffffffffffffffffffff163b6133425760608101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610e47565b60608101516011546040517f5b89dece00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690635b89dece9061339f9084908690600401614e3d565b600060405180830381600087803b1580156133b957600080fd5b505af11580156133cd573d6000803e3d6000fd5b505050505050565b60005460ff1615613442576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610e47565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861320e3390565b6000818060200190518101906134ab9190614f50565b6040517fbe9b03f1000000000000000000000000000000000000000000000000000000008152909150309063be9b03f1906134ed908490600190600401615139565b600060405180830381600087803b15801561350757600080fd5b505af115801561351b573d6000803e3d6000fd5b5050505050505050565b6000613532826020614c28565b61353d856020614c28565b61354988610144614c10565b6135539190614c10565b61355d9190614c10565b613568906000614c10565b90503681146135d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610e47565b50505050505050565b60008060005a6007546040517fe71e65ce00000000000000000000000000000000000000000000000000000000815291925060009173ffffffffffffffffffffffffffffffffffffffff9091169063e71e65ce90613646908c908c908c908c908c90600401615230565b6020604051808303816000875af1158015613665573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136899190615282565b9050600081116136c5576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a6136d19084614716565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146137415780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610e47565b60095460c08201515170010000000000000000000000000000000090910467ffffffffffffffff16108061377f57508060e00151518160c001515114155b156137c85760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610e47565b60095460a0820151516801000000000000000090910467ffffffffffffffff1610156109c85760095460a0820151516040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090920467ffffffffffffffff1660048301526024820152604401610e47565b73ffffffffffffffffffffffffffffffffffffffff81811660009081526004602052604090205416806138bd576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610e47565b919050565b6040517f72b9d10500000000000000000000000000000000000000000000000000000000815260009030906372b9d1059061390190859060040161529b565b600060405180830381600087803b15801561391b57600080fd5b505af192505050801561392c575060015b613969573d80801561395a576040519150601f19603f3d011682016040523d82523d6000602084013e61395f565b606091505b5060039392505050565b506002919050565b6000808a8a8a8a8a8a8a8a8a604051602001613995999897969594939291906152ae565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603613a9b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610e47565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000613b2284613846565b6040517fea6192a200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690529192509082169063ea6192a2906044016134ed565b604051806103e00160405280601f906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff811681146109c857600080fd5b600060208284031215613bd457600080fd5b8135613bdf81613ba0565b9392505050565b67ffffffffffffffff811681146109c857600080fd5b80356138bd81613be6565b600060208284031215613c1957600080fd5b8135613bdf81613be6565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310613c8e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b83811015613caf578181015183820152602001613c97565b838111156132cc5750506000910152565b60008151808452613cd8816020860160208601613c94565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613bdf6020830184613cc0565b600060208284031215613d2f57600080fd5b5035919050565b80356138bd81613ba0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff81118282101715613d9457613d94613d41565b60405290565b604051610100810167ffffffffffffffff81118282101715613d9457613d94613d41565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613e0557613e05613d41565b604052919050565b600067ffffffffffffffff821115613e2757613e27613d41565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112613e6457600080fd5b8135613e77613e7282613e0d565b613dbe565b818152846020838601011115613e8c57600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115613ec357613ec3613d41565b5060051b60200190565b600082601f830112613ede57600080fd5b81356020613eee613e7283613ea9565b82815260059290921b84018101918181019086841115613f0d57600080fd5b8286015b84811015613f31578035613f2481613ba0565b8352918301918301613f11565b509695505050505050565b600082601f830112613f4d57600080fd5b81356020613f5d613e7283613ea9565b82815260059290921b84018101918181019086841115613f7c57600080fd5b8286015b84811015613f315780358352918301918301613f80565b600060208284031215613fa957600080fd5b813567ffffffffffffffff80821115613fc157600080fd5b908301906101208286031215613fd657600080fd5b613fde613d70565b82358152613fee60208401613bfc565b6020820152613fff60408401613d36565b604082015261401060608401613d36565b606082015261402160808401613bfc565b608082015260a08301358281111561403857600080fd5b61404487828601613e53565b60a08301525060c08301358281111561405c57600080fd5b61406887828601613ecd565b60c08301525060e08301358281111561408057600080fd5b61408c87828601613f3c565b60e0830152506101009283013592810192909252509392505050565b600080604083850312156140bb57600080fd5b82356140c681613ba0565b915060208301356140d681613ba0565b809150509250929050565b600081518084526020808501945080840160005b8381101561412757815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016140f5565b509495945050505050565b602081526000613bdf60208301846140e1565b60006020828403121561415757600080fd5b813567ffffffffffffffff81111561416e57600080fd5b82016101208185031215613bdf57600080fd5b60008083601f84011261419357600080fd5b50813567ffffffffffffffff8111156141ab57600080fd5b6020830191508360208260051b85010111156141c657600080fd5b9250929050565b60008060008060008060008060e0898b0312156141e957600080fd5b606089018a8111156141fa57600080fd5b8998503567ffffffffffffffff8082111561421457600080fd5b818b0191508b601f83011261422857600080fd5b81358181111561423757600080fd5b8c602082850101111561424957600080fd5b6020830199508098505060808b013591508082111561426757600080fd5b6142738c838d01614181565b909750955060a08b013591508082111561428c57600080fd5b506142998b828c01614181565b999c989b50969995989497949560c00135949350505050565b600082601f8301126142c357600080fd5b813560206142d3613e7283613ea9565b82815260059290921b840181019181810190868411156142f257600080fd5b8286015b84811015613f3157803561430981613be6565b83529183019183016142f6565b600082601f83011261432757600080fd5b81356020614337613e7283613ea9565b82815260059290921b8401810191818101908684111561435657600080fd5b8286015b84811015613f3157803567ffffffffffffffff81111561437a5760008081fd5b6143888986838b0101613e53565b84525091830191830161435a565b80151581146109c857600080fd5b80356138bd81614396565b600080604083850312156143c257600080fd5b823567ffffffffffffffff808211156143da57600080fd5b9084019061010082870312156143ef57600080fd5b6143f7613d9a565b82358281111561440657600080fd5b614412888286016142b2565b82525060208301358281111561442757600080fd5b61443388828601613ecd565b60208301525060408301358281111561444b57600080fd5b61445788828601613f3c565b60408301525060608301358281111561446f57600080fd5b61447b88828601614316565b60608301525060808301358281111561449357600080fd5b61449f88828601613f3c565b60808301525060a083013560a082015260c0830135828111156144c157600080fd5b6144cd88828601613f3c565b60c08301525060e083013560e08201528094505050506144ef602084016143a4565b90509250929050565b803560ff811681146138bd57600080fd5b60008060008060008060c0878903121561452257600080fd5b863567ffffffffffffffff8082111561453a57600080fd5b6145468a838b01613ecd565b9750602089013591508082111561455c57600080fd5b6145688a838b01613ecd565b965061457660408a016144f8565b9550606089013591508082111561458c57600080fd5b6145988a838b01613e53565b94506145a660808a01613bfc565b935060a08901359150808211156145bc57600080fd5b506145c989828a01613e53565b9150509295509295509295565b600060a082840312156145e857600080fd5b60405160a0810181811067ffffffffffffffff8211171561460b5761460b613d41565b60405282358152602083013561462081613be6565b6020820152604083013561463381613be6565b6040820152606083013561464681613be6565b6060820152608083013563ffffffff8116811461466257600080fd5b60808201529392505050565b60006020828403121561468057600080fd5b8151613bdf81614396565b60006060828403121561469d57600080fd5b6040516060810181811067ffffffffffffffff821117156146c0576146c0613d41565b80604052508251815260208301516020820152604083015160408201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015614728576147286146e7565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff038211156147a8576147a86146e7565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff8316806147f2576147f26147b0565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361485e5761485e6146e7565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b80516138bd81613be6565b80516138bd81613ba0565b600082601f8301126148bb57600080fd5b81516148c9613e7282613e0d565b8181528460208386010111156148de57600080fd5b6148ef826020830160208701613c94565b949350505050565b600082601f83011261490857600080fd5b81516020614918613e7283613ea9565b82815260059290921b8401810191818101908684111561493757600080fd5b8286015b84811015613f3157805161494e81613ba0565b835291830191830161493b565b600082601f83011261496c57600080fd5b8151602061497c613e7283613ea9565b82815260059290921b8401810191818101908684111561499b57600080fd5b8286015b84811015613f31578051835291830191830161499f565b6000602082840312156149c857600080fd5b815167ffffffffffffffff808211156149e057600080fd5b9083019061012082860312156149f557600080fd5b6149fd613d70565b82518152614a0d60208401614894565b6020820152614a1e6040840161489f565b6040820152614a2f6060840161489f565b6060820152614a4060808401614894565b608082015260a083015182811115614a5757600080fd5b614a63878286016148aa565b60a08301525060c083015182811115614a7b57600080fd5b614a87878286016148f7565b60c08301525060e083015182811115614a9f57600080fd5b614aab8782860161495b565b60e0830152506101009283015192810192909252509392505050565b6000815260008251614ae0816001850160208701613c94565b9190910160010192915050565b600082614afc57614afc6147b0565b500490565b600060208284031215614b1357600080fd5b815167ffffffffffffffff80821115614b2b57600080fd5b9083019060808286031215614b3f57600080fd5b604051608081018181108382111715614b5a57614b5a613d41565b604052825182811115614b6c57600080fd5b614b78878286016148f7565b82525060208301519150614b8b82613ba0565b81602082015260408301519150614ba182614396565b8160408201526060830151606082015280935050505092915050565b600067ffffffffffffffff808316818516808303821115614be057614be06146e7565b01949350505050565b600067ffffffffffffffff808316818103614c0657614c066146e7565b6001019392505050565b60008219821115614c2357614c236146e7565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614c6057614c606146e7565b500290565b600063ffffffff808316818516808303821115614be057614be06146e7565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614cb48184018a6140e1565b90508281036080840152614cc881896140e1565b905060ff871660a084015282810360c0840152614ce58187613cc0565b905067ffffffffffffffff851660e0840152828103610100840152614d0a8185613cc0565b9c9b505050505050505050505050565b600081518084526020808501945080840160005b8381101561412757815187529582019590820190600101614d2e565b6000610120825184526020830151614d6e602086018267ffffffffffffffff169052565b506040830151614d96604086018273ffffffffffffffffffffffffffffffffffffffff169052565b506060830151614dbe606086018273ffffffffffffffffffffffffffffffffffffffff169052565b506080830151614dda608086018267ffffffffffffffff169052565b5060a08301518160a0860152614df282860182613cc0565b91505060c083015184820360c0860152614e0c82826140e1565b91505060e083015184820360e0860152614e268282614d1a565b610100948501519590940194909452509092915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006148ef6040830184614d4a565b600082601f830112614e7d57600080fd5b81516020614e8d613e7283613ea9565b82815260059290921b84018101918181019086841115614eac57600080fd5b8286015b84811015613f31578051614ec381613be6565b8352918301918301614eb0565b600082601f830112614ee157600080fd5b81516020614ef1613e7283613ea9565b82815260059290921b84018101918181019086841115614f1057600080fd5b8286015b84811015613f3157805167ffffffffffffffff811115614f345760008081fd5b614f428986838b01016148aa565b845250918301918301614f14565b600060208284031215614f6257600080fd5b815167ffffffffffffffff80821115614f7a57600080fd5b908301906101008286031215614f8f57600080fd5b614f97613d9a565b825182811115614fa657600080fd5b614fb287828601614e6c565b825250602083015182811115614fc757600080fd5b614fd3878286016148f7565b602083015250604083015182811115614feb57600080fd5b614ff78782860161495b565b60408301525060608301518281111561500f57600080fd5b61501b87828601614ed0565b60608301525060808301518281111561503357600080fd5b61503f8782860161495b565b60808301525060a083015160a082015260c08301518281111561506157600080fd5b61506d8782860161495b565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b8381101561412757815167ffffffffffffffff168752958201959082019060010161509d565b600082825180855260208086019550808260051b84010181860160005b8481101561512c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086840301895261511a838351613cc0565b988401989250908301906001016150e0565b5090979650505050505050565b6040815260008351610100806040850152615158610140850183615089565b915060208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08086850301606087015261519484836140e1565b935060408801519150808685030160808701526151b18483614d1a565b935060608801519150808685030160a08701526151ce84836150c3565b935060808801519150808685030160c08701526151eb8483614d1a565b935060a088015160e087015260c08801519150808685030183870152506152128382614d1a565b60e088015161012087015286151560208701529350613bdf92505050565b60a08152600061524360a0830188614d1a565b82810360208401526152558188614d1a565b9050856040840152828103606084015261526f8186614d1a565b9150508260808301529695505050505050565b60006020828403121561529457600080fd5b5051919050565b602081526000613bdf6020830184614d4a565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526152f58285018b6140e1565b91508382036080850152615309828a6140e1565b915060ff881660a085015283820360c08501526153268288613cc0565b90861660e08501528381036101008501529050614d0a8185613cc056fea164736f6c634300080f000a",
}

var Any2EVMSubscriptionOffRampABI = Any2EVMSubscriptionOffRampMetaData.ABI

var Any2EVMSubscriptionOffRampBin = Any2EVMSubscriptionOffRampMetaData.Bin

func DeployAny2EVMSubscriptionOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, blobVerifier common.Address, onRampAddress common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, maxTimeWithoutAFNSignal *big.Int) (common.Address, *types.Transaction, *Any2EVMSubscriptionOffRamp, error) {
	parsed, err := Any2EVMSubscriptionOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Any2EVMSubscriptionOffRampBin), backend, chainId, offRampConfig, blobVerifier, onRampAddress, afn, sourceTokens, pools, maxTimeWithoutAFNSignal)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Any2EVMSubscriptionOffRamp{Any2EVMSubscriptionOffRampCaller: Any2EVMSubscriptionOffRampCaller{contract: contract}, Any2EVMSubscriptionOffRampTransactor: Any2EVMSubscriptionOffRampTransactor{contract: contract}, Any2EVMSubscriptionOffRampFilterer: Any2EVMSubscriptionOffRampFilterer{contract: contract}}, nil
}

type Any2EVMSubscriptionOffRamp struct {
	address common.Address
	abi     abi.ABI
	Any2EVMSubscriptionOffRampCaller
	Any2EVMSubscriptionOffRampTransactor
	Any2EVMSubscriptionOffRampFilterer
}

type Any2EVMSubscriptionOffRampCaller struct {
	contract *bind.BoundContract
}

type Any2EVMSubscriptionOffRampTransactor struct {
	contract *bind.BoundContract
}

type Any2EVMSubscriptionOffRampFilterer struct {
	contract *bind.BoundContract
}

type Any2EVMSubscriptionOffRampSession struct {
	Contract     *Any2EVMSubscriptionOffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type Any2EVMSubscriptionOffRampCallerSession struct {
	Contract *Any2EVMSubscriptionOffRampCaller
	CallOpts bind.CallOpts
}

type Any2EVMSubscriptionOffRampTransactorSession struct {
	Contract     *Any2EVMSubscriptionOffRampTransactor
	TransactOpts bind.TransactOpts
}

type Any2EVMSubscriptionOffRampRaw struct {
	Contract *Any2EVMSubscriptionOffRamp
}

type Any2EVMSubscriptionOffRampCallerRaw struct {
	Contract *Any2EVMSubscriptionOffRampCaller
}

type Any2EVMSubscriptionOffRampTransactorRaw struct {
	Contract *Any2EVMSubscriptionOffRampTransactor
}

func NewAny2EVMSubscriptionOffRamp(address common.Address, backend bind.ContractBackend) (*Any2EVMSubscriptionOffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(Any2EVMSubscriptionOffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindAny2EVMSubscriptionOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRamp{address: address, abi: abi, Any2EVMSubscriptionOffRampCaller: Any2EVMSubscriptionOffRampCaller{contract: contract}, Any2EVMSubscriptionOffRampTransactor: Any2EVMSubscriptionOffRampTransactor{contract: contract}, Any2EVMSubscriptionOffRampFilterer: Any2EVMSubscriptionOffRampFilterer{contract: contract}}, nil
}

func NewAny2EVMSubscriptionOffRampCaller(address common.Address, caller bind.ContractCaller) (*Any2EVMSubscriptionOffRampCaller, error) {
	contract, err := bindAny2EVMSubscriptionOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampCaller{contract: contract}, nil
}

func NewAny2EVMSubscriptionOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*Any2EVMSubscriptionOffRampTransactor, error) {
	contract, err := bindAny2EVMSubscriptionOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampTransactor{contract: contract}, nil
}

func NewAny2EVMSubscriptionOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*Any2EVMSubscriptionOffRampFilterer, error) {
	contract, err := bindAny2EVMSubscriptionOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampFilterer{contract: contract}, nil
}

func bindAny2EVMSubscriptionOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Any2EVMSubscriptionOffRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Any2EVMSubscriptionOffRamp.Contract.Any2EVMSubscriptionOffRampCaller.contract.Call(opts, result, method, params...)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Any2EVMSubscriptionOffRampTransactor.contract.Transfer(opts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Any2EVMSubscriptionOffRampTransactor.contract.Transact(opts, method, params...)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Any2EVMSubscriptionOffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.contract.Transfer(opts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) CHAINID() (*big.Int, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.CHAINID(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) CHAINID() (*big.Int, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.CHAINID(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "SOURCE_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) SOURCECHAINID() (*big.Int, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SOURCECHAINID(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) SOURCECHAINID() (*big.Int, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SOURCECHAINID(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMSubscriptionMessage) error {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) CcipReceive(arg0 CCIPAny2EVMSubscriptionMessage) error {
	return _Any2EVMSubscriptionOffRamp.Contract.CcipReceive(&_Any2EVMSubscriptionOffRamp.CallOpts, arg0)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) CcipReceive(arg0 CCIPAny2EVMSubscriptionMessage) error {
	return _Any2EVMSubscriptionOffRamp.Contract.CcipReceive(&_Any2EVMSubscriptionOffRamp.CallOpts, arg0)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) GetAFN() (common.Address, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.GetAFN(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) GetAFN() (common.Address, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.GetAFN(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) GetBlobVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getBlobVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) GetBlobVerifier() (common.Address, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.GetBlobVerifier(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) GetBlobVerifier() (common.Address, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.GetBlobVerifier(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOffRampInterfaceOffRampConfig)).(*BaseOffRampInterfaceOffRampConfig)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.GetConfig(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.GetConfig(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.GetExecutionState(&_Any2EVMSubscriptionOffRamp.CallOpts, sequenceNumber)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.GetExecutionState(&_Any2EVMSubscriptionOffRamp.CallOpts, sequenceNumber)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getMaxSecondsWithoutAFNHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.GetPool(&_Any2EVMSubscriptionOffRamp.CallOpts, sourceToken)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.GetPool(&_Any2EVMSubscriptionOffRamp.CallOpts, sourceToken)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) GetPoolTokens() ([]common.Address, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.GetPoolTokens(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.GetPoolTokens(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "isHealthy", timeNow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.IsHealthy(&_Any2EVMSubscriptionOffRamp.CallOpts, timeNow)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.IsHealthy(&_Any2EVMSubscriptionOffRamp.CallOpts, timeNow)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) IsPool(addr common.Address) (bool, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.IsPool(&_Any2EVMSubscriptionOffRamp.CallOpts, addr)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) IsPool(addr common.Address) (bool, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.IsPool(&_Any2EVMSubscriptionOffRamp.CallOpts, addr)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _Any2EVMSubscriptionOffRamp.Contract.LatestConfigDetails(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _Any2EVMSubscriptionOffRamp.Contract.LatestConfigDetails(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _Any2EVMSubscriptionOffRamp.Contract.LatestConfigDigestAndEpoch(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _Any2EVMSubscriptionOffRamp.Contract.LatestConfigDigestAndEpoch(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) Owner() (common.Address, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Owner(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) Owner() (common.Address, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Owner(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) Paused() (bool, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Paused(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) Paused() (bool, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Paused(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) SReceiverToNonce(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "s_receiverToNonce", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) SReceiverToNonce(arg0 common.Address) (uint64, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SReceiverToNonce(&_Any2EVMSubscriptionOffRamp.CallOpts, arg0)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) SReceiverToNonce(arg0 common.Address) (uint64, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SReceiverToNonce(&_Any2EVMSubscriptionOffRamp.CallOpts, arg0)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) SRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "s_router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) SRouter() (common.Address, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SRouter(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) SRouter() (common.Address, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SRouter(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) Transmitters() ([]common.Address, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Transmitters(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) Transmitters() ([]common.Address, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Transmitters(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Any2EVMSubscriptionOffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) TypeAndVersion() (string, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.TypeAndVersion(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampCallerSession) TypeAndVersion() (string, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.TypeAndVersion(&_Any2EVMSubscriptionOffRamp.CallOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.AcceptOwnership(&_Any2EVMSubscriptionOffRamp.TransactOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.AcceptOwnership(&_Any2EVMSubscriptionOffRamp.TransactOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.AddPool(&_Any2EVMSubscriptionOffRamp.TransactOpts, token, pool)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.AddPool(&_Any2EVMSubscriptionOffRamp.TransactOpts, token, pool)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) Execute(opts *bind.TransactOpts, report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "execute", report, manualExecution)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Execute(&_Any2EVMSubscriptionOffRamp.TransactOpts, report, manualExecution)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Execute(&_Any2EVMSubscriptionOffRamp.TransactOpts, report, manualExecution)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMSubscriptionMessage) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "executeSingleMessage", message)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) ExecuteSingleMessage(message CCIPAny2EVMSubscriptionMessage) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.ExecuteSingleMessage(&_Any2EVMSubscriptionOffRamp.TransactOpts, message)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMSubscriptionMessage) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.ExecuteSingleMessage(&_Any2EVMSubscriptionOffRamp.TransactOpts, message)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "pause")
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) Pause() (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Pause(&_Any2EVMSubscriptionOffRamp.TransactOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) Pause() (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Pause(&_Any2EVMSubscriptionOffRamp.TransactOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.RemovePool(&_Any2EVMSubscriptionOffRamp.TransactOpts, token, pool)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.RemovePool(&_Any2EVMSubscriptionOffRamp.TransactOpts, token, pool)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "setAFN", afn)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SetAFN(&_Any2EVMSubscriptionOffRamp.TransactOpts, afn)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SetAFN(&_Any2EVMSubscriptionOffRamp.TransactOpts, afn)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) SetBlobVerifier(opts *bind.TransactOpts, blobVerifier common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "setBlobVerifier", blobVerifier)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SetBlobVerifier(&_Any2EVMSubscriptionOffRamp.TransactOpts, blobVerifier)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SetBlobVerifier(&_Any2EVMSubscriptionOffRamp.TransactOpts, blobVerifier)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SetConfig(&_Any2EVMSubscriptionOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SetConfig(&_Any2EVMSubscriptionOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) SetConfig0(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "setConfig0", config)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) SetConfig0(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SetConfig0(&_Any2EVMSubscriptionOffRamp.TransactOpts, config)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) SetConfig0(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SetConfig0(&_Any2EVMSubscriptionOffRamp.TransactOpts, config)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_Any2EVMSubscriptionOffRamp.TransactOpts, newTime)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_Any2EVMSubscriptionOffRamp.TransactOpts, newTime)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "setRouter", router)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SetRouter(&_Any2EVMSubscriptionOffRamp.TransactOpts, router)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.SetRouter(&_Any2EVMSubscriptionOffRamp.TransactOpts, router)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.TransferOwnership(&_Any2EVMSubscriptionOffRamp.TransactOpts, to)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.TransferOwnership(&_Any2EVMSubscriptionOffRamp.TransactOpts, to)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Transmit(&_Any2EVMSubscriptionOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Transmit(&_Any2EVMSubscriptionOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.contract.Transact(opts, "unpause")
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampSession) Unpause() (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Unpause(&_Any2EVMSubscriptionOffRamp.TransactOpts)
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _Any2EVMSubscriptionOffRamp.Contract.Unpause(&_Any2EVMSubscriptionOffRamp.TransactOpts)
}

type Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSetIterator struct {
	Event *Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet)
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
		it.Event = new(Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet)
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

func (it *Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSetIterator, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSetIterator{contract: _Any2EVMSubscriptionOffRamp.contract, event: "AFNMaxHeartbeatTimeSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet) (event.Subscription, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet)
				if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) ParseAFNMaxHeartbeatTimeSet(log types.Log) (*Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet, error) {
	event := new(Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet)
	if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMSubscriptionOffRampAFNSetIterator struct {
	Event *Any2EVMSubscriptionOffRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMSubscriptionOffRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMSubscriptionOffRampAFNSet)
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
		it.Event = new(Any2EVMSubscriptionOffRampAFNSet)
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

func (it *Any2EVMSubscriptionOffRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMSubscriptionOffRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMSubscriptionOffRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampAFNSetIterator, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampAFNSetIterator{contract: _Any2EVMSubscriptionOffRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMSubscriptionOffRampAFNSet)
				if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) ParseAFNSet(log types.Log) (*Any2EVMSubscriptionOffRampAFNSet, error) {
	event := new(Any2EVMSubscriptionOffRampAFNSet)
	if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMSubscriptionOffRampConfigSetIterator struct {
	Event *Any2EVMSubscriptionOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMSubscriptionOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMSubscriptionOffRampConfigSet)
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
		it.Event = new(Any2EVMSubscriptionOffRampConfigSet)
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

func (it *Any2EVMSubscriptionOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMSubscriptionOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMSubscriptionOffRampConfigSet struct {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampConfigSetIterator, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampConfigSetIterator{contract: _Any2EVMSubscriptionOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMSubscriptionOffRampConfigSet)
				if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) ParseConfigSet(log types.Log) (*Any2EVMSubscriptionOffRampConfigSet, error) {
	event := new(Any2EVMSubscriptionOffRampConfigSet)
	if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMSubscriptionOffRampExecutionCompletedIterator struct {
	Event *Any2EVMSubscriptionOffRampExecutionCompleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMSubscriptionOffRampExecutionCompletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMSubscriptionOffRampExecutionCompleted)
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
		it.Event = new(Any2EVMSubscriptionOffRampExecutionCompleted)
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

func (it *Any2EVMSubscriptionOffRampExecutionCompletedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMSubscriptionOffRampExecutionCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMSubscriptionOffRampExecutionCompleted struct {
	SequenceNumber uint64
	State          uint8
	Raw            types.Log
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) FilterExecutionCompleted(opts *bind.FilterOpts, sequenceNumber []uint64) (*Any2EVMSubscriptionOffRampExecutionCompletedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "ExecutionCompleted", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampExecutionCompletedIterator{contract: _Any2EVMSubscriptionOffRamp.contract, event: "ExecutionCompleted", logs: logs, sub: sub}, nil
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampExecutionCompleted, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "ExecutionCompleted", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMSubscriptionOffRampExecutionCompleted)
				if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "ExecutionCompleted", log); err != nil {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) ParseExecutionCompleted(log types.Log) (*Any2EVMSubscriptionOffRampExecutionCompleted, error) {
	event := new(Any2EVMSubscriptionOffRampExecutionCompleted)
	if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "ExecutionCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMSubscriptionOffRampOffRampConfigSetIterator struct {
	Event *Any2EVMSubscriptionOffRampOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMSubscriptionOffRampOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMSubscriptionOffRampOffRampConfigSet)
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
		it.Event = new(Any2EVMSubscriptionOffRampOffRampConfigSet)
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

func (it *Any2EVMSubscriptionOffRampOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMSubscriptionOffRampOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMSubscriptionOffRampOffRampConfigSet struct {
	Config BaseOffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampOffRampConfigSetIterator, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampOffRampConfigSetIterator{contract: _Any2EVMSubscriptionOffRamp.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMSubscriptionOffRampOffRampConfigSet)
				if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) ParseOffRampConfigSet(log types.Log) (*Any2EVMSubscriptionOffRampOffRampConfigSet, error) {
	event := new(Any2EVMSubscriptionOffRampOffRampConfigSet)
	if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMSubscriptionOffRampOffRampRouterSetIterator struct {
	Event *Any2EVMSubscriptionOffRampOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMSubscriptionOffRampOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMSubscriptionOffRampOffRampRouterSet)
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
		it.Event = new(Any2EVMSubscriptionOffRampOffRampRouterSet)
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

func (it *Any2EVMSubscriptionOffRampOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMSubscriptionOffRampOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMSubscriptionOffRampOffRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*Any2EVMSubscriptionOffRampOffRampRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampOffRampRouterSetIterator{contract: _Any2EVMSubscriptionOffRamp.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMSubscriptionOffRampOffRampRouterSet)
				if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) ParseOffRampRouterSet(log types.Log) (*Any2EVMSubscriptionOffRampOffRampRouterSet, error) {
	event := new(Any2EVMSubscriptionOffRampOffRampRouterSet)
	if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMSubscriptionOffRampOwnershipTransferRequestedIterator struct {
	Event *Any2EVMSubscriptionOffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMSubscriptionOffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMSubscriptionOffRampOwnershipTransferRequested)
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
		it.Event = new(Any2EVMSubscriptionOffRampOwnershipTransferRequested)
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

func (it *Any2EVMSubscriptionOffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMSubscriptionOffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMSubscriptionOffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMSubscriptionOffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampOwnershipTransferRequestedIterator{contract: _Any2EVMSubscriptionOffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMSubscriptionOffRampOwnershipTransferRequested)
				if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*Any2EVMSubscriptionOffRampOwnershipTransferRequested, error) {
	event := new(Any2EVMSubscriptionOffRampOwnershipTransferRequested)
	if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMSubscriptionOffRampOwnershipTransferredIterator struct {
	Event *Any2EVMSubscriptionOffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMSubscriptionOffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMSubscriptionOffRampOwnershipTransferred)
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
		it.Event = new(Any2EVMSubscriptionOffRampOwnershipTransferred)
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

func (it *Any2EVMSubscriptionOffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *Any2EVMSubscriptionOffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMSubscriptionOffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMSubscriptionOffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampOwnershipTransferredIterator{contract: _Any2EVMSubscriptionOffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMSubscriptionOffRampOwnershipTransferred)
				if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) ParseOwnershipTransferred(log types.Log) (*Any2EVMSubscriptionOffRampOwnershipTransferred, error) {
	event := new(Any2EVMSubscriptionOffRampOwnershipTransferred)
	if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMSubscriptionOffRampPausedIterator struct {
	Event *Any2EVMSubscriptionOffRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMSubscriptionOffRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMSubscriptionOffRampPaused)
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
		it.Event = new(Any2EVMSubscriptionOffRampPaused)
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

func (it *Any2EVMSubscriptionOffRampPausedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMSubscriptionOffRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMSubscriptionOffRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) FilterPaused(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampPausedIterator, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampPausedIterator{contract: _Any2EVMSubscriptionOffRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampPaused) (event.Subscription, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMSubscriptionOffRampPaused)
				if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) ParsePaused(log types.Log) (*Any2EVMSubscriptionOffRampPaused, error) {
	event := new(Any2EVMSubscriptionOffRampPaused)
	if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMSubscriptionOffRampPoolAddedIterator struct {
	Event *Any2EVMSubscriptionOffRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMSubscriptionOffRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMSubscriptionOffRampPoolAdded)
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
		it.Event = new(Any2EVMSubscriptionOffRampPoolAdded)
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

func (it *Any2EVMSubscriptionOffRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMSubscriptionOffRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMSubscriptionOffRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampPoolAddedIterator, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampPoolAddedIterator{contract: _Any2EVMSubscriptionOffRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMSubscriptionOffRampPoolAdded)
				if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) ParsePoolAdded(log types.Log) (*Any2EVMSubscriptionOffRampPoolAdded, error) {
	event := new(Any2EVMSubscriptionOffRampPoolAdded)
	if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMSubscriptionOffRampPoolRemovedIterator struct {
	Event *Any2EVMSubscriptionOffRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMSubscriptionOffRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMSubscriptionOffRampPoolRemoved)
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
		it.Event = new(Any2EVMSubscriptionOffRampPoolRemoved)
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

func (it *Any2EVMSubscriptionOffRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMSubscriptionOffRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMSubscriptionOffRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampPoolRemovedIterator, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampPoolRemovedIterator{contract: _Any2EVMSubscriptionOffRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMSubscriptionOffRampPoolRemoved)
				if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) ParsePoolRemoved(log types.Log) (*Any2EVMSubscriptionOffRampPoolRemoved, error) {
	event := new(Any2EVMSubscriptionOffRampPoolRemoved)
	if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMSubscriptionOffRampTransmittedIterator struct {
	Event *Any2EVMSubscriptionOffRampTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMSubscriptionOffRampTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMSubscriptionOffRampTransmitted)
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
		it.Event = new(Any2EVMSubscriptionOffRampTransmitted)
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

func (it *Any2EVMSubscriptionOffRampTransmittedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMSubscriptionOffRampTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMSubscriptionOffRampTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) FilterTransmitted(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampTransmittedIterator, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampTransmittedIterator{contract: _Any2EVMSubscriptionOffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampTransmitted) (event.Subscription, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMSubscriptionOffRampTransmitted)
				if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) ParseTransmitted(log types.Log) (*Any2EVMSubscriptionOffRampTransmitted, error) {
	event := new(Any2EVMSubscriptionOffRampTransmitted)
	if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMSubscriptionOffRampUnpausedIterator struct {
	Event *Any2EVMSubscriptionOffRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMSubscriptionOffRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMSubscriptionOffRampUnpaused)
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
		it.Event = new(Any2EVMSubscriptionOffRampUnpaused)
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

func (it *Any2EVMSubscriptionOffRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMSubscriptionOffRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMSubscriptionOffRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampUnpausedIterator, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &Any2EVMSubscriptionOffRampUnpausedIterator{contract: _Any2EVMSubscriptionOffRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _Any2EVMSubscriptionOffRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMSubscriptionOffRampUnpaused)
				if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRampFilterer) ParseUnpaused(log types.Log) (*Any2EVMSubscriptionOffRampUnpaused, error) {
	event := new(Any2EVMSubscriptionOffRampUnpaused)
	if err := _Any2EVMSubscriptionOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Any2EVMSubscriptionOffRamp.abi.Events["AFNMaxHeartbeatTimeSet"].ID:
		return _Any2EVMSubscriptionOffRamp.ParseAFNMaxHeartbeatTimeSet(log)
	case _Any2EVMSubscriptionOffRamp.abi.Events["AFNSet"].ID:
		return _Any2EVMSubscriptionOffRamp.ParseAFNSet(log)
	case _Any2EVMSubscriptionOffRamp.abi.Events["ConfigSet"].ID:
		return _Any2EVMSubscriptionOffRamp.ParseConfigSet(log)
	case _Any2EVMSubscriptionOffRamp.abi.Events["ExecutionCompleted"].ID:
		return _Any2EVMSubscriptionOffRamp.ParseExecutionCompleted(log)
	case _Any2EVMSubscriptionOffRamp.abi.Events["OffRampConfigSet"].ID:
		return _Any2EVMSubscriptionOffRamp.ParseOffRampConfigSet(log)
	case _Any2EVMSubscriptionOffRamp.abi.Events["OffRampRouterSet"].ID:
		return _Any2EVMSubscriptionOffRamp.ParseOffRampRouterSet(log)
	case _Any2EVMSubscriptionOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _Any2EVMSubscriptionOffRamp.ParseOwnershipTransferRequested(log)
	case _Any2EVMSubscriptionOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _Any2EVMSubscriptionOffRamp.ParseOwnershipTransferred(log)
	case _Any2EVMSubscriptionOffRamp.abi.Events["Paused"].ID:
		return _Any2EVMSubscriptionOffRamp.ParsePaused(log)
	case _Any2EVMSubscriptionOffRamp.abi.Events["PoolAdded"].ID:
		return _Any2EVMSubscriptionOffRamp.ParsePoolAdded(log)
	case _Any2EVMSubscriptionOffRamp.abi.Events["PoolRemoved"].ID:
		return _Any2EVMSubscriptionOffRamp.ParsePoolRemoved(log)
	case _Any2EVMSubscriptionOffRamp.abi.Events["Transmitted"].ID:
		return _Any2EVMSubscriptionOffRamp.ParseTransmitted(log)
	case _Any2EVMSubscriptionOffRamp.abi.Events["Unpaused"].ID:
		return _Any2EVMSubscriptionOffRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet) Topic() common.Hash {
	return common.HexToHash("0x72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251")
}

func (Any2EVMSubscriptionOffRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (Any2EVMSubscriptionOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (Any2EVMSubscriptionOffRampExecutionCompleted) Topic() common.Hash {
	return common.HexToHash("0xbca6416e78a437ab47530846568a4d78457e41bc2adc0d91a826090e2d853d1c")
}

func (Any2EVMSubscriptionOffRampOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xedc1b12e6a2ea72b9768b6c0e185d04d9c656f3a270976aa68badc3c1d550902")
}

func (Any2EVMSubscriptionOffRampOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4")
}

func (Any2EVMSubscriptionOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (Any2EVMSubscriptionOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (Any2EVMSubscriptionOffRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (Any2EVMSubscriptionOffRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (Any2EVMSubscriptionOffRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (Any2EVMSubscriptionOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (Any2EVMSubscriptionOffRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_Any2EVMSubscriptionOffRamp *Any2EVMSubscriptionOffRamp) Address() common.Address {
	return _Any2EVMSubscriptionOffRamp.address
}

type Any2EVMSubscriptionOffRampInterface interface {
	CHAINID(opts *bind.CallOpts) (*big.Int, error)

	SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error)

	CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMSubscriptionMessage) error

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetBlobVerifier(opts *bind.CallOpts) (common.Address, error)

	GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error)

	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

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

	ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMSubscriptionMessage) (*types.Transaction, error)

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

	FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSetIterator, error)

	WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet) (event.Subscription, error)

	ParseAFNMaxHeartbeatTimeSet(log types.Log) (*Any2EVMSubscriptionOffRampAFNMaxHeartbeatTimeSet, error)

	FilterAFNSet(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*Any2EVMSubscriptionOffRampAFNSet, error)

	FilterConfigSet(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*Any2EVMSubscriptionOffRampConfigSet, error)

	FilterExecutionCompleted(opts *bind.FilterOpts, sequenceNumber []uint64) (*Any2EVMSubscriptionOffRampExecutionCompletedIterator, error)

	WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampExecutionCompleted, sequenceNumber []uint64) (event.Subscription, error)

	ParseExecutionCompleted(log types.Log) (*Any2EVMSubscriptionOffRampExecutionCompleted, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*Any2EVMSubscriptionOffRampOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*Any2EVMSubscriptionOffRampOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*Any2EVMSubscriptionOffRampOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMSubscriptionOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*Any2EVMSubscriptionOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMSubscriptionOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*Any2EVMSubscriptionOffRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*Any2EVMSubscriptionOffRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*Any2EVMSubscriptionOffRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*Any2EVMSubscriptionOffRampPoolRemoved, error)

	FilterTransmitted(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*Any2EVMSubscriptionOffRampTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*Any2EVMSubscriptionOffRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Any2EVMSubscriptionOffRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*Any2EVMSubscriptionOffRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
