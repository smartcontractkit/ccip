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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SubscriptionNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlobVerifier\",\"outputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_sourceChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"s_receiverToNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_router\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"}],\"name\":\"setBlobVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005c3a38038062005c3a833981016040819052620000349162000608565b6000805460ff19168155600190899089908990899089908990899089908290829085903390819081620000ae5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000e857620000e8816200030e565b5050506001600160a01b0381166200011357604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03929092169190911790558051825114620001555760405162d8548360e71b815260040160405180910390fd5b81516200016a906004906020850190620003bf565b5060005b82518110156200024c5760008282815181106200018f576200018f62000738565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001d957620001d962000738565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600590925220805460ff1916600117905562000244816200074e565b90506200016e565b50505060809790975250505060a0929092528051600880546020840151604085015160609095015163ffffffff16600160c01b0263ffffffff60c01b196001600160401b03968716600160801b0216600160801b600160e01b031992871668010000000000000000026001600160801b03199094169690951695909517919091171691909117919091179055600780546001600160a01b039092166001600160a01b031990921691909117905550151560c05250620007769650505050505050565b336001600160a01b03821603620003685760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a5565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b82805482825590600052602060002090810192821562000417579160200282015b828111156200041757825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003e0565b506200042592915062000429565b5090565b5b808211156200042557600081556001016200042a565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b03811182821017156200047b576200047b62000440565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620004ac57620004ac62000440565b604052919050565b80516001600160401b0381168114620004cc57600080fd5b919050565b6001600160a01b0381168114620004e757600080fd5b50565b8051620004cc81620004d1565b60006001600160401b0382111562000513576200051362000440565b5060051b60200190565b600082601f8301126200052f57600080fd5b81516020620005486200054283620004f7565b62000481565b82815260059290921b840181019181810190868411156200056857600080fd5b8286015b84811015620005905780516200058281620004d1565b83529183019183016200056c565b509695505050505050565b600082601f830112620005ad57600080fd5b81516020620005c06200054283620004f7565b82815260059290921b84018101918181019086841115620005e057600080fd5b8286015b8481101562000590578051620005fa81620004d1565b8352918301918301620005e4565b600080600080600080600080888a036101608112156200062757600080fd5b8951985060208a015197506080603f19820112156200064557600080fd5b506200065062000456565b6200065e60408b01620004b4565b81526200066e60608b01620004b4565b60208201526200068160808b01620004b4565b604082015260a08a015163ffffffff811681146200069e57600080fd5b60608201529550620006b360c08a01620004ea565b9450620006c360e08a01620004ea565b9350620006d46101008a01620004ea565b6101208a01519093506001600160401b0380821115620006f357600080fd5b620007018c838d016200051d565b93506101408b01519150808211156200071957600080fd5b50620007288b828c016200059b565b9150509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b6000600182016200076f57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05161548d620007ad60003960006111ad0152600061038f01526000818161026601526135fb015261548d6000f3fe608060405234801561001057600080fd5b506004361061025c5760003560e01c806389c0656811610145578063be9b03f1116100bd578063e3d0e7121161008c578063eb511dd411610071578063eb511dd414610714578063f2fde38b14610727578063fb777fc71461073a57600080fd5b8063e3d0e712146106f3578063e402bc541461070657600080fd5b8063be9b03f1146105c9578063c0d78655146105dc578063c3f909d4146105ef578063e16e632c146106d357600080fd5b8063afcb95d711610114578063b1dc65a4116100f9578063b1dc65a41461056a578063b4069b311461057d578063bbe4f6db1461059057600080fd5b8063afcb95d71461052c578063b0f479a11461054c57600080fd5b806389c06568146104db5780638b721c71146104e35780638da5cb5b146104f6578063a639d1c01461051957600080fd5b80635c975abb116101d857806379ba5097116101a7578063814118341161018c578063814118341461048e57806381ff7048146104a35780638456cb59146104d357600080fd5b806379ba5097146104435780637c34718c1461044b57600080fd5b80635c975abb146103f25780636133dc24146103fd578063681fba161461041b578063744b92e21461043057600080fd5b8063181f5a771161022f57806328094b591161021457806328094b591461038a5780633f4ba83a146103b15780635b16ebb7146103b957600080fd5b8063181f5a77146103025780632222dd421461034b57600080fd5b80630f733bca14610261578063108ee5fc1461029b578063142a98fc146102b0578063147809b3146102ea575b600080fd5b6102887f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b6102ae6102a9366004613cb3565b61074d565b005b6102dd6102be366004613cf1565b67ffffffffffffffff1660009081526009602052604090205460ff1690565b6040516102929190613d3d565b6102f2610828565b6040519015158152602001610292565b61033e6040518060400160405280602081526020017f45564d3245564d537562736372697074696f6e4f666652616d7020312e302e3081525081565b6040516102929190613df4565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610292565b6102887f000000000000000000000000000000000000000000000000000000000000000081565b6102ae6108c2565b6102f26103c7366004613cb3565b73ffffffffffffffffffffffffffffffffffffffff1660009081526005602052604090205460ff1690565b60005460ff166102f2565b60075473ffffffffffffffffffffffffffffffffffffffff16610365565b6104236108d4565b6040516102929190613e12565b6102ae61043e366004613e6c565b6109b3565b6102ae610db2565b610475610459366004613cb3565b60106020526000908152604090205467ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610292565b610496610ed9565b6040516102929190613ef6565b600c54600a546040805163ffffffff80851682526401000000009094049093166020840152820152606001610292565b6102ae610f48565b610423610f58565b6102ae6104f136600461415f565b610fc5565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610365565b6102ae610527366004613cb3565b611021565b604080516001815260006020820181905291810191909152606001610292565b60065473ffffffffffffffffffffffffffffffffffffffff16610365565b6102ae6105783660046142be565b611070565b61036561058b366004613cb3565b611717565b61036561059e366004613cb3565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600360205260409020541690565b6102ae6105d7366004614504565b61181f565b6102ae6105ea366004613cb3565b61221b565b61068c604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260085467ffffffffffffffff80821683526801000000000000000082048116602084015270010000000000000000000000000000000082041692820192909252780100000000000000000000000000000000000000000000000090910463ffffffff16606082015290565b60408051825167ffffffffffffffff908116825260208085015182169083015283830151169181019190915260609182015163ffffffff1691810191909152608001610292565b6006546103659073ffffffffffffffffffffffffffffffffffffffff1681565b6102ae61070136600461465e565b612292565b6102ae61025c36600461472b565b6102ae610722366004613e6c565b612c6d565b6102ae610735366004613cb3565b612eb5565b6102ae610748366004614767565b612ec6565b61075561300b565b73ffffffffffffffffffffffffffffffffffffffff81166107a2576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28910160405180910390a15050565b600254604080517f46f8e6d7000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff16916346f8e6d79160048083019260209291908290030181865afa158015610898573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bc91906147f5565b15905090565b6108ca61300b565b6108d2613091565b565b60045460609067ffffffffffffffff8111156108f2576108f2613f09565b60405190808252806020026020018201604052801561091b578160200160208202803683370190505b50905060005b6004548110156109af576109686004828154811061094157610941614812565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16611717565b82828151811061097a5761097a614812565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101526109a881614870565b9050610921565b5090565b6109bb61300b565b60045460008190036109f9576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610a94576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610afd576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004610b0c6001856148a8565b81548110610b1c57610b1c614812565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600482602001516bffffffffffffffffffffffff1681548110610b6e57610b6e614812565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166004610b9d6001866148a8565b81548110610bad57610bad614812565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600483602001516bffffffffffffffffffffffff1681548110610c1b57610c1b614812565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556004805480610cbd57610cbd6148bf565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600383526040808520859055918816808552600584529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610e38576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600f805480602002602001604051908101604052809291908181526020018280548015610f3e57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610f13575b5050505050905090565b610f5061300b565b6108d2613172565b60606004805480602002602001604051908101604052809291908181526020018280548015610f3e5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610f13575050505050905090565b333014610ffe576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110158160a001518260c001518360600151613232565b61101e816132cc565b50565b61102961300b565b600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916110c691849163ffffffff851691908e908e90819084018382808284376000920191909152506133be92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600a54808252600b5460ff8082166020850152610100909104169282019290925290831461119b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610e2f565b6111a98b8b8b8b8b8b613444565b60007f000000000000000000000000000000000000000000000000000000000000000015611206576002826020015183604001516111e791906148ee565b6111f19190614942565b6111fc9060016148ee565b60ff16905061121c565b60208201516112169060016148ee565b60ff1690505b888114611285576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610e2f565b8887146112ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610e2f565b336000908152600d602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561133157611331613d0e565b600281111561134257611342613d0e565b905250905060028160200151600281111561135f5761135f613d0e565b1480156113a65750600f816000015160ff168154811061138157611381614812565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b61140c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610e2f565b505050505060008888604051611423929190614964565b60405190819003812061143a918c90602001614974565b60405160208183030381529060405280519060200120905061145a613c72565b604080518082019091526000808252602082015260005b888110156116f557600060018588846020811061149057611490614812565b61149d91901a601b6148ee565b8d8d868181106114af576114af614812565b905060200201358c8c878181106114c8576114c8614812565b9050602002013560405160008152602001604052604051611505949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611527573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600d602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156115a7576115a7613d0e565b60028111156115b8576115b8613d0e565b90525092506001836020015160028111156115d5576115d5613d0e565b1461163c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610e2f565b8251849060ff16601f811061165357611653614812565b6020020151156116bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610e2f565b600184846000015160ff16601f81106116da576116da614812565b91151560209092020152506116ee81614870565b9050611471565b5050505063ffffffff811061170c5761170c614990565b505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff80821660009081526003602052604081205490911680611778576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa1580156117f4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061181891906149bf565b9392505050565b60005460ff161561188c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610e2f565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061191d91906147f5565b15611953576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60065473ffffffffffffffffffffffffffffffffffffffff16806119a3576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608301515160008190036119e4576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff8111156119ff576119ff613f09565b604051908082528060200260200182016040528015611a28578160200160208202803683370190505b50905060008267ffffffffffffffff811115611a4657611a46613f09565b604051908082528060200260200182016040528015611ada57816020015b60408051610120810182526000808252602080830182905292820181905260608083018290526080830182905260a0830181905260c0830181905260e083015261010082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff909201910181611a645790505b50905060005b83811015611ba75786606001518181518110611afe57611afe614812565b6020026020010151806020019051810190611b199190614afe565b828281518110611b2b57611b2b614812565b602002602001018190525086606001518181518110611b4c57611b4c614812565b6020026020010151604051602001611b649190614c0f565b60405160208183030381529060405280519060200120838281518110611b8c57611b8c614812565b6020908102919091010152611ba081614870565b9050611ae0565b50600080611bc88489608001518a60a001518b60c001518c60e001516134f2565b915091506000835182611bdb9190614c35565b9050878015611c1657506008547801000000000000000000000000000000000000000000000000900463ffffffff16611c1484426148a8565b105b15611c4d576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008960400151600081518110611c6657611c66614812565b6020026020010151905060005b8781101561220e5760005a90506000878381518110611c9457611c94614812565b602002602001015190506000611cc7826020015167ffffffffffffffff1660009081526009602052604090205460ff1690565b90506002816003811115611cdd57611cdd613d0e565b03611d265760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610e2f565b60608201516040517f0cbebc2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526000918e1690630cbebc2490602401600060405180830381865afa158015611d9a573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052611de09190810190614c49565b602081015190915073ffffffffffffffffffffffffffffffffffffffff16611e525760608301516040517f8515736a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610e2f565b6080830151606084015173ffffffffffffffffffffffffffffffffffffffff16600090815260106020526040812054909167ffffffffffffffff90811691611e9c91166001614d05565b67ffffffffffffffff161490508080611ed357508160400151158015611ed357506003836003811115611ed157611ed1613d0e565b145b611f1b5760808401516040517fb0241f4a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610e2f565b611f24846135f7565b60005b8460c0015151811015611f6b57611f5a8560c001518281518110611f4d57611f4d614812565b602002602001015161375c565b50611f6481614870565b9050611f27565b5060208481015167ffffffffffffffff16600090815260099091526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611fc4611fbf866137d8565b6138fd565b60208087015167ffffffffffffffff166000908152600990915260409020805491925082917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600183600381111561201f5761201f613d0e565b0217905550846020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf826040516120629190613d3d565b60405180910390a28180156120945750600281600381111561208657612086613d0e565b148061209457508260400151155b1561210057606085015173ffffffffffffffffffffffffffffffffffffffff166000908152601060205260408120805467ffffffffffffffff16916120d883614d31565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b8f6121f7578e73ffffffffffffffffffffffffffffffffffffffff1663294d266186606001518760400151670de0b6b3a76400008c3a8f5a612142908f6148a8565b61214c9190614d58565b6121569190614d70565b6121609190614d70565b61216a9190614c35565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815273ffffffffffffffffffffffffffffffffffffffff93841660048201529290911660248301526044820152606401600060405180830381600087803b1580156121de57600080fd5b505af11580156121f2573d6000803e3d6000fd5b505050505b5050505050508061220790614870565b9050611c73565b5050505050505050505050565b61222361300b565b600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b855185518560ff16601f831115612305576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610e2f565b6000811161236f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610e2f565b8183146123fd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610e2f565b612408816003614d70565b8311612470576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610e2f565b61247861300b565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600e541561266b57600e546000906124d0906001906148a8565b90506000600e82815481106124e7576124e7614812565b6000918252602082200154600f805473ffffffffffffffffffffffffffffffffffffffff9092169350908490811061252157612521614812565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600d909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600e805491925090806125a1576125a16148bf565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600f80548061260a5761260a6148bf565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055506124b6915050565b60005b815151811015612ad0576000600d60008460000151848151811061269457612694614812565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff1660028111156126de576126de613d0e565b14612745576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610e2f565b6040805180820190915260ff821681526001602082015282518051600d916000918590811061277657612776614812565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561281757612817613d0e565b0217905550600091506128279050565b600d60008460200151848151811061284157612841614812565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561288b5761288b613d0e565b146128f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610e2f565b6040805180820190915260ff82168152602081016002815250600d60008460200151848151811061292557612925614812565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156129c6576129c6613d0e565b02179055505082518051600e9250839081106129e4576129e4614812565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600f919083908110612a6057612a60614812565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055612ac981614870565b905061266e565b506040810151600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600c80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092612b62928692908216911617614dad565b92506101000a81548163ffffffff021916908363ffffffff160217905550612bc14630600c60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613a2f565b600a81905582518051600b805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600c5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612c60988b98919763ffffffff909216969095919491939192614dcc565b60405180910390a161220e565b612c7561300b565b73ffffffffffffffffffffffffffffffffffffffff82161580612cac575073ffffffffffffffffffffffffffffffffffffffff8116155b15612ce3576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612d7f576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600581529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b612ebd61300b565b61101e81613ada565b612ece61300b565b80516008805460208085018051604080880180516060808b01805163ffffffff9081167801000000000000000000000000000000000000000000000000027fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff95861670010000000000000000000000000000000002167fffffffff000000000000000000000000ffffffffffffffffffffffffffffffff98861668010000000000000000027fffffffffffffffffffffffffffffffff00000000000000000000000000000000909c169d86169d8e179b909b17979097169990991795909517909855825198895293518416948801949094529251909116918501919091525116908201527f187b05a44b331bbf26fcef5672820acbcd72e33ec0214a0b89042db7b73b46a29060800160405180910390a150565b600054610100900473ffffffffffffffffffffffffffffffffffffffff1633146108d2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610e2f565b60005460ff166130fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610e2f565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60005460ff16156131df576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610e2f565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586131483390565b815183511461326d576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b83518110156132c6576132b684828151811061328e5761328e614812565b60200260200101518483815181106132a8576132a8614812565b602002602001015184613bd5565b6132bf81614870565b9050613270565b50505050565b606081015173ffffffffffffffffffffffffffffffffffffffff163b6132ef5750565b6006546040517fc39a285b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063c39a285b90613345908490600401614e92565b6020604051808303816000875af1158015613364573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061338891906147f5565b61101e576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff1663be9b03f1828060200190518101906133ee9190615071565b60006040518363ffffffff1660e01b815260040161340d929190615239565b600060405180830381600087803b15801561342757600080fd5b505af115801561343b573d6000803e3d6000fd5b50505050505050565b6000613451826020614d70565b61345c856020614d70565b61346888610144614d58565b6134729190614d58565b61347c9190614d58565b613487906000614d58565b905036811461343b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610e2f565b60008060005a6007546040517fe71e65ce00000000000000000000000000000000000000000000000000000000815291925060009173ffffffffffffffffffffffffffffffffffffffff9091169063e71e65ce9061355c908c908c908c908c908c90600401615330565b6020604051808303816000875af115801561357b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061359f9190615382565b9050600081116135db576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a6135e790846148a8565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146136575780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610e2f565b60085460c08201515170010000000000000000000000000000000090910467ffffffffffffffff16108061369557508060e00151518160c001515114155b156136de5760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610e2f565b60085460a0820151516801000000000000000090910467ffffffffffffffff16101561101e5760085460a0820151516040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090920467ffffffffffffffff1660048301526024820152604401610e2f565b73ffffffffffffffffffffffffffffffffffffffff81811660009081526003602052604090205416806137d3576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610e2f565b919050565b61384060405180610100016040528060008152602001600067ffffffffffffffff16815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001600081525090565b60405180610100016040528083600001518152602001836020015167ffffffffffffffff168152602001836040015160405160200161389b919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b6040516020818303038152906040528152602001836060015173ffffffffffffffffffffffffffffffffffffffff1681526020018360a0015181526020018360c0015181526020018360e0015181526020018361010001518152509050919050565b6040517f8b721c710000000000000000000000000000000000000000000000000000000081526000903090638b721c719061393c908590600401614e92565b600060405180830381600087803b15801561395657600080fd5b505af1925050508015613967575060015b613a27573d808015613995576040519150601f19603f3d011682016040523d82523d6000602084013e61399a565b606091505b506139a48161539b565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da800000000000000000000000000000000000000000000000000000000036139f55750600392915050565b6040517f2532cf4500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506002919050565b6000808a8a8a8a8a8a8a8a8a604051602001613a53999897969594939291906153eb565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603613b59576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610e2f565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000613be08461375c565b6040517fea6192a200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690529192509082169063ea6192a290604401600060405180830381600087803b158015613c5457600080fd5b505af1158015613c68573d6000803e3d6000fd5b5050505050505050565b604051806103e00160405280601f906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff8116811461101e57600080fd5b600060208284031215613cc557600080fd5b813561181881613c91565b67ffffffffffffffff8116811461101e57600080fd5b80356137d381613cd0565b600060208284031215613d0357600080fd5b813561181881613cd0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310613d78577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b83811015613d99578181015183820152602001613d81565b838111156132c65750506000910152565b60008151808452613dc2816020860160208601613d7e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006118186020830184613daa565b80356137d381613c91565b6020808252825182820181905260009190848201906040850190845b81811015613e6057835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101613e2e565b50909695505050505050565b60008060408385031215613e7f57600080fd5b8235613e8a81613c91565b91506020830135613e9a81613c91565b809150509250929050565b600081518084526020808501945080840160005b83811015613eeb57815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101613eb9565b509495945050505050565b6020815260006118186020830184613ea5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff81118282101715613f5c57613f5c613f09565b60405290565b604051610120810167ffffffffffffffff81118282101715613f5c57613f5c613f09565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613fcd57613fcd613f09565b604052919050565b600067ffffffffffffffff821115613fef57613fef613f09565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261402c57600080fd5b813561403f61403a82613fd5565b613f86565b81815284602083860101111561405457600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff82111561408b5761408b613f09565b5060051b60200190565b600082601f8301126140a657600080fd5b813560206140b661403a83614071565b82815260059290921b840181019181810190868411156140d557600080fd5b8286015b848110156140f95780356140ec81613c91565b83529183019183016140d9565b509695505050505050565b600082601f83011261411557600080fd5b8135602061412561403a83614071565b82815260059290921b8401810191818101908684111561414457600080fd5b8286015b848110156140f95780358352918301918301614148565b60006020828403121561417157600080fd5b813567ffffffffffffffff8082111561418957600080fd5b90830190610100828603121561419e57600080fd5b6141a6613f38565b823581526141b660208401613ce6565b60208201526040830135828111156141cd57600080fd5b6141d98782860161401b565b6040830152506141eb60608401613e07565b606082015260808301358281111561420257600080fd5b61420e8782860161401b565b60808301525060a08301358281111561422657600080fd5b61423287828601614095565b60a08301525060c08301358281111561424a57600080fd5b61425687828601614104565b60c08301525060e083013560e082015280935050505092915050565b60008083601f84011261428457600080fd5b50813567ffffffffffffffff81111561429c57600080fd5b6020830191508360208260051b85010111156142b757600080fd5b9250929050565b60008060008060008060008060e0898b0312156142da57600080fd5b606089018a8111156142eb57600080fd5b8998503567ffffffffffffffff8082111561430557600080fd5b818b0191508b601f83011261431957600080fd5b81358181111561432857600080fd5b8c602082850101111561433a57600080fd5b6020830199508098505060808b013591508082111561435857600080fd5b6143648c838d01614272565b909750955060a08b013591508082111561437d57600080fd5b5061438a8b828c01614272565b999c989b50969995989497949560c00135949350505050565b600082601f8301126143b457600080fd5b813560206143c461403a83614071565b82815260059290921b840181019181810190868411156143e357600080fd5b8286015b848110156140f95780356143fa81613cd0565b83529183019183016143e7565b600082601f83011261441857600080fd5b8135602061442861403a83614071565b82815260059290921b8401810191818101908684111561444757600080fd5b8286015b848110156140f957803561445e81613c91565b835291830191830161444b565b600082601f83011261447c57600080fd5b8135602061448c61403a83614071565b82815260059290921b840181019181810190868411156144ab57600080fd5b8286015b848110156140f957803567ffffffffffffffff8111156144cf5760008081fd5b6144dd8986838b010161401b565b8452509183019183016144af565b801515811461101e57600080fd5b80356137d3816144eb565b6000806040838503121561451757600080fd5b823567ffffffffffffffff8082111561452f57600080fd5b90840190610100828703121561454457600080fd5b61454c613f38565b82358281111561455b57600080fd5b614567888286016143a3565b82525060208301358281111561457c57600080fd5b61458888828601614407565b6020830152506040830135828111156145a057600080fd5b6145ac88828601614104565b6040830152506060830135828111156145c457600080fd5b6145d08882860161446b565b6060830152506080830135828111156145e857600080fd5b6145f488828601614104565b60808301525060a083013560a082015260c08301358281111561461657600080fd5b61462288828601614104565b60c08301525060e083013560e0820152809450505050614644602084016144f9565b90509250929050565b803560ff811681146137d357600080fd5b60008060008060008060c0878903121561467757600080fd5b863567ffffffffffffffff8082111561468f57600080fd5b61469b8a838b01614407565b975060208901359150808211156146b157600080fd5b6146bd8a838b01614407565b96506146cb60408a0161464d565b955060608901359150808211156146e157600080fd5b6146ed8a838b0161401b565b94506146fb60808a01613ce6565b935060a089013591508082111561471157600080fd5b5061471e89828a0161401b565b9150509295509295509295565b60006020828403121561473d57600080fd5b813567ffffffffffffffff81111561475457600080fd5b8201610100818503121561181857600080fd5b60006080828403121561477957600080fd5b6040516080810181811067ffffffffffffffff8211171561479c5761479c613f09565b60405282356147aa81613cd0565b815260208301356147ba81613cd0565b602082015260408301356147cd81613cd0565b6040820152606083013563ffffffff811681146147e957600080fd5b60608201529392505050565b60006020828403121561480757600080fd5b8151611818816144eb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036148a1576148a1614841565b5060010190565b6000828210156148ba576148ba614841565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff0382111561490b5761490b614841565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff83168061495557614955614913565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b6000602082840312156149d157600080fd5b815161181881613c91565b80516137d381613cd0565b80516137d381613c91565b600082601f830112614a0357600080fd5b8151614a1161403a82613fd5565b818152846020838601011115614a2657600080fd5b614a37826020830160208701613d7e565b949350505050565b600082601f830112614a5057600080fd5b81516020614a6061403a83614071565b82815260059290921b84018101918181019086841115614a7f57600080fd5b8286015b848110156140f9578051614a9681613c91565b8352918301918301614a83565b600082601f830112614ab457600080fd5b81516020614ac461403a83614071565b82815260059290921b84018101918181019086841115614ae357600080fd5b8286015b848110156140f95780518352918301918301614ae7565b600060208284031215614b1057600080fd5b815167ffffffffffffffff80821115614b2857600080fd5b908301906101208286031215614b3d57600080fd5b614b45613f62565b82518152614b55602084016149dc565b6020820152614b66604084016149e7565b6040820152614b77606084016149e7565b6060820152614b88608084016149dc565b608082015260a083015182811115614b9f57600080fd5b614bab878286016149f2565b60a08301525060c083015182811115614bc357600080fd5b614bcf87828601614a3f565b60c08301525060e083015182811115614be757600080fd5b614bf387828601614aa3565b60e0830152506101009283015192810192909252509392505050565b6000815260008251614c28816001850160208701613d7e565b9190910160010192915050565b600082614c4457614c44614913565b500490565b600060208284031215614c5b57600080fd5b815167ffffffffffffffff80821115614c7357600080fd5b9083019060808286031215614c8757600080fd5b604051608081018181108382111715614ca257614ca2613f09565b604052825182811115614cb457600080fd5b614cc087828601614a3f565b82525060208301519150614cd382613c91565b81602082015260408301519150614ce9826144eb565b8160408201526060830151606082015280935050505092915050565b600067ffffffffffffffff808316818516808303821115614d2857614d28614841565b01949350505050565b600067ffffffffffffffff808316818103614d4e57614d4e614841565b6001019392505050565b60008219821115614d6b57614d6b614841565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614da857614da8614841565b500290565b600063ffffffff808316818516808303821115614d2857614d28614841565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614dfc8184018a613ea5565b90508281036080840152614e108189613ea5565b905060ff871660a084015282810360c0840152614e2d8187613daa565b905067ffffffffffffffff851660e0840152828103610100840152614e528185613daa565b9c9b505050505050505050505050565b600081518084526020808501945080840160005b83811015613eeb57815187529582019590820190600101614e76565b602081528151602082015260006020830151614eba604084018267ffffffffffffffff169052565b506040830151610100806060850152614ed7610120850183613daa565b91506060850151614f00608086018273ffffffffffffffffffffffffffffffffffffffff169052565b5060808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808685030160a0870152614f3b8483613daa565b935060a08701519150808685030160c0870152614f588483613ea5565b935060c08701519150808685030160e087015250614f768382614e62565b92505060e085015181850152508091505092915050565b600082601f830112614f9e57600080fd5b81516020614fae61403a83614071565b82815260059290921b84018101918181019086841115614fcd57600080fd5b8286015b848110156140f9578051614fe481613cd0565b8352918301918301614fd1565b600082601f83011261500257600080fd5b8151602061501261403a83614071565b82815260059290921b8401810191818101908684111561503157600080fd5b8286015b848110156140f957805167ffffffffffffffff8111156150555760008081fd5b6150638986838b01016149f2565b845250918301918301615035565b60006020828403121561508357600080fd5b815167ffffffffffffffff8082111561509b57600080fd5b9083019061010082860312156150b057600080fd5b6150b8613f38565b8251828111156150c757600080fd5b6150d387828601614f8d565b8252506020830151828111156150e857600080fd5b6150f487828601614a3f565b60208301525060408301518281111561510c57600080fd5b61511887828601614aa3565b60408301525060608301518281111561513057600080fd5b61513c87828601614ff1565b60608301525060808301518281111561515457600080fd5b61516087828601614aa3565b60808301525060a083015160a082015260c08301518281111561518257600080fd5b61518e87828601614aa3565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b83811015613eeb57815167ffffffffffffffff16875295820195908201906001016151be565b600081518084526020808501808196508360051b8101915082860160005b8581101561522c57828403895261521a848351613daa565b98850198935090840190600101615202565b5091979650505050505050565b60408152600083516101008060408501526152586101408501836151aa565b915060208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808685030160608701526152948483613ea5565b935060408801519150808685030160808701526152b18483614e62565b935060608801519150808685030160a08701526152ce84836151e4565b935060808801519150808685030160c08701526152eb8483614e62565b935060a088015160e087015260c08801519150808685030183870152506153128382614e62565b60e08801516101208701528615156020870152935061181892505050565b60a08152600061534360a0830188614e62565b82810360208401526153558188614e62565b9050856040840152828103606084015261536f8186614e62565b9150508260808301529695505050505050565b60006020828403121561539457600080fd5b5051919050565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156153e35780818460040360031b1b83161693505b505050919050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526154328285018b613ea5565b91508382036080850152615446828a613ea5565b915060ff881660a085015283820360c08501526154638288613daa565b90861660e08501528381036101008501529050614e528185613daa56fea164736f6c634300080f000a",
}

