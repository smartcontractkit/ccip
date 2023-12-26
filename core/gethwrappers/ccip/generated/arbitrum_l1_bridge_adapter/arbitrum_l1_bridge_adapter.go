// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum_l1_bridge_adapter

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

var ArbitrumL1BridgeAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIL1GatewayRouter\",\"name\":\"l1GatewayRouter\",\"type\":\"address\"},{\"internalType\":\"contractIOutbox\",\"name\":\"l1Outbox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1ERC20Gateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GAS_PRICE_BID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SUBMISSION_COST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1Receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arbitrumFinalizationPayload\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20FromL2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"}],\"name\":\"getL2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60e0604052600080546001600160401b031916905534801561002057600080fd5b50604051620010b6380380620010b6833981016040819052610041916100c3565b6001600160a01b038316158061005e57506001600160a01b038216155b8061007057506001600160a01b038116155b1561008e57604051635e9c404d60e11b815260040160405180910390fd5b6001600160a01b0392831660805290821660c0521660a052610110565b6001600160a01b03811681146100c057600080fd5b50565b6000806000606084860312156100d857600080fd5b83516100e3816100ab565b60208501519093506100f4816100ab565b6040850151909250610105816100ab565b809150509250925092565b60805160a05160c051610f6f620001476000396000610469015260006101a001526000818161028d01526103e40152610f6f6000f3fe6080604052600436106100655760003560e01c8063c985069c11610043578063c985069c146100c1578063cb7a6e1614610106578063ea6c2f801461012657600080fd5b806332eb79051461006a5780635f2a9f411461009557806379a35b4b146100ac575b600080fd5b34801561007657600080fd5b506100826311e1a30081565b6040519081526020015b60405180910390f35b3480156100a157600080fd5b50610082620186a081565b6100bf6100ba366004610905565b610141565b005b3480156100cd57600080fd5b506100e16100dc366004610956565b61039c565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161008c565b34801561011257600080fd5b506100bf61012136600461097a565b610457565b34801561013257600080fd5b506100826602d79883d2000081565b61016373ffffffffffffffffffffffffffffffffffffffff8516333084610525565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660048301526024820183905285169063095ea7b3906044016020604051808303816000875af11580156101f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061021c9190610a0c565b5060006102306311e1a300620186a0610a5d565b610241906602d79883d20000610a74565b90508034101561028b576040517fe2c5a8f7000000000000000000000000000000000000000000000000000000008152600481018290523460248201526044015b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16634fb1a07b8287868787620186a06311e1a3006602d79883d20000604051806020016040528060008152506040516020016102fe929190610af5565b6040516020818303038152906040526040518963ffffffff1660e01b815260040161032f9796959493929190610b0e565b60006040518083038185885af115801561034d573d6000803e3d6000fd5b50505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526103949190810190610c5b565b505050505050565b6040517fa7e28d4800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063a7e28d4890602401602060405180830381865afa15801561042d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104519190610cd7565b92915050565b600061046582840184610dba565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166308635a958260000151836020015188888660400151876060015188608001518960a001518a60c001516040518a63ffffffff1660e01b81526004016104ec99989796959493929190610e7d565b600060405180830381600087803b15801561050657600080fd5b505af115801561051a573d6000803e3d6000fd5b505050505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526105ba9085906105c0565b50505050565b6000610622826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166106d19092919063ffffffff16565b8051909150156106cc57808060200190518101906106409190610a0c565b6106cc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610282565b505050565b60606106e084846000856106e8565b949350505050565b60608247101561077a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610282565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516107a39190610f33565b60006040518083038185875af1925050503d80600081146107e0576040519150601f19603f3d011682016040523d82523d6000602084013e6107e5565b606091505b50915091506107f687838387610801565b979650505050505050565b606083156108975782516000036108905773ffffffffffffffffffffffffffffffffffffffff85163b610890576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610282565b50816106e0565b6106e083838151156108ac5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102829190610f4f565b73ffffffffffffffffffffffffffffffffffffffff8116811461090257600080fd5b50565b6000806000806080858703121561091b57600080fd5b8435610926816108e0565b93506020850135610936816108e0565b92506040850135610946816108e0565b9396929550929360600135925050565b60006020828403121561096857600080fd5b8135610973816108e0565b9392505050565b6000806000806060858703121561099057600080fd5b843561099b816108e0565b935060208501356109ab816108e0565b9250604085013567ffffffffffffffff808211156109c857600080fd5b818701915087601f8301126109dc57600080fd5b8135818111156109eb57600080fd5b8860208285010111156109fd57600080fd5b95989497505060200194505050565b600060208284031215610a1e57600080fd5b8151801515811461097357600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808202811582820484141761045157610451610a2e565b8082018082111561045157610451610a2e565b60005b83811015610aa2578181015183820152602001610a8a565b50506000910152565b60008151808452610ac3816020860160208601610a87565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b8281526040602082015260006106e06040830184610aab565b600073ffffffffffffffffffffffffffffffffffffffff808a16835280891660208401528088166040840152508560608301528460808301528360a083015260e060c0830152610b6160e0830184610aab565b9998505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715610bc057610bc0610b6e565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610c0d57610c0d610b6e565b604052919050565b600067ffffffffffffffff821115610c2f57610c2f610b6e565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600060208284031215610c6d57600080fd5b815167ffffffffffffffff811115610c8457600080fd5b8201601f81018413610c9557600080fd5b8051610ca8610ca382610c15565b610bc6565b818152856020838501011115610cbd57600080fd5b610cce826020830160208601610a87565b95945050505050565b600060208284031215610ce957600080fd5b8151610973816108e0565b600082601f830112610d0557600080fd5b8135602067ffffffffffffffff821115610d2157610d21610b6e565b8160051b610d30828201610bc6565b9283528481018201928281019087851115610d4a57600080fd5b83870192505b848310156107f657823582529183019190830190610d50565b600082601f830112610d7a57600080fd5b8135610d88610ca382610c15565b818152846020838601011115610d9d57600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215610dcc57600080fd5b813567ffffffffffffffff80821115610de457600080fd5b9083019060e08286031215610df857600080fd5b610e00610b9d565b823582811115610e0f57600080fd5b610e1b87828601610cf4565b8252506020830135602082015260408301356040820152606083013560608201526080830135608082015260a083013560a082015260c083013582811115610e6257600080fd5b610e6e87828601610d69565b60c08301525095945050505050565b6101208082528a51908201819052600090610140830190602090818e01845b82811015610eb857815185529383019390830190600101610e9c565b50505083018b905273ffffffffffffffffffffffffffffffffffffffff8a16604084015273ffffffffffffffffffffffffffffffffffffffff891660608401528760808401528660a08401528560c08401528460e0840152828103610100840152610f238185610aab565b9c9b505050505050505050505050565b60008251610f45818460208701610a87565b9190910192915050565b6020815260006109736020830184610aab56fea164736f6c6343000813000a",
}

