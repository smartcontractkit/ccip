// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_arm_contract

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

type IRMNTaggedRoot struct {
	CommitStore common.Address
	Root        [32]byte
}

type RMNConfig struct {
	Voters               []RMNVoter
	BlessWeightThreshold uint16
	CurseWeightThreshold uint16
}

type RMNUnvoteToCurseRecord struct {
	CurseVoteAddr common.Address
	CursesHash    [32]byte
	ForceUnvote   bool
}

type RMNVoter struct {
	BlessVoteAddr   common.Address
	CurseVoteAddr   common.Address
	CurseUnvoteAddr common.Address
	BlessWeight     uint8
	CurseWeight     uint8
}

var MockARMContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"CustomError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"blessVoteAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"curseVoteAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"curseUnvoteAddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"blessWeight\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"curseWeight\",\"type\":\"uint8\"}],\"internalType\":\"structRMN.Voter[]\",\"name\":\"voters\",\"type\":\"tuple[]\"},{\"internalType\":\"uint16\",\"name\":\"blessWeightThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"curseWeightThreshold\",\"type\":\"uint16\"}],\"internalType\":\"structRMN.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structIRMN.TaggedRoot\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"isBlessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCursed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subject\",\"type\":\"bytes32\"}],\"name\":\"isCursed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"curseVoteAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"cursesHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"forceUnvote\",\"type\":\"bool\"}],\"internalType\":\"structRMN.UnvoteToCurseRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"subject\",\"type\":\"bytes32\"}],\"name\":\"ownerUnvoteToCurse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"curseVoteAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"cursesHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"forceUnvote\",\"type\":\"bool\"}],\"internalType\":\"structRMN.UnvoteToCurseRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"name\":\"ownerUnvoteToCurse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"setRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subject\",\"type\":\"bytes32\"}],\"name\":\"voteToCurse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteToCurse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b610dec806101576000396000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c8063618af128116100815780638da5cb5b1161005b5780638da5cb5b146102335780639362812d1461025b578063f2fde38b1461029f57600080fd5b8063618af128146101df57806379ba5097146102185780637a7c27491461022057600080fd5b8063447deab7116100b2578063447deab7146101565780634d6167711461019b57806358babe33146101cc57600080fd5b8063119a3527146100d9578063397796f7146101225780633f42ab731461013f575b600080fd5b6101206100e736600461078a565b600090815260066020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b005b61012a6102b2565b60405190151581526020015b60405180910390f35b610147610328565b604051610136939291906107a3565b610120610164366004610a0a565b600090815260066020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905550565b61012a6101a9366004610a4f565b5060015474010000000000000000000000000000000000000000900460ff161590565b61012a6101da36600461078a565b610469565b6101206101ed366004610a67565b50600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff169055565b6101206104f1565b61012061022e366004610aa4565b6105ee565b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610136565b610120600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000179055565b6101206102ad366004610b57565b6105fe565b6000600280546102c190610b79565b1590506103065760026040517f5a4ff6710000000000000000000000000000000000000000000000000000000081526004016102fd9190610bc6565b60405180910390fd5b5060015474010000000000000000000000000000000000000000900460ff1690565b6040805160608082018352815260006020820181905291810182905281906005546040805160038054608060208202840181019094526060830181815263ffffffff8087169664010000000090041694929392849284929184919060009085015b828210156104355760008481526020908190206040805160a08101825260038602909201805473ffffffffffffffffffffffffffffffffffffffff90811684526001808301548216858701526002909201549081169284019290925260ff74010000000000000000000000000000000000000000830481166060850152750100000000000000000000000000000000000000000090920490911660808301529083529092019101610389565b505050908252506001919091015461ffff808216602084015262010000909104166040909101529296919550919350915050565b60006002805461047890610b79565b1590506104b45760026040517f5a4ff6710000000000000000000000000000000000000000000000000000000081526004016102fd9190610bc6565b60015474010000000000000000000000000000000000000000900460ff16806104eb575060008281526006602052604090205460ff165b92915050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610572576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016102fd565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60026105fa8282610cc5565b5050565b610606610612565b61060f81610695565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610693576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016102fd565b565b3373ffffffffffffffffffffffffffffffffffffffff821603610714576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016102fd565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561079c57600080fd5b5035919050565b63ffffffff84811682528316602080830191909152606060408084018290528451848301839052805160c0860181905260009491820190859060e08801905b80831015610849578351805173ffffffffffffffffffffffffffffffffffffffff9081168452868201518116878501528782015116878401528781015160ff908116898501526080918201511690830152928401926001929092019160a0909101906107e2565b509288015161ffff908116608089015260409098015190971660a090960195909552979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156108c8576108c8610876565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561091557610915610876565b604052919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461094157600080fd5b919050565b600082601f83011261095757600080fd5b8135602067ffffffffffffffff82111561097357610973610876565b610981818360051b016108ce565b828152606092830285018201928282019190878511156109a057600080fd5b8387015b858110156109fd5781818a0312156109bc5760008081fd5b6109c46108a5565b6109cd8261091d565b8152858201358682015260408083013580151581146109ec5760008081fd5b9082015284529284019281016109a4565b5090979650505050505050565b60008060408385031215610a1d57600080fd5b823567ffffffffffffffff811115610a3457600080fd5b610a4085828601610946565b95602094909401359450505050565b600060408284031215610a6157600080fd5b50919050565b600060208284031215610a7957600080fd5b813567ffffffffffffffff811115610a9057600080fd5b610a9c84828501610946565b949350505050565b60006020808385031215610ab757600080fd5b823567ffffffffffffffff80821115610acf57600080fd5b818501915085601f830112610ae357600080fd5b813581811115610af557610af5610876565b610b25847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016108ce565b91508082528684828501011115610b3b57600080fd5b8084840185840137600090820190930192909252509392505050565b600060208284031215610b6957600080fd5b610b728261091d565b9392505050565b600181811c90821680610b8d57607f821691505b602082108103610a61577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000602080835260008454610bda81610b79565b8060208701526040600180841660008114610bfc5760018114610c3657610c66565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00851660408a0152604084151560051b8a01019550610c66565b89600052602060002060005b85811015610c5d5781548b8201860152908301908801610c42565b8a016040019650505b509398975050505050505050565b601f821115610cc0576000816000526020600020601f850160051c81016020861015610c9d5750805b601f850160051c820191505b81811015610cbc57828155600101610ca9565b5050505b505050565b815167ffffffffffffffff811115610cdf57610cdf610876565b610cf381610ced8454610b79565b84610c74565b602080601f831160018114610d465760008415610d105750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555610cbc565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015610d9357888601518255948401946001909101908401610d74565b5085821015610dcf57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c6343000818000a",
}

