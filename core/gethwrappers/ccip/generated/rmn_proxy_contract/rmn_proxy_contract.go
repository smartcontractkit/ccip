package rmn_proxy_contract

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
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
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
	_ = abi.ConvertType
)

var RMNProxyContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arm\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"arm\",\"type\":\"address\"}],\"name\":\"ARMSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getARM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arm\",\"type\":\"address\"}],\"name\":\"setARM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161084138038061084183398101604081905261002f91610255565b33806000816100855760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100b5576100b5816100cd565b5050506100c78161017660201b60201c565b50610285565b336001600160a01b038216036101255760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161007c565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b61017e6101f9565b6001600160a01b0381166101a5576040516342bcdf7f60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527fef31f568d741a833c6a9dc85a6e1c65e06fa772740d5dc94d1da21827a4e0cab9060200160405180910390a150565b6000546001600160a01b031633146102535760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161007c565b565b60006020828403121561026757600080fd5b81516001600160a01b038116811461027e57600080fd5b9392505050565b6105ad806102946000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c806379ba50971161005057806379ba5097146101615780638da5cb5b14610169578063f2fde38b1461018757610072565b8063181f5a77146100bb5780632e90aa211461010d578063458fec3b1461014c575b60025473ffffffffffffffffffffffffffffffffffffffff16803b61009657600080fd5b366000803760008036600080855af13d6000803e80156100b5573d6000f35b503d6000fd5b6100f76040518060400160405280600e81526020017f41524d50726f787920312e302e3000000000000000000000000000000000000081525081565b60405161010491906104f6565b60405180910390f35b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610104565b61015f61015a366004610563565b61019a565b005b61015f610268565b60005473ffffffffffffffffffffffffffffffffffffffff16610127565b61015f610195366004610563565b61036a565b6101a261037e565b73ffffffffffffffffffffffffffffffffffffffff81166101ef576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fef31f568d741a833c6a9dc85a6e1c65e06fa772740d5dc94d1da21827a4e0cab9060200160405180910390a150565b60015473ffffffffffffffffffffffffffffffffffffffff1633146102ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61037261037e565b61037b81610401565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103ff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016102e5565b565b3373ffffffffffffffffffffffffffffffffffffffff821603610480576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016102e5565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020808352835180602085015260005b8181101561052457858101830151858201604001528201610508565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b60006020828403121561057557600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461059957600080fd5b939250505056fea164736f6c6343000818000a",
}

var RMNProxyContractABI = RMNProxyContractMetaData.ABI

var RMNProxyContractBin = RMNProxyContractMetaData.Bin

func DeployRMNProxyContract(auth *bind.TransactOpts, backend bind.ContractBackend, arm common.Address) (common.Address, *generated.Transaction, *RMNProxyContract, error) {
	parsed, err := RMNProxyContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}
	if generated.IsZKSync(backend) {
		address, ethTx, contractBind, _ := generated.DeployContract(auth, parsed, common.FromHex(RMNProxyContractZKBin), backend, arm)
		contractReturn := &RMNProxyContract{address: address, abi: *parsed, RMNProxyContractCaller: RMNProxyContractCaller{contract: contractBind}, RMNProxyContractTransactor: RMNProxyContractTransactor{contract: contractBind}, RMNProxyContractFilterer: RMNProxyContractFilterer{contract: contractBind}}
		return address, ethTx, contractReturn, err
	}
	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RMNProxyContractBin), backend, arm)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, &generated.Transaction{Transaction: tx, HashZks: tx.Hash()}, &RMNProxyContract{address: address, abi: *parsed, RMNProxyContractCaller: RMNProxyContractCaller{contract: contract}, RMNProxyContractTransactor: RMNProxyContractTransactor{contract: contract}, RMNProxyContractFilterer: RMNProxyContractFilterer{contract: contract}}, nil
}

