// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simple_message_receiver

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

type CommonAny2EVMMessage struct {
	SourceChainId        uint64
	Sender               []byte
	Data                 []byte
	DestTokensAndAmounts []CommonEVMTokenAndAmount
}

type CommonEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

var SimpleMessageReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCommon.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b503360805260805161010a61002f6000396000604a015261010a6000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80633015b91c146037578063d5009584146048575b600080fd5b6046604236600460be565b6092565b005b7f000000000000000000000000000000000000000000000000000000000000000060405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6040517fd82ce31e3523f6eeb2d24317b2b4133001e8472729657f663b68624c45f8f3e890600090a150565b60006020828403121560cf57600080fd5b813567ffffffffffffffff81111560e557600080fd5b82016080818503121560f657600080fd5b939250505056fea164736f6c634300080f000a",
}

var SimpleMessageReceiverABI = SimpleMessageReceiverMetaData.ABI

var SimpleMessageReceiverBin = SimpleMessageReceiverMetaData.Bin

func DeploySimpleMessageReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleMessageReceiver, error) {
	parsed, err := SimpleMessageReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleMessageReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleMessageReceiver{SimpleMessageReceiverCaller: SimpleMessageReceiverCaller{contract: contract}, SimpleMessageReceiverTransactor: SimpleMessageReceiverTransactor{contract: contract}, SimpleMessageReceiverFilterer: SimpleMessageReceiverFilterer{contract: contract}}, nil
}

type SimpleMessageReceiver struct {
	address common.Address
	abi     abi.ABI
	SimpleMessageReceiverCaller
	SimpleMessageReceiverTransactor
	SimpleMessageReceiverFilterer
}

type SimpleMessageReceiverCaller struct {
	contract *bind.BoundContract
}

type SimpleMessageReceiverTransactor struct {
	contract *bind.BoundContract
}

type SimpleMessageReceiverFilterer struct {
	contract *bind.BoundContract
}

type SimpleMessageReceiverSession struct {
	Contract     *SimpleMessageReceiver
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type SimpleMessageReceiverCallerSession struct {
	Contract *SimpleMessageReceiverCaller
	CallOpts bind.CallOpts
}

type SimpleMessageReceiverTransactorSession struct {
	Contract     *SimpleMessageReceiverTransactor
	TransactOpts bind.TransactOpts
}

type SimpleMessageReceiverRaw struct {
	Contract *SimpleMessageReceiver
}

type SimpleMessageReceiverCallerRaw struct {
	Contract *SimpleMessageReceiverCaller
}

type SimpleMessageReceiverTransactorRaw struct {
	Contract *SimpleMessageReceiverTransactor
}

func NewSimpleMessageReceiver(address common.Address, backend bind.ContractBackend) (*SimpleMessageReceiver, error) {
	abi, err := abi.JSON(strings.NewReader(SimpleMessageReceiverABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindSimpleMessageReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleMessageReceiver{address: address, abi: abi, SimpleMessageReceiverCaller: SimpleMessageReceiverCaller{contract: contract}, SimpleMessageReceiverTransactor: SimpleMessageReceiverTransactor{contract: contract}, SimpleMessageReceiverFilterer: SimpleMessageReceiverFilterer{contract: contract}}, nil
}

func NewSimpleMessageReceiverCaller(address common.Address, caller bind.ContractCaller) (*SimpleMessageReceiverCaller, error) {
	contract, err := bindSimpleMessageReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleMessageReceiverCaller{contract: contract}, nil
}

func NewSimpleMessageReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleMessageReceiverTransactor, error) {
	contract, err := bindSimpleMessageReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleMessageReceiverTransactor{contract: contract}, nil
}

func NewSimpleMessageReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleMessageReceiverFilterer, error) {
	contract, err := bindSimpleMessageReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleMessageReceiverFilterer{contract: contract}, nil
}

func bindSimpleMessageReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleMessageReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_SimpleMessageReceiver *SimpleMessageReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleMessageReceiver.Contract.SimpleMessageReceiverCaller.contract.Call(opts, result, method, params...)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleMessageReceiver.Contract.SimpleMessageReceiverTransactor.contract.Transfer(opts)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleMessageReceiver.Contract.SimpleMessageReceiverTransactor.contract.Transact(opts, method, params...)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleMessageReceiver.Contract.contract.Call(opts, result, method, params...)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleMessageReceiver.Contract.contract.Transfer(opts)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleMessageReceiver.Contract.contract.Transact(opts, method, params...)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverCaller) GetManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleMessageReceiver.contract.Call(opts, &out, "getManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SimpleMessageReceiver *SimpleMessageReceiverSession) GetManager() (common.Address, error) {
	return _SimpleMessageReceiver.Contract.GetManager(&_SimpleMessageReceiver.CallOpts)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverCallerSession) GetManager() (common.Address, error) {
	return _SimpleMessageReceiver.Contract.GetManager(&_SimpleMessageReceiver.CallOpts)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverTransactor) CcipReceive(opts *bind.TransactOpts, arg0 CommonAny2EVMMessage) (*types.Transaction, error) {
	return _SimpleMessageReceiver.contract.Transact(opts, "ccipReceive", arg0)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverSession) CcipReceive(arg0 CommonAny2EVMMessage) (*types.Transaction, error) {
	return _SimpleMessageReceiver.Contract.CcipReceive(&_SimpleMessageReceiver.TransactOpts, arg0)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverTransactorSession) CcipReceive(arg0 CommonAny2EVMMessage) (*types.Transaction, error) {
	return _SimpleMessageReceiver.Contract.CcipReceive(&_SimpleMessageReceiver.TransactOpts, arg0)
}

