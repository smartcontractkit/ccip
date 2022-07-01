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

var Any2EVMTollOffRampHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlobVerifier\",\"outputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMTollOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"}],\"name\":\"setBlobVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620058e0380380620058e08339810160408190526200003491620005b3565b8787878787878787600188888888888888888282858333806000806000806101000a81548160ff02191690831515021790555060006001600160a01b0316826001600160a01b031603620000cf5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620001095762000109816200034e565b5050506001600160a01b038216158062000121575080155b156200014057604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001865760405162d8548360e71b815260040160405180910390fd5b81516200019b906005906020850190620003ff565b5060005b82518110156200027d576000828281518110620001c057620001c0620006ee565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b0316815250600460008685815181106200020a576200020a620006ee565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff19166001179055620002758162000704565b90506200019f565b50508751608090815260a0999099525050855160085550505060208301516009805460408601516060870151979096015163ffffffff16600160c01b0263ffffffff60c01b196001600160401b03988916600160801b0216600160801b600160e01b031997891668010000000000000000026001600160801b0319909316989094169790971717949094161793909317909155600780546001600160a01b039092166001600160a01b031990921691909117905550151560c052506200072c9e505050505050505050505050505050565b336001600160a01b03821603620003a85760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000c6565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b82805482825590600052602060002090810192821562000457579160200282015b828111156200045757825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000420565b506200046592915062000469565b5090565b5b808211156200046557600081556001016200046a565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715620004bb57620004bb62000480565b60405290565b80516001600160401b0381168114620004d957600080fd5b919050565b6001600160a01b0381168114620004f457600080fd5b50565b8051620004d981620004de565b600082601f8301126200051657600080fd5b815160206001600160401b038083111562000535576200053562000480565b8260051b604051601f19603f830116810181811084821117156200055d576200055d62000480565b6040529384528581018301938381019250878511156200057c57600080fd5b83870191505b84821015620005a85781516200059881620004de565b8352918301919083019062000582565b979650505050505050565b600080600080600080600080888a03610180811215620005d257600080fd5b8951985060a0601f1982011215620005e957600080fd5b50620005f462000496565b60208a015181526200060960408b01620004c1565b60208201526200061c60608b01620004c1565b60408201526200062f60808b01620004c1565b606082015260a08a015163ffffffff811681146200064c57600080fd5b608082015296506200066160c08a01620004f7565b95506200067160e08a01620004f7565b9450620006826101008a01620004f7565b6101208a01519094506001600160401b0380821115620006a157600080fd5b620006af8c838d0162000504565b94506101408b0151915080821115620006c757600080fd5b50620006d68b828c0162000504565b92505061016089015190509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b6000600182016200072557634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c0516151766200076a6000396000611137015260006104680152600081816103de01528181612cfe01526134d901526151766000f3fe608060405234801561001057600080fd5b50600436106102415760003560e01c806389c0656811610145578063b6608c3b116100bd578063c3f909d41161008c578063eb511dd411610071578063eb511dd4146106cf578063efaf6eab146106e2578063f2fde38b146106f557600080fd5b8063c3f909d4146105b4578063e3d0e712146106bc57600080fd5b8063b6608c3b14610542578063bbe4f6db14610555578063be9b03f11461058e578063c0d78655146105a157600080fd5b8063afcb95d711610114578063b0f479a1116100f9578063b0f479a1146104fe578063b1dc65a41461051c578063b57671661461052f57600080fd5b8063afcb95d7146104d6578063b034909c146104f657600080fd5b806389c065681461048a5780638bbad066146104925780638da5cb5b146104a0578063a639d1c0146104c357600080fd5b80635c975abb116101d857806379ba5097116101a757806381ff70481161018c57806381ff70481461042b5780638456cb591461045b57806385e1f4d01461046357600080fd5b806379ba50971461040e578063814118341461041657600080fd5b80635c975abb1461039d5780636133dc24146103a8578063744b92e2146103c657806374be2150146103d957600080fd5b80632222dd42116102145780632222dd42146102fa5780633f4ba83a14610339578063567c814b146103415780635b16ebb71461036457600080fd5b8063092cddc214610246578063108ee5fc1461025b578063142a98fc1461026e578063181f5a77146102b1575b600080fd5b610259610254366004613c6b565b610708565b005b610259610269366004613d87565b610764565b61029b61027c366004613dab565b67ffffffffffffffff166000908152600a602052604090205460ff1690565b6040516102a89190613df7565b60405180910390f35b6102ed6040518060400160405280601881526020017f416e793245564d546f6c6c4f666652616d7020312e302e30000000000000000081525081565b6040516102a89190613eae565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102a8565b610259610840565b61035461034f366004613ec1565b610852565b60405190151581526020016102a8565b610354610372366004613d87565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff16610354565b60075473ffffffffffffffffffffffffffffffffffffffff16610314565b6102596103d4366004613eda565b610999565b6104007f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016102a8565b610259610d98565b61041e610ebf565b6040516102a89190613f64565b600d54600b546040805163ffffffff808516825264010000000090940490931660208401528201526060016102a8565b610259610f2e565b6104007f000000000000000000000000000000000000000000000000000000000000000081565b61041e610f3e565b610259610241366004613f77565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610314565b6102596104d1366004613d87565b610fab565b6040805160018152600060208201819052918101919091526060016102a8565b600354610400565b60115473ffffffffffffffffffffffffffffffffffffffff16610314565b61025961052a366004613fff565b610ffa565b61025961053d3660046140e4565b6116a1565b610259610550366004613ec1565b6116ad565b610314610563366004613d87565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b61025961059c366004614282565b61172d565b6102596105af366004613d87565b612050565b6106626040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a081018252600854815260095467ffffffffffffffff808216602084015268010000000000000000820481169383019390935270010000000000000000000000000000000081049092166060820152780100000000000000000000000000000000000000000000000090910463ffffffff16608082015290565b6040516102a89190600060a08201905082518252602083015167ffffffffffffffff8082166020850152806040860151166040850152806060860151166060850152505063ffffffff608084015116608083015292915050565b6102596106ca3660046143dc565b6120c7565b6102596106dd366004613eda565b612aaa565b6102596106f03660046144a9565b612cf2565b610259610703366004613d87565b612ec7565b333014610741576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107588160a001518260c001518360600151612ed8565b61076181612f72565b50565b61076c613075565b73ffffffffffffffffffffffffffffffffffffffff81166107b9576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b610848613075565b6108506130fb565b565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa1580156108c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e69190614541565b1580156109935750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa15801561095e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610982919061455e565b6020015161099090846145e9565b11155b92915050565b6109a1613075565b60055460008190036109df576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610a7a576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610ae3576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006005610af26001856145e9565b81548110610b0257610b02614600565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff1681548110610b5457610b54614600565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166005610b836001866145e9565b81548110610b9357610b93614600565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110610c0157610c01614600565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480610ca357610ca361462f565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610e1e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60606010805480602002602001604051908101604052809291908181526020018280548015610f2457602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610ef9575b5050505050905090565b610f36613075565b6108506131dc565b60606005805480602002602001604051908101604052809291908181526020018280548015610f245760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610ef9575050505050905090565b610fb3613075565b600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161105091849163ffffffff851691908e908e908190840183828082843760009201919091525061329c92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600b54808252600c5460ff80821660208501526101009091041692820192909252908314611125576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610e15565b6111338b8b8b8b8b8b613322565b60007f00000000000000000000000000000000000000000000000000000000000000001561119057600282602001518360400151611171919061465e565b61117b91906146b2565b61118690600161465e565b60ff1690506111a6565b60208201516111a090600161465e565b60ff1690505b88811461120f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610e15565b888714611278576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610e15565b336000908152600e602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156112bb576112bb613dc8565b60028111156112cc576112cc613dc8565b90525090506002816020015160028111156112e9576112e9613dc8565b14801561133057506010816000015160ff168154811061130b5761130b614600565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b611396576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610e15565b5050505050600088886040516113ad9291906146d4565b6040519081900381206113c4918c906020016146e4565b6040516020818303038152906040528051906020012090506113e46139a8565b604080518082019091526000808252602082015260005b8881101561167f57600060018588846020811061141a5761141a614600565b61142791901a601b61465e565b8d8d8681811061143957611439614600565b905060200201358c8c8781811061145257611452614600565b905060200201356040516000815260200160405260405161148f949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156114b1573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600e602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561153157611531613dc8565b600281111561154257611542613dc8565b905250925060018360200151600281111561155f5761155f613dc8565b146115c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610e15565b8251849060ff16601f81106115dd576115dd614600565b602002015115611649576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610e15565b600184846000015160ff16601f811061166457611664614600565b911515602090920201525061167881614700565b90506113fb565b5050505063ffffffff811061169657611696614738565b505050505050505050565b6107616000808361329c565b6116b5613075565b806000036116ef576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c2519101610834565b60005460ff161561179a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610e15565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611807573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061182b9190614541565b15611861576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa1580156118d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f5919061455e565b905060035481602001514261190a91906145e9565b1115611942576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60115473ffffffffffffffffffffffffffffffffffffffff16611991576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608301515160008190036119d2576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff8111156119ed576119ed6139c7565b604051908082528060200260200182016040528015611a16578160200160208202803683370190505b50905060008267ffffffffffffffff811115611a3457611a346139c7565b604051908082528060200260200182016040528015611b0757816020015b611af460405180610140016040528060008152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b815260200190600190039081611a525790505b50905060005b83811015611bda5786606001518181518110611b2b57611b2b614600565b6020026020010151806020019051810190611b469190614881565b828281518110611b5857611b58614600565b6020026020010181905250600087606001518281518110611b7b57611b7b614600565b6020026020010151604051602001611b93919061499d565b60405160208183030381529060405290508080519060200120848381518110611bbe57611bbe614600565b602090810291909101015250611bd381614700565b9050611b0d565b50600080611bfb8489608001518a60a001518b60c001518c60e001516133d0565b915091506000835182611c0e91906149c3565b9050878015611c4957506009547801000000000000000000000000000000000000000000000000900463ffffffff16611c4784426145e9565b105b15611c80576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b86811015612044576000858281518110611c9f57611c9f614600565b602002602001015190506000611cd2826020015167ffffffffffffffff166000908152600a602052604090205460ff1690565b90506002816003811115611ce857611ce8613dc8565b03611d315760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610e15565b611d3a826134d5565b60005b8260a0015151811015611d8357611d708360a001518281518110611d6357611d63614600565b602002602001015161363a565b5080611d7b81614700565b915050611d3d565b506000816003811115611d9857611d98613dc8565b148015611da357508a155b15611f3a576000805b8d6020015151811015611e3e578360e0015173ffffffffffffffffffffffffffffffffffffffff168e602001518281518110611dea57611dea614600565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1603611e2e578d604001518181518110611e2357611e23614600565b602002602001015191505b611e3781614700565b9050611dac565b5080611e945760e08301516040517fce480bcc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610e15565b6000670de0b6b3a7640000823a86610120015189611eb291906149d7565b611ebc91906149ef565b611ec691906149ef565b611ed091906149c3565b9050836101000151811115611f235760208401516040517f6b830fc700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610e15565b611f378460e00151856101000151306136b6565b50505b60208281015167ffffffffffffffff166000908152600a9091526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611f8a83613753565b60208085015167ffffffffffffffff166000908152600a90915260409020805491925082917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001836003811115611fe557611fe5613dc8565b0217905550826020015167ffffffffffffffff167fbca6416e78a437ab47530846568a4d78457e41bc2adc0d91a826090e2d853d1c826040516120289190613df7565b60405180910390a25050508061203d90614700565b9050611c83565b50505050505050505050565b612058613075565b601180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b855185518560ff16601f83111561213a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610e15565b600081116121a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610e15565b818314612232576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610e15565b61223d8160036149ef565b83116122a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610e15565b6122ad613075565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600f54156124a057600f54600090612305906001906145e9565b90506000600f828154811061231c5761231c614600565b60009182526020822001546010805473ffffffffffffffffffffffffffffffffffffffff9092169350908490811061235657612356614600565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600e909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600f805491925090806123d6576123d661462f565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055601080548061243f5761243f61462f565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055506122eb915050565b60005b815151811015612905576000600e6000846000015184815181106124c9576124c9614600565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561251357612513613dc8565b1461257a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610e15565b6040805180820190915260ff821681526001602082015282518051600e91600091859081106125ab576125ab614600565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561264c5761264c613dc8565b02179055506000915061265c9050565b600e60008460200151848151811061267657612676614600565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff1660028111156126c0576126c0613dc8565b14612727576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610e15565b6040805180820190915260ff82168152602081016002815250600e60008460200151848151811061275a5761275a614600565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156127fb576127fb613dc8565b02179055505082518051600f92508390811061281957612819614600565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051601091908390811061289557612895614600565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790556128fe81614700565b90506124a3565b506040810151600c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600d80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092612997928692908216911617614a2c565b92506101000a81548163ffffffff021916908363ffffffff1602179055506129f64630600d60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613802565b600b81905582518051600c805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600d5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612a95988b98919763ffffffff909216969095919491939192614a54565b60405180910390a15050505050505050505050565b612ab2613075565b73ffffffffffffffffffffffffffffffffffffffff82161580612ae9575073ffffffffffffffffffffffffffffffffffffffff8116155b15612b20576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612bbc576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b612cfa613075565b80517f000000000000000000000000000000000000000000000000000000000000000014612d5a5780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610e15565b80516008556020810151600980546040808501516060860151608087015163ffffffff167801000000000000000000000000000000000000000000000000027fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff92831670010000000000000000000000000000000002167fffffffff000000000000000000000000ffffffffffffffffffffffffffffffff93831668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009096169290971691909117939093171693909317179055517fedc1b12e6a2ea72b9768b6c0e185d04d9c656f3a270976aa68badc3c1d55090290612ebc908390600060a08201905082518252602083015167ffffffffffffffff8082166020850152806040860151166040850152806060860151166060850152505063ffffffff608084015116608083015292915050565b60405180910390a150565b612ecf613075565b610761816138ad565b8151835114612f13576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015612f6c57612f5c848281518110612f3457612f34614600565b6020026020010151848381518110612f4e57612f4e614600565b6020026020010151846136b6565b612f6581614700565b9050612f16565b50505050565b606081015173ffffffffffffffffffffffffffffffffffffffff163b612fe25760608101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610e15565b60608101516011546040517ffd12f6e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063fd12f6e59061303f9084908690600401614c23565b600060405180830381600087803b15801561305957600080fd5b505af115801561306d573d6000803e3d6000fd5b505050505050565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610850576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610e15565b60005460ff16613167576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610e15565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60005460ff1615613249576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610e15565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586131b23390565b3073ffffffffffffffffffffffffffffffffffffffff1663be9b03f1828060200190518101906132cc9190614d9a565b60006040518363ffffffff1660e01b81526004016132eb929190614f5f565b600060405180830381600087803b15801561330557600080fd5b505af1158015613319573d6000803e3d6000fd5b50505050505050565b600061332f8260206149ef565b61333a8560206149ef565b613346886101446149d7565b61335091906149d7565b61335a91906149d7565b6133659060006149d7565b9050368114613319576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610e15565b60008060005a6007546040517fe71e65ce00000000000000000000000000000000000000000000000000000000815291925060009173ffffffffffffffffffffffffffffffffffffffff9091169063e71e65ce9061343a908c908c908c908c908c90600401615056565b6020604051808303816000875af1158015613459573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061347d91906150a8565b9050600081116134b9576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a6134c590846145e9565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146135355780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610e15565b60095460a08201515170010000000000000000000000000000000090910467ffffffffffffffff16108061357357508060c00151518160a001515114155b156135bc5760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610e15565b6009546080820151516801000000000000000090910467ffffffffffffffff161015610761576009546080820151516040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090920467ffffffffffffffff1660048301526024820152604401610e15565b73ffffffffffffffffffffffffffffffffffffffff81811660009081526004602052604090205416806136b1576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610e15565b919050565b60006136c18461363a565b6040517fea6192a200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690529192509082169063ea6192a290604401600060405180830381600087803b15801561373557600080fd5b505af1158015613749573d6000803e3d6000fd5b5050505050505050565b6040517f092cddc2000000000000000000000000000000000000000000000000000000008152600090309063092cddc2906137929085906004016150c1565b600060405180830381600087803b1580156137ac57600080fd5b505af19250505080156137bd575060015b6137fa573d8080156137eb576040519150601f19603f3d011682016040523d82523d6000602084013e6137f0565b606091505b5060039392505050565b506002919050565b6000808a8a8a8a8a8a8a8a8a604051602001613826999897969594939291906150d4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff82160361392c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610e15565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604051806103e00160405280601f906020820280368337509192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610140810167ffffffffffffffff81118282101715613a1a57613a1a6139c7565b60405290565b604051610100810167ffffffffffffffff81118282101715613a1a57613a1a6139c7565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613a8b57613a8b6139c7565b604052919050565b67ffffffffffffffff8116811461076157600080fd5b80356136b181613a93565b73ffffffffffffffffffffffffffffffffffffffff8116811461076157600080fd5b80356136b181613ab4565b600067ffffffffffffffff821115613afb57613afb6139c7565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112613b3857600080fd5b8135613b4b613b4682613ae1565b613a44565b818152846020838601011115613b6057600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115613b9757613b976139c7565b5060051b60200190565b600082601f830112613bb257600080fd5b81356020613bc2613b4683613b7d565b82815260059290921b84018101918181019086841115613be157600080fd5b8286015b84811015613c05578035613bf881613ab4565b8352918301918301613be5565b509695505050505050565b600082601f830112613c2157600080fd5b81356020613c31613b4683613b7d565b82815260059290921b84018101918181019086841115613c5057600080fd5b8286015b84811015613c055780358352918301918301613c54565b600060208284031215613c7d57600080fd5b813567ffffffffffffffff80821115613c9557600080fd5b908301906101408286031215613caa57600080fd5b613cb26139f6565b82358152613cc260208401613aa9565b6020820152613cd360408401613ad6565b6040820152613ce460608401613ad6565b6060820152608083013582811115613cfb57600080fd5b613d0787828601613b27565b60808301525060a083013582811115613d1f57600080fd5b613d2b87828601613ba1565b60a08301525060c083013582811115613d4357600080fd5b613d4f87828601613c10565b60c083015250613d6160e08401613ad6565b60e082015261010083810135908201526101209283013592810192909252509392505050565b600060208284031215613d9957600080fd5b8135613da481613ab4565b9392505050565b600060208284031215613dbd57600080fd5b8135613da481613a93565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310613e32577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b83811015613e53578181015183820152602001613e3b565b83811115612f6c5750506000910152565b60008151808452613e7c816020860160208601613e38565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613da46020830184613e64565b600060208284031215613ed357600080fd5b5035919050565b60008060408385031215613eed57600080fd5b8235613ef881613ab4565b91506020830135613f0881613ab4565b809150509250929050565b600081518084526020808501945080840160005b83811015613f5957815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101613f27565b509495945050505050565b602081526000613da46020830184613f13565b600060208284031215613f8957600080fd5b813567ffffffffffffffff811115613fa057600080fd5b82016101408185031215613da457600080fd5b60008083601f840112613fc557600080fd5b50813567ffffffffffffffff811115613fdd57600080fd5b6020830191508360208260051b8501011115613ff857600080fd5b9250929050565b60008060008060008060008060e0898b03121561401b57600080fd5b606089018a81111561402c57600080fd5b8998503567ffffffffffffffff8082111561404657600080fd5b818b0191508b601f83011261405a57600080fd5b81358181111561406957600080fd5b8c602082850101111561407b57600080fd5b6020830199508098505060808b013591508082111561409957600080fd5b6140a58c838d01613fb3565b909750955060a08b01359150808211156140be57600080fd5b506140cb8b828c01613fb3565b999c989b50969995989497949560c00135949350505050565b6000602082840312156140f657600080fd5b813567ffffffffffffffff81111561410d57600080fd5b61411984828501613b27565b949350505050565b600082601f83011261413257600080fd5b81356020614142613b4683613b7d565b82815260059290921b8401810191818101908684111561416157600080fd5b8286015b84811015613c0557803561417881613a93565b8352918301918301614165565b600082601f83011261419657600080fd5b813560206141a6613b4683613b7d565b82815260059290921b840181019181810190868411156141c557600080fd5b8286015b84811015613c055780356141dc81613ab4565b83529183019183016141c9565b600082601f8301126141fa57600080fd5b8135602061420a613b4683613b7d565b82815260059290921b8401810191818101908684111561422957600080fd5b8286015b84811015613c0557803567ffffffffffffffff81111561424d5760008081fd5b61425b8986838b0101613b27565b84525091830191830161422d565b801515811461076157600080fd5b80356136b181614269565b6000806040838503121561429557600080fd5b823567ffffffffffffffff808211156142ad57600080fd5b9084019061010082870312156142c257600080fd5b6142ca613a20565b8235828111156142d957600080fd5b6142e588828601614121565b8252506020830135828111156142fa57600080fd5b61430688828601614185565b60208301525060408301358281111561431e57600080fd5b61432a88828601613c10565b60408301525060608301358281111561434257600080fd5b61434e888286016141e9565b60608301525060808301358281111561436657600080fd5b61437288828601613c10565b60808301525060a083013560a082015260c08301358281111561439457600080fd5b6143a088828601613c10565b60c08301525060e083013560e08201528094505050506143c260208401614277565b90509250929050565b803560ff811681146136b157600080fd5b60008060008060008060c087890312156143f557600080fd5b863567ffffffffffffffff8082111561440d57600080fd5b6144198a838b01614185565b9750602089013591508082111561442f57600080fd5b61443b8a838b01614185565b965061444960408a016143cb565b9550606089013591508082111561445f57600080fd5b61446b8a838b01613b27565b945061447960808a01613aa9565b935060a089013591508082111561448f57600080fd5b5061449c89828a01613b27565b9150509295509295509295565b600060a082840312156144bb57600080fd5b60405160a0810181811067ffffffffffffffff821117156144de576144de6139c7565b6040528235815260208301356144f381613a93565b6020820152604083013561450681613a93565b6040820152606083013561451981613a93565b6060820152608083013563ffffffff8116811461453557600080fd5b60808201529392505050565b60006020828403121561455357600080fd5b8151613da481614269565b60006060828403121561457057600080fd5b6040516060810181811067ffffffffffffffff82111715614593576145936139c7565b80604052508251815260208301516020820152604083015160408201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156145fb576145fb6145ba565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff0382111561467b5761467b6145ba565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff8316806146c5576146c5614683565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614731576147316145ba565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b80516136b181613a93565b80516136b181613ab4565b600082601f83011261478e57600080fd5b815161479c613b4682613ae1565b8181528460208386010111156147b157600080fd5b614119826020830160208701613e38565b600082601f8301126147d357600080fd5b815160206147e3613b4683613b7d565b82815260059290921b8401810191818101908684111561480257600080fd5b8286015b84811015613c0557805161481981613ab4565b8352918301918301614806565b600082601f83011261483757600080fd5b81516020614847613b4683613b7d565b82815260059290921b8401810191818101908684111561486657600080fd5b8286015b84811015613c05578051835291830191830161486a565b60006020828403121561489357600080fd5b815167ffffffffffffffff808211156148ab57600080fd5b9083019061014082860312156148c057600080fd5b6148c86139f6565b825181526148d860208401614767565b60208201526148e960408401614772565b60408201526148fa60608401614772565b606082015260808301518281111561491157600080fd5b61491d8782860161477d565b60808301525060a08301518281111561493557600080fd5b614941878286016147c2565b60a08301525060c08301518281111561495957600080fd5b61496587828601614826565b60c08301525061497760e08401614772565b60e082015261010083810151908201526101209283015192810192909252509392505050565b60008152600082516149b6816001850160208701613e38565b9190910160010192915050565b6000826149d2576149d2614683565b500490565b600082198211156149ea576149ea6145ba565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614a2757614a276145ba565b500290565b600063ffffffff808316818516808303821115614a4b57614a4b6145ba565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614a848184018a613f13565b90508281036080840152614a988189613f13565b905060ff871660a084015282810360c0840152614ab58187613e64565b905067ffffffffffffffff851660e0840152828103610100840152614ada8185613e64565b9c9b505050505050505050505050565b600081518084526020808501945080840160005b83811015613f5957815187529582019590820190600101614afe565b6000610140825184526020830151614b3e602086018267ffffffffffffffff169052565b506040830151614b66604086018273ffffffffffffffffffffffffffffffffffffffff169052565b506060830151614b8e606086018273ffffffffffffffffffffffffffffffffffffffff169052565b506080830151816080860152614ba682860182613e64565b91505060a083015184820360a0860152614bc08282613f13565b91505060c083015184820360c0860152614bda8282614aea565b91505060e0830151614c0460e086018273ffffffffffffffffffffffffffffffffffffffff169052565b5061010083810151908501526101209283015192909301919091525090565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006141196040830184614b1a565b600082601f830112614c6357600080fd5b81516020614c73613b4683613b7d565b82815260059290921b84018101918181019086841115614c9257600080fd5b8286015b84811015613c05578051614ca981613a93565b8352918301918301614c96565b600082601f830112614cc757600080fd5b81516020614cd7613b4683613b7d565b82815260059290921b84018101918181019086841115614cf657600080fd5b8286015b84811015613c05578051614d0d81613ab4565b8352918301918301614cfa565b600082601f830112614d2b57600080fd5b81516020614d3b613b4683613b7d565b82815260059290921b84018101918181019086841115614d5a57600080fd5b8286015b84811015613c0557805167ffffffffffffffff811115614d7e5760008081fd5b614d8c8986838b010161477d565b845250918301918301614d5e565b600060208284031215614dac57600080fd5b815167ffffffffffffffff80821115614dc457600080fd5b908301906101008286031215614dd957600080fd5b614de1613a20565b825182811115614df057600080fd5b614dfc87828601614c52565b825250602083015182811115614e1157600080fd5b614e1d87828601614cb6565b602083015250604083015182811115614e3557600080fd5b614e4187828601614826565b604083015250606083015182811115614e5957600080fd5b614e6587828601614d1a565b606083015250608083015182811115614e7d57600080fd5b614e8987828601614826565b60808301525060a083015160a082015260c083015182811115614eab57600080fd5b614eb787828601614826565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b83811015613f5957815167ffffffffffffffff1687529582019590820190600101614ee7565b6000815180845260208085019450848260051b860182860160005b85811015614f52578383038952614f40838351613e64565b98850198925090840190600101614f28565b5090979650505050505050565b6040815260008351610100806040850152614f7e610140850183614ed3565b915060208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080868503016060870152614fba8483613f13565b93506040880151915080868503016080870152614fd78483614aea565b935060608801519150808685030160a0870152614ff48483614f0d565b935060808801519150808685030160c08701526150118483614aea565b935060a088015160e087015260c08801519150808685030183870152506150388382614aea565b60e088015161012087015286151560208701529350613da492505050565b60a08152600061506960a0830188614aea565b828103602084015261507b8188614aea565b905085604084015282810360608401526150958186614aea565b9150508260808301529695505050505050565b6000602082840312156150ba57600080fd5b5051919050565b602081526000613da46020830184614b1a565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b16604085015281606085015261511b8285018b613f13565b9150838203608085015261512f828a613f13565b915060ff881660a085015283820360c085015261514c8288613e64565b90861660e08501528381036101008501529050614ada8185613e6456fea164736f6c634300080f000a",
}

