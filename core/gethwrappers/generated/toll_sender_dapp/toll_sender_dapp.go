// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package toll_sender_dapp

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

type CommonEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

var TollSenderDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractEVM2AnyTollOnRampRouterInterface\",\"name\":\"onRampRouter\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destinationContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidAddress\",\"type\":\"address\"}],\"name\":\"InvalidDestinationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"i_destinationChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_destinationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_onRampRouter\",\"outputs\":[{\"internalType\":\"contractEVM2AnyTollOnRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"}],\"name\":\"sendTokens\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b50604051610e91380380610e9183398101604081905261002f9161006c565b6001600160a01b039283166080526001600160401b0390911660a0521660c0526100c5565b6001600160a01b038116811461006957600080fd5b50565b60008060006060848603121561008157600080fd5b835161008c81610054565b60208501519093506001600160401b03811681146100a957600080fd5b60408501519092506100ba81610054565b809150509250925092565b60805160a05160c051610d8261010f6000396000818160c3015261038b01526000818161013601526103dc01526000818161010f0152818161029901526103b30152610d826000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c806389f9ad2a1161005057806389f9ad2a1461010a578063a721719514610131578063f99131251461017157600080fd5b8063181f5a771461006c5780635c1b583a146100be575b600080fd5b6100a86040518060400160405280601481526020017f546f6c6c53656e6465724461707020312e302e3000000000000000000000000081525081565b6040516100b59190610990565b60405180910390f35b6100e57f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100b5565b6100e57f000000000000000000000000000000000000000000000000000000000000000081565b6101587f000000000000000000000000000000000000000000000000000000000000000081565b60405167ffffffffffffffff90911681526020016100b5565b61015861017f366004610a73565b600073ffffffffffffffffffffffffffffffffffffffff83166101eb576040517ffdc6604f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b60005b82518110156103685761025e333085848151811061020e5761020e610b59565b60200260200101516020015186858151811061022c5761022c610b59565b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff16610584909392919063ffffffff16565b82818151811061027057610270610b59565b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008584815181106102ca576102ca610b59565b6020026020010151602001516040518363ffffffff1660e01b815260040161031492919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af1158015610333573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103579190610b88565b5061036181610baa565b90506101ee565b506040805160a0810190915273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660c08301527f000000000000000000000000000000000000000000000000000000000000000016906338fa2f9b907f0000000000000000000000000000000000000000000000000000000000000000908060e081016040516020818303038152906040528152602001338860405160200161044a92919073ffffffffffffffffffffffffffffffffffffffff92831681529116602082015260400190565b60405160208183030381529060405281526020018681526020018660008151811061047757610477610b59565b6020026020010151815260200161051a6040518060400160405280620493e08152602001600015158152506040805182516024820152602092830151151560448083019190915282518083039091018152606490910190915290810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f97a657c90000000000000000000000000000000000000000000000000000000017905290565b8152506040518363ffffffff1660e01b815260040161053a929190610c09565b6020604051808303816000875af1158015610559573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057d9190610d2f565b9392505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261061990859061061f565b50505050565b6000610681826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166107309092919063ffffffff16565b80519091501561072b578080602001905181019061069f9190610b88565b61072b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101e2565b505050565b606061073f8484600085610747565b949350505050565b6060824710156107d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101e2565b843b610841576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101e2565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161086a9190610d59565b60006040518083038185875af1925050503d80600081146108a7576040519150601f19603f3d011682016040523d82523d6000602084013e6108ac565b606091505b50915091506108bc8282866108c7565b979650505050505050565b606083156108d657508161057d565b8251156108e65782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e29190610990565b60005b8381101561093557818101518382015260200161091d565b838111156106195750506000910152565b6000815180845261095e81602086016020860161091a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061057d6020830184610946565b803573ffffffffffffffffffffffffffffffffffffffff811681146109c757600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610a1e57610a1e6109cc565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610a6b57610a6b6109cc565b604052919050565b6000806040808486031215610a8757600080fd5b610a90846109a3565b925060208085013567ffffffffffffffff80821115610aae57600080fd5b818701915087601f830112610ac257600080fd5b813581811115610ad457610ad46109cc565b610ae2848260051b01610a24565b818152848101925060069190911b830184019089821115610b0257600080fd5b928401925b81841015610b495785848b031215610b1f5760008081fd5b610b276109fb565b610b30856109a3565b8152848601358682015283529285019291840191610b07565b8096505050505050509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215610b9a57600080fd5b8151801515811461057d57600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c02577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b6000604067ffffffffffffffff8516835260208181850152845160c083860152610c37610100860182610946565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080878403016060880152610c728383610946565b88860151888203830160808a01528051808352908601945060009350908501905b80841015610cda57610cc6828651805173ffffffffffffffffffffffffffffffffffffffff168252602090810151910152565b938501936001939093019290860190610c93565b506060890151805173ffffffffffffffffffffffffffffffffffffffff1660a08a01526020015160c08901526080890151888203830160e08a01529550610d218187610946565b9a9950505050505050505050565b600060208284031215610d4157600080fd5b815167ffffffffffffffff8116811461057d57600080fd5b60008251610d6b81846020870161091a565b919091019291505056fea164736f6c634300080f000a",
}

