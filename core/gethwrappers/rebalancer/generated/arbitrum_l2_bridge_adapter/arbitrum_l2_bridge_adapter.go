// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum_l2_bridge_adapter

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

var ArbitrumL2BridgeAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIL2GatewayRouter\",\"name\":\"l2GatewayRouter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MsgShouldNotContainValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ArbitrumL1ToL2ERC20Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"outboundTransferResult\",\"type\":\"bytes\"}],\"name\":\"ArbitrumL2ToL1ERC20Sent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"depositNativeToL1\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeFeeInNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610b6e380380610b6e83398101604081905261002f91610067565b6001600160a01b03811661005657604051635e9c404d60e11b815260040160405180910390fd5b6001600160a01b0316608052610097565b60006020828403121561007957600080fd5b81516001600160a01b038116811461009057600080fd5b9392505050565b608051610abc6100b2600039600061022e0152610abc6000f3fe60806040526004361061003f5760003560e01c80630ff98e31146100445780632e4b1fc91461005957806338314bb21461007a578063a71d98b71461009a575b600080fd5b61005761005236600461073b565b6100ba565b005b34801561006557600080fd5b50604051600081526020015b60405180910390f35b34801561008657600080fd5b506100576100953660046107a6565b610151565b6100ad6100a8366004610807565b610180565b60405161007191906108f4565b6040517f25e1606300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526064906325e1606390349060240160206040518083038185885af1158015610128573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061014d9190610907565b5050565b6040517f23e6416f2029c421e34bdb63aa5810d66bff56ad91b79590968ab5c937e03d8990600090a150505050565b606034156101c1576040517f2543d86e0000000000000000000000000000000000000000000000000000000081523460048201526024015b60405180910390fd5b6101e373ffffffffffffffffffffffffffffffffffffffff8816333087610357565b60408051602081018252600080825291517f7b3a3c8b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691637b3a3c8b91610267918b918b918b91600401610920565b6000604051808303816000875af1158015610286573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526102cc9190810190610998565b90508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167fffa51f05034aaeaa4c0a623d14cae59dd4a579d2c66ed37f38c4baef2d504c018885604051610344929190610a58565b60405180910390a4979650505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526103ec9085906103f2565b50505050565b6000610454826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166105039092919063ffffffff16565b8051909150156104fe57808060200190518101906104729190610a71565b6104fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101b8565b505050565b6060610512848460008561051a565b949350505050565b6060824710156105ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101b8565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516105d59190610a93565b60006040518083038185875af1925050503d8060008114610612576040519150601f19603f3d011682016040523d82523d6000602084013e610617565b606091505b509150915061062887838387610633565b979650505050505050565b606083156106c95782516000036106c25773ffffffffffffffffffffffffffffffffffffffff85163b6106c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101b8565b5081610512565b61051283838151156106de5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b891906108f4565b803573ffffffffffffffffffffffffffffffffffffffff8116811461073657600080fd5b919050565b60006020828403121561074d57600080fd5b61075682610712565b9392505050565b60008083601f84011261076f57600080fd5b50813567ffffffffffffffff81111561078757600080fd5b60208301915083602082850101111561079f57600080fd5b9250929050565b600080600080606085870312156107bc57600080fd5b6107c585610712565b93506107d360208601610712565b9250604085013567ffffffffffffffff8111156107ef57600080fd5b6107fb8782880161075d565b95989497509550505050565b60008060008060008060a0878903121561082057600080fd5b61082987610712565b955061083760208801610712565b945061084560408801610712565b935060608701359250608087013567ffffffffffffffff81111561086857600080fd5b61087489828a0161075d565b979a9699509497509295939492505050565b60005b838110156108a1578181015183820152602001610889565b50506000910152565b600081518084526108c2816020860160208601610886565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061075660208301846108aa565b60006020828403121561091957600080fd5b5051919050565b600073ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152508360408301526080606083015261095f60808301846108aa565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156109aa57600080fd5b815167ffffffffffffffff808211156109c257600080fd5b818401915084601f8301126109d657600080fd5b8151818111156109e8576109e8610969565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610a2e57610a2e610969565b81604052828152876020848701011115610a4757600080fd5b610628836020830160208801610886565b82815260406020820152600061051260408301846108aa565b600060208284031215610a8357600080fd5b8151801515811461075657600080fd5b60008251610aa5818460208701610886565b919091019291505056fea164736f6c6343000813000a",
}

