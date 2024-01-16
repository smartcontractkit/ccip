// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimism_l2_bridge_adapter

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

var OptimismL2BridgeAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIWrappedNative\",\"name\":\"wrappedNative\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MsgShouldNotContainValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"depositNativeToL1\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeFeeInNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedNative\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600080546001600160e01b03191673420000000000000000000000000000000000001017905534801561003657600080fd5b50604051610b51380380610b5183398101604081905261005591610066565b6001600160a01b0316608052610096565b60006020828403121561007857600080fd5b81516001600160a01b038116811461008f57600080fd5b9392505050565b608051610a936100be6000396000818160b40152818161014e01526101ce0152610a936000f3fe60806040526004361061003f5760003560e01c80630ff98e31146100445780632e4b1fc91461005957806379a35b4b1461007a578063e861e9071461008d575b600080fd5b6100576100523660046108c1565b6100de565b005b34801561006557600080fd5b50604051600081526020015b60405180910390f35b6100576100883660046108e3565b6100eb565b34801561009957600080fd5b5060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610071565b6100e881346103f1565b50565b341561012a576040517f2543d86e0000000000000000000000000000000000000000000000000000000081523460048201526024015b60405180910390fd5b61014c73ffffffffffffffffffffffffffffffffffffffff8416333084610507565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361024e576040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632e1a7d4d90602401600060405180830381600087803b15801561022757600080fd5b505af115801561023b573d6000803e3d6000fd5b5050505061024982826103f1565b6103eb565b6000546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152602481018390529084169063095ea7b3906044016020604051808303816000875af11580156102c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102eb919061092e565b506000805473ffffffffffffffffffffffffffffffffffffffff81169163a3a79548918691869186919074010000000000000000000000000000000000000000900467ffffffffffffffff1681601461034383610950565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550604051602001610389919067ffffffffffffffff91909116815260200190565b6040516020818303038152906040526040518663ffffffff1660e01b81526004016103b8959493929190610a0c565b600060405180830381600087803b1580156103d257600080fd5b505af11580156103e6573d6000803e3d6000fd5b505050505b50505050565b6000805473ffffffffffffffffffffffffffffffffffffffff81169163a3a795489173deaddeaddeaddeaddeaddeaddeaddeaddead000091869186919074010000000000000000000000000000000000000000900467ffffffffffffffff1681601461045c83610950565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040516020016104a2919067ffffffffffffffff91909116815260200190565b6040516020818303038152906040526040518663ffffffff1660e01b81526004016104d1959493929190610a0c565b600060405180830381600087803b1580156104eb57600080fd5b505af11580156104ff573d6000803e3d6000fd5b505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff8581166024830152848116604483015260648083018590528351808403909101815260849092018352602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564908401526103eb928792916000916105da918516908490610689565b80519091501561068457808060200190518101906105f8919061092e565b610684576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610121565b505050565b606061069884846000856106a0565b949350505050565b606082471015610732576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610121565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161075b9190610a57565b60006040518083038185875af1925050503d8060008114610798576040519150601f19603f3d011682016040523d82523d6000602084013e61079d565b606091505b50915091506107ae878383876107b9565b979650505050505050565b6060831561084f5782516000036108485773ffffffffffffffffffffffffffffffffffffffff85163b610848576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610121565b5081610698565b61069883838151156108645781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101219190610a73565b803573ffffffffffffffffffffffffffffffffffffffff811681146108bc57600080fd5b919050565b6000602082840312156108d357600080fd5b6108dc82610898565b9392505050565b600080600080608085870312156108f957600080fd5b61090285610898565b935061091060208601610898565b925061091e60408601610898565b9396929550929360600135925050565b60006020828403121561094057600080fd5b815180151581146108dc57600080fd5b600067ffffffffffffffff808316818103610994577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6001019392505050565b60005b838110156109b95781810151838201526020016109a1565b50506000910152565b600081518084526109da81602086016020860161099e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600073ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525084604083015263ffffffff8416606083015260a060808301526107ae60a08301846109c2565b60008251610a6981846020870161099e565b9190910192915050565b6020815260006108dc60208301846109c256fea164736f6c6343000813000a",
}

var OptimismL2BridgeAdapterABI = OptimismL2BridgeAdapterMetaData.ABI

var OptimismL2BridgeAdapterBin = OptimismL2BridgeAdapterMetaData.Bin

func DeployOptimismL2BridgeAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, wrappedNative common.Address) (common.Address, *types.Transaction, *OptimismL2BridgeAdapter, error) {
	parsed, err := OptimismL2BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptimismL2BridgeAdapterBin), backend, wrappedNative)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptimismL2BridgeAdapter{address: address, abi: *parsed, OptimismL2BridgeAdapterCaller: OptimismL2BridgeAdapterCaller{contract: contract}, OptimismL2BridgeAdapterTransactor: OptimismL2BridgeAdapterTransactor{contract: contract}, OptimismL2BridgeAdapterFilterer: OptimismL2BridgeAdapterFilterer{contract: contract}}, nil
}

type OptimismL2BridgeAdapter struct {
	address common.Address
	abi     abi.ABI
	OptimismL2BridgeAdapterCaller
	OptimismL2BridgeAdapterTransactor
	OptimismL2BridgeAdapterFilterer
}

type OptimismL2BridgeAdapterCaller struct {
	contract *bind.BoundContract
}

type OptimismL2BridgeAdapterTransactor struct {
	contract *bind.BoundContract
}

