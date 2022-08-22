// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package subscription_sender_dapp

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
)

var SubscriptionSenderDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractAny2EVMSubscriptionOnRampRouterInterface\",\"name\":\"onRampRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidAddress\",\"type\":\"address\"}],\"name\":\"InvalidDestinationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"i_destinationChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_destinationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_onRampRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMSubscriptionOnRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"sendTokens\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b50604051610e08380380610e0883398101604081905261002f91610064565b6001600160a01b0392831660805260a0919091521660c0526100a7565b6001600160a01b038116811461006157600080fd5b50565b60008060006060848603121561007957600080fd5b83516100848161004c565b60208501516040860151919450925061009c8161004c565b809150509250925092565b60805160a05160c051610d176100f16000396000818160c301526103d601526000818161013601526103aa01526000818161010f015281816102a0015261036e0152610d176000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c806389f9ad2a1161005057806389f9ad2a1461010a578063a721719514610131578063e83f967b1461016657600080fd5b8063181f5a771461006c5780635c1b583a146100be575b600080fd5b6100a86040518060400160405280601c81526020017f537562736372697074696f6e53656e6465724461707020312e302e300000000081525081565b6040516100b591906108dd565b60405180910390f35b6100e57f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100b5565b6100e57f000000000000000000000000000000000000000000000000000000000000000081565b6101587f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100b5565b610179610174366004610a22565b610192565b60405167ffffffffffffffff90911681526020016100b5565b600073ffffffffffffffffffffffffffffffffffffffff84166101fe576040517ffdc6604f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024015b60405180910390fd5b60005b835181101561036b57610269333085848151811061022157610221610af8565b602002602001015187858151811061023b5761023b610af8565b602002602001015173ffffffffffffffffffffffffffffffffffffffff166104d2909392919063ffffffff16565b83818151811061027b5761027b610af8565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008584815181106102d1576102d1610af8565b60200260200101516040518363ffffffff1660e01b815260040161031792919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af1158015610336573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061035a9190610b27565b5061036481610b49565b9050610201565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16630d58bf0c7f00000000000000000000000000000000000000000000000000000000000000006040518060a001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168152602001338960405160200161044592919073ffffffffffffffffffffffffffffffffffffffff92831681529116602082015260400190565b604051602081830303815290604052815260200187815260200186815260200160008152506040518363ffffffff1660e01b8152600401610487929190610be3565b6020604051808303816000875af11580156104a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ca9190610cc4565b949350505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261056790859061056d565b50505050565b60006105cf826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661067e9092919063ffffffff16565b80519091501561067957808060200190518101906105ed9190610b27565b610679576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101f5565b505050565b60606104ca8484600085610694565b9392505050565b606082471015610726576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101f5565b843b61078e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101f5565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516107b79190610cee565b60006040518083038185875af1925050503d80600081146107f4576040519150601f19603f3d011682016040523d82523d6000602084013e6107f9565b606091505b5091509150610809828286610814565b979650505050505050565b6060831561082357508161068d565b8251156108335782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f591906108dd565b60005b8381101561088257818101518382015260200161086a565b838111156105675750506000910152565b600081518084526108ab816020860160208601610867565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061068d6020830184610893565b73ffffffffffffffffffffffffffffffffffffffff8116811461091257600080fd5b50565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561098b5761098b610915565b604052919050565b600067ffffffffffffffff8211156109ad576109ad610915565b5060051b60200190565b600082601f8301126109c857600080fd5b813560206109dd6109d883610993565b610944565b82815260059290921b840181019181810190868411156109fc57600080fd5b8286015b84811015610a175780358352918301918301610a00565b509695505050505050565b600080600060608486031215610a3757600080fd5b8335610a42816108f0565b925060208481013567ffffffffffffffff80821115610a6057600080fd5b818701915087601f830112610a7457600080fd5b8135610a826109d882610993565b81815260059190911b8301840190848101908a831115610aa157600080fd5b938501935b82851015610ac8578435610ab9816108f0565b82529385019390850190610aa6565b965050506040870135925080831115610ae057600080fd5b5050610aee868287016109b7565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215610b3957600080fd5b8151801515811461068d57600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610ba1577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b600081518084526020808501945080840160005b83811015610bd857815187529582019590820190600101610bbc565b509495945050505050565b8281526000602060408184015273ffffffffffffffffffffffffffffffffffffffff8085511660408501528185015160a06060860152610c2660e0860182610893565b60408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087830381016080890152815180845291860193506000929091908601905b80841015610c8c57845186168252938601936001939093019290860190610c6a565b5060608901519550818882030160a0890152610ca88187610ba8565b95505050505050608084015160c0840152809150509392505050565b600060208284031215610cd657600080fd5b815167ffffffffffffffff8116811461068d57600080fd5b60008251610d00818460208701610867565b919091019291505056fea164736f6c634300080f000a",
}

var SubscriptionSenderDappABI = SubscriptionSenderDappMetaData.ABI

var SubscriptionSenderDappBin = SubscriptionSenderDappMetaData.Bin

func DeploySubscriptionSenderDapp(auth *bind.TransactOpts, backend bind.ContractBackend, onRampRouter common.Address, destinationChainId *big.Int, destinationContract common.Address) (common.Address, *types.Transaction, *SubscriptionSenderDapp, error) {
	parsed, err := SubscriptionSenderDappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SubscriptionSenderDappBin), backend, onRampRouter, destinationChainId, destinationContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SubscriptionSenderDapp{SubscriptionSenderDappCaller: SubscriptionSenderDappCaller{contract: contract}, SubscriptionSenderDappTransactor: SubscriptionSenderDappTransactor{contract: contract}, SubscriptionSenderDappFilterer: SubscriptionSenderDappFilterer{contract: contract}}, nil
}