var ArbitrumL2BridgeAdapterABI = ArbitrumL2BridgeAdapterMetaData.ABI

var ArbitrumL2BridgeAdapterBin = ArbitrumL2BridgeAdapterMetaData.Bin

func DeployArbitrumL2BridgeAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, l2GatewayRouter common.Address) (common.Address, *types.Transaction, *ArbitrumL2BridgeAdapter, error) {
	parsed, err := ArbitrumL2BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArbitrumL2BridgeAdapterBin), backend, l2GatewayRouter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbitrumL2BridgeAdapter{address: address, abi: *parsed, ArbitrumL2BridgeAdapterCaller: ArbitrumL2BridgeAdapterCaller{contract: contract}, ArbitrumL2BridgeAdapterTransactor: ArbitrumL2BridgeAdapterTransactor{contract: contract}, ArbitrumL2BridgeAdapterFilterer: ArbitrumL2BridgeAdapterFilterer{contract: contract}}, nil
}

type ArbitrumL2BridgeAdapter struct {
	address common.Address
	abi     abi.ABI
	ArbitrumL2BridgeAdapterCaller
	ArbitrumL2BridgeAdapterTransactor
	ArbitrumL2BridgeAdapterFilterer
}

type ArbitrumL2BridgeAdapterCaller struct {
	contract *bind.BoundContract
}

type ArbitrumL2BridgeAdapterTransactor struct {
	contract *bind.BoundContract
}

type ArbitrumL2BridgeAdapterFilterer struct {
	contract *bind.BoundContract
}

