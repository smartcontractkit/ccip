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
	Receiver  []byte
	Data      []byte
	Tokens    []common.Address
	Amounts   []*big.Int
	ExtraArgs []byte
}

var SubscriptionSenderDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractEVM2AnySubscriptionOnRampRouterInterface\",\"name\":\"onRampRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_destinationChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_onRampRouter\",\"outputs\":[{\"internalType\":\"contractEVM2AnySubscriptionOnRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.EVM2AnySubscriptionMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"sendMessage\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unfundSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b506040516110b63803806110b683398101604081905261002f91610045565b6001600160a01b0390911660805260a05261007f565b6000806040838503121561005857600080fd5b82516001600160a01b038116811461006f57600080fd5b6020939093015192949293505050565b60805160a051610fe86100ce6000396000818161016e01526104e101526000818161010f015281816101db01528181610287015281816103ab015281816104b401526105850152610fe86000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c806389f9ad2a1161005057806389f9ad2a1461010a57806395e712db14610156578063a72171951461016957600080fd5b8063181f5a77146100775780633c5457ce146100c957806371858313146100de575b600080fd5b6100b36040518060400160405280601c81526020017f537562736372697074696f6e53656e6465724461707020312e302e300000000081525081565b6040516100c09190610a0c565b60405180910390f35b6100dc6100d7366004610a44565b61019e565b005b6100f16100ec366004610c98565b6102fc565b60405167ffffffffffffffff90911681526020016100c0565b6101317f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100c0565b6100dc610164366004610d95565b610556565b6101907f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100c0565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660048301526024820183905283169063095ea7b3906044016020604051808303816000875af1158015610233573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102579190610dae565b506040517fc1060653000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063c106065390602401600060405180830381600087803b1580156102e057600080fd5b505af11580156102f4573d6000803e3d6000fd5b505050505050565b6040810151606082015160009190825b825181101561047657610374333084848151811061032c5761032c610dd0565b602002602001015186858151811061034657610346610dd0565b602002602001015173ffffffffffffffffffffffffffffffffffffffff166105f9909392919063ffffffff16565b82818151811061038657610386610dd0565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008484815181106103dc576103dc610dd0565b60200260200101516040518363ffffffff1660e01b815260040161042292919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af1158015610441573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104659190610dae565b5061046f81610dff565b905061030c565b506040517fcc4a9c6800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063cc4a9c689061050b907f0000000000000000000000000000000000000000000000000000000000000000908890600401610e99565b6020604051808303816000875af115801561052a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061054e9190610f95565b949350505050565b6040517f95e712db000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906395e712db90602401600060405180830381600087803b1580156105de57600080fd5b505af11580156105f2573d6000803e3d6000fd5b5050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261068e908590610694565b50505050565b60006106f6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166107aa9092919063ffffffff16565b8051909150156107a557808060200190518101906107149190610dae565b6107a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b505050565b60606107b984846000856107c3565b90505b9392505050565b606082471015610855576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161079c565b843b6108bd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161079c565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516108e69190610fbf565b60006040518083038185875af1925050503d8060008114610923576040519150601f19603f3d011682016040523d82523d6000602084013e610928565b606091505b5091509150610938828286610943565b979650505050505050565b606083156109525750816107bc565b8251156109625782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079c9190610a0c565b60005b838110156109b1578181015183820152602001610999565b8381111561068e5750506000910152565b600081518084526109da816020860160208601610996565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006107bc60208301846109c2565b73ffffffffffffffffffffffffffffffffffffffff81168114610a4157600080fd5b50565b60008060408385031215610a5757600080fd5b8235610a6281610a1f565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff81118282101715610ac257610ac2610a70565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610b0f57610b0f610a70565b604052919050565b600082601f830112610b2857600080fd5b813567ffffffffffffffff811115610b4257610b42610a70565b610b7360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610ac8565b818152846020838601011115610b8857600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115610bbf57610bbf610a70565b5060051b60200190565b600082601f830112610bda57600080fd5b81356020610bef610bea83610ba5565b610ac8565b82815260059290921b84018101918181019086841115610c0e57600080fd5b8286015b84811015610c32578035610c2581610a1f565b8352918301918301610c12565b509695505050505050565b600082601f830112610c4e57600080fd5b81356020610c5e610bea83610ba5565b82815260059290921b84018101918181019086841115610c7d57600080fd5b8286015b84811015610c325780358352918301918301610c81565b600060208284031215610caa57600080fd5b813567ffffffffffffffff80821115610cc257600080fd5b9083019060a08286031215610cd657600080fd5b610cde610a9f565b823582811115610ced57600080fd5b610cf987828601610b17565b825250602083013582811115610d0e57600080fd5b610d1a87828601610b17565b602083015250604083013582811115610d3257600080fd5b610d3e87828601610bc9565b604083015250606083013582811115610d5657600080fd5b610d6287828601610c3d565b606083015250608083013582811115610d7a57600080fd5b610d8687828601610b17565b60808301525095945050505050565b600060208284031215610da757600080fd5b5035919050565b600060208284031215610dc057600080fd5b815180151581146107bc57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610e57577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b600081518084526020808501945080840160005b83811015610e8e57815187529582019590820190600101610e72565b509495945050505050565b82815260006020604081840152835160a06040850152610ebc60e08501826109c2565b9050818501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080868403016060870152610ef783836109c2565b6040880151878203830160808901528051808352908601945060009350908501905b80841015610f4f57845173ffffffffffffffffffffffffffffffffffffffff168252938501936001939093019290850190610f19565b5060608801519450818782030160a0880152610f6b8186610e5e565b94505060808701519250808685030160c08701525050610f8b82826109c2565b9695505050505050565b600060208284031215610fa757600080fd5b815167ffffffffffffffff811681146107bc57600080fd5b60008251610fd1818460208701610996565b919091019291505056fea164736f6c634300080f000a",
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
