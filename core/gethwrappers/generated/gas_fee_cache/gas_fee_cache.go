// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gas_fee_cache

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

type GEFeeUpdate struct {
	Token          common.Address
	ChainId        uint64
	LinkPerUnitGas *big.Int
}

var FeeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"linkPerUnitGas\",\"type\":\"uint128\"}],\"internalType\":\"structGE.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"feeUpdaters\",\"type\":\"address[]\"},{\"internalType\":\"uint128\",\"name\":\"stalenessThreshold\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawalAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByUpdaterOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePassed\",\"type\":\"uint256\"}],\"name\":\"StaleFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chain\",\"type\":\"uint64\"}],\"name\":\"TokenOrChainNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChain\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"linkPerUnitGas\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"name\":\"GasFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"fee\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStalenessThreshold\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"removeFeeUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"setFeeUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"linkPerUnitGas\",\"type\":\"uint128\"}],\"internalType\":\"structGE.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"}],\"name\":\"updateFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620015f4380380620015f48339810160408190526200003491620004aa565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620001b9565b50505060005b8351811015620001555762000142848281518110620000e757620000e7620005d3565b602002602001015160000151858381518110620001085762000108620005d3565b602002602001015160200151868481518110620001295762000129620005d3565b602002602001015160400151426200026460201b60201c565b6200014d81620005e9565b9050620000c4565b5060005b8251811015620001a357620001908382815181106200017c576200017c620005d3565b60200260200101516200030260201b60201c565b6200019b81620005e9565b905062000159565b506001600160801b031660805250620006119050565b336001600160a01b03821603620002135760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040805180820182526001600160801b0384811680835284821660208085018281526001600160a01b038b166000818152600284528881206001600160401b038d1680835290855290899020975192518716600160801b029290961691909117909555855192835282015290927fe45338fd766b2619bbcd30dd0c79dcc00795e7977bab27952f0e78571a100764910160405180910390a350505050565b6001600160a01b038116156200035c576001600160a01b038116600081815260036020526040808220805460ff19166001179055517fa462a3423511fce5ad038a1130251b66e3b8c135fa2ca15777f45c72ea3037c59190a25b50565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156200039a576200039a6200035f565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620003cb57620003cb6200035f565b604052919050565b60006001600160401b03821115620003ef57620003ef6200035f565b5060051b60200190565b80516001600160a01b03811681146200041157600080fd5b919050565b80516001600160801b03811681146200041157600080fd5b600082601f8301126200044057600080fd5b81516020620004596200045383620003d3565b620003a0565b82815260059290921b840181019181810190868411156200047957600080fd5b8286015b848110156200049f576200049181620003f9565b83529183019183016200047d565b509695505050505050565b60008060006060808587031215620004c157600080fd5b84516001600160401b0380821115620004d957600080fd5b818701915087601f830112620004ee57600080fd5b81516020620005016200045383620003d3565b8281529185028401810191818101908b8411156200051e57600080fd5b948201945b838610156200058f5786868d0312156200053d5760008081fd5b6200054762000375565b6200055287620003f9565b8152838701518681168114620005685760008081fd5b8185015260406200057b88820162000416565b908201528252948601949082019062000523565b918a0151919850909450505080831115620005a957600080fd5b5050620005b9868287016200042e565b925050620005ca6040850162000416565b90509250925092565b634e487b7160e01b600052603260045260246000fd5b6000600182016200060a57634e487b7160e01b600052601160045260246000fd5b5060010190565b608051610fc0620006346000396000818161014e01526104190152610fc06000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80638e160ef411610076578063a6c94a731161005b578063a6c94a731461014c578063ae7fca1814610172578063f2fde38b1461018557600080fd5b80638e160ef4146101055780639086658e1461013957600080fd5b806301e33667146100a8578063604782e6146100bd57806379ba5097146100d05780638da5cb5b146100d8575b600080fd5b6100bb6100b6366004610bfd565b610198565b005b6100bb6100cb366004610c39565b610213565b6100bb610227565b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b610118610113366004610c6c565b61030f565b6040516fffffffffffffffffffffffffffffffff90911681526020016100fc565b6100bb610147366004610d46565b610499565b7f0000000000000000000000000000000000000000000000000000000000000000610118565b6100bb610180366004610c39565b610588565b6100bb610193366004610c39565b610599565b6101a06105aa565b73ffffffffffffffffffffffffffffffffffffffff82166101ed576040517f84c2102600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61020e73ffffffffffffffffffffffffffffffffffffffff84168383610613565b505050565b61021b6105aa565b610224816106a0565b50565b60015473ffffffffffffffffffffffffffffffffffffffff1633146102935760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260026020908152604080832067ffffffffffffffff8516845282528083208151808301909252546fffffffffffffffffffffffffffffffff8082168352700100000000000000000000000000000000909104169181018290529015806103a2575080516fffffffffffffffffffffffffffffffff16155b15610401576040517f102e3c2800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015267ffffffffffffffff8416602482015260440161028a565b60208101516fffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811691600091610449911642610e76565b90508181111561048f576040517f1f49f5b8000000000000000000000000000000000000000000000000000000008152600481018390526024810182905260440161028a565b5050519392505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633148015906104d157503360009081526003602052604090205460ff16155b15610508576040517f46f0815400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4260005b825181101561020e5761057883828151811061052a5761052a610e8d565b60200260200101516000015184838151811061054857610548610e8d565b60200260200101516020015185848151811061056657610566610e8d565b60200260200101516040015185610733565b61058181610ebc565b905061050c565b6105906105aa565b610224816107f5565b6105a16105aa565b61022481610869565b60005473ffffffffffffffffffffffffffffffffffffffff1633146106115760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161028a565b565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905261020e908490610944565b73ffffffffffffffffffffffffffffffffffffffff8116156102245773ffffffffffffffffffffffffffffffffffffffff811660008181526003602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055517fa462a3423511fce5ad038a1130251b66e3b8c135fa2ca15777f45c72ea3037c59190a250565b6040805180820182526fffffffffffffffffffffffffffffffff848116808352848216602080850182815273ffffffffffffffffffffffffffffffffffffffff8b1660008181526002845288812067ffffffffffffffff8d1680835290855290899020975192518716700100000000000000000000000000000000029290961691909117909555855192835282015290927fe45338fd766b2619bbcd30dd0c79dcc00795e7977bab27952f0e78571a100764910160405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff811660008181526003602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517f74a2c31badb27f0acfb9da3ef34c9e656ca1723881466e89a40f791f1c82ee719190a250565b3373ffffffffffffffffffffffffffffffffffffffff8216036108ce5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161028a565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006109a6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16610a369092919063ffffffff16565b80519091501561020e57808060200190518101906109c49190610ef4565b61020e5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161028a565b6060610a458484600085610a4f565b90505b9392505050565b606082471015610ac75760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161028a565b843b610b155760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161028a565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610b3e9190610f46565b60006040518083038185875af1925050503d8060008114610b7b576040519150601f19603f3d011682016040523d82523d6000602084013e610b80565b606091505b5091509150610b90828286610b9b565b979650505050505050565b60608315610baa575081610a48565b825115610bba5782518084602001fd5b8160405162461bcd60e51b815260040161028a9190610f62565b803573ffffffffffffffffffffffffffffffffffffffff81168114610bf857600080fd5b919050565b600080600060608486031215610c1257600080fd5b610c1b84610bd4565b9250610c2960208501610bd4565b9150604084013590509250925092565b600060208284031215610c4b57600080fd5b610a4882610bd4565b803567ffffffffffffffff81168114610bf857600080fd5b60008060408385031215610c7f57600080fd5b610c8883610bd4565b9150610c9660208401610c54565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715610cf157610cf1610c9f565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610d3e57610d3e610c9f565b604052919050565b60006020808385031215610d5957600080fd5b823567ffffffffffffffff80821115610d7157600080fd5b818501915085601f830112610d8557600080fd5b813581811115610d9757610d97610c9f565b610da5848260051b01610cf7565b81815284810192506060918202840185019188831115610dc457600080fd5b938501935b82851015610e3b5780858a031215610de15760008081fd5b610de9610cce565b610df286610bd4565b8152610dff878701610c54565b878201526040808701356fffffffffffffffffffffffffffffffff81168114610e285760008081fd5b9082015284529384019392850192610dc9565b50979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015610e8857610e88610e47565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610eed57610eed610e47565b5060010190565b600060208284031215610f0657600080fd5b81518015158114610a4857600080fd5b60005b83811015610f31578181015183820152602001610f19565b83811115610f40576000848401525b50505050565b60008251610f58818460208701610f16565b9190910192915050565b6020815260008251806020840152610f81816040850160208701610f16565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fea164736f6c634300080f000a",
}