type ArbitrumL2BridgeAdapterSession struct {
	Contract     *ArbitrumL2BridgeAdapter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ArbitrumL2BridgeAdapterCallerSession struct {
	Contract *ArbitrumL2BridgeAdapterCaller
	CallOpts bind.CallOpts
}

type ArbitrumL2BridgeAdapterTransactorSession struct {
	Contract     *ArbitrumL2BridgeAdapterTransactor
	TransactOpts bind.TransactOpts
}

type ArbitrumL2BridgeAdapterRaw struct {
	Contract *ArbitrumL2BridgeAdapter
}

type ArbitrumL2BridgeAdapterCallerRaw struct {
	Contract *ArbitrumL2BridgeAdapterCaller
}

type ArbitrumL2BridgeAdapterTransactorRaw struct {
	Contract *ArbitrumL2BridgeAdapterTransactor
}

func NewArbitrumL2BridgeAdapter(address common.Address, backend bind.ContractBackend) (*ArbitrumL2BridgeAdapter, error) {
	abi, err := abi.JSON(strings.NewReader(ArbitrumL2BridgeAdapterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindArbitrumL2BridgeAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL2BridgeAdapter{address: address, abi: abi, ArbitrumL2BridgeAdapterCaller: ArbitrumL2BridgeAdapterCaller{contract: contract}, ArbitrumL2BridgeAdapterTransactor: ArbitrumL2BridgeAdapterTransactor{contract: contract}, ArbitrumL2BridgeAdapterFilterer: ArbitrumL2BridgeAdapterFilterer{contract: contract}}, nil
}

func NewArbitrumL2BridgeAdapterCaller(address common.Address, caller bind.ContractCaller) (*ArbitrumL2BridgeAdapterCaller, error) {
	contract, err := bindArbitrumL2BridgeAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL2BridgeAdapterCaller{contract: contract}, nil
}

func NewArbitrumL2BridgeAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrumL2BridgeAdapterTransactor, error) {
	contract, err := bindArbitrumL2BridgeAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL2BridgeAdapterTransactor{contract: contract}, nil
}

func NewArbitrumL2BridgeAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrumL2BridgeAdapterFilterer, error) {
	contract, err := bindArbitrumL2BridgeAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL2BridgeAdapterFilterer{contract: contract}, nil
}

func bindArbitrumL2BridgeAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbitrumL2BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumL2BridgeAdapter.Contract.ArbitrumL2BridgeAdapterCaller.contract.Call(opts, result, method, params...)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.Contract.ArbitrumL2BridgeAdapterTransactor.contract.Transfer(opts)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.Contract.ArbitrumL2BridgeAdapterTransactor.contract.Transact(opts, method, params...)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumL2BridgeAdapter.Contract.contract.Call(opts, result, method, params...)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.Contract.contract.Transfer(opts)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.Contract.contract.Transact(opts, method, params...)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterCaller) GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumL2BridgeAdapter.contract.Call(opts, &out, "getBridgeFeeInNative")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _ArbitrumL2BridgeAdapter.Contract.GetBridgeFeeInNative(&_ArbitrumL2BridgeAdapter.CallOpts)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterCallerSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _ArbitrumL2BridgeAdapter.Contract.GetBridgeFeeInNative(&_ArbitrumL2BridgeAdapter.CallOpts)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterTransactor) DepositNativeToL1(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.contract.Transact(opts, "depositNativeToL1", recipient)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterSession) DepositNativeToL1(recipient common.Address) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.Contract.DepositNativeToL1(&_ArbitrumL2BridgeAdapter.TransactOpts, recipient)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterTransactorSession) DepositNativeToL1(recipient common.Address) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.Contract.DepositNativeToL1(&_ArbitrumL2BridgeAdapter.TransactOpts, recipient)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []byte) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.contract.Transact(opts, "finalizeWithdrawERC20", arg0, arg1, arg2)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, arg2 []byte) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.Contract.FinalizeWithdrawERC20(&_ArbitrumL2BridgeAdapter.TransactOpts, arg0, arg1, arg2)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterTransactorSession) FinalizeWithdrawERC20(arg0 common.Address, arg1 common.Address, arg2 []byte) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.Contract.FinalizeWithdrawERC20(&_ArbitrumL2BridgeAdapter.TransactOpts, arg0, arg1, arg2)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterTransactor) SendERC20(opts *bind.TransactOpts, localToken common.Address, remoteToken common.Address, recipient common.Address, amount *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.contract.Transact(opts, "sendERC20", localToken, remoteToken, recipient, amount, arg4)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterSession) SendERC20(localToken common.Address, remoteToken common.Address, recipient common.Address, amount *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.Contract.SendERC20(&_ArbitrumL2BridgeAdapter.TransactOpts, localToken, remoteToken, recipient, amount, arg4)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterTransactorSession) SendERC20(localToken common.Address, remoteToken common.Address, recipient common.Address, amount *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.Contract.SendERC20(&_ArbitrumL2BridgeAdapter.TransactOpts, localToken, remoteToken, recipient, amount, arg4)
}

type ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20FinalizedIterator struct {
	Event *ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20Finalized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20FinalizedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20Finalized)
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
		it.Event = new(ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20Finalized)
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

func (it *ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20FinalizedIterator) Error() error {
	return it.fail
}

