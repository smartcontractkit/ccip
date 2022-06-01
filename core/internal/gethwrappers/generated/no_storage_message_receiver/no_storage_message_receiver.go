// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package no_storage_message_receiver

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
)

type CCIPAnyToEVMTollMessage struct {
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

var NoStorageMessageReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.AnyToEVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"receiveMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50608a8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806313d3e05e14602d575b600080fd5b603b6038366004603d565b50565b005b600060208284031215604e57600080fd5b813567ffffffffffffffff811115606457600080fd5b82016101408185031215607657600080fd5b939250505056fea164736f6c634300080d000a",
}

var NoStorageMessageReceiverABI = NoStorageMessageReceiverMetaData.ABI

var NoStorageMessageReceiverBin = NoStorageMessageReceiverMetaData.Bin

func DeployNoStorageMessageReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NoStorageMessageReceiver, error) {
	parsed, err := NoStorageMessageReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NoStorageMessageReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NoStorageMessageReceiver{NoStorageMessageReceiverCaller: NoStorageMessageReceiverCaller{contract: contract}, NoStorageMessageReceiverTransactor: NoStorageMessageReceiverTransactor{contract: contract}, NoStorageMessageReceiverFilterer: NoStorageMessageReceiverFilterer{contract: contract}}, nil
}

type NoStorageMessageReceiver struct {
	address common.Address
	abi     abi.ABI
	NoStorageMessageReceiverCaller
	NoStorageMessageReceiverTransactor
	NoStorageMessageReceiverFilterer
}

type NoStorageMessageReceiverCaller struct {
	contract *bind.BoundContract
}

type NoStorageMessageReceiverTransactor struct {
	contract *bind.BoundContract
}

type NoStorageMessageReceiverFilterer struct {
	contract *bind.BoundContract
}

type NoStorageMessageReceiverSession struct {
	Contract     *NoStorageMessageReceiver
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type NoStorageMessageReceiverCallerSession struct {
	Contract *NoStorageMessageReceiverCaller
	CallOpts bind.CallOpts
}

type NoStorageMessageReceiverTransactorSession struct {
	Contract     *NoStorageMessageReceiverTransactor
	TransactOpts bind.TransactOpts
}

type NoStorageMessageReceiverRaw struct {
	Contract *NoStorageMessageReceiver
}

type NoStorageMessageReceiverCallerRaw struct {
	Contract *NoStorageMessageReceiverCaller
}

type NoStorageMessageReceiverTransactorRaw struct {
	Contract *NoStorageMessageReceiverTransactor
}

func NewNoStorageMessageReceiver(address common.Address, backend bind.ContractBackend) (*NoStorageMessageReceiver, error) {
	abi, err := abi.JSON(strings.NewReader(NoStorageMessageReceiverABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindNoStorageMessageReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NoStorageMessageReceiver{address: address, abi: abi, NoStorageMessageReceiverCaller: NoStorageMessageReceiverCaller{contract: contract}, NoStorageMessageReceiverTransactor: NoStorageMessageReceiverTransactor{contract: contract}, NoStorageMessageReceiverFilterer: NoStorageMessageReceiverFilterer{contract: contract}}, nil
}

func NewNoStorageMessageReceiverCaller(address common.Address, caller bind.ContractCaller) (*NoStorageMessageReceiverCaller, error) {
	contract, err := bindNoStorageMessageReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NoStorageMessageReceiverCaller{contract: contract}, nil
}

func NewNoStorageMessageReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*NoStorageMessageReceiverTransactor, error) {
	contract, err := bindNoStorageMessageReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NoStorageMessageReceiverTransactor{contract: contract}, nil
}

func NewNoStorageMessageReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*NoStorageMessageReceiverFilterer, error) {
	contract, err := bindNoStorageMessageReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NoStorageMessageReceiverFilterer{contract: contract}, nil
}

func bindNoStorageMessageReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NoStorageMessageReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_NoStorageMessageReceiver *NoStorageMessageReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoStorageMessageReceiver.Contract.NoStorageMessageReceiverCaller.contract.Call(opts, result, method, params...)
}

func (_NoStorageMessageReceiver *NoStorageMessageReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoStorageMessageReceiver.Contract.NoStorageMessageReceiverTransactor.contract.Transfer(opts)
}

func (_NoStorageMessageReceiver *NoStorageMessageReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoStorageMessageReceiver.Contract.NoStorageMessageReceiverTransactor.contract.Transact(opts, method, params...)
}

func (_NoStorageMessageReceiver *NoStorageMessageReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoStorageMessageReceiver.Contract.contract.Call(opts, result, method, params...)
}

func (_NoStorageMessageReceiver *NoStorageMessageReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoStorageMessageReceiver.Contract.contract.Transfer(opts)
}

func (_NoStorageMessageReceiver *NoStorageMessageReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoStorageMessageReceiver.Contract.contract.Transact(opts, method, params...)
}

func (_NoStorageMessageReceiver *NoStorageMessageReceiverTransactor) ReceiveMessage(opts *bind.TransactOpts, message CCIPAnyToEVMTollMessage) (*types.Transaction, error) {
	return _NoStorageMessageReceiver.contract.Transact(opts, "receiveMessage", message)
}

func (_NoStorageMessageReceiver *NoStorageMessageReceiverSession) ReceiveMessage(message CCIPAnyToEVMTollMessage) (*types.Transaction, error) {
	return _NoStorageMessageReceiver.Contract.ReceiveMessage(&_NoStorageMessageReceiver.TransactOpts, message)
}

func (_NoStorageMessageReceiver *NoStorageMessageReceiverTransactorSession) ReceiveMessage(message CCIPAnyToEVMTollMessage) (*types.Transaction, error) {
	return _NoStorageMessageReceiver.Contract.ReceiveMessage(&_NoStorageMessageReceiver.TransactOpts, message)
}

func (_NoStorageMessageReceiver *NoStorageMessageReceiver) Address() common.Address {
	return _NoStorageMessageReceiver.address
}

type NoStorageMessageReceiverInterface interface {
	ReceiveMessage(opts *bind.TransactOpts, message CCIPAnyToEVMTollMessage) (*types.Transaction, error)

	Address() common.Address
}
