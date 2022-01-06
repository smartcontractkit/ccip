// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package single_token_receiver

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
	SequenceNumber     *big.Int
	SourceChainId      *big.Int
	DestinationChainId *big.Int
	Sender             common.Address
	Payload            CCIPMessagePayload
}

type CCIPMessagePayload struct {
	Receiver common.Address
	Data     []byte
	Tokens   []common.Address
	Amounts  []*big.Int
	Executor common.Address
	Options  []byte
}

var EOASingleTokenReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractSingleTokenOffRamp\",\"name\":\"offRamp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"OFF_RAMP\",\"outputs\":[{\"internalType\":\"contractSingleTokenOffRamp\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"internalType\":\"structCCIP.Message\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"receiveMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516106c53803806106c583398101604081905261002f91610044565b60601b6001600160601b031916608052610074565b60006020828403121561005657600080fd5b81516001600160a01b038116811461006d57600080fd5b9392505050565b60805160601c61062661009f6000396000818160930152818161010701526101f001526106266000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063181f5a771461004657806333219c691461008e5780633a56bb71146100da575b600080fd5b604080518082018252601c81527f454f4153696e676c65546f6b656e526563656976657220312e312e3000000000602082015290516100859190610440565b60405180910390f35b6100b57f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610085565b6100ed6100e8366004610405565b6100ef565b005b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610164576040517f0af9f1b600000000000000000000000000000000000000000000000000000000815233600482015260240160405180910390fd5b60006101736080830183610587565b610181906020810190610522565b81019061018e9190610386565b91505073ffffffffffffffffffffffffffffffffffffffff8116158015906101e957506101be6080830183610587565b6101cc9060608101906104b3565b60008181106101dd576101dd6105c5565b90506020020135600014155b15610382577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166382bfefc86040518163ffffffff1660e01b815260040160206040518083038186803b15801561025457600080fd5b505afa158015610268573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061028c91906103e8565b73ffffffffffffffffffffffffffffffffffffffff1663a9059cbb826102b56080860186610587565b6102c39060608101906104b3565b60008181106102d4576102d46105c5565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b16815273ffffffffffffffffffffffffffffffffffffffff90941660048501526020029190910135602483015250604401602060405180830381600087803b15801561034857600080fd5b505af115801561035c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061038091906103bf565b505b5050565b6000806040838503121561039957600080fd5b82356103a4816105f4565b915060208301356103b4816105f4565b809150509250929050565b6000602082840312156103d157600080fd5b815180151581146103e157600080fd5b9392505050565b6000602082840312156103fa57600080fd5b81516103e1816105f4565b60006020828403121561041757600080fd5b813567ffffffffffffffff81111561042e57600080fd5b820160a081850312156103e157600080fd5b600060208083528351808285015260005b8181101561046d57858101830151858201604001528201610451565b8181111561047f576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126104e857600080fd5b83018035915067ffffffffffffffff82111561050357600080fd5b6020019150600581901b360382131561051b57600080fd5b9250929050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261055757600080fd5b83018035915067ffffffffffffffff82111561057257600080fd5b60200191503681900382131561051b57600080fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff418336030181126105bb57600080fd5b9190910192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461061657600080fd5b5056fea164736f6c6343000806000a",
}

var EOASingleTokenReceiverABI = EOASingleTokenReceiverMetaData.ABI

var EOASingleTokenReceiverBin = EOASingleTokenReceiverMetaData.Bin

func DeployEOASingleTokenReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, offRamp common.Address) (common.Address, *types.Transaction, *EOASingleTokenReceiver, error) {
	parsed, err := EOASingleTokenReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EOASingleTokenReceiverBin), backend, offRamp)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EOASingleTokenReceiver{EOASingleTokenReceiverCaller: EOASingleTokenReceiverCaller{contract: contract}, EOASingleTokenReceiverTransactor: EOASingleTokenReceiverTransactor{contract: contract}, EOASingleTokenReceiverFilterer: EOASingleTokenReceiverFilterer{contract: contract}}, nil
}

type EOASingleTokenReceiver struct {
	address common.Address
	abi     abi.ABI
	EOASingleTokenReceiverCaller
	EOASingleTokenReceiverTransactor
	EOASingleTokenReceiverFilterer
}

type EOASingleTokenReceiverCaller struct {
	contract *bind.BoundContract
}

type EOASingleTokenReceiverTransactor struct {
	contract *bind.BoundContract
}

type EOASingleTokenReceiverFilterer struct {
	contract *bind.BoundContract
}

