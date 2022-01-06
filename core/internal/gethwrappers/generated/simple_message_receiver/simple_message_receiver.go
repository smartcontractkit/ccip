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

type CCIPMessage struct {
	SequenceNumber     *big.Int
	SourceChainId      *big.Int
	DestinationChainId *big.Int
	Sender             common.Address
	Payload            CCIPMessagePayload
}

type CCIPMessagePayload struct {
	Receiver common.Address
	Data     []byte
	Tokens   []common.Address
	Amounts  []*big.Int
	Executor common.Address
	Options  []byte
}

var SimpleMessageReceiverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"receiveMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_message\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610f52806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80633a56bb711461003b578063a82a362a14610050575b600080fd5b61004e610049366004610325565b610072565b005b6100586100bc565b604051610069959493929190610660565b60405180910390f35b80600061007f8282610c59565b9050507f61ac6a56cee73e3d47b52a7e615fea81775723a1760813fdfe55a4f3946c9a51816040516100b191906104a5565b60405180910390a150565b600080546001546002546003546040805160c081019091526004805473ffffffffffffffffffffffffffffffffffffffff90811683526005805497989697959691909416949093602084019161011190610bc3565b80601f016020809104026020016040519081016040528092919081815260200182805461013d90610bc3565b801561018a5780601f1061015f5761010080835404028352916020019161018a565b820191906000526020600020905b81548152906001019060200180831161016d57829003601f168201915b50505050508152602001600282018054806020026020016040519081016040528092919081815260200182805480156101f957602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116101ce575b505050505081526020016003820180548060200260200160405190810160405280929190818152602001828054801561025157602002820191906000526020600020905b81548152602001906001019080831161023d575b5050509183525050600482015473ffffffffffffffffffffffffffffffffffffffff16602082015260058201805460409092019161028e90610bc3565b80601f01602080910402602001604051908101604052809291908181526020018280546102ba90610bc3565b80156103075780601f106102dc57610100808354040283529160200191610307565b820191906000526020600020905b8154815290600101906020018083116102ea57829003601f168201915b505050505081525050905085565b803561032081610f20565b919050565b60006020828403121561033757600080fd5b813567ffffffffffffffff81111561034e57600080fd5b820160a0818503121561036057600080fd5b9392505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83111561039957600080fd5b8260051b8083602087013760009401602001938452509192915050565b600081518084526020808501945080840160005b838110156103e6578151875295820195908201906001016103ca565b509495945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6000815180845260005b8181101561046057602081850181015186830182015201610444565b81811115610472576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600060208083528335818401528084013560408401526040840135606084015260608401356104d381610f20565b73ffffffffffffffffffffffffffffffffffffffff8082166080860152608086013591507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4186360301821261052757600080fd5b60a08581015290850190813561053c81610f20565b811660c086015261054f828401836108ce565b60c060e0880152610565610180880182846103f1565b9150506105756040840184610867565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4088840381016101008a015281845291926000929087015b828410156105d65784356105c081610f20565b86168152938701936001939093019287016105ad565b6105e36060880188610867565b98509550818a8203016101208b01526105fd818988610367565b97505061060c60808701610315565b73ffffffffffffffffffffffffffffffffffffffff81166101408b0152945061063860a08701876108ce565b9650945080898803016101608a0152505050506106568383836103f1565b9695505050505050565b85815260006020868184015285604084015273ffffffffffffffffffffffffffffffffffffffff808616606085015260a060808501528085511660a08501528185015160c0808601526106b761016086018261043a565b60408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878303810160e0890152815180845291860193506000929091908601905b8084101561071d578451861682529386019360019390930192908601906106fb565b5060608901519550818882030161010089015261073a81876103b6565b9550506080880151935061076761012088018573ffffffffffffffffffffffffffffffffffffffff169052565b60a088015193508087860301610140880152505050610786828261043a565b9998505050505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126107c857600080fd5b83018035915067ffffffffffffffff8211156107e357600080fd5b6020019150600581901b36038213156107fb57600080fd5b9250929050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261083757600080fd5b83018035915067ffffffffffffffff82111561085257600080fd5b6020019150368190038213156107fb57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261089c57600080fd5b830160208101925035905067ffffffffffffffff8111156108bc57600080fd5b8060051b36038313156107fb57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261090357600080fd5b830160208101925035905067ffffffffffffffff81111561092357600080fd5b8036038313156107fb57600080fd5b601f82111561096e57806000526020600020601f840160051c810160208510156109595750805b61096b601f850160051c830182610998565b50505b505050565b8183101561096e57806000526020600020610992838201858301610998565b50505050565b5b818110156109ad5760008155600101610999565b5050565b680100000000000000008311156109ca576109ca610c17565b80548382556109da848284610973565b50818160005260208060002060005b86811015610a4c5783356109fc81610f20565b82547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9190911617825592820192600191820191016109e9565b50505050505050565b68010000000000000000831115610a6e57610a6e610c17565b8054838255610a7e848284610973565b50818160005260208060002060005b86811015610a4c57833582559282019260019182019101610a8d565b67ffffffffffffffff831115610ac157610ac1610c17565b610ad583610acf8354610bc3565b83610932565b6000601f841160018114610b275760008515610af15750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b17835561096b565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b82811015610b765786850135825560209485019460019092019101610b56565b5086821015610bb1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b600181811c90821680610bd757607f821691505b60208210811415610c11577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008135610c5381610f20565b92915050565b81358155600160208084013582840155604084013560028401556060840135610c8181610f20565b6003840180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555060808401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41853603018112610cf957600080fd5b84018035610d0681610f20565b6004850180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555060058401610d5983830183610802565b67ffffffffffffffff811115610d7157610d71610c17565b610d8581610d7f8554610bc3565b85610932565b6000601f821160018114610dd75760008315610da15750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178555610e68565b6000858152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0841690835b82811015610e235786850135825593890193908a01908901610e06565b5084821015610e5e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88660031b161c19848701351681555b50508783891b0185555b5050505050610e7a6040820182610793565b93509150610e8c8383600687016109b1565b610e996060820182610793565b93509150610eab838360078701610a55565b610f00610eba60808301610c46565b6008860173ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff00000000000000000000000000000000000000008254161781555050565b610f0d60a0820182610802565b9350915050610992828260098601610aa9565b73ffffffffffffffffffffffffffffffffffffffff81168114610f4257600080fd5b5056fea164736f6c6343000806000a",
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

