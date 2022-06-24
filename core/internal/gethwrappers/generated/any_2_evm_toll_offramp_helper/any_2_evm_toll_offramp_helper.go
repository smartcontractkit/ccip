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

var Any2EVMTollOffRampHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"needFee\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampRelayed\",\"type\":\"uint256\"},{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structCCIP.ExecutionResult[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"executedMessages\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_router\",\"outputs\":[{\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005770380380620057708339810160408190526200003491620005ff565b6000805460ff191681558890889088908890889088908890889060019084908490879085903390819081620000b05760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ea57620000ea8162000305565b5050506001600160a01b038216158062000102575080155b156200012157604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001675760405162d8548360e71b815260040160405180910390fd5b81516200017c906005906020850190620003b6565b5060005b825181101562000260576000828281518110620001a157620001a16200071c565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001eb57620001eb6200071c565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff1916600117905580620002578162000732565b91505062000180565b50505015156080525050845160a05250505060c0929092528051600f5560208101516010805460408401516060909401516001600160401b03908116600160801b02600160801b600160c01b031995821668010000000000000000026001600160801b031990931691909416171792909216179055600e80546001600160a01b039092166001600160a01b0319909216919091179055506200075a9650505050505050565b336001600160a01b038216036200035f5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a7565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200040e579160200282015b828111156200040e57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003d7565b506200041c92915062000420565b5090565b5b808211156200041c576000815560010162000421565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b038111828210171562000472576200047262000437565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620004a357620004a362000437565b604052919050565b80516001600160401b0381168114620004c357600080fd5b919050565b6001600160a01b0381168114620004de57600080fd5b50565b8051620004c381620004c8565b60006001600160401b038211156200050a576200050a62000437565b5060051b60200190565b600082601f8301126200052657600080fd5b815160206200053f6200053983620004ee565b62000478565b82815260059290921b840181019181810190868411156200055f57600080fd5b8286015b84811015620005875780516200057981620004c8565b835291830191830162000563565b509695505050505050565b600082601f830112620005a457600080fd5b81516020620005b76200053983620004ee565b82815260059290921b84018101918181019086841115620005d757600080fd5b8286015b8481101562000587578051620005f181620004c8565b8352918301918301620005db565b600080600080600080600080888a036101608112156200061e57600080fd5b895198506080601f19820112156200063557600080fd5b50620006406200044d565b60208a015181526200065560408b01620004ab565b60208201526200066860608b01620004ab565b60408201526200067b60808b01620004ab565b606082015296506200069060a08a01620004e1565b9550620006a060c08a01620004e1565b9450620006b060e08a01620004e1565b6101008a01519094506001600160401b0380821115620006cf57600080fd5b620006dd8c838d0162000514565b94506101208b0151915080821115620006f557600080fd5b50620007048b828c0162000592565b92505061014089015190509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b6000600182016200075357634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c051614fdf6200079160003960006103da015260008181610350015261323f01526000610f3b0152614fdf6000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c806385e1f4d01161010f578063b6608c3b116100a2578063e16e632c11610071578063e16e632c14610502578063e3d0e71214610522578063eb511dd414610535578063f2fde38b1461054857600080fd5b8063b6608c3b14610483578063bbe4f6db14610496578063be9b03f1146104cf578063c0d78655146104ef57600080fd5b8063afcb95d7116100de578063afcb95d714610435578063b034909c14610455578063b1dc65a41461045d578063b57671661461047057600080fd5b806385e1f4d0146103d557806389c06568146103fc5780638bbad066146104045780638da5cb5b1461041257600080fd5b80635c975abb1161018757806379ba50971161015657806379ba509714610380578063814118341461038857806381ff70481461039d5780638456cb59146103cd57600080fd5b80635c975abb146102fd5780636edcbf3814610308578063744b92e21461033857806374be21501461034b57600080fd5b80632222dd42116101c35780632222dd421461025a5780633f4ba83a14610299578063567c814b146102a15780635b16ebb7146102c457600080fd5b8063092cddc2146101ea578063108ee5fc146101ff578063181f5a7714610212575b600080fd5b6101fd6101f83660046139f4565b61055b565b005b6101fd61020d366004613b10565b6105b7565b604080518082018252601881527f416e793245564d546f6c6c4f666652616d7020312e302e300000000000000000602082015290516102519190613baa565b60405180910390f35b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610251565b6101fd610693565b6102b46102af366004613bbd565b6106a5565b6040519015158152602001610251565b6102b46102d2366004613b10565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff166102b4565b61032b610316366004613bd6565b60116020526000908152604090205460ff1681565b6040516102519190613c5d565b6101fd610346366004613c6b565b6107ec565b6103727f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610251565b6101fd610beb565b610390610d12565b6040516102519190613cf5565b6009546007546040805163ffffffff80851682526401000000009094049093166020840152820152606001610251565b6101fd610d81565b6103727f000000000000000000000000000000000000000000000000000000000000000081565b610390610d91565b6101fd6101e5366004613d08565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610274565b604080516001815260006020820181905291810191909152606001610251565b600354610372565b6101fd61046b366004613d90565b610dfe565b6101fd61047e366004613e75565b6114a7565b6101fd610491366004613bbd565b6114b3565b6102746104a4366004613b10565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6104e26104dd366004614013565b611533565b604051610251919061415c565b6101fd6104fd366004613b10565b611da4565b600d546102749073ffffffffffffffffffffffffffffffffffffffff1681565b6101fd6105303660046141e7565b611e25565b6101fd610543366004613c6b565b61280a565b6101fd610556366004613b10565b612a52565b333014610594576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ab8160a001518260c001518360600151612a63565b6105b481612ac4565b50565b6105bf612bc7565b73ffffffffffffffffffffffffffffffffffffffff811661060c576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b61069b612bc7565b6106a3612c4d565b565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa158015610715573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061073991906142b4565b1580156107e65750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa1580156107b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d591906142d1565b602001516107e3908461435c565b11155b92915050565b6107f4612bc7565b6005546000819003610832576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906108cd576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610936576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600561094560018561435c565b8154811061095557610955614373565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff16815481106109a7576109a7614373565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1660056109d660018661435c565b815481106109e6576109e6614373565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110610a5457610a54614373565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480610af657610af66143a2565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610c71576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600c805480602002602001604051908101604052809291908181526020018280548015610d7757602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610d4c575b5050505050905090565b610d89612bc7565b6106a3612d2e565b60606005805480602002602001604051908101604052809291908181526020018280548015610d775760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610d4c575050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c013591610e5491849163ffffffff851691908e908e9081908401838280828437600092019190915250612dee92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260075480825260085460ff80821660208501526101009091041692820192909252908314610f29576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610c68565b610f378b8b8b8b8b8b61307f565b60007f000000000000000000000000000000000000000000000000000000000000000015610f9457600282602001518360400151610f7591906143d1565b610f7f9190614425565b610f8a9060016143d1565b60ff169050610faa565b6020820151610fa49060016143d1565b60ff1690505b888114611013576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610c68565b88871461107c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610c68565b336000908152600a602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156110bf576110bf613bf3565b60028111156110d0576110d0613bf3565b90525090506002816020015160028111156110ed576110ed613bf3565b1480156111345750600c816000015160ff168154811061110f5761110f614373565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b61119a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610c68565b5050505050600088886040516111b1929190614447565b6040519081900381206111c8918c90602001614457565b6040516020818303038152906040528051906020012090506111e861370e565b604080518082019091526000808252602082015260005b8881101561148557600060018588846020811061121e5761121e614373565b61122b91901a601b6143d1565b8d8d8681811061123d5761123d614373565b905060200201358c8c8781811061125657611256614373565b9050602002013560405160008152602001604052604051611293949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156112b5573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600a602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561133557611335613bf3565b600281111561134657611346613bf3565b905250925060018360200151600281111561136357611363613bf3565b146113ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610c68565b8251849060ff16601f81106113e1576113e1614373565b60200201511561144d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610c68565b600184846000015160ff16601f811061146857611468614373565b91151560209092020152508061147d81614473565b9150506111ff565b5050505063ffffffff811061149c5761149c6144ab565b505050505050505050565b6105b460008083612dee565b6114bb612bc7565b806000036114f5576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c2519101610687565b606061154160005460ff1690565b156115a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c68565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611615573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061163991906142b4565b1561166f576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa1580156116df573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061170391906142d1565b9050600354816020015142611718919061435c565b1115611750576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5473ffffffffffffffffffffffffffffffffffffffff1661179f576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608401515160008190036117e0576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff8111156117fb576117fb61372d565b604051908082528060200260200182016040528015611824578160200160208202803683370190505b50905060008267ffffffffffffffff8111156118425761184261372d565b60405190808252806020026020018201604052801561191557816020015b61190260405180610140016040528060008152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b8152602001906001900390816118605790505b50905060005b83811015611a21578760600151818151811061193957611939614373565b602002602001015180602001905181019061195491906145f4565b82828151811061196657611966614373565b6020026020010181905250600082828151811061198557611985614373565b602002602001015160405160200161199d9190614849565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526119d89160200161485c565b60405160208183030381529060405290508080519060200120848381518110611a0357611a03614373565b60209081029190910101525080611a1981614473565b91505061191b565b50600080611a42848a608001518b60a001518c60c001518d60e00151613136565b915091506000835182611a559190614882565b905060008667ffffffffffffffff811115611a7257611a7261372d565b604051908082528060200260200182016040528015611acb57816020015b611ab86040805160808101825260008082526020820181905291810182905290606082015290565b815260200190600190039081611a905790505b50905060005b87811015611d955760005a90506000878381518110611af257611af2614373565b602002602001015190506000611b25826020015167ffffffffffffffff1660009081526011602052604090205460ff1690565b90506002816003811115611b3b57611b3b613bf3565b03611b845760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610c68565b611b8d8261323b565b60005b8260a0015151811015611bd657611bc38360a001518281518110611bb657611bb6614373565b60200260200101516133a0565b5080611bce81614473565b915050611b90565b506003816003811115611beb57611beb613bf3565b14158015611bf657508d5b15611c0f57611c0f8260e001518361010001513061341c565b60208281015167ffffffffffffffff16600090815260119091526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611c5f836134b9565b60208085015167ffffffffffffffff166000908152601190915260409020805491925082917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001836003811115611cba57611cba613bf3565b02179055506000875a611ccd908761435c565b611cd79190614896565b905060006040518060800160405280866020015167ffffffffffffffff1681526020018381526020018c8152602001846003811115611d1857611d18613bf3565b815250905080888881518110611d3057611d30614373565b60200260200101819052507fbca6416e78a437ab47530846568a4d78457e41bc2adc0d91a826090e2d853d1c81600001518260600151604051611d749291906148ae565b60405180910390a15050505050508080611d8d90614473565b915050611ad1565b509a9950505050505050505050565b611dac612bc7565b600d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d49060200160405180910390a150565b855185518560ff16601f831115611e98576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610c68565b60008111611f02576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610c68565b818314611f90576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610c68565b611f9b8160036148cc565b8311612003576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610c68565b61200b612bc7565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600b54156121fe57600b546000906120639060019061435c565b90506000600b828154811061207a5761207a614373565b6000918252602082200154600c805473ffffffffffffffffffffffffffffffffffffffff909216935090849081106120b4576120b4614373565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600a909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600b80549192509080612134576121346143a2565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600c80548061219d5761219d6143a2565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550612049915050565b60005b815151811015612665576000600a60008460000151848151811061222757612227614373565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561227157612271613bf3565b146122d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610c68565b6040805180820190915260ff821681526001602082015282518051600a916000918590811061230957612309614373565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156123aa576123aa613bf3565b0217905550600091506123ba9050565b600a6000846020015184815181106123d4576123d4614373565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561241e5761241e613bf3565b14612485576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610c68565b6040805180820190915260ff82168152602081016002815250600a6000846020015184815181106124b8576124b8614373565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561255957612559613bf3565b02179055505082518051600b92508390811061257757612577614373565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600c9190839081106125f3576125f3614373565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790558061265d81614473565b915050612201565b506040810151600880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600980547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff4381168202928317855590830481169360019390926000926126f7928692908216911617614909565b92506101000a81548163ffffffff021916908363ffffffff1602179055506127564630600960009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613568565b6007819055825180516008805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560095460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05986127f5988b98919763ffffffff909216969095919491939192614931565b60405180910390a15050505050505050505050565b612812612bc7565b73ffffffffffffffffffffffffffffffffffffffff82161580612849575073ffffffffffffffffffffffffffffffffffffffff8116155b15612880576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152901561291c576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b612a5a612bc7565b6105b481613613565b60005b8351811015612abe57612aac848281518110612a8457612a84614373565b6020026020010151848381518110612a9e57612a9e614373565b60200260200101518461341c565b80612ab681614473565b915050612a66565b50505050565b606081015173ffffffffffffffffffffffffffffffffffffffff163b612b345760608101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610c68565b6060810151600d546040517ffd12f6e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063fd12f6e590612b9190849086906004016149c7565b600060405180830381600087803b158015612bab57600080fd5b505af1158015612bbf573d6000803e3d6000fd5b505050505050565b600054610100900473ffffffffffffffffffffffffffffffffffffffff1633146106a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610c68565b60005460ff16612cb9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610c68565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60005460ff1615612d9b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c68565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612d043390565b600081806020019051810190612e049190614b3e565b6040517fbe9b03f1000000000000000000000000000000000000000000000000000000008152909150600090309063be9b03f190612e49908590600190600401614cf9565b6000604051808303816000875af1158015612e68573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052612eae9190810190614df0565b905060005b826060015151811015612bbf57600083606001518281518110612ed857612ed8614373565b6020026020010151806020019051810190612ef391906145f4565b90506000612f048260e001516133a0565b90506000805b866020015151811015612f9e578360e0015173ffffffffffffffffffffffffffffffffffffffff1687602001518281518110612f4857612f48614373565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1603612f8c5786604001518181518110612f8157612f81614373565b602002602001015191505b80612f9681614473565b915050612f0a565b5080612ff45760e08301516040517fce480bcc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610c68565b6000670de0b6b3a7640000823a88888151811061301357613013614373565b60200260200101516020015161302991906148cc565b61303391906148cc565b61303d9190614882565b9050600081856101000151613052919061435c565b90506130678560e0015182876060015161341c565b5050505050808061307790614473565b915050612eb3565b600061308c8260206148cc565b6130978560206148cc565b6130a388610144614896565b6130ad9190614896565b6130b79190614896565b6130c2906000614896565b905036811461312d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610c68565b50505050505050565b60008060005a600e546040517fe71e65ce00000000000000000000000000000000000000000000000000000000815291925060009173ffffffffffffffffffffffffffffffffffffffff9091169063e71e65ce906131a0908c908c908c908c908c90600401614ed2565b6020604051808303816000875af11580156131bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131e39190614f24565b90506000811161321f576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a61322b908461435c565b9350935050509550959350505050565b80517f00000000000000000000000000000000000000000000000000000000000000001461329b5780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610c68565b60105460a08201515170010000000000000000000000000000000090910467ffffffffffffffff1610806132d957508060c00151518160a001515114155b156133225760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610c68565b6010546080820151516801000000000000000090910467ffffffffffffffff1610156105b4576010546080820151516040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090920467ffffffffffffffff1660048301526024820152604401610c68565b73ffffffffffffffffffffffffffffffffffffffff8181166000908152600460205260409020541680613417576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610c68565b919050565b6000613427846133a0565b6040517fea6192a200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690529192509082169063ea6192a290604401600060405180830381600087803b15801561349b57600080fd5b505af11580156134af573d6000803e3d6000fd5b5050505050505050565b6040517f092cddc2000000000000000000000000000000000000000000000000000000008152600090309063092cddc2906134f8908590600401614849565b600060405180830381600087803b15801561351257600080fd5b505af1925050508015613523575060015b613560573d808015613551576040519150601f19603f3d011682016040523d82523d6000602084013e613556565b606091505b5060039392505050565b506002919050565b6000808a8a8a8a8a8a8a8a8a60405160200161358c99989796959493929190614f3d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603613692576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610c68565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604051806103e00160405280601f906020820280368337509192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610140810167ffffffffffffffff811182821017156137805761378061372d565b60405290565b604051610100810167ffffffffffffffff811182821017156137805761378061372d565b6040516080810167ffffffffffffffff811182821017156137805761378061372d565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156138145761381461372d565b604052919050565b67ffffffffffffffff811681146105b457600080fd5b80356134178161381c565b73ffffffffffffffffffffffffffffffffffffffff811681146105b457600080fd5b80356134178161383d565b600067ffffffffffffffff8211156138845761388461372d565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f8301126138c157600080fd5b81356138d46138cf8261386a565b6137cd565b8181528460208386010111156138e957600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff8211156139205761392061372d565b5060051b60200190565b600082601f83011261393b57600080fd5b8135602061394b6138cf83613906565b82815260059290921b8401810191818101908684111561396a57600080fd5b8286015b8481101561398e5780356139818161383d565b835291830191830161396e565b509695505050505050565b600082601f8301126139aa57600080fd5b813560206139ba6138cf83613906565b82815260059290921b840181019181810190868411156139d957600080fd5b8286015b8481101561398e57803583529183019183016139dd565b600060208284031215613a0657600080fd5b813567ffffffffffffffff80821115613a1e57600080fd5b908301906101408286031215613a3357600080fd5b613a3b61375c565b82358152613a4b60208401613832565b6020820152613a5c6040840161385f565b6040820152613a6d6060840161385f565b6060820152608083013582811115613a8457600080fd5b613a90878286016138b0565b60808301525060a083013582811115613aa857600080fd5b613ab48782860161392a565b60a08301525060c083013582811115613acc57600080fd5b613ad887828601613999565b60c083015250613aea60e0840161385f565b60e082015261010083810135908201526101209283013592810192909252509392505050565b600060208284031215613b2257600080fd5b8135613b2d8161383d565b9392505050565b60005b83811015613b4f578181015183820152602001613b37565b83811115612abe5750506000910152565b60008151808452613b78816020860160208601613b34565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613b2d6020830184613b60565b600060208284031215613bcf57600080fd5b5035919050565b600060208284031215613be857600080fd5b8135613b2d8161381c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110613c59577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b602081016107e68284613c22565b60008060408385031215613c7e57600080fd5b8235613c898161383d565b91506020830135613c998161383d565b809150509250929050565b600081518084526020808501945080840160005b83811015613cea57815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101613cb8565b509495945050505050565b602081526000613b2d6020830184613ca4565b600060208284031215613d1a57600080fd5b813567ffffffffffffffff811115613d3157600080fd5b82016101408185031215613b2d57600080fd5b60008083601f840112613d5657600080fd5b50813567ffffffffffffffff811115613d6e57600080fd5b6020830191508360208260051b8501011115613d8957600080fd5b9250929050565b60008060008060008060008060e0898b031215613dac57600080fd5b606089018a811115613dbd57600080fd5b8998503567ffffffffffffffff80821115613dd757600080fd5b818b0191508b601f830112613deb57600080fd5b813581811115613dfa57600080fd5b8c6020828501011115613e0c57600080fd5b6020830199508098505060808b0135915080821115613e2a57600080fd5b613e368c838d01613d44565b909750955060a08b0135915080821115613e4f57600080fd5b50613e5c8b828c01613d44565b999c989b50969995989497949560c00135949350505050565b600060208284031215613e8757600080fd5b813567ffffffffffffffff811115613e9e57600080fd5b613eaa848285016138b0565b949350505050565b600082601f830112613ec357600080fd5b81356020613ed36138cf83613906565b82815260059290921b84018101918181019086841115613ef257600080fd5b8286015b8481101561398e578035613f098161381c565b8352918301918301613ef6565b600082601f830112613f2757600080fd5b81356020613f376138cf83613906565b82815260059290921b84018101918181019086841115613f5657600080fd5b8286015b8481101561398e578035613f6d8161383d565b8352918301918301613f5a565b600082601f830112613f8b57600080fd5b81356020613f9b6138cf83613906565b82815260059290921b84018101918181019086841115613fba57600080fd5b8286015b8481101561398e57803567ffffffffffffffff811115613fde5760008081fd5b613fec8986838b01016138b0565b845250918301918301613fbe565b80151581146105b457600080fd5b803561341781613ffa565b6000806040838503121561402657600080fd5b823567ffffffffffffffff8082111561403e57600080fd5b90840190610100828703121561405357600080fd5b61405b613786565b82358281111561406a57600080fd5b61407688828601613eb2565b82525060208301358281111561408b57600080fd5b61409788828601613f16565b6020830152506040830135828111156140af57600080fd5b6140bb88828601613999565b6040830152506060830135828111156140d357600080fd5b6140df88828601613f7a565b6060830152506080830135828111156140f757600080fd5b61410388828601613999565b60808301525060a083013560a082015260c08301358281111561412557600080fd5b61413188828601613999565b60c08301525060e083013560e082015280945050505061415360208401614008565b90509250929050565b602080825282518282018190526000919060409081850190868401855b828110156141c9578151805167ffffffffffffffff16855286810151878601528581015186860152606090810151906141b481870183613c22565b50506080939093019290850190600101614179565b5091979650505050505050565b803560ff8116811461341757600080fd5b60008060008060008060c0878903121561420057600080fd5b863567ffffffffffffffff8082111561421857600080fd5b6142248a838b01613f16565b9750602089013591508082111561423a57600080fd5b6142468a838b01613f16565b965061425460408a016141d6565b9550606089013591508082111561426a57600080fd5b6142768a838b016138b0565b945061428460808a01613832565b935060a089013591508082111561429a57600080fd5b506142a789828a016138b0565b9150509295509295509295565b6000602082840312156142c657600080fd5b8151613b2d81613ffa565b6000606082840312156142e357600080fd5b6040516060810181811067ffffffffffffffff821117156143065761430661372d565b80604052508251815260208301516020820152604083015160408201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561436e5761436e61432d565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff038211156143ee576143ee61432d565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff831680614438576144386143f6565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036144a4576144a461432d565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b80516134178161381c565b80516134178161383d565b600082601f83011261450157600080fd5b815161450f6138cf8261386a565b81815284602083860101111561452457600080fd5b613eaa826020830160208701613b34565b600082601f83011261454657600080fd5b815160206145566138cf83613906565b82815260059290921b8401810191818101908684111561457557600080fd5b8286015b8481101561398e57805161458c8161383d565b8352918301918301614579565b600082601f8301126145aa57600080fd5b815160206145ba6138cf83613906565b82815260059290921b840181019181810190868411156145d957600080fd5b8286015b8481101561398e57805183529183019183016145dd565b60006020828403121561460657600080fd5b815167ffffffffffffffff8082111561461e57600080fd5b90830190610140828603121561463357600080fd5b61463b61375c565b8251815261464b602084016144da565b602082015261465c604084016144e5565b604082015261466d606084016144e5565b606082015260808301518281111561468457600080fd5b614690878286016144f0565b60808301525060a0830151828111156146a857600080fd5b6146b487828601614535565b60a08301525060c0830151828111156146cc57600080fd5b6146d887828601614599565b60c0830152506146ea60e084016144e5565b60e082015261010083810151908201526101209283015192810192909252509392505050565b600081518084526020808501945080840160005b83811015613cea57815187529582019590820190600101614724565b6000610140825184526020830151614764602086018267ffffffffffffffff169052565b50604083015161478c604086018273ffffffffffffffffffffffffffffffffffffffff169052565b5060608301516147b4606086018273ffffffffffffffffffffffffffffffffffffffff169052565b5060808301518160808601526147cc82860182613b60565b91505060a083015184820360a08601526147e68282613ca4565b91505060c083015184820360c08601526148008282614710565b91505060e083015161482a60e086018273ffffffffffffffffffffffffffffffffffffffff169052565b5061010083810151908501526101209283015192909301919091525090565b602081526000613b2d6020830184614740565b6000815260008251614875816001850160208701613b34565b9190910160010192915050565b600082614891576148916143f6565b500490565b600082198211156148a9576148a961432d565b500190565b67ffffffffffffffff8316815260408101613b2d6020830184613c22565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156149045761490461432d565b500290565b600063ffffffff8083168185168083038211156149285761492861432d565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526149618184018a613ca4565b905082810360808401526149758189613ca4565b905060ff871660a084015282810360c08401526149928187613b60565b905067ffffffffffffffff851660e08401528281036101008401526149b78185613b60565b9c9b505050505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000613eaa6040830184614740565b600082601f830112614a0757600080fd5b81516020614a176138cf83613906565b82815260059290921b84018101918181019086841115614a3657600080fd5b8286015b8481101561398e578051614a4d8161381c565b8352918301918301614a3a565b600082601f830112614a6b57600080fd5b81516020614a7b6138cf83613906565b82815260059290921b84018101918181019086841115614a9a57600080fd5b8286015b8481101561398e578051614ab18161383d565b8352918301918301614a9e565b600082601f830112614acf57600080fd5b81516020614adf6138cf83613906565b82815260059290921b84018101918181019086841115614afe57600080fd5b8286015b8481101561398e57805167ffffffffffffffff811115614b225760008081fd5b614b308986838b01016144f0565b845250918301918301614b02565b600060208284031215614b5057600080fd5b815167ffffffffffffffff80821115614b6857600080fd5b908301906101008286031215614b7d57600080fd5b614b85613786565b825182811115614b9457600080fd5b614ba0878286016149f6565b825250602083015182811115614bb557600080fd5b614bc187828601614a5a565b602083015250604083015182811115614bd957600080fd5b614be587828601614599565b604083015250606083015182811115614bfd57600080fd5b614c0987828601614abe565b606083015250608083015182811115614c2157600080fd5b614c2d87828601614599565b60808301525060a083015160a082015260c083015182811115614c4f57600080fd5b614c5b87828601614599565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b83811015613cea57815167ffffffffffffffff1687529582019590820190600101614c8b565b600081518084526020808501808196508360051b8101915082860160005b858110156141c9578284038952614ce7848351613b60565b98850198935090840190600101614ccf565b6040815260008351610100806040850152614d18610140850183614c77565b915060208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080868503016060870152614d548483613ca4565b93506040880151915080868503016080870152614d718483614710565b935060608801519150808685030160a0870152614d8e8483614cb1565b935060808801519150808685030160c0870152614dab8483614710565b935060a088015160e087015260c0880151915080868503018387015250614dd28382614710565b60e088015161012087015286151560208701529350613b2d92505050565b60006020808385031215614e0357600080fd5b825167ffffffffffffffff811115614e1a57600080fd5b8301601f81018513614e2b57600080fd5b8051614e396138cf82613906565b81815260079190911b82018301908381019087831115614e5857600080fd5b928401925b82841015614ec75760808489031215614e765760008081fd5b614e7e6137aa565b8451614e898161381c565b815284860151868201526040808601519082015260608086015160048110614eb15760008081fd5b9082015282526080939093019290840190614e5d565b979650505050505050565b60a081526000614ee560a0830188614710565b8281036020840152614ef78188614710565b90508560408401528281036060840152614f118186614710565b9150508260808301529695505050505050565b600060208284031215614f3657600080fd5b5051919050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b166040850152816060850152614f848285018b613ca4565b91508382036080850152614f98828a613ca4565b915060ff881660a085015283820360c0850152614fb58288613b60565b90861660e085015283810361010085015290506149b78185613b6056fea164736f6c634300080f000a",
}