func (it *ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20FinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20Finalized struct {
	Raw types.Log
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterFilterer) FilterArbitrumL1ToL2ERC20Finalized(opts *bind.FilterOpts) (*ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20FinalizedIterator, error) {

	logs, sub, err := _ArbitrumL2BridgeAdapter.contract.FilterLogs(opts, "ArbitrumL1ToL2ERC20Finalized")
	if err != nil {
		return nil, err
	}
	return &ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20FinalizedIterator{contract: _ArbitrumL2BridgeAdapter.contract, event: "ArbitrumL1ToL2ERC20Finalized", logs: logs, sub: sub}, nil
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterFilterer) WatchArbitrumL1ToL2ERC20Finalized(opts *bind.WatchOpts, sink chan<- *ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20Finalized) (event.Subscription, error) {

	logs, sub, err := _ArbitrumL2BridgeAdapter.contract.WatchLogs(opts, "ArbitrumL1ToL2ERC20Finalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20Finalized)
				if err := _ArbitrumL2BridgeAdapter.contract.UnpackLog(event, "ArbitrumL1ToL2ERC20Finalized", log); err != nil {
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

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterFilterer) ParseArbitrumL1ToL2ERC20Finalized(log types.Log) (*ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20Finalized, error) {
	event := new(ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20Finalized)
	if err := _ArbitrumL2BridgeAdapter.contract.UnpackLog(event, "ArbitrumL1ToL2ERC20Finalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20SentIterator struct {
	Event *ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20SentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent)
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
		it.Event = new(ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent)
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

func (it *ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20SentIterator) Error() error {
	return it.fail
}

func (it *ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20SentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent struct {
	LocalToken             common.Address
	RemoteToken            common.Address
	Recipient              common.Address
	Amount                 *big.Int
	OutboundTransferResult []byte
	Raw                    types.Log
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterFilterer) FilterArbitrumL2ToL1ERC20Sent(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, recipient []common.Address) (*ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20SentIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ArbitrumL2BridgeAdapter.contract.FilterLogs(opts, "ArbitrumL2ToL1ERC20Sent", localTokenRule, remoteTokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20SentIterator{contract: _ArbitrumL2BridgeAdapter.contract, event: "ArbitrumL2ToL1ERC20Sent", logs: logs, sub: sub}, nil
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterFilterer) WatchArbitrumL2ToL1ERC20Sent(opts *bind.WatchOpts, sink chan<- *ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent, localToken []common.Address, remoteToken []common.Address, recipient []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ArbitrumL2BridgeAdapter.contract.WatchLogs(opts, "ArbitrumL2ToL1ERC20Sent", localTokenRule, remoteTokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent)
				if err := _ArbitrumL2BridgeAdapter.contract.UnpackLog(event, "ArbitrumL2ToL1ERC20Sent", log); err != nil {
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

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterFilterer) ParseArbitrumL2ToL1ERC20Sent(log types.Log) (*ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent, error) {
	event := new(ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent)
	if err := _ArbitrumL2BridgeAdapter.contract.UnpackLog(event, "ArbitrumL2ToL1ERC20Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _ArbitrumL2BridgeAdapter.abi.Events["ArbitrumL1ToL2ERC20Finalized"].ID:
		return _ArbitrumL2BridgeAdapter.ParseArbitrumL1ToL2ERC20Finalized(log)
	case _ArbitrumL2BridgeAdapter.abi.Events["ArbitrumL2ToL1ERC20Sent"].ID:
		return _ArbitrumL2BridgeAdapter.ParseArbitrumL2ToL1ERC20Sent(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20Finalized) Topic() common.Hash {
	return common.HexToHash("0x23e6416f2029c421e34bdb63aa5810d66bff56ad91b79590968ab5c937e03d89")
}

func (ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent) Topic() common.Hash {
	return common.HexToHash("0xffa51f05034aaeaa4c0a623d14cae59dd4a579d2c66ed37f38c4baef2d504c01")
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapter) Address() common.Address {
	return _ArbitrumL2BridgeAdapter.address
}

type ArbitrumL2BridgeAdapterInterface interface {
	GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error)

	DepositNativeToL1(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

	FinalizeWithdrawERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []byte) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, localToken common.Address, remoteToken common.Address, recipient common.Address, amount *big.Int, arg4 []byte) (*types.Transaction, error)

	FilterArbitrumL1ToL2ERC20Finalized(opts *bind.FilterOpts) (*ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20FinalizedIterator, error)

	WatchArbitrumL1ToL2ERC20Finalized(opts *bind.WatchOpts, sink chan<- *ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20Finalized) (event.Subscription, error)

	ParseArbitrumL1ToL2ERC20Finalized(log types.Log) (*ArbitrumL2BridgeAdapterArbitrumL1ToL2ERC20Finalized, error)

	FilterArbitrumL2ToL1ERC20Sent(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, recipient []common.Address) (*ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20SentIterator, error)

	WatchArbitrumL2ToL1ERC20Sent(opts *bind.WatchOpts, sink chan<- *ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent, localToken []common.Address, remoteToken []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseArbitrumL2ToL1ERC20Sent(log types.Log) (*ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
