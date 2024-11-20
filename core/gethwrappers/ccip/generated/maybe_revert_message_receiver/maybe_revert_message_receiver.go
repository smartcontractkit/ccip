// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package maybe_revert_message_receiver

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

// ClientAny2EVMMessage is an auto generated low-level Go binding around an user-defined struct.
type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

// ClientEVMTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

// MaybeRevertMessageReceiverMetaData contains all meta data concerning the MaybeRevertMessageReceiver contract.
var MaybeRevertMessageReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"toRevert\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"CustomError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiveRevert\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ValueReceived\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_toRevert\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"setErr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"toRevert\",\"type\":\"bool\"}],\"name\":\"setRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MaybeRevertMessageReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use MaybeRevertMessageReceiverMetaData.ABI instead.
var MaybeRevertMessageReceiverABI = MaybeRevertMessageReceiverMetaData.ABI

// MaybeRevertMessageReceiver is an auto generated Go binding around an Ethereum contract.
type MaybeRevertMessageReceiver struct {
	MaybeRevertMessageReceiverCaller     // Read-only binding to the contract
	MaybeRevertMessageReceiverTransactor // Write-only binding to the contract
	MaybeRevertMessageReceiverFilterer   // Log filterer for contract events
}

// MaybeRevertMessageReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type MaybeRevertMessageReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaybeRevertMessageReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MaybeRevertMessageReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaybeRevertMessageReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MaybeRevertMessageReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaybeRevertMessageReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MaybeRevertMessageReceiverSession struct {
	Contract     *MaybeRevertMessageReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MaybeRevertMessageReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MaybeRevertMessageReceiverCallerSession struct {
	Contract *MaybeRevertMessageReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// MaybeRevertMessageReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MaybeRevertMessageReceiverTransactorSession struct {
	Contract     *MaybeRevertMessageReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// MaybeRevertMessageReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type MaybeRevertMessageReceiverRaw struct {
	Contract *MaybeRevertMessageReceiver // Generic contract binding to access the raw methods on
}

// MaybeRevertMessageReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MaybeRevertMessageReceiverCallerRaw struct {
	Contract *MaybeRevertMessageReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// MaybeRevertMessageReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MaybeRevertMessageReceiverTransactorRaw struct {
	Contract *MaybeRevertMessageReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMaybeRevertMessageReceiver creates a new instance of MaybeRevertMessageReceiver, bound to a specific deployed contract.
func NewMaybeRevertMessageReceiver(address common.Address, backend bind.ContractBackend) (*MaybeRevertMessageReceiver, error) {
	contract, err := bindMaybeRevertMessageReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiver{MaybeRevertMessageReceiverCaller: MaybeRevertMessageReceiverCaller{contract: contract}, MaybeRevertMessageReceiverTransactor: MaybeRevertMessageReceiverTransactor{contract: contract}, MaybeRevertMessageReceiverFilterer: MaybeRevertMessageReceiverFilterer{contract: contract}}, nil
}

// NewMaybeRevertMessageReceiverCaller creates a new read-only instance of MaybeRevertMessageReceiver, bound to a specific deployed contract.
func NewMaybeRevertMessageReceiverCaller(address common.Address, caller bind.ContractCaller) (*MaybeRevertMessageReceiverCaller, error) {
	contract, err := bindMaybeRevertMessageReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverCaller{contract: contract}, nil
}

// NewMaybeRevertMessageReceiverTransactor creates a new write-only instance of MaybeRevertMessageReceiver, bound to a specific deployed contract.
func NewMaybeRevertMessageReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*MaybeRevertMessageReceiverTransactor, error) {
	contract, err := bindMaybeRevertMessageReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverTransactor{contract: contract}, nil
}

// NewMaybeRevertMessageReceiverFilterer creates a new log filterer instance of MaybeRevertMessageReceiver, bound to a specific deployed contract.
func NewMaybeRevertMessageReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*MaybeRevertMessageReceiverFilterer, error) {
	contract, err := bindMaybeRevertMessageReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverFilterer{contract: contract}, nil
}

