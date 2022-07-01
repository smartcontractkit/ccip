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

type CCIPAny2EVMTollMessage struct {
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

var Any2EVMTollOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlobVerifier\",\"outputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMTollOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"}],\"name\":\"setBlobVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200584d3803806200584d83398101604081905262000034916200058c565b6000805460ff191681556001908990899089908990899089908990899083908390869084903390819081620000b05760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ea57620000ea8162000327565b5050506001600160a01b038216158062000102575080155b156200012157604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001675760405162d8548360e71b815260040160405180910390fd5b81516200017c906005906020850190620003d8565b5060005b82518110156200025e576000828281518110620001a157620001a1620006c7565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001eb57620001eb620006c7565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff191660011790556200025681620006dd565b905062000180565b50508751608090815260a0999099525050855160085550505060208301516009805460408601516060870151979096015163ffffffff16600160c01b0263ffffffff60c01b196001600160401b03988916600160801b0216600160801b600160e01b031997891668010000000000000000026001600160801b0319909316989094169790971717949094161793909317909155600780546001600160a01b039092166001600160a01b031990921691909117905550151560c05250620007059650505050505050565b336001600160a01b03821603620003815760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a7565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b82805482825590600052602060002090810192821562000430579160200282015b828111156200043057825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003f9565b506200043e92915062000442565b5090565b5b808211156200043e576000815560010162000443565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b038111828210171562000494576200049462000459565b60405290565b80516001600160401b0381168114620004b257600080fd5b919050565b6001600160a01b0381168114620004cd57600080fd5b50565b8051620004b281620004b7565b600082601f830112620004ef57600080fd5b815160206001600160401b03808311156200050e576200050e62000459565b8260051b604051601f19603f8301168101818110848211171562000536576200053662000459565b6040529384528581018301938381019250878511156200055557600080fd5b83870191505b84821015620005815781516200057181620004b7565b835291830191908301906200055b565b979650505050505050565b600080600080600080600080888a03610180811215620005ab57600080fd5b8951985060a0601f1982011215620005c257600080fd5b50620005cd6200046f565b60208a01518152620005e260408b016200049a565b6020820152620005f560608b016200049a565b60408201526200060860808b016200049a565b606082015260a08a015163ffffffff811681146200062557600080fd5b608082015296506200063a60c08a01620004d0565b95506200064a60e08a01620004d0565b94506200065b6101008a01620004d0565b6101208a01519094506001600160401b03808211156200067a57600080fd5b620006888c838d01620004dd565b94506101408b0151915080821115620006a057600080fd5b50620006af8b828c01620004dd565b92505061016089015190509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b600060018201620006fe57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05161510a6200074360003960006111090152600061044d0152600081816103c301528181612cc4015261349f015261510a6000f3fe608060405234801561001057600080fd5b50600436106102265760003560e01c806389c065681161012a578063b6608c3b116100bd578063c3f909d41161008c578063eb511dd411610071578063eb511dd4146106a1578063efaf6eab146106b4578063f2fde38b146106c757600080fd5b8063c3f909d414610586578063e3d0e7121461068e57600080fd5b8063b6608c3b14610514578063bbe4f6db14610527578063be9b03f114610560578063c0d786551461057357600080fd5b8063afcb95d7116100f9578063afcb95d7146104bb578063b034909c146104db578063b0f479a1146104e3578063b1dc65a41461050157600080fd5b806389c065681461046f5780638bbad066146104775780638da5cb5b14610485578063a639d1c0146104a857600080fd5b80635c975abb116101bd57806379ba50971161018c57806381ff70481161017157806381ff7048146104105780638456cb591461044057806385e1f4d01461044857600080fd5b806379ba5097146103f357806381411834146103fb57600080fd5b80635c975abb146103825780636133dc241461038d578063744b92e2146103ab57806374be2150146103be57600080fd5b80632222dd42116101f95780632222dd42146102df5780633f4ba83a1461031e578063567c814b146103265780635b16ebb71461034957600080fd5b8063092cddc21461022b578063108ee5fc14610240578063142a98fc14610253578063181f5a7714610296575b600080fd5b61023e610239366004613c31565b6106da565b005b61023e61024e366004613d4d565b610736565b610280610261366004613d71565b67ffffffffffffffff166000908152600a602052604090205460ff1690565b60405161028d9190613dbd565b60405180910390f35b6102d26040518060400160405280601881526020017f416e793245564d546f6c6c4f666652616d7020312e302e30000000000000000081525081565b60405161028d9190613e74565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161028d565b61023e610812565b610339610334366004613e87565b610824565b604051901515815260200161028d565b610339610357366004613d4d565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff16610339565b60075473ffffffffffffffffffffffffffffffffffffffff166102f9565b61023e6103b9366004613ea0565b61096b565b6103e57f000000000000000000000000000000000000000000000000000000000000000081565b60405190815260200161028d565b61023e610d6a565b610403610e91565b60405161028d9190613f2a565b600d54600b546040805163ffffffff8085168252640100000000909404909316602084015282015260600161028d565b61023e610f00565b6103e57f000000000000000000000000000000000000000000000000000000000000000081565b610403610f10565b61023e610226366004613f3d565b600054610100900473ffffffffffffffffffffffffffffffffffffffff166102f9565b61023e6104b6366004613d4d565b610f7d565b60408051600181526000602082018190529181019190915260600161028d565b6003546103e5565b60115473ffffffffffffffffffffffffffffffffffffffff166102f9565b61023e61050f366004613fc5565b610fcc565b61023e610522366004613e87565b611673565b6102f9610535366004613d4d565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b61023e61056e36600461420b565b6116f3565b61023e610581366004613d4d565b612016565b6106346040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a081018252600854815260095467ffffffffffffffff808216602084015268010000000000000000820481169383019390935270010000000000000000000000000000000081049092166060820152780100000000000000000000000000000000000000000000000090910463ffffffff16608082015290565b60405161028d9190600060a08201905082518252602083015167ffffffffffffffff8082166020850152806040860151166040850152806060860151166060850152505063ffffffff608084015116608083015292915050565b61023e61069c366004614365565b61208d565b61023e6106af366004613ea0565b612a70565b61023e6106c2366004614432565b612cb8565b61023e6106d5366004613d4d565b612e8d565b333014610713576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61072a8160a001518260c001518360600151612e9e565b61073381612f38565b50565b61073e61303b565b73ffffffffffffffffffffffffffffffffffffffff811661078b576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b61081a61303b565b6108226130c1565b565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa158015610894573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b891906144ca565b1580156109655750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610930573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095491906144e7565b602001516109629084614572565b11155b92915050565b61097361303b565b60055460008190036109b1576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610a4c576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610ab5576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006005610ac4600185614572565b81548110610ad457610ad4614589565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff1681548110610b2657610b26614589565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166005610b55600186614572565b81548110610b6557610b65614589565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110610bd357610bd3614589565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480610c7557610c756145b8565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610df0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60606010805480602002602001604051908101604052809291908181526020018280548015610ef657602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610ecb575b5050505050905090565b610f0861303b565b6108226131a2565b60606005805480602002602001604051908101604052809291908181526020018280548015610ef65760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610ecb575050505050905090565b610f8561303b565b600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161102291849163ffffffff851691908e908e908190840183828082843760009201919091525061326292505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600b54808252600c5460ff808216602085015261010090910416928201929092529083146110f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610de7565b6111058b8b8b8b8b8b6132e8565b60007f0000000000000000000000000000000000000000000000000000000000000000156111625760028260200151836040015161114391906145e7565b61114d919061463b565b6111589060016145e7565b60ff169050611178565b60208201516111729060016145e7565b60ff1690505b8881146111e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610de7565b88871461124a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610de7565b336000908152600e602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561128d5761128d613d8e565b600281111561129e5761129e613d8e565b90525090506002816020015160028111156112bb576112bb613d8e565b14801561130257506010816000015160ff16815481106112dd576112dd614589565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b611368576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610de7565b50505050506000888860405161137f92919061465d565b604051908190038120611396918c9060200161466d565b6040516020818303038152906040528051906020012090506113b661396e565b604080518082019091526000808252602082015260005b888110156116515760006001858884602081106113ec576113ec614589565b6113f991901a601b6145e7565b8d8d8681811061140b5761140b614589565b905060200201358c8c8781811061142457611424614589565b9050602002013560405160008152602001604052604051611461949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611483573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600e602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561150357611503613d8e565b600281111561151457611514613d8e565b905250925060018360200151600281111561153157611531613d8e565b14611598576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610de7565b8251849060ff16601f81106115af576115af614589565b60200201511561161b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610de7565b600184846000015160ff16601f811061163657611636614589565b911515602090920201525061164a81614689565b90506113cd565b5050505063ffffffff8110611668576116686146c1565b505050505050505050565b61167b61303b565b806000036116b5576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c2519101610806565b60005460ff1615611760576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610de7565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117f191906144ca565b15611827576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa158015611897573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118bb91906144e7565b90506003548160200151426118d09190614572565b1115611908576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60115473ffffffffffffffffffffffffffffffffffffffff16611957576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060830151516000819003611998576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff8111156119b3576119b361398d565b6040519080825280602002602001820160405280156119dc578160200160208202803683370190505b50905060008267ffffffffffffffff8111156119fa576119fa61398d565b604051908082528060200260200182016040528015611acd57816020015b611aba60405180610140016040528060008152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b815260200190600190039081611a185790505b50905060005b83811015611ba05786606001518181518110611af157611af1614589565b6020026020010151806020019051810190611b0c9190614812565b828281518110611b1e57611b1e614589565b6020026020010181905250600087606001518281518110611b4157611b41614589565b6020026020010151604051602001611b59919061492e565b60405160208183030381529060405290508080519060200120848381518110611b8457611b84614589565b602090810291909101015250611b9981614689565b9050611ad3565b50600080611bc18489608001518a60a001518b60c001518c60e00151613396565b915091506000835182611bd49190614954565b9050878015611c0f57506009547801000000000000000000000000000000000000000000000000900463ffffffff16611c0d8442614572565b105b15611c46576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8681101561200a576000858281518110611c6557611c65614589565b602002602001015190506000611c98826020015167ffffffffffffffff166000908152600a602052604090205460ff1690565b90506002816003811115611cae57611cae613d8e565b03611cf75760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610de7565b611d008261349b565b60005b8260a0015151811015611d4957611d368360a001518281518110611d2957611d29614589565b6020026020010151613600565b5080611d4181614689565b915050611d03565b506000816003811115611d5e57611d5e613d8e565b148015611d6957508a155b15611f00576000805b8d6020015151811015611e04578360e0015173ffffffffffffffffffffffffffffffffffffffff168e602001518281518110611db057611db0614589565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1603611df4578d604001518181518110611de957611de9614589565b602002602001015191505b611dfd81614689565b9050611d72565b5080611e5a5760e08301516040517fce480bcc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610de7565b6000670de0b6b3a7640000823a86610120015189611e789190614968565b611e829190614980565b611e8c9190614980565b611e969190614954565b9050836101000151811115611ee95760208401516040517f6b830fc700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610de7565b611efd8460e001518561010001513061367c565b50505b60208281015167ffffffffffffffff166000908152600a9091526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611f5083613719565b60208085015167ffffffffffffffff166000908152600a90915260409020805491925082917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001836003811115611fab57611fab613d8e565b0217905550826020015167ffffffffffffffff167fbca6416e78a437ab47530846568a4d78457e41bc2adc0d91a826090e2d853d1c82604051611fee9190613dbd565b60405180910390a25050508061200390614689565b9050611c49565b50505050505050505050565b61201e61303b565b601180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b855185518560ff16601f831115612100576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610de7565b6000811161216a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610de7565b8183146121f8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610de7565b612203816003614980565b831161226b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610de7565b61227361303b565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600f541561246657600f546000906122cb90600190614572565b90506000600f82815481106122e2576122e2614589565b60009182526020822001546010805473ffffffffffffffffffffffffffffffffffffffff9092169350908490811061231c5761231c614589565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600e909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600f8054919250908061239c5761239c6145b8565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690550190556010805480612405576124056145b8565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055506122b1915050565b60005b8151518110156128cb576000600e60008460000151848151811061248f5761248f614589565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff1660028111156124d9576124d9613d8e565b14612540576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610de7565b6040805180820190915260ff821681526001602082015282518051600e916000918590811061257157612571614589565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561261257612612613d8e565b0217905550600091506126229050565b600e60008460200151848151811061263c5761263c614589565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561268657612686613d8e565b146126ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610de7565b6040805180820190915260ff82168152602081016002815250600e60008460200151848151811061272057612720614589565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156127c1576127c1613d8e565b02179055505082518051600f9250839081106127df576127df614589565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051601091908390811061285b5761285b614589565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790556128c481614689565b9050612469565b506040810151600c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600d80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff43811682029283178555908304811693600193909260009261295d9286929082169116176149bd565b92506101000a81548163ffffffff021916908363ffffffff1602179055506129bc4630600d60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a001516137c8565b600b81905582518051600c805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600d5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612a5b988b98919763ffffffff9092169690959194919391926149e5565b60405180910390a15050505050505050505050565b612a7861303b565b73ffffffffffffffffffffffffffffffffffffffff82161580612aaf575073ffffffffffffffffffffffffffffffffffffffff8116155b15612ae6576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612b82576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b612cc061303b565b80517f000000000000000000000000000000000000000000000000000000000000000014612d205780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610de7565b80516008556020810151600980546040808501516060860151608087015163ffffffff167801000000000000000000000000000000000000000000000000027fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff92831670010000000000000000000000000000000002167fffffffff000000000000000000000000ffffffffffffffffffffffffffffffff93831668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009096169290971691909117939093171693909317179055517fedc1b12e6a2ea72b9768b6c0e185d04d9c656f3a270976aa68badc3c1d55090290612e82908390600060a08201905082518252602083015167ffffffffffffffff8082166020850152806040860151166040850152806060860151166060850152505063ffffffff608084015116608083015292915050565b60405180910390a150565b612e9561303b565b61073381613873565b8151835114612ed9576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015612f3257612f22848281518110612efa57612efa614589565b6020026020010151848381518110612f1457612f14614589565b60200260200101518461367c565b612f2b81614689565b9050612edc565b50505050565b606081015173ffffffffffffffffffffffffffffffffffffffff163b612fa85760608101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610de7565b60608101516011546040517ffd12f6e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063fd12f6e5906130059084908690600401614bb4565b600060405180830381600087803b15801561301f57600080fd5b505af1158015613033573d6000803e3d6000fd5b505050505050565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610822576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610de7565b60005460ff1661312d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610de7565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60005460ff161561320f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610de7565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586131783390565b3073ffffffffffffffffffffffffffffffffffffffff1663be9b03f1828060200190518101906132929190614d2b565b60006040518363ffffffff1660e01b81526004016132b1929190614ef3565b600060405180830381600087803b1580156132cb57600080fd5b505af11580156132df573d6000803e3d6000fd5b50505050505050565b60006132f5826020614980565b613300856020614980565b61330c88610144614968565b6133169190614968565b6133209190614968565b61332b906000614968565b90503681146132df576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610de7565b60008060005a6007546040517fe71e65ce00000000000000000000000000000000000000000000000000000000815291925060009173ffffffffffffffffffffffffffffffffffffffff9091169063e71e65ce90613400908c908c908c908c908c90600401614fea565b6020604051808303816000875af115801561341f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613443919061503c565b90506000811161347f576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a61348b9084614572565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146134fb5780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610de7565b60095460a08201515170010000000000000000000000000000000090910467ffffffffffffffff16108061353957508060c00151518160a001515114155b156135825760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610de7565b6009546080820151516801000000000000000090910467ffffffffffffffff161015610733576009546080820151516040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090920467ffffffffffffffff1660048301526024820152604401610de7565b73ffffffffffffffffffffffffffffffffffffffff8181166000908152600460205260409020541680613677576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610de7565b919050565b600061368784613600565b6040517fea6192a200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690529192509082169063ea6192a290604401600060405180830381600087803b1580156136fb57600080fd5b505af115801561370f573d6000803e3d6000fd5b5050505050505050565b6040517f092cddc2000000000000000000000000000000000000000000000000000000008152600090309063092cddc290613758908590600401615055565b600060405180830381600087803b15801561377257600080fd5b505af1925050508015613783575060015b6137c0573d8080156137b1576040519150601f19603f3d011682016040523d82523d6000602084013e6137b6565b606091505b5060039392505050565b506002919050565b6000808a8a8a8a8a8a8a8a8a6040516020016137ec99989796959493929190615068565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036138f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610de7565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604051806103e00160405280601f906020820280368337509192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610140810167ffffffffffffffff811182821017156139e0576139e061398d565b60405290565b604051610100810167ffffffffffffffff811182821017156139e0576139e061398d565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613a5157613a5161398d565b604052919050565b67ffffffffffffffff8116811461073357600080fd5b803561367781613a59565b73ffffffffffffffffffffffffffffffffffffffff8116811461073357600080fd5b803561367781613a7a565b600067ffffffffffffffff821115613ac157613ac161398d565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112613afe57600080fd5b8135613b11613b0c82613aa7565b613a0a565b818152846020838601011115613b2657600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115613b5d57613b5d61398d565b5060051b60200190565b600082601f830112613b7857600080fd5b81356020613b88613b0c83613b43565b82815260059290921b84018101918181019086841115613ba757600080fd5b8286015b84811015613bcb578035613bbe81613a7a565b8352918301918301613bab565b509695505050505050565b600082601f830112613be757600080fd5b81356020613bf7613b0c83613b43565b82815260059290921b84018101918181019086841115613c1657600080fd5b8286015b84811015613bcb5780358352918301918301613c1a565b600060208284031215613c4357600080fd5b813567ffffffffffffffff80821115613c5b57600080fd5b908301906101408286031215613c7057600080fd5b613c786139bc565b82358152613c8860208401613a6f565b6020820152613c9960408401613a9c565b6040820152613caa60608401613a9c565b6060820152608083013582811115613cc157600080fd5b613ccd87828601613aed565b60808301525060a083013582811115613ce557600080fd5b613cf187828601613b67565b60a08301525060c083013582811115613d0957600080fd5b613d1587828601613bd6565b60c083015250613d2760e08401613a9c565b60e082015261010083810135908201526101209283013592810192909252509392505050565b600060208284031215613d5f57600080fd5b8135613d6a81613a7a565b9392505050565b600060208284031215613d8357600080fd5b8135613d6a81613a59565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310613df8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b83811015613e19578181015183820152602001613e01565b83811115612f325750506000910152565b60008151808452613e42816020860160208601613dfe565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613d6a6020830184613e2a565b600060208284031215613e9957600080fd5b5035919050565b60008060408385031215613eb357600080fd5b8235613ebe81613a7a565b91506020830135613ece81613a7a565b809150509250929050565b600081518084526020808501945080840160005b83811015613f1f57815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101613eed565b509495945050505050565b602081526000613d6a6020830184613ed9565b600060208284031215613f4f57600080fd5b813567ffffffffffffffff811115613f6657600080fd5b82016101408185031215613d6a57600080fd5b60008083601f840112613f8b57600080fd5b50813567ffffffffffffffff811115613fa357600080fd5b6020830191508360208260051b8501011115613fbe57600080fd5b9250929050565b60008060008060008060008060e0898b031215613fe157600080fd5b606089018a811115613ff257600080fd5b8998503567ffffffffffffffff8082111561400c57600080fd5b818b0191508b601f83011261402057600080fd5b81358181111561402f57600080fd5b8c602082850101111561404157600080fd5b6020830199508098505060808b013591508082111561405f57600080fd5b61406b8c838d01613f79565b909750955060a08b013591508082111561408457600080fd5b506140918b828c01613f79565b999c989b50969995989497949560c00135949350505050565b600082601f8301126140bb57600080fd5b813560206140cb613b0c83613b43565b82815260059290921b840181019181810190868411156140ea57600080fd5b8286015b84811015613bcb57803561410181613a59565b83529183019183016140ee565b600082601f83011261411f57600080fd5b8135602061412f613b0c83613b43565b82815260059290921b8401810191818101908684111561414e57600080fd5b8286015b84811015613bcb57803561416581613a7a565b8352918301918301614152565b600082601f83011261418357600080fd5b81356020614193613b0c83613b43565b82815260059290921b840181019181810190868411156141b257600080fd5b8286015b84811015613bcb57803567ffffffffffffffff8111156141d65760008081fd5b6141e48986838b0101613aed565b8452509183019183016141b6565b801515811461073357600080fd5b8035613677816141f2565b6000806040838503121561421e57600080fd5b823567ffffffffffffffff8082111561423657600080fd5b90840190610100828703121561424b57600080fd5b6142536139e6565b82358281111561426257600080fd5b61426e888286016140aa565b82525060208301358281111561428357600080fd5b61428f8882860161410e565b6020830152506040830135828111156142a757600080fd5b6142b388828601613bd6565b6040830152506060830135828111156142cb57600080fd5b6142d788828601614172565b6060830152506080830135828111156142ef57600080fd5b6142fb88828601613bd6565b60808301525060a083013560a082015260c08301358281111561431d57600080fd5b61432988828601613bd6565b60c08301525060e083013560e082015280945050505061434b60208401614200565b90509250929050565b803560ff8116811461367757600080fd5b60008060008060008060c0878903121561437e57600080fd5b863567ffffffffffffffff8082111561439657600080fd5b6143a28a838b0161410e565b975060208901359150808211156143b857600080fd5b6143c48a838b0161410e565b96506143d260408a01614354565b955060608901359150808211156143e857600080fd5b6143f48a838b01613aed565b945061440260808a01613a6f565b935060a089013591508082111561441857600080fd5b5061442589828a01613aed565b9150509295509295509295565b600060a0828403121561444457600080fd5b60405160a0810181811067ffffffffffffffff821117156144675761446761398d565b60405282358152602083013561447c81613a59565b6020820152604083013561448f81613a59565b604082015260608301356144a281613a59565b6060820152608083013563ffffffff811681146144be57600080fd5b60808201529392505050565b6000602082840312156144dc57600080fd5b8151613d6a816141f2565b6000606082840312156144f957600080fd5b6040516060810181811067ffffffffffffffff8211171561451c5761451c61398d565b80604052508251815260208301516020820152604083015160408201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561458457614584614543565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff0382111561460457614604614543565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff83168061464e5761464e61460c565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036146ba576146ba614543565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b805161367781613a59565b805161367781613a7a565b600082601f83011261471757600080fd5b8151614725613b0c82613aa7565b81815284602083860101111561473a57600080fd5b61474b826020830160208701613dfe565b949350505050565b600082601f83011261476457600080fd5b81516020614774613b0c83613b43565b82815260059290921b8401810191818101908684111561479357600080fd5b8286015b84811015613bcb5780516147aa81613a7a565b8352918301918301614797565b600082601f8301126147c857600080fd5b815160206147d8613b0c83613b43565b82815260059290921b840181019181810190868411156147f757600080fd5b8286015b84811015613bcb57805183529183019183016147fb565b60006020828403121561482457600080fd5b815167ffffffffffffffff8082111561483c57600080fd5b90830190610140828603121561485157600080fd5b6148596139bc565b82518152614869602084016146f0565b602082015261487a604084016146fb565b604082015261488b606084016146fb565b60608201526080830151828111156148a257600080fd5b6148ae87828601614706565b60808301525060a0830151828111156148c657600080fd5b6148d287828601614753565b60a08301525060c0830151828111156148ea57600080fd5b6148f6878286016147b7565b60c08301525061490860e084016146fb565b60e082015261010083810151908201526101209283015192810192909252509392505050565b6000815260008251614947816001850160208701613dfe565b9190910160010192915050565b6000826149635761496361460c565b500490565b6000821982111561497b5761497b614543565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156149b8576149b8614543565b500290565b600063ffffffff8083168185168083038211156149dc576149dc614543565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614a158184018a613ed9565b90508281036080840152614a298189613ed9565b905060ff871660a084015282810360c0840152614a468187613e2a565b905067ffffffffffffffff851660e0840152828103610100840152614a6b8185613e2a565b9c9b505050505050505050505050565b600081518084526020808501945080840160005b83811015613f1f57815187529582019590820190600101614a8f565b6000610140825184526020830151614acf602086018267ffffffffffffffff169052565b506040830151614af7604086018273ffffffffffffffffffffffffffffffffffffffff169052565b506060830151614b1f606086018273ffffffffffffffffffffffffffffffffffffffff169052565b506080830151816080860152614b3782860182613e2a565b91505060a083015184820360a0860152614b518282613ed9565b91505060c083015184820360c0860152614b6b8282614a7b565b91505060e0830151614b9560e086018273ffffffffffffffffffffffffffffffffffffffff169052565b5061010083810151908501526101209283015192909301919091525090565b73ffffffffffffffffffffffffffffffffffffffff8316815260406020820152600061474b6040830184614aab565b600082601f830112614bf457600080fd5b81516020614c04613b0c83613b43565b82815260059290921b84018101918181019086841115614c2357600080fd5b8286015b84811015613bcb578051614c3a81613a59565b8352918301918301614c27565b600082601f830112614c5857600080fd5b81516020614c68613b0c83613b43565b82815260059290921b84018101918181019086841115614c8757600080fd5b8286015b84811015613bcb578051614c9e81613a7a565b8352918301918301614c8b565b600082601f830112614cbc57600080fd5b81516020614ccc613b0c83613b43565b82815260059290921b84018101918181019086841115614ceb57600080fd5b8286015b84811015613bcb57805167ffffffffffffffff811115614d0f5760008081fd5b614d1d8986838b0101614706565b845250918301918301614cef565b600060208284031215614d3d57600080fd5b815167ffffffffffffffff80821115614d5557600080fd5b908301906101008286031215614d6a57600080fd5b614d726139e6565b825182811115614d8157600080fd5b614d8d87828601614be3565b825250602083015182811115614da257600080fd5b614dae87828601614c47565b602083015250604083015182811115614dc657600080fd5b614dd2878286016147b7565b604083015250606083015182811115614dea57600080fd5b614df687828601614cab565b606083015250608083015182811115614e0e57600080fd5b614e1a878286016147b7565b60808301525060a083015160a082015260c083015182811115614e3c57600080fd5b614e48878286016147b7565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b83811015613f1f57815167ffffffffffffffff1687529582019590820190600101614e78565b600081518084526020808501808196508360051b8101915082860160005b85811015614ee6578284038952614ed4848351613e2a565b98850198935090840190600101614ebc565b5091979650505050505050565b6040815260008351610100806040850152614f12610140850183614e64565b915060208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080868503016060870152614f4e8483613ed9565b93506040880151915080868503016080870152614f6b8483614a7b565b935060608801519150808685030160a0870152614f888483614e9e565b935060808801519150808685030160c0870152614fa58483614a7b565b935060a088015160e087015260c0880151915080868503018387015250614fcc8382614a7b565b60e088015161012087015286151560208701529350613d6a92505050565b60a081526000614ffd60a0830188614a7b565b828103602084015261500f8188614a7b565b905085604084015282810360608401526150298186614a7b565b9150508260808301529695505050505050565b60006020828403121561504e57600080fd5b5051919050565b602081526000613d6a6020830184614aab565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526150af8285018b613ed9565b915083820360808501526150c3828a613ed9565b915060ff881660a085015283820360c08501526150e08288613e2a565b90861660e08501528381036101008501529050614a6b8185613e2a56fea164736f6c634300080f000a",
}

