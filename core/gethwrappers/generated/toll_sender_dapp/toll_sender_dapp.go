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

var TollSenderDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractEVM2AnyTollOnRampRouterInterface\",\"name\":\"onRampRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidAddress\",\"type\":\"address\"}],\"name\":\"InvalidDestinationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"i_destinationChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_destinationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_onRampRouter\",\"outputs\":[{\"internalType\":\"contractEVM2AnyTollOnRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"sendTokens\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b50604051610f1d380380610f1d83398101604081905261002f91610064565b6001600160a01b0392831660805260a0919091521660c0526100a7565b6001600160a01b038116811461006157600080fd5b50565b60008060006060848603121561007957600080fd5b83516100848161004c565b60208501516040860151919450925061009c8161004c565b809150509250925092565b60805160a05160c051610e2c6100f16000396000818160c3015261038e01526000818161013601526103e001526000818161010f015281816102a001526103b70152610e2c6000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c806389f9ad2a1161005057806389f9ad2a1461010a578063a721719514610131578063e83f967b1461016657600080fd5b8063181f5a771461006c5780635c1b583a146100be575b600080fd5b6100a86040518060400160405280601481526020017f546f6c6c53656e6465724461707020312e302e3000000000000000000000000081525081565b6040516100b591906109a2565b60405180910390f35b6100e57f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100b5565b6100e57f000000000000000000000000000000000000000000000000000000000000000081565b6101587f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100b5565b610179610174366004610aeb565b610192565b60405167ffffffffffffffff90911681526020016100b5565b600073ffffffffffffffffffffffffffffffffffffffff84166101fe576040517ffdc6604f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024015b60405180910390fd5b60005b835181101561036b57610269333085848151811061022157610221610bbd565b602002602001015187858151811061023b5761023b610bbd565b602002602001015173ffffffffffffffffffffffffffffffffffffffff16610597909392919063ffffffff16565b83818151811061027b5761027b610bbd565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008584815181106102d1576102d1610bbd565b60200260200101516040518363ffffffff1660e01b815260040161031792919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af1158015610336573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061035a9190610bec565b5061036481610c0e565b9050610201565b506040805160e0810190915273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081166101008301527f000000000000000000000000000000000000000000000000000000000000000016906370587681907f0000000000000000000000000000000000000000000000000000000000000000908061012081016040516020818303038152906040528152602001338960405160200161044f92919073ffffffffffffffffffffffffffffffffffffffff92831681529116602082015260400190565b60405160208183030381529060405281526020018781526020018681526020018760008151811061048257610482610bbd565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200161052c6040518060200160405280620493e081525060408051915160248084019190915281518084039091018152604490920190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f97a657c90000000000000000000000000000000000000000000000000000000017905290565b8152506040518363ffffffff1660e01b815260040161054c929190610ca8565b6020604051808303816000875af115801561056b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061058f9190610dd9565b949350505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261062c908590610632565b50505050565b6000610694826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166107439092919063ffffffff16565b80519091501561073e57808060200190518101906106b29190610bec565b61073e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101f5565b505050565b606061058f8484600085610759565b9392505050565b6060824710156107eb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101f5565b843b610853576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101f5565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161087c9190610e03565b60006040518083038185875af1925050503d80600081146108b9576040519150601f19603f3d011682016040523d82523d6000602084013e6108be565b606091505b50915091506108ce8282866108d9565b979650505050505050565b606083156108e8575081610752565b8251156108f85782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f591906109a2565b60005b8381101561094757818101518382015260200161092f565b8381111561062c5750506000910152565b6000815180845261097081602086016020860161092c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006107526020830184610958565b803573ffffffffffffffffffffffffffffffffffffffff811681146109d957600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610a5457610a546109de565b604052919050565b600067ffffffffffffffff821115610a7657610a766109de565b5060051b60200190565b600082601f830112610a9157600080fd5b81356020610aa6610aa183610a5c565b610a0d565b82815260059290921b84018101918181019086841115610ac557600080fd5b8286015b84811015610ae05780358352918301918301610ac9565b509695505050505050565b600080600060608486031215610b0057600080fd5b610b09846109b5565b925060208085013567ffffffffffffffff80821115610b2757600080fd5b818701915087601f830112610b3b57600080fd5b8135610b49610aa182610a5c565b81815260059190911b8301840190848101908a831115610b6857600080fd5b938501935b82851015610b8d57610b7e856109b5565b82529385019390850190610b6d565b965050506040870135925080831115610ba557600080fd5b5050610bb386828701610a80565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215610bfe57600080fd5b8151801515811461075257600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c66577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b600081518084526020808501945080840160005b83811015610c9d57815187529582019590820190600101610c81565b509495945050505050565b82815260006020604081840152835160e06040850152610ccc610120850182610958565b9050818501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080868403016060870152610d078383610958565b6040880151878203830160808901528051808352908601945060009350908501905b80841015610d5f57845173ffffffffffffffffffffffffffffffffffffffff168252938501936001939093019290850190610d29565b5060608801519450818782030160a0880152610d7b8186610c6d565b94505060808701519250610da760c087018473ffffffffffffffffffffffffffffffffffffffff169052565b60a087015160e087015260c0870151925080868503016101008701525050610dcf8282610958565b9695505050505050565b600060208284031215610deb57600080fd5b815167ffffffffffffffff8116811461075257600080fd5b60008251610e1581846020870161092c565b919091019291505056fea164736f6c634300080f000a",
}

var TollSenderDappABI = TollSenderDappMetaData.ABI

var TollSenderDappBin = TollSenderDappMetaData.Bin

func DeployTollSenderDapp(auth *bind.TransactOpts, backend bind.ContractBackend, onRampRouter common.Address, destinationChainId *big.Int, destinationContract common.Address) (common.Address, *types.Transaction, *TollSenderDapp, error) {
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

func (_TollSenderDapp *TollSenderDappCaller) IDestinationChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TollSenderDapp.contract.Call(opts, &out, "i_destinationChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TollSenderDapp *TollSenderDappSession) IDestinationChainId() (*big.Int, error) {
	return _TollSenderDapp.Contract.IDestinationChainId(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCallerSession) IDestinationChainId() (*big.Int, error) {
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

func (_TollSenderDapp *TollSenderDappTransactor) SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TollSenderDapp.contract.Transact(opts, "sendTokens", destinationAddress, tokens, amounts)
}

func (_TollSenderDapp *TollSenderDappSession) SendTokens(destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.SendTokens(&_TollSenderDapp.TransactOpts, destinationAddress, tokens, amounts)
}

func (_TollSenderDapp *TollSenderDappTransactorSession) SendTokens(destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.SendTokens(&_TollSenderDapp.TransactOpts, destinationAddress, tokens, amounts)
}

func (_TollSenderDapp *TollSenderDapp) Address() common.Address {
	return _TollSenderDapp.address
}

type TollSenderDappInterface interface {
	IDestinationChainId(opts *bind.CallOpts) (*big.Int, error)

	IDestinationContract(opts *bind.CallOpts) (common.Address, error)

	IOnRampRouter(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error)

	Address() common.Address
}
