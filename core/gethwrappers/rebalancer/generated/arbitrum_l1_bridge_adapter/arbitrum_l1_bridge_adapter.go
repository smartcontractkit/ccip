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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIL1GatewayRouter\",\"name\":\"l1GatewayRouter\",\"type\":\"address\"},{\"internalType\":\"contractIOutbox\",\"name\":\"l1Outbox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1ERC20Gateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BridgeAddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wanted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InsufficientEthValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MsgShouldNotContainValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MsgValueDoesNotMatchAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GAS_PRICE_BID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SUBMISSION_COST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1Receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arbitrumFinalizationPayload\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeFeeInNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"}],\"name\":\"getL2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60e0604052600080546001600160401b031916905534801561002057600080fd5b50604051620010e5380380620010e5833981016040819052610041916100c3565b6001600160a01b038316158061005e57506001600160a01b038216155b8061007057506001600160a01b038116155b1561008e57604051635e9c404d60e11b815260040160405180910390fd5b6001600160a01b0392831660805290821660c0521660a052610110565b6001600160a01b03811681146100c057600080fd5b50565b6000806000606084860312156100d857600080fd5b83516100e3816100ab565b60208501519093506100f4816100ab565b6040850151909250610105816100ab565b809150509250925092565b60805160a05160c051610f9e62000147600039600061019c015260006102b701526000818161038a01526104e10152610f9e6000f3fe6080604052600436106100705760003560e01c80635f2a9f411161004e5780635f2a9f41146100d757806379a35b4b146100ee578063c985069c14610101578063ea6c2f801461014657600080fd5b80632e4b1fc91461007557806332eb79051461009d57806338314bb2146100b5575b600080fd5b34801561008157600080fd5b5061008a610161565b6040519081526020015b60405180910390f35b3480156100a957600080fd5b5061008a6311e1a30081565b3480156100c157600080fd5b506100d56100d0366004610934565b61018a565b005b3480156100e357600080fd5b5061008a620186a081565b6100d56100fc3660046109c6565b610258565b34801561010d57600080fd5b5061012161011c366004610a17565b610499565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610094565b34801561015257600080fd5b5061008a6602d79883d2000081565b60006101746311e1a300620186a0610a6a565b610185906602d79883d20000610a81565b905090565b600061019882840184610c4c565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166308635a958260000151836020015188888660400151876060015188608001518960a001518a60c001516040518a63ffffffff1660e01b815260040161021f99989796959493929190610d7d565b600060405180830381600087803b15801561023957600080fd5b505af115801561024d573d6000803e3d6000fd5b505050505050505050565b61027a73ffffffffffffffffffffffffffffffffffffffff8516333084610554565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660048301526024820183905285169063095ea7b3906044016020604051808303816000875af115801561030f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103339190610e33565b50600061033e610161565b905080341015610388576040517fe2c5a8f7000000000000000000000000000000000000000000000000000000008152600481018290523460248201526044015b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16634fb1a07b3487868787620186a06311e1a3006602d79883d20000604051806020016040528060008152506040516020016103fb929190610e55565b6040516020818303038152906040526040518963ffffffff1660e01b815260040161042c9796959493929190610e6e565b60006040518083038185885af115801561044a573d6000803e3d6000fd5b50505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526104919190810190610ece565b505050505050565b6040517fa7e28d4800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063a7e28d4890602401602060405180830381865afa15801561052a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061054e9190610f45565b92915050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526105e99085906105ef565b50505050565b6000610651826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166107009092919063ffffffff16565b8051909150156106fb578080602001905181019061066f9190610e33565b6106fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161037f565b505050565b606061070f8484600085610717565b949350505050565b6060824710156107a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161037f565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516107d29190610f62565b60006040518083038185875af1925050503d806000811461080f576040519150601f19603f3d011682016040523d82523d6000602084013e610814565b606091505b509150915061082587838387610830565b979650505050505050565b606083156108c65782516000036108bf5773ffffffffffffffffffffffffffffffffffffffff85163b6108bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161037f565b508161070f565b61070f83838151156108db5781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037f9190610f7e565b73ffffffffffffffffffffffffffffffffffffffff8116811461093157600080fd5b50565b6000806000806060858703121561094a57600080fd5b84356109558161090f565b935060208501356109658161090f565b9250604085013567ffffffffffffffff8082111561098257600080fd5b818701915087601f83011261099657600080fd5b8135818111156109a557600080fd5b8860208285010111156109b757600080fd5b95989497505060200194505050565b600080600080608085870312156109dc57600080fd5b84356109e78161090f565b935060208501356109f78161090f565b92506040850135610a078161090f565b9396929550929360600135925050565b600060208284031215610a2957600080fd5b8135610a348161090f565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808202811582820484141761054e5761054e610a3b565b8082018082111561054e5761054e610a3b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715610ae657610ae6610a94565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610b3357610b33610a94565b604052919050565b600082601f830112610b4c57600080fd5b8135602067ffffffffffffffff821115610b6857610b68610a94565b8160051b610b77828201610aec565b9283528481018201928281019087851115610b9157600080fd5b83870192505b8483101561082557823582529183019190830190610b97565b600067ffffffffffffffff821115610bca57610bca610a94565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112610c0757600080fd5b8135610c1a610c1582610bb0565b610aec565b818152846020838601011115610c2f57600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215610c5e57600080fd5b813567ffffffffffffffff80821115610c7657600080fd5b9083019060e08286031215610c8a57600080fd5b610c92610ac3565b823582811115610ca157600080fd5b610cad87828601610b3b565b8252506020830135602082015260408301356040820152606083013560608201526080830135608082015260a083013560a082015260c083013582811115610cf457600080fd5b610d0087828601610bf6565b60c08301525095945050505050565b60005b83811015610d2a578181015183820152602001610d12565b50506000910152565b60008151808452610d4b816020860160208601610d0f565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6101208082528a51908201819052600090610140830190602090818e01845b82811015610db857815185529383019390830190600101610d9c565b50505083018b905273ffffffffffffffffffffffffffffffffffffffff8a16604084015273ffffffffffffffffffffffffffffffffffffffff891660608401528760808401528660a08401528560c08401528460e0840152828103610100840152610e238185610d33565b9c9b505050505050505050505050565b600060208284031215610e4557600080fd5b81518015158114610a3457600080fd5b82815260406020820152600061070f6040830184610d33565b600073ffffffffffffffffffffffffffffffffffffffff808a16835280891660208401528088166040840152508560608301528460808301528360a083015260e060c0830152610ec160e0830184610d33565b9998505050505050505050565b600060208284031215610ee057600080fd5b815167ffffffffffffffff811115610ef757600080fd5b8201601f81018413610f0857600080fd5b8051610f16610c1582610bb0565b818152856020838501011115610f2b57600080fd5b610f3c826020830160208601610d0f565b95945050505050565b600060208284031215610f5757600080fd5b8151610a348161090f565b60008251610f74818460208701610d0f565b9190910192915050565b602081526000610a346020830184610d3356fea164736f6c6343000813000a",
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

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCaller) GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumL1BridgeAdapter.contract.Call(opts, &out, "getBridgeFeeInNative")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GetBridgeFeeInNative(&_ArbitrumL1BridgeAdapter.CallOpts)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterCallerSession) GetBridgeFeeInNative() (*big.Int, error) {
	return _ArbitrumL1BridgeAdapter.Contract.GetBridgeFeeInNative(&_ArbitrumL1BridgeAdapter.CallOpts)
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

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, l2Sender common.Address, l1Receiver common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.contract.Transact(opts, "finalizeWithdrawERC20", l2Sender, l1Receiver, arbitrumFinalizationPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterSession) FinalizeWithdrawERC20(l2Sender common.Address, l1Receiver common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.FinalizeWithdrawERC20(&_ArbitrumL1BridgeAdapter.TransactOpts, l2Sender, l1Receiver, arbitrumFinalizationPayload)
}

func (_ArbitrumL1BridgeAdapter *ArbitrumL1BridgeAdapterTransactorSession) FinalizeWithdrawERC20(l2Sender common.Address, l1Receiver common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error) {
	return _ArbitrumL1BridgeAdapter.Contract.FinalizeWithdrawERC20(&_ArbitrumL1BridgeAdapter.TransactOpts, l2Sender, l1Receiver, arbitrumFinalizationPayload)
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

	GetBridgeFeeInNative(opts *bind.CallOpts) (*big.Int, error)

	GetL2Token(opts *bind.CallOpts, l1Token common.Address) (common.Address, error)

	FinalizeWithdrawERC20(opts *bind.TransactOpts, l2Sender common.Address, l1Receiver common.Address, arbitrumFinalizationPayload []byte) (*types.Transaction, error)

	SendERC20(opts *bind.TransactOpts, l1Token common.Address, arg1 common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	Address() common.Address
}
