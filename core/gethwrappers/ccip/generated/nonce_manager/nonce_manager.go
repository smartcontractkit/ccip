// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nonce_manager

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

type NonceManagerPreviousRamps struct {
	PrevOnRamp common.Address
}

type NonceManagerPreviousRampsArgs struct {
	ChainSelector uint64
	PrevRamps     NonceManagerPreviousRamps
}

var NonceManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidRampUpdate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOnRamp\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"}],\"name\":\"PreviousOnRampUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"}],\"internalType\":\"structNonceManager.PreviousRamps\",\"name\":\"prevRamps\",\"type\":\"tuple\"}],\"internalType\":\"structNonceManager.PreviousRampsArgs[]\",\"name\":\"prevRampsArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyRampUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"name\":\"getOutboundNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"getPrevRamps\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"}],\"internalType\":\"structNonceManager.PreviousRamps\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"}],\"name\":\"incrementOutboundNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b610bdb806101576000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80637a2dbcbc1161005b5780637a2dbcbc146101855780638da5cb5b146101985780638f1eb7e6146101b6578063f2fde38b146101c957600080fd5b806322d94aa71461008d57806331b89ff3146101105780634ce927231461013c57806379ba50971461017b575b600080fd5b6100e661009b3660046109cd565b604080516020808201835260009182905267ffffffffffffffff9390931681526003835281902081519283019091525473ffffffffffffffffffffffffffffffffffffffff16815290565b604051905173ffffffffffffffffffffffffffffffffffffffff1681526020015b60405180910390f35b61012361011e3660046109ea565b6101dc565b60405167ffffffffffffffff9091168152602001610107565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610107565b61018361032d565b005b6101236101933660046109ea565b61042f565b60005473ffffffffffffffffffffffffffffffffffffffff16610156565b6101836101c4366004610a91565b610622565b6101836101d7366004610b06565b61082b565b67ffffffffffffffff831660009081526004602052604080822090518291906102089086908690610b23565b9081526040519081900360200190205467ffffffffffffffff16905060008190036103235767ffffffffffffffff851660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1680156103215773ffffffffffffffffffffffffffffffffffffffff811663856c824761028b86880188610b06565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381865afa1580156102f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103189190610b33565b92505050610326565b505b90505b9392505050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146103b3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60025460009073ffffffffffffffffffffffffffffffffffffffff163314610483576040517fdaee2fdc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff841660009081526004602052604080822090516104ac9086908690610b23565b90815260405190819003602001902080549091506000906104d89067ffffffffffffffff166001610b50565b90508067ffffffffffffffff166000036105e75767ffffffffffffffff861660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1680156105e55773ffffffffffffffffffffffffffffffffffffffff811663856c824761054a87890189610b06565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381865afa1580156105b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105d79190610b33565b6105e2906001610b50565b91505b505b81547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff821617909155949350505050565b61062a61083f565b73ffffffffffffffffffffffffffffffffffffffff8316156106bd57600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091556040519081527fb3d81ef1b70b9650b688ffeba63e1b58c36aa4c936ba7d84a295f3b16045b3149060200160405180910390a15b60005b8181101561082557368383838181106106db576106db610b9f565b604002919091019150600090506003816106f860208501856109cd565b67ffffffffffffffff1681526020810191909152604001600020805490915073ffffffffffffffffffffffffffffffffffffffff1615610764576040517fbb16aec100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107746040830160208401610b06565b81547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff919091161781556107c160208301836109cd565b815460405173ffffffffffffffffffffffffffffffffffffffff909116815267ffffffffffffffff91909116907f89d2355e2829b1e15855fec87fb400638aebc9f03728949d702d3b5d4ea999549060200160405180910390a250506001016106c0565b50505050565b61083361083f565b61083c816108c2565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146108c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016103aa565b565b3373ffffffffffffffffffffffffffffffffffffffff821603610941576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016103aa565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b67ffffffffffffffff8116811461083c57600080fd5b6000602082840312156109df57600080fd5b8135610326816109b7565b6000806000604084860312156109ff57600080fd5b8335610a0a816109b7565b9250602084013567ffffffffffffffff80821115610a2757600080fd5b818601915086601f830112610a3b57600080fd5b813581811115610a4a57600080fd5b876020828501011115610a5c57600080fd5b6020830194508093505050509250925092565b73ffffffffffffffffffffffffffffffffffffffff8116811461083c57600080fd5b600080600060408486031215610aa657600080fd5b8335610ab181610a6f565b9250602084013567ffffffffffffffff80821115610ace57600080fd5b818601915086601f830112610ae257600080fd5b813581811115610af157600080fd5b8760208260061b8501011115610a5c57600080fd5b600060208284031215610b1857600080fd5b813561032681610a6f565b8183823760009101908152919050565b600060208284031215610b4557600080fd5b8151610326816109b7565b67ffffffffffffffff818116838216019080821115610b98577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c6343000818000a",
}