var ArbitrumL1BridgeAdapterABI = ArbitrumL1BridgeAdapterMetaData.ABI

var ArbitrumL1BridgeAdapterBin = ArbitrumL1BridgeAdapterMetaData.Bin

func DeployArbitrumL1BridgeAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, l1GatewayRouter common.Address, l1Outbox common.Address, l1ERC20Gateway common.Address) (common.Address, *types.Transaction, *ArbitrumL1BridgeAdapter, error) {
	parsed, err := ArbitrumL1BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArbitrumL1BridgeAdapterBin), backend, l1GatewayRouter, l1Outbox, l1ERC20Gateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbitrumL1BridgeAdapter{address: address, abi: *parsed, ArbitrumL1BridgeAdapterCaller: ArbitrumL1BridgeAdapterCaller{contract: contract}, ArbitrumL1BridgeAdapterTransactor: ArbitrumL1BridgeAdapterTransactor{contract: contract}, ArbitrumL1BridgeAdapterFilterer: ArbitrumL1BridgeAdapterFilterer{contract: contract}}, nil
}

type ArbitrumL1BridgeAdapter struct {
	address common.Address
	abi     abi.ABI
	ArbitrumL1BridgeAdapterCaller
	ArbitrumL1BridgeAdapterTransactor
	ArbitrumL1BridgeAdapterFilterer
}

type ArbitrumL1BridgeAdapterCaller struct {
	contract *bind.BoundContract
}

type ArbitrumL1BridgeAdapterTransactor struct {
	contract *bind.BoundContract
}

type ArbitrumL1BridgeAdapterFilterer struct {
	contract *bind.BoundContract
}