var EVM2EVMSubscriptionOffRampABI = EVM2EVMSubscriptionOffRampMetaData.ABI

var EVM2EVMSubscriptionOffRampBin = EVM2EVMSubscriptionOffRampMetaData.Bin

func DeployEVM2EVMSubscriptionOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, blobVerifier common.Address, onRampAddress common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address) (common.Address, *types.Transaction, *EVM2EVMSubscriptionOffRamp, error) {
	parsed, err := EVM2EVMSubscriptionOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMSubscriptionOffRampBin), backend, sourceChainId, chainId, offRampConfig, blobVerifier, onRampAddress, afn, sourceTokens, pools)
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

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) IChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "i_chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) IChainId() (*big.Int, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.IChainId(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) IChainId() (*big.Int, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.IChainId(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCaller) ISourceChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOffRamp.contract.Call(opts, &out, "i_sourceChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampSession) ISourceChainId() (*big.Int, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.ISourceChainId(&_EVM2EVMSubscriptionOffRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOffRamp *EVM2EVMSubscriptionOffRampCallerSession) ISourceChainId() (*big.Int, error) {
	return _EVM2EVMSubscriptionOffRamp.Contract.ISourceChainId(&_EVM2EVMSubscriptionOffRamp.CallOpts)
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
	CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessage) error

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetBlobVerifier(opts *bind.CallOpts) (common.Address, error)

	GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IChainId(opts *bind.CallOpts) (*big.Int, error)

	ISourceChainId(opts *bind.CallOpts) (*big.Int, error)

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

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

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

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
