// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package message_hasher

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
	_ = abi.ConvertType
)

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type InternalAny2EVMRampMessage struct {
	Header          InternalRampMessageHeader
	Sender          []byte
	Data            []byte
	Receiver        common.Address
	GasLimit        *big.Int
	TokenAmounts    []ClientEVMTokenAmount
	SourceTokenData [][]byte
}

type InternalRampMessageHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	Nonce               uint64
}

var MessageHasherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structInternal.RampMessageHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"}],\"internalType\":\"structInternal.Any2EVMRampMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610876806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80633b61b52e14610030575b600080fd5b61004361003e366004610552565b610055565b60405190815260200160405180910390f35b60006100618383610068565b9392505050565b8151602080820151604092830151925160009384936100ae937f2425b0b9f9054c76ff151b0a175b18f37a4a4e82013a72e9f15c9caa095ed21f939092918891016106e5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052805160209182012086518051888401516060808b0151908401516080808d015195015195976101159794969395929491939101610722565b604051602081830303815290604052805190602001208560400151805190602001208660a0015160405160200161014c9190610782565b604051602081830303815290604052805190602001208760c0015160405160200161017791906107e7565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e00160405160208183030381529060405280519060200120905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610247576102476101f5565b60405290565b60405160e0810167ffffffffffffffff81118282101715610247576102476101f5565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156102b7576102b76101f5565b604052919050565b803567ffffffffffffffff811681146102d757600080fd5b919050565b600060a082840312156102ee57600080fd5b60405160a0810181811067ffffffffffffffff82111715610311576103116101f5565b60405282358152905080610327602084016102bf565b6020820152610338604084016102bf565b6040820152610349606084016102bf565b606082015261035a608084016102bf565b60808201525092915050565b600082601f83011261037757600080fd5b813567ffffffffffffffff811115610391576103916101f5565b6103c260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610270565b8181528460208386010111156103d757600080fd5b816020850160208301376000918101602001919091529392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146102d757600080fd5b600067ffffffffffffffff821115610432576104326101f5565b5060051b60200190565b600082601f83011261044d57600080fd5b8135602061046261045d83610418565b610270565b82815260069290921b8401810191818101908684111561048157600080fd5b8286015b848110156104c7576040818903121561049e5760008081fd5b6104a6610224565b6104af826103f4565b81528185013585820152835291830191604001610485565b509695505050505050565b600082601f8301126104e357600080fd5b813560206104f361045d83610418565b82815260059290921b8401810191818101908684111561051257600080fd5b8286015b848110156104c757803567ffffffffffffffff8111156105365760008081fd5b6105448986838b0101610366565b845250918301918301610516565b6000806040838503121561056557600080fd5b823567ffffffffffffffff8082111561057d57600080fd5b90840190610160828703121561059257600080fd5b61059a61024d565b6105a487846102dc565b815260a0830135828111156105b857600080fd5b6105c488828601610366565b60208301525060c0830135828111156105dc57600080fd5b6105e888828601610366565b6040830152506105fa60e084016103f4565b606082015261010083013560808201526101208301358281111561061d57600080fd5b6106298882860161043c565b60a0830152506101408301358281111561064257600080fd5b61064e888286016104d2565b60c0830152509350602085013591508082111561066a57600080fd5b5061067785828601610366565b9150509250929050565b6000815180845260005b818110156106a75760208185018101518683018201520161068b565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b848152600067ffffffffffffffff8086166020840152808516604084015250608060608301526107186080830184610681565b9695505050505050565b86815260c06020820152600061073b60c0830188610681565b73ffffffffffffffffffffffffffffffffffffffff9690961660408301525067ffffffffffffffff9384166060820152608081019290925290911660a09091015292915050565b602080825282518282018190526000919060409081850190868401855b828110156107da578151805173ffffffffffffffffffffffffffffffffffffffff16855286015186850152928401929085019060010161079f565b5091979650505050505050565b600060208083016020845280855180835260408601915060408160051b87010192506020870160005b8281101561085c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc088860301845261084a858351610681565b94509285019290850190600101610810565b509297965050505050505056fea164736f6c6343000818000a",
}