var Any2EVMTollOffRampHelperABI = Any2EVMTollOffRampHelperMetaData.ABI

var Any2EVMTollOffRampHelperBin = Any2EVMTollOffRampHelperMetaData.Bin

func DeployAny2EVMTollOffRampHelper(auth *bind.TransactOpts, backend bind.ContractBackend, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, blobVerifier common.Address, onRampAddress common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, maxTimeWithoutAFNSignal *big.Int) (common.Address, *types.Transaction, *Any2EVMTollOffRampHelper, error) {
	parsed, err := Any2EVMTollOffRampHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Any2EVMTollOffRampHelperBin), backend, chainId, offRampConfig, blobVerifier, onRampAddress, afn, sourceTokens, pools, maxTimeWithoutAFNSignal)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Any2EVMTollOffRampHelper{Any2EVMTollOffRampHelperCaller: Any2EVMTollOffRampHelperCaller{contract: contract}, Any2EVMTollOffRampHelperTransactor: Any2EVMTollOffRampHelperTransactor{contract: contract}, Any2EVMTollOffRampHelperFilterer: Any2EVMTollOffRampHelperFilterer{contract: contract}}, nil
}

type Any2EVMTollOffRampHelper struct {
	address common.Address
	abi     abi.ABI
	Any2EVMTollOffRampHelperCaller
	Any2EVMTollOffRampHelperTransactor
	Any2EVMTollOffRampHelperFilterer
}