var Any2EVMTollOffRampABI = Any2EVMTollOffRampMetaData.ABI

var Any2EVMTollOffRampBin = Any2EVMTollOffRampMetaData.Bin

func DeployAny2EVMTollOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, blobVerifier common.Address, onRampAddress common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, maxTimeWithoutAFNSignal *big.Int) (common.Address, *types.Transaction, *Any2EVMTollOffRamp, error) {
	parsed, err := Any2EVMTollOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Any2EVMTollOffRampBin), backend, chainId, offRampConfig, blobVerifier, onRampAddress, afn, sourceTokens, pools, maxTimeWithoutAFNSignal)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Any2EVMTollOffRamp{Any2EVMTollOffRampCaller: Any2EVMTollOffRampCaller{contract: contract}, Any2EVMTollOffRampTransactor: Any2EVMTollOffRampTransactor{contract: contract}, Any2EVMTollOffRampFilterer: Any2EVMTollOffRampFilterer{contract: contract}}, nil
}

type Any2EVMTollOffRamp struct {
	address common.Address
	abi     abi.ABI
	Any2EVMTollOffRampCaller
	Any2EVMTollOffRampTransactor
	Any2EVMTollOffRampFilterer
}

type Any2EVMTollOffRampCaller struct {
	contract *bind.BoundContract
}

