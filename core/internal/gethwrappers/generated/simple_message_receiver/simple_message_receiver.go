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
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated"
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

type CCIPAny2EVMTollMessage struct {
	SourceChainId  *big.Int
	SequenceNumber uint64
	Sender         common.Address
	Receiver       common.Address
	Data           []byte
	Tokens         []common.Address
	Amounts        []*big.Int
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	GasLimit       *big.Int
}

var SimpleMessageReceiverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"msg\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"receiveMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_message\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611080806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806313d3e05e14610046578063a82a362a1461005b578063ce6d41de14610080575b600080fd5b610059610054366004610473565b610095565b005b6100636100df565b604051610077989796959493929190610521565b60405180910390f35b6100886101de565b6040516100779190610611565b8060006100a28282610b7f565b9050507f64b6aed55f61343abd7d2a5eb9972910673a45be58b6ff6675dc83bb7221d078816040516100d49190610f0d565b60405180910390a150565b6000805460015460025460038054939467ffffffffffffffff8416946801000000000000000090940473ffffffffffffffffffffffffffffffffffffffff9081169493169290919061013090610833565b80601f016020809104026020016040519081016040528092919081815260200182805461015c90610833565b80156101a95780601f1061017e576101008083540402835291602001916101a9565b820191906000526020600020905b81548152906001019060200180831161018c57829003601f168201915b5050505060068301546007840154600890940154929373ffffffffffffffffffffffffffffffffffffffff9091169290915088565b61028060405180610140016040528060008152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b604080516101408101825260008054825260015467ffffffffffffffff8116602084015273ffffffffffffffffffffffffffffffffffffffff680100000000000000009091048116938301939093526002549092166060820152600380549192916080840191906102f090610833565b80601f016020809104026020016040519081016040528092919081815260200182805461031c90610833565b80156103695780601f1061033e57610100808354040283529160200191610369565b820191906000526020600020905b81548152906001019060200180831161034c57829003601f168201915b50505050508152602001600482018054806020026020016040519081016040528092919081815260200182805480156103d857602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116103ad575b505050505081526020016005820180548060200260200160405190810160405280929190818152602001828054801561043057602002820191906000526020600020905b81548152602001906001019080831161041c575b5050509183525050600682015473ffffffffffffffffffffffffffffffffffffffff16602082015260078201546040820152600890910154606090910152919050565b60006020828403121561048557600080fd5b813567ffffffffffffffff81111561049c57600080fd5b820161014081850312156104af57600080fd5b9392505050565b6000815180845260005b818110156104dc576020818501810151868301820152016104c0565b818111156104ee576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60006101008a835267ffffffffffffffff8a16602084015273ffffffffffffffffffffffffffffffffffffffff808a1660408501528089166060850152816080850152610570828501896104b6565b961660a0840152505060c081019290925260e09091015295945050505050565b600081518084526020808501945080840160005b838110156105d657815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016105a4565b509495945050505050565b600081518084526020808501945080840160005b838110156105d6578151875295820195908201906001016105f5565b602081528151602082015260006020830151610639604084018267ffffffffffffffff169052565b50604083015173ffffffffffffffffffffffffffffffffffffffff8116606084015250606083015173ffffffffffffffffffffffffffffffffffffffff811660808401525060808301516101408060a085015261069a6101608501836104b6565b915060a08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808685030160c08701526106d68483610590565b935060c08701519150808685030160e0870152506106f483826105e1565b92505060e08501516101006107208187018373ffffffffffffffffffffffffffffffffffffffff169052565b860151610120868101919091529095015193019290925250919050565b67ffffffffffffffff8116811461075357600080fd5b50565b600081356107638161073d565b92915050565b73ffffffffffffffffffffffffffffffffffffffff8116811461075357600080fd5b6000813561076381610769565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126107cd57600080fd5b83018035915067ffffffffffffffff8211156107e857600080fd5b6020019150368190038213156107fd57600080fd5b9250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c9082168061084757607f821691505b602082108103610880577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b5b8181101561089b5760008155600101610887565b5050565b601f8211156108db57806000526020600020601f840160051c810160208510156108c65750805b6108d8601f850160051c830182610886565b50505b505050565b67ffffffffffffffff8311156108f8576108f8610804565b61090c836109068354610833565b8361089f565b6000601f84116001811461095e57600085156109285750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b1783556108d8565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b828110156109ad578685013582556020948501946001909201910161098d565b50868210156109e8577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610a2f57600080fd5b83018035915067ffffffffffffffff821115610a4a57600080fd5b6020019150600581901b36038213156107fd57600080fd5b818310156108db57806000526020600020610a81838201858301610886565b50505050565b68010000000000000000831115610aa057610aa0610804565b8054838255610ab0848284610a62565b50818160005260208060002060005b86811015610b22578335610ad281610769565b82547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff919091161782559282019260019182019101610abf565b50505050505050565b68010000000000000000831115610b4457610b44610804565b8054838255610b54848284610a62565b50818160005260208060002060005b86811015610b2257833582559282019260019182019101610b63565b8135815560018101610bcc610b9660208501610756565b82547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff91909116178255565b610c28610bdb6040850161078b565b82547fffffffff0000000000000000000000000000000000000000ffffffffffffffff1660409190911b7bffffffffffffffffffffffffffffffffffffffff000000000000000016178255565b50610c7e610c386060840161078b565b6002830173ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff00000000000000000000000000000000000000008254161781555050565b610c8b6080830183610798565b610c998183600386016108e0565b5050610ca860a08301836109fa565b610cb6818360048601610a87565b5050610cc560c08301836109fa565b610cd3818360058601610b2b565b5050610d2a610ce460e0840161078b565b6006830173ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff00000000000000000000000000000000000000008254161781555050565b610100820135600782015561012082013560088201555050565b8035610d4f8161073d565b919050565b8035610d4f81610769565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610d9457600080fd5b830160208101925035905067ffffffffffffffff811115610db457600080fd5b8036038313156107fd57600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610e4157600080fd5b830160208101925035905067ffffffffffffffff811115610e6157600080fd5b8060051b36038313156107fd57600080fd5b8183526000602080850194508260005b858110156105d6578135610e9681610769565b73ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101610e83565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115610ef057600080fd5b8260051b8083602087013760009401602001938452509192915050565b60208152813560208201526000610f2660208401610d44565b67ffffffffffffffff8116604084015250610f4360408401610d54565b73ffffffffffffffffffffffffffffffffffffffff8116606084015250610f6c60608401610d54565b73ffffffffffffffffffffffffffffffffffffffff8116608084015250610f966080840184610d5f565b6101408060a0860152610fae61016086018385610dc3565b9250610fbd60a0870187610e0c565b92507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808786030160c0880152610ff5858584610e73565b945061100460c0890189610e0c565b94509150808786030160e08801525061101e848483610ebe565b93505061102d60e08701610d54565b91506101006110538187018473ffffffffffffffffffffffffffffffffffffffff169052565b86013561012086810191909152909501359490930193909352509091905056fea164736f6c634300080d000a",
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

