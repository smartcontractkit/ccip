// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package maybe_revert_message_receiver

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

var MaybeRevertMessageReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"toRevert\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCommon.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_toRevert\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"toRevert\",\"type\":\"bool\"}],\"name\":\"setRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161024638038061024683398101604081905261002f9161005d565b600080546001600160a81b0319163360ff60a01b191617600160a01b92151592909202919091179055610086565b60006020828403121561006f57600080fd5b8151801515811461007f57600080fd5b9392505050565b6101b1806100956000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80633015b91c146100465780635100fc211461005b5780638fb5f17114610094575b600080fd5b610059610054366004610140565b6100ec565b005b6000546100809074010000000000000000000000000000000000000000900460ff1681565b604051901515815260200160405180910390f35b6100596100a2366004610182565b6000805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b60005474010000000000000000000000000000000000000000900460ff161561011457600080fd5b6040517fd82ce31e3523f6eeb2d24317b2b4133001e8472729657f663b68624c45f8f3e890600090a150565b60006020828403121561015257600080fd5b813567ffffffffffffffff81111561016957600080fd5b82016080818503121561017b57600080fd5b9392505050565b60006020828403121561019457600080fd5b8135801515811461017b57600080fdfea164736f6c634300080f000a",
}

var MaybeRevertMessageReceiverABI = MaybeRevertMessageReceiverMetaData.ABI

var MaybeRevertMessageReceiverBin = MaybeRevertMessageReceiverMetaData.Bin

func DeployMaybeRevertMessageReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, toRevert bool) (common.Address, *types.Transaction, *MaybeRevertMessageReceiver, error) {
	parsed, err := MaybeRevertMessageReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MaybeRevertMessageReceiverBin), backend, toRevert)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MaybeRevertMessageReceiver{MaybeRevertMessageReceiverCaller: MaybeRevertMessageReceiverCaller{contract: contract}, MaybeRevertMessageReceiverTransactor: MaybeRevertMessageReceiverTransactor{contract: contract}, MaybeRevertMessageReceiverFilterer: MaybeRevertMessageReceiverFilterer{contract: contract}}, nil
}

type MaybeRevertMessageReceiver struct {
	address common.Address
	abi     abi.ABI
	MaybeRevertMessageReceiverCaller
	MaybeRevertMessageReceiverTransactor
	MaybeRevertMessageReceiverFilterer
}

type MaybeRevertMessageReceiverCaller struct {
	contract *bind.BoundContract
}

type MaybeRevertMessageReceiverTransactor struct {
	contract *bind.BoundContract
}

type MaybeRevertMessageReceiverFilterer struct {
	contract *bind.BoundContract
}

