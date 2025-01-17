// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package maybe_revert_message_receiver

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

type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

var MaybeRevertMessageReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"toRevert\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"CustomError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiveRevert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ValueReceived\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_toRevert\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"setErr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"toRevert\",\"type\":\"bool\"}],\"name\":\"setRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161105938038061105983398101604081905261002f91610049565b336080526000805460ff1916911515919091179055610072565b60006020828403121561005b57600080fd5b8151801515811461006b57600080fd5b9392505050565b608051610f9b6100be60003960008181610286015281816103f8015281816104b301528181610557015281816105f8015281816106b00152818161073a01526108720152610f9b6000f3fe6080604052600436106100745760003560e01c80635100fc211161004e5780635100fc211461015b57806377f5b0e61461017557806385572ffb146101955780638fb5f171146101b557600080fd5b806301ffc9a7146100f057806306b091f91461012457806324600fc31461014657600080fd5b366100eb5760005460ff16156100b6576040517f3085b8db00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040513481527fe12e3b7047ff60a2dd763cf536a43597e5ce7fe7aa7476345bd4cd079912bcef9060200160405180910390a1005b600080fd5b3480156100fc57600080fd5b5061011061010b3660046108fa565b6101d5565b604051901515815260200160405180910390f35b34801561013057600080fd5b5061014461013f36600461096c565b61026e565b005b34801561015257600080fd5b5061014461053f565b34801561016757600080fd5b506000546101109060ff1681565b34801561018157600080fd5b506101446101903660046109c5565b610722565b3480156101a157600080fd5b506101446101b0366004610a94565b6107a1565b3480156101c157600080fd5b506101446101d0366004610ae0565b61085a565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f85572ffb00000000000000000000000000000000000000000000000000000000148061026857507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146102dd576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152829060009073ffffffffffffffffffffffffffffffffffffffff8316906370a0823190602401602060405180830381865afa15801561034c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103709190610afd565b9050828110156103bb576040517fcf47918100000000000000000000000000000000000000000000000000000000815260048101829052602481018490526044015b60405180910390fd5b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081166004830152602482018590526000919084169063a9059cbb906044016020604051808303816000875af1158015610454573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104789190610b16565b9050806104b1576040517f90b8ec1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f6337ed398c0e8467698c581374fdce4db14922df487b5a39483079f5f59b60a48660405161053091815260200190565b60405180910390a35050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146105ae576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4760008190036105f4576040517fcf47918100000000000000000000000000000000000000000000000000000000815260006004820152600160248201526044016103b2565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d806000811461066e576040519150601f19603f3d011682016040523d82523d6000602084013e610673565b606091505b50509050806106ae576040517f90b8ec1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167feaff4b37086828766ad3268786972c0cd24259d4c87a80f9d3963a3c3d999b0d8360405161071691815260200190565b60405180910390a25050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610791576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600161079d8282610bd7565b5050565b60005460ff16156107e15760016040517f5a4ff6710000000000000000000000000000000000000000000000000000000081526004016103b29190610cf1565b7f707732b700184c0ab3b799f43f03de9b3606a144cfb367f98291044e71972cdc81356108146040840160208501610d9f565b6108216040850185610dc9565b61082e6060870187610dc9565b61083b6080890189610e35565b60405161084f989796959493929190610ee6565b60405180910390a150565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146108c9576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b60006020828403121561090c57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461093c57600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461096757600080fd5b919050565b6000806040838503121561097f57600080fd5b61098883610943565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156109d757600080fd5b813567ffffffffffffffff808211156109ef57600080fd5b818401915084601f830112610a0357600080fd5b813581811115610a1557610a15610996565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610a5b57610a5b610996565b81604052828152876020848701011115610a7457600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208284031215610aa657600080fd5b813567ffffffffffffffff811115610abd57600080fd5b820160a0818503121561093c57600080fd5b8015158114610add57600080fd5b50565b600060208284031215610af257600080fd5b813561093c81610acf565b600060208284031215610b0f57600080fd5b5051919050565b600060208284031215610b2857600080fd5b815161093c81610acf565b600181811c90821680610b4757607f821691505b602082108103610b80577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f821115610bd2576000816000526020600020601f850160051c81016020861015610baf5750805b601f850160051c820191505b81811015610bce57828155600101610bbb565b5050505b505050565b815167ffffffffffffffff811115610bf157610bf1610996565b610c0581610bff8454610b33565b84610b86565b602080601f831160018114610c585760008415610c225750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555610bce565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015610ca557888601518255948401946001909101908401610c86565b5085821015610ce157878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b6000602080835260008454610d0581610b33565b8060208701526040600180841660008114610d275760018114610d6157610d91565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00851660408a0152604084151560051b8a01019550610d91565b89600052602060002060005b85811015610d885781548b8201860152908301908801610d6d565b8a016040019650505b509398975050505050505050565b600060208284031215610db157600080fd5b813567ffffffffffffffff8116811461093c57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610dfe57600080fd5b83018035915067ffffffffffffffff821115610e1957600080fd5b602001915036819003821315610e2e57600080fd5b9250929050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610e6a57600080fd5b83018035915067ffffffffffffffff821115610e8557600080fd5b6020019150600681901b3603821315610e2e57600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8881526000602067ffffffffffffffff8a166020840152604060a06040850152610f1460a085018a8c610e9d565b8481036060860152610f2781898b610e9d565b858103608087015286815287915060200160005b87811015610f7c5773ffffffffffffffffffffffffffffffffffffffff610f6184610943565b16825282850135858301529183019190830190600101610f3b565b509d9c5050505050505050505050505056fea164736f6c6343000818000a",
}