func (_SimpleMessageReceiver *SimpleMessageReceiverCaller) GetMessage(opts *bind.CallOpts) (CCIPAny2EVMTollMessage, error) {
	var out []interface{}
	err := _SimpleMessageReceiver.contract.Call(opts, &out, "getMessage")

	if err != nil {
		return *new(CCIPAny2EVMTollMessage), err
	}

	out0 := *abi.ConvertType(out[0], new(CCIPAny2EVMTollMessage)).(*CCIPAny2EVMTollMessage)

	return out0, err

}

func (_SimpleMessageReceiver *SimpleMessageReceiverSession) GetMessage() (CCIPAny2EVMTollMessage, error) {
	return _SimpleMessageReceiver.Contract.GetMessage(&_SimpleMessageReceiver.CallOpts)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverCallerSession) GetMessage() (CCIPAny2EVMTollMessage, error) {
	return _SimpleMessageReceiver.Contract.GetMessage(&_SimpleMessageReceiver.CallOpts)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverCaller) SMessage(opts *bind.CallOpts) (SMessage,

	error) {
	var out []interface{}
	err := _SimpleMessageReceiver.contract.Call(opts, &out, "s_message")

	outstruct := new(SMessage)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SequenceNumber = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Sender = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Receiver = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Data = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.FeeToken = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.FeeTokenAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.GasLimit = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_SimpleMessageReceiver *SimpleMessageReceiverSession) SMessage() (SMessage,

	error) {
	return _SimpleMessageReceiver.Contract.SMessage(&_SimpleMessageReceiver.CallOpts)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverCallerSession) SMessage() (SMessage,

	error) {
	return _SimpleMessageReceiver.Contract.SMessage(&_SimpleMessageReceiver.CallOpts)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverTransactor) ReceiveMessage(opts *bind.TransactOpts, message CCIPAny2EVMTollMessage) (*types.Transaction, error) {
	return _SimpleMessageReceiver.contract.Transact(opts, "receiveMessage", message)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverSession) ReceiveMessage(message CCIPAny2EVMTollMessage) (*types.Transaction, error) {
	return _SimpleMessageReceiver.Contract.ReceiveMessage(&_SimpleMessageReceiver.TransactOpts, message)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverTransactorSession) ReceiveMessage(message CCIPAny2EVMTollMessage) (*types.Transaction, error) {
	return _SimpleMessageReceiver.Contract.ReceiveMessage(&_SimpleMessageReceiver.TransactOpts, message)
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
	Message CCIPAny2EVMTollMessage
	Raw     types.Log
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

type SMessage struct {
	SourceChainId  *big.Int
	SequenceNumber uint64
	Sender         common.Address
	Receiver       common.Address
	Data           []byte
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	GasLimit       *big.Int
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
	return common.HexToHash("0x64b6aed55f61343abd7d2a5eb9972910673a45be58b6ff6675dc83bb7221d078")
}

func (_SimpleMessageReceiver *SimpleMessageReceiver) Address() common.Address {
	return _SimpleMessageReceiver.address
}

type SimpleMessageReceiverInterface interface {
	GetMessage(opts *bind.CallOpts) (CCIPAny2EVMTollMessage, error)

	SMessage(opts *bind.CallOpts) (SMessage,

		error)

	ReceiveMessage(opts *bind.TransactOpts, message CCIPAny2EVMTollMessage) (*types.Transaction, error)

	FilterMessageReceived(opts *bind.FilterOpts) (*SimpleMessageReceiverMessageReceivedIterator, error)

	WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *SimpleMessageReceiverMessageReceived) (event.Subscription, error)

	ParseMessageReceived(log types.Log) (*SimpleMessageReceiverMessageReceived, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