var MockARMContractABI = MockARMContractMetaData.ABI

var MockARMContractBin = MockARMContractMetaData.Bin

func DeployMockARMContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockARMContract, error) {
	parsed, err := MockARMContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockARMContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockARMContract{address: address, abi: *parsed, MockARMContractCaller: MockARMContractCaller{contract: contract}, MockARMContractTransactor: MockARMContractTransactor{contract: contract}, MockARMContractFilterer: MockARMContractFilterer{contract: contract}}, nil
}

type MockARMContract struct {
	address common.Address
	abi     abi.ABI
	MockARMContractCaller
	MockARMContractTransactor
	MockARMContractFilterer
}

type MockARMContractCaller struct {
	contract *bind.BoundContract
}

type MockARMContractTransactor struct {
	contract *bind.BoundContract
}

type MockARMContractFilterer struct {
	contract *bind.BoundContract
}

type MockARMContractSession struct {
	Contract     *MockARMContract
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockARMContractCallerSession struct {
	Contract *MockARMContractCaller
	CallOpts bind.CallOpts
}

type MockARMContractTransactorSession struct {
	Contract     *MockARMContractTransactor
	TransactOpts bind.TransactOpts
}

type MockARMContractRaw struct {
	Contract *MockARMContract
}

type MockARMContractCallerRaw struct {
	Contract *MockARMContractCaller
}

type MockARMContractTransactorRaw struct {
	Contract *MockARMContractTransactor
}

func NewMockARMContract(address common.Address, backend bind.ContractBackend) (*MockARMContract, error) {
	abi, err := abi.JSON(strings.NewReader(MockARMContractABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockARMContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockARMContract{address: address, abi: abi, MockARMContractCaller: MockARMContractCaller{contract: contract}, MockARMContractTransactor: MockARMContractTransactor{contract: contract}, MockARMContractFilterer: MockARMContractFilterer{contract: contract}}, nil
}

func NewMockARMContractCaller(address common.Address, caller bind.ContractCaller) (*MockARMContractCaller, error) {
	contract, err := bindMockARMContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockARMContractCaller{contract: contract}, nil
}

func NewMockARMContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MockARMContractTransactor, error) {
	contract, err := bindMockARMContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockARMContractTransactor{contract: contract}, nil
}

func NewMockARMContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MockARMContractFilterer, error) {
	contract, err := bindMockARMContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockARMContractFilterer{contract: contract}, nil
}

func bindMockARMContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockARMContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockARMContract *MockARMContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockARMContract.Contract.MockARMContractCaller.contract.Call(opts, result, method, params...)
}