var Any2EVMTollOffRampHelperABI = Any2EVMTollOffRampHelperMetaData.ABI

var Any2EVMTollOffRampHelperBin = Any2EVMTollOffRampHelperMetaData.Bin

func DeployAny2EVMTollOffRampHelper(auth *bind.TransactOpts, backend bind.ContractBackend, chainId *big.Int, offRampConfig TollOffRampInterfaceOffRampConfig, blobVerifier common.Address, onRampAddress common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, maxTimeWithoutAFNSignal *big.Int) (common.Address, *types.Transaction, *Any2EVMTollOffRampHelper, error) {
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) ExecutedMessages(opts *bind.CallOpts, arg0 uint64) (uint8, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "executedMessages", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) ExecutedMessages(arg0 uint64) (uint8, error) {
	return _Any2EVMTollOffRampHelper.Contract.ExecutedMessages(&_Any2EVMTollOffRampHelper.CallOpts, arg0)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) ExecutedMessages(arg0 uint64) (uint8, error) {
	return _Any2EVMTollOffRampHelper.Contract.ExecutedMessages(&_Any2EVMTollOffRampHelper.CallOpts, arg0)
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCaller) SRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampHelper.contract.Call(opts, &out, "s_router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) SRouter() (common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.SRouter(&_Any2EVMTollOffRampHelper.CallOpts)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperCallerSession) SRouter() (common.Address, error) {
	return _Any2EVMTollOffRampHelper.Contract.SRouter(&_Any2EVMTollOffRampHelper.CallOpts)
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) Execute(opts *bind.TransactOpts, report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "execute", report, needFee)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) Execute(report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.Execute(&_Any2EVMTollOffRampHelper.TransactOpts, report, needFee)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) Execute(report CCIPExecutionReport, needFee bool) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.Execute(&_Any2EVMTollOffRampHelper.TransactOpts, report, needFee)
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.SetConfig(&_Any2EVMTollOffRampHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Any2EVMTollOffRampHelper.Contract.SetConfig(&_Any2EVMTollOffRampHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterExecutionCompleted(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperExecutionCompletedIterator, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "ExecutionCompleted")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperExecutionCompletedIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "ExecutionCompleted", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperExecutionCompleted) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "ExecutionCompleted")
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
	Config TollOffRampInterfaceOffRampConfig
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

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperOffRampRouterSetIterator, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "OffRampRouterSet")
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperOffRampRouterSetIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperOffRampRouterSet) (event.Subscription, error) {

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "OffRampRouterSet")
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
	return common.HexToHash("0xf0d733e2ae2689a0e5857664088b68b5fc1b4cbeb757cd5397882d46f5791952")
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

	Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

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

	FilterExecutionCompleted(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperExecutionCompletedIterator, error)

	WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperExecutionCompleted) (event.Subscription, error)

	ParseExecutionCompleted(log types.Log) (*Any2EVMTollOffRampHelperExecutionCompleted, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*Any2EVMTollOffRampHelperOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts) (*Any2EVMTollOffRampHelperOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperOffRampRouterSet) (event.Subscription, error)

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