var NonceManagerABI = NonceManagerMetaData.ABI

var NonceManagerBin = NonceManagerMetaData.Bin

func DeployNonceManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NonceManager, error) {
	parsed, err := NonceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NonceManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NonceManager{address: address, abi: *parsed, NonceManagerCaller: NonceManagerCaller{contract: contract}, NonceManagerTransactor: NonceManagerTransactor{contract: contract}, NonceManagerFilterer: NonceManagerFilterer{contract: contract}}, nil
}

type NonceManager struct {
	address common.Address
	abi     abi.ABI
	NonceManagerCaller
	NonceManagerTransactor
	NonceManagerFilterer
}

type NonceManagerCaller struct {
	contract *bind.BoundContract
}

type NonceManagerTransactor struct {
	contract *bind.BoundContract
}

type NonceManagerFilterer struct {
	contract *bind.BoundContract
}

type NonceManagerSession struct {
	Contract     *NonceManager
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type NonceManagerCallerSession struct {
	Contract *NonceManagerCaller
	CallOpts bind.CallOpts
}

type NonceManagerTransactorSession struct {
	Contract     *NonceManagerTransactor
	TransactOpts bind.TransactOpts
}

type NonceManagerRaw struct {
	Contract *NonceManager
}

type NonceManagerCallerRaw struct {
	Contract *NonceManagerCaller
}

type NonceManagerTransactorRaw struct {
	Contract *NonceManagerTransactor
}

func NewNonceManager(address common.Address, backend bind.ContractBackend) (*NonceManager, error) {
	abi, err := abi.JSON(strings.NewReader(NonceManagerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindNonceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NonceManager{address: address, abi: abi, NonceManagerCaller: NonceManagerCaller{contract: contract}, NonceManagerTransactor: NonceManagerTransactor{contract: contract}, NonceManagerFilterer: NonceManagerFilterer{contract: contract}}, nil
}

func NewNonceManagerCaller(address common.Address, caller bind.ContractCaller) (*NonceManagerCaller, error) {
	contract, err := bindNonceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NonceManagerCaller{contract: contract}, nil
}

func NewNonceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NonceManagerTransactor, error) {
	contract, err := bindNonceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NonceManagerTransactor{contract: contract}, nil
}

func NewNonceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NonceManagerFilterer, error) {
	contract, err := bindNonceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NonceManagerFilterer{contract: contract}, nil
}

func bindNonceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NonceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_NonceManager *NonceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonceManager.Contract.NonceManagerCaller.contract.Call(opts, result, method, params...)
}

func (_NonceManager *NonceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonceManager.Contract.NonceManagerTransactor.contract.Transfer(opts)
}

func (_NonceManager *NonceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonceManager.Contract.NonceManagerTransactor.contract.Transact(opts, method, params...)
}

func (_NonceManager *NonceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonceManager.Contract.contract.Call(opts, result, method, params...)
}

func (_NonceManager *NonceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonceManager.Contract.contract.Transfer(opts)
}

func (_NonceManager *NonceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonceManager.Contract.contract.Transact(opts, method, params...)
}

func (_NonceManager *NonceManagerCaller) GetOnRamp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NonceManager.contract.Call(opts, &out, "getOnRamp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_NonceManager *NonceManagerSession) GetOnRamp() (common.Address, error) {
	return _NonceManager.Contract.GetOnRamp(&_NonceManager.CallOpts)
}

func (_NonceManager *NonceManagerCallerSession) GetOnRamp() (common.Address, error) {
	return _NonceManager.Contract.GetOnRamp(&_NonceManager.CallOpts)
}