var TollSenderDappABI = TollSenderDappMetaData.ABI

var TollSenderDappBin = TollSenderDappMetaData.Bin

func DeployTollSenderDapp(auth *bind.TransactOpts, backend bind.ContractBackend, onRampRouter common.Address, destinationChainId uint64, destinationContract common.Address) (common.Address, *types.Transaction, *TollSenderDapp, error) {
	parsed, err := TollSenderDappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TollSenderDappBin), backend, onRampRouter, destinationChainId, destinationContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TollSenderDapp{TollSenderDappCaller: TollSenderDappCaller{contract: contract}, TollSenderDappTransactor: TollSenderDappTransactor{contract: contract}, TollSenderDappFilterer: TollSenderDappFilterer{contract: contract}}, nil
}

type TollSenderDapp struct {
	address common.Address
	abi     abi.ABI
	TollSenderDappCaller
	TollSenderDappTransactor
	TollSenderDappFilterer
}

type TollSenderDappCaller struct {
	contract *bind.BoundContract
}

type TollSenderDappTransactor struct {
	contract *bind.BoundContract
}

type TollSenderDappFilterer struct {
	contract *bind.BoundContract
}

type TollSenderDappSession struct {
	Contract     *TollSenderDapp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TollSenderDappCallerSession struct {
	Contract *TollSenderDappCaller
	CallOpts bind.CallOpts
}

type TollSenderDappTransactorSession struct {
	Contract     *TollSenderDappTransactor
	TransactOpts bind.TransactOpts
}

type TollSenderDappRaw struct {
	Contract *TollSenderDapp
}

type TollSenderDappCallerRaw struct {
	Contract *TollSenderDappCaller
}

type TollSenderDappTransactorRaw struct {
	Contract *TollSenderDappTransactor
}

func NewTollSenderDapp(address common.Address, backend bind.ContractBackend) (*TollSenderDapp, error) {
	abi, err := abi.JSON(strings.NewReader(TollSenderDappABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindTollSenderDapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TollSenderDapp{address: address, abi: abi, TollSenderDappCaller: TollSenderDappCaller{contract: contract}, TollSenderDappTransactor: TollSenderDappTransactor{contract: contract}, TollSenderDappFilterer: TollSenderDappFilterer{contract: contract}}, nil
}

func NewTollSenderDappCaller(address common.Address, caller bind.ContractCaller) (*TollSenderDappCaller, error) {
	contract, err := bindTollSenderDapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TollSenderDappCaller{contract: contract}, nil
}

func NewTollSenderDappTransactor(address common.Address, transactor bind.ContractTransactor) (*TollSenderDappTransactor, error) {
	contract, err := bindTollSenderDapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TollSenderDappTransactor{contract: contract}, nil
}

func NewTollSenderDappFilterer(address common.Address, filterer bind.ContractFilterer) (*TollSenderDappFilterer, error) {
	contract, err := bindTollSenderDapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TollSenderDappFilterer{contract: contract}, nil
}

func bindTollSenderDapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TollSenderDappABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_TollSenderDapp *TollSenderDappRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TollSenderDapp.Contract.TollSenderDappCaller.contract.Call(opts, result, method, params...)
}

func (_TollSenderDapp *TollSenderDappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.TollSenderDappTransactor.contract.Transfer(opts)
}

func (_TollSenderDapp *TollSenderDappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.TollSenderDappTransactor.contract.Transact(opts, method, params...)
}

func (_TollSenderDapp *TollSenderDappCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TollSenderDapp.Contract.contract.Call(opts, result, method, params...)
}

func (_TollSenderDapp *TollSenderDappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.contract.Transfer(opts)
}

func (_TollSenderDapp *TollSenderDappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.contract.Transact(opts, method, params...)
}

func (_TollSenderDapp *TollSenderDappCaller) IDestinationChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TollSenderDapp.contract.Call(opts, &out, "i_destinationChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_TollSenderDapp *TollSenderDappSession) IDestinationChainId() (uint64, error) {
	return _TollSenderDapp.Contract.IDestinationChainId(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCallerSession) IDestinationChainId() (uint64, error) {
	return _TollSenderDapp.Contract.IDestinationChainId(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCaller) IDestinationContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TollSenderDapp.contract.Call(opts, &out, "i_destinationContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TollSenderDapp *TollSenderDappSession) IDestinationContract() (common.Address, error) {
	return _TollSenderDapp.Contract.IDestinationContract(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCallerSession) IDestinationContract() (common.Address, error) {
	return _TollSenderDapp.Contract.IDestinationContract(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCaller) IOnRampRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TollSenderDapp.contract.Call(opts, &out, "i_onRampRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TollSenderDapp *TollSenderDappSession) IOnRampRouter() (common.Address, error) {
	return _TollSenderDapp.Contract.IOnRampRouter(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCallerSession) IOnRampRouter() (common.Address, error) {
	return _TollSenderDapp.Contract.IOnRampRouter(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TollSenderDapp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_TollSenderDapp *TollSenderDappSession) TypeAndVersion() (string, error) {
	return _TollSenderDapp.Contract.TypeAndVersion(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCallerSession) TypeAndVersion() (string, error) {
	return _TollSenderDapp.Contract.TypeAndVersion(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappTransactor) SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokensAndAmounts []CommonEVMTokenAndAmount) (*types.Transaction, error) {
	return _TollSenderDapp.contract.Transact(opts, "sendTokens", destinationAddress, tokensAndAmounts)
}

func (_TollSenderDapp *TollSenderDappSession) SendTokens(destinationAddress common.Address, tokensAndAmounts []CommonEVMTokenAndAmount) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.SendTokens(&_TollSenderDapp.TransactOpts, destinationAddress, tokensAndAmounts)
}

func (_TollSenderDapp *TollSenderDappTransactorSession) SendTokens(destinationAddress common.Address, tokensAndAmounts []CommonEVMTokenAndAmount) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.SendTokens(&_TollSenderDapp.TransactOpts, destinationAddress, tokensAndAmounts)
}

func (_TollSenderDapp *TollSenderDapp) Address() common.Address {
	return _TollSenderDapp.address
}

type TollSenderDappInterface interface {
	IDestinationChainId(opts *bind.CallOpts) (uint64, error)

	IDestinationContract(opts *bind.CallOpts) (common.Address, error)

	IOnRampRouter(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokensAndAmounts []CommonEVMTokenAndAmount) (*types.Transaction, error)

	Address() common.Address
}
