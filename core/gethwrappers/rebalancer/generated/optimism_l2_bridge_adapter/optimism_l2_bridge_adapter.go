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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIWrappedNative\",\"name\":\"wrappedNative\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MsgShouldNotContainValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"depositNativeToL1\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeFeeInNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedNative\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60c0604052734200000000000000000000000000000000000010608052600080546001600160401b031916905534801561003857600080fd5b50604051610c59380380610c5983398101604081905261005791610068565b6001600160a01b031660a052610098565b60006020828403121561007a57600080fd5b81516001600160a01b038116811461009157600080fd5b9392505050565b60805160a051610b816100d86000396000818160f00152818161018a015261020a0152600081816102c20152818161035a01526104650152610b816000f3fe60806040526004361061005a5760003560e01c806338314bb21161004357806338314bb21461009557806379a35b4b146100b6578063e861e907146100c957600080fd5b80630ff98e311461005f5780632e4b1fc914610074575b600080fd5b61007261006d366004610921565b61011a565b005b34801561008057600080fd5b50604051600081526020015b60405180910390f35b3480156100a157600080fd5b506100726100b0366004610943565b50505050565b6100726100c43660046109d1565b610127565b3480156100d557600080fd5b5060405173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815260200161008c565b610124813461044a565b50565b3415610166576040517f2543d86e0000000000000000000000000000000000000000000000000000000081523460048201526024015b60405180910390fd5b61018873ffffffffffffffffffffffffffffffffffffffff8416333084610567565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610285576040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632e1a7d4d90602401600060405180830381600087803b15801561026357600080fd5b505af1158015610277573d6000803e3d6000fd5b505050506100b0828261044a565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660048301526024820183905284169063095ea7b3906044016020604051808303816000875af115801561031a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061033e9190610a1c565b506000805473ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163a3a79548918691869186919067ffffffffffffffff16818061039d83610a3e565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040516020016103e3919067ffffffffffffffff91909116815260200190565b6040516020818303038152906040526040518663ffffffff1660e01b8152600401610412959493929190610afa565b600060405180830381600087803b15801561042c57600080fd5b505af1158015610440573d6000803e3d6000fd5b5050505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163a3a795489173deaddeaddeaddeaddeaddeaddeaddeaddead000091869186919067ffffffffffffffff1681806104bc83610a3e565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550604051602001610502919067ffffffffffffffff91909116815260200190565b6040516020818303038152906040526040518663ffffffff1660e01b8152600401610531959493929190610afa565b600060405180830381600087803b15801561054b57600080fd5b505af115801561055f573d6000803e3d6000fd5b505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff8581166024830152848116604483015260648083018590528351808403909101815260849092018352602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564908401526100b09287929160009161063a9185169084906106e9565b8051909150156106e457808060200190518101906106589190610a1c565b6106e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161015d565b505050565b60606106f88484600085610700565b949350505050565b606082471015610792576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161015d565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516107bb9190610b45565b60006040518083038185875af1925050503d80600081146107f8576040519150601f19603f3d011682016040523d82523d6000602084013e6107fd565b606091505b509150915061080e87838387610819565b979650505050505050565b606083156108af5782516000036108a85773ffffffffffffffffffffffffffffffffffffffff85163b6108a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161015d565b50816106f8565b6106f883838151156108c45781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015d9190610b61565b803573ffffffffffffffffffffffffffffffffffffffff8116811461091c57600080fd5b919050565b60006020828403121561093357600080fd5b61093c826108f8565b9392505050565b6000806000806060858703121561095957600080fd5b610962856108f8565b9350610970602086016108f8565b9250604085013567ffffffffffffffff8082111561098d57600080fd5b818701915087601f8301126109a157600080fd5b8135818111156109b057600080fd5b8860208285010111156109c257600080fd5b95989497505060200194505050565b600080600080608085870312156109e757600080fd5b6109f0856108f8565b93506109fe602086016108f8565b9250610a0c604086016108f8565b9396929550929360600135925050565b600060208284031215610a2e57600080fd5b8151801515811461093c57600080fd5b600067ffffffffffffffff808316818103610a82577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6001019392505050565b60005b83811015610aa7578181015183820152602001610a8f565b50506000910152565b60008151808452610ac8816020860160208601610a8c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600073ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525084604083015263ffffffff8416606083015260a0608083015261080e60a0830184610ab0565b60008251610b57818460208701610a8c565b9190910192915050565b60208152600061093c6020830184610ab056fea164736f6c6343000813000a",
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

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []byte) (*types.Transaction, error) {
	return _OptimismL2BridgeAdapter.contract.Transact(opts, "finalizeWithdrawERC20", arg0, arg1, arg2)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, arg2 []byte) (*types.Transaction, error) {
	return _OptimismL2BridgeAdapter.Contract.FinalizeWithdrawERC20(&_OptimismL2BridgeAdapter.TransactOpts, arg0, arg1, arg2)
}

func (_OptimismL2BridgeAdapter *OptimismL2BridgeAdapterTransactorSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, arg2 []byte) (*types.Transaction, error) {
	return _OptimismL2BridgeAdapter.Contract.FinalizeWithdrawERC20(&_OptimismL2BridgeAdapter.TransactOpts, arg0, arg1, arg2)
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

	FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []byte) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, arg0 common.Address, l2Token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	Address() common.Address
}