type Any2EVMTollOffRampTransactor struct {
	contract *bind.BoundContract
}

type Any2EVMTollOffRampFilterer struct {
	contract *bind.BoundContract
}

type Any2EVMTollOffRampSession struct {
	Contract     *Any2EVMTollOffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type Any2EVMTollOffRampCallerSession struct {
	Contract *Any2EVMTollOffRampCaller
	CallOpts bind.CallOpts
}

type Any2EVMTollOffRampTransactorSession struct {
	Contract     *Any2EVMTollOffRampTransactor
	TransactOpts bind.TransactOpts
}

type Any2EVMTollOffRampRaw struct {
	Contract *Any2EVMTollOffRamp
}

type Any2EVMTollOffRampCallerRaw struct {
	Contract *Any2EVMTollOffRampCaller
}

type Any2EVMTollOffRampTransactorRaw struct {
	Contract *Any2EVMTollOffRampTransactor
}

func NewAny2EVMTollOffRamp(address common.Address, backend bind.ContractBackend) (*Any2EVMTollOffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(Any2EVMTollOffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindAny2EVMTollOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRamp{address: address, abi: abi, Any2EVMTollOffRampCaller: Any2EVMTollOffRampCaller{contract: contract}, Any2EVMTollOffRampTransactor: Any2EVMTollOffRampTransactor{contract: contract}, Any2EVMTollOffRampFilterer: Any2EVMTollOffRampFilterer{contract: contract}}, nil
}

func NewAny2EVMTollOffRampCaller(address common.Address, caller bind.ContractCaller) (*Any2EVMTollOffRampCaller, error) {
	contract, err := bindAny2EVMTollOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampCaller{contract: contract}, nil
}

func NewAny2EVMTollOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*Any2EVMTollOffRampTransactor, error) {
	contract, err := bindAny2EVMTollOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampTransactor{contract: contract}, nil
}

func NewAny2EVMTollOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*Any2EVMTollOffRampFilterer, error) {
	contract, err := bindAny2EVMTollOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampFilterer{contract: contract}, nil
}

