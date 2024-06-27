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

var MockARMContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"CustomError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structIRMN.TaggedRoot\",\"name\":\"taggedRoot\",\"type\":\"tuple\"}],\"name\":\"isBlessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"subject\",\"type\":\"bytes16\"}],\"name\":\"isCursed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCursed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"cursed\",\"type\":\"bool\"}],\"name\":\"setChainCursed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"cursed\",\"type\":\"bool\"}],\"name\":\"setGlobalCursed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"revertErr\",\"type\":\"bytes\"}],\"name\":\"setIsCursedRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structIRMN.TaggedRoot\",\"name\":\"taggedRoot\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"blessed\",\"type\":\"bool\"}],\"name\":\"setTaggedRootBlessed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b610add806101576000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c80638da5cb5b11610076578063da558ddd1161005b578063da558ddd1461018a578063dc46bc091461019d578063f2fde38b1461020857600080fd5b80638da5cb5b1461014f578063d6ff2ef61461017757600080fd5b80634d616771116100a75780634d616771146100f357806376eb8ae91461010657806379ba50971461014757600080fd5b80632cbc26bb146100c3578063397796f7146100eb575b600080fd5b6100d66100d1366004610679565b61021b565b60405190151581526020015b60405180910390f35b6100d66102b9565b6100d66101013660046106da565b610310565b61014561011436600461070b565b600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b005b610145610363565b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e2565b610145610185366004610726565b610460565b610145610198366004610798565b610472565b6101456101ab3660046107cc565b60809190911b77ffffffffffffffff0000000000000000000000000000000016600090815260046020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b610145610216366004610805565b6104ed565b6000806002805461022b9061083b565b905011156102715760026040517f5a4ff6710000000000000000000000000000000000000000000000000000000081526004016102689190610888565b60405180910390fd5b60035460ff16806102b357507fffffffffffffffffffffffffffffffff00000000000000000000000000000000821660009081526004602052604090205460ff165b92915050565b600080600280546102c99061083b565b905011156103065760026040517f5a4ff6710000000000000000000000000000000000000000000000000000000081526004016102689190610888565b5060035460ff1690565b60006005816103226020850185610805565b73ffffffffffffffffffffffffffffffffffffffff16815260208082019290925260409081016000908120948301358152939091529091205460ff16919050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146103e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610268565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b600261046d8284836109b5565b505050565b80600560006104846020860186610805565b73ffffffffffffffffffffffffffffffffffffffff168152602080820192909252604090810160009081209583013581529490915290922080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169215159290921790915550565b6104f5610501565b6104fe81610584565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610582576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610268565b565b3373ffffffffffffffffffffffffffffffffffffffff821603610603576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610268565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561068b57600080fd5b81357fffffffffffffffffffffffffffffffff00000000000000000000000000000000811681146106bb57600080fd5b9392505050565b6000604082840312156106d457600080fd5b50919050565b6000604082840312156106ec57600080fd5b6106bb83836106c2565b8035801515811461070657600080fd5b919050565b60006020828403121561071d57600080fd5b6106bb826106f6565b6000806020838503121561073957600080fd5b823567ffffffffffffffff8082111561075157600080fd5b818501915085601f83011261076557600080fd5b81358181111561077457600080fd5b86602082850101111561078657600080fd5b60209290920196919550909350505050565b600080606083850312156107ab57600080fd5b6107b584846106c2565b91506107c3604084016106f6565b90509250929050565b600080604083850312156107df57600080fd5b823567ffffffffffffffff811681146107f757600080fd5b91506107c3602084016106f6565b60006020828403121561081757600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146106bb57600080fd5b600181811c9082168061084f57607f821691505b6020821081036106d4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060208083526000845461089c8161083b565b80602087015260406001808416600081146108be57600181146108f857610928565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00851660408a0152604084151560051b8a01019550610928565b89600052602060002060005b8581101561091f5781548b8201860152908301908801610904565b8a016040019650505b509398975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82111561046d576000816000526020600020601f850160051c8101602086101561098e5750805b601f850160051c820191505b818110156109ad5782815560010161099a565b505050505050565b67ffffffffffffffff8311156109cd576109cd610936565b6109e1836109db835461083b565b83610965565b6000601f841160018114610a3357600085156109fd5750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355610ac9565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b82811015610a825786850135825560209485019460019092019101610a62565b5086821015610abd577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555b505050505056fea164736f6c6343000818000a",
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

