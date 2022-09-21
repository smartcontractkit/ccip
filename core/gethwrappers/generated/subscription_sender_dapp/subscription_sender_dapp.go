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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractEVM2AnySubscriptionOnRampRouterInterface\",\"name\":\"onRampRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidAddress\",\"type\":\"address\"}],\"name\":\"InvalidDestinationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"i_destinationChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_destinationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_onRampRouter\",\"outputs\":[{\"internalType\":\"contractEVM2AnySubscriptionOnRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"sendTokens\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b50604051610c28380380610c2883398101604081905261002f91610064565b6001600160a01b0392831660805260a0919091521660c0526100a7565b6001600160a01b038116811461006157600080fd5b50565b60008060006060848603121561007957600080fd5b83516100848161004c565b60208501516040860151919450925061009c8161004c565b809150509250925092565b60805160a05160c051610b376100f16000396000818160c3015261037b015260008181610129015261034f0152600081816101020152818161025f01526103200152610b376000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c806389f9ad2a1161005057806389f9ad2a146100fd578063a721719514610124578063e83f967b1461015957600080fd5b8063181f5a771461006c5780635c1b583a146100be575b600080fd5b6100a86040518060400160405280601c81526020017f537562736372697074696f6e53656e6465724461707020312e302e300000000081525081565b6040516100b591906107bd565b60405180910390f35b6100e57f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100b5565b6100e57f000000000000000000000000000000000000000000000000000000000000000081565b61014b7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100b5565b61016c6101673660046108be565b610185565b60405167ffffffffffffffff90911681526020016100b5565b60006001600160a01b0384166101d7576040517ffdc6604f0000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024015b60405180910390fd5b60005b835181101561031d5761023533308584815181106101fa576101fa610994565b602002602001015187858151811061021457610214610994565b60200260200101516001600160a01b031661045f909392919063ffffffff16565b83818151811061024757610247610994565b60200260200101516001600160a01b031663095ea7b37f000000000000000000000000000000000000000000000000000000000000000085848151811061029057610290610994565b60200260200101516040518363ffffffff1660e01b81526004016102c99291906001600160a01b03929092168252602082015260400190565b6020604051808303816000875af11580156102e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030c91906109aa565b50610316816109cc565b90506101da565b507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630d58bf0c7f00000000000000000000000000000000000000000000000000000000000000006040518060a001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200133896040516020016103d09291906001600160a01b0392831681529116602082015260400190565b6040516020818303038152906040528152602001878152602001868152602001620493e08152506040518363ffffffff1660e01b8152600401610414929190610a2e565b6020604051808303816000875af1158015610433573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104579190610ae4565b949350505050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526104e79085906104ed565b50505050565b6000610542826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166105d79092919063ffffffff16565b8051909150156105d2578080602001905181019061056091906109aa565b6105d25760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101ce565b505050565b606061045784846000856105ed565b9392505050565b6060824710156106655760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101ce565b843b6106b35760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101ce565b600080866001600160a01b031685876040516106cf9190610b0e565b60006040518083038185875af1925050503d806000811461070c576040519150601f19603f3d011682016040523d82523d6000602084013e610711565b606091505b509150915061072182828661072c565b979650505050505050565b6060831561073b5750816105e6565b82511561074b5782518084602001fd5b8160405162461bcd60e51b81526004016101ce91906107bd565b60005b83811015610780578181015183820152602001610768565b838111156104e75750506000910152565b600081518084526107a9816020860160208601610765565b601f01601f19169290920160200192915050565b6020815260006105e66020830184610791565b6001600160a01b03811681146107e557600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610827576108276107e8565b604052919050565b600067ffffffffffffffff821115610849576108496107e8565b5060051b60200190565b600082601f83011261086457600080fd5b813560206108796108748361082f565b6107fe565b82815260059290921b8401810191818101908684111561089857600080fd5b8286015b848110156108b3578035835291830191830161089c565b509695505050505050565b6000806000606084860312156108d357600080fd5b83356108de816107d0565b925060208481013567ffffffffffffffff808211156108fc57600080fd5b818701915087601f83011261091057600080fd5b813561091e6108748261082f565b81815260059190911b8301840190848101908a83111561093d57600080fd5b938501935b82851015610964578435610955816107d0565b82529385019390850190610942565b96505050604087013592508083111561097c57600080fd5b505061098a86828701610853565b9150509250925092565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156109bc57600080fd5b815180151581146105e657600080fd5b6000600182016109ec57634e487b7160e01b600052601160045260246000fd5b5060010190565b600081518084526020808501945080840160005b83811015610a2357815187529582019590820190600101610a07565b509495945050505050565b828152600060206040818401526001600160a01b038085511660408501528185015160a06060860152610a6460e0860182610791565b6040870151603f1987830381016080890152815180845291860193506000929091908601905b80841015610aac57845186168252938601936001939093019290860190610a8a565b5060608901519550818882030160a0890152610ac881876109f3565b95505050505050608084015160c0840152809150509392505050565b600060208284031215610af657600080fd5b815167ffffffffffffffff811681146105e657600080fd5b60008251610b20818460208701610765565b919091019291505056fea164736f6c634300080f000a",
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