type SimpleMessageReceiverMessageReceivedIterator struct {
	Event *SimpleMessageReceiverMessageReceived

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SimpleMessageReceiverMessageReceivedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleMessageReceiverMessageReceived)
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
		it.Event = new(SimpleMessageReceiverMessageReceived)
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

func (it *SimpleMessageReceiverMessageReceivedIterator) Error() error {
	return it.fail
}

func (it *SimpleMessageReceiverMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SimpleMessageReceiverMessageReceived struct {
	Raw types.Log
}

func (_SimpleMessageReceiver *SimpleMessageReceiverFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*SimpleMessageReceiverMessageReceivedIterator, error) {

	logs, sub, err := _SimpleMessageReceiver.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &SimpleMessageReceiverMessageReceivedIterator{contract: _SimpleMessageReceiver.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

func (_SimpleMessageReceiver *SimpleMessageReceiverFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *SimpleMessageReceiverMessageReceived) (event.Subscription, error) {

	logs, sub, err := _SimpleMessageReceiver.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SimpleMessageReceiverMessageReceived)
				if err := _SimpleMessageReceiver.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

func (_SimpleMessageReceiver *SimpleMessageReceiverFilterer) ParseMessageReceived(log types.Log) (*SimpleMessageReceiverMessageReceived, error) {
	event := new(SimpleMessageReceiverMessageReceived)
	if err := _SimpleMessageReceiver.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_SimpleMessageReceiver *SimpleMessageReceiver) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _SimpleMessageReceiver.abi.Events["MessageReceived"].ID:
		return _SimpleMessageReceiver.ParseMessageReceived(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (SimpleMessageReceiverMessageReceived) Topic() common.Hash {
	return common.HexToHash("0xd82ce31e3523f6eeb2d24317b2b4133001e8472729657f663b68624c45f8f3e8")
}

func (_SimpleMessageReceiver *SimpleMessageReceiver) Address() common.Address {
	return _SimpleMessageReceiver.address
}

type SimpleMessageReceiverInterface interface {
	GetManager(opts *bind.CallOpts) (common.Address, error)

	CcipReceive(opts *bind.TransactOpts, arg0 CommonAny2EVMMessage) (*types.Transaction, error)

	FilterMessageReceived(opts *bind.FilterOpts) (*SimpleMessageReceiverMessageReceivedIterator, error)

	WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *SimpleMessageReceiverMessageReceived) (event.Subscription, error)

	ParseMessageReceived(log types.Log) (*SimpleMessageReceiverMessageReceived, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
