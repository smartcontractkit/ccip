// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_usdc_token_transmitter

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

var MockUSDCTransmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"localDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_shouldSucceed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"shouldSucceed\",\"type\":\"bool\"}],\"name\":\"setShouldSucceed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b506040516102f23803806102f283398101604081905261002f91610069565b63ffffffff9182166080521660a0526000805460ff1916600117905561009c565b805163ffffffff8116811461006457600080fd5b919050565b6000806040838503121561007c57600080fd5b61008583610050565b915061009360208401610050565b90509250929050565b60805160a0516102336100bf600039600060e301526000606e01526102336000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80637a642935116100505780637a642935146100d45780638d3638f4146100e15780639e31ddb61461010757600080fd5b806354fd4d501461006c57806357ecfd28146100a8575b600080fd5b7f00000000000000000000000000000000000000000000000000000000000000005b60405163ffffffff90911681526020015b60405180910390f35b6100c46100b6366004610191565b60005460ff16949350505050565b604051901515815260200161009f565b6000546100c49060ff1681565b7f000000000000000000000000000000000000000000000000000000000000000061008e565b6101466101153660046101fd565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b005b60008083601f84011261015a57600080fd5b50813567ffffffffffffffff81111561017257600080fd5b60208301915083602082850101111561018a57600080fd5b9250929050565b600080600080604085870312156101a757600080fd5b843567ffffffffffffffff808211156101bf57600080fd5b6101cb88838901610148565b909650945060208701359150808211156101e457600080fd5b506101f187828801610148565b95989497509550505050565b60006020828403121561020f57600080fd5b8135801515811461021f57600080fd5b939250505056fea164736f6c6343000813000a",
}

var MockUSDCTransmitterABI = MockUSDCTransmitterMetaData.ABI

var MockUSDCTransmitterBin = MockUSDCTransmitterMetaData.Bin

func DeployMockUSDCTransmitter(auth *bind.TransactOpts, backend bind.ContractBackend, version uint32, localDomain uint32) (common.Address, *types.Transaction, *MockUSDCTransmitter, error) {
	parsed, err := MockUSDCTransmitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockUSDCTransmitterBin), backend, version, localDomain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockUSDCTransmitter{address: address, abi: *parsed, MockUSDCTransmitterCaller: MockUSDCTransmitterCaller{contract: contract}, MockUSDCTransmitterTransactor: MockUSDCTransmitterTransactor{contract: contract}, MockUSDCTransmitterFilterer: MockUSDCTransmitterFilterer{contract: contract}}, nil
}

type MockUSDCTransmitter struct {
	address common.Address
	abi     abi.ABI
	MockUSDCTransmitterCaller
	MockUSDCTransmitterTransactor
	MockUSDCTransmitterFilterer
}

type MockUSDCTransmitterCaller struct {
	contract *bind.BoundContract
}

type MockUSDCTransmitterTransactor struct {
	contract *bind.BoundContract
}

type MockUSDCTransmitterFilterer struct {
	contract *bind.BoundContract
}

