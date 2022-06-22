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

type CCIPExecutionResult struct {
	SequenceNumber   uint64
	GasUsed          *big.Int
	TimestampRelayed *big.Int
	State            uint8
}

type TollOffRampInterfaceOffRampConfig struct {
	SourceChainId         *big.Int
	ExecutionDelaySeconds uint64
	MaxDataSize           uint64
	MaxTokensLength       uint64
}

var Any2EVMTollOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"needFee\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampRelayed\",\"type\":\"uint256\"},{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structCCIP.ExecutionResult[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"executedMessages\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_router\",\"outputs\":[{\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005694380380620056948339810160408190526200003491620005e6565b6000805460ff1916815560019084908490879085903390819081620000a05760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000da57620000da81620002ec565b5050506001600160a01b0382161580620000f2575080155b156200011157604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001575760405162d8548360e71b815260040160405180910390fd5b81516200016c9060059060208501906200039d565b5060005b82518110156200025057600082828151811062000191576200019162000703565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001db57620001db62000703565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff1916600117905580620002478162000719565b91505062000170565b50505015156080525050845160a05250505060c0929092528051600f5560208101516010805460408401516060909401516001600160401b03908116600160801b02600160801b600160c01b031995821668010000000000000000026001600160801b031990931691909416171792909216179055600e80546001600160a01b039092166001600160a01b031990921691909117905562000741565b336001600160a01b03821603620003465760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000097565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620003f5579160200282015b82811115620003f557825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003be565b506200040392915062000407565b5090565b5b8082111562000403576000815560010162000408565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b03811182821017156200045957620004596200041e565b60405290565b604051601f8201601f191681016001600160401b03811182821017156200048a576200048a6200041e565b604052919050565b80516001600160401b0381168114620004aa57600080fd5b919050565b6001600160a01b0381168114620004c557600080fd5b50565b8051620004aa81620004af565b60006001600160401b03821115620004f157620004f16200041e565b5060051b60200190565b600082601f8301126200050d57600080fd5b81516020620005266200052083620004d5565b6200045f565b82815260059290921b840181019181810190868411156200054657600080fd5b8286015b848110156200056e5780516200056081620004af565b83529183019183016200054a565b509695505050505050565b600082601f8301126200058b57600080fd5b815160206200059e6200052083620004d5565b82815260059290921b84018101918181019086841115620005be57600080fd5b8286015b848110156200056e578051620005d881620004af565b8352918301918301620005c2565b600080600080600080600080888a036101608112156200060557600080fd5b895198506080601f19820112156200061c57600080fd5b506200062762000434565b60208a015181526200063c60408b0162000492565b60208201526200064f60608b0162000492565b60408201526200066260808b0162000492565b606082015296506200067760a08a01620004c8565b95506200068760c08a01620004c8565b94506200069760e08a01620004c8565b6101008a01519094506001600160401b0380821115620006b657600080fd5b620006c48c838d01620004fb565b94506101208b0151915080821115620006dc57600080fd5b50620006eb8b828c0162000579565b92505061014089015190509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b6000600182016200073a57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c051614f1c6200077860003960006103cf015260008181610345015261321501526000610f1d0152614f1c6000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c806385e1f4d011610104578063b6608c3b116100a2578063e16e632c11610071578063e16e632c146104e4578063e3d0e71214610504578063eb511dd414610517578063f2fde38b1461052a57600080fd5b8063b6608c3b14610465578063bbe4f6db14610478578063be9b03f1146104b1578063c0d78655146104d157600080fd5b80638da5cb5b116100de5780638da5cb5b14610407578063afcb95d71461042a578063b034909c1461044a578063b1dc65a41461045257600080fd5b806385e1f4d0146103ca57806389c06568146103f15780638bbad066146103f957600080fd5b80635c975abb1161017c57806379ba50971161014b57806379ba509714610375578063814118341461037d57806381ff7048146103925780638456cb59146103c257600080fd5b80635c975abb146102f25780636edcbf38146102fd578063744b92e21461032d57806374be21501461034057600080fd5b80632222dd42116101b85780632222dd421461024f5780633f4ba83a1461028e578063567c814b146102965780635b16ebb7146102b957600080fd5b8063092cddc2146101df578063108ee5fc146101f4578063181f5a7714610207575b600080fd5b6101f26101ed3660046139ca565b61053d565b005b6101f2610202366004613ae6565b610599565b604080518082018252601881527f416e793245564d546f6c6c4f666652616d7020312e302e300000000000000000602082015290516102469190613b80565b60405180910390f35b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610246565b6101f2610675565b6102a96102a4366004613b93565b610687565b6040519015158152602001610246565b6102a96102c7366004613ae6565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff166102a9565b61032061030b366004613bac565b60116020526000908152604090205460ff1681565b6040516102469190613c33565b6101f261033b366004613c41565b6107ce565b6103677f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610246565b6101f2610bcd565b610385610cf4565b6040516102469190613ccb565b6009546007546040805163ffffffff80851682526401000000009094049093166020840152820152606001610246565b6101f2610d63565b6103677f000000000000000000000000000000000000000000000000000000000000000081565b610385610d73565b6101f26101da366004613cde565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610269565b604080516001815260006020820181905291810191909152606001610246565b600354610367565b6101f2610460366004613d66565b610de0565b6101f2610473366004613b93565b611489565b610269610486366004613ae6565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6104c46104bf366004613fac565b611509565b60405161024691906140f5565b6101f26104df366004613ae6565b611d7a565b600d546102699073ffffffffffffffffffffffffffffffffffffffff1681565b6101f2610512366004614180565b611dfb565b6101f2610525366004613c41565b6127e0565b6101f2610538366004613ae6565b612a28565b333014610576576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61058d8160a001518260c001518360600151612a39565b61059681612a9a565b50565b6105a1612b9d565b73ffffffffffffffffffffffffffffffffffffffff81166105ee576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b61067d612b9d565b610685612c23565b565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa1580156106f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071b919061424d565b1580156107c85750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610793573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107b7919061426a565b602001516107c590846142f5565b11155b92915050565b6107d6612b9d565b6005546000819003610814576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906108af576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610918576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060056109276001856142f5565b815481106109375761093761430c565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff16815481106109895761098961430c565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1660056109b86001866142f5565b815481106109c8576109c861430c565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110610a3657610a3661430c565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480610ad857610ad861433b565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610c53576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600c805480602002602001604051908101604052809291908181526020018280548015610d5957602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610d2e575b5050505050905090565b610d6b612b9d565b610685612d04565b60606005805480602002602001604051908101604052809291908181526020018280548015610d595760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610d2e575050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c013591610e3691849163ffffffff851691908e908e9081908401838280828437600092019190915250612dc492505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260075480825260085460ff80821660208501526101009091041692820192909252908314610f0b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610c4a565b610f198b8b8b8b8b8b613055565b60007f000000000000000000000000000000000000000000000000000000000000000015610f7657600282602001518360400151610f57919061436a565b610f6191906143be565b610f6c90600161436a565b60ff169050610f8c565b6020820151610f8690600161436a565b60ff1690505b888114610ff5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610c4a565b88871461105e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610c4a565b336000908152600a602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156110a1576110a1613bc9565b60028111156110b2576110b2613bc9565b90525090506002816020015160028111156110cf576110cf613bc9565b1480156111165750600c816000015160ff16815481106110f1576110f161430c565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b61117c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610c4a565b5050505050600088886040516111939291906143e0565b6040519081900381206111aa918c906020016143f0565b6040516020818303038152906040528051906020012090506111ca6136e4565b604080518082019091526000808252602082015260005b888110156114675760006001858884602081106112005761120061430c565b61120d91901a601b61436a565b8d8d8681811061121f5761121f61430c565b905060200201358c8c878181106112385761123861430c565b9050602002013560405160008152602001604052604051611275949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611297573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600a602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561131757611317613bc9565b600281111561132857611328613bc9565b905250925060018360200151600281111561134557611345613bc9565b146113ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610c4a565b8251849060ff16601f81106113c3576113c361430c565b60200201511561142f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610c4a565b600184846000015160ff16601f811061144a5761144a61430c565b91151560209092020152508061145f8161440c565b9150506111e1565b5050505063ffffffff811061147e5761147e614444565b505050505050505050565b611491612b9d565b806000036114cb576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c2519101610669565b606061151760005460ff1690565b1561157e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c4a565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061160f919061424d565b15611645576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa1580156116b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116d9919061426a565b90506003548160200151426116ee91906142f5565b1115611726576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5473ffffffffffffffffffffffffffffffffffffffff16611775576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608401515160008190036117b6576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff8111156117d1576117d1613703565b6040519080825280602002602001820160405280156117fa578160200160208202803683370190505b50905060008267ffffffffffffffff81111561181857611818613703565b6040519080825280602002602001820160405280156118eb57816020015b6118d860405180610140016040528060008152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b8152602001906001900390816118365790505b50905060005b838110156119f7578760600151818151811061190f5761190f61430c565b602002602001015180602001905181019061192a9190614595565b82828151811061193c5761193c61430c565b6020026020010181905250600082828151811061195b5761195b61430c565b602002602001015160405160200161197391906147ea565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526119ae916020016147fd565b604051602081830303815290604052905080805190602001208483815181106119d9576119d961430c565b602090810291909101015250806119ef8161440c565b9150506118f1565b50600080611a18848a608001518b60a001518c60c001518d60e0015161310c565b915091506000835182611a2b9190614823565b905060008667ffffffffffffffff811115611a4857611a48613703565b604051908082528060200260200182016040528015611aa157816020015b611a8e6040805160808101825260008082526020820181905291810182905290606082015290565b815260200190600190039081611a665790505b50905060005b87811015611d6b5760005a90506000878381518110611ac857611ac861430c565b602002602001015190506000611afb826020015167ffffffffffffffff1660009081526011602052604090205460ff1690565b90506002816003811115611b1157611b11613bc9565b03611b5a5760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610c4a565b611b6382613211565b60005b8260a0015151811015611bac57611b998360a001518281518110611b8c57611b8c61430c565b6020026020010151613376565b5080611ba48161440c565b915050611b66565b506003816003811115611bc157611bc1613bc9565b14158015611bcc57508d5b15611be557611be58260e00151836101000151306133f2565b60208281015167ffffffffffffffff16600090815260119091526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611c358361348f565b60208085015167ffffffffffffffff166000908152601190915260409020805491925082917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001836003811115611c9057611c90613bc9565b02179055506000875a611ca390876142f5565b611cad9190614837565b905060006040518060800160405280866020015167ffffffffffffffff1681526020018381526020018c8152602001846003811115611cee57611cee613bc9565b815250905080888881518110611d0657611d0661430c565b60200260200101819052507fbca6416e78a437ab47530846568a4d78457e41bc2adc0d91a826090e2d853d1c81600001518260600151604051611d4a92919061484f565b60405180910390a15050505050508080611d639061440c565b915050611aa7565b509a9950505050505050505050565b611d82612b9d565b600d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d49060200160405180910390a150565b855185518560ff16601f831115611e6e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610c4a565b60008111611ed8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610c4a565b818314611f66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610c4a565b611f7181600361486d565b8311611fd9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610c4a565b611fe1612b9d565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600b54156121d457600b54600090612039906001906142f5565b90506000600b82815481106120505761205061430c565b6000918252602082200154600c805473ffffffffffffffffffffffffffffffffffffffff9092169350908490811061208a5761208a61430c565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600a909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600b8054919250908061210a5761210a61433b565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600c8054806121735761217361433b565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690550190555061201f915050565b60005b81515181101561263b576000600a6000846000015184815181106121fd576121fd61430c565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561224757612247613bc9565b146122ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610c4a565b6040805180820190915260ff821681526001602082015282518051600a91600091859081106122df576122df61430c565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561238057612380613bc9565b0217905550600091506123909050565b600a6000846020015184815181106123aa576123aa61430c565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff1660028111156123f4576123f4613bc9565b1461245b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610c4a565b6040805180820190915260ff82168152602081016002815250600a60008460200151848151811061248e5761248e61430c565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561252f5761252f613bc9565b02179055505082518051600b92508390811061254d5761254d61430c565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600c9190839081106125c9576125c961430c565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055806126338161440c565b9150506121d7565b506040810151600880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600980547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff4381168202928317855590830481169360019390926000926126cd9286929082169116176148aa565b92506101000a81548163ffffffff021916908363ffffffff16021790555061272c4630600960009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a0015161353e565b6007819055825180516008805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560095460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05986127cb988b98919763ffffffff9092169690959194919391926148d2565b60405180910390a15050505050505050505050565b6127e8612b9d565b73ffffffffffffffffffffffffffffffffffffffff8216158061281f575073ffffffffffffffffffffffffffffffffffffffff8116155b15612856576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156128f2576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b612a30612b9d565b610596816135e9565b60005b8351811015612a9457612a82848281518110612a5a57612a5a61430c565b6020026020010151848381518110612a7457612a7461430c565b6020026020010151846133f2565b80612a8c8161440c565b915050612a3c565b50505050565b606081015173ffffffffffffffffffffffffffffffffffffffff163b612b0a5760608101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610c4a565b6060810151600d546040517ffd12f6e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063fd12f6e590612b679084908690600401614968565b600060405180830381600087803b158015612b8157600080fd5b505af1158015612b95573d6000803e3d6000fd5b505050505050565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610685576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610c4a565b60005460ff16612c8f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610c4a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60005460ff1615612d71576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c4a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612cda3390565b600081806020019051810190612dda9190614a7b565b6040517fbe9b03f1000000000000000000000000000000000000000000000000000000008152909150600090309063be9b03f190612e1f908590600190600401614c36565b6000604051808303816000875af1158015612e3e573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052612e849190810190614d2d565b905060005b826060015151811015612b9557600083606001518281518110612eae57612eae61430c565b6020026020010151806020019051810190612ec99190614595565b90506000612eda8260e00151613376565b90506000805b866020015151811015612f74578360e0015173ffffffffffffffffffffffffffffffffffffffff1687602001518281518110612f1e57612f1e61430c565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1603612f625786604001518181518110612f5757612f5761430c565b602002602001015191505b80612f6c8161440c565b915050612ee0565b5080612fca5760e08301516040517fce480bcc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610c4a565b6000670de0b6b3a7640000823a888881518110612fe957612fe961430c565b602002602001015160200151612fff919061486d565b613009919061486d565b6130139190614823565b905060008185610100015161302891906142f5565b905061303d8560e001518287606001516133f2565b5050505050808061304d9061440c565b915050612e89565b600061306282602061486d565b61306d85602061486d565b61307988610144614837565b6130839190614837565b61308d9190614837565b613098906000614837565b9050368114613103576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610c4a565b50505050505050565b60008060005a600e546040517fe71e65ce00000000000000000000000000000000000000000000000000000000815291925060009173ffffffffffffffffffffffffffffffffffffffff9091169063e71e65ce90613176908c908c908c908c908c90600401614e0f565b6020604051808303816000875af1158015613195573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131b99190614e61565b9050600081116131f5576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a61320190846142f5565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146132715780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610c4a565b60105460a08201515170010000000000000000000000000000000090910467ffffffffffffffff1610806132af57508060c00151518160a001515114155b156132f85760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610c4a565b6010546080820151516801000000000000000090910467ffffffffffffffff161015610596576010546080820151516040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090920467ffffffffffffffff1660048301526024820152604401610c4a565b73ffffffffffffffffffffffffffffffffffffffff81811660009081526004602052604090205416806133ed576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610c4a565b919050565b60006133fd84613376565b6040517fea6192a200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690529192509082169063ea6192a290604401600060405180830381600087803b15801561347157600080fd5b505af1158015613485573d6000803e3d6000fd5b5050505050505050565b6040517f092cddc2000000000000000000000000000000000000000000000000000000008152600090309063092cddc2906134ce9085906004016147ea565b600060405180830381600087803b1580156134e857600080fd5b505af19250505080156134f9575060015b613536573d808015613527576040519150601f19603f3d011682016040523d82523d6000602084013e61352c565b606091505b5060039392505050565b506002919050565b6000808a8a8a8a8a8a8a8a8a60405160200161356299989796959493929190614e7a565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603613668576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610c4a565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604051806103e00160405280601f906020820280368337509192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610140810167ffffffffffffffff8111828210171561375657613756613703565b60405290565b604051610100810167ffffffffffffffff8111828210171561375657613756613703565b6040516080810167ffffffffffffffff8111828210171561375657613756613703565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156137ea576137ea613703565b604052919050565b67ffffffffffffffff8116811461059657600080fd5b80356133ed816137f2565b73ffffffffffffffffffffffffffffffffffffffff8116811461059657600080fd5b80356133ed81613813565b600067ffffffffffffffff82111561385a5761385a613703565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261389757600080fd5b81356138aa6138a582613840565b6137a3565b8181528460208386010111156138bf57600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff8211156138f6576138f6613703565b5060051b60200190565b600082601f83011261391157600080fd5b813560206139216138a5836138dc565b82815260059290921b8401810191818101908684111561394057600080fd5b8286015b8481101561396457803561395781613813565b8352918301918301613944565b509695505050505050565b600082601f83011261398057600080fd5b813560206139906138a5836138dc565b82815260059290921b840181019181810190868411156139af57600080fd5b8286015b8481101561396457803583529183019183016139b3565b6000602082840312156139dc57600080fd5b813567ffffffffffffffff808211156139f457600080fd5b908301906101408286031215613a0957600080fd5b613a11613732565b82358152613a2160208401613808565b6020820152613a3260408401613835565b6040820152613a4360608401613835565b6060820152608083013582811115613a5a57600080fd5b613a6687828601613886565b60808301525060a083013582811115613a7e57600080fd5b613a8a87828601613900565b60a08301525060c083013582811115613aa257600080fd5b613aae8782860161396f565b60c083015250613ac060e08401613835565b60e082015261010083810135908201526101209283013592810192909252509392505050565b600060208284031215613af857600080fd5b8135613b0381613813565b9392505050565b60005b83811015613b25578181015183820152602001613b0d565b83811115612a945750506000910152565b60008151808452613b4e816020860160208601613b0a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613b036020830184613b36565b600060208284031215613ba557600080fd5b5035919050565b600060208284031215613bbe57600080fd5b8135613b03816137f2565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110613c2f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b602081016107c88284613bf8565b60008060408385031215613c5457600080fd5b8235613c5f81613813565b91506020830135613c6f81613813565b809150509250929050565b600081518084526020808501945080840160005b83811015613cc057815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101613c8e565b509495945050505050565b602081526000613b036020830184613c7a565b600060208284031215613cf057600080fd5b813567ffffffffffffffff811115613d0757600080fd5b82016101408185031215613b0357600080fd5b60008083601f840112613d2c57600080fd5b50813567ffffffffffffffff811115613d4457600080fd5b6020830191508360208260051b8501011115613d5f57600080fd5b9250929050565b60008060008060008060008060e0898b031215613d8257600080fd5b606089018a811115613d9357600080fd5b8998503567ffffffffffffffff80821115613dad57600080fd5b818b0191508b601f830112613dc157600080fd5b813581811115613dd057600080fd5b8c6020828501011115613de257600080fd5b6020830199508098505060808b0135915080821115613e0057600080fd5b613e0c8c838d01613d1a565b909750955060a08b0135915080821115613e2557600080fd5b50613e328b828c01613d1a565b999c989b50969995989497949560c00135949350505050565b600082601f830112613e5c57600080fd5b81356020613e6c6138a5836138dc565b82815260059290921b84018101918181019086841115613e8b57600080fd5b8286015b84811015613964578035613ea2816137f2565b8352918301918301613e8f565b600082601f830112613ec057600080fd5b81356020613ed06138a5836138dc565b82815260059290921b84018101918181019086841115613eef57600080fd5b8286015b84811015613964578035613f0681613813565b8352918301918301613ef3565b600082601f830112613f2457600080fd5b81356020613f346138a5836138dc565b82815260059290921b84018101918181019086841115613f5357600080fd5b8286015b8481101561396457803567ffffffffffffffff811115613f775760008081fd5b613f858986838b0101613886565b845250918301918301613f57565b801515811461059657600080fd5b80356133ed81613f93565b60008060408385031215613fbf57600080fd5b823567ffffffffffffffff80821115613fd757600080fd5b908401906101008287031215613fec57600080fd5b613ff461375c565b82358281111561400357600080fd5b61400f88828601613e4b565b82525060208301358281111561402457600080fd5b61403088828601613eaf565b60208301525060408301358281111561404857600080fd5b6140548882860161396f565b60408301525060608301358281111561406c57600080fd5b61407888828601613f13565b60608301525060808301358281111561409057600080fd5b61409c8882860161396f565b60808301525060a083013560a082015260c0830135828111156140be57600080fd5b6140ca8882860161396f565b60c08301525060e083013560e08201528094505050506140ec60208401613fa1565b90509250929050565b602080825282518282018190526000919060409081850190868401855b82811015614162578151805167ffffffffffffffff168552868101518786015285810151868601526060908101519061414d81870183613bf8565b50506080939093019290850190600101614112565b5091979650505050505050565b803560ff811681146133ed57600080fd5b60008060008060008060c0878903121561419957600080fd5b863567ffffffffffffffff808211156141b157600080fd5b6141bd8a838b01613eaf565b975060208901359150808211156141d357600080fd5b6141df8a838b01613eaf565b96506141ed60408a0161416f565b9550606089013591508082111561420357600080fd5b61420f8a838b01613886565b945061421d60808a01613808565b935060a089013591508082111561423357600080fd5b5061424089828a01613886565b9150509295509295509295565b60006020828403121561425f57600080fd5b8151613b0381613f93565b60006060828403121561427c57600080fd5b6040516060810181811067ffffffffffffffff8211171561429f5761429f613703565b80604052508251815260208301516020820152604083015160408201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015614307576143076142c6565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff03821115614387576143876142c6565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff8316806143d1576143d161438f565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361443d5761443d6142c6565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b80516133ed816137f2565b80516133ed81613813565b600082601f83011261449a57600080fd5b81516144a86138a582613840565b8181528460208386010111156144bd57600080fd5b6144ce826020830160208701613b0a565b949350505050565b600082601f8301126144e757600080fd5b815160206144f76138a5836138dc565b82815260059290921b8401810191818101908684111561451657600080fd5b8286015b8481101561396457805161452d81613813565b835291830191830161451a565b600082601f83011261454b57600080fd5b8151602061455b6138a5836138dc565b82815260059290921b8401810191818101908684111561457a57600080fd5b8286015b84811015613964578051835291830191830161457e565b6000602082840312156145a757600080fd5b815167ffffffffffffffff808211156145bf57600080fd5b9083019061014082860312156145d457600080fd5b6145dc613732565b825181526145ec60208401614473565b60208201526145fd6040840161447e565b604082015261460e6060840161447e565b606082015260808301518281111561462557600080fd5b61463187828601614489565b60808301525060a08301518281111561464957600080fd5b614655878286016144d6565b60a08301525060c08301518281111561466d57600080fd5b6146798782860161453a565b60c08301525061468b60e0840161447e565b60e082015261010083810151908201526101209283015192810192909252509392505050565b600081518084526020808501945080840160005b83811015613cc0578151875295820195908201906001016146c5565b6000610140825184526020830151614705602086018267ffffffffffffffff169052565b50604083015161472d604086018273ffffffffffffffffffffffffffffffffffffffff169052565b506060830151614755606086018273ffffffffffffffffffffffffffffffffffffffff169052565b50608083015181608086015261476d82860182613b36565b91505060a083015184820360a08601526147878282613c7a565b91505060c083015184820360c08601526147a182826146b1565b91505060e08301516147cb60e086018273ffffffffffffffffffffffffffffffffffffffff169052565b5061010083810151908501526101209283015192909301919091525090565b602081526000613b0360208301846146e1565b6000815260008251614816816001850160208701613b0a565b9190910160010192915050565b6000826148325761483261438f565b500490565b6000821982111561484a5761484a6142c6565b500190565b67ffffffffffffffff8316815260408101613b036020830184613bf8565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156148a5576148a56142c6565b500290565b600063ffffffff8083168185168083038211156148c9576148c96142c6565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526149028184018a613c7a565b905082810360808401526149168189613c7a565b905060ff871660a084015282810360c08401526149338187613b36565b905067ffffffffffffffff851660e08401528281036101008401526149588185613b36565b9c9b505050505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006144ce60408301846146e1565b600082601f8301126149a857600080fd5b815160206149b86138a5836138dc565b82815260059290921b840181019181810190868411156149d757600080fd5b8286015b848110156139645780516149ee816137f2565b83529183019183016149db565b600082601f830112614a0c57600080fd5b81516020614a1c6138a5836138dc565b82815260059290921b84018101918181019086841115614a3b57600080fd5b8286015b8481101561396457805167ffffffffffffffff811115614a5f5760008081fd5b614a6d8986838b0101614489565b845250918301918301614a3f565b600060208284031215614a8d57600080fd5b815167ffffffffffffffff80821115614aa557600080fd5b908301906101008286031215614aba57600080fd5b614ac261375c565b825182811115614ad157600080fd5b614add87828601614997565b825250602083015182811115614af257600080fd5b614afe878286016144d6565b602083015250604083015182811115614b1657600080fd5b614b228782860161453a565b604083015250606083015182811115614b3a57600080fd5b614b46878286016149fb565b606083015250608083015182811115614b5e57600080fd5b614b6a8782860161453a565b60808301525060a083015160a082015260c083015182811115614b8c57600080fd5b614b988782860161453a565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b83811015613cc057815167ffffffffffffffff1687529582019590820190600101614bc8565b600081518084526020808501808196508360051b8101915082860160005b85811015614162578284038952614c24848351613b36565b98850198935090840190600101614c0c565b6040815260008351610100806040850152614c55610140850183614bb4565b915060208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080868503016060870152614c918483613c7a565b93506040880151915080868503016080870152614cae84836146b1565b935060608801519150808685030160a0870152614ccb8483614bee565b935060808801519150808685030160c0870152614ce884836146b1565b935060a088015160e087015260c0880151915080868503018387015250614d0f83826146b1565b60e088015161012087015286151560208701529350613b0392505050565b60006020808385031215614d4057600080fd5b825167ffffffffffffffff811115614d5757600080fd5b8301601f81018513614d6857600080fd5b8051614d766138a5826138dc565b81815260079190911b82018301908381019087831115614d9557600080fd5b928401925b82841015614e045760808489031215614db35760008081fd5b614dbb613780565b8451614dc6816137f2565b815284860151868201526040808601519082015260608086015160048110614dee5760008081fd5b9082015282526080939093019290840190614d9a565b979650505050505050565b60a081526000614e2260a08301886146b1565b8281036020840152614e3481886146b1565b90508560408401528281036060840152614e4e81866146b1565b9150508260808301529695505050505050565b600060208284031215614e7357600080fd5b5051919050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b166040850152816060850152614ec18285018b613c7a565b91508382036080850152614ed5828a613c7a565b915060ff881660a085015283820360c0850152614ef28288613b36565b90861660e085015283810361010085015290506149588185613b3656fea164736f6c634300080d000a",
}

