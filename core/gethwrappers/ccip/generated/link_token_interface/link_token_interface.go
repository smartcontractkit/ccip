// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package link_token_interface

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

var LinkTokenInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"decimalPlaces\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalTokensIssued\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var LinkTokenInterfaceABI = LinkTokenInterfaceMetaData.ABI

type LinkTokenInterface struct {
	address common.Address
	abi     abi.ABI
	LinkTokenInterfaceCaller
	LinkTokenInterfaceTransactor
	LinkTokenInterfaceFilterer
}

type LinkTokenInterfaceCaller struct {
	contract *bind.BoundContract
}

type LinkTokenInterfaceTransactor struct {
	contract *bind.BoundContract
}

type LinkTokenInterfaceFilterer struct {
	contract *bind.BoundContract
}

type LinkTokenInterfaceSession struct {
	Contract     *LinkTokenInterface
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type LinkTokenInterfaceCallerSession struct {
	Contract *LinkTokenInterfaceCaller
	CallOpts bind.CallOpts
}

type LinkTokenInterfaceTransactorSession struct {
	Contract     *LinkTokenInterfaceTransactor
	TransactOpts bind.TransactOpts
}

type LinkTokenInterfaceRaw struct {
	Contract *LinkTokenInterface
}

type LinkTokenInterfaceCallerRaw struct {
	Contract *LinkTokenInterfaceCaller
}

type LinkTokenInterfaceTransactorRaw struct {
	Contract *LinkTokenInterfaceTransactor
}

func NewLinkTokenInterface(address common.Address, backend bind.ContractBackend) (*LinkTokenInterface, error) {
	abi, err := abi.JSON(strings.NewReader(LinkTokenInterfaceABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindLinkTokenInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LinkTokenInterface{address: address, abi: abi, LinkTokenInterfaceCaller: LinkTokenInterfaceCaller{contract: contract}, LinkTokenInterfaceTransactor: LinkTokenInterfaceTransactor{contract: contract}, LinkTokenInterfaceFilterer: LinkTokenInterfaceFilterer{contract: contract}}, nil
}

func NewLinkTokenInterfaceCaller(address common.Address, caller bind.ContractCaller) (*LinkTokenInterfaceCaller, error) {
	contract, err := bindLinkTokenInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LinkTokenInterfaceCaller{contract: contract}, nil
}

func NewLinkTokenInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*LinkTokenInterfaceTransactor, error) {
	contract, err := bindLinkTokenInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LinkTokenInterfaceTransactor{contract: contract}, nil
}

func NewLinkTokenInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*LinkTokenInterfaceFilterer, error) {
	contract, err := bindLinkTokenInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LinkTokenInterfaceFilterer{contract: contract}, nil
}

func bindLinkTokenInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LinkTokenInterfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_LinkTokenInterface *LinkTokenInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LinkTokenInterface.Contract.LinkTokenInterfaceCaller.contract.Call(opts, result, method, params...)
}

func (_LinkTokenInterface *LinkTokenInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.LinkTokenInterfaceTransactor.contract.Transfer(opts)
}

func (_LinkTokenInterface *LinkTokenInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.LinkTokenInterfaceTransactor.contract.Transact(opts, method, params...)
}

func (_LinkTokenInterface *LinkTokenInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LinkTokenInterface.Contract.contract.Call(opts, result, method, params...)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.contract.Transfer(opts)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.contract.Transact(opts, method, params...)
}

func (_LinkTokenInterface *LinkTokenInterfaceCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LinkTokenInterface.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_LinkTokenInterface *LinkTokenInterfaceSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LinkTokenInterface.Contract.Allowance(&_LinkTokenInterface.CallOpts, owner, spender)
}

func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LinkTokenInterface.Contract.Allowance(&_LinkTokenInterface.CallOpts, owner, spender)
}

func (_LinkTokenInterface *LinkTokenInterfaceCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LinkTokenInterface.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_LinkTokenInterface *LinkTokenInterfaceSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LinkTokenInterface.Contract.BalanceOf(&_LinkTokenInterface.CallOpts, owner)
}

func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LinkTokenInterface.Contract.BalanceOf(&_LinkTokenInterface.CallOpts, owner)
}

func (_LinkTokenInterface *LinkTokenInterfaceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LinkTokenInterface.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_LinkTokenInterface *LinkTokenInterfaceSession) Decimals() (uint8, error) {
	return _LinkTokenInterface.Contract.Decimals(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) Decimals() (uint8, error) {
	return _LinkTokenInterface.Contract.Decimals(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LinkTokenInterface.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_LinkTokenInterface *LinkTokenInterfaceSession) Name() (string, error) {
	return _LinkTokenInterface.Contract.Name(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) Name() (string, error) {
	return _LinkTokenInterface.Contract.Name(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LinkTokenInterface.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_LinkTokenInterface *LinkTokenInterfaceSession) Symbol() (string, error) {
	return _LinkTokenInterface.Contract.Symbol(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) Symbol() (string, error) {
	return _LinkTokenInterface.Contract.Symbol(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LinkTokenInterface.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_LinkTokenInterface *LinkTokenInterfaceSession) TotalSupply() (*big.Int, error) {
	return _LinkTokenInterface.Contract.TotalSupply(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) TotalSupply() (*big.Int, error) {
	return _LinkTokenInterface.Contract.TotalSupply(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "approve", spender, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.Approve(&_LinkTokenInterface.TransactOpts, spender, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.Approve(&_LinkTokenInterface.TransactOpts, spender, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "decreaseApproval", spender, addedValue)
}

func (_LinkTokenInterface *LinkTokenInterfaceSession) DecreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.DecreaseApproval(&_LinkTokenInterface.TransactOpts, spender, addedValue)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) DecreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.DecreaseApproval(&_LinkTokenInterface.TransactOpts, spender, addedValue)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "increaseApproval", spender, subtractedValue)
}

func (_LinkTokenInterface *LinkTokenInterfaceSession) IncreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.IncreaseApproval(&_LinkTokenInterface.TransactOpts, spender, subtractedValue)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) IncreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.IncreaseApproval(&_LinkTokenInterface.TransactOpts, spender, subtractedValue)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "transfer", to, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.Transfer(&_LinkTokenInterface.TransactOpts, to, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.Transfer(&_LinkTokenInterface.TransactOpts, to, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "transferAndCall", to, value, data)
}

func (_LinkTokenInterface *LinkTokenInterfaceSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.TransferAndCall(&_LinkTokenInterface.TransactOpts, to, value, data)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.TransferAndCall(&_LinkTokenInterface.TransactOpts, to, value, data)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "transferFrom", from, to, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.TransferFrom(&_LinkTokenInterface.TransactOpts, from, to, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.TransferFrom(&_LinkTokenInterface.TransactOpts, from, to, value)
}

func (_LinkTokenInterface *LinkTokenInterface) Address() common.Address {
	return _LinkTokenInterface.address
}

type LinkTokenInterfaceInterface interface {
	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	Name(opts *bind.CallOpts) (string, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error)

	DecreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error)

	IncreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error)

	TransferAndCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error)

	Address() common.Address
}
