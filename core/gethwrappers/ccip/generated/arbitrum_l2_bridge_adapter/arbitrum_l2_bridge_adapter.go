// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum_l2_bridge_adapter

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

var ArbitrumL2BridgeAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIL2GatewayRouter\",\"name\":\"l2GatewayRouter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"depositNativeToL1\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161092338038061092383398101604081905261002f91610067565b6001600160a01b03811661005657604051635e9c404d60e11b815260040160405180910390fd5b6001600160a01b0316608052610097565b60006020828403121561007957600080fd5b81516001600160a01b038116811461009057600080fd5b9392505050565b6080516108716100b2600039600061015901526108716000f3fe6080604052600436106100295760003560e01c80630ff98e311461002e57806379a35b4b14610043575b600080fd5b61004161003c3660046105e7565b610056565b005b610041610051366004610609565b6100ed565b6040517f25e1606300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526064906325e1606390349060240160206040518083038185885af11580156100c4573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906100e99190610654565b5050565b61010f73ffffffffffffffffffffffffffffffffffffffff84163330846101fe565b604080516020810182526000815290517f7b3a3c8b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691637b3a3c8b91610192918891879187916004016106db565b6000604051808303816000875af11580156101b1573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526101f79190810190610753565b5050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052610293908590610299565b50505050565b60006102fb826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166103af9092919063ffffffff16565b8051909150156103aa57808060200190518101906103199190610813565b6103aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b505050565b60606103be84846000856103c6565b949350505050565b606082471015610458576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016103a1565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516104819190610835565b60006040518083038185875af1925050503d80600081146104be576040519150601f19603f3d011682016040523d82523d6000602084013e6104c3565b606091505b50915091506104d4878383876104df565b979650505050505050565b6060831561057557825160000361056e5773ffffffffffffffffffffffffffffffffffffffff85163b61056e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103a1565b50816103be565b6103be838381511561058a5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a19190610851565b803573ffffffffffffffffffffffffffffffffffffffff811681146105e257600080fd5b919050565b6000602082840312156105f957600080fd5b610602826105be565b9392505050565b6000806000806080858703121561061f57600080fd5b610628856105be565b9350610636602086016105be565b9250610644604086016105be565b9396929550929360600135925050565b60006020828403121561066657600080fd5b5051919050565b60005b83811015610688578181015183820152602001610670565b50506000910152565b600081518084526106a981602086016020860161066d565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600073ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152508360408301526080606083015261071a6080830184610691565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561076557600080fd5b815167ffffffffffffffff8082111561077d57600080fd5b818401915084601f83011261079157600080fd5b8151818111156107a3576107a3610724565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156107e9576107e9610724565b8160405282815287602084870101111561080257600080fd5b6104d483602083016020880161066d565b60006020828403121561082557600080fd5b8151801515811461060257600080fd5b6000825161084781846020870161066d565b9190910192915050565b602081526000610602602083018461069156fea164736f6c6343000813000a",
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

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterTransactor) DepositNativeToL1(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.contract.Transact(opts, "depositNativeToL1", recipient)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterSession) DepositNativeToL1(recipient common.Address) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.Contract.DepositNativeToL1(&_ArbitrumL2BridgeAdapter.TransactOpts, recipient)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterTransactorSession) DepositNativeToL1(recipient common.Address) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.Contract.DepositNativeToL1(&_ArbitrumL2BridgeAdapter.TransactOpts, recipient)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterTransactor) SendERC20(opts *bind.TransactOpts, l1Token common.Address, l2Token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.contract.Transact(opts, "sendERC20", l1Token, l2Token, recipient, amount)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterSession) SendERC20(l1Token common.Address, l2Token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.Contract.SendERC20(&_ArbitrumL2BridgeAdapter.TransactOpts, l1Token, l2Token, recipient, amount)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapterTransactorSession) SendERC20(l1Token common.Address, l2Token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbitrumL2BridgeAdapter.Contract.SendERC20(&_ArbitrumL2BridgeAdapter.TransactOpts, l1Token, l2Token, recipient, amount)
}

func (_ArbitrumL2BridgeAdapter *ArbitrumL2BridgeAdapter) Address() common.Address {
	return _ArbitrumL2BridgeAdapter.address
}

type ArbitrumL2BridgeAdapterInterface interface {
	DepositNativeToL1(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, l1Token common.Address, l2Token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	Address() common.Address
}
