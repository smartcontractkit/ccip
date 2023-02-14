// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package receiver_dapp

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

type ClientAny2EVMMessage struct {
	MessageId        [32]byte
	SourceChainId    uint64
	Sender           []byte
	Data             []byte
	DestTokenAmounts []ClientEVMTokenAmount
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

var ReceiverDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516108e93803806108e983398101604081905261002f91610070565b806001600160a01b03811661005e576040516335fdcccd60e21b81526000600482015260240160405180910390fd5b6001600160a01b0316608052506100a0565b60006020828403121561008257600080fd5b81516001600160a01b038116811461009957600080fd5b9392505050565b6080516108286100c16000396000818160f101526101cc01526108286000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806301ffc9a714610051578063181f5a771461007957806385572ffb146100c2578063b0f479a1146100d7575b600080fd5b61006461005f3660046103a3565b61011b565b60405190151581526020015b60405180910390f35b6100b56040518060400160405280601281526020017f52656365697665724461707020322e302e30000000000000000000000000000081525081565b60405161007091906103ec565b6100d56100d0366004610441565b6101b4565b005b60405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610070565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f85572ffb0000000000000000000000000000000000000000000000000000000014806101ae57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610229576040517fd7f7333400000000000000000000000000000000000000000000000000000000815233600482015260240160405180910390fd5b61023a61023582610684565b61023d565b50565b61023a816060015182608001516000828060200190518101906102609190610731565b91505060005b825181101561039d5760008382815181106102835761028361076b565b6020026020010151602001519050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156102cd57508015155b1561038c578382815181106102e4576102e461076b565b6020908102919091010151516040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018490529091169063a9059cbb906044016020604051808303816000875af1158015610366573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061038a919061079a565b505b50610396816107bc565b9050610266565b50505050565b6000602082840312156103b557600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146103e557600080fd5b9392505050565b600060208083528351808285015260005b81811015610419578581018301518582016040015282016103fd565b8181111561042b576000604083870101525b50601f01601f1916929092016040019392505050565b60006020828403121561045357600080fd5b813567ffffffffffffffff81111561046a57600080fd5b820160a081850312156103e557600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156104ce576104ce61047c565b60405290565b60405160a0810167ffffffffffffffff811182821017156104ce576104ce61047c565b604051601f8201601f1916810167ffffffffffffffff811182821017156105205761052061047c565b604052919050565b803567ffffffffffffffff8116811461054057600080fd5b919050565b600082601f83011261055657600080fd5b813567ffffffffffffffff8111156105705761057061047c565b6105836020601f19601f840116016104f7565b81815284602083860101111561059857600080fd5b816020850160208301376000918101602001919091529392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461023a57600080fd5b600082601f8301126105e857600080fd5b8135602067ffffffffffffffff8211156106045761060461047c565b610612818360051b016104f7565b82815260069290921b8401810191818101908684111561063157600080fd5b8286015b84811015610679576040818903121561064e5760008081fd5b6106566104ab565b8135610661816105b5565b81528185013585820152835291830191604001610635565b509695505050505050565b600060a0823603121561069657600080fd5b61069e6104d4565b823581526106ae60208401610528565b6020820152604083013567ffffffffffffffff808211156106ce57600080fd5b6106da36838701610545565b604084015260608501359150808211156106f357600080fd5b6106ff36838701610545565b6060840152608085013591508082111561071857600080fd5b50610725368286016105d7565b60808301525092915050565b6000806040838503121561074457600080fd5b825161074f816105b5565b6020840151909250610760816105b5565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156107ac57600080fd5b815180151581146103e557600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610814577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c634300080f000a",
}

var ReceiverDappABI = ReceiverDappMetaData.ABI

var ReceiverDappBin = ReceiverDappMetaData.Bin

func DeployReceiverDapp(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address) (common.Address, *types.Transaction, *ReceiverDapp, error) {
	parsed, err := ReceiverDappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReceiverDappBin), backend, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReceiverDapp{ReceiverDappCaller: ReceiverDappCaller{contract: contract}, ReceiverDappTransactor: ReceiverDappTransactor{contract: contract}, ReceiverDappFilterer: ReceiverDappFilterer{contract: contract}}, nil
}

