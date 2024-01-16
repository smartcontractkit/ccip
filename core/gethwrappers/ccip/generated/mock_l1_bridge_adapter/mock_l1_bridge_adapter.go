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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MsgShouldNotContainValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1Receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"bridgeSpecificPayload\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20FromL2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeFeeInNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"provideLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610c41380380610c4183398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051610ba16100a0600039600081816101560152818161020d015281816102df01526103f40152610ba16000f3fe6080604052600436106100655760003560e01c806379a35b4b1161004357806379a35b4b146100cd578063cb7a6e16146100e0578063eb521a4c1461010157600080fd5b80630184f8461461006a5780630a861f2a1461008c5780632e4b1fc9146100ac575b600080fd5b34801561007657600080fd5b5061008a6100853660046108cd565b610121565b005b34801561009857600080fd5b5061008a6100a73660046109a6565b6101dd565b3480156100b857600080fd5b50600060405190815260200160405180910390f35b61008a6100db3660046109e8565b610336565b3480156100ec57600080fd5b5061008a6100fb366004610a33565b50505050565b34801561010d57600080fd5b5061008a61011c3660046109a6565b6103da565b6040517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018390527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a9059cbb906044016020604051808303816000875af11580156101b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101d89190610ac1565b505050565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015281907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610269573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061028d9190610aea565b10156102c5576040517fbb55fd2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61030673ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016338361044c565b604051819033907fc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf984017171990600090a350565b6040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810182905273ffffffffffffffffffffffffffffffffffffffff8516906323b872dd906064016020604051808303816000875af11580156103af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d39190610ac1565b5050505050565b61041c73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084610520565b604051819033907fc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb31208890600090a350565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526101d89084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261057e565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526100fb9085907f23b872dd000000000000000000000000000000000000000000000000000000009060840161049e565b60006105e0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661068f9092919063ffffffff16565b8051909150156101d857808060200190518101906105fe9190610ac1565b6101d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b606061069e84846000856106a6565b949350505050565b606082471015610738576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610686565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516107619190610b27565b60006040518083038185875af1925050503d806000811461079e576040519150601f19603f3d011682016040523d82523d6000602084013e6107a3565b606091505b50915091506107b4878383876107bf565b979650505050505050565b6060831561085557825160000361084e5773ffffffffffffffffffffffffffffffffffffffff85163b61084e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610686565b508161069e565b61069e838381511561086a5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106869190610b43565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080604083850312156108e057600080fd5b82359150602083013567ffffffffffffffff808211156108ff57600080fd5b818501915085601f83011261091357600080fd5b8135818111156109255761092561089e565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561096b5761096b61089e565b8160405282815288602084870101111561098457600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b6000602082840312156109b857600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146109e357600080fd5b919050565b600080600080608085870312156109fe57600080fd5b610a07856109bf565b9350610a15602086016109bf565b9250610a23604086016109bf565b9396929550929360600135925050565b60008060008060608587031215610a4957600080fd5b610a52856109bf565b9350610a60602086016109bf565b9250604085013567ffffffffffffffff80821115610a7d57600080fd5b818701915087601f830112610a9157600080fd5b813581811115610aa057600080fd5b886020828501011115610ab257600080fd5b95989497505060200194505050565b600060208284031215610ad357600080fd5b81518015158114610ae357600080fd5b9392505050565b600060208284031215610afc57600080fd5b5051919050565b60005b83811015610b1e578181015183820152602001610b06565b50506000910152565b60008251610b39818460208701610b03565b9190910192915050565b6020815260008251806020840152610b62816040850160208701610b03565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fea164736f6c6343000813000a",
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

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactor) FinalizeWithdraw(opts *bind.TransactOpts, amount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.contract.Transact(opts, "finalizeWithdraw", amount, arg1)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) FinalizeWithdraw(amount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.FinalizeWithdraw(&_MockL1BridgeAdapter.TransactOpts, amount, arg1)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorSession) FinalizeWithdraw(amount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.FinalizeWithdraw(&_MockL1BridgeAdapter.TransactOpts, amount, arg1)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactor) FinalizeWithdrawERC20FromL2(opts *bind.TransactOpts, l2Sender common.Address, l1Receiver common.Address, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.contract.Transact(opts, "finalizeWithdrawERC20FromL2", l2Sender, l1Receiver, bridgeSpecificPayload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) FinalizeWithdrawERC20FromL2(l2Sender common.Address, l1Receiver common.Address, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.FinalizeWithdrawERC20FromL2(&_MockL1BridgeAdapter.TransactOpts, l2Sender, l1Receiver, bridgeSpecificPayload)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorSession) FinalizeWithdrawERC20FromL2(l2Sender common.Address, l1Receiver common.Address, bridgeSpecificPayload []byte) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.FinalizeWithdrawERC20FromL2(&_MockL1BridgeAdapter.TransactOpts, l2Sender, l1Receiver, bridgeSpecificPayload)
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

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactor) SendERC20(opts *bind.TransactOpts, l1Token common.Address, arg1 common.Address, arg2 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.contract.Transact(opts, "sendERC20", l1Token, arg1, arg2, amount)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterSession) SendERC20(l1Token common.Address, arg1 common.Address, arg2 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.SendERC20(&_MockL1BridgeAdapter.TransactOpts, l1Token, arg1, arg2, amount)
}

func (_MockL1BridgeAdapter *MockL1BridgeAdapterTransactorSession) SendERC20(l1Token common.Address, arg1 common.Address, arg2 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockL1BridgeAdapter.Contract.SendERC20(&_MockL1BridgeAdapter.TransactOpts, l1Token, arg1, arg2, amount)
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

	FinalizeWithdraw(opts *bind.TransactOpts, amount *big.Int, arg1 []byte) (*types.Transaction, error)

	FinalizeWithdrawERC20FromL2(opts *bind.TransactOpts, l2Sender common.Address, l1Receiver common.Address, bridgeSpecificPayload []byte) (*types.Transaction, error)

	ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, l1Token common.Address, arg1 common.Address, arg2 common.Address, amount *big.Int) (*types.Transaction, error)

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