var FeeManagerABI = FeeManagerMetaData.ABI

var FeeManagerBin = FeeManagerMetaData.Bin

func DeployFeeManager(auth *bind.TransactOpts, backend bind.ContractBackend, feeUpdates []GEFeeUpdate, feeUpdaters []common.Address, stalenessThreshold *big.Int) (common.Address, *types.Transaction, *FeeManager, error) {
	parsed, err := FeeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeManagerBin), backend, feeUpdates, feeUpdaters, stalenessThreshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeManager{FeeManagerCaller: FeeManagerCaller{contract: contract}, FeeManagerTransactor: FeeManagerTransactor{contract: contract}, FeeManagerFilterer: FeeManagerFilterer{contract: contract}}, nil
}

type FeeManager struct {
	address common.Address
	abi     abi.ABI
	FeeManagerCaller
	FeeManagerTransactor
	FeeManagerFilterer
}

type FeeManagerCaller struct {
	contract *bind.BoundContract
}

type FeeManagerTransactor struct {
	contract *bind.BoundContract
}

type FeeManagerFilterer struct {
	contract *bind.BoundContract
}

type FeeManagerSession struct {
	Contract     *FeeManager
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type FeeManagerCallerSession struct {
	Contract *FeeManagerCaller
	CallOpts bind.CallOpts
}

type FeeManagerTransactorSession struct {
	Contract     *FeeManagerTransactor
	TransactOpts bind.TransactOpts
}

type FeeManagerRaw struct {
	Contract *FeeManager
}

type FeeManagerCallerRaw struct {
	Contract *FeeManagerCaller
}

type FeeManagerTransactorRaw struct {
	Contract *FeeManagerTransactor
}

func NewFeeManager(address common.Address, backend bind.ContractBackend) (*FeeManager, error) {
	abi, err := abi.JSON(strings.NewReader(FeeManagerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindFeeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeManager{address: address, abi: abi, FeeManagerCaller: FeeManagerCaller{contract: contract}, FeeManagerTransactor: FeeManagerTransactor{contract: contract}, FeeManagerFilterer: FeeManagerFilterer{contract: contract}}, nil
}

func NewFeeManagerCaller(address common.Address, caller bind.ContractCaller) (*FeeManagerCaller, error) {
	contract, err := bindFeeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeManagerCaller{contract: contract}, nil
}

func NewFeeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeManagerTransactor, error) {
	contract, err := bindFeeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeManagerTransactor{contract: contract}, nil
}

func NewFeeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeManagerFilterer, error) {
	contract, err := bindFeeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeManagerFilterer{contract: contract}, nil
}

func bindFeeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_FeeManager *FeeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeManager.Contract.FeeManagerCaller.contract.Call(opts, result, method, params...)
}

