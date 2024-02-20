// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_l1_bridge_adapter

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

var MockL1BridgeAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MsgShouldNotContainValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"remoteSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"localReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"bridgeSpecificPayload\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeFeeInNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"provideLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"bridgeSpecificPayload\",\"type\":\"bytes\"}],\"name\":\"sendERC20\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600160005534801561001557600080fd5b50604051610c00380380610c0083398101604081905261003491610045565b6001600160a01b0316608052610075565b60006020828403121561005757600080fd5b81516001600160a01b038116811461006e57600080fd5b9392505050565b608051610b5b6100a560003960008181610132015281816102040152818161028601526103ce0152610b5b6000f3fe60806040526004361061005a5760003560e01c806338314bb21161004357806338314bb2146100a2578063a71d98b7146100c2578063eb521a4c146100e257600080fd5b80630a861f2a1461005f5780632e4b1fc914610081575b600080fd5b34801561006b57600080fd5b5061007f61007a366004610883565b610102565b005b34801561008d57600080fd5b50604051600081526020015b60405180910390f35b3480156100ae57600080fd5b5061007f6100bd36600461090e565b61025b565b6100d56100d036600461096f565b6102b4565b6040516100999190610a5c565b3480156100ee57600080fd5b5061007f6100fd366004610883565b6103b4565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015281907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa15801561018e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101b29190610a76565b10156101ea576040517fbb55fd2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61022b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163383610426565b604051819033907fc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf984017171990600090a350565b600061026982840184610a8f565b5090506102ad73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168583610426565b5050505050565b6040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810184905260609073ffffffffffffffffffffffffffffffffffffffff8816906323b872dd906064016020604051808303816000875af1158015610330573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103549190610ab1565b5060008054818061036483610ad3565b9190505560405160200161037a91815260200190565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815291905298975050505050505050565b6103f673ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163330846104ff565b604051819033907fc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb31208890600090a350565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526104fa9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610563565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff8085166024830152831660448201526064810182905261055d9085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401610478565b50505050565b60006105c5826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166106749092919063ffffffff16565b8051909150156104fa57808060200190518101906105e39190610ab1565b6104fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6060610683848460008561068b565b949350505050565b60608247101561071d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161066b565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516107469190610b32565b60006040518083038185875af1925050503d8060008114610783576040519150601f19603f3d011682016040523d82523d6000602084013e610788565b606091505b5091509150610799878383876107a4565b979650505050505050565b6060831561083a5782516000036108335773ffffffffffffffffffffffffffffffffffffffff85163b610833576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161066b565b5081610683565b610683838381511561084f5781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066b9190610a5c565b60006020828403121561089557600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146108c057600080fd5b919050565b60008083601f8401126108d757600080fd5b50813567ffffffffffffffff8111156108ef57600080fd5b60208301915083602082850101111561090757600080fd5b9250929050565b6000806000806060858703121561092457600080fd5b61092d8561089c565b935061093b6020860161089c565b9250604085013567ffffffffffffffff81111561095757600080fd5b610963878288016108c5565b95989497509550505050565b60008060008060008060a0878903121561098857600080fd5b6109918761089c565b955061099f6020880161089c565b94506109ad6040880161089c565b935060608701359250608087013567ffffffffffffffff8111156109d057600080fd5b6109dc89828a016108c5565b979a9699509497509295939492505050565b60005b83811015610a095781810151838201526020016109f1565b50506000910152565b60008151808452610a2a8160208601602086016109ee565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610a6f6020830184610a12565b9392505050565b600060208284031215610a8857600080fd5b5051919050565b60008060408385031215610aa257600080fd5b50508035926020909101359150565b600060208284031215610ac357600080fd5b81518015158114610a6f57600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610b2b577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b60008251610b448184602087016109ee565b919091019291505056fea164736f6c6343000813000a",
}

var MockL1BridgeAdapterABI = MockL1BridgeAdapterMetaData.ABI

var MockL1BridgeAdapterBin = MockL1BridgeAdapterMetaData.Bin

func DeployMockL1BridgeAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *MockL1BridgeAdapter, error) {
	parsed, err := MockL1BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockL1BridgeAdapterBin), backend, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockL1BridgeAdapter{address: address, abi: *parsed, MockL1BridgeAdapterCaller: MockL1BridgeAdapterCaller{contract: contract}, MockL1BridgeAdapterTransactor: MockL1BridgeAdapterTransactor{contract: contract}, MockL1BridgeAdapterFilterer: MockL1BridgeAdapterFilterer{contract: contract}}, nil
}

type MockL1BridgeAdapter struct {
	address common.Address
	abi     abi.ABI
	MockL1BridgeAdapterCaller
	MockL1BridgeAdapterTransactor
	MockL1BridgeAdapterFilterer
}

type MockL1BridgeAdapterCaller struct {
	contract *bind.BoundContract
}