// bindMaybeRevertMessageReceiver binds a generic wrapper to an already deployed contract.
func bindMaybeRevertMessageReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MaybeRevertMessageReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MaybeRevertMessageReceiver.Contract.MaybeRevertMessageReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.MaybeRevertMessageReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.MaybeRevertMessageReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MaybeRevertMessageReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.contract.Transact(opts, method, params...)
}

// SToRevert is a free data retrieval call binding the contract method 0x5100fc21.
//
// Solidity: function s_toRevert() view returns(bool)
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverCaller) SToRevert(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MaybeRevertMessageReceiver.contract.Call(opts, &out, "s_toRevert")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SToRevert is a free data retrieval call binding the contract method 0x5100fc21.
//
// Solidity: function s_toRevert() view returns(bool)
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) SToRevert() (bool, error) {
	return _MaybeRevertMessageReceiver.Contract.SToRevert(&_MaybeRevertMessageReceiver.CallOpts)
}

// SToRevert is a free data retrieval call binding the contract method 0x5100fc21.
//
// Solidity: function s_toRevert() view returns(bool)
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverCallerSession) SToRevert() (bool, error) {
	return _MaybeRevertMessageReceiver.Contract.SToRevert(&_MaybeRevertMessageReceiver.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MaybeRevertMessageReceiver.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MaybeRevertMessageReceiver.Contract.SupportsInterface(&_MaybeRevertMessageReceiver.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MaybeRevertMessageReceiver.Contract.SupportsInterface(&_MaybeRevertMessageReceiver.CallOpts, interfaceId)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.CcipReceive(&_MaybeRevertMessageReceiver.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.CcipReceive(&_MaybeRevertMessageReceiver.TransactOpts, message)
}

// SetErr is a paid mutator transaction binding the contract method 0x77f5b0e6.
//
// Solidity: function setErr(bytes err) returns()
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactor) SetErr(opts *bind.TransactOpts, err []byte) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.contract.Transact(opts, "setErr", err)
}

// SetErr is a paid mutator transaction binding the contract method 0x77f5b0e6.
//
// Solidity: function setErr(bytes err) returns()
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) SetErr(err []byte) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.SetErr(&_MaybeRevertMessageReceiver.TransactOpts, err)
}

// SetErr is a paid mutator transaction binding the contract method 0x77f5b0e6.
//
// Solidity: function setErr(bytes err) returns()
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorSession) SetErr(err []byte) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.SetErr(&_MaybeRevertMessageReceiver.TransactOpts, err)
}

// SetRevert is a paid mutator transaction binding the contract method 0x8fb5f171.
//
// Solidity: function setRevert(bool toRevert) returns()
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactor) SetRevert(opts *bind.TransactOpts, toRevert bool) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.contract.Transact(opts, "setRevert", toRevert)
}

// SetRevert is a paid mutator transaction binding the contract method 0x8fb5f171.
//
// Solidity: function setRevert(bool toRevert) returns()
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) SetRevert(toRevert bool) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.SetRevert(&_MaybeRevertMessageReceiver.TransactOpts, toRevert)
}