var MaybeRevertMessageReceiverABI = MaybeRevertMessageReceiverMetaData.ABI

var MaybeRevertMessageReceiverBin = MaybeRevertMessageReceiverMetaData.Bin

func DeployMaybeRevertMessageReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, toRevert bool) (common.Address, *types.Transaction, *MaybeRevertMessageReceiver, error) {
	parsed, err := MaybeRevertMessageReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MaybeRevertMessageReceiverBin), backend, toRevert)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MaybeRevertMessageReceiver{address: address, abi: *parsed, MaybeRevertMessageReceiverCaller: MaybeRevertMessageReceiverCaller{contract: contract}, MaybeRevertMessageReceiverTransactor: MaybeRevertMessageReceiverTransactor{contract: contract}, MaybeRevertMessageReceiverFilterer: MaybeRevertMessageReceiverFilterer{contract: contract}}, nil
}

type MaybeRevertMessageReceiver struct {
	address common.Address
	abi     abi.ABI
	MaybeRevertMessageReceiverCaller
	MaybeRevertMessageReceiverTransactor
	MaybeRevertMessageReceiverFilterer
}

type MaybeRevertMessageReceiverCaller struct {
	contract *bind.BoundContract
}

type MaybeRevertMessageReceiverTransactor struct {
	contract *bind.BoundContract
}

type MaybeRevertMessageReceiverFilterer struct {
	contract *bind.BoundContract
}

