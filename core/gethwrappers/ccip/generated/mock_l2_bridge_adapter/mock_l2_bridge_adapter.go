// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_l2_bridge_adapter

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

var MockL2BridgeAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MsgShouldNotContainValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdraw\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeFeeInNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610352806100206000396000f3fe6080604052600436106100345760003560e01c80630184f846146100395780632e4b1fc91461005b57806379a35b4b1461007c575b600080fd5b34801561004557600080fd5b506100596100543660046101cf565b61008f565b005b34801561006757600080fd5b50600060405190815260200160405180910390f35b61005961008a3660046102d1565b6100fc565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100f39060208082526004908201527f746f646f00000000000000000000000000000000000000000000000000000000604082015260600190565b60405180910390fd5b6040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810182905273ffffffffffffffffffffffffffffffffffffffff8416906323b872dd906064016020604051808303816000875af1158015610175573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610199919061031c565b5050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080604083850312156101e257600080fd5b82359150602083013567ffffffffffffffff8082111561020157600080fd5b818501915085601f83011261021557600080fd5b813581811115610227576102276101a0565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561026d5761026d6101a0565b8160405282815288602084870101111561028657600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b803573ffffffffffffffffffffffffffffffffffffffff811681146102cc57600080fd5b919050565b600080600080608085870312156102e757600080fd5b6102f0856102a8565b93506102fe602086016102a8565b925061030c604086016102a8565b9396929550929360600135925050565b60006020828403121561032e57600080fd5b8151801515811461033e57600080fd5b939250505056fea164736f6c6343000813000a",
}

var MockL2BridgeAdapterABI = MockL2BridgeAdapterMetaData.ABI

var MockL2BridgeAdapterBin = MockL2BridgeAdapterMetaData.Bin

func DeployMockL2BridgeAdapter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockL2BridgeAdapter, error) {
	parsed, err := MockL2BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockL2BridgeAdapterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockL2BridgeAdapter{address: address, abi: *parsed, MockL2BridgeAdapterCaller: MockL2BridgeAdapterCaller{contract: contract}, MockL2BridgeAdapterTransactor: MockL2BridgeAdapterTransactor{contract: contract}, MockL2BridgeAdapterFilterer: MockL2BridgeAdapterFilterer{contract: contract}}, nil
}

type MockL2BridgeAdapter struct {
	address common.Address
	abi     abi.ABI
	MockL2BridgeAdapterCaller
	MockL2BridgeAdapterTransactor
	MockL2BridgeAdapterFilterer
}

type MockL2BridgeAdapterCaller struct {
	contract *bind.BoundContract
}

type MockL2BridgeAdapterTransactor struct {
	contract *bind.BoundContract
}

type MockL2BridgeAdapterFilterer struct {
	contract *bind.BoundContract
}

type MockL2BridgeAdapterSession struct {
	Contract     *MockL2BridgeAdapter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockL2BridgeAdapterCallerSession struct {
	Contract *MockL2BridgeAdapterCaller
	CallOpts bind.CallOpts
}

type MockL2BridgeAdapterTransactorSession struct {
	Contract     *MockL2BridgeAdapterTransactor
	TransactOpts bind.TransactOpts
}

type MockL2BridgeAdapterRaw struct {
	Contract *MockL2BridgeAdapter
}

type MockL2BridgeAdapterCallerRaw struct {
	Contract *MockL2BridgeAdapterCaller
}

type MockL2BridgeAdapterTransactorRaw struct {
	Contract *MockL2BridgeAdapterTransactor
}

func NewMockL2BridgeAdapter(address common.Address, backend bind.ContractBackend) (*MockL2BridgeAdapter, error) {
	abi, err := abi.JSON(strings.NewReader(MockL2BridgeAdapterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockL2BridgeAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockL2BridgeAdapter{address: address, abi: abi, MockL2BridgeAdapterCaller: MockL2BridgeAdapterCaller{contract: contract}, MockL2BridgeAdapterTransactor: MockL2BridgeAdapterTransactor{contract: contract}, MockL2BridgeAdapterFilterer: MockL2BridgeAdapterFilterer{contract: contract}}, nil
}

func NewMockL2BridgeAdapterCaller(address common.Address, caller bind.ContractCaller) (*MockL2BridgeAdapterCaller, error) {
	contract, err := bindMockL2BridgeAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockL2BridgeAdapterCaller{contract: contract}, nil
}

func NewMockL2BridgeAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*MockL2BridgeAdapterTransactor, error) {
	contract, err := bindMockL2BridgeAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockL2BridgeAdapterTransactor{contract: contract}, nil
}

func NewMockL2BridgeAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*MockL2BridgeAdapterFilterer, error) {
	contract, err := bindMockL2BridgeAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockL2BridgeAdapterFilterer{contract: contract}, nil
}

func bindMockL2BridgeAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockL2BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL2BridgeAdapter.Contract.MockL2BridgeAdapterCaller.contract.Call(opts, result, method, params...)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.Contract.MockL2BridgeAdapterTransactor.contract.Transfer(opts)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.Contract.MockL2BridgeAdapterTransactor.contract.Transact(opts, method, params...)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL2BridgeAdapter.Contract.contract.Call(opts, result, method, params...)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.Contract.contract.Transfer(opts)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.Contract.contract.Transact(opts, method, params...)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterCaller) FinalizeWithdraw(opts *bind.CallOpts, arg0 *big.Int, arg1 []byte) error {
	var out []interface{}
	err := _MockL2BridgeAdapter.contract.Call(opts, &out, "finalizeWithdraw", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterSession) FinalizeWithdraw(arg0 *big.Int, arg1 []byte) error {
	return _MockL2BridgeAdapter.Contract.FinalizeWithdraw(&_MockL2BridgeAdapter.CallOpts, arg0, arg1)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterCallerSession) FinalizeWithdraw(arg0 *big.Int, arg1 []byte) error {
	return _MockL2BridgeAdapter.Contract.FinalizeWithdraw(&_MockL2BridgeAdapter.CallOpts, arg0, arg1)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterCaller) GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockL2BridgeAdapter.contract.Call(opts, &out, "getBridgeFeeInNative")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _MockL2BridgeAdapter.Contract.GetBridgeFeeInNative(&_MockL2BridgeAdapter.CallOpts)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterCallerSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _MockL2BridgeAdapter.Contract.GetBridgeFeeInNative(&_MockL2BridgeAdapter.CallOpts)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterTransactor) SendERC20(opts *bind.TransactOpts, arg0 common.Address, l2token common.Address, arg2 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.contract.Transact(opts, "sendERC20", arg0, l2token, arg2, amount)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterSession) SendERC20(arg0 common.Address, l2token common.Address, arg2 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.Contract.SendERC20(&_MockL2BridgeAdapter.TransactOpts, arg0, l2token, arg2, amount)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapterTransactorSession) SendERC20(arg0 common.Address, l2token common.Address, arg2 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockL2BridgeAdapter.Contract.SendERC20(&_MockL2BridgeAdapter.TransactOpts, arg0, l2token, arg2, amount)
}

func (_MockL2BridgeAdapter *MockL2BridgeAdapter) Address() common.Address {
	return _MockL2BridgeAdapter.address
}

type MockL2BridgeAdapterInterface interface {
	FinalizeWithdraw(opts *bind.CallOpts, arg0 *big.Int, arg1 []byte) error

	GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error)

	SendERC20(opts *bind.TransactOpts, arg0 common.Address, l2token common.Address, arg2 common.Address, amount *big.Int) (*types.Transaction, error)

	Address() common.Address
}
