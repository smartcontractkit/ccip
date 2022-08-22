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

var EVM2EVMTollOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlobVerifier\",\"outputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_sourceChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_router\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"}],\"name\":\"setBlobVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620059ab380380620059ab833981016040819052620000349162000608565b6000805460ff19168155600190899089908990899089908990899089908290829085903390819081620000ae5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000e857620000e8816200030e565b5050506001600160a01b0381166200011357604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03929092169190911790558051825114620001555760405162d8548360e71b815260040160405180910390fd5b81516200016a906004906020850190620003bf565b5060005b82518110156200024c5760008282815181106200018f576200018f62000738565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001d957620001d962000738565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600590925220805460ff1916600117905562000244816200074e565b90506200016e565b50505060809790975250505060a0929092528051600880546020840151604085015160609095015163ffffffff16600160c01b0263ffffffff60c01b196001600160401b03968716600160801b0216600160801b600160e01b031992871668010000000000000000026001600160801b03199094169690951695909517919091171691909117919091179055600780546001600160a01b039092166001600160a01b031990921691909117905550151560c05250620007769650505050505050565b336001600160a01b03821603620003685760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a5565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b82805482825590600052602060002090810192821562000417579160200282015b828111156200041757825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003e0565b506200042592915062000429565b5090565b5b808211156200042557600081556001016200042a565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b03811182821017156200047b576200047b62000440565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620004ac57620004ac62000440565b604052919050565b80516001600160401b0381168114620004cc57600080fd5b919050565b6001600160a01b0381168114620004e757600080fd5b50565b8051620004cc81620004d1565b60006001600160401b0382111562000513576200051362000440565b5060051b60200190565b600082601f8301126200052f57600080fd5b81516020620005486200054283620004f7565b62000481565b82815260059290921b840181019181810190868411156200056857600080fd5b8286015b84811015620005905780516200058281620004d1565b83529183019183016200056c565b509695505050505050565b600082601f830112620005ad57600080fd5b81516020620005c06200054283620004f7565b82815260059290921b84018101918181019086841115620005e057600080fd5b8286015b8481101562000590578051620005fa81620004d1565b8352918301918301620005e4565b600080600080600080600080888a036101608112156200062757600080fd5b8951985060208a015197506080603f19820112156200064557600080fd5b506200065062000456565b6200065e60408b01620004b4565b81526200066e60608b01620004b4565b60208201526200068160808b01620004b4565b604082015260a08a015163ffffffff811681146200069e57600080fd5b60608201529550620006b360c08a01620004ea565b9450620006c360e08a01620004ea565b9350620006d46101008a01620004ea565b6101208a01519093506001600160401b0380821115620006f357600080fd5b620007018c838d016200051d565b93506101408b01519150808211156200071957600080fd5b50620007288b828c016200059b565b9150509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b6000600182016200076f57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c0516151fe620007ad600039600061114f0152600061037401526000818161024b015261346701526151fe6000f3fe608060405234801561001057600080fd5b50600436106102415760003560e01c806389c0656811610145578063be9b03f1116100bd578063e3d0e7121161008c578063eb511dd411610071578063eb511dd4146106b6578063f2fde38b146106c9578063fb777fc7146106dc57600080fd5b8063e3d0e71214610695578063e402bc54146106a857600080fd5b8063be9b03f11461056b578063c0d786551461057e578063c3f909d414610591578063e16e632c1461067557600080fd5b8063afcb95d711610114578063b1dc65a4116100f9578063b1dc65a41461050c578063b4069b311461051f578063bbe4f6db1461053257600080fd5b8063afcb95d7146104ce578063b0f479a1146104ee57600080fd5b806389c065681461047d5780638b721c71146104855780638da5cb5b14610498578063a639d1c0146104bb57600080fd5b80635b16ebb7116101d8578063744b92e2116101a7578063814118341161018c578063814118341461043057806381ff7048146104455780638456cb591461047557600080fd5b8063744b92e21461041557806379ba50971461042857600080fd5b80635b16ebb71461039e5780635c975abb146103d75780636133dc24146103e2578063681fba161461040057600080fd5b8063181f5a7711610214578063181f5a77146102e75780632222dd421461033057806328094b591461036f5780633f4ba83a1461039657600080fd5b80630f733bca14610246578063108ee5fc14610280578063142a98fc14610295578063147809b3146102cf575b600080fd5b61026d7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b61029361028e366004613b1f565b6106ef565b005b6102c26102a3366004613b5d565b67ffffffffffffffff1660009081526009602052604090205460ff1690565b6040516102779190613ba9565b6102d76107ca565b6040519015158152602001610277565b6103236040518060400160405280601881526020017f45564d3245564d546f6c6c4f666652616d7020312e302e30000000000000000081525081565b6040516102779190613c60565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610277565b61026d7f000000000000000000000000000000000000000000000000000000000000000081565b610293610864565b6102d76103ac366004613b1f565b73ffffffffffffffffffffffffffffffffffffffff1660009081526005602052604090205460ff1690565b60005460ff166102d7565b60075473ffffffffffffffffffffffffffffffffffffffff1661034a565b610408610876565b6040516102779190613c7e565b610293610423366004613cd8565b610955565b610293610d54565b610438610e7b565b6040516102779190613d62565b600c54600a546040805163ffffffff80851682526401000000009094049093166020840152820152606001610277565b610293610eea565b610408610efa565b610293610493366004613fcb565b610f67565b600054610100900473ffffffffffffffffffffffffffffffffffffffff1661034a565b6102936104c9366004613b1f565b610fc3565b604080516001815260006020820181905291810191909152606001610277565b60065473ffffffffffffffffffffffffffffffffffffffff1661034a565b61029361051a36600461412a565b611012565b61034a61052d366004613b1f565b6116b9565b61034a610540366004613b1f565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600360205260409020541690565b61029361057936600461430c565b6117c1565b61029361058c366004613b1f565b61207f565b61062e604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260085467ffffffffffffffff80821683526801000000000000000082048116602084015270010000000000000000000000000000000082041692820192909252780100000000000000000000000000000000000000000000000090910463ffffffff16606082015290565b60408051825167ffffffffffffffff908116825260208085015182169083015283830151169181019190915260609182015163ffffffff1691810191909152608001610277565b60065461034a9073ffffffffffffffffffffffffffffffffffffffff1681565b6102936106a3366004614466565b6120f6565b610293610241366004614533565b6102936106c4366004613cd8565b612ad9565b6102936106d7366004613b1f565b612d21565b6102936106ea36600461456f565b612d32565b6106f7612e77565b73ffffffffffffffffffffffffffffffffffffffff8116610744576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28910160405180910390a15050565b600254604080517f46f8e6d7000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff16916346f8e6d79160048083019260209291908290030181865afa15801561083a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085e91906145fd565b15905090565b61086c612e77565b610874612efd565b565b60045460609067ffffffffffffffff81111561089457610894613d75565b6040519080825280602002602001820160405280156108bd578160200160208202803683370190505b50905060005b6004548110156109515761090a600482815481106108e3576108e361461a565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166116b9565b82828151811061091c5761091c61461a565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015261094a81614678565b90506108c3565b5090565b61095d612e77565b600454600081900361099b576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610a36576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610a9f576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004610aae6001856146b0565b81548110610abe57610abe61461a565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600482602001516bffffffffffffffffffffffff1681548110610b1057610b1061461a565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff166004610b3f6001866146b0565b81548110610b4f57610b4f61461a565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600483602001516bffffffffffffffffffffffff1681548110610bbd57610bbd61461a565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556004805480610c5f57610c5f6146c7565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600383526040808520859055918816808552600584529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610dda576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600f805480602002602001604051908101604052809291908181526020018280548015610ee057602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610eb5575b5050505050905090565b610ef2612e77565b610874612fde565b60606004805480602002602001604051908101604052809291908181526020018280548015610ee05760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610eb5575050505050905090565b333014610fa0576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fb78160a001518260c00151836060015161309e565b610fc081613138565b50565b610fcb612e77565b600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161106891849163ffffffff851691908e908e908190840183828082843760009201919091525061322a92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600a54808252600b5460ff8082166020850152610100909104169282019290925290831461113d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610dd1565b61114b8b8b8b8b8b8b6132b0565b60007f0000000000000000000000000000000000000000000000000000000000000000156111a85760028260200151836040015161118991906146f6565b611193919061474a565b61119e9060016146f6565b60ff1690506111be565b60208201516111b89060016146f6565b60ff1690505b888114611227576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610dd1565b888714611290576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610dd1565b336000908152600d602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156112d3576112d3613b7a565b60028111156112e4576112e4613b7a565b905250905060028160200151600281111561130157611301613b7a565b1480156113485750600f816000015160ff16815481106113235761132361461a565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b6113ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610dd1565b5050505050600088886040516113c592919061476c565b6040519081900381206113dc918c9060200161477c565b6040516020818303038152906040528051906020012090506113fc613ade565b604080518082019091526000808252602082015260005b888110156116975760006001858884602081106114325761143261461a565b61143f91901a601b6146f6565b8d8d868181106114515761145161461a565b905060200201358c8c8781811061146a5761146a61461a565b90506020020135604051600081526020016040526040516114a7949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156114c9573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600d602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561154957611549613b7a565b600281111561155a5761155a613b7a565b905250925060018360200151600281111561157757611577613b7a565b146115de576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610dd1565b8251849060ff16601f81106115f5576115f561461a565b602002015115611661576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610dd1565b600184846000015160ff16601f811061167c5761167c61461a565b911515602090920201525061169081614678565b9050611413565b5050505063ffffffff81106116ae576116ae614798565b505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8082166000908152600360205260408120549091168061171a576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611796573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117ba91906147d2565b9392505050565b60005460ff161561182e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610dd1565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561189b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118bf91906145fd565b156118f5576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60065473ffffffffffffffffffffffffffffffffffffffff16611944576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060820151516000819003611985576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff8111156119a0576119a0613d75565b6040519080825280602002602001820160405280156119c9578160200160208202803683370190505b50905060008267ffffffffffffffff8111156119e7576119e7613d75565b604051908082528060200260200182016040528015611aba57816020015b611aa760405180610140016040528060008152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b815260200190600190039081611a055790505b50905060005b83811015611b875785606001518181518110611ade57611ade61461a565b6020026020010151806020019051810190611af99190614906565b828281518110611b0b57611b0b61461a565b602002602001018190525085606001518181518110611b2c57611b2c61461a565b6020026020010151604051602001611b449190614a22565b60405160208183030381529060405280519060200120838281518110611b6c57611b6c61461a565b6020908102919091010152611b8081614678565b9050611ac0565b50600080611ba88488608001518960a001518a60c001518b60e0015161335e565b915091506000835182611bbb9190614a48565b9050868015611bf657506008547801000000000000000000000000000000000000000000000000900463ffffffff16611bf484426146b0565b105b15611c2d576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b868110156116ae576000858281518110611c4c57611c4c61461a565b602002602001015190506000611c7f826020015167ffffffffffffffff1660009081526009602052604090205460ff1690565b90506002816003811115611c9557611c95613b7a565b03611cde5760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610dd1565b611ce782613463565b60005b8260a0015151811015611d2e57611d1d8360a001518281518110611d1057611d1061461a565b60200260200101516135c8565b50611d2781614678565b9050611cea565b506000816003811115611d4357611d43613b7a565b148015611d4e575089155b15611f6d57600080611d638460e001516135c8565b73ffffffffffffffffffffffffffffffffffffffff166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611dad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dd191906147d2565b905060005b8d6020015151811015611e64578173ffffffffffffffffffffffffffffffffffffffff168e602001518281518110611e1057611e1061461a565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1603611e54578d604001518181518110611e4957611e4961461a565b602002602001015192505b611e5d81614678565b9050611dd6565b5081611eb4576040517fce480bcc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610dd1565b6000670de0b6b3a7640000833a8761012001518a611ed29190614a5c565b611edc9190614a74565b611ee69190614a74565b611ef09190614a48565b9050846101000151811115611f555760208501516101008601516040517f394a2c2700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9092166004830152602482018390526044820152606401610dd1565b611f698560e0015186610100015130613644565b5050505b60208281015167ffffffffffffffff16600090815260099091526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611fc5611fc0846136e1565b613806565b60208085015167ffffffffffffffff166000908152600990915260409020805491925082917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600183600381111561202057612020613b7a565b0217905550826020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf826040516120639190613ba9565b60405180910390a25050508061207890614678565b9050611c30565b612087612e77565b600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b855185518560ff16601f831115612169576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610dd1565b600081116121d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610dd1565b818314612261576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610dd1565b61226c816003614a74565b83116122d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610dd1565b6122dc612e77565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600e54156124cf57600e54600090612334906001906146b0565b90506000600e828154811061234b5761234b61461a565b6000918252602082200154600f805473ffffffffffffffffffffffffffffffffffffffff909216935090849081106123855761238561461a565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600d909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600e80549192509080612405576124056146c7565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600f80548061246e5761246e6146c7565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690550190555061231a915050565b60005b815151811015612934576000600d6000846000015184815181106124f8576124f861461a565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561254257612542613b7a565b146125a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610dd1565b6040805180820190915260ff821681526001602082015282518051600d91600091859081106125da576125da61461a565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561267b5761267b613b7a565b02179055506000915061268b9050565b600d6000846020015184815181106126a5576126a561461a565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff1660028111156126ef576126ef613b7a565b14612756576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610dd1565b6040805180820190915260ff82168152602081016002815250600d6000846020015184815181106127895761278961461a565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561282a5761282a613b7a565b02179055505082518051600e9250839081106128485761284861461a565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600f9190839081106128c4576128c461461a565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691909117905561292d81614678565b90506124d2565b506040810151600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600c80547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff4381168202928317855590830481169360019390926000926129c6928692908216911617614ab1565b92506101000a81548163ffffffff021916908363ffffffff160217905550612a254630600c60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613938565b600a81905582518051600b805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055600c5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612ac4988b98919763ffffffff909216969095919491939192614ad9565b60405180910390a15050505050505050505050565b612ae1612e77565b73ffffffffffffffffffffffffffffffffffffffff82161580612b18575073ffffffffffffffffffffffffffffffffffffffff8116155b15612b4f576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612beb576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600581529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b612d29612e77565b610fc0816139e3565b612d3a612e77565b80516008805460208085018051604080880180516060808b01805163ffffffff9081167801000000000000000000000000000000000000000000000000027fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff95861670010000000000000000000000000000000002167fffffffff000000000000000000000000ffffffffffffffffffffffffffffffff98861668010000000000000000027fffffffffffffffffffffffffffffffff00000000000000000000000000000000909c169d86169d8e179b909b17979097169990991795909517909855825198895293518416948801949094529251909116918501919091525116908201527f187b05a44b331bbf26fcef5672820acbcd72e33ec0214a0b89042db7b73b46a29060800160405180910390a150565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610874576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610dd1565b60005460ff16612f69576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610dd1565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60005460ff161561304b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610dd1565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612fb43390565b81518351146130d9576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015613132576131228482815181106130fa576130fa61461a565b60200260200101518483815181106131145761311461461a565b602002602001015184613644565b61312b81614678565b90506130dc565b50505050565b606081015173ffffffffffffffffffffffffffffffffffffffff163b61315b5750565b6006546040517fc39a285b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063c39a285b906131b1908490600401614b9f565b6020604051808303816000875af11580156131d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131f491906145fd565b610fc0576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff1663be9b03f18280602001905181019061325a9190614de2565b60006040518363ffffffff1660e01b8152600401613279929190614faa565b600060405180830381600087803b15801561329357600080fd5b505af11580156132a7573d6000803e3d6000fd5b50505050505050565b60006132bd826020614a74565b6132c8856020614a74565b6132d488610144614a5c565b6132de9190614a5c565b6132e89190614a5c565b6132f3906000614a5c565b90503681146132a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610dd1565b60008060005a6007546040517fe71e65ce00000000000000000000000000000000000000000000000000000000815291925060009173ffffffffffffffffffffffffffffffffffffffff9091169063e71e65ce906133c8908c908c908c908c908c906004016150a1565b6020604051808303816000875af11580156133e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061340b91906150f3565b905060008111613447576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a61345390846146b0565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146134c35780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610dd1565b60085460a08201515170010000000000000000000000000000000090910467ffffffffffffffff16108061350157508060c00151518160a001515114155b1561354a5760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610dd1565b6008546080820151516801000000000000000090910467ffffffffffffffff161015610fc0576008546080820151516040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090920467ffffffffffffffff1660048301526024820152604401610dd1565b73ffffffffffffffffffffffffffffffffffffffff818116600090815260036020526040902054168061363f576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610dd1565b919050565b600061364f846135c8565b6040517fea6192a200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690529192509082169063ea6192a290604401600060405180830381600087803b1580156136c357600080fd5b505af11580156136d7573d6000803e3d6000fd5b5050505050505050565b61374960405180610100016040528060008152602001600067ffffffffffffffff16815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001600081525090565b60405180610100016040528083600001518152602001836020015167ffffffffffffffff16815260200183604001516040516020016137a4919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b6040516020818303038152906040528152602001836060015173ffffffffffffffffffffffffffffffffffffffff168152602001836080015181526020018360a0015181526020018360c0015181526020018361012001518152509050919050565b6040517f8b721c710000000000000000000000000000000000000000000000000000000081526000903090638b721c7190613845908590600401614b9f565b600060405180830381600087803b15801561385f57600080fd5b505af1925050508015613870575060015b613930573d80801561389e576040519150601f19603f3d011682016040523d82523d6000602084013e6138a3565b606091505b506138ad8161510c565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da800000000000000000000000000000000000000000000000000000000036138fe5750600392915050565b6040517f2532cf4500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506002919050565b6000808a8a8a8a8a8a8a8a8a60405160200161395c9998979695949392919061515c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603613a62576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610dd1565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604051806103e00160405280601f906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff81168114610fc057600080fd5b600060208284031215613b3157600080fd5b81356117ba81613afd565b67ffffffffffffffff81168114610fc057600080fd5b803561363f81613b3c565b600060208284031215613b6f57600080fd5b81356117ba81613b3c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310613be4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b83811015613c05578181015183820152602001613bed565b838111156131325750506000910152565b60008151808452613c2e816020860160208601613bea565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006117ba6020830184613c16565b803561363f81613afd565b6020808252825182820181905260009190848201906040850190845b81811015613ccc57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101613c9a565b50909695505050505050565b60008060408385031215613ceb57600080fd5b8235613cf681613afd565b91506020830135613d0681613afd565b809150509250929050565b600081518084526020808501945080840160005b83811015613d5757815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101613d25565b509495945050505050565b6020815260006117ba6020830184613d11565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff81118282101715613dc857613dc8613d75565b60405290565b604051610140810167ffffffffffffffff81118282101715613dc857613dc8613d75565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613e3957613e39613d75565b604052919050565b600067ffffffffffffffff821115613e5b57613e5b613d75565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112613e9857600080fd5b8135613eab613ea682613e41565b613df2565b818152846020838601011115613ec057600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115613ef757613ef7613d75565b5060051b60200190565b600082601f830112613f1257600080fd5b81356020613f22613ea683613edd565b82815260059290921b84018101918181019086841115613f4157600080fd5b8286015b84811015613f65578035613f5881613afd565b8352918301918301613f45565b509695505050505050565b600082601f830112613f8157600080fd5b81356020613f91613ea683613edd565b82815260059290921b84018101918181019086841115613fb057600080fd5b8286015b84811015613f655780358352918301918301613fb4565b600060208284031215613fdd57600080fd5b813567ffffffffffffffff80821115613ff557600080fd5b90830190610100828603121561400a57600080fd5b614012613da4565b8235815261402260208401613b52565b602082015260408301358281111561403957600080fd5b61404587828601613e87565b60408301525061405760608401613c73565b606082015260808301358281111561406e57600080fd5b61407a87828601613e87565b60808301525060a08301358281111561409257600080fd5b61409e87828601613f01565b60a08301525060c0830135828111156140b657600080fd5b6140c287828601613f70565b60c08301525060e083013560e082015280935050505092915050565b60008083601f8401126140f057600080fd5b50813567ffffffffffffffff81111561410857600080fd5b6020830191508360208260051b850101111561412357600080fd5b9250929050565b60008060008060008060008060e0898b03121561414657600080fd5b606089018a81111561415757600080fd5b8998503567ffffffffffffffff8082111561417157600080fd5b818b0191508b601f83011261418557600080fd5b81358181111561419457600080fd5b8c60208285010111156141a657600080fd5b6020830199508098505060808b01359150808211156141c457600080fd5b6141d08c838d016140de565b909750955060a08b01359150808211156141e957600080fd5b506141f68b828c016140de565b999c989b50969995989497949560c00135949350505050565b600082601f83011261422057600080fd5b81356020614230613ea683613edd565b82815260059290921b8401810191818101908684111561424f57600080fd5b8286015b84811015613f6557803561426681613b3c565b8352918301918301614253565b600082601f83011261428457600080fd5b81356020614294613ea683613edd565b82815260059290921b840181019181810190868411156142b357600080fd5b8286015b84811015613f6557803567ffffffffffffffff8111156142d75760008081fd5b6142e58986838b0101613e87565b8452509183019183016142b7565b8015158114610fc057600080fd5b803561363f816142f3565b6000806040838503121561431f57600080fd5b823567ffffffffffffffff8082111561433757600080fd5b90840190610100828703121561434c57600080fd5b614354613da4565b82358281111561436357600080fd5b61436f8882860161420f565b82525060208301358281111561438457600080fd5b61439088828601613f01565b6020830152506040830135828111156143a857600080fd5b6143b488828601613f70565b6040830152506060830135828111156143cc57600080fd5b6143d888828601614273565b6060830152506080830135828111156143f057600080fd5b6143fc88828601613f70565b60808301525060a083013560a082015260c08301358281111561441e57600080fd5b61442a88828601613f70565b60c08301525060e083013560e082015280945050505061444c60208401614301565b90509250929050565b803560ff8116811461363f57600080fd5b60008060008060008060c0878903121561447f57600080fd5b863567ffffffffffffffff8082111561449757600080fd5b6144a38a838b01613f01565b975060208901359150808211156144b957600080fd5b6144c58a838b01613f01565b96506144d360408a01614455565b955060608901359150808211156144e957600080fd5b6144f58a838b01613e87565b945061450360808a01613b52565b935060a089013591508082111561451957600080fd5b5061452689828a01613e87565b9150509295509295509295565b60006020828403121561454557600080fd5b813567ffffffffffffffff81111561455c57600080fd5b820161010081850312156117ba57600080fd5b60006080828403121561458157600080fd5b6040516080810181811067ffffffffffffffff821117156145a4576145a4613d75565b60405282356145b281613b3c565b815260208301356145c281613b3c565b602082015260408301356145d581613b3c565b6040820152606083013563ffffffff811681146145f157600080fd5b60608201529392505050565b60006020828403121561460f57600080fd5b81516117ba816142f3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036146a9576146a9614649565b5060010190565b6000828210156146c2576146c2614649565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff0382111561471357614713614649565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff83168061475d5761475d61471b565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b805161363f81613afd565b6000602082840312156147e457600080fd5b81516117ba81613afd565b805161363f81613b3c565b600082601f83011261480b57600080fd5b8151614819613ea682613e41565b81815284602083860101111561482e57600080fd5b61483f826020830160208701613bea565b949350505050565b600082601f83011261485857600080fd5b81516020614868613ea683613edd565b82815260059290921b8401810191818101908684111561488757600080fd5b8286015b84811015613f6557805161489e81613afd565b835291830191830161488b565b600082601f8301126148bc57600080fd5b815160206148cc613ea683613edd565b82815260059290921b840181019181810190868411156148eb57600080fd5b8286015b84811015613f6557805183529183019183016148ef565b60006020828403121561491857600080fd5b815167ffffffffffffffff8082111561493057600080fd5b90830190610140828603121561494557600080fd5b61494d613dce565b8251815261495d602084016147ef565b602082015261496e604084016147c7565b604082015261497f606084016147c7565b606082015260808301518281111561499657600080fd5b6149a2878286016147fa565b60808301525060a0830151828111156149ba57600080fd5b6149c687828601614847565b60a08301525060c0830151828111156149de57600080fd5b6149ea878286016148ab565b60c0830152506149fc60e084016147c7565b60e082015261010083810151908201526101209283015192810192909252509392505050565b6000815260008251614a3b816001850160208701613bea565b9190910160010192915050565b600082614a5757614a5761471b565b500490565b60008219821115614a6f57614a6f614649565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614aac57614aac614649565b500290565b600063ffffffff808316818516808303821115614ad057614ad0614649565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614b098184018a613d11565b90508281036080840152614b1d8189613d11565b905060ff871660a084015282810360c0840152614b3a8187613c16565b905067ffffffffffffffff851660e0840152828103610100840152614b5f8185613c16565b9c9b505050505050505050505050565b600081518084526020808501945080840160005b83811015613d5757815187529582019590820190600101614b83565b602081528151602082015260006020830151614bc7604084018267ffffffffffffffff169052565b506040830151610100806060850152614be4610120850183613c16565b91506060850151614c0d608086018273ffffffffffffffffffffffffffffffffffffffff169052565b5060808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808685030160a0870152614c488483613c16565b935060a08701519150808685030160c0870152614c658483613d11565b935060c08701519150808685030160e087015250614c838382614b6f565b92505060e085015181850152508091505092915050565b600082601f830112614cab57600080fd5b81516020614cbb613ea683613edd565b82815260059290921b84018101918181019086841115614cda57600080fd5b8286015b84811015613f65578051614cf181613b3c565b8352918301918301614cde565b600082601f830112614d0f57600080fd5b81516020614d1f613ea683613edd565b82815260059290921b84018101918181019086841115614d3e57600080fd5b8286015b84811015613f65578051614d5581613afd565b8352918301918301614d42565b600082601f830112614d7357600080fd5b81516020614d83613ea683613edd565b82815260059290921b84018101918181019086841115614da257600080fd5b8286015b84811015613f6557805167ffffffffffffffff811115614dc65760008081fd5b614dd48986838b01016147fa565b845250918301918301614da6565b600060208284031215614df457600080fd5b815167ffffffffffffffff80821115614e0c57600080fd5b908301906101008286031215614e2157600080fd5b614e29613da4565b825182811115614e3857600080fd5b614e4487828601614c9a565b825250602083015182811115614e5957600080fd5b614e6587828601614cfe565b602083015250604083015182811115614e7d57600080fd5b614e89878286016148ab565b604083015250606083015182811115614ea157600080fd5b614ead87828601614d62565b606083015250608083015182811115614ec557600080fd5b614ed1878286016148ab565b60808301525060a083015160a082015260c083015182811115614ef357600080fd5b614eff878286016148ab565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b83811015613d5757815167ffffffffffffffff1687529582019590820190600101614f2f565b600081518084526020808501808196508360051b8101915082860160005b85811015614f9d578284038952614f8b848351613c16565b98850198935090840190600101614f73565b5091979650505050505050565b6040815260008351610100806040850152614fc9610140850183614f1b565b915060208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808685030160608701526150058483613d11565b935060408801519150808685030160808701526150228483614b6f565b935060608801519150808685030160a087015261503f8483614f55565b935060808801519150808685030160c087015261505c8483614b6f565b935060a088015160e087015260c08801519150808685030183870152506150838382614b6f565b60e0880151610120870152861515602087015293506117ba92505050565b60a0815260006150b460a0830188614b6f565b82810360208401526150c68188614b6f565b905085604084015282810360608401526150e08186614b6f565b9150508260808301529695505050505050565b60006020828403121561510557600080fd5b5051919050565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156151545780818460040360031b1b83161693505b505050919050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526151a38285018b613d11565b915083820360808501526151b7828a613d11565b915060ff881660a085015283820360c08501526151d48288613c16565b90861660e08501528381036101008501529050614b5f8185613c1656fea164736f6c634300080f000a",
}

