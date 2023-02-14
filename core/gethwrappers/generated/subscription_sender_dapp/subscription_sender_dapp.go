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

type CCIPEVM2AnySubscriptionMessage struct {
	Receiver         []byte
	Data             []byte
	TokenAmounts []CCIPEVMTokenAmount
	ExtraArgs        []byte
}

type CCIPEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

var SubscriptionSenderDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractEVM2AnySubscriptionOnRampRouterInterface\",\"name\":\"onRampRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_destinationChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_onRampRouter\",\"outputs\":[{\"internalType\":\"contractEVM2AnySubscriptionOnRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.EVM2AnySubscriptionMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"sendMessage\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unfundSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161102738038061102783398101604081905261002f91610045565b6001600160a01b0390911660805260a05261007f565b6000806040838503121561005857600080fd5b82516001600160a01b038116811461006f57600080fd5b6020939093015192949293505050565b60805160a051610f596100ce6000396000818161016e01526104eb01526000818161010f015281816101db01528181610287015281816103b1015281816104be015261058e0152610f596000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c806389f9ad2a1161005057806389f9ad2a1461010a57806395e712db14610156578063a72171951461016957600080fd5b8063181f5a77146100775780633c5457ce146100c957806374c0f26c146100de575b600080fd5b6100b36040518060400160405280601c81526020017f537562736372697074696f6e53656e6465724461707020312e302e300000000081525081565b6040516100c09190610a13565b60405180910390f35b6100dc6100d7366004610a4b565b61019e565b005b6100f16100ec366004610c7c565b6102fc565b60405167ffffffffffffffff90911681526020016100c0565b6101317f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100c0565b6100dc610164366004610d55565b61055f565b6101907f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100c0565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660048301526024820183905283169063095ea7b3906044016020604051808303816000875af1158015610233573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102579190610d6e565b506040517fc1060653000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063c106065390602401600060405180830381600087803b1580156102e057600080fd5b505af11580156102f4573d6000803e3d6000fd5b505050505050565b6040810151600090815b815181101561048057610376333084848151811061032657610326610d90565b60200260200101516020015185858151811061034457610344610d90565b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff16610602909392919063ffffffff16565b81818151811061038857610388610d90565b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008484815181106103e2576103e2610d90565b6020026020010151602001516040518363ffffffff1660e01b815260040161042c92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af115801561044b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061046f9190610d6e565b5061047981610dbf565b9050610306565b506040517f97fa090a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906397fa090a90610515907f0000000000000000000000000000000000000000000000000000000000000000908790600401610e1e565b6020604051808303816000875af1158015610534573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105589190610f06565b9392505050565b6040517f95e712db000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906395e712db90602401600060405180830381600087803b1580156105e757600080fd5b505af11580156105fb573d6000803e3d6000fd5b5050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261069790859061069d565b50505050565b60006106ff826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166107b39092919063ffffffff16565b8051909150156107ae578080602001905181019061071d9190610d6e565b6107ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b505050565b60606107c284846000856107ca565b949350505050565b60608247101561085c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016107a5565b843b6108c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016107a5565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516108ed9190610f30565b60006040518083038185875af1925050503d806000811461092a576040519150601f19603f3d011682016040523d82523d6000602084013e61092f565b606091505b509150915061093f82828661094a565b979650505050505050565b60608315610959575081610558565b8251156109695782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107a59190610a13565b60005b838110156109b85781810151838201526020016109a0565b838111156106975750506000910152565b600081518084526109e181602086016020860161099d565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061055860208301846109c9565b73ffffffffffffffffffffffffffffffffffffffff81168114610a4857600080fd5b50565b60008060408385031215610a5e57600080fd5b8235610a6981610a26565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610ac957610ac9610a77565b60405290565b6040516080810167ffffffffffffffff81118282101715610ac957610ac9610a77565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610b3957610b39610a77565b604052919050565b600082601f830112610b5257600080fd5b813567ffffffffffffffff811115610b6c57610b6c610a77565b610b9d60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610af2565b818152846020838601011115610bb257600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112610be057600080fd5b8135602067ffffffffffffffff821115610bfc57610bfc610a77565b610c0a818360051b01610af2565b82815260069290921b84018101918181019086841115610c2957600080fd5b8286015b84811015610c715760408189031215610c465760008081fd5b610c4e610aa6565b8135610c5981610a26565b81528185013585820152835291830191604001610c2d565b509695505050505050565b600060208284031215610c8e57600080fd5b813567ffffffffffffffff80821115610ca657600080fd5b9083019060808286031215610cba57600080fd5b610cc2610acf565b823582811115610cd157600080fd5b610cdd87828601610b41565b825250602083013582811115610cf257600080fd5b610cfe87828601610b41565b602083015250604083013582811115610d1657600080fd5b610d2287828601610bcf565b604083015250606083013582811115610d3a57600080fd5b610d4687828601610b41565b60608301525095945050505050565b600060208284031215610d6757600080fd5b5035919050565b600060208284031215610d8057600080fd5b8151801515811461055857600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610e17577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b60006040848352602081818501528451608083860152610e4160c08601826109c9565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080878403016060880152610e7c83836109c9565b88860151888203830160808a01528051808352908601945060009350908501905b80841015610edc578451805173ffffffffffffffffffffffffffffffffffffffff16835286015186830152938501936001939093019290860190610e9d565b5060608901519550818882030160a0890152610ef881876109c9565b9a9950505050505050505050565b600060208284031215610f1857600080fd5b815167ffffffffffffffff8116811461055857600080fd5b60008251610f4281846020870161099d565b919091019291505056fea164736f6c634300080f000a",
}