type MockUSDCTransmitterSession struct {
	Contract     *MockUSDCTransmitter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockUSDCTransmitterCallerSession struct {
	Contract *MockUSDCTransmitterCaller
	CallOpts bind.CallOpts
}

type MockUSDCTransmitterTransactorSession struct {
	Contract     *MockUSDCTransmitterTransactor
	TransactOpts bind.TransactOpts
}

type MockUSDCTransmitterRaw struct {
	Contract *MockUSDCTransmitter
}

type MockUSDCTransmitterCallerRaw struct {
	Contract *MockUSDCTransmitterCaller
}

type MockUSDCTransmitterTransactorRaw struct {
	Contract *MockUSDCTransmitterTransactor
}

func NewMockUSDCTransmitter(address common.Address, backend bind.ContractBackend) (*MockUSDCTransmitter, error) {
	abi, err := abi.JSON(strings.NewReader(MockUSDCTransmitterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockUSDCTransmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockUSDCTransmitter{address: address, abi: abi, MockUSDCTransmitterCaller: MockUSDCTransmitterCaller{contract: contract}, MockUSDCTransmitterTransactor: MockUSDCTransmitterTransactor{contract: contract}, MockUSDCTransmitterFilterer: MockUSDCTransmitterFilterer{contract: contract}}, nil
}

func NewMockUSDCTransmitterCaller(address common.Address, caller bind.ContractCaller) (*MockUSDCTransmitterCaller, error) {
	contract, err := bindMockUSDCTransmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockUSDCTransmitterCaller{contract: contract}, nil
}

func NewMockUSDCTransmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*MockUSDCTransmitterTransactor, error) {
	contract, err := bindMockUSDCTransmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockUSDCTransmitterTransactor{contract: contract}, nil
}

func NewMockUSDCTransmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*MockUSDCTransmitterFilterer, error) {
	contract, err := bindMockUSDCTransmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockUSDCTransmitterFilterer{contract: contract}, nil
}

func bindMockUSDCTransmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockUSDCTransmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockUSDCTransmitter *MockUSDCTransmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockUSDCTransmitter.Contract.MockUSDCTransmitterCaller.contract.Call(opts, result, method, params...)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockUSDCTransmitter.Contract.MockUSDCTransmitterTransactor.contract.Transfer(opts)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockUSDCTransmitter.Contract.MockUSDCTransmitterTransactor.contract.Transact(opts, method, params...)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockUSDCTransmitter.Contract.contract.Call(opts, result, method, params...)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockUSDCTransmitter.Contract.contract.Transfer(opts)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockUSDCTransmitter.Contract.contract.Transact(opts, method, params...)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MockUSDCTransmitter.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_MockUSDCTransmitter *MockUSDCTransmitterSession) LocalDomain() (uint32, error) {
	return _MockUSDCTransmitter.Contract.LocalDomain(&_MockUSDCTransmitter.CallOpts)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterCallerSession) LocalDomain() (uint32, error) {
	return _MockUSDCTransmitter.Contract.LocalDomain(&_MockUSDCTransmitter.CallOpts)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterCaller) ReceiveMessage(opts *bind.CallOpts, arg0 []byte, arg1 []byte) (bool, error) {
	var out []interface{}
	err := _MockUSDCTransmitter.contract.Call(opts, &out, "receiveMessage", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockUSDCTransmitter *MockUSDCTransmitterSession) ReceiveMessage(arg0 []byte, arg1 []byte) (bool, error) {
	return _MockUSDCTransmitter.Contract.ReceiveMessage(&_MockUSDCTransmitter.CallOpts, arg0, arg1)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterCallerSession) ReceiveMessage(arg0 []byte, arg1 []byte) (bool, error) {
	return _MockUSDCTransmitter.Contract.ReceiveMessage(&_MockUSDCTransmitter.CallOpts, arg0, arg1)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterCaller) SShouldSucceed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockUSDCTransmitter.contract.Call(opts, &out, "s_shouldSucceed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockUSDCTransmitter *MockUSDCTransmitterSession) SShouldSucceed() (bool, error) {
	return _MockUSDCTransmitter.Contract.SShouldSucceed(&_MockUSDCTransmitter.CallOpts)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterCallerSession) SShouldSucceed() (bool, error) {
	return _MockUSDCTransmitter.Contract.SShouldSucceed(&_MockUSDCTransmitter.CallOpts)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterCaller) Version(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MockUSDCTransmitter.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_MockUSDCTransmitter *MockUSDCTransmitterSession) Version() (uint32, error) {
	return _MockUSDCTransmitter.Contract.Version(&_MockUSDCTransmitter.CallOpts)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterCallerSession) Version() (uint32, error) {
	return _MockUSDCTransmitter.Contract.Version(&_MockUSDCTransmitter.CallOpts)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterTransactor) SetShouldSucceed(opts *bind.TransactOpts, shouldSucceed bool) (*types.Transaction, error) {
	return _MockUSDCTransmitter.contract.Transact(opts, "setShouldSucceed", shouldSucceed)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterSession) SetShouldSucceed(shouldSucceed bool) (*types.Transaction, error) {
	return _MockUSDCTransmitter.Contract.SetShouldSucceed(&_MockUSDCTransmitter.TransactOpts, shouldSucceed)
}

func (_MockUSDCTransmitter *MockUSDCTransmitterTransactorSession) SetShouldSucceed(shouldSucceed bool) (*types.Transaction, error) {
	return _MockUSDCTransmitter.Contract.SetShouldSucceed(&_MockUSDCTransmitter.TransactOpts, shouldSucceed)
}

func (_MockUSDCTransmitter *MockUSDCTransmitter) Address() common.Address {
	return _MockUSDCTransmitter.address
}

type MockUSDCTransmitterInterface interface {
	LocalDomain(opts *bind.CallOpts) (uint32, error)

	ReceiveMessage(opts *bind.CallOpts, arg0 []byte, arg1 []byte) (bool, error)

	SShouldSucceed(opts *bind.CallOpts) (bool, error)

	Version(opts *bind.CallOpts) (uint32, error)

	SetShouldSucceed(opts *bind.TransactOpts, shouldSucceed bool) (*types.Transaction, error)

	Address() common.Address
}