type ReceiverDapp struct {
	address common.Address
	abi     abi.ABI
	ReceiverDappCaller
	ReceiverDappTransactor
	ReceiverDappFilterer
}

type ReceiverDappCaller struct {
	contract *bind.BoundContract
}

type ReceiverDappTransactor struct {
	contract *bind.BoundContract
}

type ReceiverDappFilterer struct {
	contract *bind.BoundContract
}

type ReceiverDappSession struct {
	Contract     *ReceiverDapp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ReceiverDappCallerSession struct {
	Contract *ReceiverDappCaller
	CallOpts bind.CallOpts
}

type ReceiverDappTransactorSession struct {
	Contract     *ReceiverDappTransactor
	TransactOpts bind.TransactOpts
}

type ReceiverDappRaw struct {
	Contract *ReceiverDapp
}

type ReceiverDappCallerRaw struct {
	Contract *ReceiverDappCaller
}

type ReceiverDappTransactorRaw struct {
	Contract *ReceiverDappTransactor
}

func NewReceiverDapp(address common.Address, backend bind.ContractBackend) (*ReceiverDapp, error) {
	abi, err := abi.JSON(strings.NewReader(ReceiverDappABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindReceiverDapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiverDapp{address: address, abi: abi, ReceiverDappCaller: ReceiverDappCaller{contract: contract}, ReceiverDappTransactor: ReceiverDappTransactor{contract: contract}, ReceiverDappFilterer: ReceiverDappFilterer{contract: contract}}, nil
}

func NewReceiverDappCaller(address common.Address, caller bind.ContractCaller) (*ReceiverDappCaller, error) {
	contract, err := bindReceiverDapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverDappCaller{contract: contract}, nil
}

func NewReceiverDappTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiverDappTransactor, error) {
	contract, err := bindReceiverDapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiverDappTransactor{contract: contract}, nil
}

func NewReceiverDappFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiverDappFilterer, error) {
	contract, err := bindReceiverDapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiverDappFilterer{contract: contract}, nil
}

func bindReceiverDapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReceiverDappABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_ReceiverDapp *ReceiverDappRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiverDapp.Contract.ReceiverDappCaller.contract.Call(opts, result, method, params...)
}

func (_ReceiverDapp *ReceiverDappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.ReceiverDappTransactor.contract.Transfer(opts)
}

func (_ReceiverDapp *ReceiverDappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.ReceiverDappTransactor.contract.Transact(opts, method, params...)
}

func (_ReceiverDapp *ReceiverDappCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiverDapp.Contract.contract.Call(opts, result, method, params...)
}

func (_ReceiverDapp *ReceiverDappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.contract.Transfer(opts)
}

func (_ReceiverDapp *ReceiverDappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.contract.Transact(opts, method, params...)
}

func (_ReceiverDapp *ReceiverDappCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReceiverDapp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ReceiverDapp *ReceiverDappSession) GetRouter() (common.Address, error) {
	return _ReceiverDapp.Contract.GetRouter(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappCallerSession) GetRouter() (common.Address, error) {
	return _ReceiverDapp.Contract.GetRouter(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ReceiverDapp.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ReceiverDapp *ReceiverDappSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ReceiverDapp.Contract.SupportsInterface(&_ReceiverDapp.CallOpts, interfaceId)
}

func (_ReceiverDapp *ReceiverDappCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ReceiverDapp.Contract.SupportsInterface(&_ReceiverDapp.CallOpts, interfaceId)
}

func (_ReceiverDapp *ReceiverDappCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ReceiverDapp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ReceiverDapp *ReceiverDappSession) TypeAndVersion() (string, error) {
	return _ReceiverDapp.Contract.TypeAndVersion(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappCallerSession) TypeAndVersion() (string, error) {
	return _ReceiverDapp.Contract.TypeAndVersion(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _ReceiverDapp.contract.Transact(opts, "ccipReceive", message)
}

func (_ReceiverDapp *ReceiverDappSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.CcipReceive(&_ReceiverDapp.TransactOpts, message)
}

func (_ReceiverDapp *ReceiverDappTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.CcipReceive(&_ReceiverDapp.TransactOpts, message)
}

func (_ReceiverDapp *ReceiverDapp) Address() common.Address {
	return _ReceiverDapp.address
}

type ReceiverDappInterface interface {
	GetRouter(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	Address() common.Address
}