var SubscriptionSenderDappABI = SubscriptionSenderDappMetaData.ABI

var SubscriptionSenderDappBin = SubscriptionSenderDappMetaData.Bin

func DeploySubscriptionSenderDapp(auth *bind.TransactOpts, backend bind.ContractBackend, onRampRouter common.Address, destinationChainId *big.Int) (common.Address, *types.Transaction, *SubscriptionSenderDapp, error) {
	parsed, err := SubscriptionSenderDappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SubscriptionSenderDappBin), backend, onRampRouter, destinationChainId)
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

func (_SubscriptionSenderDapp *SubscriptionSenderDappTransactor) FundSubscription(opts *bind.TransactOpts, feeToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.contract.Transact(opts, "fundSubscription", feeToken, amount)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappSession) FundSubscription(feeToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.FundSubscription(&_SubscriptionSenderDapp.TransactOpts, feeToken, amount)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappTransactorSession) FundSubscription(feeToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.FundSubscription(&_SubscriptionSenderDapp.TransactOpts, feeToken, amount)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappTransactor) SendMessage(opts *bind.TransactOpts, message CCIPEVM2AnySubscriptionMessage) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.contract.Transact(opts, "sendMessage", message)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappSession) SendMessage(message CCIPEVM2AnySubscriptionMessage) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.SendMessage(&_SubscriptionSenderDapp.TransactOpts, message)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappTransactorSession) SendMessage(message CCIPEVM2AnySubscriptionMessage) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.SendMessage(&_SubscriptionSenderDapp.TransactOpts, message)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappTransactor) UnfundSubscription(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.contract.Transact(opts, "unfundSubscription", amount)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappSession) UnfundSubscription(amount *big.Int) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.UnfundSubscription(&_SubscriptionSenderDapp.TransactOpts, amount)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappTransactorSession) UnfundSubscription(amount *big.Int) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.UnfundSubscription(&_SubscriptionSenderDapp.TransactOpts, amount)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDapp) Address() common.Address {
	return _SubscriptionSenderDapp.address
}

type SubscriptionSenderDappInterface interface {
	IDestinationChainId(opts *bind.CallOpts) (*big.Int, error)

	IOnRampRouter(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	FundSubscription(opts *bind.TransactOpts, feeToken common.Address, amount *big.Int) (*types.Transaction, error)

	SendMessage(opts *bind.TransactOpts, message CCIPEVM2AnySubscriptionMessage) (*types.Transaction, error)

	UnfundSubscription(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Address() common.Address
}
