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
	SourceChainId  *big.Int
	SequenceNumber uint64
	Sender         common.Address
	Payload        CCIPMessagePayload
}

type CCIPMessagePayload struct {
	Tokens             []common.Address
	Amounts            []*big.Int
	DestinationChainId *big.Int
	Receiver           common.Address
	Executor           common.Address
	Data               []byte
}

var SimpleMessageReceiverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"receiveMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_message\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610e09806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806302b5e9fa1461003b578063a82a362a14610050575b600080fd5b61004e610049366004610290565b610071565b005b6100586100bb565b6040516100689493929190610378565b60405180910390f35b80600061007e8282610864565b9050507f5e1b08f8d3326c43d00b511ea7532b721cbd06679cf8aad01ade2ce0a9d4a080816040516100b09190610c29565b60405180910390a150565b60008054600154604080516002805460208102830160e090810190945260c08301818152959667ffffffffffffffff8616966801000000000000000090960473ffffffffffffffffffffffffffffffffffffffff169590948492849184018282801561015d57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610132575b50505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156101b557602002820191906000526020600020905b8154815260200190600101908083116101a1575b505050918352505060028201546020820152600382015473ffffffffffffffffffffffffffffffffffffffff90811660408301526004830154166060820152600582018054608090920191610209906106b7565b80601f0160208091040260200160405190810160405280929190818152602001828054610235906106b7565b80156102825780601f1061025757610100808354040283529160200191610282565b820191906000526020600020905b81548152906001019060200180831161026557829003601f168201915b505050505081525050905084565b6000602082840312156102a257600080fd5b813567ffffffffffffffff8111156102b957600080fd5b8201608081850312156102cb57600080fd5b9392505050565b600081518084526020808501945080840160005b83811015610302578151875295820195908201906001016102e6565b509495945050505050565b6000815180845260005b8181101561033357602081850181015186830182015201610317565b81811115610345576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b8481526000602067ffffffffffffffff86168184015273ffffffffffffffffffffffffffffffffffffffff85166040840152608060608401526101408301845160c06080860152818151808452610160870191508483019350600092505b8083101561040e57835173ffffffffffffffffffffffffffffffffffffffff16825260208201915084840193506001830192506103d6565b508387015193507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff809250828682030160a087015261044c81856102d2565b93505050604085015160c0850152606085015161048160e086018273ffffffffffffffffffffffffffffffffffffffff169052565b50608085015173ffffffffffffffffffffffffffffffffffffffff1661010085015260a085015184830382016101208601526104bd838261030d565b9998505050505050505050565b67ffffffffffffffff811681146104e057600080fd5b50565b73ffffffffffffffffffffffffffffffffffffffff811681146104e057600080fd5b60008135610512816104e3565b92915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261054d57600080fd5b83018035915067ffffffffffffffff82111561056857600080fd5b6020019150600581901b360382131561058057600080fd5b9250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b5b818110156105cb57600081556001016105b7565b5050565b818310156105f0578060005260206000206105ee8382018583016105b6565b505b505050565b6801000000000000000083111561060e5761060e610587565b805483825561061e8482846105cf565b50818160005260208060002060005b868110156106495783358255928201926001918201910161062d565b50505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261068757600080fd5b83018035915067ffffffffffffffff8211156106a257600080fd5b60200191503681900382131561058057600080fd5b600181811c908216806106cb57607f821691505b602082108103610704577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f8211156105f057806000526020600020601f840160051c810160208510156107315750805b610743601f850160051c8301826105b6565b5050505050565b67ffffffffffffffff83111561076257610762610587565b6107768361077083546106b7565b8361070a565b6000601f8411600181146107c857600085156107925750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355610743565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b8281101561081757868501358255602094850194600190920191016107f7565b5086821015610852577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b81358155600180820160208085013561087c816104ca565b8254604087013561088c816104e3565b7bffffffffffffffffffffffffffffffffffffffff00000000000000008160401b1667ffffffffffffffff84167fffffffff0000000000000000000000000000000000000000000000000000000084161717855550505060028401915060608501357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4186360301811261091e57600080fd5b850161092a8180610518565b6801000000000000000081111561094357610943610587565b84548186556109538282886105cf565b506000948552838520945b818110156109c0578235610971816104e3565b86547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff919091161786559486019491840191860161095e565b5050506109cf82820182610518565b945092506109e18484600388016105f5565b60408101356004860155610a406109fa60608301610505565b6005870173ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff00000000000000000000000000000000000000008254161781555050565b610a95610a4f60808301610505565b6006870173ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff00000000000000000000000000000000000000008254161781555050565b610aa260a0820182610652565b9450925050506105ee82826007860161074a565b8035610ac1816104e3565b919050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610afb57600080fd5b830160208101925035905067ffffffffffffffff811115610b1b57600080fd5b8060051b360383131561058057600080fd5b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115610b5f57600080fd5b8260051b8083602087013760009401602001938452509192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610bb157600080fd5b830160208101925035905067ffffffffffffffff811115610bd157600080fd5b80360383131561058057600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6000602080835283358184015280840135610c43816104ca565b67ffffffffffffffff81166040850152506040840135610c62816104e3565b73ffffffffffffffffffffffffffffffffffffffff811660608501525060608401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41853603018112610cb457600080fd5b60808481015284016101608401610ccb8280610ac6565b60c060a0880152918290529060009061018087015b81831015610d1e578335610cf3816104e3565b73ffffffffffffffffffffffffffffffffffffffff1681529285019260019290920191602001610ce0565b610d2a86860186610ac6565b965093507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff609250828882030160c0890152610d66818786610b2d565b95505050604083013560e0870152610d8060608401610ab6565b73ffffffffffffffffffffffffffffffffffffffff81166101008801529150610dab60808401610ab6565b73ffffffffffffffffffffffffffffffffffffffff81166101208801529150610dd760a0840184610b7c565b93509150808685030161014087015250610df2838383610be0565b969550505050505056fea164736f6c634300080d000a",
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

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SequenceNumber = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Sender = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Payload = *abi.ConvertType(out[3], new(CCIPMessagePayload)).(*CCIPMessagePayload)

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
	SourceChainId  *big.Int
	SequenceNumber uint64
	Sender         common.Address
	Payload        CCIPMessagePayload
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
	return common.HexToHash("0x5e1b08f8d3326c43d00b511ea7532b721cbd06679cf8aad01ade2ce0a9d4a080")
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
