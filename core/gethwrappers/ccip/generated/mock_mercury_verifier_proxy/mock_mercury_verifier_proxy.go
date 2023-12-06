// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_mercury_verifier_proxy

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

var MockMercuryVerifierProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"parameterPayload\",\"type\":\"bytes\"}],\"name\":\"verifyBulk\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"verifiedReports\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610223806100206000396000f3fe60806040526004361061001e5760003560e01c8063f873a61c14610023575b600080fd5b610036610031366004610083565b61004c565b6040516100439190610148565b60405180910390f35b60408051600080825260208201909252606091610079565b60608152602001906001900390816100645790505b5095945050505050565b6000806000806040858703121561009957600080fd5b843567ffffffffffffffff808211156100b157600080fd5b818701915087601f8301126100c557600080fd5b8135818111156100d457600080fd5b8860208260051b85010111156100e957600080fd5b60209283019650945090860135908082111561010457600080fd5b818701915087601f83011261011857600080fd5b81358181111561012757600080fd5b88602082850101111561013957600080fd5b95989497505060200194505050565b6000602080830181845280855180835260408601915060408160051b87010192508387016000805b83811015610208577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc089870301855282518051808852835b818110156101c3578281018a01518982018b015289016101a8565b508781018901849052601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016909601870195509386019391860191600101610170565b50939897505050505050505056fea164736f6c6343000813000a",
}

var MockMercuryVerifierProxyABI = MockMercuryVerifierProxyMetaData.ABI

var MockMercuryVerifierProxyBin = MockMercuryVerifierProxyMetaData.Bin

func DeployMockMercuryVerifierProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockMercuryVerifierProxy, error) {
	parsed, err := MockMercuryVerifierProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockMercuryVerifierProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockMercuryVerifierProxy{MockMercuryVerifierProxyCaller: MockMercuryVerifierProxyCaller{contract: contract}, MockMercuryVerifierProxyTransactor: MockMercuryVerifierProxyTransactor{contract: contract}, MockMercuryVerifierProxyFilterer: MockMercuryVerifierProxyFilterer{contract: contract}}, nil
}

type MockMercuryVerifierProxy struct {
	address common.Address
	abi     abi.ABI
	MockMercuryVerifierProxyCaller
	MockMercuryVerifierProxyTransactor
	MockMercuryVerifierProxyFilterer
}

type MockMercuryVerifierProxyCaller struct {
	contract *bind.BoundContract
}

type MockMercuryVerifierProxyTransactor struct {
	contract *bind.BoundContract
}

type MockMercuryVerifierProxyFilterer struct {
	contract *bind.BoundContract
}

type MockMercuryVerifierProxySession struct {
	Contract     *MockMercuryVerifierProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockMercuryVerifierProxyCallerSession struct {
	Contract *MockMercuryVerifierProxyCaller
	CallOpts bind.CallOpts
}

type MockMercuryVerifierProxyTransactorSession struct {
	Contract     *MockMercuryVerifierProxyTransactor
	TransactOpts bind.TransactOpts
}

type MockMercuryVerifierProxyRaw struct {
	Contract *MockMercuryVerifierProxy
}

type MockMercuryVerifierProxyCallerRaw struct {
	Contract *MockMercuryVerifierProxyCaller
}

type MockMercuryVerifierProxyTransactorRaw struct {
	Contract *MockMercuryVerifierProxyTransactor
}

func NewMockMercuryVerifierProxy(address common.Address, backend bind.ContractBackend) (*MockMercuryVerifierProxy, error) {
	abi, err := abi.JSON(strings.NewReader(MockMercuryVerifierProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockMercuryVerifierProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockMercuryVerifierProxy{address: address, abi: abi, MockMercuryVerifierProxyCaller: MockMercuryVerifierProxyCaller{contract: contract}, MockMercuryVerifierProxyTransactor: MockMercuryVerifierProxyTransactor{contract: contract}, MockMercuryVerifierProxyFilterer: MockMercuryVerifierProxyFilterer{contract: contract}}, nil
}

func NewMockMercuryVerifierProxyCaller(address common.Address, caller bind.ContractCaller) (*MockMercuryVerifierProxyCaller, error) {
	contract, err := bindMockMercuryVerifierProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockMercuryVerifierProxyCaller{contract: contract}, nil
}

func NewMockMercuryVerifierProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*MockMercuryVerifierProxyTransactor, error) {
	contract, err := bindMockMercuryVerifierProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockMercuryVerifierProxyTransactor{contract: contract}, nil
}

func NewMockMercuryVerifierProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*MockMercuryVerifierProxyFilterer, error) {
	contract, err := bindMockMercuryVerifierProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockMercuryVerifierProxyFilterer{contract: contract}, nil
}

func bindMockMercuryVerifierProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockMercuryVerifierProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockMercuryVerifierProxy *MockMercuryVerifierProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockMercuryVerifierProxy.Contract.MockMercuryVerifierProxyCaller.contract.Call(opts, result, method, params...)
}

func (_MockMercuryVerifierProxy *MockMercuryVerifierProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockMercuryVerifierProxy.Contract.MockMercuryVerifierProxyTransactor.contract.Transfer(opts)
}

func (_MockMercuryVerifierProxy *MockMercuryVerifierProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockMercuryVerifierProxy.Contract.MockMercuryVerifierProxyTransactor.contract.Transact(opts, method, params...)
}

func (_MockMercuryVerifierProxy *MockMercuryVerifierProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockMercuryVerifierProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_MockMercuryVerifierProxy *MockMercuryVerifierProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockMercuryVerifierProxy.Contract.contract.Transfer(opts)
}

func (_MockMercuryVerifierProxy *MockMercuryVerifierProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockMercuryVerifierProxy.Contract.contract.Transact(opts, method, params...)
}

func (_MockMercuryVerifierProxy *MockMercuryVerifierProxyTransactor) VerifyBulk(opts *bind.TransactOpts, payloads [][]byte, parameterPayload []byte) (*types.Transaction, error) {
	return _MockMercuryVerifierProxy.contract.Transact(opts, "verifyBulk", payloads, parameterPayload)
}

func (_MockMercuryVerifierProxy *MockMercuryVerifierProxySession) VerifyBulk(payloads [][]byte, parameterPayload []byte) (*types.Transaction, error) {
	return _MockMercuryVerifierProxy.Contract.VerifyBulk(&_MockMercuryVerifierProxy.TransactOpts, payloads, parameterPayload)
}

func (_MockMercuryVerifierProxy *MockMercuryVerifierProxyTransactorSession) VerifyBulk(payloads [][]byte, parameterPayload []byte) (*types.Transaction, error) {
	return _MockMercuryVerifierProxy.Contract.VerifyBulk(&_MockMercuryVerifierProxy.TransactOpts, payloads, parameterPayload)
}

func (_MockMercuryVerifierProxy *MockMercuryVerifierProxy) Address() common.Address {
	return _MockMercuryVerifierProxy.address
}

type MockMercuryVerifierProxyInterface interface {
	VerifyBulk(opts *bind.TransactOpts, payloads [][]byte, parameterPayload []byte) (*types.Transaction, error)

	Address() common.Address
}
