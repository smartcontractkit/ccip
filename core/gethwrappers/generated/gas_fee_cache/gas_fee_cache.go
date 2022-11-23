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

type CCIPFeeUpdate struct {
	ChainId  *big.Int
	GasPrice *big.Int
}

var GasFeeCacheMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"feeUpdaters\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chain\",\"type\":\"uint256\"}],\"name\":\"ChainNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByUpdaterOrOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"FeeUpdaterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linkPerUnitGas\",\"type\":\"uint256\"}],\"name\":\"GasFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"removeFeeUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeUpdater\",\"type\":\"address\"}],\"name\":\"setFeeUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"}],\"name\":\"updateFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000e6c38038062000e6c8339810160408190526200003491620003f0565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be8162000221565b50505060005b82518110156200013a57828181518110620000e357620000e3620004de565b60200260200101516020015160026000858481518110620001085762000108620004de565b602002602001015160000151815260200190815260200160002081905550806200013290620004f4565b9050620000c4565b5060005b81518110156200021857600160036000848481518110620001635762000163620004de565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff0219169083151502179055507fa462a3423511fce5ad038a1130251b66e3b8c135fa2ca15777f45c72ea3037c5828281518110620001d857620001d8620004de565b6020026020010151604051620001fd91906001600160a01b0391909116815260200190565b60405180910390a16200021081620004f4565b90506200013e565b5050506200051c565b336001600160a01b038216036200027b5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715620003075762000307620002cc565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620003385762000338620002cc565b604052919050565b60006001600160401b038211156200035c576200035c620002cc565b5060051b60200190565b600082601f8301126200037857600080fd5b81516020620003916200038b8362000340565b6200030d565b82815260059290921b84018101918181019086841115620003b157600080fd5b8286015b84811015620003e55780516001600160a01b0381168114620003d75760008081fd5b8352918301918301620003b5565b509695505050505050565b60008060408084860312156200040557600080fd5b83516001600160401b03808211156200041d57600080fd5b818601915086601f8301126200043257600080fd5b81516020620004456200038b8362000340565b82815260069290921b8401810191818101908a8411156200046557600080fd5b948201945b83861015620004aa5786868c031215620004845760008081fd5b6200048e620002e2565b865181528387015184820152825294860194908201906200046a565b91890151919750909450505080831115620004c457600080fd5b5050620004d48582860162000366565b9150509250929050565b634e487b7160e01b600052603260045260246000fd5b6000600182016200051557634e487b7160e01b600052601160045260246000fd5b5060010190565b610940806200052c6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100b2578063ae7fca18146100df578063f2fde38b146100f2578063fcee45f41461010557600080fd5b806345ef670614610082578063604782e61461009757806379ba5097146100aa575b600080fd5b610095610090366004610782565b610126565b005b6100956100a536600461084f565b6102d5565b610095610380565b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100956100ed36600461084f565b61047d565b61009561010036600461084f565b610501565b61011861011336600461088c565b610512565b6040519081526020016100d6565b60005473ffffffffffffffffffffffffffffffffffffffff16331480159061015e57503360009081526003602052604090205460ff16155b15610195576040517f46f0815400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526003602052604090205460ff166101e5576040517faf0026bf0000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b805160005b818110156102d057828181518110610204576102046108a5565b60200260200101516020015160026000858481518110610226576102266108a5565b6020026020010151600001518152602001908152602001600020819055507fa0523da05fa5623ace9527e06812e2f0c03aeec53b39c73bff9488dd090fc56d838281518110610277576102776108a5565b602002602001015160000151848381518110610295576102956108a5565b6020026020010151602001516040516102b8929190918252602082015260400190565b60405180910390a16102c9816108d4565b90506101ea565b505050565b6102dd610563565b73ffffffffffffffffffffffffffffffffffffffff81161561037d5773ffffffffffffffffffffffffffffffffffffffff811660008181526003602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590519182527fa462a3423511fce5ad038a1130251b66e3b8c135fa2ca15777f45c72ea3037c591015b60405180910390a15b50565b60015473ffffffffffffffffffffffffffffffffffffffff163314610401576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016101dc565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610485610563565b73ffffffffffffffffffffffffffffffffffffffff811660008181526003602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905590519182527f74a2c31badb27f0acfb9da3ef34c9e656ca1723881466e89a40f791f1c82ee719101610374565b610509610563565b61037d816105e6565b6000818152600260205260408120549081900361055e576040517f264e42cf000000000000000000000000000000000000000000000000000000008152600481018390526024016101dc565b919050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146105e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016101dc565b565b3373ffffffffffffffffffffffffffffffffffffffff821603610665576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016101dc565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561072d5761072d6106db565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561077a5761077a6106db565b604052919050565b6000602080838503121561079557600080fd5b823567ffffffffffffffff808211156107ad57600080fd5b818501915085601f8301126107c157600080fd5b8135818111156107d3576107d36106db565b6107e1848260051b01610733565b818152848101925060069190911b83018401908782111561080157600080fd5b928401925b81841015610844576040848903121561081f5760008081fd5b61082761070a565b843581528585013586820152835260409093019291840191610806565b979650505050505050565b60006020828403121561086157600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461088557600080fd5b9392505050565b60006020828403121561089e57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361092c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c634300080f000a",
}

var GasFeeCacheABI = GasFeeCacheMetaData.ABI

var GasFeeCacheBin = GasFeeCacheMetaData.Bin

func DeployGasFeeCache(auth *bind.TransactOpts, backend bind.ContractBackend, feeUpdates []CCIPFeeUpdate, feeUpdaters []common.Address) (common.Address, *types.Transaction, *GasFeeCache, error) {
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

func (_GasFeeCache *GasFeeCacheCaller) GetFee(opts *bind.CallOpts, destChainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GasFeeCache.contract.Call(opts, &out, "getFee", destChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_GasFeeCache *GasFeeCacheSession) GetFee(destChainId *big.Int) (*big.Int, error) {
	return _GasFeeCache.Contract.GetFee(&_GasFeeCache.CallOpts, destChainId)
}

func (_GasFeeCache *GasFeeCacheCallerSession) GetFee(destChainId *big.Int) (*big.Int, error) {
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

func (_GasFeeCache *GasFeeCacheTransactor) UpdateFees(opts *bind.TransactOpts, feeUpdates []CCIPFeeUpdate) (*types.Transaction, error) {
	return _GasFeeCache.contract.Transact(opts, "updateFees", feeUpdates)
}

func (_GasFeeCache *GasFeeCacheSession) UpdateFees(feeUpdates []CCIPFeeUpdate) (*types.Transaction, error) {
	return _GasFeeCache.Contract.UpdateFees(&_GasFeeCache.TransactOpts, feeUpdates)
}

func (_GasFeeCache *GasFeeCacheTransactorSession) UpdateFees(feeUpdates []CCIPFeeUpdate) (*types.Transaction, error) {
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
	DestChain      *big.Int
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
	return common.HexToHash("0xa0523da05fa5623ace9527e06812e2f0c03aeec53b39c73bff9488dd090fc56d")
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
	GetFee(opts *bind.CallOpts, destChainId *big.Int) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error)

	SetFeeUpdater(opts *bind.TransactOpts, feeUpdater common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateFees(opts *bind.TransactOpts, feeUpdates []CCIPFeeUpdate) (*types.Transaction, error)

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