var Any2EVMTollOffRampABI = Any2EVMTollOffRampMetaData.ABI

var Any2EVMTollOffRampBin = Any2EVMTollOffRampMetaData.Bin

func DeployAny2EVMTollOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, chainId *big.Int, offRampConfig TollOffRampInterfaceOffRampConfig, blobVerifier common.Address, onRampAddress common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, maxTimeWithoutAFNSignal *big.Int) (common.Address, *types.Transaction, *Any2EVMTollOffRamp, error) {
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) ExecutedMessages(opts *bind.CallOpts, arg0 uint64) (uint8, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "executedMessages", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) ExecutedMessages(arg0 uint64) (uint8, error) {
	return _Any2EVMTollOffRamp.Contract.ExecutedMessages(&_Any2EVMTollOffRamp.CallOpts, arg0)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) ExecutedMessages(arg0 uint64) (uint8, error) {
	return _Any2EVMTollOffRamp.Contract.ExecutedMessages(&_Any2EVMTollOffRamp.CallOpts, arg0)
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCaller) SRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRamp.contract.Call(opts, &out, "s_router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) SRouter() (common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.SRouter(&_Any2EVMTollOffRamp.CallOpts)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampCallerSession) SRouter() (common.Address, error) {
	return _Any2EVMTollOffRamp.Contract.SRouter(&_Any2EVMTollOffRamp.CallOpts)
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) Execute(opts *bind.TransactOpts, report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "execute", report, needFee)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) Execute(report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.Execute(&_Any2EVMTollOffRamp.TransactOpts, report, needFee)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) Execute(report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.Execute(&_Any2EVMTollOffRamp.TransactOpts, report, needFee)
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.SetConfig(&_Any2EVMTollOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRamp.Contract.SetConfig(&_Any2EVMTollOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterExecutionCompleted(opts *bind.FilterOpts) (*Any2EVMTollOffRampExecutionCompletedIterator, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "ExecutionCompleted")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampExecutionCompletedIterator{contract: _Any2EVMTollOffRamp.contract, event: "ExecutionCompleted", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampExecutionCompleted) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "ExecutionCompleted")
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
	Config TollOffRampInterfaceOffRampConfig
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

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampOffRampRouterSetIterator, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "OffRampRouterSet")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampOffRampRouterSetIterator{contract: _Any2EVMTollOffRamp.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampOffRampRouterSet) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "OffRampRouterSet")
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
	return common.HexToHash("0xf0d733e2ae2689a0e5857664088b68b5fc1b4cbeb757cd5397882d46f5791952")
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

	ExecutedMessages(opts *bind.CallOpts, arg0 uint64) (uint8, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

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

	SRouter(opts *bind.CallOpts) (common.Address, error)

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, report CCIPExecutionReport, needFee bool) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMTollMessage) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

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

	FilterExecutionCompleted(opts *bind.FilterOpts) (*Any2EVMTollOffRampExecutionCompletedIterator, error)

	WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampExecutionCompleted) (event.Subscription, error)

	ParseExecutionCompleted(log types.Log) (*Any2EVMTollOffRampExecutionCompleted, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*Any2EVMTollOffRampOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampOffRampRouterSet) (event.Subscription, error)

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
