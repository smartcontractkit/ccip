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
	ChainId        uint64
	LinkPerUnitGas *big.Int
}

var FeeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"linkPerUnitGas\",\"type\":\"uint128\"}],\"internalType\":\"structGE.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"feeUpdaters\",\"type\":\"address[]\"},{\"internalType\":\"uint128\",\"name\":\"stalenessThreshold\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chain\",\"type\":\"uint64\"}],\"name\":\"ChainNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByUpdaterOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePassed\",\"type\":\"uint256\"}],\"name\":\"StaleFee\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"destChain\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"linkPerUnitGas\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"name\":\"GasFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"fee\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStalenessThreshold\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"removeFeeUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"setFeeUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"linkPerUnitGas\",\"type\":\"uint128\"}],\"internalType\":\"structGE.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"}],\"name\":\"updateFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620010a5380380620010a583398101604081905262000034916200047a565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be8162000198565b50505060005b8351811015620001345762000121848281518110620000e757620000e762000595565b60200260200101516000015185838151811062000108576200010862000595565b602002602001015160200151426200024360201b60201c565b6200012c81620005ab565b9050620000c4565b5060005b825181101562000182576200016f8382815181106200015b576200015b62000595565b6020026020010151620002ce60201b60201c565b6200017a81620005ab565b905062000138565b506001600160801b031660805250620005d39050565b336001600160a01b03821603620001f25760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040805180820182526001600160801b0384811680835284821660208085018281526001600160401b038a16600081815260028452889020965191518616600160801b029190951617909455845192835292820152918201527f3c3ba0a429c06596b6598f026ff79b9c782cdbb3287276341ccc3221909d556a9060600160405180910390a1505050565b6001600160a01b0381161562000336576001600160a01b038116600081815260036020908152604091829020805460ff1916600117905590519182527fa462a3423511fce5ad038a1130251b66e3b8c135fa2ca15777f45c72ea3037c5910160405180910390a15b50565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171562000374576200037462000339565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620003a557620003a562000339565b604052919050565b60006001600160401b03821115620003c957620003c962000339565b5060051b60200190565b80516001600160801b0381168114620003eb57600080fd5b919050565b600082601f8301126200040257600080fd5b815160206200041b6200041583620003ad565b6200037a565b82815260059290921b840181019181810190868411156200043b57600080fd5b8286015b848110156200046f5780516001600160a01b0381168114620004615760008081fd5b83529183019183016200043f565b509695505050505050565b6000806000606084860312156200049057600080fd5b83516001600160401b0380821115620004a857600080fd5b818601915086601f830112620004bd57600080fd5b81516020620004d06200041583620003ad565b82815260069290921b8401810191818101908a841115620004f057600080fd5b948201945b8386101562000552576040868c031215620005105760008081fd5b6200051a6200034f565b865186811681146200052c5760008081fd5b81526200053b878501620003d3565b8185015282526040959095019490820190620004f5565b918901519197509093505050808211156200056c57600080fd5b506200057b86828701620003f0565b9250506200058c60408501620003d3565b90509250925092565b634e487b7160e01b600052603260045260246000fd5b600060018201620005cc57634e487b7160e01b600052601160045260246000fd5b5060010190565b608051610aaf620005f66000396000818161012001526101c60152610aaf6000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638da5cb5b1161005b5780638da5cb5b146100f6578063a6c94a731461011e578063ae7fca1814610144578063f2fde38b1461015757600080fd5b80631982b1d01461008d57806337170831146100c6578063604782e6146100db57806379ba5097146100ee575b600080fd5b6100a061009b366004610802565b61016a565b6040516fffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100d96100d43660046108cb565b6102a4565b005b6100d96100e93660046109bf565b61037a565b6100d961038e565b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100bd565b7f00000000000000000000000000000000000000000000000000000000000000006100a0565b6100d96101523660046109bf565b61048b565b6100d96101653660046109bf565b61049c565b67ffffffffffffffff811660009081526002602090815260408083208151808301909252546fffffffffffffffffffffffffffffffff8082168352700100000000000000000000000000000000909104811692820183905290917f00000000000000000000000000000000000000000000000000000000000000009091169083906101f59042610a24565b905081811115610240576040517f1f49f5b800000000000000000000000000000000000000000000000000000000815260048101839052602481018290526044015b60405180910390fd5b825193506fffffffffffffffffffffffffffffffff841660000361029c576040517f2e59db3a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610237565b505050919050565b60005473ffffffffffffffffffffffffffffffffffffffff1633148015906102dc57503360009081526003602052604090205460ff16155b15610313576040517f46f0815400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4260005b82518110156103755761036583828151811061033557610335610a3b565b60200260200101516000015184838151811061035357610353610a3b565b602002602001015160200151846104ad565b61036e81610a6a565b9050610317565b505050565b61038261054f565b61038b816105d2565b50565b60015473ffffffffffffffffffffffffffffffffffffffff16331461040f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610237565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61049361054f565b61038b81610674565b6104a461054f565b61038b816106f0565b6040805180820182526fffffffffffffffffffffffffffffffff848116808352848216602080850182815267ffffffffffffffff8a16600081815260028452889020965191518616700100000000000000000000000000000000029190951617909455845192835292820152918201527f3c3ba0a429c06596b6598f026ff79b9c782cdbb3287276341ccc3221909d556a9060600160405180910390a1505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146105d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610237565b565b73ffffffffffffffffffffffffffffffffffffffff81161561038b5773ffffffffffffffffffffffffffffffffffffffff811660008181526003602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590519182527fa462a3423511fce5ad038a1130251b66e3b8c135fa2ca15777f45c72ea3037c591015b60405180910390a150565b73ffffffffffffffffffffffffffffffffffffffff811660008181526003602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905590519182527f74a2c31badb27f0acfb9da3ef34c9e656ca1723881466e89a40f791f1c82ee719101610669565b3373ffffffffffffffffffffffffffffffffffffffff82160361076f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610237565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b803567ffffffffffffffff811681146107fd57600080fd5b919050565b60006020828403121561081457600080fd5b61081d826107e5565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561087657610876610824565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156108c3576108c3610824565b604052919050565b600060208083850312156108de57600080fd5b823567ffffffffffffffff808211156108f657600080fd5b818501915085601f83011261090a57600080fd5b81358181111561091c5761091c610824565b61092a848260051b0161087c565b818152848101925060069190911b83018401908782111561094a57600080fd5b928401925b818410156109b457604084890312156109685760008081fd5b610970610853565b610979856107e5565b8152858501356fffffffffffffffffffffffffffffffff8116811461099e5760008081fd5b818701528352604093909301929184019161094f565b979650505050505050565b6000602082840312156109d157600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461081d57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015610a3657610a366109f5565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610a9b57610a9b6109f5565b506001019056fea164736f6c634300080f000a",
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