func (_FeeManager *FeeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeManager.Contract.FeeManagerTransactor.contract.Transfer(opts)
}

func (_FeeManager *FeeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeManager.Contract.FeeManagerTransactor.contract.Transact(opts, method, params...)
}

func (_FeeManager *FeeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeManager.Contract.contract.Call(opts, result, method, params...)
}

func (_FeeManager *FeeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeManager.Contract.contract.Transfer(opts)
}

func (_FeeManager *FeeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeManager.Contract.contract.Transact(opts, method, params...)
}

func (_FeeManager *FeeManagerCaller) GetFee(opts *bind.CallOpts, token common.Address, destChainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "getFee", token, destChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeManager *FeeManagerSession) GetFee(token common.Address, destChainId uint64) (*big.Int, error) {
	return _FeeManager.Contract.GetFee(&_FeeManager.CallOpts, token, destChainId)
}

func (_FeeManager *FeeManagerCallerSession) GetFee(token common.Address, destChainId uint64) (*big.Int, error) {
	return _FeeManager.Contract.GetFee(&_FeeManager.CallOpts, token, destChainId)
}

func (_FeeManager *FeeManagerCaller) GetStalenessThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "getStalenessThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeManager *FeeManagerSession) GetStalenessThreshold() (*big.Int, error) {
	return _FeeManager.Contract.GetStalenessThreshold(&_FeeManager.CallOpts)
}

