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

var GasFeeCacheMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"linkPerUnitGas\",\"type\":\"uint256\"}],\"internalType\":\"structGE.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"feeUpdaters\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chain\",\"type\":\"uint64\"}],\"name\":\"ChainNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByUpdaterOrOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"destChain\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linkPerUnitGas\",\"type\":\"uint256\"}],\"name\":\"GasFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"removeFeeUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"setFeeUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"linkPerUnitGas\",\"type\":\"uint256\"}],\"internalType\":\"structGE.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"}],\"name\":\"updateFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000ed038038062000ed08339810160408190526200003491620003f9565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be816200022a565b50505060005b82518110156200014357828181518110620000e357620000e3620004f7565b60200260200101516020015160026000858481518110620001085762000108620004f7565b6020026020010151600001516001600160401b0316815260200190815260200160002081905550806200013b906200050d565b9050620000c4565b5060005b815181101562000221576001600360008484815181106200016c576200016c620004f7565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff0219169083151502179055507fa462a3423511fce5ad038a1130251b66e3b8c135fa2ca15777f45c72ea3037c5828281518110620001e157620001e1620004f7565b60200260200101516040516200020691906001600160a01b0391909116815260200190565b60405180910390a162000219816200050d565b905062000147565b50505062000535565b336001600160a01b03821603620002845760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715620003105762000310620002d5565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620003415762000341620002d5565b604052919050565b60006001600160401b03821115620003655762000365620002d5565b5060051b60200190565b600082601f8301126200038157600080fd5b815160206200039a620003948362000349565b62000316565b82815260059290921b84018101918181019086841115620003ba57600080fd5b8286015b84811015620003ee5780516001600160a01b0381168114620003e05760008081fd5b8352918301918301620003be565b509695505050505050565b60008060408084860312156200040e57600080fd5b83516001600160401b03808211156200042657600080fd5b818601915086601f8301126200043b57600080fd5b815160206200044e620003948362000349565b82815260069290921b8401810191818101908a8411156200046e57600080fd5b948201945b83861015620004c35786868c0312156200048d5760008081fd5b62000497620002eb565b86518681168114620004a95760008081fd5b815286840151848201528252948601949082019062000473565b91890151919750909450505080831115620004dd57600080fd5b5050620004ed858286016200036f565b9150509250929050565b634e487b7160e01b600052603260045260246000fd5b6000600182016200052e57634e487b7160e01b600052601160045260246000fd5b5060010190565b61098b80620005456000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806379ba50971161005b57806379ba5097146100d05780638da5cb5b146100d8578063ae7fca1814610100578063f2fde38b1461011357600080fd5b80631982b1d014610082578063604782e6146100a85780636e42efb8146100bd575b600080fd5b61009561009036600461071d565b610126565b6040519081526020015b60405180910390f35b6100bb6100b636600461073f565b610190565b005b6100bb6100cb36600461081c565b61023b565b6100bb6103fb565b60005460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161009f565b6100bb61010e36600461073f565b6104f8565b6100bb61012136600461073f565b61057c565b67ffffffffffffffff81166000908152600260205260408120549081900361018b576040517f2e59db3a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff831660048201526024015b60405180910390fd5b919050565b61019861058d565b73ffffffffffffffffffffffffffffffffffffffff8116156102385773ffffffffffffffffffffffffffffffffffffffff811660008181526003602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590519182527fa462a3423511fce5ad038a1130251b66e3b8c135fa2ca15777f45c72ea3037c591015b60405180910390a15b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331480159061027357503360009081526003602052604090205460ff16155b156102aa576040517f46f0815400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526003602052604090205460ff166102f5576040517faf0026bf000000000000000000000000000000000000000000000000000000008152336004820152602401610182565b805160005b818110156103f657828181518110610314576103146108f0565b60200260200101516020015160026000858481518110610336576103366108f0565b60200260200101516000015167ffffffffffffffff168152602001908152602001600020819055507fe4a6427b8d1dbd7bde164395b80d8c8d04264cc8f303268b79eba905b7218350838281518110610391576103916108f0565b6020026020010151600001518483815181106103af576103af6108f0565b6020026020010151602001516040516103de92919067ffffffffffffffff929092168252602082015260400190565b60405180910390a16103ef8161091f565b90506102fa565b505050565b60015473ffffffffffffffffffffffffffffffffffffffff16331461047c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610182565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61050061058d565b73ffffffffffffffffffffffffffffffffffffffff811660008181526003602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905590519182527f74a2c31badb27f0acfb9da3ef34c9e656ca1723881466e89a40f791f1c82ee71910161022f565b61058461058d565b61023881610610565b60005473ffffffffffffffffffffffffffffffffffffffff16331461060e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610182565b565b3373ffffffffffffffffffffffffffffffffffffffff82160361068f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610182565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b803567ffffffffffffffff8116811461018b57600080fd5b60006020828403121561072f57600080fd5b61073882610705565b9392505050565b60006020828403121561075157600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461073857600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156107c7576107c7610775565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561081457610814610775565b604052919050565b6000602080838503121561082f57600080fd5b823567ffffffffffffffff8082111561084757600080fd5b818501915085601f83011261085b57600080fd5b81358181111561086d5761086d610775565b61087b848260051b016107cd565b818152848101925060069190911b83018401908782111561089b57600080fd5b928401925b818410156108e557604084890312156108b95760008081fd5b6108c16107a4565b6108ca85610705565b815284860135868201528352604090930192918401916108a0565b979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610977577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c634300080f000a",
}

