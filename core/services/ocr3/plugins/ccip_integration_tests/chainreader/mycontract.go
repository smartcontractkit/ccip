// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SimpleEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"emitEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eventCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEventCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610204806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806371be2e4a146100465780637b0cb83914610064578063d9e48f5c1461006e575b600080fd5b61004e61008c565b60405161005b9190610108565b60405180910390f35b61006c610092565b005b6100766100e6565b6040516100839190610108565b60405180910390f35b60005481565b6000808154809291906100a490610152565b919050555060016000546100b8919061019a565b7f12d199749b3f4c44df8d9386c63d725b7756ec47204f3aa0bf05ea832f89effb60405160405180910390a2565b60008054905090565b6000819050919050565b610102816100ef565b82525050565b600060208201905061011d60008301846100f9565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061015d826100ef565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361018f5761018e610123565b5b600182019050919050565b60006101a5826100ef565b91506101b0836100ef565b92508282019050808211156101c8576101c7610123565b5b9291505056fea2646970667358221220f450e91be013d47a00d1cec0a32b103774ad11d9b184c211adf123824193d47f64736f6c63430008120033",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// MainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MainMetaData.Bin instead.
var MainBin = MainMetaData.Bin

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// EventCount is a free data retrieval call binding the contract method 0x71be2e4a.
//
// Solidity: function eventCount() view returns(uint256)
func (_Main *MainCaller) EventCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "eventCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventCount is a free data retrieval call binding the contract method 0x71be2e4a.
//
// Solidity: function eventCount() view returns(uint256)
func (_Main *MainSession) EventCount() (*big.Int, error) {
	return _Main.Contract.EventCount(&_Main.CallOpts)
}

// EventCount is a free data retrieval call binding the contract method 0x71be2e4a.
//
// Solidity: function eventCount() view returns(uint256)
func (_Main *MainCallerSession) EventCount() (*big.Int, error) {
	return _Main.Contract.EventCount(&_Main.CallOpts)
}

// GetEventCount is a free data retrieval call binding the contract method 0xd9e48f5c.
//
// Solidity: function getEventCount() view returns(uint256)
func (_Main *MainCaller) GetEventCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getEventCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEventCount is a free data retrieval call binding the contract method 0xd9e48f5c.
//
// Solidity: function getEventCount() view returns(uint256)
func (_Main *MainSession) GetEventCount() (*big.Int, error) {
	return _Main.Contract.GetEventCount(&_Main.CallOpts)
}

// GetEventCount is a free data retrieval call binding the contract method 0xd9e48f5c.
//
// Solidity: function getEventCount() view returns(uint256)
func (_Main *MainCallerSession) GetEventCount() (*big.Int, error) {
	return _Main.Contract.GetEventCount(&_Main.CallOpts)
}

// EmitEvent is a paid mutator transaction binding the contract method 0x7b0cb839.
//
// Solidity: function emitEvent() returns()
func (_Main *MainTransactor) EmitEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "emitEvent")
}

// EmitEvent is a paid mutator transaction binding the contract method 0x7b0cb839.
//
// Solidity: function emitEvent() returns()
func (_Main *MainSession) EmitEvent() (*types.Transaction, error) {
	return _Main.Contract.EmitEvent(&_Main.TransactOpts)
}

// EmitEvent is a paid mutator transaction binding the contract method 0x7b0cb839.
//
// Solidity: function emitEvent() returns()
func (_Main *MainTransactorSession) EmitEvent() (*types.Transaction, error) {
	return _Main.Contract.EmitEvent(&_Main.TransactOpts)
}

// MainSimpleEventIterator is returned from FilterSimpleEvent and is used to iterate over the raw logs and unpacked data for SimpleEvent events raised by the Main contract.
type MainSimpleEventIterator struct {
	Event *MainSimpleEvent // Event containing the contract specifics and raw log

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
func (it *MainSimpleEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainSimpleEvent)
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
		it.Event = new(MainSimpleEvent)
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
func (it *MainSimpleEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainSimpleEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainSimpleEvent represents a SimpleEvent event raised by the Main contract.
type MainSimpleEvent struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSimpleEvent is a free log retrieval operation binding the contract event 0x12d199749b3f4c44df8d9386c63d725b7756ec47204f3aa0bf05ea832f89effb.
//
// Solidity: event SimpleEvent(uint256 indexed value)
func (_Main *MainFilterer) FilterSimpleEvent(opts *bind.FilterOpts, value []*big.Int) (*MainSimpleEventIterator, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "SimpleEvent", valueRule)
	if err != nil {
		return nil, err
	}
	return &MainSimpleEventIterator{contract: _Main.contract, event: "SimpleEvent", logs: logs, sub: sub}, nil
}

// WatchSimpleEvent is a free log subscription operation binding the contract event 0x12d199749b3f4c44df8d9386c63d725b7756ec47204f3aa0bf05ea832f89effb.
//
// Solidity: event SimpleEvent(uint256 indexed value)
func (_Main *MainFilterer) WatchSimpleEvent(opts *bind.WatchOpts, sink chan<- *MainSimpleEvent, value []*big.Int) (event.Subscription, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "SimpleEvent", valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainSimpleEvent)
				if err := _Main.contract.UnpackLog(event, "SimpleEvent", log); err != nil {
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

// ParseSimpleEvent is a log parse operation binding the contract event 0x12d199749b3f4c44df8d9386c63d725b7756ec47204f3aa0bf05ea832f89effb.
//
// Solidity: event SimpleEvent(uint256 indexed value)
func (_Main *MainFilterer) ParseSimpleEvent(log types.Log) (*MainSimpleEvent, error) {
	event := new(MainSimpleEvent)
	if err := _Main.contract.UnpackLog(event, "SimpleEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