type ArbitrumL1BridgeAdapterSession struct {
	Contract     *ArbitrumL1BridgeAdapter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ArbitrumL1BridgeAdapterCallerSession struct {
	Contract *ArbitrumL1BridgeAdapterCaller
	CallOpts bind.CallOpts
}

type ArbitrumL1BridgeAdapterTransactorSession struct {
	Contract     *ArbitrumL1BridgeAdapterTransactor
	TransactOpts bind.TransactOpts
}

type ArbitrumL1BridgeAdapterRaw struct {
	Contract *ArbitrumL1BridgeAdapter
}

type ArbitrumL1BridgeAdapterCallerRaw struct {
	Contract *ArbitrumL1BridgeAdapterCaller
}

type ArbitrumL1BridgeAdapterTransactorRaw struct {
	Contract *ArbitrumL1BridgeAdapterTransactor
}

func NewArbitrumL1BridgeAdapter(address common.Address, backend bind.ContractBackend) (*ArbitrumL1BridgeAdapter, error) {
	abi, err := abi.JSON(strings.NewReader(ArbitrumL1BridgeAdapterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindArbitrumL1BridgeAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapter{address: address, abi: abi, ArbitrumL1BridgeAdapterCaller: ArbitrumL1BridgeAdapterCaller{contract: contract}, ArbitrumL1BridgeAdapterTransactor: ArbitrumL1BridgeAdapterTransactor{contract: contract}, ArbitrumL1BridgeAdapterFilterer: ArbitrumL1BridgeAdapterFilterer{contract: contract}}, nil
}

func NewArbitrumL1BridgeAdapterCaller(address common.Address, caller bind.ContractCaller) (*ArbitrumL1BridgeAdapterCaller, error) {
	contract, err := bindArbitrumL1BridgeAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapterCaller{contract: contract}, nil
}

func NewArbitrumL1BridgeAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrumL1BridgeAdapterTransactor, error) {
	contract, err := bindArbitrumL1BridgeAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapterTransactor{contract: contract}, nil
}

func NewArbitrumL1BridgeAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrumL1BridgeAdapterFilterer, error) {
	contract, err := bindArbitrumL1BridgeAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrumL1BridgeAdapterFilterer{contract: contract}, nil
}

func bindArbitrumL1BridgeAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbitrumL1BridgeAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumL1BridgeAdapter.Contract.ArbitrumL1BridgeAdapterCaller.contract.Call(opts, result, method, params...)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.ArbitrumL1BridgeAdapterTransactor.contract.Transfer(opts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.ArbitrumL1BridgeAdapterTransactor.contract.Transact(opts, method, params...)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumL1BridgeAdapter.Contract.contract.Call(opts, result, method, params...)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.contract.Transfer(opts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.contract.Transact(opts, method, params...)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) GASPRICEBID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "GAS_PRICE_BID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) GASPRICEBID() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GASPRICEBID(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) GASPRICEBID() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GASPRICEBID(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) MAXGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "MAX_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) MAXGAS() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.MAXGAS(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) MAXGAS() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.MAXGAS(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) MAXSUBMISSIONCOST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "MAX_SUBMISSION_COST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) MAXSUBMISSIONCOST() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.MAXSUBMISSIONCOST(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) MAXSUBMISSIONCOST() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.MAXSUBMISSIONCOST(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) GetL2Token(opts *bind.CallOpts, l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "getL2Token", l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) GetL2Token(l1Token common.Address) (common.Address, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GetL2Token(&_ArbitrumL1BridgeAdapter.CallOpts, l1Token)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) GetL2Token(l1Token common.Address) (common.Address, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GetL2Token(&_ArbitrumL1BridgeAdapter.CallOpts, l1Token)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactor) FinalizeWithdrawERC20FromL2(opts *bind.TransactOpts, l2Sender common.Address, l1Receiver common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.contract.Transact(opts, "finalizeWithdrawERC20FromL2", l2Sender, l1Receiver, arbitrumFinalizationPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) FinalizeWithdrawERC20FromL2(l2Sender common.Address, l1Receiver common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.FinalizeWithdrawERC20FromL2(&_ArbitrumL1BridgeAdapter.TransactOpts, l2Sender, l1Receiver, arbitrumFinalizationPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactorSession) FinalizeWithdrawERC20FromL2(l2Sender common.Address, l1Receiver common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.FinalizeWithdrawERC20FromL2(&_ArbitrumL1BridgeAdapter.TransactOpts, l2Sender, l1Receiver, arbitrumFinalizationPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactor) SendERC20(opts *bind.TransactOpts, l1Token common.Address, arg1 common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.contract.Transact(opts, "sendERC20", l1Token, arg1, recipient, amount)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) SendERC20(l1Token common.Address, arg1 common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.SendERC20(&_ArbitrumL1BridgeAdapter.TransactOpts, l1Token, arg1, recipient, amount)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactorSession) SendERC20(l1Token common.Address, arg1 common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.SendERC20(&_ArbitrumL1BridgeAdapter.TransactOpts, l1Token, arg1, recipient, amount)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapter) Address() common.Address {
	return _ArbitrumL1BridgeAdapter.address
}

type ArbitrumL1BridgeAdapterInterface interface {
	GASPRICEBID(opts *bind.CallOpts) (*big.Int, error)

	MAXGAS(opts *bind.CallOpts) (*big.Int, error)

	MAXSUBMISSIONCOST(opts *bind.CallOpts) (*big.Int, error)

	GetL2Token(opts *bind.CallOpts, l1Token common.Address) (common.Address, error)

	FinalizeWithdrawERC20FromL2(opts *bind.TransactOpts, l2Sender common.Address, l1Receiver common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, l1Token common.Address, arg1 common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	Address() common.Address
}