var GasFeeCacheABI = GasFeeCacheMetaData.ABI

var GasFeeCacheBin = GasFeeCacheMetaData.Bin

func DeployGasFeeCache(auth *bind.TransactOpts, backend bind.ContractBackend, feeUpdates []GEFeeUpdate, feeUpdaters []common.Address) (common.Address, *types.Transaction, *GasFeeCache, error) {
	parsed, err := GasFeeCacheMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GasFeeCacheBin), backend, feeUpdates, feeUpdaters)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasFeeCache{GasFeeCacheCaller: GasFeeCacheCaller{contract: contract}, GasFeeCacheTransactor: GasFeeCacheTransactor{contract: contract}, GasFeeCacheFilterer: GasFeeCacheFilterer{contract: contract}}, nil
}

type GasFeeCache struct {
	address common.Address
	abi     abi.ABI
	GasFeeCacheCaller
	GasFeeCacheTransactor
	GasFeeCacheFilterer
}

type GasFeeCacheCaller struct {
	contract *bind.BoundContract
}

type GasFeeCacheTransactor struct {
	contract *bind.BoundContract
}

type GasFeeCacheFilterer struct {
	contract *bind.BoundContract
}

type GasFeeCacheSession struct {
	Contract     *GasFeeCache
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type GasFeeCacheCallerSession struct {
	Contract *GasFeeCacheCaller
	CallOpts bind.CallOpts
}

type GasFeeCacheTransactorSession struct {
	Contract     *GasFeeCacheTransactor
	TransactOpts bind.TransactOpts
}

type GasFeeCacheRaw struct {
	Contract *GasFeeCache
}

type GasFeeCacheCallerRaw struct {
	Contract *GasFeeCacheCaller
}

type GasFeeCacheTransactorRaw struct {
	Contract *GasFeeCacheTransactor
}

func NewGasFeeCache(address common.Address, backend bind.ContractBackend) (*GasFeeCache, error) {
	abi, err := abi.JSON(strings.NewReader(GasFeeCacheABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindGasFeeCache(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasFeeCache{address: address, abi: abi, GasFeeCacheCaller: GasFeeCacheCaller{contract: contract}, GasFeeCacheTransactor: GasFeeCacheTransactor{contract: contract}, GasFeeCacheFilterer: GasFeeCacheFilterer{contract: contract}}, nil
}

func NewGasFeeCacheCaller(address common.Address, caller bind.ContractCaller) (*GasFeeCacheCaller, error) {
	contract, err := bindGasFeeCache(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasFeeCacheCaller{contract: contract}, nil
}

func NewGasFeeCacheTransactor(address common.Address, transactor bind.ContractTransactor) (*GasFeeCacheTransactor, error) {
	contract, err := bindGasFeeCache(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasFeeCacheTransactor{contract: contract}, nil
}

func NewGasFeeCacheFilterer(address common.Address, filterer bind.ContractFilterer) (*GasFeeCacheFilterer, error) {
	contract, err := bindGasFeeCache(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasFeeCacheFilterer{contract: contract}, nil
}

func bindGasFeeCache(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GasFeeCacheABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_GasFeeCache *GasFeeCacheRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasFeeCache.Contract.GasFeeCacheCaller.contract.Call(opts, result, method, params...)
}

func (_GasFeeCache *GasFeeCacheRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasFeeCache.Contract.GasFeeCacheTransactor.contract.Transfer(opts)
}

func (_GasFeeCache *GasFeeCacheRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasFeeCache.Contract.GasFeeCacheTransactor.contract.Transact(opts, method, params...)
}

func (_GasFeeCache *GasFeeCacheCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasFeeCache.Contract.contract.Call(opts, result, method, params...)
}

func (_GasFeeCache *GasFeeCacheTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasFeeCache.Contract.contract.Transfer(opts)
}

func (_GasFeeCache *GasFeeCacheTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasFeeCache.Contract.contract.Transact(opts, method, params...)
}

func (_GasFeeCache *GasFeeCacheCaller) GetFee(opts *bind.CallOpts, destChainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _GasFeeCache.contract.Call(opts, &out, "getFee", destChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_GasFeeCache *GasFeeCacheSession) GetFee(destChainId uint64) (*big.Int, error) {
	return _GasFeeCache.Contract.GetFee(&_GasFeeCache.CallOpts, destChainId)
}

func (_GasFeeCache *GasFeeCacheCallerSession) GetFee(destChainId uint64) (*big.Int, error) {
	return _GasFeeCache.Contract.GetFee(&_GasFeeCache.CallOpts, destChainId)
}

func (_GasFeeCache *GasFeeCacheCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasFeeCache.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_GasFeeCache *GasFeeCacheSession) Owner() (common.Address, error) {
	return _GasFeeCache.Contract.Owner(&_GasFeeCache.CallOpts)
}

func (_GasFeeCache *GasFeeCacheCallerSession) Owner() (common.Address, error) {
	return _GasFeeCache.Contract.Owner(&_GasFeeCache.CallOpts)
}

func (_GasFeeCache *GasFeeCacheTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasFeeCache.contract.Transact(opts, "acceptOwnership")
}

func (_GasFeeCache *GasFeeCacheSession) AcceptOwnership() (*types.Transaction, error) {
	return _GasFeeCache.Contract.AcceptOwnership(&_GasFeeCache.TransactOpts)
}

func (_GasFeeCache *GasFeeCacheTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _GasFeeCache.Contract.AcceptOwnership(&_GasFeeCache.TransactOpts)
}

func (_GasFeeCache *GasFeeCacheTransactor) RemoveFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error) {
	return _GasFeeCache.contract.Transact(opts, "removeFeeUpdater", feeUpdater)
}

func (_GasFeeCache *GasFeeCacheSession) RemoveFeeUpdater(feeUpdater common.Address) (*types.Transaction, error) {
	return _GasFeeCache.Contract.RemoveFeeUpdater(&_GasFeeCache.TransactOpts, feeUpdater)
}

func (_GasFeeCache *GasFeeCacheTransactorSession) RemoveFeeUpdater(feeUpdater common.Address) (*types.Transaction, error) {
	return _GasFeeCache.Contract.RemoveFeeUpdater(&_GasFeeCache.TransactOpts, feeUpdater)
}

func (_GasFeeCache *GasFeeCacheTransactor) SetFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error) {
	return _GasFeeCache.contract.Transact(opts, "setFeeUpdater", feeUpdater)
}

func (_GasFeeCache *GasFeeCacheSession) SetFeeUpdater(feeUpdater common.Address) (*types.Transaction, error) {
	return _GasFeeCache.Contract.SetFeeUpdater(&_GasFeeCache.TransactOpts, feeUpdater)
}

func (_GasFeeCache *GasFeeCacheTransactorSession) SetFeeUpdater(feeUpdater common.Address) (*types.Transaction, error) {
	return _GasFeeCache.Contract.SetFeeUpdater(&_GasFeeCache.TransactOpts, feeUpdater)
}

func (_GasFeeCache *GasFeeCacheTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _GasFeeCache.contract.Transact(opts, "transferOwnership", to)
}

func (_GasFeeCache *GasFeeCacheSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _GasFeeCache.Contract.TransferOwnership(&_GasFeeCache.TransactOpts, to)
}

func (_GasFeeCache *GasFeeCacheTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _GasFeeCache.Contract.TransferOwnership(&_GasFeeCache.TransactOpts, to)
}

func (_GasFeeCache *GasFeeCacheTransactor) UpdateFees(opts *bind.TransactOpts, feeUpdates []GEFeeUpdate) (*types.Transaction, error) {
	return _GasFeeCache.contract.Transact(opts, "updateFees", feeUpdates)
}

func (_GasFeeCache *GasFeeCacheSession) UpdateFees(feeUpdates []GEFeeUpdate) (*types.Transaction, error) {
	return _GasFeeCache.Contract.UpdateFees(&_GasFeeCache.TransactOpts, feeUpdates)
}

func (_GasFeeCache *GasFeeCacheTransactorSession) UpdateFees(feeUpdates []GEFeeUpdate) (*types.Transaction, error) {
	return _GasFeeCache.Contract.UpdateFees(&_GasFeeCache.TransactOpts, feeUpdates)
}

type GasFeeCacheFeeUpdaterRemovedIterator struct {
	Event *GasFeeCacheFeeUpdaterRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GasFeeCacheFeeUpdaterRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasFeeCacheFeeUpdaterRemoved)
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
		it.Event = new(GasFeeCacheFeeUpdaterRemoved)
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

func (it *GasFeeCacheFeeUpdaterRemovedIterator) Error() error {
	return it.fail
}

func (it *GasFeeCacheFeeUpdaterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GasFeeCacheFeeUpdaterRemoved struct {
	FeeUpdater common.Address
	Raw        types.Log
}

func (_GasFeeCache *GasFeeCacheFilterer) FilterFeeUpdaterRemoved(opts *bind.FilterOpts) (*GasFeeCacheFeeUpdaterRemovedIterator, error) {

	logs, sub, err := _GasFeeCache.contract.FilterLogs(opts, "FeeUpdaterRemoved")
	if err != nil {
		return nil, err
	}
	return &GasFeeCacheFeeUpdaterRemovedIterator{contract: _GasFeeCache.contract, event: "FeeUpdaterRemoved", logs: logs, sub: sub}, nil
}

func (_GasFeeCache *GasFeeCacheFilterer) WatchFeeUpdaterRemoved(opts *bind.WatchOpts, sink chan<- *GasFeeCacheFeeUpdaterRemoved) (event.Subscription, error) {

	logs, sub, err := _GasFeeCache.contract.WatchLogs(opts, "FeeUpdaterRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GasFeeCacheFeeUpdaterRemoved)
				if err := _GasFeeCache.contract.UnpackLog(event, "FeeUpdaterRemoved", log); err != nil {
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

func (_GasFeeCache *GasFeeCacheFilterer) ParseFeeUpdaterRemoved(log types.Log) (*GasFeeCacheFeeUpdaterRemoved, error) {
	event := new(GasFeeCacheFeeUpdaterRemoved)
	if err := _GasFeeCache.contract.UnpackLog(event, "FeeUpdaterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GasFeeCacheFeeUpdaterSetIterator struct {
	Event *GasFeeCacheFeeUpdaterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GasFeeCacheFeeUpdaterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasFeeCacheFeeUpdaterSet)
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
		it.Event = new(GasFeeCacheFeeUpdaterSet)
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

func (it *GasFeeCacheFeeUpdaterSetIterator) Error() error {
	return it.fail
}

func (it *GasFeeCacheFeeUpdaterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GasFeeCacheFeeUpdaterSet struct {
	FeeUpdater common.Address
	Raw        types.Log
}

func (_GasFeeCache *GasFeeCacheFilterer) FilterFeeUpdaterSet(opts *bind.FilterOpts) (*GasFeeCacheFeeUpdaterSetIterator, error) {

	logs, sub, err := _GasFeeCache.contract.FilterLogs(opts, "FeeUpdaterSet")
	if err != nil {
		return nil, err
	}
	return &GasFeeCacheFeeUpdaterSetIterator{contract: _GasFeeCache.contract, event: "FeeUpdaterSet", logs: logs, sub: sub}, nil
}

func (_GasFeeCache *GasFeeCacheFilterer) WatchFeeUpdaterSet(opts *bind.WatchOpts, sink chan<- *GasFeeCacheFeeUpdaterSet) (event.Subscription, error) {

	logs, sub, err := _GasFeeCache.contract.WatchLogs(opts, "FeeUpdaterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GasFeeCacheFeeUpdaterSet)
				if err := _GasFeeCache.contract.UnpackLog(event, "FeeUpdaterSet", log); err != nil {
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

func (_GasFeeCache *GasFeeCacheFilterer) ParseFeeUpdaterSet(log types.Log) (*GasFeeCacheFeeUpdaterSet, error) {
	event := new(GasFeeCacheFeeUpdaterSet)
	if err := _GasFeeCache.contract.UnpackLog(event, "FeeUpdaterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GasFeeCacheGasFeeUpdatedIterator struct {
	Event *GasFeeCacheGasFeeUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GasFeeCacheGasFeeUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasFeeCacheGasFeeUpdated)
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
		it.Event = new(GasFeeCacheGasFeeUpdated)
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

func (it *GasFeeCacheGasFeeUpdatedIterator) Error() error {
	return it.fail
}

func (it *GasFeeCacheGasFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GasFeeCacheGasFeeUpdated struct {
	DestChain      uint64
	LinkPerUnitGas *big.Int
	Raw            types.Log
}

func (_GasFeeCache *GasFeeCacheFilterer) FilterGasFeeUpdated(opts *bind.FilterOpts) (*GasFeeCacheGasFeeUpdatedIterator, error) {

	logs, sub, err := _GasFeeCache.contract.FilterLogs(opts, "GasFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &GasFeeCacheGasFeeUpdatedIterator{contract: _GasFeeCache.contract, event: "GasFeeUpdated", logs: logs, sub: sub}, nil
}

func (_GasFeeCache *GasFeeCacheFilterer) WatchGasFeeUpdated(opts *bind.WatchOpts, sink chan<- *GasFeeCacheGasFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _GasFeeCache.contract.WatchLogs(opts, "GasFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GasFeeCacheGasFeeUpdated)
				if err := _GasFeeCache.contract.UnpackLog(event, "GasFeeUpdated", log); err != nil {
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

func (_GasFeeCache *GasFeeCacheFilterer) ParseGasFeeUpdated(log types.Log) (*GasFeeCacheGasFeeUpdated, error) {
	event := new(GasFeeCacheGasFeeUpdated)
	if err := _GasFeeCache.contract.UnpackLog(event, "GasFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GasFeeCacheOwnershipTransferRequestedIterator struct {
	Event *GasFeeCacheOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GasFeeCacheOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasFeeCacheOwnershipTransferRequested)
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
		it.Event = new(GasFeeCacheOwnershipTransferRequested)
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

func (it *GasFeeCacheOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *GasFeeCacheOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GasFeeCacheOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_GasFeeCache *GasFeeCacheFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GasFeeCacheOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GasFeeCache.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GasFeeCacheOwnershipTransferRequestedIterator{contract: _GasFeeCache.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_GasFeeCache *GasFeeCacheFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *GasFeeCacheOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GasFeeCache.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GasFeeCacheOwnershipTransferRequested)
				if err := _GasFeeCache.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_GasFeeCache *GasFeeCacheFilterer) ParseOwnershipTransferRequested(log types.Log) (*GasFeeCacheOwnershipTransferRequested, error) {
	event := new(GasFeeCacheOwnershipTransferRequested)
	if err := _GasFeeCache.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GasFeeCacheOwnershipTransferredIterator struct {
	Event *GasFeeCacheOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GasFeeCacheOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasFeeCacheOwnershipTransferred)
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
		it.Event = new(GasFeeCacheOwnershipTransferred)
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

func (it *GasFeeCacheOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *GasFeeCacheOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GasFeeCacheOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_GasFeeCache *GasFeeCacheFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GasFeeCacheOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GasFeeCache.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GasFeeCacheOwnershipTransferredIterator{contract: _GasFeeCache.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_GasFeeCache *GasFeeCacheFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GasFeeCacheOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GasFeeCache.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GasFeeCacheOwnershipTransferred)
				if err := _GasFeeCache.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_GasFeeCache *GasFeeCacheFilterer) ParseOwnershipTransferred(log types.Log) (*GasFeeCacheOwnershipTransferred, error) {
	event := new(GasFeeCacheOwnershipTransferred)
	if err := _GasFeeCache.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_GasFeeCache *GasFeeCache) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _GasFeeCache.abi.Events["FeeUpdaterRemoved"].ID:
		return _GasFeeCache.ParseFeeUpdaterRemoved(log)
	case _GasFeeCache.abi.Events["FeeUpdaterSet"].ID:
		return _GasFeeCache.ParseFeeUpdaterSet(log)
	case _GasFeeCache.abi.Events["GasFeeUpdated"].ID:
		return _GasFeeCache.ParseGasFeeUpdated(log)
	case _GasFeeCache.abi.Events["OwnershipTransferRequested"].ID:
		return _GasFeeCache.ParseOwnershipTransferRequested(log)
	case _GasFeeCache.abi.Events["OwnershipTransferred"].ID:
		return _GasFeeCache.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (GasFeeCacheFeeUpdaterRemoved) Topic() common.Hash {
	return common.HexToHash("0x74a2c31badb27f0acfb9da3ef34c9e656ca1723881466e89a40f791f1c82ee71")
}

func (GasFeeCacheFeeUpdaterSet) Topic() common.Hash {
	return common.HexToHash("0xa462a3423511fce5ad038a1130251b66e3b8c135fa2ca15777f45c72ea3037c5")
}

func (GasFeeCacheGasFeeUpdated) Topic() common.Hash {
	return common.HexToHash("0xe4a6427b8d1dbd7bde164395b80d8c8d04264cc8f303268b79eba905b7218350")
}

func (GasFeeCacheOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (GasFeeCacheOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_GasFeeCache *GasFeeCache) Address() common.Address {
	return _GasFeeCache.address
}

type GasFeeCacheInterface interface {
	GetFee(opts *bind.CallOpts, destChainId uint64) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error)

	SetFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateFees(opts *bind.TransactOpts, feeUpdates []GEFeeUpdate) (*types.Transaction, error)

	FilterFeeUpdaterRemoved(opts *bind.FilterOpts) (*GasFeeCacheFeeUpdaterRemovedIterator, error)

	WatchFeeUpdaterRemoved(opts *bind.WatchOpts, sink chan<- *GasFeeCacheFeeUpdaterRemoved) (event.Subscription, error)

	ParseFeeUpdaterRemoved(log types.Log) (*GasFeeCacheFeeUpdaterRemoved, error)

	FilterFeeUpdaterSet(opts *bind.FilterOpts) (*GasFeeCacheFeeUpdaterSetIterator, error)

	WatchFeeUpdaterSet(opts *bind.WatchOpts, sink chan<- *GasFeeCacheFeeUpdaterSet) (event.Subscription, error)

	ParseFeeUpdaterSet(log types.Log) (*GasFeeCacheFeeUpdaterSet, error)

	FilterGasFeeUpdated(opts *bind.FilterOpts) (*GasFeeCacheGasFeeUpdatedIterator, error)

	WatchGasFeeUpdated(opts *bind.WatchOpts, sink chan<- *GasFeeCacheGasFeeUpdated) (event.Subscription, error)

	ParseGasFeeUpdated(log types.Log) (*GasFeeCacheGasFeeUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GasFeeCacheOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *GasFeeCacheOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*GasFeeCacheOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GasFeeCacheOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GasFeeCacheOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*GasFeeCacheOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
