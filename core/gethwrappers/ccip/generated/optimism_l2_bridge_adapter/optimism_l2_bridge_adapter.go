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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIWrappedNative\",\"name\":\"wrappedNative\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"depositNativeToL1\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedNative\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600080546001600160e01b03191673420000000000000000000000000000000000001017905534801561003657600080fd5b50604051610af0380380610af083398101604081905261005591610066565b6001600160a01b0316608052610096565b60006020828403121561007857600080fd5b81516001600160a01b038116811461008f57600080fd5b9392505050565b608051610a336100bd6000396000818160700152818160e901526101690152610a336000f3fe6080604052600436106100345760003560e01c80630ff98e311461003957806379a35b4b1461004e578063e861e90714610061575b600080fd5b61004c610047366004610861565b6100b8565b005b61004c61005c366004610883565b6100c5565b34801561006d57600080fd5b507f000000000000000000000000000000000000000000000000000000000000000060405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100c2813461038c565b50565b6100e773ffffffffffffffffffffffffffffffffffffffff84163330846104a2565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036101e9576040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632e1a7d4d90602401600060405180830381600087803b1580156101c257600080fd5b505af11580156101d6573d6000803e3d6000fd5b505050506101e4828261038c565b610386565b6000546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152602481018390529084169063095ea7b3906044016020604051808303816000875af1158015610262573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061028691906108ce565b506000805473ffffffffffffffffffffffffffffffffffffffff81169163a3a79548918691869186919074010000000000000000000000000000000000000000900467ffffffffffffffff168160146102de836108f0565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550604051602001610324919067ffffffffffffffff91909116815260200190565b6040516020818303038152906040526040518663ffffffff1660e01b81526004016103539594939291906109ac565b600060405180830381600087803b15801561036d57600080fd5b505af1158015610381573d6000803e3d6000fd5b505050505b50505050565b6000805473ffffffffffffffffffffffffffffffffffffffff81169163a3a795489173deaddeaddeaddeaddeaddeaddeaddeaddead000091869186919074010000000000000000000000000000000000000000900467ffffffffffffffff168160146103f7836108f0565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060405160200161043d919067ffffffffffffffff91909116815260200190565b6040516020818303038152906040526040518663ffffffff1660e01b815260040161046c9594939291906109ac565b600060405180830381600087803b15801561048657600080fd5b505af115801561049a573d6000803e3d6000fd5b505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff8581166024830152848116604483015260648083018590528351808403909101815260849092018352602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65649084015261038692879291600091610575918516908490610629565b805190915015610624578080602001905181019061059391906108ce565b610624576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b505050565b60606106388484600085610640565b949350505050565b6060824710156106d2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161061b565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516106fb91906109f7565b60006040518083038185875af1925050503d8060008114610738576040519150601f19603f3d011682016040523d82523d6000602084013e61073d565b606091505b509150915061074e87838387610759565b979650505050505050565b606083156107ef5782516000036107e85773ffffffffffffffffffffffffffffffffffffffff85163b6107e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161061b565b5081610638565b61063883838151156108045781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061b9190610a13565b803573ffffffffffffffffffffffffffffffffffffffff8116811461085c57600080fd5b919050565b60006020828403121561087357600080fd5b61087c82610838565b9392505050565b6000806000806080858703121561089957600080fd5b6108a285610838565b93506108b060208601610838565b92506108be60408601610838565b9396929550929360600135925050565b6000602082840312156108e057600080fd5b8151801515811461087c57600080fd5b600067ffffffffffffffff808316818103610934577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6001019392505050565b60005b83811015610959578181015183820152602001610941565b50506000910152565b6000815180845261097a81602086016020860161093e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600073ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525084604083015263ffffffff8416606083015260a0608083015261074e60a0830184610962565b60008251610a0981846020870161093e565b9190910192915050565b60208152600061087c602083018461096256fea164736f6c6343000813000a",
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
	GetWrappedNative(opts *bind.CallOpts) (common.Address, error)

	DepositNativeToL1(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, arg0 common.Address, l2Token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	Address() common.Address
}
