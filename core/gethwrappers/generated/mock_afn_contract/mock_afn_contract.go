// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_afn_contract

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

type AFNUnvoteToCurseRecord struct {
	CurseVoteAddr common.Address
	CursesHash    [32]byte
	ForceUnvote   bool
}

type IAFNTaggedRoot struct {
	CommitStore common.Address
	Root        [32]byte
}

var MockAFNContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structIAFN.TaggedRoot\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"isBlessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCursed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"curseVoteAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"cursesHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"forceUnvote\",\"type\":\"bool\"}],\"internalType\":\"structAFN.UnvoteToCurseRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"name\":\"ownerUnvoteToCurse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voteToCurse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506102e7806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063119a352714610051578063397796f71461008f5780634d616771146100aa578063618af128146100c3575b600080fd5b61008d61005f3660046100fc565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b005b60005460ff165b604051901515815260200160405180910390f35b6100966100b8366004610115565b5060005460ff161590565b61008d6100d13660046101d4565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b60006020828403121561010e57600080fd5b5035919050565b60006040828403121561012757600080fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561017f5761017f61012d565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156101cc576101cc61012d565b604052919050565b600060208083850312156101e757600080fd5b823567ffffffffffffffff808211156101ff57600080fd5b818501915085601f83011261021357600080fd5b8135818111156102255761022561012d565b610233848260051b01610185565b8181528481019250606091820284018501918883111561025257600080fd5b938501935b828510156102ce5780858a03121561026f5760008081fd5b61027761015c565b853573ffffffffffffffffffffffffffffffffffffffff8116811461029c5760008081fd5b8152858701358782015260408087013580151581146102bb5760008081fd5b9082015284529384019392850192610257565b5097965050505050505056fea164736f6c634300080f000a",
}

var MockAFNContractABI = MockAFNContractMetaData.ABI

var MockAFNContractBin = MockAFNContractMetaData.Bin

func DeployMockAFNContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockAFNContract, error) {
	parsed, err := MockAFNContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockAFNContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockAFNContract{MockAFNContractCaller: MockAFNContractCaller{contract: contract}, MockAFNContractTransactor: MockAFNContractTransactor{contract: contract}, MockAFNContractFilterer: MockAFNContractFilterer{contract: contract}}, nil
}

type MockAFNContract struct {
	address common.Address
	abi     abi.ABI
	MockAFNContractCaller
	MockAFNContractTransactor
	MockAFNContractFilterer
}

type MockAFNContractCaller struct {
	contract *bind.BoundContract
}

type MockAFNContractTransactor struct {
	contract *bind.BoundContract
}

type MockAFNContractFilterer struct {
	contract *bind.BoundContract
}

type MockAFNContractSession struct {
	Contract     *MockAFNContract
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockAFNContractCallerSession struct {
	Contract *MockAFNContractCaller
	CallOpts bind.CallOpts
}

type MockAFNContractTransactorSession struct {
	Contract     *MockAFNContractTransactor
	TransactOpts bind.TransactOpts
}

type MockAFNContractRaw struct {
	Contract *MockAFNContract
}

type MockAFNContractCallerRaw struct {
	Contract *MockAFNContractCaller
}

type MockAFNContractTransactorRaw struct {
	Contract *MockAFNContractTransactor
}

func NewMockAFNContract(address common.Address, backend bind.ContractBackend) (*MockAFNContract, error) {
	abi, err := abi.JSON(strings.NewReader(MockAFNContractABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockAFNContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockAFNContract{address: address, abi: abi, MockAFNContractCaller: MockAFNContractCaller{contract: contract}, MockAFNContractTransactor: MockAFNContractTransactor{contract: contract}, MockAFNContractFilterer: MockAFNContractFilterer{contract: contract}}, nil
}

func NewMockAFNContractCaller(address common.Address, caller bind.ContractCaller) (*MockAFNContractCaller, error) {
	contract, err := bindMockAFNContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockAFNContractCaller{contract: contract}, nil
}

func NewMockAFNContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MockAFNContractTransactor, error) {
	contract, err := bindMockAFNContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockAFNContractTransactor{contract: contract}, nil
}

func NewMockAFNContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MockAFNContractFilterer, error) {
	contract, err := bindMockAFNContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockAFNContractFilterer{contract: contract}, nil
}

func bindMockAFNContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockAFNContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockAFNContract *MockAFNContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockAFNContract.Contract.MockAFNContractCaller.contract.Call(opts, result, method, params...)
}

