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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampRelayed\",\"type\":\"uint256\"},{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"indexed\":true,\"internalType\":\"structCCIP.ExecutionResult\",\"name\":\"results\",\"type\":\"tuple\"}],\"name\":\"ExecutionCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTollOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"needFee\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampRelayed\",\"type\":\"uint256\"},{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structCCIP.ExecutionResult[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"executedMessages\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_router\",\"outputs\":[{\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620055da380380620055da8339810160408190526200003491620005ff565b6000805460ff191681558890889088908890889088908890889060019084908490879085903390819081620000b05760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ea57620000ea8162000305565b5050506001600160a01b038216158062000102575080155b156200012157604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0393909316929092179091556003558051825114620001675760405162d8548360e71b815260040160405180910390fd5b81516200017c906005906020850190620003b6565b5060005b825181101562000260576000828281518110620001a157620001a16200071c565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001eb57620001eb6200071c565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff1916600117905580620002578162000732565b91505062000180565b50505015156080525050845160a05250505060c0929092528051600f5560208101516010805460408401516060909401516001600160401b03908116600160801b02600160801b600160c01b031995821668010000000000000000026001600160801b031990931691909416171792909216179055600e80546001600160a01b039092166001600160a01b0319909216919091179055506200075a9650505050505050565b336001600160a01b038216036200035f5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a7565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200040e579160200282015b828111156200040e57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003d7565b506200041c92915062000420565b5090565b5b808211156200041c576000815560010162000421565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b038111828210171562000472576200047262000437565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620004a357620004a362000437565b604052919050565b80516001600160401b0381168114620004c357600080fd5b919050565b6001600160a01b0381168114620004de57600080fd5b50565b8051620004c381620004c8565b60006001600160401b038211156200050a576200050a62000437565b5060051b60200190565b600082601f8301126200052657600080fd5b815160206200053f6200053983620004ee565b62000478565b82815260059290921b840181019181810190868411156200055f57600080fd5b8286015b84811015620005875780516200057981620004c8565b835291830191830162000563565b509695505050505050565b600082601f830112620005a457600080fd5b81516020620005b76200053983620004ee565b82815260059290921b84018101918181019086841115620005d757600080fd5b8286015b8481101562000587578051620005f181620004c8565b8352918301918301620005db565b600080600080600080600080888a036101608112156200061e57600080fd5b895198506080601f19820112156200063557600080fd5b50620006406200044d565b60208a015181526200065560408b01620004ab565b60208201526200066860608b01620004ab565b60408201526200067b60808b01620004ab565b606082015296506200069060a08a01620004e1565b9550620006a060c08a01620004e1565b9450620006b060e08a01620004e1565b6101008a01519094506001600160401b0380821115620006cf57600080fd5b620006dd8c838d0162000514565b94506101208b0151915080821115620006f557600080fd5b50620007048b828c0162000592565b92505061014089015190509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b6000600182016200075357634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c051614e496200079160003960006103da015260008181610350015261311601526000610f3b0152614e496000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c806385e1f4d01161010f578063b6608c3b116100a2578063e16e632c11610071578063e16e632c14610502578063e3d0e71214610522578063eb511dd414610535578063f2fde38b1461054857600080fd5b8063b6608c3b14610483578063bbe4f6db14610496578063be9b03f1146104cf578063c0d78655146104ef57600080fd5b8063afcb95d7116100de578063afcb95d714610435578063b034909c14610455578063b1dc65a41461045d578063b57671661461047057600080fd5b806385e1f4d0146103d557806389c06568146103fc5780638bbad066146104045780638da5cb5b1461041257600080fd5b80635c975abb1161018757806379ba50971161015657806379ba509714610380578063814118341461038857806381ff70481461039d5780638456cb59146103cd57600080fd5b80635c975abb146102fd5780636edcbf3814610308578063744b92e21461033857806374be21501461034b57600080fd5b80632222dd42116101c35780632222dd421461025a5780633f4ba83a14610299578063567c814b146102a15780635b16ebb7146102c457600080fd5b8063092cddc2146101ea578063108ee5fc146101ff578063181f5a7714610212575b600080fd5b6101fd6101f83660046138cb565b61055b565b005b6101fd61020d3660046139e7565b6105b7565b604080518082018252601881527f416e793245564d546f6c6c4f666652616d7020312e302e300000000000000000602082015290516102519190613a81565b60405180910390f35b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610251565b6101fd610693565b6102b46102af366004613a94565b6106a5565b6040519015158152602001610251565b6102b46102d23660046139e7565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff166102b4565b61032b610316366004613aad565b60116020526000908152604090205460ff1681565b6040516102519190613b34565b6101fd610346366004613b42565b6107ec565b6103727f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610251565b6101fd610beb565b610390610d12565b6040516102519190613bcc565b6009546007546040805163ffffffff80851682526401000000009094049093166020840152820152606001610251565b6101fd610d81565b6103727f000000000000000000000000000000000000000000000000000000000000000081565b610390610d91565b6101fd6101e5366004613bdf565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610274565b604080516001815260006020820181905291810191909152606001610251565b600354610372565b6101fd61046b366004613c67565b610dfe565b6101fd61047e366004613d4c565b6114a7565b6101fd610491366004613a94565b6114b3565b6102746104a43660046139e7565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6104e26104dd366004613eea565b611533565b6040516102519190614033565b6101fd6104fd3660046139e7565b611d60565b600d546102749073ffffffffffffffffffffffffffffffffffffffff1681565b6101fd6105303660046140be565b611de1565b6101fd610543366004613b42565b6127c6565b6101fd6105563660046139e7565b612a0e565b333014610594576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ab8160a001518260c001518360600151612a1f565b6105b481612a80565b50565b6105bf612b83565b73ffffffffffffffffffffffffffffffffffffffff811661060c576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b61069b612b83565b6106a3612c09565b565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa158015610715573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610739919061418b565b1580156107e65750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa1580156107b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d591906141a8565b602001516107e39084614233565b11155b92915050565b6107f4612b83565b6005546000819003610832576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906108cd576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610936576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006005610945600185614233565b815481106109555761095561424a565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff16815481106109a7576109a761424a565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1660056109d6600186614233565b815481106109e6576109e661424a565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff1681548110610a5457610a5461424a565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480610af657610af6614279565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610c71576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600c805480602002602001604051908101604052809291908181526020018280548015610d7757602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610d4c575b5050505050905090565b610d89612b83565b6106a3612cea565b60606005805480602002602001604051908101604052809291908181526020018280548015610d775760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610d4c575050505050905090565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c013591610e5491849163ffffffff851691908e908e9081908401838280828437600092019190915250612daa92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260075480825260085460ff80821660208501526101009091041692820192909252908314610f29576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610c68565b610f378b8b8b8b8b8b612f56565b60007f000000000000000000000000000000000000000000000000000000000000000015610f9457600282602001518360400151610f7591906142a8565b610f7f91906142fc565b610f8a9060016142a8565b60ff169050610faa565b6020820151610fa49060016142a8565b60ff1690505b888114611013576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610c68565b88871461107c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610c68565b336000908152600a602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156110bf576110bf613aca565b60028111156110d0576110d0613aca565b90525090506002816020015160028111156110ed576110ed613aca565b1480156111345750600c816000015160ff168154811061110f5761110f61424a565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b61119a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610c68565b5050505050600088886040516111b192919061431e565b6040519081900381206111c8918c9060200161432e565b6040516020818303038152906040528051906020012090506111e86135e5565b604080518082019091526000808252602082015260005b8881101561148557600060018588846020811061121e5761121e61424a565b61122b91901a601b6142a8565b8d8d8681811061123d5761123d61424a565b905060200201358c8c878181106112565761125661424a565b9050602002013560405160008152602001604052604051611293949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156112b5573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff81166000908152600a602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561133557611335613aca565b600281111561134657611346613aca565b905250925060018360200151600281111561136357611363613aca565b146113ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401610c68565b8251849060ff16601f81106113e1576113e161424a565b60200201511561144d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401610c68565b600184846000015160ff16601f81106114685761146861424a565b91151560209092020152508061147d8161434a565b9150506111ff565b5050505063ffffffff811061149c5761149c614382565b505050505050505050565b6105b460008083612daa565b6114bb612b83565b806000036114f5576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c2519101610687565b606061154160005460ff1690565b156115a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c68565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611615573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611639919061418b565b1561166f576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa1580156116df573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061170391906141a8565b90506003548160200151426117189190614233565b1115611750576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5473ffffffffffffffffffffffffffffffffffffffff1661179f576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608401515160008190036117e0576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff8111156117fb576117fb613604565b604051908082528060200260200182016040528015611824578160200160208202803683370190505b50905060008267ffffffffffffffff81111561184257611842613604565b60405190808252806020026020018201604052801561191557816020015b61190260405180610140016040528060008152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b8152602001906001900390816118605790505b50905060005b838110156119e057876060015181815181106119395761193961424a565b602002602001015180602001905181019061195491906144cb565b8282815181106119665761196661424a565b60200260200101819052508181815181106119835761198361424a565b602002602001015160405160200161199b9190614720565b604051602081830303815290604052805190602001208382815181106119c3576119c361424a565b6020908102919091010152806119d88161434a565b91505061191b565b50600080611a01848a608001518b60a001518c60c001518d60e0015161300d565b915091506000835182611a149190614733565b905060008667ffffffffffffffff811115611a3157611a31613604565b604051908082528060200260200182016040528015611a8a57816020015b611a776040805160808101825260008082526020820181905291810182905290606082015290565b815260200190600190039081611a4f5790505b50905060005b87811015611d515760005a90506000878381518110611ab157611ab161424a565b602002602001015190506000611ae4826020015167ffffffffffffffff1660009081526011602052604090205460ff1690565b90506002816003811115611afa57611afa613aca565b03611b435760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610c68565b611b4c82613112565b60005b8260a0015151811015611b9557611b828360a001518281518110611b7557611b7561424a565b6020026020010151613277565b5080611b8d8161434a565b915050611b4f565b506003816003811115611baa57611baa613aca565b14158015611bb557508d5b15611bce57611bce8260e00151836101000151306132f3565b60208281015167ffffffffffffffff16600090815260119091526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055611c1e83613390565b60208085015167ffffffffffffffff166000908152601190915260409020805491925082917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001836003811115611c7957611c79613aca565b02179055506000875a611c8c9087614233565b611c969190614747565b905060006040518060800160405280866020015167ffffffffffffffff1681526020018381526020018c8152602001846003811115611cd757611cd7613aca565b815250905080888881518110611cef57611cef61424a565b602002602001018190525080604051611d08919061475f565b604051908190038120907f0127bb0341b85f0846a2f2de50be702b328557f51ec8ca98a05bc12edfcfb8a390600090a25050505050508080611d499061434a565b915050611a90565b509a9950505050505050505050565b611d68612b83565b600d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d49060200160405180910390a150565b855185518560ff16601f831115611e54576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401610c68565b60008111611ebe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610c68565b818314611f4c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610c68565b611f5781600361479a565b8311611fbf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610c68565b611fc7612b83565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b600b54156121ba57600b5460009061201f90600190614233565b90506000600b82815481106120365761203661424a565b6000918252602082200154600c805473ffffffffffffffffffffffffffffffffffffffff909216935090849081106120705761207061424a565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff8581168452600a909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055600b805491925090806120f0576120f0614279565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055600c80548061215957612159614279565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905501905550612005915050565b60005b815151811015612621576000600a6000846000015184815181106121e3576121e361424a565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff16600281111561222d5761222d613aca565b14612294576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610c68565b6040805180820190915260ff821681526001602082015282518051600a91600091859081106122c5576122c561424a565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561236657612366613aca565b0217905550600091506123769050565b600a6000846020015184815181106123905761239061424a565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002054610100900460ff1660028111156123da576123da613aca565b14612441576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610c68565b6040805180820190915260ff82168152602081016002815250600a6000846020015184815181106124745761247461424a565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561251557612515613aca565b02179055505082518051600b9250839081106125335761253361424a565b602090810291909101810151825460018101845560009384529282902090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909316929092179091558201518051600c9190839081106125af576125af61424a565b60209081029190910181015182546001810184556000938452919092200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055806126198161434a565b9150506121bd565b506040810151600880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff909216919091179055600980547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff4381168202928317855590830481169360019390926000926126b39286929082169116176147d7565b92506101000a81548163ffffffff021916908363ffffffff1602179055506127124630600960009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a0015161343f565b6007819055825180516008805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560095460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05986127b1988b98919763ffffffff9092169690959194919391926147ff565b60405180910390a15050505050505050505050565b6127ce612b83565b73ffffffffffffffffffffffffffffffffffffffff82161580612805575073ffffffffffffffffffffffffffffffffffffffff8116155b1561283c576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156128d8576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b612a16612b83565b6105b4816134ea565b60005b8351811015612a7a57612a68848281518110612a4057612a4061424a565b6020026020010151848381518110612a5a57612a5a61424a565b6020026020010151846132f3565b80612a728161434a565b915050612a22565b50505050565b606081015173ffffffffffffffffffffffffffffffffffffffff163b612af05760608101516040517f9cfea58300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610c68565b6060810151600d546040517ffd12f6e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063fd12f6e590612b4d9084908690600401614895565b600060405180830381600087803b158015612b6757600080fd5b505af1158015612b7b573d6000803e3d6000fd5b505050505050565b600054610100900473ffffffffffffffffffffffffffffffffffffffff1633146106a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610c68565b60005460ff16612c75576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610c68565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60005460ff1615612d57576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c68565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612cc03390565b600081806020019051810190612dc091906149a8565b6040517fbe9b03f1000000000000000000000000000000000000000000000000000000008152909150600090309063be9b03f190612e05908590600190600401614b63565b6000604051808303816000875af1158015612e24573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052612e6a9190810190614c5a565b905060005b826060015151811015612b7b57600083606001518281518110612e9457612e9461424a565b6020026020010151806020019051810190612eaf91906144cb565b90506000612ec08260e00151613277565b9050600085604001518481518110612eda57612eda61424a565b60200260200101513a868681518110612ef557612ef561424a565b602002602001015160200151612f0b919061479a565b612f15919061479a565b9050600081846101000151612f2a9190614233565b9050612f3f8460e001518286606001516132f3565b505050508080612f4e9061434a565b915050612e6f565b6000612f6382602061479a565b612f6e85602061479a565b612f7a88610144614747565b612f849190614747565b612f8e9190614747565b612f99906000614747565b9050368114613004576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610c68565b50505050505050565b60008060005a600e546040517fe71e65ce00000000000000000000000000000000000000000000000000000000815291925060009173ffffffffffffffffffffffffffffffffffffffff9091169063e71e65ce90613077908c908c908c908c908c90600401614d3c565b6020604051808303816000875af1158015613096573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130ba9190614d8e565b9050600081116130f6576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a6131029084614233565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146131725780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401610c68565b60105460a08201515170010000000000000000000000000000000090910467ffffffffffffffff1610806131b057508060c00151518160a001515114155b156131f95760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610c68565b6010546080820151516801000000000000000090910467ffffffffffffffff1610156105b4576010546080820151516040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090920467ffffffffffffffff1660048301526024820152604401610c68565b73ffffffffffffffffffffffffffffffffffffffff81811660009081526004602052604090205416806132ee576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610c68565b919050565b60006132fe84613277565b6040517fea6192a200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690529192509082169063ea6192a290604401600060405180830381600087803b15801561337257600080fd5b505af1158015613386573d6000803e3d6000fd5b5050505050505050565b6040517f092cddc2000000000000000000000000000000000000000000000000000000008152600090309063092cddc2906133cf908590600401614720565b600060405180830381600087803b1580156133e957600080fd5b505af19250505080156133fa575060015b613437573d808015613428576040519150601f19603f3d011682016040523d82523d6000602084013e61342d565b606091505b5060039392505050565b506002919050565b6000808a8a8a8a8a8a8a8a8a60405160200161346399989796959493929190614da7565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603613569576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610c68565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604051806103e00160405280601f906020820280368337509192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610140810167ffffffffffffffff8111828210171561365757613657613604565b60405290565b604051610100810167ffffffffffffffff8111828210171561365757613657613604565b6040516080810167ffffffffffffffff8111828210171561365757613657613604565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156136eb576136eb613604565b604052919050565b67ffffffffffffffff811681146105b457600080fd5b80356132ee816136f3565b73ffffffffffffffffffffffffffffffffffffffff811681146105b457600080fd5b80356132ee81613714565b600067ffffffffffffffff82111561375b5761375b613604565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261379857600080fd5b81356137ab6137a682613741565b6136a4565b8181528460208386010111156137c057600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff8211156137f7576137f7613604565b5060051b60200190565b600082601f83011261381257600080fd5b813560206138226137a6836137dd565b82815260059290921b8401810191818101908684111561384157600080fd5b8286015b8481101561386557803561385881613714565b8352918301918301613845565b509695505050505050565b600082601f83011261388157600080fd5b813560206138916137a6836137dd565b82815260059290921b840181019181810190868411156138b057600080fd5b8286015b8481101561386557803583529183019183016138b4565b6000602082840312156138dd57600080fd5b813567ffffffffffffffff808211156138f557600080fd5b90830190610140828603121561390a57600080fd5b613912613633565b8235815261392260208401613709565b602082015261393360408401613736565b604082015261394460608401613736565b606082015260808301358281111561395b57600080fd5b61396787828601613787565b60808301525060a08301358281111561397f57600080fd5b61398b87828601613801565b60a08301525060c0830135828111156139a357600080fd5b6139af87828601613870565b60c0830152506139c160e08401613736565b60e082015261010083810135908201526101209283013592810192909252509392505050565b6000602082840312156139f957600080fd5b8135613a0481613714565b9392505050565b60005b83811015613a26578181015183820152602001613a0e565b83811115612a7a5750506000910152565b60008151808452613a4f816020860160208601613a0b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613a046020830184613a37565b600060208284031215613aa657600080fd5b5035919050565b600060208284031215613abf57600080fd5b8135613a04816136f3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110613b30577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b602081016107e68284613af9565b60008060408385031215613b5557600080fd5b8235613b6081613714565b91506020830135613b7081613714565b809150509250929050565b600081518084526020808501945080840160005b83811015613bc157815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101613b8f565b509495945050505050565b602081526000613a046020830184613b7b565b600060208284031215613bf157600080fd5b813567ffffffffffffffff811115613c0857600080fd5b82016101408185031215613a0457600080fd5b60008083601f840112613c2d57600080fd5b50813567ffffffffffffffff811115613c4557600080fd5b6020830191508360208260051b8501011115613c6057600080fd5b9250929050565b60008060008060008060008060e0898b031215613c8357600080fd5b606089018a811115613c9457600080fd5b8998503567ffffffffffffffff80821115613cae57600080fd5b818b0191508b601f830112613cc257600080fd5b813581811115613cd157600080fd5b8c6020828501011115613ce357600080fd5b6020830199508098505060808b0135915080821115613d0157600080fd5b613d0d8c838d01613c1b565b909750955060a08b0135915080821115613d2657600080fd5b50613d338b828c01613c1b565b999c989b50969995989497949560c00135949350505050565b600060208284031215613d5e57600080fd5b813567ffffffffffffffff811115613d7557600080fd5b613d8184828501613787565b949350505050565b600082601f830112613d9a57600080fd5b81356020613daa6137a6836137dd565b82815260059290921b84018101918181019086841115613dc957600080fd5b8286015b84811015613865578035613de0816136f3565b8352918301918301613dcd565b600082601f830112613dfe57600080fd5b81356020613e0e6137a6836137dd565b82815260059290921b84018101918181019086841115613e2d57600080fd5b8286015b84811015613865578035613e4481613714565b8352918301918301613e31565b600082601f830112613e6257600080fd5b81356020613e726137a6836137dd565b82815260059290921b84018101918181019086841115613e9157600080fd5b8286015b8481101561386557803567ffffffffffffffff811115613eb55760008081fd5b613ec38986838b0101613787565b845250918301918301613e95565b80151581146105b457600080fd5b80356132ee81613ed1565b60008060408385031215613efd57600080fd5b823567ffffffffffffffff80821115613f1557600080fd5b908401906101008287031215613f2a57600080fd5b613f3261365d565b823582811115613f4157600080fd5b613f4d88828601613d89565b825250602083013582811115613f6257600080fd5b613f6e88828601613ded565b602083015250604083013582811115613f8657600080fd5b613f9288828601613870565b604083015250606083013582811115613faa57600080fd5b613fb688828601613e51565b606083015250608083013582811115613fce57600080fd5b613fda88828601613870565b60808301525060a083013560a082015260c083013582811115613ffc57600080fd5b61400888828601613870565b60c08301525060e083013560e082015280945050505061402a60208401613edf565b90509250929050565b602080825282518282018190526000919060409081850190868401855b828110156140a0578151805167ffffffffffffffff168552868101518786015285810151868601526060908101519061408b81870183613af9565b50506080939093019290850190600101614050565b5091979650505050505050565b803560ff811681146132ee57600080fd5b60008060008060008060c087890312156140d757600080fd5b863567ffffffffffffffff808211156140ef57600080fd5b6140fb8a838b01613ded565b9750602089013591508082111561411157600080fd5b61411d8a838b01613ded565b965061412b60408a016140ad565b9550606089013591508082111561414157600080fd5b61414d8a838b01613787565b945061415b60808a01613709565b935060a089013591508082111561417157600080fd5b5061417e89828a01613787565b9150509295509295509295565b60006020828403121561419d57600080fd5b8151613a0481613ed1565b6000606082840312156141ba57600080fd5b6040516060810181811067ffffffffffffffff821117156141dd576141dd613604565b80604052508251815260208301516020820152604083015160408201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561424557614245614204565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff038211156142c5576142c5614204565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff83168061430f5761430f6142cd565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361437b5761437b614204565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b80516132ee816136f3565b80516132ee81613714565b600082601f8301126143d857600080fd5b81516143e66137a682613741565b8181528460208386010111156143fb57600080fd5b613d81826020830160208701613a0b565b600082601f83011261441d57600080fd5b8151602061442d6137a6836137dd565b82815260059290921b8401810191818101908684111561444c57600080fd5b8286015b8481101561386557805161446381613714565b8352918301918301614450565b600082601f83011261448157600080fd5b815160206144916137a6836137dd565b82815260059290921b840181019181810190868411156144b057600080fd5b8286015b8481101561386557805183529183019183016144b4565b6000602082840312156144dd57600080fd5b815167ffffffffffffffff808211156144f557600080fd5b90830190610140828603121561450a57600080fd5b614512613633565b82518152614522602084016143b1565b6020820152614533604084016143bc565b6040820152614544606084016143bc565b606082015260808301518281111561455b57600080fd5b614567878286016143c7565b60808301525060a08301518281111561457f57600080fd5b61458b8782860161440c565b60a08301525060c0830151828111156145a357600080fd5b6145af87828601614470565b60c0830152506145c160e084016143bc565b60e082015261010083810151908201526101209283015192810192909252509392505050565b600081518084526020808501945080840160005b83811015613bc1578151875295820195908201906001016145fb565b600061014082518452602083015161463b602086018267ffffffffffffffff169052565b506040830151614663604086018273ffffffffffffffffffffffffffffffffffffffff169052565b50606083015161468b606086018273ffffffffffffffffffffffffffffffffffffffff169052565b5060808301518160808601526146a382860182613a37565b91505060a083015184820360a08601526146bd8282613b7b565b91505060c083015184820360c08601526146d782826145e7565b91505060e083015161470160e086018273ffffffffffffffffffffffffffffffffffffffff169052565b5061010083810151908501526101209283015192909301919091525090565b602081526000613a046020830184614617565b600082614742576147426142cd565b500490565b6000821982111561475a5761475a614204565b500190565b67ffffffffffffffff82511681526020820151602082015260408201516040820152614792606082016060840151613af9565b608001919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156147d2576147d2614204565b500290565b600063ffffffff8083168185168083038211156147f6576147f6614204565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b1660408501525080606084015261482f8184018a613b7b565b905082810360808401526148438189613b7b565b905060ff871660a084015282810360c08401526148608187613a37565b905067ffffffffffffffff851660e08401528281036101008401526148858185613a37565b9c9b505050505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000613d816040830184614617565b600082601f8301126148d557600080fd5b815160206148e56137a6836137dd565b82815260059290921b8401810191818101908684111561490457600080fd5b8286015b8481101561386557805161491b816136f3565b8352918301918301614908565b600082601f83011261493957600080fd5b815160206149496137a6836137dd565b82815260059290921b8401810191818101908684111561496857600080fd5b8286015b8481101561386557805167ffffffffffffffff81111561498c5760008081fd5b61499a8986838b01016143c7565b84525091830191830161496c565b6000602082840312156149ba57600080fd5b815167ffffffffffffffff808211156149d257600080fd5b9083019061010082860312156149e757600080fd5b6149ef61365d565b8251828111156149fe57600080fd5b614a0a878286016148c4565b825250602083015182811115614a1f57600080fd5b614a2b8782860161440c565b602083015250604083015182811115614a4357600080fd5b614a4f87828601614470565b604083015250606083015182811115614a6757600080fd5b614a7387828601614928565b606083015250608083015182811115614a8b57600080fd5b614a9787828601614470565b60808301525060a083015160a082015260c083015182811115614ab957600080fd5b614ac587828601614470565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b83811015613bc157815167ffffffffffffffff1687529582019590820190600101614af5565b600081518084526020808501808196508360051b8101915082860160005b858110156140a0578284038952614b51848351613a37565b98850198935090840190600101614b39565b6040815260008351610100806040850152614b82610140850183614ae1565b915060208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080868503016060870152614bbe8483613b7b565b93506040880151915080868503016080870152614bdb84836145e7565b935060608801519150808685030160a0870152614bf88483614b1b565b935060808801519150808685030160c0870152614c1584836145e7565b935060a088015160e087015260c0880151915080868503018387015250614c3c83826145e7565b60e088015161012087015286151560208701529350613a0492505050565b60006020808385031215614c6d57600080fd5b825167ffffffffffffffff811115614c8457600080fd5b8301601f81018513614c9557600080fd5b8051614ca36137a6826137dd565b81815260079190911b82018301908381019087831115614cc257600080fd5b928401925b82841015614d315760808489031215614ce05760008081fd5b614ce8613681565b8451614cf3816136f3565b815284860151868201526040808601519082015260608086015160048110614d1b5760008081fd5b9082015282526080939093019290840190614cc7565b979650505050505050565b60a081526000614d4f60a08301886145e7565b8281036020840152614d6181886145e7565b90508560408401528281036060840152614d7b81866145e7565b9150508260808301529695505050505050565b600060208284031215614da057600080fd5b5051919050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b166040850152816060850152614dee8285018b613b7b565b91508382036080850152614e02828a613b7b565b915060ff881660a085015283820360c0850152614e1f8288613a37565b90861660e085015283810361010085015290506148858185613a3756fea164736f6c634300080d000a",
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
	Results CCIPExecutionResult
	Raw     types.Log
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) FilterExecutionCompleted(opts *bind.FilterOpts, results []CCIPExecutionResult) (*Any2EVMTollOffRampHelperExecutionCompletedIterator, error) {

	var resultsRule []interface{}
	for _, resultsItem := range results {
		resultsRule = append(resultsRule, resultsItem)
	}

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.FilterLogs(opts, "ExecutionCompleted", resultsRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampHelperExecutionCompletedIterator{contract: _Any2EVMTollOffRampHelper.contract, event: "ExecutionCompleted", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampHelper *Any2EVMTollOffRampHelperFilterer) WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperExecutionCompleted, results []CCIPExecutionResult) (event.Subscription, error) {

	var resultsRule []interface{}
	for _, resultsItem := range results {
		resultsRule = append(resultsRule, resultsItem)
	}

	logs, sub, err := _Any2EVMTollOffRampHelper.contract.WatchLogs(opts, "ExecutionCompleted", resultsRule)
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
	return common.HexToHash("0x0127bb0341b85f0846a2f2de50be702b328557f51ec8ca98a05bc12edfcfb8a3")
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

	FilterExecutionCompleted(opts *bind.FilterOpts, results []CCIPExecutionResult) (*Any2EVMTollOffRampHelperExecutionCompletedIterator, error)

	WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampHelperExecutionCompleted, results []CCIPExecutionResult) (event.Subscription, error)

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