func (_SimpleMessageReceiver *SimpleMessageReceiverCaller) SMessage(opts *bind.CallOpts) (SMessage,

	error) {
	var out []interface{}
	err := _SimpleMessageReceiver.contract.Call(opts, &out, "s_message")

	outstruct := new(SMessage)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SequenceNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SourceChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DestinationChainId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Sender = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Payload = *abi.ConvertType(out[4], new(CCIPMessagePayload)).(*CCIPMessagePayload)

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

func (_SimpleMessageReceiver *SimpleMessageReceiverTransactor) ReceiveMessage(opts *bind.TransactOpts, message CCIPMessage) (*types.Transaction, error) {
	return _SimpleMessageReceiver.contract.Transact(opts, "receiveMessage", message)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverSession) ReceiveMessage(message CCIPMessage) (*types.Transaction, error) {
	return _SimpleMessageReceiver.Contract.ReceiveMessage(&_SimpleMessageReceiver.TransactOpts, message)
}

func (_SimpleMessageReceiver *SimpleMessageReceiverTransactorSession) ReceiveMessage(message CCIPMessage) (*types.Transaction, error) {
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
	Message CCIPMessage
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
	SequenceNumber     *big.Int
	SourceChainId      *big.Int
	DestinationChainId *big.Int
	Sender             common.Address
	Payload            CCIPMessagePayload
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
	return common.HexToHash("0x61ac6a56cee73e3d47b52a7e615fea81775723a1760813fdfe55a4f3946c9a51")
}

func (_SimpleMessageReceiver *SimpleMessageReceiver) Address() common.Address {
	return _SimpleMessageReceiver.address
}

type SimpleMessageReceiverInterface interface {
	SMessage(opts *bind.CallOpts) (SMessage,

		error)

	ReceiveMessage(opts *bind.TransactOpts, message CCIPMessage) (*types.Transaction, error)

	FilterMessageReceived(opts *bind.FilterOpts) (*SimpleMessageReceiverMessageReceivedIterator, error)

	WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *SimpleMessageReceiverMessageReceived) (event.Subscription, error)

	ParseMessageReceived(log types.Log) (*SimpleMessageReceiverMessageReceived, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