type MaybeRevertMessageReceiverSession struct {
	Contract     *MaybeRevertMessageReceiver
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MaybeRevertMessageReceiverCallerSession struct {
	Contract *MaybeRevertMessageReceiverCaller
	CallOpts bind.CallOpts
}

type MaybeRevertMessageReceiverTransactorSession struct {
	Contract     *MaybeRevertMessageReceiverTransactor
	TransactOpts bind.TransactOpts
}

type MaybeRevertMessageReceiverRaw struct {
	Contract *MaybeRevertMessageReceiver
}

type MaybeRevertMessageReceiverCallerRaw struct {
	Contract *MaybeRevertMessageReceiverCaller
}

type MaybeRevertMessageReceiverTransactorRaw struct {
	Contract *MaybeRevertMessageReceiverTransactor
}

func NewMaybeRevertMessageReceiver(address common.Address, backend bind.ContractBackend) (*MaybeRevertMessageReceiver, error) {
	abi, err := abi.JSON(strings.NewReader(MaybeRevertMessageReceiverABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMaybeRevertMessageReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiver{address: address, abi: abi, MaybeRevertMessageReceiverCaller: MaybeRevertMessageReceiverCaller{contract: contract}, MaybeRevertMessageReceiverTransactor: MaybeRevertMessageReceiverTransactor{contract: contract}, MaybeRevertMessageReceiverFilterer: MaybeRevertMessageReceiverFilterer{contract: contract}}, nil
}

func NewMaybeRevertMessageReceiverCaller(address common.Address, caller bind.ContractCaller) (*MaybeRevertMessageReceiverCaller, error) {
	contract, err := bindMaybeRevertMessageReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverCaller{contract: contract}, nil
}

func NewMaybeRevertMessageReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*MaybeRevertMessageReceiverTransactor, error) {
	contract, err := bindMaybeRevertMessageReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverTransactor{contract: contract}, nil
}

func NewMaybeRevertMessageReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*MaybeRevertMessageReceiverFilterer, error) {
	contract, err := bindMaybeRevertMessageReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverFilterer{contract: contract}, nil
}

func bindMaybeRevertMessageReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MaybeRevertMessageReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MaybeRevertMessageReceiver.Contract.MaybeRevertMessageReceiverCaller.contract.Call(opts, result, method, params...)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.MaybeRevertMessageReceiverTransactor.contract.Transfer(opts)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.MaybeRevertMessageReceiverTransactor.contract.Transact(opts, method, params...)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MaybeRevertMessageReceiver.Contract.contract.Call(opts, result, method, params...)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.contract.Transfer(opts)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.contract.Transact(opts, method, params...)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverCaller) SToRevert(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MaybeRevertMessageReceiver.contract.Call(opts, &out, "s_toRevert")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) SToRevert() (bool, error) {
	return _MaybeRevertMessageReceiver.Contract.SToRevert(&_MaybeRevertMessageReceiver.CallOpts)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverCallerSession) SToRevert() (bool, error) {
	return _MaybeRevertMessageReceiver.Contract.SToRevert(&_MaybeRevertMessageReceiver.CallOpts)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MaybeRevertMessageReceiver.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MaybeRevertMessageReceiver.Contract.SupportsInterface(&_MaybeRevertMessageReceiver.CallOpts, interfaceId)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MaybeRevertMessageReceiver.Contract.SupportsInterface(&_MaybeRevertMessageReceiver.CallOpts, interfaceId)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.contract.Transact(opts, "ccipReceive", message)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.CcipReceive(&_MaybeRevertMessageReceiver.TransactOpts, message)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.CcipReceive(&_MaybeRevertMessageReceiver.TransactOpts, message)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactor) SetErr(opts *bind.TransactOpts, err []byte) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.contract.Transact(opts, "setErr", err)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) SetErr(err []byte) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.SetErr(&_MaybeRevertMessageReceiver.TransactOpts, err)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorSession) SetErr(err []byte) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.SetErr(&_MaybeRevertMessageReceiver.TransactOpts, err)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactor) SetRevert(opts *bind.TransactOpts, toRevert bool) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.contract.Transact(opts, "setRevert", toRevert)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) SetRevert(toRevert bool) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.SetRevert(&_MaybeRevertMessageReceiver.TransactOpts, toRevert)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorSession) SetRevert(toRevert bool) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.SetRevert(&_MaybeRevertMessageReceiver.TransactOpts, toRevert)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactor) WithdrawFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.contract.Transact(opts, "withdrawFunds")
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) WithdrawFunds() (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.WithdrawFunds(&_MaybeRevertMessageReceiver.TransactOpts)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorSession) WithdrawFunds() (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.WithdrawFunds(&_MaybeRevertMessageReceiver.TransactOpts)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactor) WithdrawTokens(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.contract.Transact(opts, "withdrawTokens", token, amount)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) WithdrawTokens(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.WithdrawTokens(&_MaybeRevertMessageReceiver.TransactOpts, token, amount)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorSession) WithdrawTokens(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.WithdrawTokens(&_MaybeRevertMessageReceiver.TransactOpts, token, amount)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.contract.RawTransact(opts, nil)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverSession) Receive() (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.Receive(&_MaybeRevertMessageReceiver.TransactOpts)
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverTransactorSession) Receive() (*types.Transaction, error) {
	return _MaybeRevertMessageReceiver.Contract.Receive(&_MaybeRevertMessageReceiver.TransactOpts)
}

type MaybeRevertMessageReceiverFundsWithdrawnIterator struct {
	Event *MaybeRevertMessageReceiverFundsWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MaybeRevertMessageReceiverFundsWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaybeRevertMessageReceiverFundsWithdrawn)
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
		it.Event = new(MaybeRevertMessageReceiverFundsWithdrawn)
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

func (it *MaybeRevertMessageReceiverFundsWithdrawnIterator) Error() error {
	return it.fail
}

func (it *MaybeRevertMessageReceiverFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MaybeRevertMessageReceiverFundsWithdrawn struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*MaybeRevertMessageReceiverFundsWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MaybeRevertMessageReceiver.contract.FilterLogs(opts, "FundsWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverFundsWithdrawnIterator{contract: _MaybeRevertMessageReceiver.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *MaybeRevertMessageReceiverFundsWithdrawn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MaybeRevertMessageReceiver.contract.WatchLogs(opts, "FundsWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MaybeRevertMessageReceiverFundsWithdrawn)
				if err := _MaybeRevertMessageReceiver.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
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

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) ParseFundsWithdrawn(log types.Log) (*MaybeRevertMessageReceiverFundsWithdrawn, error) {
	event := new(MaybeRevertMessageReceiverFundsWithdrawn)
	if err := _MaybeRevertMessageReceiver.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MaybeRevertMessageReceiverMessageReceivedIterator struct {
	Event *MaybeRevertMessageReceiverMessageReceived

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MaybeRevertMessageReceiverMessageReceivedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaybeRevertMessageReceiverMessageReceived)
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
		it.Event = new(MaybeRevertMessageReceiverMessageReceived)
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

func (it *MaybeRevertMessageReceiverMessageReceivedIterator) Error() error {
	return it.fail
}

func (it *MaybeRevertMessageReceiverMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MaybeRevertMessageReceiverMessageReceived struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
	Raw                 types.Log
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*MaybeRevertMessageReceiverMessageReceivedIterator, error) {

	logs, sub, err := _MaybeRevertMessageReceiver.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverMessageReceivedIterator{contract: _MaybeRevertMessageReceiver.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *MaybeRevertMessageReceiverMessageReceived) (event.Subscription, error) {

	logs, sub, err := _MaybeRevertMessageReceiver.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MaybeRevertMessageReceiverMessageReceived)
				if err := _MaybeRevertMessageReceiver.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) ParseMessageReceived(log types.Log) (*MaybeRevertMessageReceiverMessageReceived, error) {
	event := new(MaybeRevertMessageReceiverMessageReceived)
	if err := _MaybeRevertMessageReceiver.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MaybeRevertMessageReceiverTokensWithdrawnIterator struct {
	Event *MaybeRevertMessageReceiverTokensWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MaybeRevertMessageReceiverTokensWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaybeRevertMessageReceiverTokensWithdrawn)
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
		it.Event = new(MaybeRevertMessageReceiverTokensWithdrawn)
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

func (it *MaybeRevertMessageReceiverTokensWithdrawnIterator) Error() error {
	return it.fail
}

func (it *MaybeRevertMessageReceiverTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MaybeRevertMessageReceiverTokensWithdrawn struct {
	Token  common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, token []common.Address, owner []common.Address) (*MaybeRevertMessageReceiverTokensWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MaybeRevertMessageReceiver.contract.FilterLogs(opts, "TokensWithdrawn", tokenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverTokensWithdrawnIterator{contract: _MaybeRevertMessageReceiver.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *MaybeRevertMessageReceiverTokensWithdrawn, token []common.Address, owner []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MaybeRevertMessageReceiver.contract.WatchLogs(opts, "TokensWithdrawn", tokenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MaybeRevertMessageReceiverTokensWithdrawn)
				if err := _MaybeRevertMessageReceiver.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) ParseTokensWithdrawn(log types.Log) (*MaybeRevertMessageReceiverTokensWithdrawn, error) {
	event := new(MaybeRevertMessageReceiverTokensWithdrawn)
	if err := _MaybeRevertMessageReceiver.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MaybeRevertMessageReceiverValueReceivedIterator struct {
	Event *MaybeRevertMessageReceiverValueReceived

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MaybeRevertMessageReceiverValueReceivedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaybeRevertMessageReceiverValueReceived)
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
		it.Event = new(MaybeRevertMessageReceiverValueReceived)
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

func (it *MaybeRevertMessageReceiverValueReceivedIterator) Error() error {
	return it.fail
}

func (it *MaybeRevertMessageReceiverValueReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MaybeRevertMessageReceiverValueReceived struct {
	Amount *big.Int
	Raw    types.Log
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) FilterValueReceived(opts *bind.FilterOpts) (*MaybeRevertMessageReceiverValueReceivedIterator, error) {

	logs, sub, err := _MaybeRevertMessageReceiver.contract.FilterLogs(opts, "ValueReceived")
	if err != nil {
		return nil, err
	}
	return &MaybeRevertMessageReceiverValueReceivedIterator{contract: _MaybeRevertMessageReceiver.contract, event: "ValueReceived", logs: logs, sub: sub}, nil
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) WatchValueReceived(opts *bind.WatchOpts, sink chan<- *MaybeRevertMessageReceiverValueReceived) (event.Subscription, error) {

	logs, sub, err := _MaybeRevertMessageReceiver.contract.WatchLogs(opts, "ValueReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MaybeRevertMessageReceiverValueReceived)
				if err := _MaybeRevertMessageReceiver.contract.UnpackLog(event, "ValueReceived", log); err != nil {
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

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiverFilterer) ParseValueReceived(log types.Log) (*MaybeRevertMessageReceiverValueReceived, error) {
	event := new(MaybeRevertMessageReceiverValueReceived)
	if err := _MaybeRevertMessageReceiver.contract.UnpackLog(event, "ValueReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiver) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MaybeRevertMessageReceiver.abi.Events["FundsWithdrawn"].ID:
		return _MaybeRevertMessageReceiver.ParseFundsWithdrawn(log)
	case _MaybeRevertMessageReceiver.abi.Events["MessageReceived"].ID:
		return _MaybeRevertMessageReceiver.ParseMessageReceived(log)
	case _MaybeRevertMessageReceiver.abi.Events["TokensWithdrawn"].ID:
		return _MaybeRevertMessageReceiver.ParseTokensWithdrawn(log)
	case _MaybeRevertMessageReceiver.abi.Events["ValueReceived"].ID:
		return _MaybeRevertMessageReceiver.ParseValueReceived(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MaybeRevertMessageReceiverFundsWithdrawn) Topic() common.Hash {
	return common.HexToHash("0xeaff4b37086828766ad3268786972c0cd24259d4c87a80f9d3963a3c3d999b0d")
}

func (MaybeRevertMessageReceiverMessageReceived) Topic() common.Hash {
	return common.HexToHash("0x707732b700184c0ab3b799f43f03de9b3606a144cfb367f98291044e71972cdc")
}

func (MaybeRevertMessageReceiverTokensWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x6337ed398c0e8467698c581374fdce4db14922df487b5a39483079f5f59b60a4")
}

func (MaybeRevertMessageReceiverValueReceived) Topic() common.Hash {
	return common.HexToHash("0xe12e3b7047ff60a2dd763cf536a43597e5ce7fe7aa7476345bd4cd079912bcef")
}

func (_MaybeRevertMessageReceiver *MaybeRevertMessageReceiver) Address() common.Address {
	return _MaybeRevertMessageReceiver.address
}

type MaybeRevertMessageReceiverInterface interface {
	SToRevert(opts *bind.CallOpts) (bool, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	SetErr(opts *bind.TransactOpts, err []byte) (*types.Transaction, error)

	SetRevert(opts *bind.TransactOpts, toRevert bool) (*types.Transaction, error)

	WithdrawFunds(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawTokens(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterFundsWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*MaybeRevertMessageReceiverFundsWithdrawnIterator, error)

	WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *MaybeRevertMessageReceiverFundsWithdrawn, owner []common.Address) (event.Subscription, error)

	ParseFundsWithdrawn(log types.Log) (*MaybeRevertMessageReceiverFundsWithdrawn, error)

	FilterMessageReceived(opts *bind.FilterOpts) (*MaybeRevertMessageReceiverMessageReceivedIterator, error)

	WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *MaybeRevertMessageReceiverMessageReceived) (event.Subscription, error)

	ParseMessageReceived(log types.Log) (*MaybeRevertMessageReceiverMessageReceived, error)

	FilterTokensWithdrawn(opts *bind.FilterOpts, token []common.Address, owner []common.Address) (*MaybeRevertMessageReceiverTokensWithdrawnIterator, error)

	WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *MaybeRevertMessageReceiverTokensWithdrawn, token []common.Address, owner []common.Address) (event.Subscription, error)

	ParseTokensWithdrawn(log types.Log) (*MaybeRevertMessageReceiverTokensWithdrawn, error)

	FilterValueReceived(opts *bind.FilterOpts) (*MaybeRevertMessageReceiverValueReceivedIterator, error)

	WatchValueReceived(opts *bind.WatchOpts, sink chan<- *MaybeRevertMessageReceiverValueReceived) (event.Subscription, error)

	ParseValueReceived(log types.Log) (*MaybeRevertMessageReceiverValueReceived, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