type MockL1BridgeAdapterTransactor struct {
	contract *bind.BoundContract
}

type MockL1BridgeAdapterFilterer struct {
	contract *bind.BoundContract
}

type MockL1BridgeAdapterSession struct {
	Contract     *MockL1BridgeAdapter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockL1BridgeAdapterCallerSession struct {
	Contract *MockL1BridgeAdapterCaller
	CallOpts bind.CallOpts
}

type MockL1BridgeAdapterTransactorSession struct {
	Contract     *MockL1BridgeAdapterTransactor
	TransactOpts bind.TransactOpts
}

type MockL1BridgeAdapterRaw struct {
	Contract *MockL1BridgeAdapter
}

type MockL1BridgeAdapterCallerRaw struct {
	Contract *MockL1BridgeAdapterCaller
}

type MockL1BridgeAdapterTransactorRaw struct {
	Contract *MockL1BridgeAdapterTransactor
}

func NewMockL1BridgeAdapter(address common.Address, backend bind.ContractBackend) (*MockL1BridgeAdapter, error) {
	abi, err := abi.JSON(strings.NewReader(MockL1BridgeAdapterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockL1BridgeAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockL1BridgeAdapter{address: address, abi: abi, MockL1BridgeAdapterCaller: MockL1BridgeAdapterCaller{contract: contract}, MockL1BridgeAdapterTransactor: MockL1BridgeAdapterTransactor{contract: contract}, MockL1BridgeAdapterFilterer: MockL1BridgeAdapterFilterer{contract: contract}}, nil
}

func NewMockL1BridgeAdapterCaller(address common.Address, caller bind.ContractCaller) (*MockL1BridgeAdapterCaller, error) {
	contract, err := bindMockL1BridgeAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockL1BridgeAdapterCaller{contract: contract}, nil
}

func NewMockL1BridgeAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*MockL1BridgeAdapterTransactor, error) {
	contract, err := bindMockL1BridgeAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockL1BridgeAdapterTransactor{contract: contract}, nil
}

func NewMockL1BridgeAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*MockL1BridgeAdapterFilterer, error) {
	contract, err := bindMockL1BridgeAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockL1BridgeAdapterFilterer{contract: contract}, nil
}

func bindMockL1BridgeAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockL1BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL1BridgeAdapter.Contract.MockL1BridgeAdapterCaller.contract.Call(opts, result, method, params...)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.MockL1BridgeAdapterTransactor.contract.Transfer(opts)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.MockL1BridgeAdapterTransactor.contract.Transact(opts, method, params...)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL1BridgeAdapter.Contract.contract.Call(opts, result, method, params...)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.contract.Transfer(opts)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.contract.Transact(opts, method, params...)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterCaller) GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockL1BridgeAdapter.contract.Call(opts, &out, "getBridgeFeeInNative")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _MockL1BridgeAdapter.Contract.GetBridgeFeeInNative(&_MockL1BridgeAdapter.CallOpts)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterCallerSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _MockL1BridgeAdapter.Contract.GetBridgeFeeInNative(&_MockL1BridgeAdapter.CallOpts)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, remoteSender common.Address, localReceiver common.Address, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.contract.Transact(opts, "finalizeWithdrawERC20", remoteSender, localReceiver, bridgeSpecificPayload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) FinalizeWithdrawERC20(remoteSender common.Address, localReceiver common.Address, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.FinalizeWithdrawERC20(&_MockL1BridgeAdapter.TransactOpts, remoteSender, localReceiver, bridgeSpecificPayload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorSession) FinalizeWithdrawERC20(remoteSender common.Address, localReceiver common.Address, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.FinalizeWithdrawERC20(&_MockL1BridgeAdapter.TransactOpts, remoteSender, localReceiver, bridgeSpecificPayload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactor) ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.contract.Transact(opts, "provideLiquidity", amount)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.ProvideLiquidity(&_MockL1BridgeAdapter.TransactOpts, amount)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.ProvideLiquidity(&_MockL1BridgeAdapter.TransactOpts, amount)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactor) SendERC20(opts *bind.TransactOpts, localToken common.Address, remoteToken common.Address, remoteReceiver common.Address, amount *big.Int, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.contract.Transact(opts, "sendERC20", localToken, remoteToken, remoteReceiver, amount, bridgeSpecificPayload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) SendERC20(localToken common.Address, remoteToken common.Address, remoteReceiver common.Address, amount *big.Int, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.SendERC20(&_MockL1BridgeAdapter.TransactOpts, localToken, remoteToken, remoteReceiver, amount, bridgeSpecificPayload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorSession) SendERC20(localToken common.Address, remoteToken common.Address, remoteReceiver common.Address, amount *big.Int, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.SendERC20(&_MockL1BridgeAdapter.TransactOpts, localToken, remoteToken, remoteReceiver, amount, bridgeSpecificPayload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactor) WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.contract.Transact(opts, "withdrawLiquidity", amount)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.WithdrawLiquidity(&_MockL1BridgeAdapter.TransactOpts, amount)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.WithdrawLiquidity(&_MockL1BridgeAdapter.TransactOpts, amount)
}

type MockL1BridgeAdapterLiquidityAddedIterator struct {
	Event *MockL1BridgeAdapterLiquidityAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockL1BridgeAdapterLiquidityAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL1BridgeAdapterLiquidityAdded)
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
		it.Event = new(MockL1BridgeAdapterLiquidityAdded)
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

func (it *MockL1BridgeAdapterLiquidityAddedIterator) Error() error {
	return it.fail
}

func (it *MockL1BridgeAdapterLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockL1BridgeAdapterLiquidityAdded struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*MockL1BridgeAdapterLiquidityAddedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _MockL1BridgeAdapter.contract.FilterLogs(opts, "LiquidityAdded", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &MockL1BridgeAdapterLiquidityAddedIterator{contract: _MockL1BridgeAdapter.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *MockL1BridgeAdapterLiquidityAdded, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _MockL1BridgeAdapter.contract.WatchLogs(opts, "LiquidityAdded", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockL1BridgeAdapterLiquidityAdded)
				if err := _MockL1BridgeAdapter.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

func (_MockL1BridgeAdapter *MockL1BridgeAdapterFilterer) ParseLiquidityAdded(log types.Log) (*MockL1BridgeAdapterLiquidityAdded, error) {
	event := new(MockL1BridgeAdapterLiquidityAdded)
	if err := _MockL1BridgeAdapter.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockL1BridgeAdapterLiquidityRemovedIterator struct {
	Event *MockL1BridgeAdapterLiquidityRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockL1BridgeAdapterLiquidityRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL1BridgeAdapterLiquidityRemoved)
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
		it.Event = new(MockL1BridgeAdapterLiquidityRemoved)
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

func (it *MockL1BridgeAdapterLiquidityRemovedIterator) Error() error {
	return it.fail
}

func (it *MockL1BridgeAdapterLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockL1BridgeAdapterLiquidityRemoved struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*MockL1BridgeAdapterLiquidityRemovedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _MockL1BridgeAdapter.contract.FilterLogs(opts, "LiquidityRemoved", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &MockL1BridgeAdapterLiquidityRemovedIterator{contract: _MockL1BridgeAdapter.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *MockL1BridgeAdapterLiquidityRemoved, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _MockL1BridgeAdapter.contract.WatchLogs(opts, "LiquidityRemoved", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockL1BridgeAdapterLiquidityRemoved)
				if err := _MockL1BridgeAdapter.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

func (_MockL1BridgeAdapter *MockL1BridgeAdapterFilterer) ParseLiquidityRemoved(log types.Log) (*MockL1BridgeAdapterLiquidityRemoved, error) {
	event := new(MockL1BridgeAdapterLiquidityRemoved)
	if err := _MockL1BridgeAdapter.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MockL1BridgeAdapter.abi.Events["LiquidityAdded"].ID:
		return _MockL1BridgeAdapter.ParseLiquidityAdded(log)
	case _MockL1BridgeAdapter.abi.Events["LiquidityRemoved"].ID:
		return _MockL1BridgeAdapter.ParseLiquidityRemoved(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MockL1BridgeAdapterLiquidityAdded) Topic() common.Hash {
	return common.HexToHash("0xc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088")
}

func (MockL1BridgeAdapterLiquidityRemoved) Topic() common.Hash {
	return common.HexToHash("0xc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719")
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapter) Address() common.Address {
	return _MockL1BridgeAdapter.address
}

type MockL1BridgeAdapterInterface interface {
	GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error)

	FinalizeWithdrawERC20(opts *bind.TransactOpts, remoteSender common.Address, localReceiver common.Address, bridgeSpecificPayload []byte) (*types.Transaction, error)

	ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, localToken common.Address, remoteToken common.Address, remoteReceiver common.Address, amount *big.Int, bridgeSpecificPayload []byte) (*types.Transaction, error)

	WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*MockL1BridgeAdapterLiquidityAddedIterator, error)

	WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *MockL1BridgeAdapterLiquidityAdded, provider []common.Address, amount []*big.Int) (event.Subscription, error)

	ParseLiquidityAdded(log types.Log) (*MockL1BridgeAdapterLiquidityAdded, error)

	FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*MockL1BridgeAdapterLiquidityRemovedIterator, error)

	WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *MockL1BridgeAdapterLiquidityRemoved, provider []common.Address, amount []*big.Int) (event.Subscription, error)

	ParseLiquidityRemoved(log types.Log) (*MockL1BridgeAdapterLiquidityRemoved, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