type Any2EVMTollOffRampHelperCaller struct {
	contract *bind.BoundContract
}

type Any2EVMTollOffRampHelperTransactor struct {
	contract *bind.BoundContract
}

type Any2EVMTollOffRampHelperFilterer struct {
	contract *bind.BoundContract
}

type Any2EVMTollOffRampHelperSession struct {
	Contract     *Any2EVMTollOffRampHelper
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type Any2EVMTollOffRampHelperCallerSession struct {
	Contract *Any2EVMTollOffRampHelperCaller
	CallOpts bind.CallOpts
}

type Any2EVMTollOffRampHelperTransactorSession struct {
	Contract     *Any2EVMTollOffRampHelperTransactor
	TransactOpts bind.TransactOpts
}

type Any2EVMTollOffRampHelperRaw struct {
	Contract *Any2EVMTollOffRampHelper
}

type Any2EVMTollOffRampHelperCallerRaw struct {
	Contract *Any2EVMTollOffRampHelperCaller
}

type Any2EVMTollOffRampHelperTransactorRaw struct {
	Contract *Any2EVMTollOffRampHelperTransactor
}

func NewAny2EVMTollOffRampHelper(address common.Address, backend bind.ContractBackend) (*Any2EVMTollOffRampHelper, error) {
	abi, err := abi.JSON(strings.NewReader(Any2EVMTollOffRampHelperABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindAny2EVMTollOffRampHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelper{address: address, abi: abi, Any2EVMTollOffRampHelperCaller: Any2EVMTollOffRampHelperCaller{contract: contract}, Any2EVMTollOffRampHelperTransactor: Any2EVMTollOffRampHelperTransactor{contract: contract}, Any2EVMTollOffRampHelperFilterer: Any2EVMTollOffRampHelperFilterer{contract: contract}}, nil
}

func NewAny2EVMTollOffRampHelperCaller(address common.Address, caller bind.ContractCaller) (*Any2EVMTollOffRampHelperCaller, error) {
	contract, err := bindAny2EVMTollOffRampHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperCaller{contract: contract}, nil
}

func NewAny2EVMTollOffRampHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*Any2EVMTollOffRampHelperTransactor, error) {
	contract, err := bindAny2EVMTollOffRampHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperTransactor{contract: contract}, nil
}

func NewAny2EVMTollOffRampHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*Any2EVMTollOffRampHelperFilterer, error) {
	contract, err := bindAny2EVMTollOffRampHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperFilterer{contract: contract}, nil
}

func bindAny2EVMTollOffRampHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Any2EVMTollOffRampHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Any2EVMTollOffRampHelper.Contract.Any2EVMTollOffRampHelperCaller.contract.Call(opts, result, method, params...)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.Any2EVMTollOffRampHelperTransactor.contract.Transfer(opts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.Any2EVMTollOffRampHelperTransactor.contract.Transact(opts, method, params...)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Any2EVMTollOffRampHelper.Contract.contract.Call(opts, result, method, params...)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.contract.Transfer(opts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.contract.Transact(opts, method, params...)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) CHAINID() (*big.Int, error) {
	return _Any2EVMTollOffRampHelper.Contract.CHAINID(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) CHAINID() (*big.Int, error) {
	return _Any2EVMTollOffRampHelper.Contract.CHAINID(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) SOURCECHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "SOURCE_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) SOURCECHAINID() (*big.Int, error) {
	return _Any2EVMTollOffRampHelper.Contract.SOURCECHAINID(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) SOURCECHAINID() (*big.Int, error) {
	return _Any2EVMTollOffRampHelper.Contract.SOURCECHAINID(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMTollMessage) error {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) CcipReceive(arg0 CCIPAny2EVMTollMessage) error {
	return _Any2EVMTollOffRampHelper.Contract.CcipReceive(&_Any2EVMTollOffRampHelper.CallOpts, arg0)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) CcipReceive(arg0 CCIPAny2EVMTollMessage) error {
	return _Any2EVMTollOffRampHelper.Contract.CcipReceive(&_Any2EVMTollOffRampHelper.CallOpts, arg0)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) GetAFN() (common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetAFN(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) GetAFN() (common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetAFN(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) GetBlobVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "getBlobVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) GetBlobVerifier() (common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetBlobVerifier(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) GetBlobVerifier() (common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetBlobVerifier(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOffRampInterfaceOffRampConfig)).(*BaseOffRampInterfaceOffRampConfig)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetConfig(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetConfig(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetExecutionState(&_Any2EVMTollOffRampHelper.CallOpts, sequenceNumber)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetExecutionState(&_Any2EVMTollOffRampHelper.CallOpts, sequenceNumber)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "getMaxSecondsWithoutAFNHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetPool(&_Any2EVMTollOffRampHelper.CallOpts, sourceToken)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetPool(&_Any2EVMTollOffRampHelper.CallOpts, sourceToken)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) GetPoolTokens() ([]common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetPoolTokens(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetPoolTokens(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) GetRouter() (common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetRouter(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) GetRouter() (common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.GetRouter(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "isHealthy", timeNow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _Any2EVMTollOffRampHelper.Contract.IsHealthy(&_Any2EVMTollOffRampHelper.CallOpts, timeNow)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _Any2EVMTollOffRampHelper.Contract.IsHealthy(&_Any2EVMTollOffRampHelper.CallOpts, timeNow)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) IsPool(addr common.Address) (bool, error) {
	return _Any2EVMTollOffRampHelper.Contract.IsPool(&_Any2EVMTollOffRampHelper.CallOpts, addr)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) IsPool(addr common.Address) (bool, error) {
	return _Any2EVMTollOffRampHelper.Contract.IsPool(&_Any2EVMTollOffRampHelper.CallOpts, addr)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _Any2EVMTollOffRampHelper.Contract.LatestConfigDetails(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _Any2EVMTollOffRampHelper.Contract.LatestConfigDetails(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _Any2EVMTollOffRampHelper.Contract.LatestConfigDigestAndEpoch(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _Any2EVMTollOffRampHelper.Contract.LatestConfigDigestAndEpoch(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) Owner() (common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.Owner(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) Owner() (common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.Owner(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) Paused() (bool, error) {
	return _Any2EVMTollOffRampHelper.Contract.Paused(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) Paused() (bool, error) {
	return _Any2EVMTollOffRampHelper.Contract.Paused(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) Transmitters() ([]common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.Transmitters(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) Transmitters() ([]common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.Transmitters(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) TypeAndVersion() (string, error) {
	return _Any2EVMTollOffRampHelper.Contract.TypeAndVersion(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) TypeAndVersion() (string, error) {
	return _Any2EVMTollOffRampHelper.Contract.TypeAndVersion(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "acceptOwnership")
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) AcceptOwnership() (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.AcceptOwnership(&_Any2EVMTollOffRampHelper.TransactOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.AcceptOwnership(&_Any2EVMTollOffRampHelper.TransactOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "addPool", token, pool)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.AddPool(&_Any2EVMTollOffRampHelper.TransactOpts, token, pool)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.AddPool(&_Any2EVMTollOffRampHelper.TransactOpts, token, pool)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) Execute(opts *bind.TransactOpts, report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "execute", report, manualExecution)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.Execute(&_Any2EVMTollOffRampHelper.TransactOpts, report, manualExecution)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.Execute(&_Any2EVMTollOffRampHelper.TransactOpts, report, manualExecution)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMTollMessage) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "executeSingleMessage", message)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) ExecuteSingleMessage(message CCIPAny2EVMTollMessage) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.ExecuteSingleMessage(&_Any2EVMTollOffRampHelper.TransactOpts, message)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMTollMessage) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.ExecuteSingleMessage(&_Any2EVMTollOffRampHelper.TransactOpts, message)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "pause")
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) Pause() (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.Pause(&_Any2EVMTollOffRampHelper.TransactOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) Pause() (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.Pause(&_Any2EVMTollOffRampHelper.TransactOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "removePool", token, pool)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.RemovePool(&_Any2EVMTollOffRampHelper.TransactOpts, token, pool)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.RemovePool(&_Any2EVMTollOffRampHelper.TransactOpts, token, pool)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "report", executableMessages)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) Report(executableMessages []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.Report(&_Any2EVMTollOffRampHelper.TransactOpts, executableMessages)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) Report(executableMessages []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.Report(&_Any2EVMTollOffRampHelper.TransactOpts, executableMessages)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "setAFN", afn)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.SetAFN(&_Any2EVMTollOffRampHelper.TransactOpts, afn)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.SetAFN(&_Any2EVMTollOffRampHelper.TransactOpts, afn)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) SetBlobVerifier(opts *bind.TransactOpts, blobVerifier common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "setBlobVerifier", blobVerifier)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.SetBlobVerifier(&_Any2EVMTollOffRampHelper.TransactOpts, blobVerifier)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.SetBlobVerifier(&_Any2EVMTollOffRampHelper.TransactOpts, blobVerifier)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.SetConfig(&_Any2EVMTollOffRampHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.SetConfig(&_Any2EVMTollOffRampHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) SetConfig0(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "setConfig0", config)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) SetConfig0(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.SetConfig0(&_Any2EVMTollOffRampHelper.TransactOpts, config)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) SetConfig0(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.SetConfig0(&_Any2EVMTollOffRampHelper.TransactOpts, config)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_Any2EVMTollOffRampHelper.TransactOpts, newTime)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_Any2EVMTollOffRampHelper.TransactOpts, newTime)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "setRouter", router)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.SetRouter(&_Any2EVMTollOffRampHelper.TransactOpts, router)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.SetRouter(&_Any2EVMTollOffRampHelper.TransactOpts, router)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "transferOwnership", to)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.TransferOwnership(&_Any2EVMTollOffRampHelper.TransactOpts, to)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.TransferOwnership(&_Any2EVMTollOffRampHelper.TransactOpts, to)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.Transmit(&_Any2EVMTollOffRampHelper.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.Transmit(&_Any2EVMTollOffRampHelper.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "unpause")
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) Unpause() (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.Unpause(&_Any2EVMTollOffRampHelper.TransactOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) Unpause() (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.Unpause(&_Any2EVMTollOffRampHelper.TransactOpts)
}

type Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSetIterator struct {
	Event *Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet)
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
		it.Event = new(Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet)
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

func (it *Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSetIterator, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSetIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "AFNMaxHeartbeatTimeSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet)
				if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) ParseAFNMaxHeartbeatTimeSet(log types.Log) (*Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet, error) {
	event := new(Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet)
	if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampHelperAFNSetIterator struct {
	Event *Any2EVMTollOffRampHelperAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampHelperAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampHelperAFNSet)
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
		it.Event = new(Any2EVMTollOffRampHelperAFNSet)
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

func (it *Any2EVMTollOffRampHelperAFNSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampHelperAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampHelperAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterAFNSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperAFNSetIterator, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperAFNSetIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperAFNSet) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampHelperAFNSet)
				if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) ParseAFNSet(log types.Log) (*Any2EVMTollOffRampHelperAFNSet, error) {
	event := new(Any2EVMTollOffRampHelperAFNSet)
	if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampHelperConfigSetIterator struct {
	Event *Any2EVMTollOffRampHelperConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampHelperConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampHelperConfigSet)
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
		it.Event = new(Any2EVMTollOffRampHelperConfigSet)
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

func (it *Any2EVMTollOffRampHelperConfigSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampHelperConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampHelperConfigSet struct {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterConfigSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperConfigSetIterator, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperConfigSetIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperConfigSet) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampHelperConfigSet)
				if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) ParseConfigSet(log types.Log) (*Any2EVMTollOffRampHelperConfigSet, error) {
	event := new(Any2EVMTollOffRampHelperConfigSet)
	if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampHelperExecutionCompletedIterator struct {
	Event *Any2EVMTollOffRampHelperExecutionCompleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampHelperExecutionCompletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampHelperExecutionCompleted)
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
		it.Event = new(Any2EVMTollOffRampHelperExecutionCompleted)
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

func (it *Any2EVMTollOffRampHelperExecutionCompletedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampHelperExecutionCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampHelperExecutionCompleted struct {
	SequenceNumber uint64
	State          uint8
	Raw            types.Log
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterExecutionCompleted(opts *bind.FilterOpts, sequenceNumber []uint64) (*Any2EVMTollOffRampHelperExecutionCompletedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "ExecutionCompleted", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperExecutionCompletedIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "ExecutionCompleted", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperExecutionCompleted, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "ExecutionCompleted", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampHelperExecutionCompleted)
				if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "ExecutionCompleted", log); err != nil {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) ParseExecutionCompleted(log types.Log) (*Any2EVMTollOffRampHelperExecutionCompleted, error) {
	event := new(Any2EVMTollOffRampHelperExecutionCompleted)
	if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "ExecutionCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampHelperOffRampConfigSetIterator struct {
	Event *Any2EVMTollOffRampHelperOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampHelperOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampHelperOffRampConfigSet)
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
		it.Event = new(Any2EVMTollOffRampHelperOffRampConfigSet)
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

func (it *Any2EVMTollOffRampHelperOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampHelperOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampHelperOffRampConfigSet struct {
	Config BaseOffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperOffRampConfigSetIterator, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperOffRampConfigSetIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampHelperOffRampConfigSet)
				if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) ParseOffRampConfigSet(log types.Log) (*Any2EVMTollOffRampHelperOffRampConfigSet, error) {
	event := new(Any2EVMTollOffRampHelperOffRampConfigSet)
	if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampHelperOffRampRouterSetIterator struct {
	Event *Any2EVMTollOffRampHelperOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampHelperOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampHelperOffRampRouterSet)
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
		it.Event = new(Any2EVMTollOffRampHelperOffRampRouterSet)
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

func (it *Any2EVMTollOffRampHelperOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampHelperOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampHelperOffRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*Any2EVMTollOffRampHelperOffRampRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperOffRampRouterSetIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperOffRampRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampHelperOffRampRouterSet)
				if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) ParseOffRampRouterSet(log types.Log) (*Any2EVMTollOffRampHelperOffRampRouterSet, error) {
	event := new(Any2EVMTollOffRampHelperOffRampRouterSet)
	if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampHelperOwnershipTransferRequestedIterator struct {
	Event *Any2EVMTollOffRampHelperOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampHelperOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampHelperOwnershipTransferRequested)
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
		it.Event = new(Any2EVMTollOffRampHelperOwnershipTransferRequested)
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

func (it *Any2EVMTollOffRampHelperOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampHelperOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampHelperOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampHelperOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperOwnershipTransferRequestedIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampHelperOwnershipTransferRequested)
				if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) ParseOwnershipTransferRequested(log types.Log) (*Any2EVMTollOffRampHelperOwnershipTransferRequested, error) {
	event := new(Any2EVMTollOffRampHelperOwnershipTransferRequested)
	if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampHelperOwnershipTransferredIterator struct {
	Event *Any2EVMTollOffRampHelperOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampHelperOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampHelperOwnershipTransferred)
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
		it.Event = new(Any2EVMTollOffRampHelperOwnershipTransferred)
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

func (it *Any2EVMTollOffRampHelperOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampHelperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampHelperOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampHelperOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperOwnershipTransferredIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampHelperOwnershipTransferred)
				if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) ParseOwnershipTransferred(log types.Log) (*Any2EVMTollOffRampHelperOwnershipTransferred, error) {
	event := new(Any2EVMTollOffRampHelperOwnershipTransferred)
	if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampHelperPausedIterator struct {
	Event *Any2EVMTollOffRampHelperPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampHelperPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampHelperPaused)
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
		it.Event = new(Any2EVMTollOffRampHelperPaused)
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

func (it *Any2EVMTollOffRampHelperPausedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampHelperPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampHelperPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterPaused(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperPausedIterator, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperPausedIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperPaused) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampHelperPaused)
				if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) ParsePaused(log types.Log) (*Any2EVMTollOffRampHelperPaused, error) {
	event := new(Any2EVMTollOffRampHelperPaused)
	if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampHelperPoolAddedIterator struct {
	Event *Any2EVMTollOffRampHelperPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampHelperPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampHelperPoolAdded)
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
		it.Event = new(Any2EVMTollOffRampHelperPoolAdded)
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

func (it *Any2EVMTollOffRampHelperPoolAddedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampHelperPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampHelperPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperPoolAddedIterator, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperPoolAddedIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperPoolAdded) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampHelperPoolAdded)
				if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) ParsePoolAdded(log types.Log) (*Any2EVMTollOffRampHelperPoolAdded, error) {
	event := new(Any2EVMTollOffRampHelperPoolAdded)
	if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampHelperPoolRemovedIterator struct {
	Event *Any2EVMTollOffRampHelperPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampHelperPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampHelperPoolRemoved)
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
		it.Event = new(Any2EVMTollOffRampHelperPoolRemoved)
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

func (it *Any2EVMTollOffRampHelperPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampHelperPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampHelperPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperPoolRemovedIterator, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperPoolRemovedIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampHelperPoolRemoved)
				if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) ParsePoolRemoved(log types.Log) (*Any2EVMTollOffRampHelperPoolRemoved, error) {
	event := new(Any2EVMTollOffRampHelperPoolRemoved)
	if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampHelperTransmittedIterator struct {
	Event *Any2EVMTollOffRampHelperTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampHelperTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampHelperTransmitted)
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
		it.Event = new(Any2EVMTollOffRampHelperTransmitted)
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

func (it *Any2EVMTollOffRampHelperTransmittedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampHelperTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampHelperTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterTransmitted(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperTransmittedIterator, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperTransmittedIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperTransmitted) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampHelperTransmitted)
				if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) ParseTransmitted(log types.Log) (*Any2EVMTollOffRampHelperTransmitted, error) {
	event := new(Any2EVMTollOffRampHelperTransmitted)
	if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampHelperUnpausedIterator struct {
	Event *Any2EVMTollOffRampHelperUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampHelperUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampHelperUnpaused)
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
		it.Event = new(Any2EVMTollOffRampHelperUnpaused)
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

func (it *Any2EVMTollOffRampHelperUnpausedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampHelperUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampHelperUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterUnpaused(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperUnpausedIterator, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperUnpausedIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperUnpaused) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampHelperUnpaused)
				if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) ParseUnpaused(log types.Log) (*Any2EVMTollOffRampHelperUnpaused, error) {
	event := new(Any2EVMTollOffRampHelperUnpaused)
	if err := _Any2EVMTollOffRampHelper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelper) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Any2EVMTollOffRampHelper.abi.Events["AFNMaxHeartbeatTimeSet"].ID:
		return _Any2EVMTollOffRampHelper.ParseAFNMaxHeartbeatTimeSet(log)
	case _Any2EVMTollOffRampHelper.abi.Events["AFNSet"].ID:
		return _Any2EVMTollOffRampHelper.ParseAFNSet(log)
	case _Any2EVMTollOffRampHelper.abi.Events["ConfigSet"].ID:
		return _Any2EVMTollOffRampHelper.ParseConfigSet(log)
	case _Any2EVMTollOffRampHelper.abi.Events["ExecutionCompleted"].ID:
		return _Any2EVMTollOffRampHelper.ParseExecutionCompleted(log)
	case _Any2EVMTollOffRampHelper.abi.Events["OffRampConfigSet"].ID:
		return _Any2EVMTollOffRampHelper.ParseOffRampConfigSet(log)
	case _Any2EVMTollOffRampHelper.abi.Events["OffRampRouterSet"].ID:
		return _Any2EVMTollOffRampHelper.ParseOffRampRouterSet(log)
	case _Any2EVMTollOffRampHelper.abi.Events["OwnershipTransferRequested"].ID:
		return _Any2EVMTollOffRampHelper.ParseOwnershipTransferRequested(log)
	case _Any2EVMTollOffRampHelper.abi.Events["OwnershipTransferred"].ID:
		return _Any2EVMTollOffRampHelper.ParseOwnershipTransferred(log)
	case _Any2EVMTollOffRampHelper.abi.Events["Paused"].ID:
		return _Any2EVMTollOffRampHelper.ParsePaused(log)
	case _Any2EVMTollOffRampHelper.abi.Events["PoolAdded"].ID:
		return _Any2EVMTollOffRampHelper.ParsePoolAdded(log)
	case _Any2EVMTollOffRampHelper.abi.Events["PoolRemoved"].ID:
		return _Any2EVMTollOffRampHelper.ParsePoolRemoved(log)
	case _Any2EVMTollOffRampHelper.abi.Events["Transmitted"].ID:
		return _Any2EVMTollOffRampHelper.ParseTransmitted(log)
	case _Any2EVMTollOffRampHelper.abi.Events["Unpaused"].ID:
		return _Any2EVMTollOffRampHelper.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet) Topic() common.Hash {
	return common.HexToHash("0x72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251")
}

func (Any2EVMTollOffRampHelperAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (Any2EVMTollOffRampHelperConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (Any2EVMTollOffRampHelperExecutionCompleted) Topic() common.Hash {
	return common.HexToHash("0xbca6416e78a437ab47530846568a4d78457e41bc2adc0d91a826090e2d853d1c")
}

func (Any2EVMTollOffRampHelperOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xedc1b12e6a2ea72b9768b6c0e185d04d9c656f3a270976aa68badc3c1d550902")
}

func (Any2EVMTollOffRampHelperOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4")
}

func (Any2EVMTollOffRampHelperOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (Any2EVMTollOffRampHelperOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (Any2EVMTollOffRampHelperPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (Any2EVMTollOffRampHelperPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (Any2EVMTollOffRampHelperPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (Any2EVMTollOffRampHelperTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (Any2EVMTollOffRampHelperUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelper) Address() common.Address {
	return _Any2EVMTollOffRampHelper.address
}

type Any2EVMTollOffRampHelperInterface interface {
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

	FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSetIterator, error)

	WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet) (event.Subscription, error)

	ParseAFNMaxHeartbeatTimeSet(log types.Log) (*Any2EVMTollOffRampHelperAFNMaxHeartbeatTimeSet, error)

	FilterAFNSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*Any2EVMTollOffRampHelperAFNSet, error)

	FilterConfigSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*Any2EVMTollOffRampHelperConfigSet, error)

	FilterExecutionCompleted(opts *bind.FilterOpts, sequenceNumber []uint64) (*Any2EVMTollOffRampHelperExecutionCompletedIterator, error)

	WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperExecutionCompleted, sequenceNumber []uint64) (event.Subscription, error)

	ParseExecutionCompleted(log types.Log) (*Any2EVMTollOffRampHelperExecutionCompleted, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*Any2EVMTollOffRampHelperOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*Any2EVMTollOffRampHelperOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperOffRampRouterSet, router []common.Address) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*Any2EVMTollOffRampHelperOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampHelperOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*Any2EVMTollOffRampHelperOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampHelperOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*Any2EVMTollOffRampHelperOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*Any2EVMTollOffRampHelperPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*Any2EVMTollOffRampHelperPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*Any2EVMTollOffRampHelperPoolRemoved, error)

	FilterTransmitted(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*Any2EVMTollOffRampHelperTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*Any2EVMTollOffRampHelperUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
