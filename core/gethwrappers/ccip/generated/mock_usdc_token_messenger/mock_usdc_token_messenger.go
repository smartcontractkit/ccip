// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_usdc_token_messenger

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

var MockUSDCTokenMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destinationTokenMessenger\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"}],\"name\":\"DepositForBurn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"}],\"name\":\"depositForBurnWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_destinationTokenMessenger\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localMessageTransmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBodyVersion\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_nonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161053638038061053683398101604081905261002f9161006e565b63ffffffff909116608052600080546001600160a01b0390921668010000000000000000026001600160e01b03199092169190911760011790556100bd565b6000806040838503121561008157600080fd5b825163ffffffff8116811461009557600080fd5b60208401519092506001600160a01b03811681146100b257600080fd5b809150509250929050565b60805161045e6100d86000396000610127015261045e6000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80637eccf63e116100505780637eccf63e146100f05780639cdbb1811461011d578063f856ddb61461015157600080fd5b80632c1219211461006c5780636665a0bb146100bb575b600080fd5b60005468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1660405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100e27f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f681565b6040519081526020016100b2565b6000546101049067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016100b2565b60405163ffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526020016100b2565b61010461015f366004610366565b6040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810186905260009073ffffffffffffffffffffffffffffffffffffffff8416906323b872dd906064016020604051808303816000875af11580156101db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101ff91906103da565b506040517f42966c680000000000000000000000000000000000000000000000000000000081526004810187905273ffffffffffffffffffffffffffffffffffffffff8416906342966c6890602401600060405180830381600087803b15801561026857600080fd5b505af115801561027c573d6000803e3d6000fd5b5050600054604080518a81526020810189905263ffffffff8a16818301527f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f6606082015260808101879052905133945073ffffffffffffffffffffffffffffffffffffffff8816935067ffffffffffffffff909216917f2fa9ca894982930190727e75500a97d8dc500233a5065e0f3126c48fbe0343c09181900360a00190a46000805467ffffffffffffffff16908061033583610403565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550905095945050505050565b600080600080600060a0868803121561037e57600080fd5b85359450602086013563ffffffff8116811461039957600080fd5b935060408601359250606086013573ffffffffffffffffffffffffffffffffffffffff811681146103c957600080fd5b949793965091946080013592915050565b6000602082840312156103ec57600080fd5b815180151581146103fc57600080fd5b9392505050565b600067ffffffffffffffff808316818103610447577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600101939250505056fea164736f6c6343000813000a",
}

var MockUSDCTokenMessengerABI = MockUSDCTokenMessengerMetaData.ABI

var MockUSDCTokenMessengerBin = MockUSDCTokenMessengerMetaData.Bin

func DeployMockUSDCTokenMessenger(auth *bind.TransactOpts, backend bind.ContractBackend, version uint32, transmitter common.Address) (common.Address, *types.Transaction, *MockUSDCTokenMessenger, error) {
	parsed, err := MockUSDCTokenMessengerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockUSDCTokenMessengerBin), backend, version, transmitter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockUSDCTokenMessenger{address: address, abi: *parsed, MockUSDCTokenMessengerCaller: MockUSDCTokenMessengerCaller{contract: contract}, MockUSDCTokenMessengerTransactor: MockUSDCTokenMessengerTransactor{contract: contract}, MockUSDCTokenMessengerFilterer: MockUSDCTokenMessengerFilterer{contract: contract}}, nil
}

type MockUSDCTokenMessenger struct {
	address common.Address
	abi     abi.ABI
	MockUSDCTokenMessengerCaller
	MockUSDCTokenMessengerTransactor
	MockUSDCTokenMessengerFilterer
}

type MockUSDCTokenMessengerCaller struct {
	contract *bind.BoundContract
}

type MockUSDCTokenMessengerTransactor struct {
	contract *bind.BoundContract
}

type MockUSDCTokenMessengerFilterer struct {
	contract *bind.BoundContract
}