type RMNProxyContract struct {
	address common.Address
	abi     abi.ABI
	RMNProxyContractCaller
	RMNProxyContractTransactor
	RMNProxyContractFilterer
}

type RMNProxyContractCaller struct {
	contract *bind.BoundContract
}

type RMNProxyContractTransactor struct {
	contract *bind.BoundContract
}

type RMNProxyContractFilterer struct {
	contract *bind.BoundContract
}

type RMNProxyContractSession struct {
	Contract     *RMNProxyContract
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type RMNProxyContractCallerSession struct {
	Contract *RMNProxyContractCaller
	CallOpts bind.CallOpts
}

type RMNProxyContractTransactorSession struct {
	Contract     *RMNProxyContractTransactor
	TransactOpts bind.TransactOpts
}

type RMNProxyContractRaw struct {
	Contract *RMNProxyContract
}

type RMNProxyContractCallerRaw struct {
	Contract *RMNProxyContractCaller
}

type RMNProxyContractTransactorRaw struct {
	Contract *RMNProxyContractTransactor
}

func NewRMNProxyContract(address common.Address, backend bind.ContractBackend) (*RMNProxyContract, error) {
	abi, err := abi.JSON(strings.NewReader(RMNProxyContractABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindRMNProxyContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RMNProxyContract{address: address, abi: abi, RMNProxyContractCaller: RMNProxyContractCaller{contract: contract}, RMNProxyContractTransactor: RMNProxyContractTransactor{contract: contract}, RMNProxyContractFilterer: RMNProxyContractFilterer{contract: contract}}, nil
}

func NewRMNProxyContractCaller(address common.Address, caller bind.ContractCaller) (*RMNProxyContractCaller, error) {
	contract, err := bindRMNProxyContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RMNProxyContractCaller{contract: contract}, nil
}

func NewRMNProxyContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RMNProxyContractTransactor, error) {
	contract, err := bindRMNProxyContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RMNProxyContractTransactor{contract: contract}, nil
}

func NewRMNProxyContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RMNProxyContractFilterer, error) {
	contract, err := bindRMNProxyContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RMNProxyContractFilterer{contract: contract}, nil
}

func bindRMNProxyContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RMNProxyContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_RMNProxyContract *RMNProxyContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RMNProxyContract.Contract.RMNProxyContractCaller.contract.Call(opts, result, method, params...)
}

func (_RMNProxyContract *RMNProxyContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RMNProxyContract.Contract.RMNProxyContractTransactor.contract.Transfer(opts)
}

func (_RMNProxyContract *RMNProxyContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RMNProxyContract.Contract.RMNProxyContractTransactor.contract.Transact(opts, method, params...)
}

func (_RMNProxyContract *RMNProxyContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RMNProxyContract.Contract.contract.Call(opts, result, method, params...)
}

func (_RMNProxyContract *RMNProxyContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RMNProxyContract.Contract.contract.Transfer(opts)
}

func (_RMNProxyContract *RMNProxyContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RMNProxyContract.Contract.contract.Transact(opts, method, params...)
}

func (_RMNProxyContract *RMNProxyContractCaller) GetARM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RMNProxyContract.contract.Call(opts, &out, "getARM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RMNProxyContract *RMNProxyContractSession) GetARM() (common.Address, error) {
	return _RMNProxyContract.Contract.GetARM(&_RMNProxyContract.CallOpts)
}

func (_RMNProxyContract *RMNProxyContractCallerSession) GetARM() (common.Address, error) {
	return _RMNProxyContract.Contract.GetARM(&_RMNProxyContract.CallOpts)
}

func (_RMNProxyContract *RMNProxyContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RMNProxyContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RMNProxyContract *RMNProxyContractSession) Owner() (common.Address, error) {
	return _RMNProxyContract.Contract.Owner(&_RMNProxyContract.CallOpts)
}