type MaybeRevertMessageReceiverSession struct {
	Contract     *MaybeRevertMessageReceiver
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MaybeRevertMessageReceiverCallerSession struct {
	Contract *MaybeRevertMessageReceiverCaller
	CallOpts bind.CallOpts
}

type MaybeRevertMessageReceiverTransactorSession struct {
	Contract     *MaybeRevertMessageReceiverTransactor
	TransactOpts bind.TransactOpts
}

type MaybeRevertMessageReceiverRaw struct {
	Contract *MaybeRevertMessageReceiver
}

type MaybeRevertMessageReceiverCallerRaw struct {
	Contract *MaybeRevertMessageReceiverCaller
}

type MaybeRevertMessageReceiverTransactorRaw struct {
	Contract *MaybeRevertMessageReceiverTransactor
}

func NewMaybeRevertMessageReceiver(address common.Address, backend bind.ContractBackend) (*MaybeRevertMessageReceiver, error) {
	abi, err := abi.JSON(strings.NewReader(MaybeRevertMessageReceiverABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMaybeRevertMessageReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiver{address: address, abi: abi, MaybeRevertMessageReceiverCaller: MaybeRevertMessageReceiverCaller{contract: contract}, MaybeRevertMessageReceiverTransactor: MaybeRevertMessageReceiverTransactor{contract: contract}, MaybeRevertMessageReceiverFilterer: MaybeRevertMessageReceiverFilterer{contract: contract}}, nil
}

func NewMaybeRevertMessageReceiverCaller(address common.Address, caller bind.ContractCaller) (*MaybeRevertMessageReceiverCaller, error) {
	contract, err := bindMaybeRevertMessageReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverCaller{contract: contract}, nil
}

func NewMaybeRevertMessageReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*MaybeRevertMessageReceiverTransactor, error) {
	contract, err := bindMaybeRevertMessageReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverTransactor{contract: contract}, nil
}

func NewMaybeRevertMessageReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*MaybeRevertMessageReceiverFilterer, error) {
	contract, err := bindMaybeRevertMessageReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverFilterer{contract: contract}, nil
}

func bindMaybeRevertMessageReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MaybeRevertMessageReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MaybeRevertMessageReceiver.Contract.MaybeRevertMessageReceiverCaller.contract.Call(opts, result, method, params...)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.MaybeRevertMessageReceiverTransactor.contract.Transfer(opts)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.MaybeRevertMessageReceiverTransactor.contract.Transact(opts, method, params...)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MaybeRevertMessageReceiver.Contract.contract.Call(opts, result, method, params...)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.contract.Transfer(opts)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.contract.Transact(opts, method, params...)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverCaller) SToRevert(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MaybeRevertMessageReceiver.contract.Call(opts, &out, "s_toRevert")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) SToRevert() (bool, error) {
	return _MaybeRevertMessageReceiver.Contract.SToRevert(&_MaybeRevertMessageReceiver.CallOpts)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverCallerSession) SToRevert() (bool, error) {
	return _MaybeRevertMessageReceiver.Contract.SToRevert(&_MaybeRevertMessageReceiver.CallOpts)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactor) CcipReceive(opts *bind.TransactOpts, arg0 CommonAny2EVMMessage) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.contract.Transact(opts, "ccipReceive", arg0)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) CcipReceive(arg0 CommonAny2EVMMessage) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.CcipReceive(&_MaybeRevertMessageReceiver.TransactOpts, arg0)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorSession) CcipReceive(arg0 CommonAny2EVMMessage) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.CcipReceive(&_MaybeRevertMessageReceiver.TransactOpts, arg0)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactor) SetRevert(opts *bind.TransactOpts, toRevert bool) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.contract.Transact(opts, "setRevert", toRevert)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) SetRevert(toRevert bool) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.SetRevert(&_MaybeRevertMessageReceiver.TransactOpts, toRevert)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorSession) SetRevert(toRevert bool) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.SetRevert(&_MaybeRevertMessageReceiver.TransactOpts, toRevert)
}

type MaybeRevertMessageReceiverMessageReceivedIterator struct {
	Event *MaybeRevertMessageReceiverMessageReceived

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MaybeRevertMessageReceiverMessageReceivedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

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

func (it *MaybeRevertMessageReceiverMessageReceivedIterator) Error() error {
	return it.fail
}

func (it *MaybeRevertMessageReceiverMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MaybeRevertMessageReceiverMessageReceived struct {
	Raw types.Log
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*MaybeRevertMessageReceiverMessageReceivedIterator, error) {

	logs, sub, err := _MaybeRevertMessageReceiver.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverMessageReceivedIterator{contract: _MaybeRevertMessageReceiver.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

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

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) ParseMessageReceived(log types.Log) (*MaybeRevertMessageReceiverMessageReceived, error) {
	event := new(MaybeRevertMessageReceiverMessageReceived)
	if err := _MaybeRevertMessageReceiver.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiver) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MaybeRevertMessageReceiver.abi.Events["MessageReceived"].ID:
		return _MaybeRevertMessageReceiver.ParseMessageReceived(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MaybeRevertMessageReceiverMessageReceived) Topic() common.Hash {
	return common.HexToHash("0xd82ce31e3523f6eeb2d24317b2b4133001e8472729657f663b68624c45f8f3e8")
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiver) Address() common.Address {
	return _MaybeRevertMessageReceiver.address
}

type MaybeRevertMessageReceiverInterface interface {
	SToRevert(opts *bind.CallOpts) (bool, error)

	CcipReceive(opts *bind.TransactOpts, arg0 CommonAny2EVMMessage) (*types.Transaction, error)

	SetRevert(opts *bind.TransactOpts, toRevert bool) (*types.Transaction, error)

	FilterMessageReceived(opts *bind.FilterOpts) (*MaybeRevertMessageReceiverMessageReceivedIterator, error)

	WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *MaybeRevertMessageReceiverMessageReceived) (event.Subscription, error)

	ParseMessageReceived(log types.Log) (*MaybeRevertMessageReceiverMessageReceived, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
