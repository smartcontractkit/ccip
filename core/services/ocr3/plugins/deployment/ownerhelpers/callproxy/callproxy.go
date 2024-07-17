// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package callproxy

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
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

// CallproxyMetaData contains all meta data concerning the Callproxy contract.
var CallproxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"TargetSet\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161013e38038061013e83398101604081905261002f91610077565b6001600160a01b03811660808190526040519081527f3bfb4bbf112628248058745a3c57e35b13369386e474b8e56c552f3063a4a1969060200160405180910390a1506100a7565b60006020828403121561008957600080fd5b81516001600160a01b03811681146100a057600080fd5b9392505050565b608051607f6100bf600039600060060152607f6000f3fe60806040527f0000000000000000000000000000000000000000000000000000000000000000366000803760008036600034855af13d6000803e80156043573d6000f35b503d6000fdfea26469706673582212202974aca3a8ae03528c7df03132603029149d639b2cd6de0ce90e33abd7a3eb9064736f6c63430008130033",
}

// CallproxyABI is the input ABI used to generate the binding from.
// Deprecated: Use CallproxyMetaData.ABI instead.
var CallproxyABI = CallproxyMetaData.ABI

// CallproxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CallproxyMetaData.Bin instead.
var CallproxyBin = CallproxyMetaData.Bin

// DeployCallproxy deploys a new Ethereum contract, binding an instance of Callproxy to it.
func DeployCallproxy(auth *bind.TransactOpts, backend bind.ContractBackend, target common.Address) (common.Address, *types.Transaction, *Callproxy, error) {
	parsed, err := CallproxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CallproxyBin), backend, target)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Callproxy{CallproxyCaller: CallproxyCaller{contract: contract}, CallproxyTransactor: CallproxyTransactor{contract: contract}, CallproxyFilterer: CallproxyFilterer{contract: contract}}, nil
}

// Callproxy is an auto generated Go binding around an Ethereum contract.
type Callproxy struct {
	CallproxyCaller     // Read-only binding to the contract
	CallproxyTransactor // Write-only binding to the contract
	CallproxyFilterer   // Log filterer for contract events
}

// CallproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type CallproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CallproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallproxySession struct {
	Contract     *Callproxy        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallproxyCallerSession struct {
	Contract *CallproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CallproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallproxyTransactorSession struct {
	Contract     *CallproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CallproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type CallproxyRaw struct {
	Contract *Callproxy // Generic contract binding to access the raw methods on
}

// CallproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallproxyCallerRaw struct {
	Contract *CallproxyCaller // Generic read-only contract binding to access the raw methods on
}

// CallproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallproxyTransactorRaw struct {
	Contract *CallproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCallproxy creates a new instance of Callproxy, bound to a specific deployed contract.
func NewCallproxy(address common.Address, backend bind.ContractBackend) (*Callproxy, error) {
	contract, err := bindCallproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Callproxy{CallproxyCaller: CallproxyCaller{contract: contract}, CallproxyTransactor: CallproxyTransactor{contract: contract}, CallproxyFilterer: CallproxyFilterer{contract: contract}}, nil
}

// NewCallproxyCaller creates a new read-only instance of Callproxy, bound to a specific deployed contract.
func NewCallproxyCaller(address common.Address, caller bind.ContractCaller) (*CallproxyCaller, error) {
	contract, err := bindCallproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallproxyCaller{contract: contract}, nil
}

// NewCallproxyTransactor creates a new write-only instance of Callproxy, bound to a specific deployed contract.
func NewCallproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*CallproxyTransactor, error) {
	contract, err := bindCallproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallproxyTransactor{contract: contract}, nil
}

// NewCallproxyFilterer creates a new log filterer instance of Callproxy, bound to a specific deployed contract.
func NewCallproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*CallproxyFilterer, error) {
	contract, err := bindCallproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallproxyFilterer{contract: contract}, nil
}

// bindCallproxy binds a generic wrapper to an already deployed contract.
func bindCallproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CallproxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Callproxy *CallproxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Callproxy.Contract.CallproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Callproxy *CallproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Callproxy.Contract.CallproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Callproxy *CallproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Callproxy.Contract.CallproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Callproxy *CallproxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Callproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Callproxy *CallproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Callproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Callproxy *CallproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Callproxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Callproxy *CallproxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Callproxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Callproxy *CallproxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Callproxy.Contract.Fallback(&_Callproxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Callproxy *CallproxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Callproxy.Contract.Fallback(&_Callproxy.TransactOpts, calldata)
}

// CallproxyTargetSetIterator is returned from FilterTargetSet and is used to iterate over the raw logs and unpacked data for TargetSet events raised by the Callproxy contract.
type CallproxyTargetSetIterator struct {
	Event *CallproxyTargetSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CallproxyTargetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CallproxyTargetSet)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CallproxyTargetSet)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CallproxyTargetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CallproxyTargetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CallproxyTargetSet represents a TargetSet event raised by the Callproxy contract.
type CallproxyTargetSet struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTargetSet is a free log retrieval operation binding the contract event 0x3bfb4bbf112628248058745a3c57e35b13369386e474b8e56c552f3063a4a196.
//
// Solidity: event TargetSet(address target)
func (_Callproxy *CallproxyFilterer) FilterTargetSet(opts *bind.FilterOpts) (*CallproxyTargetSetIterator, error) {

	logs, sub, err := _Callproxy.contract.FilterLogs(opts, "TargetSet")
	if err != nil {
		return nil, err
	}
	return &CallproxyTargetSetIterator{contract: _Callproxy.contract, event: "TargetSet", logs: logs, sub: sub}, nil
}

// WatchTargetSet is a free log subscription operation binding the contract event 0x3bfb4bbf112628248058745a3c57e35b13369386e474b8e56c552f3063a4a196.
//
// Solidity: event TargetSet(address target)
func (_Callproxy *CallproxyFilterer) WatchTargetSet(opts *bind.WatchOpts, sink chan<- *CallproxyTargetSet) (event.Subscription, error) {

	logs, sub, err := _Callproxy.contract.WatchLogs(opts, "TargetSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CallproxyTargetSet)
				if err := _Callproxy.contract.UnpackLog(event, "TargetSet", log); err != nil {
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

// ParseTargetSet is a log parse operation binding the contract event 0x3bfb4bbf112628248058745a3c57e35b13369386e474b8e56c552f3063a4a196.
//
// Solidity: event TargetSet(address target)
func (_Callproxy *CallproxyFilterer) ParseTargetSet(log types.Log) (*CallproxyTargetSet, error) {
	event := new(CallproxyTargetSet)
	if err := _Callproxy.contract.UnpackLog(event, "TargetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
