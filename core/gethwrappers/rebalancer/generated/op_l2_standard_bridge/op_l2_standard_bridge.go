// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package op_l2_standard_bridge

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

var OpL2StandardBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"DepositFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"DepositFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1TokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l1Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l1Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var OpL2StandardBridgeABI = OpL2StandardBridgeMetaData.ABI

type OpL2StandardBridge struct {
	address common.Address
	abi     abi.ABI
	OpL2StandardBridgeCaller
	OpL2StandardBridgeTransactor
	OpL2StandardBridgeFilterer
}

type OpL2StandardBridgeCaller struct {
	contract *bind.BoundContract
}

type OpL2StandardBridgeTransactor struct {
	contract *bind.BoundContract
}

type OpL2StandardBridgeFilterer struct {
	contract *bind.BoundContract
}

type OpL2StandardBridgeSession struct {
	Contract     *OpL2StandardBridge
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OpL2StandardBridgeCallerSession struct {
	Contract *OpL2StandardBridgeCaller
	CallOpts bind.CallOpts
}

type OpL2StandardBridgeTransactorSession struct {
	Contract     *OpL2StandardBridgeTransactor
	TransactOpts bind.TransactOpts
}

type OpL2StandardBridgeRaw struct {
	Contract *OpL2StandardBridge
}

type OpL2StandardBridgeCallerRaw struct {
	Contract *OpL2StandardBridgeCaller
}

type OpL2StandardBridgeTransactorRaw struct {
	Contract *OpL2StandardBridgeTransactor
}

func NewOpL2StandardBridge(address common.Address, backend bind.ContractBackend) (*OpL2StandardBridge, error) {
	abi, err := abi.JSON(strings.NewReader(OpL2StandardBridgeABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOpL2StandardBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpL2StandardBridge{address: address, abi: abi, OpL2StandardBridgeCaller: OpL2StandardBridgeCaller{contract: contract}, OpL2StandardBridgeTransactor: OpL2StandardBridgeTransactor{contract: contract}, OpL2StandardBridgeFilterer: OpL2StandardBridgeFilterer{contract: contract}}, nil
}

func NewOpL2StandardBridgeCaller(address common.Address, caller bind.ContractCaller) (*OpL2StandardBridgeCaller, error) {
	contract, err := bindOpL2StandardBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpL2StandardBridgeCaller{contract: contract}, nil
}

func NewOpL2StandardBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*OpL2StandardBridgeTransactor, error) {
	contract, err := bindOpL2StandardBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpL2StandardBridgeTransactor{contract: contract}, nil
}

func NewOpL2StandardBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*OpL2StandardBridgeFilterer, error) {
	contract, err := bindOpL2StandardBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpL2StandardBridgeFilterer{contract: contract}, nil
}

func bindOpL2StandardBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OpL2StandardBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OpL2StandardBridge *OpL2StandardBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpL2StandardBridge.Contract.OpL2StandardBridgeCaller.contract.Call(opts, result, method, params...)
}

func (_OpL2StandardBridge *OpL2StandardBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpL2StandardBridge.Contract.OpL2StandardBridgeTransactor.contract.Transfer(opts)
}

func (_OpL2StandardBridge *OpL2StandardBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpL2StandardBridge.Contract.OpL2StandardBridgeTransactor.contract.Transact(opts, method, params...)
}

func (_OpL2StandardBridge *OpL2StandardBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpL2StandardBridge.Contract.contract.Call(opts, result, method, params...)
}

func (_OpL2StandardBridge *OpL2StandardBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpL2StandardBridge.Contract.contract.Transfer(opts)
}

func (_OpL2StandardBridge *OpL2StandardBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpL2StandardBridge.Contract.contract.Transact(opts, method, params...)
}

func (_OpL2StandardBridge *OpL2StandardBridgeCaller) L1TokenBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpL2StandardBridge.contract.Call(opts, &out, "l1TokenBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OpL2StandardBridge *OpL2StandardBridgeSession) L1TokenBridge() (common.Address, error) {
	return _OpL2StandardBridge.Contract.L1TokenBridge(&_OpL2StandardBridge.CallOpts)
}

func (_OpL2StandardBridge *OpL2StandardBridgeCallerSession) L1TokenBridge() (common.Address, error) {
	return _OpL2StandardBridge.Contract.L1TokenBridge(&_OpL2StandardBridge.CallOpts)
}

