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

type CCIPMessage struct {
	SequenceNumber *big.Int
	SourceChainId  *big.Int
	Sender         common.Address
	Payload        CCIPMessagePayload
}

type CCIPMessagePayload struct {
	Tokens             []common.Address
	Amounts            []*big.Int
	DestinationChainId *big.Int
	Receiver           common.Address
	Executor           common.Address
	Data               []byte
	Options            []byte
}

var ReceiverDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractOffRamp\",\"name\":\"offRamp\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"OFF_RAMP\",\"outputs\":[{\"internalType\":\"contractOffRamp\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"receiveMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b506040516106c43803806106c483398101604081905261002f9161005e565b6001600160a01b039182166080521660a052610098565b6001600160a01b038116811461005b57600080fd5b50565b6000806040838503121561007157600080fd5b825161007c81610046565b602084015190925061008d81610046565b809150509250929050565b60805160a0516105fb6100c96000396000818160ea0152610293015260008181609e015261013901526105fb6000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063181f5a771461005157806333219c691461009957806382bfefc8146100e55780639c5984681461010c575b600080fd5b604080518082018252601281527f52656365697665724461707020302e302e31000000000000000000000000000060208201529051610090919061031a565b60405180910390f35b6100c07f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610090565b6100c07f000000000000000000000000000000000000000000000000000000000000000081565b61011f61011a36600461038d565b610121565b005b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610196576040517f0af9f1b600000000000000000000000000000000000000000000000000000000815233600482015260240160405180910390fd5b60006101a560608301836103cf565b6101b39060a081019061040d565b8101906101c091906104a2565b91505060005b6101d360608401846103cf565b6101dd90806104d5565b90508110156103155760006101f560608501856103cf565b6102039060208101906104d5565b838181106102135761021361053d565b602002919091013591505073ffffffffffffffffffffffffffffffffffffffff83161580159061024257508015155b15610302576040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018390527f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303816000875af11580156102dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610300919061056c565b505b508061030d8161058e565b9150506101c6565b505050565b600060208083528351808285015260005b818110156103475785810183015185820160400152820161032b565b81811115610359576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60006020828403121561039f57600080fd5b813567ffffffffffffffff8111156103b657600080fd5b8201608081850312156103c857600080fd5b9392505050565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2183360301811261040357600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261044257600080fd5b83018035915067ffffffffffffffff82111561045d57600080fd5b60200191503681900382131561047257600080fd5b9250929050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461049d57600080fd5b919050565b600080604083850312156104b557600080fd5b6104be83610479565b91506104cc60208401610479565b90509250929050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261050a57600080fd5b83018035915067ffffffffffffffff82111561052557600080fd5b6020019150600581901b360382131561047257600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561057e57600080fd5b815180151581146103c857600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156105e7577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c634300080c000a",
}

var ReceiverDappABI = ReceiverDappMetaData.ABI

var ReceiverDappBin = ReceiverDappMetaData.Bin

func DeployReceiverDapp(auth *bind.TransactOpts, backend bind.ContractBackend, offRamp common.Address, token common.Address) (common.Address, *types.Transaction, *ReceiverDapp, error) {
	parsed, err := ReceiverDappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReceiverDappBin), backend, offRamp, token)
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

func (_ReceiverDapp *ReceiverDappCaller) OFFRAMP(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReceiverDapp.contract.Call(opts, &out, "OFF_RAMP")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ReceiverDapp *ReceiverDappSession) OFFRAMP() (common.Address, error) {
	return _ReceiverDapp.Contract.OFFRAMP(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappCallerSession) OFFRAMP() (common.Address, error) {
	return _ReceiverDapp.Contract.OFFRAMP(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappCaller) TOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReceiverDapp.contract.Call(opts, &out, "TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ReceiverDapp *ReceiverDappSession) TOKEN() (common.Address, error) {
	return _ReceiverDapp.Contract.TOKEN(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappCallerSession) TOKEN() (common.Address, error) {
	return _ReceiverDapp.Contract.TOKEN(&_ReceiverDapp.CallOpts)
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

func (_ReceiverDapp *ReceiverDappTransactor) ReceiveMessage(opts *bind.TransactOpts, message CCIPMessage) (*types.Transaction, error) {
	return _ReceiverDapp.contract.Transact(opts, "receiveMessage", message)
}

func (_ReceiverDapp *ReceiverDappSession) ReceiveMessage(message CCIPMessage) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.ReceiveMessage(&_ReceiverDapp.TransactOpts, message)
}

func (_ReceiverDapp *ReceiverDappTransactorSession) ReceiveMessage(message CCIPMessage) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.ReceiveMessage(&_ReceiverDapp.TransactOpts, message)
}

func (_ReceiverDapp *ReceiverDapp) Address() common.Address {
	return _ReceiverDapp.address
}

type ReceiverDappInterface interface {
	OFFRAMP(opts *bind.CallOpts) (common.Address, error)

	TOKEN(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	ReceiveMessage(opts *bind.TransactOpts, message CCIPMessage) (*types.Transaction, error)

	Address() common.Address
}
