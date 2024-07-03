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

type InternalEVM2EVMMessage struct {
	SourceChainSelector uint64
	Sender              common.Address
	Receiver            common.Address
	SequenceNumber      uint64
	GasLimit            *big.Int
	Strict              bool
	Nonce               uint64
	FeeToken            common.Address
	FeeTokenAmount      *big.Int
	Data                []byte
	TokenAmounts        []ClientEVMTokenAmount
	SourceTokenData     [][]byte
	MessageId           [32]byte
}

type InternalRampMessageHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	Nonce               uint64
}

var MessageHasherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structInternal.RampMessageHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"}],\"internalType\":\"structInternal.Any2EVMRampMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610b3b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80633b61b52e1461003b578063c9f1502914610060575b600080fd5b61004e6100493660046106a8565b610073565b60405190815260200160405180910390f35b61004e61006e3660046107e7565b610086565b600061007f8383610092565b9392505050565b600061007f838361021f565b8151602080820151604092830151925160009384936100d8937f2425b0b9f9054c76ff151b0a175b18f37a4a4e82013a72e9f15c9caa095ed21f939092918891016109aa565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052805160209182012086518051888401516060808b0151908401516080808d0151950151959761013f97949693959294919391016109e7565b604051602081830303815290604052805190602001208560400151805190602001208660a001516040516020016101769190610a47565b604051602081830303815290604052805190602001208760c001516040516020016101a19190610aac565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e00160405160208183030381529060405280519060200120905092915050565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b61010001516040516020016102c298979695949392919073ffffffffffffffffffffffffffffffffffffffff9889168152968816602088015267ffffffffffffffff95861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b60405160208183030381529060405280519060200120856101200151805190602001208661014001516040516020016102fb9190610a47565b604051602081830303815290604052805190602001208761016001516040516020016101a19190610aac565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561037957610379610327565b60405290565b60405160e0810167ffffffffffffffff8111828210171561037957610379610327565b6040516101a0810167ffffffffffffffff8111828210171561037957610379610327565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561040d5761040d610327565b604052919050565b803567ffffffffffffffff8116811461042d57600080fd5b919050565b600060a0828403121561044457600080fd5b60405160a0810181811067ffffffffffffffff8211171561046757610467610327565b6040528235815290508061047d60208401610415565b602082015261048e60408401610415565b604082015261049f60608401610415565b60608201526104b060808401610415565b60808201525092915050565b600082601f8301126104cd57600080fd5b813567ffffffffffffffff8111156104e7576104e7610327565b61051860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016103c6565b81815284602083860101111561052d57600080fd5b816020850160208301376000918101602001919091529392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461042d57600080fd5b600067ffffffffffffffff82111561058857610588610327565b5060051b60200190565b600082601f8301126105a357600080fd5b813560206105b86105b38361056e565b6103c6565b82815260069290921b840181019181810190868411156105d757600080fd5b8286015b8481101561061d57604081890312156105f45760008081fd5b6105fc610356565b6106058261054a565b815281850135858201528352918301916040016105db565b509695505050505050565b600082601f83011261063957600080fd5b813560206106496105b38361056e565b82815260059290921b8401810191818101908684111561066857600080fd5b8286015b8481101561061d57803567ffffffffffffffff81111561068c5760008081fd5b61069a8986838b01016104bc565b84525091830191830161066c565b600080604083850312156106bb57600080fd5b823567ffffffffffffffff808211156106d357600080fd5b9084019061016082870312156106e857600080fd5b6106f061037f565b6106fa8784610432565b815260a08301358281111561070e57600080fd5b61071a888286016104bc565b60208301525060c08301358281111561073257600080fd5b61073e888286016104bc565b60408301525061075060e0840161054a565b606082015261010083013560808201526101208301358281111561077357600080fd5b61077f88828601610592565b60a0830152506101408301358281111561079857600080fd5b6107a488828601610628565b60c083015250935060208501359150808211156107c057600080fd5b506107cd858286016104bc565b9150509250929050565b8035801515811461042d57600080fd5b600080604083850312156107fa57600080fd5b823567ffffffffffffffff8082111561081257600080fd5b908401906101a0828703121561082757600080fd5b61082f6103a2565b61083883610415565b81526108466020840161054a565b60208201526108576040840161054a565b604082015261086860608401610415565b60608201526080830135608082015261088360a084016107d7565b60a082015261089460c08401610415565b60c08201526108a560e0840161054a565b60e0820152610100838101359082015261012080840135838111156108c957600080fd5b6108d5898287016104bc565b82840152505061014080840135838111156108ef57600080fd5b6108fb89828701610592565b828401525050610160808401358381111561091557600080fd5b61092189828701610628565b9183019190915250610180928301359281019290925250946020939093013593505050565b6000815180845260005b8181101561096c57602081850181015186830182015201610950565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b848152600067ffffffffffffffff8086166020840152808516604084015250608060608301526109dd6080830184610946565b9695505050505050565b86815260c060208201526000610a0060c0830188610946565b73ffffffffffffffffffffffffffffffffffffffff9690961660408301525067ffffffffffffffff9384166060820152608081019290925290911660a09091015292915050565b602080825282518282018190526000919060409081850190868401855b82811015610a9f578151805173ffffffffffffffffffffffffffffffffffffffff168552860151868501529284019290850190600101610a64565b5091979650505050505050565b600060208083016020845280855180835260408601915060408160051b87010192506020870160005b82811015610b21577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0888603018452610b0f858351610946565b94509285019290850190600101610ad5565b509297965050505050505056fea164736f6c6343000818000a",
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

func (_MessageHasher *MessageHasherCaller) Hash0(opts *bind.CallOpts, message InternalEVM2EVMMessage, metadataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageHasher.contract.Call(opts, &out, "hash0", message, metadataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MessageHasher *MessageHasherSession) Hash0(message InternalEVM2EVMMessage, metadataHash [32]byte) ([32]byte, error) {
	return _MessageHasher.Contract.Hash0(&_MessageHasher.CallOpts, message, metadataHash)
}

func (_MessageHasher *MessageHasherCallerSession) Hash0(message InternalEVM2EVMMessage, metadataHash [32]byte) ([32]byte, error) {
	return _MessageHasher.Contract.Hash0(&_MessageHasher.CallOpts, message, metadataHash)
}

func (_MessageHasher *MessageHasher) Address() common.Address {
	return _MessageHasher.address
}

type MessageHasherInterface interface {
	Hash(opts *bind.CallOpts, message InternalAny2EVMRampMessage, onRamp []byte) ([32]byte, error)

	Hash0(opts *bind.CallOpts, message InternalEVM2EVMMessage, metadataHash [32]byte) ([32]byte, error)

	Address() common.Address
}