func (_FeeManager *FeeManagerCallerSession) GetStalenessThreshold() (*big.Int, error) {
	return _FeeManager.Contract.GetStalenessThreshold(&_FeeManager.CallOpts)
}

func (_FeeManager *FeeManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FeeManager *FeeManagerSession) Owner() (common.Address, error) {
	return _FeeManager.Contract.Owner(&_FeeManager.CallOpts)
}

func (_FeeManager *FeeManagerCallerSession) Owner() (common.Address, error) {
	return _FeeManager.Contract.Owner(&_FeeManager.CallOpts)
}

func (_FeeManager *FeeManagerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "acceptOwnership")
}

func (_FeeManager *FeeManagerSession) AcceptOwnership() (*types.Transaction, error) {
	return _FeeManager.Contract.AcceptOwnership(&_FeeManager.TransactOpts)
}

func (_FeeManager *FeeManagerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FeeManager.Contract.AcceptOwnership(&_FeeManager.TransactOpts)
}

func (_FeeManager *FeeManagerTransactor) RemoveFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "removeFeeUpdater", feeUpdater)
}

func (_FeeManager *FeeManagerSession) RemoveFeeUpdater(feeUpdater common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.RemoveFeeUpdater(&_FeeManager.TransactOpts, feeUpdater)
}

func (_FeeManager *FeeManagerTransactorSession) RemoveFeeUpdater(feeUpdater common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.RemoveFeeUpdater(&_FeeManager.TransactOpts, feeUpdater)
}

func (_FeeManager *FeeManagerTransactor) SetFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "setFeeUpdater", feeUpdater)
}

func (_FeeManager *FeeManagerSession) SetFeeUpdater(feeUpdater common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.SetFeeUpdater(&_FeeManager.TransactOpts, feeUpdater)
}

func (_FeeManager *FeeManagerTransactorSession) SetFeeUpdater(feeUpdater common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.SetFeeUpdater(&_FeeManager.TransactOpts, feeUpdater)
}

func (_FeeManager *FeeManagerTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "transferOwnership", to)
}

func (_FeeManager *FeeManagerSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.TransferOwnership(&_FeeManager.TransactOpts, to)
}

func (_FeeManager *FeeManagerTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.TransferOwnership(&_FeeManager.TransactOpts, to)
}

func (_FeeManager *FeeManagerTransactor) UpdateFees(opts *bind.TransactOpts, feeUpdates []GEFeeUpdate) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "updateFees", feeUpdates)
}

func (_FeeManager *FeeManagerSession) UpdateFees(feeUpdates []GEFeeUpdate) (*types.Transaction, error) {
	return _FeeManager.Contract.UpdateFees(&_FeeManager.TransactOpts, feeUpdates)
}