func (_OpL2StandardBridge *OpL2StandardBridgeCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpL2StandardBridge.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OpL2StandardBridge *OpL2StandardBridgeSession) Messenger() (common.Address, error) {
	return _OpL2StandardBridge.Contract.Messenger(&_OpL2StandardBridge.CallOpts)
}

func (_OpL2StandardBridge *OpL2StandardBridgeCallerSession) Messenger() (common.Address, error) {
	return _OpL2StandardBridge.Contract.Messenger(&_OpL2StandardBridge.CallOpts)
}

func (_OpL2StandardBridge *OpL2StandardBridgeTransactor) FinalizeDeposit(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _OpL2StandardBridge.contract.Transact(opts, "finalizeDeposit", _l1Token, _l2Token, _from, _to, _amount, _data)
}

func (_OpL2StandardBridge *OpL2StandardBridgeSession) FinalizeDeposit(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _OpL2StandardBridge.Contract.FinalizeDeposit(&_OpL2StandardBridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

func (_OpL2StandardBridge *OpL2StandardBridgeTransactorSession) FinalizeDeposit(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _OpL2StandardBridge.Contract.FinalizeDeposit(&_OpL2StandardBridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

func (_OpL2StandardBridge *OpL2StandardBridgeTransactor) Withdraw(opts *bind.TransactOpts, _l2Token common.Address, _amount *big.Int, _l1Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OpL2StandardBridge.contract.Transact(opts, "withdraw", _l2Token, _amount, _l1Gas, _data)
}

func (_OpL2StandardBridge *OpL2StandardBridgeSession) Withdraw(_l2Token common.Address, _amount *big.Int, _l1Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OpL2StandardBridge.Contract.Withdraw(&_OpL2StandardBridge.TransactOpts, _l2Token, _amount, _l1Gas, _data)
}

func (_OpL2StandardBridge *OpL2StandardBridgeTransactorSession) Withdraw(_l2Token common.Address, _amount *big.Int, _l1Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OpL2StandardBridge.Contract.Withdraw(&_OpL2StandardBridge.TransactOpts, _l2Token, _amount, _l1Gas, _data)
}

func (_OpL2StandardBridge *OpL2StandardBridgeTransactor) WithdrawTo(opts *bind.TransactOpts, _l2Token common.Address, _to common.Address, _amount *big.Int, _l1Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OpL2StandardBridge.contract.Transact(opts, "withdrawTo", _l2Token, _to, _amount, _l1Gas, _data)
}

func (_OpL2StandardBridge *OpL2StandardBridgeSession) WithdrawTo(_l2Token common.Address, _to common.Address, _amount *big.Int, _l1Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OpL2StandardBridge.Contract.WithdrawTo(&_OpL2StandardBridge.TransactOpts, _l2Token, _to, _amount, _l1Gas, _data)
}

func (_OpL2StandardBridge *OpL2StandardBridgeTransactorSession) WithdrawTo(_l2Token common.Address, _to common.Address, _amount *big.Int, _l1Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OpL2StandardBridge.Contract.WithdrawTo(&_OpL2StandardBridge.TransactOpts, _l2Token, _to, _amount, _l1Gas, _data)
}

type OpL2StandardBridgeDepositFailedIterator struct {
	Event *OpL2StandardBridgeDepositFailed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OpL2StandardBridgeDepositFailedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpL2StandardBridgeDepositFailed)
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
		it.Event = new(OpL2StandardBridgeDepositFailed)
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

func (it *OpL2StandardBridgeDepositFailedIterator) Error() error {
	return it.fail
}

func (it *OpL2StandardBridgeDepositFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OpL2StandardBridgeDepositFailed struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log
}

func (_OpL2StandardBridge *OpL2StandardBridgeFilterer) FilterDepositFailed(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*OpL2StandardBridgeDepositFailedIterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _OpL2StandardBridge.contract.FilterLogs(opts, "DepositFailed", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &OpL2StandardBridgeDepositFailedIterator{contract: _OpL2StandardBridge.contract, event: "DepositFailed", logs: logs, sub: sub}, nil
}

func (_OpL2StandardBridge *OpL2StandardBridgeFilterer) WatchDepositFailed(opts *bind.WatchOpts, sink chan<- *OpL2StandardBridgeDepositFailed, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _OpL2StandardBridge.contract.WatchLogs(opts, "DepositFailed", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OpL2StandardBridgeDepositFailed)
				if err := _OpL2StandardBridge.contract.UnpackLog(event, "DepositFailed", log); err != nil {
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

func (_OpL2StandardBridge *OpL2StandardBridgeFilterer) ParseDepositFailed(log types.Log) (*OpL2StandardBridgeDepositFailed, error) {
	event := new(OpL2StandardBridgeDepositFailed)
	if err := _OpL2StandardBridge.contract.UnpackLog(event, "DepositFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OpL2StandardBridgeDepositFinalizedIterator struct {
	Event *OpL2StandardBridgeDepositFinalized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OpL2StandardBridgeDepositFinalizedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpL2StandardBridgeDepositFinalized)
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
		it.Event = new(OpL2StandardBridgeDepositFinalized)
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

func (it *OpL2StandardBridgeDepositFinalizedIterator) Error() error {
	return it.fail
}

func (it *OpL2StandardBridgeDepositFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OpL2StandardBridgeDepositFinalized struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log
}

func (_OpL2StandardBridge *OpL2StandardBridgeFilterer) FilterDepositFinalized(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*OpL2StandardBridgeDepositFinalizedIterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _OpL2StandardBridge.contract.FilterLogs(opts, "DepositFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &OpL2StandardBridgeDepositFinalizedIterator{contract: _OpL2StandardBridge.contract, event: "DepositFinalized", logs: logs, sub: sub}, nil
}

func (_OpL2StandardBridge *OpL2StandardBridgeFilterer) WatchDepositFinalized(opts *bind.WatchOpts, sink chan<- *OpL2StandardBridgeDepositFinalized, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _OpL2StandardBridge.contract.WatchLogs(opts, "DepositFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OpL2StandardBridgeDepositFinalized)
				if err := _OpL2StandardBridge.contract.UnpackLog(event, "DepositFinalized", log); err != nil {
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

func (_OpL2StandardBridge *OpL2StandardBridgeFilterer) ParseDepositFinalized(log types.Log) (*OpL2StandardBridgeDepositFinalized, error) {
	event := new(OpL2StandardBridgeDepositFinalized)
	if err := _OpL2StandardBridge.contract.UnpackLog(event, "DepositFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OpL2StandardBridgeERC20BridgeFinalizedIterator struct {
	Event *OpL2StandardBridgeERC20BridgeFinalized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OpL2StandardBridgeERC20BridgeFinalizedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpL2StandardBridgeERC20BridgeFinalized)
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
		it.Event = new(OpL2StandardBridgeERC20BridgeFinalized)
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

func (it *OpL2StandardBridgeERC20BridgeFinalizedIterator) Error() error {
	return it.fail
}

func (it *OpL2StandardBridgeERC20BridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OpL2StandardBridgeERC20BridgeFinalized struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         types.Log
}

func (_OpL2StandardBridge *OpL2StandardBridgeFilterer) FilterERC20BridgeFinalized(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*OpL2StandardBridgeERC20BridgeFinalizedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OpL2StandardBridge.contract.FilterLogs(opts, "ERC20BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &OpL2StandardBridgeERC20BridgeFinalizedIterator{contract: _OpL2StandardBridge.contract, event: "ERC20BridgeFinalized", logs: logs, sub: sub}, nil
}

func (_OpL2StandardBridge *OpL2StandardBridgeFilterer) WatchERC20BridgeFinalized(opts *bind.WatchOpts, sink chan<- *OpL2StandardBridgeERC20BridgeFinalized, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OpL2StandardBridge.contract.WatchLogs(opts, "ERC20BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OpL2StandardBridgeERC20BridgeFinalized)
				if err := _OpL2StandardBridge.contract.UnpackLog(event, "ERC20BridgeFinalized", log); err != nil {
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

func (_OpL2StandardBridge *OpL2StandardBridgeFilterer) ParseERC20BridgeFinalized(log types.Log) (*OpL2StandardBridgeERC20BridgeFinalized, error) {
	event := new(OpL2StandardBridgeERC20BridgeFinalized)
	if err := _OpL2StandardBridge.contract.UnpackLog(event, "ERC20BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OpL2StandardBridgeWithdrawalInitiatedIterator struct {
	Event *OpL2StandardBridgeWithdrawalInitiated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OpL2StandardBridgeWithdrawalInitiatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpL2StandardBridgeWithdrawalInitiated)
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
		it.Event = new(OpL2StandardBridgeWithdrawalInitiated)
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

func (it *OpL2StandardBridgeWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

func (it *OpL2StandardBridgeWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OpL2StandardBridgeWithdrawalInitiated struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log
}

func (_OpL2StandardBridge *OpL2StandardBridgeFilterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*OpL2StandardBridgeWithdrawalInitiatedIterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _OpL2StandardBridge.contract.FilterLogs(opts, "WithdrawalInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &OpL2StandardBridgeWithdrawalInitiatedIterator{contract: _OpL2StandardBridge.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

func (_OpL2StandardBridge *OpL2StandardBridgeFilterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *OpL2StandardBridgeWithdrawalInitiated, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _OpL2StandardBridge.contract.WatchLogs(opts, "WithdrawalInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OpL2StandardBridgeWithdrawalInitiated)
				if err := _OpL2StandardBridge.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
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

func (_OpL2StandardBridge *OpL2StandardBridgeFilterer) ParseWithdrawalInitiated(log types.Log) (*OpL2StandardBridgeWithdrawalInitiated, error) {
	event := new(OpL2StandardBridgeWithdrawalInitiated)
	if err := _OpL2StandardBridge.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_OpL2StandardBridge *OpL2StandardBridge) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OpL2StandardBridge.abi.Events["DepositFailed"].ID:
		return _OpL2StandardBridge.ParseDepositFailed(log)
	case _OpL2StandardBridge.abi.Events["DepositFinalized"].ID:
		return _OpL2StandardBridge.ParseDepositFinalized(log)
	case _OpL2StandardBridge.abi.Events["ERC20BridgeFinalized"].ID:
		return _OpL2StandardBridge.ParseERC20BridgeFinalized(log)
	case _OpL2StandardBridge.abi.Events["WithdrawalInitiated"].ID:
		return _OpL2StandardBridge.ParseWithdrawalInitiated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (OpL2StandardBridgeDepositFailed) Topic() common.Hash {
	return common.HexToHash("0x7ea89a4591614515571c2b51f5ea06494056f261c10ab1ed8c03c7590d87bce0")
}

func (OpL2StandardBridgeDepositFinalized) Topic() common.Hash {
	return common.HexToHash("0xb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89")
}

func (OpL2StandardBridgeERC20BridgeFinalized) Topic() common.Hash {
	return common.HexToHash("0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd")
}

func (OpL2StandardBridgeWithdrawalInitiated) Topic() common.Hash {
	return common.HexToHash("0x73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e")
}

func (_OpL2StandardBridge *OpL2StandardBridge) Address() common.Address {
	return _OpL2StandardBridge.address
}

type OpL2StandardBridgeInterface interface {
	L1TokenBridge(opts *bind.CallOpts) (common.Address, error)

	Messenger(opts *bind.CallOpts) (common.Address, error)

	FinalizeDeposit(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, _l2Token common.Address, _amount *big.Int, _l1Gas uint32, _data []byte) (*types.Transaction, error)

	WithdrawTo(opts *bind.TransactOpts, _l2Token common.Address, _to common.Address, _amount *big.Int, _l1Gas uint32, _data []byte) (*types.Transaction, error)

	FilterDepositFailed(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*OpL2StandardBridgeDepositFailedIterator, error)

	WatchDepositFailed(opts *bind.WatchOpts, sink chan<- *OpL2StandardBridgeDepositFailed, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error)

	ParseDepositFailed(log types.Log) (*OpL2StandardBridgeDepositFailed, error)

	FilterDepositFinalized(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*OpL2StandardBridgeDepositFinalizedIterator, error)

	WatchDepositFinalized(opts *bind.WatchOpts, sink chan<- *OpL2StandardBridgeDepositFinalized, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error)

	ParseDepositFinalized(log types.Log) (*OpL2StandardBridgeDepositFinalized, error)

	FilterERC20BridgeFinalized(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*OpL2StandardBridgeERC20BridgeFinalizedIterator, error)

	WatchERC20BridgeFinalized(opts *bind.WatchOpts, sink chan<- *OpL2StandardBridgeERC20BridgeFinalized, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error)

	ParseERC20BridgeFinalized(log types.Log) (*OpL2StandardBridgeERC20BridgeFinalized, error)

	FilterWithdrawalInitiated(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*OpL2StandardBridgeWithdrawalInitiatedIterator, error)

	WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *OpL2StandardBridgeWithdrawalInitiated, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error)

	ParseWithdrawalInitiated(log types.Log) (*OpL2StandardBridgeWithdrawalInitiated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