type SubscriptionSenderDapp struct {
	address common.Address
	abi     abi.ABI
	SubscriptionSenderDappCaller
	SubscriptionSenderDappTransactor
	SubscriptionSenderDappFilterer
}

type SubscriptionSenderDappCaller struct {
	contract *bind.BoundContract
}

type SubscriptionSenderDappTransactor struct {
	contract *bind.BoundContract
}

type SubscriptionSenderDappFilterer struct {
	contract *bind.BoundContract
}

type SubscriptionSenderDappSession struct {
	Contract     *SubscriptionSenderDapp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type SubscriptionSenderDappCallerSession struct {
	Contract *SubscriptionSenderDappCaller
	CallOpts bind.CallOpts
}

type SubscriptionSenderDappTransactorSession struct {
	Contract     *SubscriptionSenderDappTransactor
	TransactOpts bind.TransactOpts
}

type SubscriptionSenderDappRaw struct {
	Contract *SubscriptionSenderDapp
}

type SubscriptionSenderDappCallerRaw struct {
	Contract *SubscriptionSenderDappCaller
}

type SubscriptionSenderDappTransactorRaw struct {
	Contract *SubscriptionSenderDappTransactor
}

func NewSubscriptionSenderDapp(address common.Address, backend bind.ContractBackend) (*SubscriptionSenderDapp, error) {
	abi, err := abi.JSON(strings.NewReader(SubscriptionSenderDappABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindSubscriptionSenderDapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubscriptionSenderDapp{address: address, abi: abi, SubscriptionSenderDappCaller: SubscriptionSenderDappCaller{contract: contract}, SubscriptionSenderDappTransactor: SubscriptionSenderDappTransactor{contract: contract}, SubscriptionSenderDappFilterer: SubscriptionSenderDappFilterer{contract: contract}}, nil
}

func NewSubscriptionSenderDappCaller(address common.Address, caller bind.ContractCaller) (*SubscriptionSenderDappCaller, error) {
	contract, err := bindSubscriptionSenderDapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionSenderDappCaller{contract: contract}, nil
}

func NewSubscriptionSenderDappTransactor(address common.Address, transactor bind.ContractTransactor) (*SubscriptionSenderDappTransactor, error) {
	contract, err := bindSubscriptionSenderDapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionSenderDappTransactor{contract: contract}, nil
}

func NewSubscriptionSenderDappFilterer(address common.Address, filterer bind.ContractFilterer) (*SubscriptionSenderDappFilterer, error) {
	contract, err := bindSubscriptionSenderDapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubscriptionSenderDappFilterer{contract: contract}, nil
}

func bindSubscriptionSenderDapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubscriptionSenderDappABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubscriptionSenderDapp.Contract.SubscriptionSenderDappCaller.contract.Call(opts, result, method, params...)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.SubscriptionSenderDappTransactor.contract.Transfer(opts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.SubscriptionSenderDappTransactor.contract.Transact(opts, method, params...)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubscriptionSenderDapp.Contract.contract.Call(opts, result, method, params...)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.contract.Transfer(opts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.contract.Transact(opts, method, params...)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCaller) IDestinationChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubscriptionSenderDapp.contract.Call(opts, &out, "i_destinationChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SubscriptionSenderDapp *SubscriptionSenderDappSession) IDestinationChainId() (*big.Int, error) {
	return _SubscriptionSenderDapp.Contract.IDestinationChainId(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCallerSession) IDestinationChainId() (*big.Int, error) {
	return _SubscriptionSenderDapp.Contract.IDestinationChainId(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCaller) IDestinationContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SubscriptionSenderDapp.contract.Call(opts, &out, "i_destinationContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SubscriptionSenderDapp *SubscriptionSenderDappSession) IDestinationContract() (common.Address, error) {
	return _SubscriptionSenderDapp.Contract.IDestinationContract(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCallerSession) IDestinationContract() (common.Address, error) {
	return _SubscriptionSenderDapp.Contract.IDestinationContract(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCaller) IOnRampRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SubscriptionSenderDapp.contract.Call(opts, &out, "i_onRampRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SubscriptionSenderDapp *SubscriptionSenderDappSession) IOnRampRouter() (common.Address, error) {
	return _SubscriptionSenderDapp.Contract.IOnRampRouter(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCallerSession) IOnRampRouter() (common.Address, error) {
	return _SubscriptionSenderDapp.Contract.IOnRampRouter(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SubscriptionSenderDapp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_SubscriptionSenderDapp *SubscriptionSenderDappSession) TypeAndVersion() (string, error) {
	return _SubscriptionSenderDapp.Contract.TypeAndVersion(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCallerSession) TypeAndVersion() (string, error) {
	return _SubscriptionSenderDapp.Contract.TypeAndVersion(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappTransactor) SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.contract.Transact(opts, "sendTokens", destinationAddress, tokens, amounts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappSession) SendTokens(destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.SendTokens(&_SubscriptionSenderDapp.TransactOpts, destinationAddress, tokens, amounts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappTransactorSession) SendTokens(destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.SendTokens(&_SubscriptionSenderDapp.TransactOpts, destinationAddress, tokens, amounts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDapp) Address() common.Address {
	return _SubscriptionSenderDapp.address
}

type SubscriptionSenderDappInterface interface {
	IDestinationChainId(opts *bind.CallOpts) (*big.Int, error)

	IDestinationContract(opts *bind.CallOpts) (common.Address, error)

	IOnRampRouter(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error)

	Address() common.Address
}
