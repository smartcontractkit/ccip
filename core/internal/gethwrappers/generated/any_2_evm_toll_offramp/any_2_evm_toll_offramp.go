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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampRelayed\",\"type\":\"uint256\"},{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"indexed\":true,\"internalType\":\"structCCIP.ExecutionResult\",\"name\":\"results\",\"type\":\"tuple\"}],\"name\":\"ExecutionCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"needFee\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampRelayed\",\"type\":\"uint256\"},{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structCCIP.ExecutionResult[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"executedMessages\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_router\",\"outputs\":[{\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620055f4380380620055f48339810160408190526200003491620005e6565b6000805460ff1916815560019084908490879085903390819081620000a05760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000da57620000da81620002ec565b5050506001600160a01b0382161580620000f2575080155b156200011157604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001575760405162d8548360e71b815260040160405180910390fd5b81516200016c9060059060208501906200039d565b5060005b82518110156200025057600082828151811062000191576200019162000703565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001db57620001db62000703565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff1916600117905580620002478162000719565b91505062000170565b50505015156080525050845160a05250505060c0929092528051600f5560208101516010805460408401516060909401516001600160401b03908116600160801b02600160801b600160c01b031995821668010000000000000000026001600160801b031990931691909416171792909216179055600e80546001600160a01b039092166001600160a01b031990921691909117905562000741565b336001600160a01b03821603620003465760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000097565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620003f5579160200282015b82811115620003f557825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003be565b506200040392915062000407565b5090565b5b8082111562000403576000815560010162000408565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b03811182821017156200045957620004596200041e565b60405290565b604051601f8201601f191681016001600160401b03811182821017156200048a576200048a6200041e565b604052919050565b80516001600160401b0381168114620004aa57600080fd5b919050565b6001600160a01b0381168114620004c557600080fd5b50565b8051620004aa81620004af565b60006001600160401b03821115620004f157620004f16200041e565b5060051b60200190565b600082601f8301126200050d57600080fd5b81516020620005266200052083620004d5565b6200045f565b82815260059290921b840181019181810190868411156200054657600080fd5b8286015b848110156200056e5780516200056081620004af565b83529183019183016200054a565b509695505050505050565b600082601f8301126200058b57600080fd5b815160206200059e6200052083620004d5565b82815260059290921b84018101918181019086841115620005be57600080fd5b8286015b848110156200056e578051620005d881620004af565b8352918301918301620005c2565b600080600080600080600080888a036101608112156200060557600080fd5b895198506080601f19820112156200061c57600080fd5b506200062762000434565b60208a015181526200063c60408b0162000492565b60208201526200064f60608b0162000492565b60408201526200066260808b0162000492565b606082015296506200067760a08a01620004c8565b95506200068760c08a01620004c8565b94506200069760e08a01620004c8565b6101008a01519094506001600160401b0380821115620006b657600080fd5b620006c48c838d01620004fb565b94506101208b0151915080821115620006dc57600080fd5b50620006eb8b828c0162000579565b92505061014089015190509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b6000600182016200073a57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c051614e7c6200077860003960006103cf01526000818161034501526130ec01526000610f1d0152614e7c6000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c806385e1f4d011610104578063b6608c3b116100a2578063e16e632c11610071578063e16e632c146104e4578063e3d0e71214610504578063eb511dd414610517578063f2fde38b1461052a57600080fd5b8063b6608c3b14610465578063bbe4f6db14610478578063be9b03f1146104b1578063c0d78655146104d157600080fd5b80638da5cb5b116100de5780638da5cb5b14610407578063afcb95d71461042a578063b034909c1461044a578063b1dc65a41461045257600080fd5b806385e1f4d0146103ca57806389c06568146103f15780638bbad066146103f957600080fd5b80635c975abb1161017c57806379ba50971161014b57806379ba509714610375578063814118341461037d57806381ff7048146103925780638456cb59146103c257600080fd5b80635c975abb146102f25780636edcbf38146102fd578063744b92e21461032d57806374be21501461034057600080fd5b80632222dd42116101b85780632222dd421461024f5780633f4ba83a1461028e578063567c814b146102965780635b16ebb7146102b957600080fd5b8063092cddc2146101df578063108ee5fc146101f4578063181f5a7714610207575b600080fd5b6101f26101ed3660046138a1565b61053d565b005b6101f26102023660046139bd565b610599565b604080518082018252601881527f416e793245564d546f6c6c4f666652616d7020312e302e300000000000000000602082015290516102469190613a57565b60405180910390f35b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610246565b6101f2610675565b6102a96102a4366004613a6a565b610687565b6040519015158152602001610246565b6102a96102c73660046139bd565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff166102a9565b61032061030b366004613a83565b60116020526000908152604090205460ff1681565b6040516102469190613b0a565b6101f261033b366004613b18565b6107ce565b6103677f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610246565b6101f2610bcd565b610385610cf4565b6040516102469190613ba2565b6009546007546040805163ffffffff80851682526401000000009094049093166020840152820152606001610246565b6101f2610d63565b6103677f000000000000000000000000000000000000000000000000000000000000000081565b610385610d73565b6101f26101da366004613bb5565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610269565b604080516001815260006020820181905291810191909152606001610246565b600354610367565b6101f2610460366004613c3d565b610de0565b6101f2610473366004613a6a565b611489565b6102696104863660046139bd565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6104c46104bf366004613e83565b611509565b6040516102469190613fcc565b6101f26104df3660046139bd565b611d36565b600d546102699073ffffffffffffffffffffffffffffffffffffffff1681565b6101f2610512366004614057565b611db7565b6101f2610525366004613b18565b61279c565b6101f26105383660046139bd565b6129e4565b333014610576576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61058d8160a001518260c0015183606001516129f5565b61059681612a56565b50565b6105a1612b59565b73ffffffffffffffffffffffffffffffffffffffff81166105ee576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b61067d612b59565b610685612bdf565b565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa1580156106f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071b9190614124565b1580156107c85750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610793573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107b79190614141565b602001516107c590846141cc565b11155b92915050565b6107d6612b59565b6005546000819003610814576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906108af576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610918576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060056109276001856141cc565b81548110610937576109376141e3565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff1681548110610989576109896141e3565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1660056109b86001866141cc565b815481106109c8576109c86141e3565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110610a3657610a366141e3565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480610ad857610ad8614212565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610c53576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600c805480602002602001604051908101604052809291908181526020018280548015610d5957602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610d2e575b5050505050905090565b610d6b612b59565b610685612cc0565b60606005805480602002602001604051908101604052809291908181526020018280548015610d595760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610d2e575050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c013591610e3691849163ffffffff851691908e908e9081908401838280828437600092019190915250612d8092505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260075480825260085460ff80821660208501526101009091041692820192909252908314610f0b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610c4a565b610f198b8b8b8b8b8b612f2c565b60007f000000000000000000000000000000000000000000000000000000000000000015610f7657600282602001518360400151610f579190614241565b610f619190614295565b610f6c906001614241565b60ff169050610f8c565b6020820151610f86906001614241565b60ff1690505b888114610ff5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610c4a565b88871461105e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610c4a565b336000908152600a602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156110a1576110a1613aa0565b60028111156110b2576110b2613aa0565b90525090506002816020015160028111156110cf576110cf613aa0565b1480156111165750600c816000015160ff16815481106110f1576110f16141e3565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b61117c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610c4a565b5050505050600088886040516111939291906142b7565b6040519081900381206111aa918c906020016142c7565b6040516020818303038152906040528051906020012090506111ca6135bb565b604080518082019091526000808252602082015260005b88811015611467576000600185888460208110611200576112006141e3565b61120d91901a601b614241565b8d8d8681811061121f5761121f6141e3565b905060200201358c8c87818110611238576112386141e3565b9050602002013560405160008152602001604052604051611275949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611297573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600a602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561131757611317613aa0565b600281111561132857611328613aa0565b905250925060018360200151600281111561134557611345613aa0565b146113ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610c4a565b8251849060ff16601f81106113c3576113c36141e3565b60200201511561142f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610c4a565b600184846000015160ff16601f811061144a5761144a6141e3565b91151560209092020152508061145f816142e3565b9150506111e1565b5050505063ffffffff811061147e5761147e61431b565b505050505050505050565b611491612b59565b806000036114cb576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c2519101610669565b606061151760005460ff1690565b1561157e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c4a565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061160f9190614124565b15611645576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa1580156116b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116d99190614141565b90506003548160200151426116ee91906141cc565b1115611726576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5473ffffffffffffffffffffffffffffffffffffffff16611775576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608401515160008190036117b6576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff8111156117d1576117d16135da565b6040519080825280602002602001820160405280156117fa578160200160208202803683370190505b50905060008267ffffffffffffffff811115611818576118186135da565b6040519080825280602002602001820160405280156118eb57816020015b6118d860405180610140016040528060008152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b8152602001906001900390816118365790505b50905060005b838110156119b6578760600151818151811061190f5761190f6141e3565b602002602001015180602001905181019061192a919061446c565b82828151811061193c5761193c6141e3565b6020026020010181905250818181518110611959576119596141e3565b602002602001015160405160200161197191906146c1565b60405160208183030381529060405280519060200120838281518110611999576119996141e3565b6020908102919091010152806119ae816142e3565b9150506118f1565b506000806119d7848a608001518b60a001518c60c001518d60e00151612fe3565b9150915060008351826119ea91906146d4565b905060008667ffffffffffffffff811115611a0757611a076135da565b604051908082528060200260200182016040528015611a6057816020015b611a4d6040805160808101825260008082526020820181905291810182905290606082015290565b815260200190600190039081611a255790505b50905060005b87811015611d275760005a90506000878381518110611a8757611a876141e3565b602002602001015190506000611aba826020015167ffffffffffffffff1660009081526011602052604090205460ff1690565b90506002816003811115611ad057611ad0613aa0565b03611b195760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610c4a565b611b22826130e8565b60005b8260a0015151811015611b6b57611b588360a001518281518110611b4b57611b4b6141e3565b602002602001015161324d565b5080611b63816142e3565b915050611b25565b506003816003811115611b8057611b80613aa0565b14158015611b8b57508d5b15611ba457611ba48260e00151836101000151306132c9565b60208281015167ffffffffffffffff16600090815260119091526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611bf483613366565b60208085015167ffffffffffffffff166000908152601190915260409020805491925082917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001836003811115611c4f57611c4f613aa0565b02179055506000875a611c6290876141cc565b611c6c91906146e8565b905060006040518060800160405280866020015167ffffffffffffffff1681526020018381526020018c8152602001846003811115611cad57611cad613aa0565b815250905080888881518110611cc557611cc56141e3565b602002602001018190525080604051611cde9190614700565b604051908190038120907f0127bb0341b85f0846a2f2de50be702b328557f51ec8ca98a05bc12edfcfb8a390600090a25050505050508080611d1f906142e3565b915050611a66565b509a9950505050505050505050565b611d3e612b59565b600d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d49060200160405180910390a150565b855185518560ff16601f831115611e2a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610c4a565b60008111611e94576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610c4a565b818314611f22576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610c4a565b611f2d81600361473b565b8311611f95576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610c4a565b611f9d612b59565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600b541561219057600b54600090611ff5906001906141cc565b90506000600b828154811061200c5761200c6141e3565b6000918252602082200154600c805473ffffffffffffffffffffffffffffffffffffffff90921693509084908110612046576120466141e3565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600a909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600b805491925090806120c6576120c6614212565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600c80548061212f5761212f614212565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550611fdb915050565b60005b8151518110156125f7576000600a6000846000015184815181106121b9576121b96141e3565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561220357612203613aa0565b1461226a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610c4a565b6040805180820190915260ff821681526001602082015282518051600a916000918590811061229b5761229b6141e3565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561233c5761233c613aa0565b02179055506000915061234c9050565b600a600084602001518481518110612366576123666141e3565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff1660028111156123b0576123b0613aa0565b14612417576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610c4a565b6040805180820190915260ff82168152602081016002815250600a60008460200151848151811061244a5761244a6141e3565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156124eb576124eb613aa0565b02179055505082518051600b925083908110612509576125096141e3565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600c919083908110612585576125856141e3565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055806125ef816142e3565b915050612193565b506040810151600880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600980547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092612689928692908216911617614778565b92506101000a81548163ffffffff021916908363ffffffff1602179055506126e84630600960009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613415565b6007819055825180516008805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560095460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612787988b98919763ffffffff9092169690959194919391926147a0565b60405180910390a15050505050505050505050565b6127a4612b59565b73ffffffffffffffffffffffffffffffffffffffff821615806127db575073ffffffffffffffffffffffffffffffffffffffff8116155b15612812576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156128ae576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b6129ec612b59565b610596816134c0565b60005b8351811015612a5057612a3e848281518110612a1657612a166141e3565b6020026020010151848381518110612a3057612a306141e3565b6020026020010151846132c9565b80612a48816142e3565b9150506129f8565b50505050565b606081015173ffffffffffffffffffffffffffffffffffffffff163b612ac65760608101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610c4a565b6060810151600d546040517ffd12f6e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063fd12f6e590612b239084908690600401614836565b600060405180830381600087803b158015612b3d57600080fd5b505af1158015612b51573d6000803e3d6000fd5b505050505050565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610685576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610c4a565b60005460ff16612c4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610c4a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60005460ff1615612d2d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c4a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612c963390565b600081806020019051810190612d9691906149ad565b6040517fbe9b03f1000000000000000000000000000000000000000000000000000000008152909150600090309063be9b03f190612ddb908590600190600401614b96565b6000604051808303816000875af1158015612dfa573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052612e409190810190614c8d565b905060005b826060015151811015612b5157600083606001518281518110612e6a57612e6a6141e3565b6020026020010151806020019051810190612e85919061446c565b90506000612e968260e0015161324d565b9050600085604001518481518110612eb057612eb06141e3565b60200260200101513a868681518110612ecb57612ecb6141e3565b602002602001015160200151612ee1919061473b565b612eeb919061473b565b9050600081846101000151612f0091906141cc565b9050612f158460e001518286606001516132c9565b505050508080612f24906142e3565b915050612e45565b6000612f3982602061473b565b612f4485602061473b565b612f50886101446146e8565b612f5a91906146e8565b612f6491906146e8565b612f6f9060006146e8565b9050368114612fda576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610c4a565b50505050505050565b60008060005a600e546040517fe71e65ce00000000000000000000000000000000000000000000000000000000815291925060009173ffffffffffffffffffffffffffffffffffffffff9091169063e71e65ce9061304d908c908c908c908c908c90600401614d6f565b6020604051808303816000875af115801561306c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130909190614dc1565b9050600081116130cc576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a6130d890846141cc565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146131485780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610c4a565b60105460a08201515170010000000000000000000000000000000090910467ffffffffffffffff16108061318657508060c00151518160a001515114155b156131cf5760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610c4a565b6010546080820151516801000000000000000090910467ffffffffffffffff161015610596576010546080820151516040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090920467ffffffffffffffff1660048301526024820152604401610c4a565b73ffffffffffffffffffffffffffffffffffffffff81811660009081526004602052604090205416806132c4576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610c4a565b919050565b60006132d48461324d565b6040517fea6192a200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690529192509082169063ea6192a290604401600060405180830381600087803b15801561334857600080fd5b505af115801561335c573d6000803e3d6000fd5b5050505050505050565b6040517f092cddc2000000000000000000000000000000000000000000000000000000008152600090309063092cddc2906133a59085906004016146c1565b600060405180830381600087803b1580156133bf57600080fd5b505af19250505080156133d0575060015b61340d573d8080156133fe576040519150601f19603f3d011682016040523d82523d6000602084013e613403565b606091505b5060039392505050565b506002919050565b6000808a8a8a8a8a8a8a8a8a60405160200161343999989796959493929190614dda565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff82160361353f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610c4a565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604051806103e00160405280601f906020820280368337509192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610140810167ffffffffffffffff8111828210171561362d5761362d6135da565b60405290565b604051610100810167ffffffffffffffff8111828210171561362d5761362d6135da565b6040516080810167ffffffffffffffff8111828210171561362d5761362d6135da565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156136c1576136c16135da565b604052919050565b67ffffffffffffffff8116811461059657600080fd5b80356132c4816136c9565b73ffffffffffffffffffffffffffffffffffffffff8116811461059657600080fd5b80356132c4816136ea565b600067ffffffffffffffff821115613731576137316135da565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261376e57600080fd5b813561378161377c82613717565b61367a565b81815284602083860101111561379657600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff8211156137cd576137cd6135da565b5060051b60200190565b600082601f8301126137e857600080fd5b813560206137f861377c836137b3565b82815260059290921b8401810191818101908684111561381757600080fd5b8286015b8481101561383b57803561382e816136ea565b835291830191830161381b565b509695505050505050565b600082601f83011261385757600080fd5b8135602061386761377c836137b3565b82815260059290921b8401810191818101908684111561388657600080fd5b8286015b8481101561383b578035835291830191830161388a565b6000602082840312156138b357600080fd5b813567ffffffffffffffff808211156138cb57600080fd5b9083019061014082860312156138e057600080fd5b6138e8613609565b823581526138f8602084016136df565b60208201526139096040840161370c565b604082015261391a6060840161370c565b606082015260808301358281111561393157600080fd5b61393d8782860161375d565b60808301525060a08301358281111561395557600080fd5b613961878286016137d7565b60a08301525060c08301358281111561397957600080fd5b61398587828601613846565b60c08301525061399760e0840161370c565b60e082015261010083810135908201526101209283013592810192909252509392505050565b6000602082840312156139cf57600080fd5b81356139da816136ea565b9392505050565b60005b838110156139fc5781810151838201526020016139e4565b83811115612a505750506000910152565b60008151808452613a258160208601602086016139e1565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006139da6020830184613a0d565b600060208284031215613a7c57600080fd5b5035919050565b600060208284031215613a9557600080fd5b81356139da816136c9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110613b06577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b602081016107c88284613acf565b60008060408385031215613b2b57600080fd5b8235613b36816136ea565b91506020830135613b46816136ea565b809150509250929050565b600081518084526020808501945080840160005b83811015613b9757815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101613b65565b509495945050505050565b6020815260006139da6020830184613b51565b600060208284031215613bc757600080fd5b813567ffffffffffffffff811115613bde57600080fd5b820161014081850312156139da57600080fd5b60008083601f840112613c0357600080fd5b50813567ffffffffffffffff811115613c1b57600080fd5b6020830191508360208260051b8501011115613c3657600080fd5b9250929050565b60008060008060008060008060e0898b031215613c5957600080fd5b606089018a811115613c6a57600080fd5b8998503567ffffffffffffffff80821115613c8457600080fd5b818b0191508b601f830112613c9857600080fd5b813581811115613ca757600080fd5b8c6020828501011115613cb957600080fd5b6020830199508098505060808b0135915080821115613cd757600080fd5b613ce38c838d01613bf1565b909750955060a08b0135915080821115613cfc57600080fd5b50613d098b828c01613bf1565b999c989b50969995989497949560c00135949350505050565b600082601f830112613d3357600080fd5b81356020613d4361377c836137b3565b82815260059290921b84018101918181019086841115613d6257600080fd5b8286015b8481101561383b578035613d79816136c9565b8352918301918301613d66565b600082601f830112613d9757600080fd5b81356020613da761377c836137b3565b82815260059290921b84018101918181019086841115613dc657600080fd5b8286015b8481101561383b578035613ddd816136ea565b8352918301918301613dca565b600082601f830112613dfb57600080fd5b81356020613e0b61377c836137b3565b82815260059290921b84018101918181019086841115613e2a57600080fd5b8286015b8481101561383b57803567ffffffffffffffff811115613e4e5760008081fd5b613e5c8986838b010161375d565b845250918301918301613e2e565b801515811461059657600080fd5b80356132c481613e6a565b60008060408385031215613e9657600080fd5b823567ffffffffffffffff80821115613eae57600080fd5b908401906101008287031215613ec357600080fd5b613ecb613633565b823582811115613eda57600080fd5b613ee688828601613d22565b825250602083013582811115613efb57600080fd5b613f0788828601613d86565b602083015250604083013582811115613f1f57600080fd5b613f2b88828601613846565b604083015250606083013582811115613f4357600080fd5b613f4f88828601613dea565b606083015250608083013582811115613f6757600080fd5b613f7388828601613846565b60808301525060a083013560a082015260c083013582811115613f9557600080fd5b613fa188828601613846565b60c08301525060e083013560e0820152809450505050613fc360208401613e78565b90509250929050565b602080825282518282018190526000919060409081850190868401855b82811015614039578151805167ffffffffffffffff168552868101518786015285810151868601526060908101519061402481870183613acf565b50506080939093019290850190600101613fe9565b5091979650505050505050565b803560ff811681146132c457600080fd5b60008060008060008060c0878903121561407057600080fd5b863567ffffffffffffffff8082111561408857600080fd5b6140948a838b01613d86565b975060208901359150808211156140aa57600080fd5b6140b68a838b01613d86565b96506140c460408a01614046565b955060608901359150808211156140da57600080fd5b6140e68a838b0161375d565b94506140f460808a016136df565b935060a089013591508082111561410a57600080fd5b5061411789828a0161375d565b9150509295509295509295565b60006020828403121561413657600080fd5b81516139da81613e6a565b60006060828403121561415357600080fd5b6040516060810181811067ffffffffffffffff82111715614176576141766135da565b80604052508251815260208301516020820152604083015160408201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156141de576141de61419d565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff0382111561425e5761425e61419d565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff8316806142a8576142a8614266565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036143145761431461419d565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b80516132c4816136c9565b80516132c4816136ea565b600082601f83011261437157600080fd5b815161437f61377c82613717565b81815284602083860101111561439457600080fd5b6143a58260208301602087016139e1565b949350505050565b600082601f8301126143be57600080fd5b815160206143ce61377c836137b3565b82815260059290921b840181019181810190868411156143ed57600080fd5b8286015b8481101561383b578051614404816136ea565b83529183019183016143f1565b600082601f83011261442257600080fd5b8151602061443261377c836137b3565b82815260059290921b8401810191818101908684111561445157600080fd5b8286015b8481101561383b5780518352918301918301614455565b60006020828403121561447e57600080fd5b815167ffffffffffffffff8082111561449657600080fd5b9083019061014082860312156144ab57600080fd5b6144b3613609565b825181526144c36020840161434a565b60208201526144d460408401614355565b60408201526144e560608401614355565b60608201526080830151828111156144fc57600080fd5b61450887828601614360565b60808301525060a08301518281111561452057600080fd5b61452c878286016143ad565b60a08301525060c08301518281111561454457600080fd5b61455087828601614411565b60c08301525061456260e08401614355565b60e082015261010083810151908201526101209283015192810192909252509392505050565b600081518084526020808501945080840160005b83811015613b975781518752958201959082019060010161459c565b60006101408251845260208301516145dc602086018267ffffffffffffffff169052565b506040830151614604604086018273ffffffffffffffffffffffffffffffffffffffff169052565b50606083015161462c606086018273ffffffffffffffffffffffffffffffffffffffff169052565b50608083015181608086015261464482860182613a0d565b91505060a083015184820360a086015261465e8282613b51565b91505060c083015184820360c08601526146788282614588565b91505060e08301516146a260e086018273ffffffffffffffffffffffffffffffffffffffff169052565b5061010083810151908501526101209283015192909301919091525090565b6020815260006139da60208301846145b8565b6000826146e3576146e3614266565b500490565b600082198211156146fb576146fb61419d565b500190565b67ffffffffffffffff82511681526020820151602082015260408201516040820152614733606082016060840151613acf565b608001919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156147735761477361419d565b500290565b600063ffffffff8083168185168083038211156147975761479761419d565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526147d08184018a613b51565b905082810360808401526147e48189613b51565b905060ff871660a084015282810360c08401526148018187613a0d565b905067ffffffffffffffff851660e08401528281036101008401526148268185613a0d565b9c9b505050505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006143a560408301846145b8565b600082601f83011261487657600080fd5b8151602061488661377c836137b3565b82815260059290921b840181019181810190868411156148a557600080fd5b8286015b8481101561383b5780516148bc816136c9565b83529183019183016148a9565b600082601f8301126148da57600080fd5b815160206148ea61377c836137b3565b82815260059290921b8401810191818101908684111561490957600080fd5b8286015b8481101561383b578051614920816136ea565b835291830191830161490d565b600082601f83011261493e57600080fd5b8151602061494e61377c836137b3565b82815260059290921b8401810191818101908684111561496d57600080fd5b8286015b8481101561383b57805167ffffffffffffffff8111156149915760008081fd5b61499f8986838b0101614360565b845250918301918301614971565b6000602082840312156149bf57600080fd5b815167ffffffffffffffff808211156149d757600080fd5b9083019061010082860312156149ec57600080fd5b6149f4613633565b825182811115614a0357600080fd5b614a0f87828601614865565b825250602083015182811115614a2457600080fd5b614a30878286016148c9565b602083015250604083015182811115614a4857600080fd5b614a5487828601614411565b604083015250606083015182811115614a6c57600080fd5b614a788782860161492d565b606083015250608083015182811115614a9057600080fd5b614a9c87828601614411565b60808301525060a083015160a082015260c083015182811115614abe57600080fd5b614aca87828601614411565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b83811015613b9757815167ffffffffffffffff1687529582019590820190600101614afa565b600082825180855260208086019550808260051b84010181860160005b84811015614b89577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868403018952614b77838351613a0d565b98840198925090830190600101614b3d565b5090979650505050505050565b6040815260008351610100806040850152614bb5610140850183614ae6565b915060208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080868503016060870152614bf18483613b51565b93506040880151915080868503016080870152614c0e8483614588565b935060608801519150808685030160a0870152614c2b8483614b20565b935060808801519150808685030160c0870152614c488483614588565b935060a088015160e087015260c0880151915080868503018387015250614c6f8382614588565b60e0880151610120870152861515602087015293506139da92505050565b60006020808385031215614ca057600080fd5b825167ffffffffffffffff811115614cb757600080fd5b8301601f81018513614cc857600080fd5b8051614cd661377c826137b3565b81815260079190911b82018301908381019087831115614cf557600080fd5b928401925b82841015614d645760808489031215614d135760008081fd5b614d1b613657565b8451614d26816136c9565b815284860151868201526040808601519082015260608086015160048110614d4e5760008081fd5b9082015282526080939093019290840190614cfa565b979650505050505050565b60a081526000614d8260a0830188614588565b8281036020840152614d948188614588565b90508560408401528281036060840152614dae8186614588565b9150508260808301529695505050505050565b600060208284031215614dd357600080fd5b5051919050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b166040850152816060850152614e218285018b613b51565b91508382036080850152614e35828a613b51565b915060ff881660a085015283820360c0850152614e528288613a0d565b90861660e085015283810361010085015290506148268185613a0d56fea164736f6c634300080d000a",
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
	Results CCIPExecutionResult
	Raw     types.Log
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) FilterExecutionCompleted(opts *bind.FilterOpts, results []CCIPExecutionResult) (*Any2EVMTollOffRampExecutionCompletedIterator, error) {

	var resultsRule []interface{}
	for _, resultsItem := range results {
		resultsRule = append(resultsRule, resultsItem)
	}

	logs, sub, err := _Any2EVMTollOffRamp.contract.FilterLogs(opts, "ExecutionCompleted", resultsRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampExecutionCompletedIterator{contract: _Any2EVMTollOffRamp.contract, event: "ExecutionCompleted", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRamp *Any2EVMTollOffRampFilterer) WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampExecutionCompleted, results []CCIPExecutionResult) (event.Subscription, error) {

	var resultsRule []interface{}
	for _, resultsItem := range results {
		resultsRule = append(resultsRule, resultsItem)
	}

	logs, sub, err := _Any2EVMTollOffRamp.contract.WatchLogs(opts, "ExecutionCompleted", resultsRule)
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
	return common.HexToHash("0x0127bb0341b85f0846a2f2de50be702b328557f51ec8ca98a05bc12edfcfb8a3")
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

	FilterExecutionCompleted(opts *bind.FilterOpts, results []CCIPExecutionResult) (*Any2EVMTollOffRampExecutionCompletedIterator, error)

	WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampExecutionCompleted, results []CCIPExecutionResult) (event.Subscription, error)

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