var EVM2EVMTollOffRampABI = EVM2EVMTollOffRampMetaData.ABI

var EVM2EVMTollOffRampBin = EVM2EVMTollOffRampMetaData.Bin

func DeployEVM2EVMTollOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, blobVerifier common.Address, onRampAddress common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address) (common.Address, *types.Transaction, *EVM2EVMTollOffRamp, error) {
	parsed, err := EVM2EVMTollOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMTollOffRampBin), backend, sourceChainId, chainId, offRampConfig, blobVerifier, onRampAddress, afn, sourceTokens, pools)
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessage) error {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) CcipReceive(arg0 CCIPAny2EVMMessage) error {
	return _EVM2EVMTollOffRamp.Contract.CcipReceive(&_EVM2EVMTollOffRamp.CallOpts, arg0)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) CcipReceive(arg0 CCIPAny2EVMMessage) error {
	return _EVM2EVMTollOffRamp.Contract.CcipReceive(&_EVM2EVMTollOffRamp.CallOpts, arg0)
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) IChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "i_chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) IChainId() (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.IChainId(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) IChainId() (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.IChainId(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) ISourceChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "i_sourceChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) ISourceChainId() (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.ISourceChainId(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) ISourceChainId() (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.ISourceChainId(&_EVM2EVMTollOffRamp.CallOpts)
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) IsPool(addr common.Address) (bool, error) {
	return _EVM2EVMTollOffRamp.Contract.IsPool(&_EVM2EVMTollOffRamp.CallOpts, addr)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) IsPool(addr common.Address) (bool, error) {
	return _EVM2EVMTollOffRamp.Contract.IsPool(&_EVM2EVMTollOffRamp.CallOpts, addr)
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) SRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "s_router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SRouter() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.SRouter(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) SRouter() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.SRouter(&_EVM2EVMTollOffRamp.CallOpts)
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessage) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "executeSingleMessage", message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) ExecuteSingleMessage(message CCIPAny2EVMMessage) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRamp.TransactOpts, message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMMessage) (*types.Transaction, error) {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig(&_EVM2EVMTollOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig(&_EVM2EVMTollOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetConfig0(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setConfig0", config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetConfig0(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig0(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetConfig0(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig0(&_EVM2EVMTollOffRamp.TransactOpts, config)
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

func (EVM2EVMTollOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMTollOffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf")
}

func (EVM2EVMTollOffRampOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x187b05a44b331bbf26fcef5672820acbcd72e33ec0214a0b89042db7b73b46a2")
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

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMTollOffRampAFNSet, error)

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

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMTollOffRampTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMTollOffRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