func (_FeeManager *FeeManagerCaller) GetFee(opts *bind.CallOpts, destChainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "getFee", destChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeManager *FeeManagerSession) GetFee(destChainId uint64) (*big.Int, error) {
	return _FeeManager.Contract.GetFee(&_FeeManager.CallOpts, destChainId)
}

func (_FeeManager *FeeManagerCallerSession) GetFee(destChainId uint64) (*big.Int, error) {
	return _FeeManager.Contract.GetFee(&_FeeManager.CallOpts, destChainId)
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

func (_FeeManager *FeeManagerFilterer) FilterFeeUpdaterRemoved(opts *bind.FilterOpts) (*FeeManagerFeeUpdaterRemovedIterator, error) {

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "FeeUpdaterRemoved")
	if err != nil {
		return nil, err
	}
	return &FeeManagerFeeUpdaterRemovedIterator{contract: _FeeManager.contract, event: "FeeUpdaterRemoved", logs: logs, sub: sub}, nil
}

func (_FeeManager *FeeManagerFilterer) WatchFeeUpdaterRemoved(opts *bind.WatchOpts, sink chan<- *FeeManagerFeeUpdaterRemoved) (event.Subscription, error) {

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "FeeUpdaterRemoved")
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

func (_FeeManager *FeeManagerFilterer) FilterFeeUpdaterSet(opts *bind.FilterOpts) (*FeeManagerFeeUpdaterSetIterator, error) {

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "FeeUpdaterSet")
	if err != nil {
		return nil, err
	}
	return &FeeManagerFeeUpdaterSetIterator{contract: _FeeManager.contract, event: "FeeUpdaterSet", logs: logs, sub: sub}, nil
}

func (_FeeManager *FeeManagerFilterer) WatchFeeUpdaterSet(opts *bind.WatchOpts, sink chan<- *FeeManagerFeeUpdaterSet) (event.Subscription, error) {

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "FeeUpdaterSet")
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
	DestChain      uint64
	LinkPerUnitGas *big.Int
	Timestamp      *big.Int
	Raw            types.Log
}

func (_FeeManager *FeeManagerFilterer) FilterGasFeeUpdated(opts *bind.FilterOpts) (*FeeManagerGasFeeUpdatedIterator, error) {

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "GasFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &FeeManagerGasFeeUpdatedIterator{contract: _FeeManager.contract, event: "GasFeeUpdated", logs: logs, sub: sub}, nil
}

func (_FeeManager *FeeManagerFilterer) WatchGasFeeUpdated(opts *bind.WatchOpts, sink chan<- *FeeManagerGasFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "GasFeeUpdated")
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
	return common.HexToHash("0x3c3ba0a429c06596b6598f026ff79b9c782cdbb3287276341ccc3221909d556a")
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
	GetFee(opts *bind.CallOpts, destChainId uint64) (*big.Int, error)

	GetStalenessThreshold(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error)

	SetFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateFees(opts *bind.TransactOpts, feeUpdates []GEFeeUpdate) (*types.Transaction, error)

	FilterFeeUpdaterRemoved(opts *bind.FilterOpts) (*FeeManagerFeeUpdaterRemovedIterator, error)

	WatchFeeUpdaterRemoved(opts *bind.WatchOpts, sink chan<- *FeeManagerFeeUpdaterRemoved) (event.Subscription, error)

	ParseFeeUpdaterRemoved(log types.Log) (*FeeManagerFeeUpdaterRemoved, error)

	FilterFeeUpdaterSet(opts *bind.FilterOpts) (*FeeManagerFeeUpdaterSetIterator, error)

	WatchFeeUpdaterSet(opts *bind.WatchOpts, sink chan<- *FeeManagerFeeUpdaterSet) (event.Subscription, error)

	ParseFeeUpdaterSet(log types.Log) (*FeeManagerFeeUpdaterSet, error)

	FilterGasFeeUpdated(opts *bind.FilterOpts) (*FeeManagerGasFeeUpdatedIterator, error)

	WatchGasFeeUpdated(opts *bind.WatchOpts, sink chan<- *FeeManagerGasFeeUpdated) (event.Subscription, error)

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