func (_RMNProxyContract *RMNProxyContractCallerSession) Owner() (common.Address, error) {
	return _RMNProxyContract.Contract.Owner(&_RMNProxyContract.CallOpts)
}

func (_RMNProxyContract *RMNProxyContractCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RMNProxyContract.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_RMNProxyContract *RMNProxyContractSession) TypeAndVersion() (string, error) {
	return _RMNProxyContract.Contract.TypeAndVersion(&_RMNProxyContract.CallOpts)
}

func (_RMNProxyContract *RMNProxyContractCallerSession) TypeAndVersion() (string, error) {
	return _RMNProxyContract.Contract.TypeAndVersion(&_RMNProxyContract.CallOpts)
}

func (_RMNProxyContract *RMNProxyContractTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RMNProxyContract.contract.Transact(opts, "acceptOwnership")
}

func (_RMNProxyContract *RMNProxyContractSession) AcceptOwnership() (*types.Transaction, error) {
	return _RMNProxyContract.Contract.AcceptOwnership(&_RMNProxyContract.TransactOpts)
}

func (_RMNProxyContract *RMNProxyContractTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _RMNProxyContract.Contract.AcceptOwnership(&_RMNProxyContract.TransactOpts)
}

func (_RMNProxyContract *RMNProxyContractTransactor) SetARM(opts *bind.TransactOpts, arm common.Address) (*types.Transaction, error) {
	return _RMNProxyContract.contract.Transact(opts, "setARM", arm)
}

func (_RMNProxyContract *RMNProxyContractSession) SetARM(arm common.Address) (*types.Transaction, error) {
	return _RMNProxyContract.Contract.SetARM(&_RMNProxyContract.TransactOpts, arm)
}

func (_RMNProxyContract *RMNProxyContractTransactorSession) SetARM(arm common.Address) (*types.Transaction, error) {
	return _RMNProxyContract.Contract.SetARM(&_RMNProxyContract.TransactOpts, arm)
}

func (_RMNProxyContract *RMNProxyContractTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RMNProxyContract.contract.Transact(opts, "transferOwnership", to)
}

func (_RMNProxyContract *RMNProxyContractSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RMNProxyContract.Contract.TransferOwnership(&_RMNProxyContract.TransactOpts, to)
}

func (_RMNProxyContract *RMNProxyContractTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RMNProxyContract.Contract.TransferOwnership(&_RMNProxyContract.TransactOpts, to)
}

func (_RMNProxyContract *RMNProxyContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RMNProxyContract.contract.RawTransact(opts, calldata)
}

func (_RMNProxyContract *RMNProxyContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RMNProxyContract.Contract.Fallback(&_RMNProxyContract.TransactOpts, calldata)
}

func (_RMNProxyContract *RMNProxyContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RMNProxyContract.Contract.Fallback(&_RMNProxyContract.TransactOpts, calldata)
}

type RMNProxyContractARMSetIterator struct {
	Event *RMNProxyContractARMSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNProxyContractARMSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNProxyContractARMSet)
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
		it.Event = new(RMNProxyContractARMSet)
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

func (it *RMNProxyContractARMSetIterator) Error() error {
	return it.fail
}