func bindAny2EVMTollOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Any2EVMTollOffRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Any2EVMTollOffRamp.Contract.Any2EVMTollOffRampCaller.contract.Call(opts, result, method, params...)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.Any2EVMTollOffRampTransactor.contract.Transfer(opts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.Any2EVMTollOffRampTransactor.contract.Transact(opts, method, params...)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Any2EVMTollOffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.contract.Transfer(opts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) CHAINID() (*big.Int, error) {
	return _Any2EVMTollOffRamp.Contract.CHAINID(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) CHAINID() (*big.Int, error) {
	return _Any2EVMTollOffRamp.Contract.CHAINID(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "SOURCE_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) SOURCECHAINID() (*big.Int, error) {
	return _Any2EVMTollOffRamp.Contract.SOURCECHAINID(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) SOURCECHAINID() (*big.Int, error) {
	return _Any2EVMTollOffRamp.Contract.SOURCECHAINID(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMTollMessage) error {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) CcipReceive(arg0 CCIPAny2EVMTollMessage) error {
	return _Any2EVMTollOffRamp.Contract.CcipReceive(&_Any2EVMTollOffRamp.CallOpts, arg0)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) CcipReceive(arg0 CCIPAny2EVMTollMessage) error {
	return _Any2EVMTollOffRamp.Contract.CcipReceive(&_Any2EVMTollOffRamp.CallOpts, arg0)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) GetAFN() (common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.GetAFN(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) GetAFN() (common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.GetAFN(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) GetBlobVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "getBlobVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) GetBlobVerifier() (common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.GetBlobVerifier(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) GetBlobVerifier() (common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.GetBlobVerifier(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOffRampInterfaceOffRampConfig)).(*BaseOffRampInterfaceOffRampConfig)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _Any2EVMTollOffRamp.Contract.GetConfig(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _Any2EVMTollOffRamp.Contract.GetConfig(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _Any2EVMTollOffRamp.Contract.GetExecutionState(&_Any2EVMTollOffRamp.CallOpts, sequenceNumber)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _Any2EVMTollOffRamp.Contract.GetExecutionState(&_Any2EVMTollOffRamp.CallOpts, sequenceNumber)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "getMaxSecondsWithoutAFNHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _Any2EVMTollOffRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _Any2EVMTollOffRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.GetPool(&_Any2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.GetPool(&_Any2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) GetPoolTokens() ([]common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.GetPoolTokens(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.GetPoolTokens(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) GetRouter() (common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.GetRouter(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) GetRouter() (common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.GetRouter(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "isHealthy", timeNow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _Any2EVMTollOffRamp.Contract.IsHealthy(&_Any2EVMTollOffRamp.CallOpts, timeNow)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _Any2EVMTollOffRamp.Contract.IsHealthy(&_Any2EVMTollOffRamp.CallOpts, timeNow)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) IsPool(addr common.Address) (bool, error) {
	return _Any2EVMTollOffRamp.Contract.IsPool(&_Any2EVMTollOffRamp.CallOpts, addr)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) IsPool(addr common.Address) (bool, error) {
	return _Any2EVMTollOffRamp.Contract.IsPool(&_Any2EVMTollOffRamp.CallOpts, addr)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _Any2EVMTollOffRamp.Contract.LatestConfigDetails(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _Any2EVMTollOffRamp.Contract.LatestConfigDetails(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _Any2EVMTollOffRamp.Contract.LatestConfigDigestAndEpoch(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _Any2EVMTollOffRamp.Contract.LatestConfigDigestAndEpoch(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) Owner() (common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.Owner(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) Owner() (common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.Owner(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) Paused() (bool, error) {
	return _Any2EVMTollOffRamp.Contract.Paused(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) Paused() (bool, error) {
	return _Any2EVMTollOffRamp.Contract.Paused(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) Transmitters() ([]common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.Transmitters(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) Transmitters() ([]common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.Transmitters(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) TypeAndVersion() (string, error) {
	return _Any2EVMTollOffRamp.Contract.TypeAndVersion(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) TypeAndVersion() (string, error) {
	return _Any2EVMTollOffRamp.Contract.TypeAndVersion(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.AcceptOwnership(&_Any2EVMTollOffRamp.TransactOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.AcceptOwnership(&_Any2EVMTollOffRamp.TransactOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.AddPool(&_Any2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.AddPool(&_Any2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) Execute(opts *bind.TransactOpts, report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "execute", report, manualExecution)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.Execute(&_Any2EVMTollOffRamp.TransactOpts, report, manualExecution)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.Execute(&_Any2EVMTollOffRamp.TransactOpts, report, manualExecution)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMTollMessage) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "executeSingleMessage", message)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) ExecuteSingleMessage(message CCIPAny2EVMTollMessage) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.ExecuteSingleMessage(&_Any2EVMTollOffRamp.TransactOpts, message)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMTollMessage) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.ExecuteSingleMessage(&_Any2EVMTollOffRamp.TransactOpts, message)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "pause")
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) Pause() (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.Pause(&_Any2EVMTollOffRamp.TransactOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) Pause() (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.Pause(&_Any2EVMTollOffRamp.TransactOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.RemovePool(&_Any2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.RemovePool(&_Any2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "setAFN", afn)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.SetAFN(&_Any2EVMTollOffRamp.TransactOpts, afn)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.SetAFN(&_Any2EVMTollOffRamp.TransactOpts, afn)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) SetBlobVerifier(opts *bind.TransactOpts, blobVerifier common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "setBlobVerifier", blobVerifier)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.SetBlobVerifier(&_Any2EVMTollOffRamp.TransactOpts, blobVerifier)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.SetBlobVerifier(&_Any2EVMTollOffRamp.TransactOpts, blobVerifier)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.SetConfig(&_Any2EVMTollOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.SetConfig(&_Any2EVMTollOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) SetConfig0(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "setConfig0", config)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) SetConfig0(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.SetConfig0(&_Any2EVMTollOffRamp.TransactOpts, config)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) SetConfig0(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.SetConfig0(&_Any2EVMTollOffRamp.TransactOpts, config)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_Any2EVMTollOffRamp.TransactOpts, newTime)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_Any2EVMTollOffRamp.TransactOpts, newTime)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "setRouter", router)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.SetRouter(&_Any2EVMTollOffRamp.TransactOpts, router)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.SetRouter(&_Any2EVMTollOffRamp.TransactOpts, router)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.TransferOwnership(&_Any2EVMTollOffRamp.TransactOpts, to)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.TransferOwnership(&_Any2EVMTollOffRamp.TransactOpts, to)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.Transmit(&_Any2EVMTollOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.Transmit(&_Any2EVMTollOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "unpause")
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) Unpause() (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.Unpause(&_Any2EVMTollOffRamp.TransactOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.Unpause(&_Any2EVMTollOffRamp.TransactOpts)
}

type Any2EVMTollOffRampAFNMaxHeartbeatTimeSetIterator struct {
	Event *Any2EVMTollOffRampAFNMaxHeartbeatTimeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampAFNMaxHeartbeatTimeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampAFNMaxHeartbeatTimeSet)
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
		it.Event = new(Any2EVMTollOffRampAFNMaxHeartbeatTimeSet)
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

func (it *Any2EVMTollOffRampAFNMaxHeartbeatTimeSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampAFNMaxHeartbeatTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampAFNMaxHeartbeatTimeSet struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampAFNMaxHeartbeatTimeSetIterator, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampAFNMaxHeartbeatTimeSetIterator{contract: _Any2EVMTollOffRamp.contract, event: "AFNMaxHeartbeatTimeSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampAFNMaxHeartbeatTimeSet) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampAFNMaxHeartbeatTimeSet)
				if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) ParseAFNMaxHeartbeatTimeSet(log types.Log) (*Any2EVMTollOffRampAFNMaxHeartbeatTimeSet, error) {
	event := new(Any2EVMTollOffRampAFNMaxHeartbeatTimeSet)
	if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampAFNSetIterator struct {
	Event *Any2EVMTollOffRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampAFNSet)
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
		it.Event = new(Any2EVMTollOffRampAFNSet)
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

func (it *Any2EVMTollOffRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampAFNSetIterator, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampAFNSetIterator{contract: _Any2EVMTollOffRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampAFNSet)
				if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) ParseAFNSet(log types.Log) (*Any2EVMTollOffRampAFNSet, error) {
	event := new(Any2EVMTollOffRampAFNSet)
	if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampConfigSetIterator struct {
	Event *Any2EVMTollOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampConfigSet)
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
		it.Event = new(Any2EVMTollOffRampConfigSet)
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

func (it *Any2EVMTollOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampConfigSet struct {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampConfigSetIterator, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampConfigSetIterator{contract: _Any2EVMTollOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampConfigSet)
				if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) ParseConfigSet(log types.Log) (*Any2EVMTollOffRampConfigSet, error) {
	event := new(Any2EVMTollOffRampConfigSet)
	if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampExecutionCompletedIterator struct {
	Event *Any2EVMTollOffRampExecutionCompleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampExecutionCompletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampExecutionCompleted)
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
		it.Event = new(Any2EVMTollOffRampExecutionCompleted)
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

func (it *Any2EVMTollOffRampExecutionCompletedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampExecutionCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampExecutionCompleted struct {
	SequenceNumber uint64
	State          uint8
	Raw            types.Log
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterExecutionCompleted(opts *bind.FilterOpts, sequenceNumber []uint64) (*Any2EVMTollOffRampExecutionCompletedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "ExecutionCompleted", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampExecutionCompletedIterator{contract: _Any2EVMTollOffRamp.contract, event: "ExecutionCompleted", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampExecutionCompleted, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "ExecutionCompleted", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampExecutionCompleted)
				if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "ExecutionCompleted", log); err != nil {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) ParseExecutionCompleted(log types.Log) (*Any2EVMTollOffRampExecutionCompleted, error) {
	event := new(Any2EVMTollOffRampExecutionCompleted)
	if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "ExecutionCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampOffRampConfigSetIterator struct {
	Event *Any2EVMTollOffRampOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampOffRampConfigSet)
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
		it.Event = new(Any2EVMTollOffRampOffRampConfigSet)
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

func (it *Any2EVMTollOffRampOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampOffRampConfigSet struct {
	Config BaseOffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampOffRampConfigSetIterator, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampOffRampConfigSetIterator{contract: _Any2EVMTollOffRamp.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampOffRampConfigSet)
				if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) ParseOffRampConfigSet(log types.Log) (*Any2EVMTollOffRampOffRampConfigSet, error) {
	event := new(Any2EVMTollOffRampOffRampConfigSet)
	if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampOffRampRouterSetIterator struct {
	Event *Any2EVMTollOffRampOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampOffRampRouterSet)
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
		it.Event = new(Any2EVMTollOffRampOffRampRouterSet)
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

func (it *Any2EVMTollOffRampOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampOffRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*Any2EVMTollOffRampOffRampRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampOffRampRouterSetIterator{contract: _Any2EVMTollOffRamp.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampOffRampRouterSet)
				if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) ParseOffRampRouterSet(log types.Log) (*Any2EVMTollOffRampOffRampRouterSet, error) {
	event := new(Any2EVMTollOffRampOffRampRouterSet)
	if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampOwnershipTransferRequestedIterator struct {
	Event *Any2EVMTollOffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampOwnershipTransferRequested)
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
		it.Event = new(Any2EVMTollOffRampOwnershipTransferRequested)
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

func (it *Any2EVMTollOffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampOwnershipTransferRequestedIterator{contract: _Any2EVMTollOffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampOwnershipTransferRequested)
				if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*Any2EVMTollOffRampOwnershipTransferRequested, error) {
	event := new(Any2EVMTollOffRampOwnershipTransferRequested)
	if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampOwnershipTransferredIterator struct {
	Event *Any2EVMTollOffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampOwnershipTransferred)
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
		it.Event = new(Any2EVMTollOffRampOwnershipTransferred)
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

func (it *Any2EVMTollOffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampOwnershipTransferredIterator{contract: _Any2EVMTollOffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampOwnershipTransferred)
				if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) ParseOwnershipTransferred(log types.Log) (*Any2EVMTollOffRampOwnershipTransferred, error) {
	event := new(Any2EVMTollOffRampOwnershipTransferred)
	if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampPausedIterator struct {
	Event *Any2EVMTollOffRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampPaused)
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
		it.Event = new(Any2EVMTollOffRampPaused)
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

func (it *Any2EVMTollOffRampPausedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterPaused(opts *bind.FilterOpts) (*Any2EVMTollOffRampPausedIterator, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampPausedIterator{contract: _Any2EVMTollOffRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampPaused) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampPaused)
				if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) ParsePaused(log types.Log) (*Any2EVMTollOffRampPaused, error) {
	event := new(Any2EVMTollOffRampPaused)
	if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampPoolAddedIterator struct {
	Event *Any2EVMTollOffRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampPoolAdded)
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
		it.Event = new(Any2EVMTollOffRampPoolAdded)
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

func (it *Any2EVMTollOffRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*Any2EVMTollOffRampPoolAddedIterator, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampPoolAddedIterator{contract: _Any2EVMTollOffRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampPoolAdded)
				if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) ParsePoolAdded(log types.Log) (*Any2EVMTollOffRampPoolAdded, error) {
	event := new(Any2EVMTollOffRampPoolAdded)
	if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampPoolRemovedIterator struct {
	Event *Any2EVMTollOffRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampPoolRemoved)
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
		it.Event = new(Any2EVMTollOffRampPoolRemoved)
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

func (it *Any2EVMTollOffRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*Any2EVMTollOffRampPoolRemovedIterator, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampPoolRemovedIterator{contract: _Any2EVMTollOffRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampPoolRemoved)
				if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) ParsePoolRemoved(log types.Log) (*Any2EVMTollOffRampPoolRemoved, error) {
	event := new(Any2EVMTollOffRampPoolRemoved)
	if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampTransmittedIterator struct {
	Event *Any2EVMTollOffRampTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampTransmitted)
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
		it.Event = new(Any2EVMTollOffRampTransmitted)
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

func (it *Any2EVMTollOffRampTransmittedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterTransmitted(opts *bind.FilterOpts) (*Any2EVMTollOffRampTransmittedIterator, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampTransmittedIterator{contract: _Any2EVMTollOffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampTransmitted) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampTransmitted)
				if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) ParseTransmitted(log types.Log) (*Any2EVMTollOffRampTransmitted, error) {
	event := new(Any2EVMTollOffRampTransmitted)
	if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampUnpausedIterator struct {
	Event *Any2EVMTollOffRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampUnpaused)
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
		it.Event = new(Any2EVMTollOffRampUnpaused)
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

func (it *Any2EVMTollOffRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*Any2EVMTollOffRampUnpausedIterator, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampUnpausedIterator{contract: _Any2EVMTollOffRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampUnpaused)
				if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) ParseUnpaused(log types.Log) (*Any2EVMTollOffRampUnpaused, error) {
	event := new(Any2EVMTollOffRampUnpaused)
	if err := _Any2EVMTollOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Any2EVMTollOffRamp.abi.Events["AFNMaxHeartbeatTimeSet"].ID:
		return _Any2EVMTollOffRamp.ParseAFNMaxHeartbeatTimeSet(log)
	case _Any2EVMTollOffRamp.abi.Events["AFNSet"].ID:
		return _Any2EVMTollOffRamp.ParseAFNSet(log)
	case _Any2EVMTollOffRamp.abi.Events["ConfigSet"].ID:
		return _Any2EVMTollOffRamp.ParseConfigSet(log)
	case _Any2EVMTollOffRamp.abi.Events["ExecutionCompleted"].ID:
		return _Any2EVMTollOffRamp.ParseExecutionCompleted(log)
	case _Any2EVMTollOffRamp.abi.Events["OffRampConfigSet"].ID:
		return _Any2EVMTollOffRamp.ParseOffRampConfigSet(log)
	case _Any2EVMTollOffRamp.abi.Events["OffRampRouterSet"].ID:
		return _Any2EVMTollOffRamp.ParseOffRampRouterSet(log)
	case _Any2EVMTollOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _Any2EVMTollOffRamp.ParseOwnershipTransferRequested(log)
	case _Any2EVMTollOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _Any2EVMTollOffRamp.ParseOwnershipTransferred(log)
	case _Any2EVMTollOffRamp.abi.Events["Paused"].ID:
		return _Any2EVMTollOffRamp.ParsePaused(log)
	case _Any2EVMTollOffRamp.abi.Events["PoolAdded"].ID:
		return _Any2EVMTollOffRamp.ParsePoolAdded(log)
	case _Any2EVMTollOffRamp.abi.Events["PoolRemoved"].ID:
		return _Any2EVMTollOffRamp.ParsePoolRemoved(log)
	case _Any2EVMTollOffRamp.abi.Events["Transmitted"].ID:
		return _Any2EVMTollOffRamp.ParseTransmitted(log)
	case _Any2EVMTollOffRamp.abi.Events["Unpaused"].ID:
		return _Any2EVMTollOffRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (Any2EVMTollOffRampAFNMaxHeartbeatTimeSet) Topic() common.Hash {
	return common.HexToHash("0x72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251")
}

func (Any2EVMTollOffRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (Any2EVMTollOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (Any2EVMTollOffRampExecutionCompleted) Topic() common.Hash {
	return common.HexToHash("0xbca6416e78a437ab47530846568a4d78457e41bc2adc0d91a826090e2d853d1c")
}

func (Any2EVMTollOffRampOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xedc1b12e6a2ea72b9768b6c0e185d04d9c656f3a270976aa68badc3c1d550902")
}

func (Any2EVMTollOffRampOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4")
}

func (Any2EVMTollOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (Any2EVMTollOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (Any2EVMTollOffRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (Any2EVMTollOffRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (Any2EVMTollOffRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (Any2EVMTollOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (Any2EVMTollOffRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRamp) Address() common.Address {
	return _Any2EVMTollOffRamp.address
}

type Any2EVMTollOffRampInterface interface {
	CHAINID(opts *bind.CallOpts) (*big.Int, error)

	SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error)

	CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMTollMessage) error

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetBlobVerifier(opts *bind.CallOpts) (common.Address, error)

	GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error)

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

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMTollMessage) (*types.Transaction, error)

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

	FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampAFNMaxHeartbeatTimeSetIterator, error)

	WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampAFNMaxHeartbeatTimeSet) (event.Subscription, error)

	ParseAFNMaxHeartbeatTimeSet(log types.Log) (*Any2EVMTollOffRampAFNMaxHeartbeatTimeSet, error)

	FilterAFNSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*Any2EVMTollOffRampAFNSet, error)

	FilterConfigSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*Any2EVMTollOffRampConfigSet, error)

	FilterExecutionCompleted(opts *bind.FilterOpts, sequenceNumber []uint64) (*Any2EVMTollOffRampExecutionCompletedIterator, error)

	WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampExecutionCompleted, sequenceNumber []uint64) (event.Subscription, error)

	ParseExecutionCompleted(log types.Log) (*Any2EVMTollOffRampExecutionCompleted, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*Any2EVMTollOffRampOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*Any2EVMTollOffRampOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*Any2EVMTollOffRampOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*Any2EVMTollOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*Any2EVMTollOffRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*Any2EVMTollOffRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*Any2EVMTollOffRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*Any2EVMTollOffRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*Any2EVMTollOffRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*Any2EVMTollOffRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*Any2EVMTollOffRampPoolRemoved, error)

	FilterTransmitted(opts *bind.FilterOpts) (*Any2EVMTollOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*Any2EVMTollOffRampTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*Any2EVMTollOffRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*Any2EVMTollOffRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
