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

type CCIPAny2EVMSubscriptionMessage struct {
	SourceChainId  *big.Int
	SequenceNumber uint64
	Sender         common.Address
	Receiver       common.Address
	Nonce          uint64
	Data           []byte
	Tokens         []common.Address
	Amounts        []*big.Int
	GasLimit       *big.Int
}

type CCIPAny2EVMTollMessage struct {
	SourceChainId  *big.Int
	SequenceNumber uint64
	Sender         common.Address
	Receiver       common.Address
	Data           []byte
	Tokens         []common.Address
	Amounts        []*big.Int
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	GasLimit       *big.Int
}

var ReceiverDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractAny2EVMTollOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deliverer\",\"type\":\"address\"}],\"name\":\"InvalidDeliverer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"contractAny2EVMTollOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMSubscriptionMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSubscriptionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b506040516108de3803806108de83398101604081905261002f91610070565b6001600160a01b039182166080521660a052600080546001600160a01b031916331790556100aa565b6001600160a01b038116811461006d57600080fd5b50565b6000806040838503121561008357600080fd5b825161008e81610058565b602084015190925061009f81610058565b809150509250929050565b60805160a0516107fb6100e36000396000818161011a015261049501526000818160ce0152818161019a01526102db01526107fb6000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c80638bbad066116100505780638bbad0661461013c578063a988980814610151578063e2a92e281461016457600080fd5b8063181f5a771461007757806332fe7b26146100c957806382bfefc814610115575b600080fd5b6100b36040518060400160405280601281526020017f52656365697665724461707020312e302e30000000000000000000000000000081525081565b6040516100c0919061051c565b60405180910390f35b6100f07f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100c0565b6100f07f000000000000000000000000000000000000000000000000000000000000000081565b61014f61014a36600461058f565b610182565b005b61014f61015f3660046105d2565b6102c3565b60005473ffffffffffffffffffffffffffffffffffffffff166100f0565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146101f8576040517f0af9f1b60000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6102c0610208608083018361060e565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061024a9250505060a084018461067a565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506102899250505060c085018561067a565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506103c592505050565b50565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610334576040517f0af9f1b60000000000000000000000000000000000000000000000000000000081523360048201526024016101ef565b6102c061034460a083018361060e565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506103869250505060c084018461067a565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506102899250505060e085018561067a565b6000838060200190518101906103db919061070b565b91505060005b83518110156105155760008382815181106103fe576103fe61073e565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561044457508015155b15610504576040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018390527f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303816000875af11580156104de573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610502919061076d565b505b5061050e8161078f565b90506103e1565b5050505050565b600060208083528351808285015260005b818110156105495785810183015185820160400152820161052d565b8181111561055b576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b6000602082840312156105a157600080fd5b813567ffffffffffffffff8111156105b857600080fd5b820161014081850312156105cb57600080fd5b9392505050565b6000602082840312156105e457600080fd5b813567ffffffffffffffff8111156105fb57600080fd5b820161012081850312156105cb57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261064357600080fd5b83018035915067ffffffffffffffff82111561065e57600080fd5b60200191503681900382131561067357600080fd5b9250929050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126106af57600080fd5b83018035915067ffffffffffffffff8211156106ca57600080fd5b6020019150600581901b360382131561067357600080fd5b805173ffffffffffffffffffffffffffffffffffffffff8116811461070657600080fd5b919050565b6000806040838503121561071e57600080fd5b610727836106e2565b9150610735602084016106e2565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561077f57600080fd5b815180151581146105cb57600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036107e7577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c634300080f000a",
}

var ReceiverDappABI = ReceiverDappMetaData.ABI

var ReceiverDappBin = ReceiverDappMetaData.Bin

func DeployReceiverDapp(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address, token common.Address) (common.Address, *types.Transaction, *ReceiverDapp, error) {
	parsed, err := ReceiverDappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReceiverDappBin), backend, router, token)
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

func (_ReceiverDapp *ReceiverDappCaller) ROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReceiverDapp.contract.Call(opts, &out, "ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ReceiverDapp *ReceiverDappSession) ROUTER() (common.Address, error) {
	return _ReceiverDapp.Contract.ROUTER(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappCallerSession) ROUTER() (common.Address, error) {
	return _ReceiverDapp.Contract.ROUTER(&_ReceiverDapp.CallOpts)
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

func (_ReceiverDapp *ReceiverDappCaller) GetSubscriptionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReceiverDapp.contract.Call(opts, &out, "getSubscriptionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ReceiverDapp *ReceiverDappSession) GetSubscriptionManager() (common.Address, error) {
	return _ReceiverDapp.Contract.GetSubscriptionManager(&_ReceiverDapp.CallOpts)
}

func (_ReceiverDapp *ReceiverDappCallerSession) GetSubscriptionManager() (common.Address, error) {
	return _ReceiverDapp.Contract.GetSubscriptionManager(&_ReceiverDapp.CallOpts)
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

func (_ReceiverDapp *ReceiverDappTransactor) CcipReceive(opts *bind.TransactOpts, message CCIPAny2EVMTollMessage) (*types.Transaction, error) {
	return _ReceiverDapp.contract.Transact(opts, "ccipReceive", message)
}

func (_ReceiverDapp *ReceiverDappSession) CcipReceive(message CCIPAny2EVMTollMessage) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.CcipReceive(&_ReceiverDapp.TransactOpts, message)
}

func (_ReceiverDapp *ReceiverDappTransactorSession) CcipReceive(message CCIPAny2EVMTollMessage) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.CcipReceive(&_ReceiverDapp.TransactOpts, message)
}

func (_ReceiverDapp *ReceiverDappTransactor) CcipReceive0(opts *bind.TransactOpts, message CCIPAny2EVMSubscriptionMessage) (*types.Transaction, error) {
	return _ReceiverDapp.contract.Transact(opts, "ccipReceive0", message)
}

func (_ReceiverDapp *ReceiverDappSession) CcipReceive0(message CCIPAny2EVMSubscriptionMessage) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.CcipReceive0(&_ReceiverDapp.TransactOpts, message)
}

func (_ReceiverDapp *ReceiverDappTransactorSession) CcipReceive0(message CCIPAny2EVMSubscriptionMessage) (*types.Transaction, error) {
	return _ReceiverDapp.Contract.CcipReceive0(&_ReceiverDapp.TransactOpts, message)
}

func (_ReceiverDapp *ReceiverDapp) Address() common.Address {
	return _ReceiverDapp.address
}

type ReceiverDappInterface interface {
	ROUTER(opts *bind.CallOpts) (common.Address, error)

	TOKEN(opts *bind.CallOpts) (common.Address, error)

	GetSubscriptionManager(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	CcipReceive(opts *bind.TransactOpts, message CCIPAny2EVMTollMessage) (*types.Transaction, error)

	CcipReceive0(opts *bind.TransactOpts, message CCIPAny2EVMSubscriptionMessage) (*types.Transaction, error)

	Address() common.Address
}