func (_MockARMContract *MockARMContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockARMContract.Contract.MockARMContractTransactor.contract.Transfer(opts)
}

func (_MockARMContract *MockARMContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockARMContract.Contract.MockARMContractTransactor.contract.Transact(opts, method, params...)
}

func (_MockARMContract *MockARMContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockARMContract.Contract.contract.Call(opts, result, method, params...)
}

func (_MockARMContract *MockARMContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockARMContract.Contract.contract.Transfer(opts)
}

func (_MockARMContract *MockARMContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockARMContract.Contract.contract.Transact(opts, method, params...)
}

func (_MockARMContract *MockARMContractCaller) GetConfigDetails(opts *bind.CallOpts) (GetConfigDetails,

	error) {
	var out []interface{}
	err := _MockARMContract.contract.Call(opts, &out, "getConfigDetails")

	outstruct := new(GetConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Version = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Config = *abi.ConvertType(out[2], new(RMNConfig)).(*RMNConfig)

	return *outstruct, err

}

func (_MockARMContract *MockARMContractSession) GetConfigDetails() (GetConfigDetails,

	error) {
	return _MockARMContract.Contract.GetConfigDetails(&_MockARMContract.CallOpts)
}

func (_MockARMContract *MockARMContractCallerSession) GetConfigDetails() (GetConfigDetails,

	error) {
	return _MockARMContract.Contract.GetConfigDetails(&_MockARMContract.CallOpts)
}

func (_MockARMContract *MockARMContractCaller) IsBlessed(opts *bind.CallOpts, arg0 IRMNTaggedRoot) (bool, error) {
	var out []interface{}
	err := _MockARMContract.contract.Call(opts, &out, "isBlessed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockARMContract *MockARMContractSession) IsBlessed(arg0 IRMNTaggedRoot) (bool, error) {
	return _MockARMContract.Contract.IsBlessed(&_MockARMContract.CallOpts, arg0)
}

func (_MockARMContract *MockARMContractCallerSession) IsBlessed(arg0 IRMNTaggedRoot) (bool, error) {
	return _MockARMContract.Contract.IsBlessed(&_MockARMContract.CallOpts, arg0)
}

func (_MockARMContract *MockARMContractCaller) IsCursed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockARMContract.contract.Call(opts, &out, "isCursed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockARMContract *MockARMContractSession) IsCursed() (bool, error) {
	return _MockARMContract.Contract.IsCursed(&_MockARMContract.CallOpts)
}

func (_MockARMContract *MockARMContractCallerSession) IsCursed() (bool, error) {
	return _MockARMContract.Contract.IsCursed(&_MockARMContract.CallOpts)
}

func (_MockARMContract *MockARMContractCaller) IsCursed0(opts *bind.CallOpts, subject [32]byte) (bool, error) {
	var out []interface{}
	err := _MockARMContract.contract.Call(opts, &out, "isCursed0", subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockARMContract *MockARMContractSession) IsCursed0(subject [32]byte) (bool, error) {
	return _MockARMContract.Contract.IsCursed0(&_MockARMContract.CallOpts, subject)
}

func (_MockARMContract *MockARMContractCallerSession) IsCursed0(subject [32]byte) (bool, error) {
	return _MockARMContract.Contract.IsCursed0(&_MockARMContract.CallOpts, subject)
}

func (_MockARMContract *MockARMContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockARMContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockARMContract *MockARMContractSession) Owner() (common.Address, error) {
	return _MockARMContract.Contract.Owner(&_MockARMContract.CallOpts)
}

func (_MockARMContract *MockARMContractCallerSession) Owner() (common.Address, error) {
	return _MockARMContract.Contract.Owner(&_MockARMContract.CallOpts)
}

func (_MockARMContract *MockARMContractTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockARMContract.contract.Transact(opts, "acceptOwnership")
}

func (_MockARMContract *MockARMContractSession) AcceptOwnership() (*types.Transaction, error) {
	return _MockARMContract.Contract.AcceptOwnership(&_MockARMContract.TransactOpts)
}

func (_MockARMContract *MockARMContractTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MockARMContract.Contract.AcceptOwnership(&_MockARMContract.TransactOpts)
}

func (_MockARMContract *MockARMContractTransactor) OwnerUnvoteToCurse(opts *bind.TransactOpts, arg0 []RMNUnvoteToCurseRecord, subject [32]byte) (*types.Transaction, error) {
	return _MockARMContract.contract.Transact(opts, "ownerUnvoteToCurse", arg0, subject)
}

func (_MockARMContract *MockARMContractSession) OwnerUnvoteToCurse(arg0 []RMNUnvoteToCurseRecord, subject [32]byte) (*types.Transaction, error) {
	return _MockARMContract.Contract.OwnerUnvoteToCurse(&_MockARMContract.TransactOpts, arg0, subject)
}

func (_MockARMContract *MockARMContractTransactorSession) OwnerUnvoteToCurse(arg0 []RMNUnvoteToCurseRecord, subject [32]byte) (*types.Transaction, error) {
	return _MockARMContract.Contract.OwnerUnvoteToCurse(&_MockARMContract.TransactOpts, arg0, subject)
}

func (_MockARMContract *MockARMContractTransactor) OwnerUnvoteToCurse0(opts *bind.TransactOpts, arg0 []RMNUnvoteToCurseRecord) (*types.Transaction, error) {
	return _MockARMContract.contract.Transact(opts, "ownerUnvoteToCurse0", arg0)
}

func (_MockARMContract *MockARMContractSession) OwnerUnvoteToCurse0(arg0 []RMNUnvoteToCurseRecord) (*types.Transaction, error) {
	return _MockARMContract.Contract.OwnerUnvoteToCurse0(&_MockARMContract.TransactOpts, arg0)
}

func (_MockARMContract *MockARMContractTransactorSession) OwnerUnvoteToCurse0(arg0 []RMNUnvoteToCurseRecord) (*types.Transaction, error) {
	return _MockARMContract.Contract.OwnerUnvoteToCurse0(&_MockARMContract.TransactOpts, arg0)
}

func (_MockARMContract *MockARMContractTransactor) SetRevert(opts *bind.TransactOpts, err []byte) (*types.Transaction, error) {
	return _MockARMContract.contract.Transact(opts, "setRevert", err)
}

func (_MockARMContract *MockARMContractSession) SetRevert(err []byte) (*types.Transaction, error) {
	return _MockARMContract.Contract.SetRevert(&_MockARMContract.TransactOpts, err)
}

func (_MockARMContract *MockARMContractTransactorSession) SetRevert(err []byte) (*types.Transaction, error) {
	return _MockARMContract.Contract.SetRevert(&_MockARMContract.TransactOpts, err)
}

func (_MockARMContract *MockARMContractTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MockARMContract.contract.Transact(opts, "transferOwnership", to)
}

func (_MockARMContract *MockARMContractSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MockARMContract.Contract.TransferOwnership(&_MockARMContract.TransactOpts, to)
}

func (_MockARMContract *MockARMContractTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MockARMContract.Contract.TransferOwnership(&_MockARMContract.TransactOpts, to)
}

func (_MockARMContract *MockARMContractTransactor) VoteToCurse(opts *bind.TransactOpts, subject [32]byte) (*types.Transaction, error) {
	return _MockARMContract.contract.Transact(opts, "voteToCurse", subject)
}

func (_MockARMContract *MockARMContractSession) VoteToCurse(subject [32]byte) (*types.Transaction, error) {
	return _MockARMContract.Contract.VoteToCurse(&_MockARMContract.TransactOpts, subject)
}

func (_MockARMContract *MockARMContractTransactorSession) VoteToCurse(subject [32]byte) (*types.Transaction, error) {
	return _MockARMContract.Contract.VoteToCurse(&_MockARMContract.TransactOpts, subject)
}

func (_MockARMContract *MockARMContractTransactor) VoteToCurse0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockARMContract.contract.Transact(opts, "voteToCurse0")
}

func (_MockARMContract *MockARMContractSession) VoteToCurse0() (*types.Transaction, error) {
	return _MockARMContract.Contract.VoteToCurse0(&_MockARMContract.TransactOpts)
}

func (_MockARMContract *MockARMContractTransactorSession) VoteToCurse0() (*types.Transaction, error) {
	return _MockARMContract.Contract.VoteToCurse0(&_MockARMContract.TransactOpts)
}

type MockARMContractOwnershipTransferRequestedIterator struct {
	Event *MockARMContractOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockARMContractOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockARMContractOwnershipTransferRequested)
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
		it.Event = new(MockARMContractOwnershipTransferRequested)
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

func (it *MockARMContractOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *MockARMContractOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockARMContractOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MockARMContract *MockARMContractFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockARMContractOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockARMContract.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockARMContractOwnershipTransferRequestedIterator{contract: _MockARMContract.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_MockARMContract *MockARMContractFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MockARMContractOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockARMContract.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockARMContractOwnershipTransferRequested)
				if err := _MockARMContract.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_MockARMContract *MockARMContractFilterer) ParseOwnershipTransferRequested(log types.Log) (*MockARMContractOwnershipTransferRequested, error) {
	event := new(MockARMContractOwnershipTransferRequested)
	if err := _MockARMContract.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockARMContractOwnershipTransferredIterator struct {
	Event *MockARMContractOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockARMContractOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockARMContractOwnershipTransferred)
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
		it.Event = new(MockARMContractOwnershipTransferred)
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

func (it *MockARMContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *MockARMContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockARMContractOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MockARMContract *MockARMContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockARMContractOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockARMContract.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockARMContractOwnershipTransferredIterator{contract: _MockARMContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_MockARMContract *MockARMContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockARMContractOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockARMContract.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockARMContractOwnershipTransferred)
				if err := _MockARMContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_MockARMContract *MockARMContractFilterer) ParseOwnershipTransferred(log types.Log) (*MockARMContractOwnershipTransferred, error) {
	event := new(MockARMContractOwnershipTransferred)
	if err := _MockARMContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetConfigDetails struct {
	Version     uint32
	BlockNumber uint32
	Config      RMNConfig
}

func (_MockARMContract *MockARMContract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MockARMContract.abi.Events["OwnershipTransferRequested"].ID:
		return _MockARMContract.ParseOwnershipTransferRequested(log)
	case _MockARMContract.abi.Events["OwnershipTransferred"].ID:
		return _MockARMContract.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MockARMContractOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (MockARMContractOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_MockARMContract *MockARMContract) Address() common.Address {
	return _MockARMContract.address
}

type MockARMContractInterface interface {
	GetConfigDetails(opts *bind.CallOpts) (GetConfigDetails,

		error)

	IsBlessed(opts *bind.CallOpts, arg0 IRMNTaggedRoot) (bool, error)

	IsCursed(opts *bind.CallOpts) (bool, error)

	IsCursed0(opts *bind.CallOpts, subject [32]byte) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	OwnerUnvoteToCurse(opts *bind.TransactOpts, arg0 []RMNUnvoteToCurseRecord, subject [32]byte) (*types.Transaction, error)

	OwnerUnvoteToCurse0(opts *bind.TransactOpts, arg0 []RMNUnvoteToCurseRecord) (*types.Transaction, error)

	SetRevert(opts *bind.TransactOpts, err []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	VoteToCurse(opts *bind.TransactOpts, subject [32]byte) (*types.Transaction, error)

	VoteToCurse0(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockARMContractOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MockARMContractOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*MockARMContractOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockARMContractOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockARMContractOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*MockARMContractOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