func (_NonceManager *NonceManagerCaller) GetOutboundNonce(opts *bind.CallOpts, destChainSelector uint64, sender []byte) (uint64, error) {
	var out []interface{}
	err := _NonceManager.contract.Call(opts, &out, "getOutboundNonce", destChainSelector, sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_NonceManager *NonceManagerSession) GetOutboundNonce(destChainSelector uint64, sender []byte) (uint64, error) {
	return _NonceManager.Contract.GetOutboundNonce(&_NonceManager.CallOpts, destChainSelector, sender)
}

func (_NonceManager *NonceManagerCallerSession) GetOutboundNonce(destChainSelector uint64, sender []byte) (uint64, error) {
	return _NonceManager.Contract.GetOutboundNonce(&_NonceManager.CallOpts, destChainSelector, sender)
}

func (_NonceManager *NonceManagerCaller) GetPrevRamps(opts *bind.CallOpts, chainSelector uint64) (NonceManagerPreviousRamps, error) {
	var out []interface{}
	err := _NonceManager.contract.Call(opts, &out, "getPrevRamps", chainSelector)

	if err != nil {
		return *new(NonceManagerPreviousRamps), err
	}

	out0 := *abi.ConvertType(out[0], new(NonceManagerPreviousRamps)).(*NonceManagerPreviousRamps)

	return out0, err

}

func (_NonceManager *NonceManagerSession) GetPrevRamps(chainSelector uint64) (NonceManagerPreviousRamps, error) {
	return _NonceManager.Contract.GetPrevRamps(&_NonceManager.CallOpts, chainSelector)
}

func (_NonceManager *NonceManagerCallerSession) GetPrevRamps(chainSelector uint64) (NonceManagerPreviousRamps, error) {
	return _NonceManager.Contract.GetPrevRamps(&_NonceManager.CallOpts, chainSelector)
}

func (_NonceManager *NonceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NonceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_NonceManager *NonceManagerSession) Owner() (common.Address, error) {
	return _NonceManager.Contract.Owner(&_NonceManager.CallOpts)
}

func (_NonceManager *NonceManagerCallerSession) Owner() (common.Address, error) {
	return _NonceManager.Contract.Owner(&_NonceManager.CallOpts)
}

func (_NonceManager *NonceManagerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonceManager.contract.Transact(opts, "acceptOwnership")
}

func (_NonceManager *NonceManagerSession) AcceptOwnership() (*types.Transaction, error) {
	return _NonceManager.Contract.AcceptOwnership(&_NonceManager.TransactOpts)
}

func (_NonceManager *NonceManagerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _NonceManager.Contract.AcceptOwnership(&_NonceManager.TransactOpts)
}

func (_NonceManager *NonceManagerTransactor) ApplyRampUpdates(opts *bind.TransactOpts, onRamp common.Address, prevRampsArgs []NonceManagerPreviousRampsArgs) (*types.Transaction, error) {
	return _NonceManager.contract.Transact(opts, "applyRampUpdates", onRamp, prevRampsArgs)
}

func (_NonceManager *NonceManagerSession) ApplyRampUpdates(onRamp common.Address, prevRampsArgs []NonceManagerPreviousRampsArgs) (*types.Transaction, error) {
	return _NonceManager.Contract.ApplyRampUpdates(&_NonceManager.TransactOpts, onRamp, prevRampsArgs)
}

func (_NonceManager *NonceManagerTransactorSession) ApplyRampUpdates(onRamp common.Address, prevRampsArgs []NonceManagerPreviousRampsArgs) (*types.Transaction, error) {
	return _NonceManager.Contract.ApplyRampUpdates(&_NonceManager.TransactOpts, onRamp, prevRampsArgs)
}

func (_NonceManager *NonceManagerTransactor) IncrementOutboundNonce(opts *bind.TransactOpts, destChainSelector uint64, sender []byte) (*types.Transaction, error) {
	return _NonceManager.contract.Transact(opts, "incrementOutboundNonce", destChainSelector, sender)
}

func (_NonceManager *NonceManagerSession) IncrementOutboundNonce(destChainSelector uint64, sender []byte) (*types.Transaction, error) {
	return _NonceManager.Contract.IncrementOutboundNonce(&_NonceManager.TransactOpts, destChainSelector, sender)
}

func (_NonceManager *NonceManagerTransactorSession) IncrementOutboundNonce(destChainSelector uint64, sender []byte) (*types.Transaction, error) {
	return _NonceManager.Contract.IncrementOutboundNonce(&_NonceManager.TransactOpts, destChainSelector, sender)
}

func (_NonceManager *NonceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _NonceManager.contract.Transact(opts, "transferOwnership", to)
}

func (_NonceManager *NonceManagerSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _NonceManager.Contract.TransferOwnership(&_NonceManager.TransactOpts, to)
}

func (_NonceManager *NonceManagerTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _NonceManager.Contract.TransferOwnership(&_NonceManager.TransactOpts, to)
}

type NonceManagerOnRampUpdatedIterator struct {
	Event *NonceManagerOnRampUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NonceManagerOnRampUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NonceManagerOnRampUpdated)
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
		it.Event = new(NonceManagerOnRampUpdated)
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

func (it *NonceManagerOnRampUpdatedIterator) Error() error {
	return it.fail
}

func (it *NonceManagerOnRampUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NonceManagerOnRampUpdated struct {
	OnRamp common.Address
	Raw    types.Log
}

func (_NonceManager *NonceManagerFilterer) FilterOnRampUpdated(opts *bind.FilterOpts) (*NonceManagerOnRampUpdatedIterator, error) {

	logs, sub, err := _NonceManager.contract.FilterLogs(opts, "OnRampUpdated")
	if err != nil {
		return nil, err
	}
	return &NonceManagerOnRampUpdatedIterator{contract: _NonceManager.contract, event: "OnRampUpdated", logs: logs, sub: sub}, nil
}

func (_NonceManager *NonceManagerFilterer) WatchOnRampUpdated(opts *bind.WatchOpts, sink chan<- *NonceManagerOnRampUpdated) (event.Subscription, error) {

	logs, sub, err := _NonceManager.contract.WatchLogs(opts, "OnRampUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NonceManagerOnRampUpdated)
				if err := _NonceManager.contract.UnpackLog(event, "OnRampUpdated", log); err != nil {
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

func (_NonceManager *NonceManagerFilterer) ParseOnRampUpdated(log types.Log) (*NonceManagerOnRampUpdated, error) {
	event := new(NonceManagerOnRampUpdated)
	if err := _NonceManager.contract.UnpackLog(event, "OnRampUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NonceManagerOwnershipTransferRequestedIterator struct {
	Event *NonceManagerOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NonceManagerOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NonceManagerOwnershipTransferRequested)
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
		it.Event = new(NonceManagerOwnershipTransferRequested)
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

func (it *NonceManagerOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *NonceManagerOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NonceManagerOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_NonceManager *NonceManagerFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NonceManagerOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NonceManager.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NonceManagerOwnershipTransferRequestedIterator{contract: _NonceManager.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_NonceManager *NonceManagerFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *NonceManagerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NonceManager.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NonceManagerOwnershipTransferRequested)
				if err := _NonceManager.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_NonceManager *NonceManagerFilterer) ParseOwnershipTransferRequested(log types.Log) (*NonceManagerOwnershipTransferRequested, error) {
	event := new(NonceManagerOwnershipTransferRequested)
	if err := _NonceManager.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NonceManagerOwnershipTransferredIterator struct {
	Event *NonceManagerOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NonceManagerOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NonceManagerOwnershipTransferred)
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
		it.Event = new(NonceManagerOwnershipTransferred)
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

func (it *NonceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *NonceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NonceManagerOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_NonceManager *NonceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NonceManagerOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NonceManager.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NonceManagerOwnershipTransferredIterator{contract: _NonceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_NonceManager *NonceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NonceManagerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NonceManager.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NonceManagerOwnershipTransferred)
				if err := _NonceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_NonceManager *NonceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*NonceManagerOwnershipTransferred, error) {
	event := new(NonceManagerOwnershipTransferred)
	if err := _NonceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NonceManagerPreviousOnRampUpdatedIterator struct {
	Event *NonceManagerPreviousOnRampUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NonceManagerPreviousOnRampUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NonceManagerPreviousOnRampUpdated)
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
		it.Event = new(NonceManagerPreviousOnRampUpdated)
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

func (it *NonceManagerPreviousOnRampUpdatedIterator) Error() error {
	return it.fail
}

func (it *NonceManagerPreviousOnRampUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NonceManagerPreviousOnRampUpdated struct {
	DestChainSelector uint64
	PrevOnRamp        common.Address
	Raw               types.Log
}

func (_NonceManager *NonceManagerFilterer) FilterPreviousOnRampUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*NonceManagerPreviousOnRampUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _NonceManager.contract.FilterLogs(opts, "PreviousOnRampUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &NonceManagerPreviousOnRampUpdatedIterator{contract: _NonceManager.contract, event: "PreviousOnRampUpdated", logs: logs, sub: sub}, nil
}

func (_NonceManager *NonceManagerFilterer) WatchPreviousOnRampUpdated(opts *bind.WatchOpts, sink chan<- *NonceManagerPreviousOnRampUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _NonceManager.contract.WatchLogs(opts, "PreviousOnRampUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NonceManagerPreviousOnRampUpdated)
				if err := _NonceManager.contract.UnpackLog(event, "PreviousOnRampUpdated", log); err != nil {
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

func (_NonceManager *NonceManagerFilterer) ParsePreviousOnRampUpdated(log types.Log) (*NonceManagerPreviousOnRampUpdated, error) {
	event := new(NonceManagerPreviousOnRampUpdated)
	if err := _NonceManager.contract.UnpackLog(event, "PreviousOnRampUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_NonceManager *NonceManager) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _NonceManager.abi.Events["OnRampUpdated"].ID:
		return _NonceManager.ParseOnRampUpdated(log)
	case _NonceManager.abi.Events["OwnershipTransferRequested"].ID:
		return _NonceManager.ParseOwnershipTransferRequested(log)
	case _NonceManager.abi.Events["OwnershipTransferred"].ID:
		return _NonceManager.ParseOwnershipTransferred(log)
	case _NonceManager.abi.Events["PreviousOnRampUpdated"].ID:
		return _NonceManager.ParsePreviousOnRampUpdated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (NonceManagerOnRampUpdated) Topic() common.Hash {
	return common.HexToHash("0xb3d81ef1b70b9650b688ffeba63e1b58c36aa4c936ba7d84a295f3b16045b314")
}

func (NonceManagerOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (NonceManagerOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (NonceManagerPreviousOnRampUpdated) Topic() common.Hash {
	return common.HexToHash("0x89d2355e2829b1e15855fec87fb400638aebc9f03728949d702d3b5d4ea99954")
}

func (_NonceManager *NonceManager) Address() common.Address {
	return _NonceManager.address
}

type NonceManagerInterface interface {
	GetOnRamp(opts *bind.CallOpts) (common.Address, error)

	GetOutboundNonce(opts *bind.CallOpts, destChainSelector uint64, sender []byte) (uint64, error)

	GetPrevRamps(opts *bind.CallOpts, chainSelector uint64) (NonceManagerPreviousRamps, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyRampUpdates(opts *bind.TransactOpts, onRamp common.Address, prevRampsArgs []NonceManagerPreviousRampsArgs) (*types.Transaction, error)

	IncrementOutboundNonce(opts *bind.TransactOpts, destChainSelector uint64, sender []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOnRampUpdated(opts *bind.FilterOpts) (*NonceManagerOnRampUpdatedIterator, error)

	WatchOnRampUpdated(opts *bind.WatchOpts, sink chan<- *NonceManagerOnRampUpdated) (event.Subscription, error)

	ParseOnRampUpdated(log types.Log) (*NonceManagerOnRampUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NonceManagerOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *NonceManagerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*NonceManagerOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NonceManagerOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NonceManagerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*NonceManagerOwnershipTransferred, error)

	FilterPreviousOnRampUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*NonceManagerPreviousOnRampUpdatedIterator, error)

	WatchPreviousOnRampUpdated(opts *bind.WatchOpts, sink chan<- *NonceManagerPreviousOnRampUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParsePreviousOnRampUpdated(log types.Log) (*NonceManagerPreviousOnRampUpdated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