type OptimismL2BridgeAdapterFilterer struct {
	contract *bind.BoundContract
}

type OptimismL2BridgeAdapterSession struct {
	Contract     *OptimismL2BridgeAdapter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OptimismL2BridgeAdapterCallerSession struct {
	Contract *OptimismL2BridgeAdapterCaller
	CallOpts bind.CallOpts
}

type OptimismL2BridgeAdapterTransactorSession struct {
	Contract     *OptimismL2BridgeAdapterTransactor
	TransactOpts bind.TransactOpts
}

type OptimismL2BridgeAdapterRaw struct {
	Contract *OptimismL2BridgeAdapter
}

type OptimismL2BridgeAdapterCallerRaw struct {
	Contract *OptimismL2BridgeAdapterCaller
}

type OptimismL2BridgeAdapterTransactorRaw struct {
	Contract *OptimismL2BridgeAdapterTransactor
}

func NewOptimismL2BridgeAdapter(address common.Address, backend bind.ContractBackend) (*OptimismL2BridgeAdapter, error) {
	abi, err := abi.JSON(strings.NewReader(OptimismL2BridgeAdapterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOptimismL2BridgeAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimismL2BridgeAdapter{address: address, abi: abi, OptimismL2BridgeAdapterCaller: OptimismL2BridgeAdapterCaller{contract: contract}, OptimismL2BridgeAdapterTransactor: OptimismL2BridgeAdapterTransactor{contract: contract}, OptimismL2BridgeAdapterFilterer: OptimismL2BridgeAdapterFilterer{contract: contract}}, nil
}

func NewOptimismL2BridgeAdapterCaller(address common.Address, caller bind.ContractCaller) (*OptimismL2BridgeAdapterCaller, error) {
	contract, err := bindOptimismL2BridgeAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismL2BridgeAdapterCaller{contract: contract}, nil
}

func NewOptimismL2BridgeAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimismL2BridgeAdapterTransactor, error) {
	contract, err := bindOptimismL2BridgeAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismL2BridgeAdapterTransactor{contract: contract}, nil
}

func NewOptimismL2BridgeAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimismL2BridgeAdapterFilterer, error) {
	contract, err := bindOptimismL2BridgeAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimismL2BridgeAdapterFilterer{contract: contract}, nil
}

func bindOptimismL2BridgeAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimismL2BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismL2BridgeAdapter.Contract.OptimismL2BridgeAdapterCaller.contract.Call(opts, result, method, params...)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismL2BridgeAdapter.Contract.OptimismL2BridgeAdapterTransactor.contract.Transfer(opts)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismL2BridgeAdapter.Contract.OptimismL2BridgeAdapterTransactor.contract.Transact(opts, method, params...)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismL2BridgeAdapter.Contract.contract.Call(opts, result, method, params...)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismL2BridgeAdapter.Contract.contract.Transfer(opts)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismL2BridgeAdapter.Contract.contract.Transact(opts, method, params...)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterCaller) GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OptimismL2BridgeAdapter.contract.Call(opts, &out, "getBridgeFeeInNative")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _OptimismL2BridgeAdapter.Contract.GetBridgeFeeInNative(&_OptimismL2BridgeAdapter.CallOpts)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterCallerSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _OptimismL2BridgeAdapter.Contract.GetBridgeFeeInNative(&_OptimismL2BridgeAdapter.CallOpts)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterCaller) GetWrappedNative(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismL2BridgeAdapter.contract.Call(opts, &out, "getWrappedNative")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterSession) GetWrappedNative() (common.Address, error) {
	return _OptimismL2BridgeAdapter.Contract.GetWrappedNative(&_OptimismL2BridgeAdapter.CallOpts)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterCallerSession) GetWrappedNative() (common.Address, error) {
	return _OptimismL2BridgeAdapter.Contract.GetWrappedNative(&_OptimismL2BridgeAdapter.CallOpts)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterTransactor) DepositNativeToL1(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _OptimismL2BridgeAdapter.contract.Transact(opts, "depositNativeToL1", recipient)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterSession) DepositNativeToL1(recipient common.Address) (*types.Transaction, error) {
	return _OptimismL2BridgeAdapter.Contract.DepositNativeToL1(&_OptimismL2BridgeAdapter.TransactOpts, recipient)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterTransactorSession) DepositNativeToL1(recipient common.Address) (*types.Transaction, error) {
	return _OptimismL2BridgeAdapter.Contract.DepositNativeToL1(&_OptimismL2BridgeAdapter.TransactOpts, recipient)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterTransactor) SendERC20(opts *bind.TransactOpts, arg0 common.Address, l2Token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OptimismL2BridgeAdapter.contract.Transact(opts, "sendERC20", arg0, l2Token, recipient, amount)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterSession) SendERC20(arg0 common.Address, l2Token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OptimismL2BridgeAdapter.Contract.SendERC20(&_OptimismL2BridgeAdapter.TransactOpts, arg0, l2Token, recipient, amount)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterTransactorSession) SendERC20(arg0 common.Address, l2Token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OptimismL2BridgeAdapter.Contract.SendERC20(&_OptimismL2BridgeAdapter.TransactOpts, arg0, l2Token, recipient, amount)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapter) Address() common.Address {
	return _OptimismL2BridgeAdapter.address
}

type OptimismL2BridgeAdapterInterface interface {
	GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error)

	GetWrappedNative(opts *bind.CallOpts) (common.Address, error)

	DepositNativeToL1(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, arg0 common.Address, l2Token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	Address() common.Address
}