type MockUSDCTokenMessengerSession struct {
	Contract     *MockUSDCTokenMessenger
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockUSDCTokenMessengerCallerSession struct {
	Contract *MockUSDCTokenMessengerCaller
	CallOpts bind.CallOpts
}

type MockUSDCTokenMessengerTransactorSession struct {
	Contract     *MockUSDCTokenMessengerTransactor
	TransactOpts bind.TransactOpts
}

type MockUSDCTokenMessengerRaw struct {
	Contract *MockUSDCTokenMessenger
}

type MockUSDCTokenMessengerCallerRaw struct {
	Contract *MockUSDCTokenMessengerCaller
}

type MockUSDCTokenMessengerTransactorRaw struct {
	Contract *MockUSDCTokenMessengerTransactor
}

func NewMockUSDCTokenMessenger(address common.Address, backend bind.ContractBackend) (*MockUSDCTokenMessenger, error) {
	abi, err := abi.JSON(strings.NewReader(MockUSDCTokenMessengerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockUSDCTokenMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockUSDCTokenMessenger{address: address, abi: abi, MockUSDCTokenMessengerCaller: MockUSDCTokenMessengerCaller{contract: contract}, MockUSDCTokenMessengerTransactor: MockUSDCTokenMessengerTransactor{contract: contract}, MockUSDCTokenMessengerFilterer: MockUSDCTokenMessengerFilterer{contract: contract}}, nil
}

func NewMockUSDCTokenMessengerCaller(address common.Address, caller bind.ContractCaller) (*MockUSDCTokenMessengerCaller, error) {
	contract, err := bindMockUSDCTokenMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockUSDCTokenMessengerCaller{contract: contract}, nil
}

func NewMockUSDCTokenMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*MockUSDCTokenMessengerTransactor, error) {
	contract, err := bindMockUSDCTokenMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockUSDCTokenMessengerTransactor{contract: contract}, nil
}

func NewMockUSDCTokenMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*MockUSDCTokenMessengerFilterer, error) {
	contract, err := bindMockUSDCTokenMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockUSDCTokenMessengerFilterer{contract: contract}, nil
}

func bindMockUSDCTokenMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockUSDCTokenMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockUSDCTokenMessenger.Contract.MockUSDCTokenMessengerCaller.contract.Call(opts, result, method, params...)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockUSDCTokenMessenger.Contract.MockUSDCTokenMessengerTransactor.contract.Transfer(opts)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockUSDCTokenMessenger.Contract.MockUSDCTokenMessengerTransactor.contract.Transact(opts, method, params...)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockUSDCTokenMessenger.Contract.contract.Call(opts, result, method, params...)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockUSDCTokenMessenger.Contract.contract.Transfer(opts)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockUSDCTokenMessenger.Contract.contract.Transact(opts, method, params...)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerCaller) IDestinationTokenMessenger(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockUSDCTokenMessenger.contract.Call(opts, &out, "i_destinationTokenMessenger")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerSession) IDestinationTokenMessenger() ([32]byte, error) {
	return _MockUSDCTokenMessenger.Contract.IDestinationTokenMessenger(&_MockUSDCTokenMessenger.CallOpts)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerCallerSession) IDestinationTokenMessenger() ([32]byte, error) {
	return _MockUSDCTokenMessenger.Contract.IDestinationTokenMessenger(&_MockUSDCTokenMessenger.CallOpts)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerCaller) LocalMessageTransmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockUSDCTokenMessenger.contract.Call(opts, &out, "localMessageTransmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerSession) LocalMessageTransmitter() (common.Address, error) {
	return _MockUSDCTokenMessenger.Contract.LocalMessageTransmitter(&_MockUSDCTokenMessenger.CallOpts)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerCallerSession) LocalMessageTransmitter() (common.Address, error) {
	return _MockUSDCTokenMessenger.Contract.LocalMessageTransmitter(&_MockUSDCTokenMessenger.CallOpts)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerCaller) MessageBodyVersion(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MockUSDCTokenMessenger.contract.Call(opts, &out, "messageBodyVersion")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerSession) MessageBodyVersion() (uint32, error) {
	return _MockUSDCTokenMessenger.Contract.MessageBodyVersion(&_MockUSDCTokenMessenger.CallOpts)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerCallerSession) MessageBodyVersion() (uint32, error) {
	return _MockUSDCTokenMessenger.Contract.MessageBodyVersion(&_MockUSDCTokenMessenger.CallOpts)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerCaller) SNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MockUSDCTokenMessenger.contract.Call(opts, &out, "s_nonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerSession) SNonce() (uint64, error) {
	return _MockUSDCTokenMessenger.Contract.SNonce(&_MockUSDCTokenMessenger.CallOpts)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerCallerSession) SNonce() (uint64, error) {
	return _MockUSDCTokenMessenger.Contract.SNonce(&_MockUSDCTokenMessenger.CallOpts)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerTransactor) DepositForBurnWithCaller(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _MockUSDCTokenMessenger.contract.Transact(opts, "depositForBurnWithCaller", amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerSession) DepositForBurnWithCaller(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _MockUSDCTokenMessenger.Contract.DepositForBurnWithCaller(&_MockUSDCTokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerTransactorSession) DepositForBurnWithCaller(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _MockUSDCTokenMessenger.Contract.DepositForBurnWithCaller(&_MockUSDCTokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

type MockUSDCTokenMessengerDepositForBurnIterator struct {
	Event *MockUSDCTokenMessengerDepositForBurn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockUSDCTokenMessengerDepositForBurnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUSDCTokenMessengerDepositForBurn)
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
		it.Event = new(MockUSDCTokenMessengerDepositForBurn)
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

func (it *MockUSDCTokenMessengerDepositForBurnIterator) Error() error {
	return it.fail
}

func (it *MockUSDCTokenMessengerDepositForBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockUSDCTokenMessengerDepositForBurn struct {
	Nonce                     uint64
	BurnToken                 common.Address
	Amount                    *big.Int
	Depositor                 common.Address
	MintRecipient             [32]byte
	DestinationDomain         uint32
	DestinationTokenMessenger [32]byte
	DestinationCaller         [32]byte
	Raw                       types.Log
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerFilterer) FilterDepositForBurn(opts *bind.FilterOpts, nonce []uint64, burnToken []common.Address, depositor []common.Address) (*MockUSDCTokenMessengerDepositForBurnIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var burnTokenRule []interface{}
	for _, burnTokenItem := range burnToken {
		burnTokenRule = append(burnTokenRule, burnTokenItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _MockUSDCTokenMessenger.contract.FilterLogs(opts, "DepositForBurn", nonceRule, burnTokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &MockUSDCTokenMessengerDepositForBurnIterator{contract: _MockUSDCTokenMessenger.contract, event: "DepositForBurn", logs: logs, sub: sub}, nil
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerFilterer) WatchDepositForBurn(opts *bind.WatchOpts, sink chan<- *MockUSDCTokenMessengerDepositForBurn, nonce []uint64, burnToken []common.Address, depositor []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var burnTokenRule []interface{}
	for _, burnTokenItem := range burnToken {
		burnTokenRule = append(burnTokenRule, burnTokenItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _MockUSDCTokenMessenger.contract.WatchLogs(opts, "DepositForBurn", nonceRule, burnTokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockUSDCTokenMessengerDepositForBurn)
				if err := _MockUSDCTokenMessenger.contract.UnpackLog(event, "DepositForBurn", log); err != nil {
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

func (_MockUSDCTokenMessenger *MockUSDCTokenMessengerFilterer) ParseDepositForBurn(log types.Log) (*MockUSDCTokenMessengerDepositForBurn, error) {
	event := new(MockUSDCTokenMessengerDepositForBurn)
	if err := _MockUSDCTokenMessenger.contract.UnpackLog(event, "DepositForBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessenger) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MockUSDCTokenMessenger.abi.Events["DepositForBurn"].ID:
		return _MockUSDCTokenMessenger.ParseDepositForBurn(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MockUSDCTokenMessengerDepositForBurn) Topic() common.Hash {
	return common.HexToHash("0x2fa9ca894982930190727e75500a97d8dc500233a5065e0f3126c48fbe0343c0")
}

func (_MockUSDCTokenMessenger *MockUSDCTokenMessenger) Address() common.Address {
	return _MockUSDCTokenMessenger.address
}

type MockUSDCTokenMessengerInterface interface {
	IDestinationTokenMessenger(opts *bind.CallOpts) ([32]byte, error)

	LocalMessageTransmitter(opts *bind.CallOpts) (common.Address, error)

	MessageBodyVersion(opts *bind.CallOpts) (uint32, error)

	SNonce(opts *bind.CallOpts) (uint64, error)

	DepositForBurnWithCaller(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error)

	FilterDepositForBurn(opts *bind.FilterOpts, nonce []uint64, burnToken []common.Address, depositor []common.Address) (*MockUSDCTokenMessengerDepositForBurnIterator, error)

	WatchDepositForBurn(opts *bind.WatchOpts, sink chan<- *MockUSDCTokenMessengerDepositForBurn, nonce []uint64, burnToken []common.Address, depositor []common.Address) (event.Subscription, error)

	ParseDepositForBurn(log types.Log) (*MockUSDCTokenMessengerDepositForBurn, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