func (_MockARMContract *MockARMContractCaller) IsBlessed(opts *bind.CallOpts, taggedRoot IRMNTaggedRoot) (bool, error) {
	var out []interface{}
	err := _MockARMContract.contract.Call(opts, &out, "isBlessed", taggedRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockARMContract *MockARMContractSession) IsBlessed(taggedRoot IRMNTaggedRoot) (bool, error) {
	return _MockARMContract.Contract.IsBlessed(&_MockARMContract.CallOpts, taggedRoot)
}

func (_MockARMContract *MockARMContractCallerSession) IsBlessed(taggedRoot IRMNTaggedRoot) (bool, error) {
	return _MockARMContract.Contract.IsBlessed(&_MockARMContract.CallOpts, taggedRoot)
}

func (_MockARMContract *MockARMContractCaller) IsCursed(opts *bind.CallOpts, subject [16]byte) (bool, error) {
	var out []interface{}
	err := _MockARMContract.contract.Call(opts, &out, "isCursed", subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockARMContract *MockARMContractSession) IsCursed(subject [16]byte) (bool, error) {
	return _MockARMContract.Contract.IsCursed(&_MockARMContract.CallOpts, subject)
}

func (_MockARMContract *MockARMContractCallerSession) IsCursed(subject [16]byte) (bool, error) {
	return _MockARMContract.Contract.IsCursed(&_MockARMContract.CallOpts, subject)
}

func (_MockARMContract *MockARMContractCaller) IsCursed0(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockARMContract.contract.Call(opts, &out, "isCursed0")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockARMContract *MockARMContractSession) IsCursed0() (bool, error) {
	return _MockARMContract.Contract.IsCursed0(&_MockARMContract.CallOpts)
}

func (_MockARMContract *MockARMContractCallerSession) IsCursed0() (bool, error) {
	return _MockARMContract.Contract.IsCursed0(&_MockARMContract.CallOpts)
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

func (_MockARMContract *MockARMContractTransactor) SetChainCursed(opts *bind.TransactOpts, chainSelector uint64, cursed bool) (*types.Transaction, error) {
	return _MockARMContract.contract.Transact(opts, "setChainCursed", chainSelector, cursed)
}

func (_MockARMContract *MockARMContractSession) SetChainCursed(chainSelector uint64, cursed bool) (*types.Transaction, error) {
	return _MockARMContract.Contract.SetChainCursed(&_MockARMContract.TransactOpts, chainSelector, cursed)
}

func (_MockARMContract *MockARMContractTransactorSession) SetChainCursed(chainSelector uint64, cursed bool) (*types.Transaction, error) {
	return _MockARMContract.Contract.SetChainCursed(&_MockARMContract.TransactOpts, chainSelector, cursed)
}

func (_MockARMContract *MockARMContractTransactor) SetGlobalCursed(opts *bind.TransactOpts, cursed bool) (*types.Transaction, error) {
	return _MockARMContract.contract.Transact(opts, "setGlobalCursed", cursed)
}

func (_MockARMContract *MockARMContractSession) SetGlobalCursed(cursed bool) (*types.Transaction, error) {
	return _MockARMContract.Contract.SetGlobalCursed(&_MockARMContract.TransactOpts, cursed)
}

func (_MockARMContract *MockARMContractTransactorSession) SetGlobalCursed(cursed bool) (*types.Transaction, error) {
	return _MockARMContract.Contract.SetGlobalCursed(&_MockARMContract.TransactOpts, cursed)
}

func (_MockARMContract *MockARMContractTransactor) SetIsCursedRevert(opts *bind.TransactOpts, revertErr []byte) (*types.Transaction, error) {
	return _MockARMContract.contract.Transact(opts, "setIsCursedRevert", revertErr)
}

func (_MockARMContract *MockARMContractSession) SetIsCursedRevert(revertErr []byte) (*types.Transaction, error) {
	return _MockARMContract.Contract.SetIsCursedRevert(&_MockARMContract.TransactOpts, revertErr)
}

func (_MockARMContract *MockARMContractTransactorSession) SetIsCursedRevert(revertErr []byte) (*types.Transaction, error) {
	return _MockARMContract.Contract.SetIsCursedRevert(&_MockARMContract.TransactOpts, revertErr)
}

func (_MockARMContract *MockARMContractTransactor) SetTaggedRootBlessed(opts *bind.TransactOpts, taggedRoot IRMNTaggedRoot, blessed bool) (*types.Transaction, error) {
	return _MockARMContract.contract.Transact(opts, "setTaggedRootBlessed", taggedRoot, blessed)
}

func (_MockARMContract *MockARMContractSession) SetTaggedRootBlessed(taggedRoot IRMNTaggedRoot, blessed bool) (*types.Transaction, error) {
	return _MockARMContract.Contract.SetTaggedRootBlessed(&_MockARMContract.TransactOpts, taggedRoot, blessed)
}

func (_MockARMContract *MockARMContractTransactorSession) SetTaggedRootBlessed(taggedRoot IRMNTaggedRoot, blessed bool) (*types.Transaction, error) {
	return _MockARMContract.Contract.SetTaggedRootBlessed(&_MockARMContract.TransactOpts, taggedRoot, blessed)
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
	IsBlessed(opts *bind.CallOpts, taggedRoot IRMNTaggedRoot) (bool, error)

	IsCursed(opts *bind.CallOpts, subject [16]byte) (bool, error)

	IsCursed0(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetChainCursed(opts *bind.TransactOpts, chainSelector uint64, cursed bool) (*types.Transaction, error)

	SetGlobalCursed(opts *bind.TransactOpts, cursed bool) (*types.Transaction, error)

	SetIsCursedRevert(opts *bind.TransactOpts, revertErr []byte) (*types.Transaction, error)

	SetTaggedRootBlessed(opts *bind.TransactOpts, taggedRoot IRMNTaggedRoot, blessed bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockARMContractOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MockARMContractOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*MockARMContractOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockARMContractOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockARMContractOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*MockARMContractOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