func (_MockAFNContract *MockAFNContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockAFNContract.Contract.MockAFNContractTransactor.contract.Transfer(opts)
}

func (_MockAFNContract *MockAFNContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockAFNContract.Contract.MockAFNContractTransactor.contract.Transact(opts, method, params...)
}

func (_MockAFNContract *MockAFNContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockAFNContract.Contract.contract.Call(opts, result, method, params...)
}

func (_MockAFNContract *MockAFNContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockAFNContract.Contract.contract.Transfer(opts)
}

func (_MockAFNContract *MockAFNContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockAFNContract.Contract.contract.Transact(opts, method, params...)
}

func (_MockAFNContract *MockAFNContractCaller) IsBlessed(opts *bind.CallOpts, arg0 IAFNTaggedRoot) (bool, error) {
	var out []interface{}
	err := _MockAFNContract.contract.Call(opts, &out, "isBlessed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockAFNContract *MockAFNContractSession) IsBlessed(arg0 IAFNTaggedRoot) (bool, error) {
	return _MockAFNContract.Contract.IsBlessed(&_MockAFNContract.CallOpts, arg0)
}

func (_MockAFNContract *MockAFNContractCallerSession) IsBlessed(arg0 IAFNTaggedRoot) (bool, error) {
	return _MockAFNContract.Contract.IsBlessed(&_MockAFNContract.CallOpts, arg0)
}

func (_MockAFNContract *MockAFNContractCaller) IsCursed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockAFNContract.contract.Call(opts, &out, "isCursed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockAFNContract *MockAFNContractSession) IsCursed() (bool, error) {
	return _MockAFNContract.Contract.IsCursed(&_MockAFNContract.CallOpts)
}

func (_MockAFNContract *MockAFNContractCallerSession) IsCursed() (bool, error) {
	return _MockAFNContract.Contract.IsCursed(&_MockAFNContract.CallOpts)
}

func (_MockAFNContract *MockAFNContractTransactor) OwnerUnvoteToCurse(opts *bind.TransactOpts, arg0 []AFNUnvoteToCurseRecord) (*types.Transaction, error) {
	return _MockAFNContract.contract.Transact(opts, "ownerUnvoteToCurse", arg0)
}

func (_MockAFNContract *MockAFNContractSession) OwnerUnvoteToCurse(arg0 []AFNUnvoteToCurseRecord) (*types.Transaction, error) {
	return _MockAFNContract.Contract.OwnerUnvoteToCurse(&_MockAFNContract.TransactOpts, arg0)
}

func (_MockAFNContract *MockAFNContractTransactorSession) OwnerUnvoteToCurse(arg0 []AFNUnvoteToCurseRecord) (*types.Transaction, error) {
	return _MockAFNContract.Contract.OwnerUnvoteToCurse(&_MockAFNContract.TransactOpts, arg0)
}

func (_MockAFNContract *MockAFNContractTransactor) VoteToCurse(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error) {
	return _MockAFNContract.contract.Transact(opts, "voteToCurse", arg0)
}

func (_MockAFNContract *MockAFNContractSession) VoteToCurse(arg0 [32]byte) (*types.Transaction, error) {
	return _MockAFNContract.Contract.VoteToCurse(&_MockAFNContract.TransactOpts, arg0)
}

func (_MockAFNContract *MockAFNContractTransactorSession) VoteToCurse(arg0 [32]byte) (*types.Transaction, error) {
	return _MockAFNContract.Contract.VoteToCurse(&_MockAFNContract.TransactOpts, arg0)
}

func (_MockAFNContract *MockAFNContract) Address() common.Address {
	return _MockAFNContract.address
}

type MockAFNContractInterface interface {
	IsBlessed(opts *bind.CallOpts, arg0 IAFNTaggedRoot) (bool, error)

	IsCursed(opts *bind.CallOpts) (bool, error)

	OwnerUnvoteToCurse(opts *bind.TransactOpts, arg0 []AFNUnvoteToCurseRecord) (*types.Transaction, error)

	VoteToCurse(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error)

	Address() common.Address
}
