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
	SequenceNumber *big.Int
	SourceChainId  *big.Int
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
	Options            []byte
}

var SimpleMessageReceiverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"receiveMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_message\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610e9e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80639c5984681461003b578063a82a362a14610050575b600080fd5b61004e610049366004610315565b610071565b005b6100586100bb565b60405161006894939291906103fd565b60405180910390f35b80600061007e82826108e7565b9050507f2f76d78617340e21d20e107fdeb8fd7498a1930f7ca15e0a74e80624bf548899816040516100b09190610cab565b60405180910390a150565b600080546001546002546040805160038054610100602082028401810190945260e083018181529697959673ffffffffffffffffffffffffffffffffffffffff90951695949293919284929091849184018282801561015057602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610125575b50505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156101a857602002820191906000526020600020905b815481526020019060010190808311610194575b505050918352505060028201546020820152600382015473ffffffffffffffffffffffffffffffffffffffff908116604083015260048301541660608201526005820180546080909201916101fc90610739565b80601f016020809104026020016040519081016040528092919081815260200182805461022890610739565b80156102755780601f1061024a57610100808354040283529160200191610275565b820191906000526020600020905b81548152906001019060200180831161025857829003601f168201915b5050505050815260200160068201805461028e90610739565b80601f01602080910402602001604051908101604052809291908181526020018280546102ba90610739565b80156103075780601f106102dc57610100808354040283529160200191610307565b820191906000526020600020905b8154815290600101906020018083116102ea57829003601f168201915b505050505081525050905084565b60006020828403121561032757600080fd5b813567ffffffffffffffff81111561033e57600080fd5b82016080818503121561035057600080fd5b9392505050565b600081518084526020808501945080840160005b838110156103875781518752958201959082019060010161036b565b509495945050505050565b6000815180845260005b818110156103b85760208185018101518683018201520161039c565b818111156103ca576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b84815260006020858184015273ffffffffffffffffffffffffffffffffffffffff85166040840152608060608401526101608301845160e06080860152818151808452610180870191508483019350600092505b8083101561048957835173ffffffffffffffffffffffffffffffffffffffff1682526020820191508484019350600183019250610451565b508387015193507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff809250828682030160a08701526104c78185610357565b93505050604085015160c085015260608501516104fc60e086018273ffffffffffffffffffffffffffffffffffffffff169052565b50608085015173ffffffffffffffffffffffffffffffffffffffff1661010085015260a085015184830382016101208601526105388382610392565b92505060c085015181858403016101408601526105558382610392565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461058457600080fd5b50565b6000813561059481610562565b92915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126105cf57600080fd5b83018035915067ffffffffffffffff8211156105ea57600080fd5b6020019150600581901b360382131561060257600080fd5b9250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b5b8181101561064d5760008155600101610639565b5050565b8183101561067257806000526020600020610670838201858301610638565b505b505050565b6801000000000000000083111561069057610690610609565b80548382556106a0848284610651565b50818160005260208060002060005b868110156106cb578335825592820192600191820191016106af565b50505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261070957600080fd5b83018035915067ffffffffffffffff82111561072457600080fd5b60200191503681900382131561060257600080fd5b600181811c9082168061074d57607f821691505b60208210811415610787577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f82111561067257806000526020600020601f840160051c810160208510156107b45750805b6107c6601f850160051c830182610638565b5050505050565b67ffffffffffffffff8311156107e5576107e5610609565b6107f9836107f38354610739565b8361078d565b6000601f84116001811461084b57600085156108155750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b1783556107c6565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b8281101561089a578685013582556020948501946001909201910161087a565b50868210156108d5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b81358155600160208084013582840155604084013561090581610562565b6002840180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316179055506003830160608501357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2186360301811261098157600080fd5b850161098d818061059a565b680100000000000000008111156109a6576109a6610609565b83548185556109b6828287610651565b506000938452848420935b81811015610a235782356109d481610562565b85547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff91909116178555938601939185019186016109c1565b505050610a328382018261059a565b94509250610a44848460048801610677565b60408101356005860155610aa3610a5d60608301610587565b6006870173ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff00000000000000000000000000000000000000008254161781555050565b610af8610ab260808301610587565b6007870173ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff00000000000000000000000000000000000000008254161781555050565b610b0560a08201826106d4565b94509250610b178484600888016107cd565b610b2460c08201826106d4565b9450925050506106708282600986016107cd565b8035610b4381610562565b919050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610b7d57600080fd5b830160208101925035905067ffffffffffffffff811115610b9d57600080fd5b8060051b360383131561060257600080fd5b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115610be157600080fd5b8260051b8083602087013760009401602001938452509192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610c3357600080fd5b830160208101925035905067ffffffffffffffff811115610c5357600080fd5b80360383131561060257600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b600060208083528335818401528084013560408401526040840135610ccf81610562565b73ffffffffffffffffffffffffffffffffffffffff811660608501525060608401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff21853603018112610d2157600080fd5b60808481015284016101808401610d388280610b48565b60e060a088015291829052906000906101a087015b81831015610d8b578335610d6081610562565b73ffffffffffffffffffffffffffffffffffffffff1681529285019260019290920191602001610d4d565b610d9786860186610b48565b965093507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff609250828882030160c0890152610dd3818786610baf565b95505050604083013560e0870152610ded60608401610b38565b73ffffffffffffffffffffffffffffffffffffffff81166101008801529150610e1860808401610b38565b73ffffffffffffffffffffffffffffffffffffffff81166101208801529150610e4460a0840184610bfe565b92508187860301610140880152610e5c858483610c62565b945050610e6c60c0840184610bfe565b93509150808685030161016087015250610e87838383610c62565b969550505050505056fea164736f6c634300080c000a",
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
	SequenceNumber *big.Int
	SourceChainId  *big.Int
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
	return common.HexToHash("0x2f76d78617340e21d20e107fdeb8fd7498a1930f7ca15e0a74e80624bf548899")
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