func (it *RMNProxyContractARMSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNProxyContractARMSet struct {
	Arm common.Address
	Raw types.Log
}

func (_RMNProxyContract *RMNProxyContractFilterer) FilterARMSet(opts *bind.FilterOpts) (*RMNProxyContractARMSetIterator, error) {

	logs, sub, err := _RMNProxyContract.contract.FilterLogs(opts, "ARMSet")
	if err != nil {
		return nil, err
	}
	return &RMNProxyContractARMSetIterator{contract: _RMNProxyContract.contract, event: "ARMSet", logs: logs, sub: sub}, nil
}

func (_RMNProxyContract *RMNProxyContractFilterer) WatchARMSet(opts *bind.WatchOpts, sink chan<- *RMNProxyContractARMSet) (event.Subscription, error) {

	logs, sub, err := _RMNProxyContract.contract.WatchLogs(opts, "ARMSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNProxyContractARMSet)
				if err := _RMNProxyContract.contract.UnpackLog(event, "ARMSet", log); err != nil {
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

func (_RMNProxyContract *RMNProxyContractFilterer) ParseARMSet(log types.Log) (*RMNProxyContractARMSet, error) {
	event := new(RMNProxyContractARMSet)
	if err := _RMNProxyContract.contract.UnpackLog(event, "ARMSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNProxyContractOwnershipTransferRequestedIterator struct {
	Event *RMNProxyContractOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNProxyContractOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNProxyContractOwnershipTransferRequested)
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
		it.Event = new(RMNProxyContractOwnershipTransferRequested)
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

func (it *RMNProxyContractOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *RMNProxyContractOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNProxyContractOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RMNProxyContract *RMNProxyContractFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNProxyContractOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMNProxyContract.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RMNProxyContractOwnershipTransferRequestedIterator{contract: _RMNProxyContract.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_RMNProxyContract *RMNProxyContractFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RMNProxyContractOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMNProxyContract.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNProxyContractOwnershipTransferRequested)
				if err := _RMNProxyContract.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_RMNProxyContract *RMNProxyContractFilterer) ParseOwnershipTransferRequested(log types.Log) (*RMNProxyContractOwnershipTransferRequested, error) {
	event := new(RMNProxyContractOwnershipTransferRequested)
	if err := _RMNProxyContract.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNProxyContractOwnershipTransferredIterator struct {
	Event *RMNProxyContractOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNProxyContractOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNProxyContractOwnershipTransferred)
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
		it.Event = new(RMNProxyContractOwnershipTransferred)
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

func (it *RMNProxyContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *RMNProxyContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNProxyContractOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RMNProxyContract *RMNProxyContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNProxyContractOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMNProxyContract.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RMNProxyContractOwnershipTransferredIterator{contract: _RMNProxyContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_RMNProxyContract *RMNProxyContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RMNProxyContractOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMNProxyContract.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNProxyContractOwnershipTransferred)
				if err := _RMNProxyContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_RMNProxyContract *RMNProxyContractFilterer) ParseOwnershipTransferred(log types.Log) (*RMNProxyContractOwnershipTransferred, error) {
	event := new(RMNProxyContractOwnershipTransferred)
	if err := _RMNProxyContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_RMNProxyContract *RMNProxyContract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _RMNProxyContract.abi.Events["ARMSet"].ID:
		return _RMNProxyContract.ParseARMSet(log)
	case _RMNProxyContract.abi.Events["OwnershipTransferRequested"].ID:
		return _RMNProxyContract.ParseOwnershipTransferRequested(log)
	case _RMNProxyContract.abi.Events["OwnershipTransferred"].ID:
		return _RMNProxyContract.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (RMNProxyContractARMSet) Topic() common.Hash {
	return common.HexToHash("0xef31f568d741a833c6a9dc85a6e1c65e06fa772740d5dc94d1da21827a4e0cab")
}

func (RMNProxyContractOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (RMNProxyContractOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_RMNProxyContract *RMNProxyContract) Address() common.Address {
	return _RMNProxyContract.address
}

type RMNProxyContractInterface interface {
	GetARM(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetARM(opts *bind.TransactOpts, arm common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error)

	FilterARMSet(opts *bind.FilterOpts) (*RMNProxyContractARMSetIterator, error)

	WatchARMSet(opts *bind.WatchOpts, sink chan<- *RMNProxyContractARMSet) (event.Subscription, error)

	ParseARMSet(log types.Log) (*RMNProxyContractARMSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNProxyContractOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RMNProxyContractOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*RMNProxyContractOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNProxyContractOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RMNProxyContractOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*RMNProxyContractOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}

var RMNProxyContractZKBin = ("0x0004000000000002000100000000000200000000030100190000006004300270000000670340019700030000003103550002000000010355000000670040019d00000001002001900000002c0000c13d0000008002000039000000400020043f000000040030008c000000830000413d000000000201043b000000e002200270000000740020009c000000630000213d000000780020009c000000cf0000613d000000790020009c000000df0000613d0000007a0020009c000000830000c13d000000240030008c0000018c0000413d0000000002000416000000000002004b0000018c0000c13d0000000401100370000000000101043b0000006a0010009c0000018c0000213d000000000200041a0000006a022001970000000003000411000000000023004b000001550000c13d000000000001004b0000016b0000c13d0000006f01000041000000800010043f00000084010000410000019b000104300000000002000416000000000002004b0000018c0000c13d0000001f0230003900000068022001970000008002200039000000400020043f0000001f0430018f000000690530019800000080025000390000003d0000613d0000008006000039000000000701034f000000007807043c0000000006860436000000000026004b000000390000c13d000000000004004b0000004a0000613d000000000151034f0000000304400210000000000502043300000000054501cf000000000545022f000000000101043b0000010004400089000000000141022f00000000014101cf000000000151019f0000000000120435000000200030008c0000018c0000413d000000800100043d0000006a0010009c0000018c0000213d0000000003000411000000000003004b000001060000c13d000000400100043d00000044021000390000007103000041000000000032043500000024021000390000001803000039000000000032043500000072020000410000000000210435000000040210003900000020030000390000000000320435000000670010009c0000006701008041000000400110021000000073011001c70000019b00010430000000750020009c000000e50000613d000000760020009c000000fe0000613d000000770020009c000000830000c13d000000240030008c0000018c0000413d0000000002000416000000000002004b0000018c0000c13d0000000401100370000000000601043b0000006a0060009c0000018c0000213d000000000100041a0000006a011001970000000005000411000000000015004b000001550000c13d000000000056004b0000017a0000c13d0000007201000041000000800010043f0000002001000039000000840010043f0000001701000039000000a40010043f0000007d01000041000000c40010043f0000007e010000410000019b000104300000000001000416000000000001004b0000018c0000c13d0000000201000039000000000101041a000000870200004100000000002004430000006a01100197000100000001001d00000004001004430000000001000414000000670010009c0000006701008041000000c00110021000000088011001c70000800202000039019901940000040f0000000100200190000000ce0000613d000000000101043b000000000001004b0000018c0000613d0000000204000367000000000100003100000089031001980000001f0510018f000000a40000613d000000000604034f0000000007000019000000006806043c0000000007870436000000000037004b000000a00000c13d000000000005004b000000b10000613d000000000434034f0000000305500210000000000603043300000000065601cf000000000656022f000000000404043b0000010005500089000000000454022f00000000045401cf000000000464019f000000000043043500000000030004140000000102000029000000040020008c000001140000c13d0000000303000367000000010100003100000089021001980000001f0410018f000000c00000613d000000000503034f0000000006000019000000005705043c0000000006760436000000000026004b000000bc0000c13d000000000004004b000001370000613d000000000323034f0000000304400210000000000502043300000000054501cf000000000545022f000000000303043b0000010004400089000000000343022f00000000034301cf000000000353019f0000000000320435000001370000013d000000000001042f0000000001000416000000000001004b0000018c0000c13d000000c001000039000000400010043f0000000e01000039000000800010043f0000008502000041000000a00020043f0000002003000039000000c00030043f000000e00010043f000001000020043f0000010e0000043f00000086010000410000019a0001042e0000000001000416000000000001004b0000018c0000c13d0000000201000039000000000101041a000001020000013d0000000001000416000000000001004b0000018c0000c13d0000000101000039000000000201041a0000006a032001970000000006000411000000000036004b0000015f0000c13d000000000300041a0000006b04300197000000000464019f000000000040041b0000006b02200197000000000021041b00000000010004140000006a05300197000000670010009c0000006701008041000000c0011002100000007b011001c70000800d0200003900000003030000390000008104000041000001870000013d0000000001000416000000000001004b0000018c0000c13d000000000100041a0000006a01100197000000800010043f0000007f010000410000019a0001042e0000006a02100198000000000100041a0000006b01100197000000000131019f000000000010041b000000400100043d0000013b0000c13d0000006f020000410000000000210435000000670010009c0000006701008041000000400110021000000070011001c70000019b000104300000006001100210000000670030009c0000006703008041000000c00330021000000000011300190199018f0000040f0003000000010355000000000301001900000060033002700000001f0530018f000100670030019d0000006904300198000001270000613d000000000601034f0000000007000019000000006806043c0000000007870436000000000047004b000001230000c13d000000000005004b000001340000613d000000000141034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f000000000014043500000067013001970000000100200190000001690000613d000000670010009c000000670100804100000060011002100000019a0001042e0000000203000039000000000403041a0000006b04400197000000000424019f000000000043041b0000000000210435000000670010009c000000670100804100000040011002100000000002000414000000670020009c0000006702008041000000c002200210000000000112019f0000006c011001c70000800d0200003900000001030000390000006d040000410199018f0000040f00000001002001900000018c0000613d0000002001000039000001000010044300000120000004430000006e010000410000019a0001042e0000007201000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f0000008201000041000000c40010043f0000007e010000410000019b000104300000007201000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f0000008001000041000000c40010043f0000007e010000410000019b0001043000000060011002100000019b000104300000000202000039000000000302041a0000006b03300197000000000313019f000000000032041b000000800010043f0000000001000414000000670010009c0000006701008041000000c00110021000000083011001c70000800d0200003900000001030000390000006d04000041000001870000013d0000000101000039000000000201041a0000006b02200197000000000262019f000000000021041b0000000001000414000000670010009c0000006701008041000000c0011002100000007b011001c70000800d0200003900000003030000390000007c040000410199018f0000040f00000001002001900000018c0000613d00000000010000190000019a0001042e00000000010000190000019b00010430000000000001042f00000192002104210000000102000039000000000001042d0000000002000019000000000001042d00000197002104230000000102000039000000000001042d0000000002000019000000000001042d00000199000004320000019a0001042e0000019b0001043000000000000000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000000000000000000000000001ffffffe000000000000000000000000000000000000000000000000000000000ffffffe0000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000000000000000000000000000000000000200000000000000000000000000000000000020000000000000000000000000ef31f568d741a833c6a9dc85a6e1c65e06fa772740d5dc94d1da21827a4e0cab00000002000000000000000000000000000000400000010000000000000000008579befe00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000043616e6e6f7420736574206f776e657220746f207a65726f000000000000000008c379a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000000000000000000000079ba50960000000000000000000000000000000000000000000000000000000079ba5097000000000000000000000000000000000000000000000000000000008da5cb5b00000000000000000000000000000000000000000000000000000000f2fde38b00000000000000000000000000000000000000000000000000000000181f5a77000000000000000000000000000000000000000000000000000000002e90aa2100000000000000000000000000000000000000000000000000000000458fec3b0200000000000000000000000000000000000000000000000000000000000000ed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127843616e6e6f74207472616e7366657220746f2073656c66000000000000000000000000000000000000000000000000000000006400000080000000000000000000000000000000000000000000000000000000200000008000000000000000004d7573742062652070726f706f736564206f776e6572000000000000000000008be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e04f6e6c792063616c6c61626c65206279206f776e6572000000000000000000000200000000000000000000000000000000000020000000800000000000000000000000000000000000000000000000000000000400000080000000000000000041524d50726f787920312e302e300000000000000000000000000000000000000000000000000000000000000000000000000060000000c000000000000000001806aa1896bbf26568e884a7374b41e002500962caba6a15023a8d90e8508b830200000200000000000000000000000000000024000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00000000000000000000000000000000000000000000000000000000000000000")