// SetRevert is a paid mutator transaction binding the contract method 0x8fb5f171.
//
// Solidity: function setRevert(bool toRevert) returns()
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorSession) SetRevert(toRevert bool) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.SetRevert(&_MaybeRevertMessageReceiver.TransactOpts, toRevert)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) Receive() (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.Receive(&_MaybeRevertMessageReceiver.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorSession) Receive() (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.Receive(&_MaybeRevertMessageReceiver.TransactOpts)
}

// MaybeRevertMessageReceiverMessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the MaybeRevertMessageReceiver contract.
type MaybeRevertMessageReceiverMessageReceivedIterator struct {
	Event *MaybeRevertMessageReceiverMessageReceived // Event containing the contract specifics and raw log

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
func (it *MaybeRevertMessageReceiverMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaybeRevertMessageReceiverMessageReceived)
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
		it.Event = new(MaybeRevertMessageReceiverMessageReceived)
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
func (it *MaybeRevertMessageReceiverMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaybeRevertMessageReceiverMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaybeRevertMessageReceiverMessageReceived represents a MessageReceived event raised by the MaybeRevertMessageReceiver contract.
type MaybeRevertMessageReceiverMessageReceived struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0x707732b700184c0ab3b799f43f03de9b3606a144cfb367f98291044e71972cdc.
//
// Solidity: event MessageReceived(bytes32 messageId, uint64 sourceChainSelector, bytes sender, bytes data, (address,uint256)[] destTokenAmounts)
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*MaybeRevertMessageReceiverMessageReceivedIterator, error) {

	logs, sub, err := _MaybeRevertMessageReceiver.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverMessageReceivedIterator{contract: _MaybeRevertMessageReceiver.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0x707732b700184c0ab3b799f43f03de9b3606a144cfb367f98291044e71972cdc.
//
// Solidity: event MessageReceived(bytes32 messageId, uint64 sourceChainSelector, bytes sender, bytes data, (address,uint256)[] destTokenAmounts)
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *MaybeRevertMessageReceiverMessageReceived) (event.Subscription, error) {

	logs, sub, err := _MaybeRevertMessageReceiver.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaybeRevertMessageReceiverMessageReceived)
				if err := _MaybeRevertMessageReceiver.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

// ParseMessageReceived is a log parse operation binding the contract event 0x707732b700184c0ab3b799f43f03de9b3606a144cfb367f98291044e71972cdc.
//
// Solidity: event MessageReceived(bytes32 messageId, uint64 sourceChainSelector, bytes sender, bytes data, (address,uint256)[] destTokenAmounts)
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) ParseMessageReceived(log types.Log) (*MaybeRevertMessageReceiverMessageReceived, error) {
	event := new(MaybeRevertMessageReceiverMessageReceived)
	if err := _MaybeRevertMessageReceiver.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaybeRevertMessageReceiverValueReceivedIterator is returned from FilterValueReceived and is used to iterate over the raw logs and unpacked data for ValueReceived events raised by the MaybeRevertMessageReceiver contract.
type MaybeRevertMessageReceiverValueReceivedIterator struct {
	Event *MaybeRevertMessageReceiverValueReceived // Event containing the contract specifics and raw log

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
func (it *MaybeRevertMessageReceiverValueReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaybeRevertMessageReceiverValueReceived)
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
		it.Event = new(MaybeRevertMessageReceiverValueReceived)
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
func (it *MaybeRevertMessageReceiverValueReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaybeRevertMessageReceiverValueReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaybeRevertMessageReceiverValueReceived represents a ValueReceived event raised by the MaybeRevertMessageReceiver contract.
type MaybeRevertMessageReceiverValueReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterValueReceived is a free log retrieval operation binding the contract event 0xe12e3b7047ff60a2dd763cf536a43597e5ce7fe7aa7476345bd4cd079912bcef.
//
// Solidity: event ValueReceived(uint256 amount)
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) FilterValueReceived(opts *bind.FilterOpts) (*MaybeRevertMessageReceiverValueReceivedIterator, error) {

	logs, sub, err := _MaybeRevertMessageReceiver.contract.FilterLogs(opts, "ValueReceived")
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverValueReceivedIterator{contract: _MaybeRevertMessageReceiver.contract, event: "ValueReceived", logs: logs, sub: sub}, nil
}

// WatchValueReceived is a free log subscription operation binding the contract event 0xe12e3b7047ff60a2dd763cf536a43597e5ce7fe7aa7476345bd4cd079912bcef.
//
// Solidity: event ValueReceived(uint256 amount)
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) WatchValueReceived(opts *bind.WatchOpts, sink chan<- *MaybeRevertMessageReceiverValueReceived) (event.Subscription, error) {

	logs, sub, err := _MaybeRevertMessageReceiver.contract.WatchLogs(opts, "ValueReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaybeRevertMessageReceiverValueReceived)
				if err := _MaybeRevertMessageReceiver.contract.UnpackLog(event, "ValueReceived", log); err != nil {
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

// ParseValueReceived is a log parse operation binding the contract event 0xe12e3b7047ff60a2dd763cf536a43597e5ce7fe7aa7476345bd4cd079912bcef.
//
// Solidity: event ValueReceived(uint256 amount)
func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) ParseValueReceived(log types.Log) (*MaybeRevertMessageReceiverValueReceived, error) {
	event := new(MaybeRevertMessageReceiverValueReceived)
	if err := _MaybeRevertMessageReceiver.contract.UnpackLog(event, "ValueReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