func (_FeeManager *FeeManagerTransactorSession) UpdateFees(feeUpdates []GEFeeUpdate) (*types.Transaction, error) {
	return _FeeManager.Contract.UpdateFees(&_FeeManager.TransactOpts, feeUpdates)
}

func (_FeeManager *FeeManagerTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "withdrawToken", token, to, amount)
}

func (_FeeManager *FeeManagerSession) WithdrawToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FeeManager.Contract.WithdrawToken(&_FeeManager.TransactOpts, token, to, amount)
}

func (_FeeManager *FeeManagerTransactorSession) WithdrawToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FeeManager.Contract.WithdrawToken(&_FeeManager.TransactOpts, token, to, amount)
}

type FeeManagerFeeUpdaterRemovedIterator struct {
	Event *FeeManagerFeeUpdaterRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeManagerFeeUpdaterRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerFeeUpdaterRemoved)
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
		it.Event = new(FeeManagerFeeUpdaterRemoved)
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

func (it *FeeManagerFeeUpdaterRemovedIterator) Error() error {
	return it.fail
}

func (it *FeeManagerFeeUpdaterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeManagerFeeUpdaterRemoved struct {
	FeeUpdater common.Address
	Raw        types.Log
}

func (_FeeManager *FeeManagerFilterer) FilterFeeUpdaterRemoved(opts *bind.FilterOpts, feeUpdater []common.Address) (*FeeManagerFeeUpdaterRemovedIterator, error) {

	var feeUpdaterRule []interface{}
	for _, feeUpdaterItem := range feeUpdater {
		feeUpdaterRule = append(feeUpdaterRule, feeUpdaterItem)
	}

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "FeeUpdaterRemoved", feeUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &FeeManagerFeeUpdaterRemovedIterator{contract: _FeeManager.contract, event: "FeeUpdaterRemoved", logs: logs, sub: sub}, nil
}

func (_FeeManager *FeeManagerFilterer) WatchFeeUpdaterRemoved(opts *bind.WatchOpts, sink chan<- *FeeManagerFeeUpdaterRemoved, feeUpdater []common.Address) (event.Subscription, error) {

	var feeUpdaterRule []interface{}
	for _, feeUpdaterItem := range feeUpdater {
		feeUpdaterRule = append(feeUpdaterRule, feeUpdaterItem)
	}

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "FeeUpdaterRemoved", feeUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeManagerFeeUpdaterRemoved)
				if err := _FeeManager.contract.UnpackLog(event, "FeeUpdaterRemoved", log); err != nil {
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

func (_FeeManager *FeeManagerFilterer) ParseFeeUpdaterRemoved(log types.Log) (*FeeManagerFeeUpdaterRemoved, error) {
	event := new(FeeManagerFeeUpdaterRemoved)
	if err := _FeeManager.contract.UnpackLog(event, "FeeUpdaterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeManagerFeeUpdaterSetIterator struct {
	Event *FeeManagerFeeUpdaterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeManagerFeeUpdaterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerFeeUpdaterSet)
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
		it.Event = new(FeeManagerFeeUpdaterSet)
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

func (it *FeeManagerFeeUpdaterSetIterator) Error() error {
	return it.fail
}

func (it *FeeManagerFeeUpdaterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeManagerFeeUpdaterSet struct {
	FeeUpdater common.Address
	Raw        types.Log
}

func (_FeeManager *FeeManagerFilterer) FilterFeeUpdaterSet(opts *bind.FilterOpts, feeUpdater []common.Address) (*FeeManagerFeeUpdaterSetIterator, error) {

	var feeUpdaterRule []interface{}
	for _, feeUpdaterItem := range feeUpdater {
		feeUpdaterRule = append(feeUpdaterRule, feeUpdaterItem)
	}

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "FeeUpdaterSet", feeUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &FeeManagerFeeUpdaterSetIterator{contract: _FeeManager.contract, event: "FeeUpdaterSet", logs: logs, sub: sub}, nil
}

func (_FeeManager *FeeManagerFilterer) WatchFeeUpdaterSet(opts *bind.WatchOpts, sink chan<- *FeeManagerFeeUpdaterSet, feeUpdater []common.Address) (event.Subscription, error) {

	var feeUpdaterRule []interface{}
	for _, feeUpdaterItem := range feeUpdater {
		feeUpdaterRule = append(feeUpdaterRule, feeUpdaterItem)
	}

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "FeeUpdaterSet", feeUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeManagerFeeUpdaterSet)
				if err := _FeeManager.contract.UnpackLog(event, "FeeUpdaterSet", log); err != nil {
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

func (_FeeManager *FeeManagerFilterer) ParseFeeUpdaterSet(log types.Log) (*FeeManagerFeeUpdaterSet, error) {
	event := new(FeeManagerFeeUpdaterSet)
	if err := _FeeManager.contract.UnpackLog(event, "FeeUpdaterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeManagerGasFeeUpdatedIterator struct {
	Event *FeeManagerGasFeeUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeManagerGasFeeUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerGasFeeUpdated)
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
		it.Event = new(FeeManagerGasFeeUpdated)
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

func (it *FeeManagerGasFeeUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeManagerGasFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeManagerGasFeeUpdated struct {
	Token          common.Address
	DestChain      uint64
	LinkPerUnitGas *big.Int
	Timestamp      *big.Int
	Raw            types.Log
}

func (_FeeManager *FeeManagerFilterer) FilterGasFeeUpdated(opts *bind.FilterOpts, token []common.Address, destChain []uint64) (*FeeManagerGasFeeUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var destChainRule []interface{}
	for _, destChainItem := range destChain {
		destChainRule = append(destChainRule, destChainItem)
	}

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "GasFeeUpdated", tokenRule, destChainRule)
	if err != nil {
		return nil, err
	}
	return &FeeManagerGasFeeUpdatedIterator{contract: _FeeManager.contract, event: "GasFeeUpdated", logs: logs, sub: sub}, nil
}

func (_FeeManager *FeeManagerFilterer) WatchGasFeeUpdated(opts *bind.WatchOpts, sink chan<- *FeeManagerGasFeeUpdated, token []common.Address, destChain []uint64) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var destChainRule []interface{}
	for _, destChainItem := range destChain {
		destChainRule = append(destChainRule, destChainItem)
	}

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "GasFeeUpdated", tokenRule, destChainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeManagerGasFeeUpdated)
				if err := _FeeManager.contract.UnpackLog(event, "GasFeeUpdated", log); err != nil {
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

func (_FeeManager *FeeManagerFilterer) ParseGasFeeUpdated(log types.Log) (*FeeManagerGasFeeUpdated, error) {
	event := new(FeeManagerGasFeeUpdated)
	if err := _FeeManager.contract.UnpackLog(event, "GasFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeManagerOwnershipTransferRequestedIterator struct {
	Event *FeeManagerOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeManagerOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerOwnershipTransferRequested)
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
		it.Event = new(FeeManagerOwnershipTransferRequested)
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

func (it *FeeManagerOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *FeeManagerOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeManagerOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FeeManager *FeeManagerFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeManagerOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeeManagerOwnershipTransferRequestedIterator{contract: _FeeManager.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_FeeManager *FeeManagerFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FeeManagerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeManagerOwnershipTransferRequested)
				if err := _FeeManager.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_FeeManager *FeeManagerFilterer) ParseOwnershipTransferRequested(log types.Log) (*FeeManagerOwnershipTransferRequested, error) {
	event := new(FeeManagerOwnershipTransferRequested)
	if err := _FeeManager.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeManagerOwnershipTransferredIterator struct {
	Event *FeeManagerOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeManagerOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerOwnershipTransferred)
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
		it.Event = new(FeeManagerOwnershipTransferred)
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

func (it *FeeManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *FeeManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeManagerOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FeeManager *FeeManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeManagerOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeeManagerOwnershipTransferredIterator{contract: _FeeManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_FeeManager *FeeManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeManagerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeManagerOwnershipTransferred)
				if err := _FeeManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_FeeManager *FeeManagerFilterer) ParseOwnershipTransferred(log types.Log) (*FeeManagerOwnershipTransferred, error) {
	event := new(FeeManagerOwnershipTransferred)
	if err := _FeeManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_FeeManager *FeeManager) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _FeeManager.abi.Events["FeeUpdaterRemoved"].ID:
		return _FeeManager.ParseFeeUpdaterRemoved(log)
	case _FeeManager.abi.Events["FeeUpdaterSet"].ID:
		return _FeeManager.ParseFeeUpdaterSet(log)
	case _FeeManager.abi.Events["GasFeeUpdated"].ID:
		return _FeeManager.ParseGasFeeUpdated(log)
	case _FeeManager.abi.Events["OwnershipTransferRequested"].ID:
		return _FeeManager.ParseOwnershipTransferRequested(log)
	case _FeeManager.abi.Events["OwnershipTransferred"].ID:
		return _FeeManager.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (FeeManagerFeeUpdaterRemoved) Topic() common.Hash {
	return common.HexToHash("0x74a2c31badb27f0acfb9da3ef34c9e656ca1723881466e89a40f791f1c82ee71")
}

func (FeeManagerFeeUpdaterSet) Topic() common.Hash {
	return common.HexToHash("0xa462a3423511fce5ad038a1130251b66e3b8c135fa2ca15777f45c72ea3037c5")
}

func (FeeManagerGasFeeUpdated) Topic() common.Hash {
	return common.HexToHash("0xe45338fd766b2619bbcd30dd0c79dcc00795e7977bab27952f0e78571a100764")
}

func (FeeManagerOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (FeeManagerOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_FeeManager *FeeManager) Address() common.Address {
	return _FeeManager.address
}

type FeeManagerInterface interface {
	GetFee(opts *bind.CallOpts, token common.Address, destChainId uint64) (*big.Int, error)

	GetStalenessThreshold(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error)

	SetFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateFees(opts *bind.TransactOpts, feeUpdates []GEFeeUpdate) (*types.Transaction, error)

	WithdrawToken(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	FilterFeeUpdaterRemoved(opts *bind.FilterOpts, feeUpdater []common.Address) (*FeeManagerFeeUpdaterRemovedIterator, error)

	WatchFeeUpdaterRemoved(opts *bind.WatchOpts, sink chan<- *FeeManagerFeeUpdaterRemoved, feeUpdater []common.Address) (event.Subscription, error)

	ParseFeeUpdaterRemoved(log types.Log) (*FeeManagerFeeUpdaterRemoved, error)

	FilterFeeUpdaterSet(opts *bind.FilterOpts, feeUpdater []common.Address) (*FeeManagerFeeUpdaterSetIterator, error)

	WatchFeeUpdaterSet(opts *bind.WatchOpts, sink chan<- *FeeManagerFeeUpdaterSet, feeUpdater []common.Address) (event.Subscription, error)

	ParseFeeUpdaterSet(log types.Log) (*FeeManagerFeeUpdaterSet, error)

	FilterGasFeeUpdated(opts *bind.FilterOpts, token []common.Address, destChain []uint64) (*FeeManagerGasFeeUpdatedIterator, error)

	WatchGasFeeUpdated(opts *bind.WatchOpts, sink chan<- *FeeManagerGasFeeUpdated, token []common.Address, destChain []uint64) (event.Subscription, error)

	ParseGasFeeUpdated(log types.Log) (*FeeManagerGasFeeUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeManagerOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FeeManagerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*FeeManagerOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeManagerOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeManagerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*FeeManagerOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