type EOASingleTokenReceiverSession struct {
	Contract     *EOASingleTokenReceiver
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EOASingleTokenReceiverCallerSession struct {
	Contract *EOASingleTokenReceiverCaller
	CallOpts bind.CallOpts
}

type EOASingleTokenReceiverTransactorSession struct {
	Contract     *EOASingleTokenReceiverTransactor
	TransactOpts bind.TransactOpts
}

type EOASingleTokenReceiverRaw struct {
	Contract *EOASingleTokenReceiver
}

type EOASingleTokenReceiverCallerRaw struct {
	Contract *EOASingleTokenReceiverCaller
}

type EOASingleTokenReceiverTransactorRaw struct {
	Contract *EOASingleTokenReceiverTransactor
}

func NewEOASingleTokenReceiver(address common.Address, backend bind.ContractBackend) (*EOASingleTokenReceiver, error) {
	abi, err := abi.JSON(strings.NewReader(EOASingleTokenReceiverABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEOASingleTokenReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EOASingleTokenReceiver{address: address, abi: abi, EOASingleTokenReceiverCaller: EOASingleTokenReceiverCaller{contract: contract}, EOASingleTokenReceiverTransactor: EOASingleTokenReceiverTransactor{contract: contract}, EOASingleTokenReceiverFilterer: EOASingleTokenReceiverFilterer{contract: contract}}, nil
}

func NewEOASingleTokenReceiverCaller(address common.Address, caller bind.ContractCaller) (*EOASingleTokenReceiverCaller, error) {
	contract, err := bindEOASingleTokenReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EOASingleTokenReceiverCaller{contract: contract}, nil
}

func NewEOASingleTokenReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*EOASingleTokenReceiverTransactor, error) {
	contract, err := bindEOASingleTokenReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EOASingleTokenReceiverTransactor{contract: contract}, nil
}

func NewEOASingleTokenReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*EOASingleTokenReceiverFilterer, error) {
	contract, err := bindEOASingleTokenReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EOASingleTokenReceiverFilterer{contract: contract}, nil
}

func bindEOASingleTokenReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EOASingleTokenReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EOASingleTokenReceiver.Contract.EOASingleTokenReceiverCaller.contract.Call(opts, result, method, params...)
}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EOASingleTokenReceiver.Contract.EOASingleTokenReceiverTransactor.contract.Transfer(opts)
}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EOASingleTokenReceiver.Contract.EOASingleTokenReceiverTransactor.contract.Transact(opts, method, params...)
}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EOASingleTokenReceiver.Contract.contract.Call(opts, result, method, params...)
}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EOASingleTokenReceiver.Contract.contract.Transfer(opts)
}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EOASingleTokenReceiver.Contract.contract.Transact(opts, method, params...)
}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverCaller) OFFRAMP(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EOASingleTokenReceiver.contract.Call(opts, &out, "OFF_RAMP")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverSession) OFFRAMP() (common.Address, error) {
	return _EOASingleTokenReceiver.Contract.OFFRAMP(&_EOASingleTokenReceiver.CallOpts)
}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverCallerSession) OFFRAMP() (common.Address, error) {
	return _EOASingleTokenReceiver.Contract.OFFRAMP(&_EOASingleTokenReceiver.CallOpts)
}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EOASingleTokenReceiver.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverSession) TypeAndVersion() (string, error) {
	return _EOASingleTokenReceiver.Contract.TypeAndVersion(&_EOASingleTokenReceiver.CallOpts)
}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverCallerSession) TypeAndVersion() (string, error) {
	return _EOASingleTokenReceiver.Contract.TypeAndVersion(&_EOASingleTokenReceiver.CallOpts)
}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverTransactor) ReceiveMessage(opts *bind.TransactOpts, message CCIPMessage) (*types.Transaction, error) {
	return _EOASingleTokenReceiver.contract.Transact(opts, "receiveMessage", message)
}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverSession) ReceiveMessage(message CCIPMessage) (*types.Transaction, error) {
	return _EOASingleTokenReceiver.Contract.ReceiveMessage(&_EOASingleTokenReceiver.TransactOpts, message)
}

func (_EOASingleTokenReceiver *EOASingleTokenReceiverTransactorSession) ReceiveMessage(message CCIPMessage) (*types.Transaction, error) {
	return _EOASingleTokenReceiver.Contract.ReceiveMessage(&_EOASingleTokenReceiver.TransactOpts, message)
}

func (_EOASingleTokenReceiver *EOASingleTokenReceiver) Address() common.Address {
	return _EOASingleTokenReceiver.address
}

type EOASingleTokenReceiverInterface interface {
	OFFRAMP(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	ReceiveMessage(opts *bind.TransactOpts, message CCIPMessage) (*types.Transaction, error)

	Address() common.Address
}