var MessageHasherABI = MessageHasherMetaData.ABI

var MessageHasherBin = MessageHasherMetaData.Bin

func DeployMessageHasher(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageHasher, error) {
	parsed, err := MessageHasherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageHasherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageHasher{address: address, abi: *parsed, MessageHasherCaller: MessageHasherCaller{contract: contract}, MessageHasherTransactor: MessageHasherTransactor{contract: contract}, MessageHasherFilterer: MessageHasherFilterer{contract: contract}}, nil
}

type MessageHasher struct {
	address common.Address
	abi     abi.ABI
	MessageHasherCaller
	MessageHasherTransactor
	MessageHasherFilterer
}

type MessageHasherCaller struct {
	contract *bind.BoundContract
}

type MessageHasherTransactor struct {
	contract *bind.BoundContract
}

type MessageHasherFilterer struct {
	contract *bind.BoundContract
}

type MessageHasherSession struct {
	Contract     *MessageHasher
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MessageHasherCallerSession struct {
	Contract *MessageHasherCaller
	CallOpts bind.CallOpts
}

type MessageHasherTransactorSession struct {
	Contract     *MessageHasherTransactor
	TransactOpts bind.TransactOpts
}

type MessageHasherRaw struct {
	Contract *MessageHasher
}

type MessageHasherCallerRaw struct {
	Contract *MessageHasherCaller
}

type MessageHasherTransactorRaw struct {
	Contract *MessageHasherTransactor
}

func NewMessageHasher(address common.Address, backend bind.ContractBackend) (*MessageHasher, error) {
	abi, err := abi.JSON(strings.NewReader(MessageHasherABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMessageHasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageHasher{address: address, abi: abi, MessageHasherCaller: MessageHasherCaller{contract: contract}, MessageHasherTransactor: MessageHasherTransactor{contract: contract}, MessageHasherFilterer: MessageHasherFilterer{contract: contract}}, nil
}

func NewMessageHasherCaller(address common.Address, caller bind.ContractCaller) (*MessageHasherCaller, error) {
	contract, err := bindMessageHasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageHasherCaller{contract: contract}, nil
}

func NewMessageHasherTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageHasherTransactor, error) {
	contract, err := bindMessageHasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageHasherTransactor{contract: contract}, nil
}

func NewMessageHasherFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageHasherFilterer, error) {
	contract, err := bindMessageHasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageHasherFilterer{contract: contract}, nil
}

func bindMessageHasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MessageHasherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MessageHasher *MessageHasherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageHasher.Contract.MessageHasherCaller.contract.Call(opts, result, method, params...)
}

func (_MessageHasher *MessageHasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageHasher.Contract.MessageHasherTransactor.contract.Transfer(opts)
}

func (_MessageHasher *MessageHasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageHasher.Contract.MessageHasherTransactor.contract.Transact(opts, method, params...)
}

func (_MessageHasher *MessageHasherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageHasher.Contract.contract.Call(opts, result, method, params...)
}

func (_MessageHasher *MessageHasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageHasher.Contract.contract.Transfer(opts)
}

func (_MessageHasher *MessageHasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageHasher.Contract.contract.Transact(opts, method, params...)
}

func (_MessageHasher *MessageHasherCaller) Hash(opts *bind.CallOpts, message InternalAny2EVMRampMessage, onRamp []byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "hash", message, onRamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) Hash(message InternalAny2EVMRampMessage, onRamp []byte) ([32]byte, error) {
	return _MessageHasher.Contract.Hash(&_MessageHasher.CallOpts, message, onRamp)
}

func (_MessageHasher *MessageHasherCallerSession) Hash(message InternalAny2EVMRampMessage, onRamp []byte) ([32]byte, error) {
	return _MessageHasher.Contract.Hash(&_MessageHasher.CallOpts, message, onRamp)
}

func (_MessageHasher *MessageHasher) Address() common.Address {
	return _MessageHasher.address
}

type MessageHasherInterface interface {
	Hash(opts *bind.CallOpts, message InternalAny2EVMRampMessage, onRamp []byte) ([32]byte, error)

	Address() common.Address
}
