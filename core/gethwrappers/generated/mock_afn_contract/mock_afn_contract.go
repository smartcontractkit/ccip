// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_afn_contract

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
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated"
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

var MockAFNContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyVoted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"InvalidVoter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustRecoverFromBadSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RecoveryNotNecessary\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AFNBadSignal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"parties\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goodQuorum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"badQuorum\",\"type\":\"uint256\"}],\"name\":\"AFNConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RecoveredFromBadSignal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"name\":\"RootBlessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"VoteBad\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"VoteToBless\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"badSignalReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBadVotersAndVotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfigVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParticipants\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getVotesToBlessRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getWeightByParticipant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWeightThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blessing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"badSignal\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"hasVotedBad\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"hasVotedToBlessRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isBlessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverFromBadSignal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_badSignal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"parties\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"goodQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"badQuorum\",\"type\":\"uint256\"}],\"name\":\"setAFNConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteBad\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"name\":\"voteToBlessRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610596806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806351856565116100975780639dc6edc7116100665780639dc6edc7146101dd578063c3453fa5146101e4578063f438c9c0146101f6578063ff888fb11461020a57600080fd5b806351856565146101905780635aa68ac0146101a457806379adb16e146101b35780638e1d4e61146101cf57600080fd5b80633cd4f669116100d35780633cd4f6691461015357806342ac28061461016257806346f8e6d71461016f578063508ede921461017a57600080fd5b80632cb145d4146100fa5780632ea953711461010e578063365f15ec14610137575b600080fd5b61010c6000805460ff19166001179055565b005b61012261011c36600461024c565b50600090565b60405190151581526020015b60405180910390f35b61010c61014536600461026e565b50506000805460ff19169055565b61010c6000805460ff19169055565b6000546101229060ff1681565b60005460ff16610122565b6101226101883660046102e3565b600092915050565b61010c61019e36600461041a565b50505050565b606060405161012e919061053b565b6101c161011c36600461024c565b60405190815260200161012e565b6101c161011c36600461054e565b60006101c1565b6060600060405161012e929190610567565b60408051600080825260208201520161012e565b61012261021836600461054e565b5060005460ff161590565b803573ffffffffffffffffffffffffffffffffffffffff8116811461024757600080fd5b919050565b60006020828403121561025e57600080fd5b61026782610223565b9392505050565b6000806020838503121561028157600080fd5b823567ffffffffffffffff8082111561029957600080fd5b818501915085601f8301126102ad57600080fd5b8135818111156102bc57600080fd5b8660208260051b85010111156102d157600080fd5b60209290920196919550909350505050565b600080604083850312156102f657600080fd5b6102ff83610223565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156103835761038361030d565b604052919050565b600067ffffffffffffffff8211156103a5576103a561030d565b5060051b60200190565b600082601f8301126103c057600080fd5b813560206103d56103d08361038b565b61033c565b82815260059290921b840181019181810190868411156103f457600080fd5b8286015b8481101561040f57803583529183019183016103f8565b509695505050505050565b6000806000806080858703121561043057600080fd5b843567ffffffffffffffff8082111561044857600080fd5b818701915087601f83011261045c57600080fd5b8135602061046c6103d08361038b565b82815260059290921b8401810191818101908b84111561048b57600080fd5b948201945b838610156104b0576104a186610223565b82529482019490820190610490565b985050880135925050808211156104c657600080fd5b506104d3878288016103af565b949794965050505060408301359260600135919050565b600081518084526020808501945080840160005b8381101561053057815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016104fe565b509495945050505050565b60208152600061026760208301846104ea565b60006020828403121561056057600080fd5b5035919050565b60408152600061057a60408301856104ea565b9050826020830152939250505056fea164736f6c634300080f000a",
}

var MockAFNContractABI = MockAFNContractMetaData.ABI

var MockAFNContractBin = MockAFNContractMetaData.Bin

func DeployMockAFNContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockAFNContract, error) {
	parsed, err := MockAFNContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockAFNContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockAFNContract{MockAFNContractCaller: MockAFNContractCaller{contract: contract}, MockAFNContractTransactor: MockAFNContractTransactor{contract: contract}, MockAFNContractFilterer: MockAFNContractFilterer{contract: contract}}, nil
}

type MockAFNContract struct {
	address common.Address
	abi     abi.ABI
	MockAFNContractCaller
	MockAFNContractTransactor
	MockAFNContractFilterer
}

type MockAFNContractCaller struct {
	contract *bind.BoundContract
}

type MockAFNContractTransactor struct {
	contract *bind.BoundContract
}

type MockAFNContractFilterer struct {
	contract *bind.BoundContract
}

type MockAFNContractSession struct {
	Contract     *MockAFNContract
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockAFNContractCallerSession struct {
	Contract *MockAFNContractCaller
	CallOpts bind.CallOpts
}

type MockAFNContractTransactorSession struct {
	Contract     *MockAFNContractTransactor
	TransactOpts bind.TransactOpts
}

type MockAFNContractRaw struct {
	Contract *MockAFNContract
}

type MockAFNContractCallerRaw struct {
	Contract *MockAFNContractCaller
}

type MockAFNContractTransactorRaw struct {
	Contract *MockAFNContractTransactor
}

func NewMockAFNContract(address common.Address, backend bind.ContractBackend) (*MockAFNContract, error) {
	abi, err := abi.JSON(strings.NewReader(MockAFNContractABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockAFNContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockAFNContract{address: address, abi: abi, MockAFNContractCaller: MockAFNContractCaller{contract: contract}, MockAFNContractTransactor: MockAFNContractTransactor{contract: contract}, MockAFNContractFilterer: MockAFNContractFilterer{contract: contract}}, nil
}

func NewMockAFNContractCaller(address common.Address, caller bind.ContractCaller) (*MockAFNContractCaller, error) {
	contract, err := bindMockAFNContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockAFNContractCaller{contract: contract}, nil
}

func NewMockAFNContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MockAFNContractTransactor, error) {
	contract, err := bindMockAFNContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockAFNContractTransactor{contract: contract}, nil
}

func NewMockAFNContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MockAFNContractFilterer, error) {
	contract, err := bindMockAFNContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockAFNContractFilterer{contract: contract}, nil
}

func bindMockAFNContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockAFNContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_MockAFNContract *MockAFNContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockAFNContract.Contract.MockAFNContractCaller.contract.Call(opts, result, method, params...)
}

func (_MockAFNContract *MockAFNContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockAFNContract.Contract.MockAFNContractTransactor.contract.Transfer(opts)
}

func (_MockAFNContract *MockAFNContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockAFNContract.Contract.MockAFNContractTransactor.contract.Transact(opts, method, params...)
}

func (_MockAFNContract *MockAFNContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockAFNContract.Contract.contract.Call(opts, result, method, params...)
}

func (_MockAFNContract *MockAFNContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockAFNContract.Contract.contract.Transfer(opts)
}

func (_MockAFNContract *MockAFNContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockAFNContract.Contract.contract.Transact(opts, method, params...)
}

func (_MockAFNContract *MockAFNContractCaller) BadSignalReceived(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockAFNContract.contract.Call(opts, &out, "badSignalReceived")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockAFNContract *MockAFNContractSession) BadSignalReceived() (bool, error) {
	return _MockAFNContract.Contract.BadSignalReceived(&_MockAFNContract.CallOpts)
}

func (_MockAFNContract *MockAFNContractCallerSession) BadSignalReceived() (bool, error) {
	return _MockAFNContract.Contract.BadSignalReceived(&_MockAFNContract.CallOpts)
}

func (_MockAFNContract *MockAFNContractCaller) GetBadVotersAndVotes(opts *bind.CallOpts) (GetBadVotersAndVotes,

	error) {
	var out []interface{}
	err := _MockAFNContract.contract.Call(opts, &out, "getBadVotersAndVotes")

	outstruct := new(GetBadVotersAndVotes)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Voters = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Votes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_MockAFNContract *MockAFNContractSession) GetBadVotersAndVotes() (GetBadVotersAndVotes,

	error) {
	return _MockAFNContract.Contract.GetBadVotersAndVotes(&_MockAFNContract.CallOpts)
}

func (_MockAFNContract *MockAFNContractCallerSession) GetBadVotersAndVotes() (GetBadVotersAndVotes,

	error) {
	return _MockAFNContract.Contract.GetBadVotersAndVotes(&_MockAFNContract.CallOpts)
}

func (_MockAFNContract *MockAFNContractCaller) GetConfigVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockAFNContract.contract.Call(opts, &out, "getConfigVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockAFNContract *MockAFNContractSession) GetConfigVersion() (*big.Int, error) {
	return _MockAFNContract.Contract.GetConfigVersion(&_MockAFNContract.CallOpts)
}

func (_MockAFNContract *MockAFNContractCallerSession) GetConfigVersion() (*big.Int, error) {
	return _MockAFNContract.Contract.GetConfigVersion(&_MockAFNContract.CallOpts)
}

func (_MockAFNContract *MockAFNContractCaller) GetVotesToBlessRoot(opts *bind.CallOpts, root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MockAFNContract.contract.Call(opts, &out, "getVotesToBlessRoot", root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockAFNContract *MockAFNContractSession) GetVotesToBlessRoot(root [32]byte) (*big.Int, error) {
	return _MockAFNContract.Contract.GetVotesToBlessRoot(&_MockAFNContract.CallOpts, root)
}

func (_MockAFNContract *MockAFNContractCallerSession) GetVotesToBlessRoot(root [32]byte) (*big.Int, error) {
	return _MockAFNContract.Contract.GetVotesToBlessRoot(&_MockAFNContract.CallOpts, root)
}

func (_MockAFNContract *MockAFNContractCaller) GetWeightByParticipant(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockAFNContract.contract.Call(opts, &out, "getWeightByParticipant", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockAFNContract *MockAFNContractSession) GetWeightByParticipant(arg0 common.Address) (*big.Int, error) {
	return _MockAFNContract.Contract.GetWeightByParticipant(&_MockAFNContract.CallOpts, arg0)
}

func (_MockAFNContract *MockAFNContractCallerSession) GetWeightByParticipant(arg0 common.Address) (*big.Int, error) {
	return _MockAFNContract.Contract.GetWeightByParticipant(&_MockAFNContract.CallOpts, arg0)
}

func (_MockAFNContract *MockAFNContractCaller) HasVotedBad(opts *bind.CallOpts, participant common.Address) (bool, error) {
	var out []interface{}
	err := _MockAFNContract.contract.Call(opts, &out, "hasVotedBad", participant)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockAFNContract *MockAFNContractSession) HasVotedBad(participant common.Address) (bool, error) {
	return _MockAFNContract.Contract.HasVotedBad(&_MockAFNContract.CallOpts, participant)
}

func (_MockAFNContract *MockAFNContractCallerSession) HasVotedBad(participant common.Address) (bool, error) {
	return _MockAFNContract.Contract.HasVotedBad(&_MockAFNContract.CallOpts, participant)
}

func (_MockAFNContract *MockAFNContractCaller) HasVotedToBlessRoot(opts *bind.CallOpts, participant common.Address, root [32]byte) (bool, error) {
	var out []interface{}
	err := _MockAFNContract.contract.Call(opts, &out, "hasVotedToBlessRoot", participant, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockAFNContract *MockAFNContractSession) HasVotedToBlessRoot(participant common.Address, root [32]byte) (bool, error) {
	return _MockAFNContract.Contract.HasVotedToBlessRoot(&_MockAFNContract.CallOpts, participant, root)
}

func (_MockAFNContract *MockAFNContractCallerSession) HasVotedToBlessRoot(participant common.Address, root [32]byte) (bool, error) {
	return _MockAFNContract.Contract.HasVotedToBlessRoot(&_MockAFNContract.CallOpts, participant, root)
}

func (_MockAFNContract *MockAFNContractCaller) IsBlessed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _MockAFNContract.contract.Call(opts, &out, "isBlessed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockAFNContract *MockAFNContractSession) IsBlessed(arg0 [32]byte) (bool, error) {
	return _MockAFNContract.Contract.IsBlessed(&_MockAFNContract.CallOpts, arg0)
}

func (_MockAFNContract *MockAFNContractCallerSession) IsBlessed(arg0 [32]byte) (bool, error) {
	return _MockAFNContract.Contract.IsBlessed(&_MockAFNContract.CallOpts, arg0)
}

func (_MockAFNContract *MockAFNContractCaller) SBadSignal(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockAFNContract.contract.Call(opts, &out, "s_badSignal")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockAFNContract *MockAFNContractSession) SBadSignal() (bool, error) {
	return _MockAFNContract.Contract.SBadSignal(&_MockAFNContract.CallOpts)
}

func (_MockAFNContract *MockAFNContractCallerSession) SBadSignal() (bool, error) {
	return _MockAFNContract.Contract.SBadSignal(&_MockAFNContract.CallOpts)
}

func (_MockAFNContract *MockAFNContractTransactor) GetParticipants(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockAFNContract.contract.Transact(opts, "getParticipants")
}

func (_MockAFNContract *MockAFNContractSession) GetParticipants() (*types.Transaction, error) {
	return _MockAFNContract.Contract.GetParticipants(&_MockAFNContract.TransactOpts)
}

func (_MockAFNContract *MockAFNContractTransactorSession) GetParticipants() (*types.Transaction, error) {
	return _MockAFNContract.Contract.GetParticipants(&_MockAFNContract.TransactOpts)
}

func (_MockAFNContract *MockAFNContractTransactor) GetWeightThresholds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockAFNContract.contract.Transact(opts, "getWeightThresholds")
}

func (_MockAFNContract *MockAFNContractSession) GetWeightThresholds() (*types.Transaction, error) {
	return _MockAFNContract.Contract.GetWeightThresholds(&_MockAFNContract.TransactOpts)
}

func (_MockAFNContract *MockAFNContractTransactorSession) GetWeightThresholds() (*types.Transaction, error) {
	return _MockAFNContract.Contract.GetWeightThresholds(&_MockAFNContract.TransactOpts)
}

func (_MockAFNContract *MockAFNContractTransactor) RecoverFromBadSignal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockAFNContract.contract.Transact(opts, "recoverFromBadSignal")
}

func (_MockAFNContract *MockAFNContractSession) RecoverFromBadSignal() (*types.Transaction, error) {
	return _MockAFNContract.Contract.RecoverFromBadSignal(&_MockAFNContract.TransactOpts)
}

func (_MockAFNContract *MockAFNContractTransactorSession) RecoverFromBadSignal() (*types.Transaction, error) {
	return _MockAFNContract.Contract.RecoverFromBadSignal(&_MockAFNContract.TransactOpts)
}

func (_MockAFNContract *MockAFNContractTransactor) SetAFNConfig(opts *bind.TransactOpts, parties []common.Address, weights []*big.Int, goodQuorum *big.Int, badQuorum *big.Int) (*types.Transaction, error) {
	return _MockAFNContract.contract.Transact(opts, "setAFNConfig", parties, weights, goodQuorum, badQuorum)
}

func (_MockAFNContract *MockAFNContractSession) SetAFNConfig(parties []common.Address, weights []*big.Int, goodQuorum *big.Int, badQuorum *big.Int) (*types.Transaction, error) {
	return _MockAFNContract.Contract.SetAFNConfig(&_MockAFNContract.TransactOpts, parties, weights, goodQuorum, badQuorum)
}

func (_MockAFNContract *MockAFNContractTransactorSession) SetAFNConfig(parties []common.Address, weights []*big.Int, goodQuorum *big.Int, badQuorum *big.Int) (*types.Transaction, error) {
	return _MockAFNContract.Contract.SetAFNConfig(&_MockAFNContract.TransactOpts, parties, weights, goodQuorum, badQuorum)
}

func (_MockAFNContract *MockAFNContractTransactor) VoteBad(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockAFNContract.contract.Transact(opts, "voteBad")
}

func (_MockAFNContract *MockAFNContractSession) VoteBad() (*types.Transaction, error) {
	return _MockAFNContract.Contract.VoteBad(&_MockAFNContract.TransactOpts)
}

func (_MockAFNContract *MockAFNContractTransactorSession) VoteBad() (*types.Transaction, error) {
	return _MockAFNContract.Contract.VoteBad(&_MockAFNContract.TransactOpts)
}

func (_MockAFNContract *MockAFNContractTransactor) VoteToBlessRoots(opts *bind.TransactOpts, arg0 [][32]byte) (*types.Transaction, error) {
	return _MockAFNContract.contract.Transact(opts, "voteToBlessRoots", arg0)
}

func (_MockAFNContract *MockAFNContractSession) VoteToBlessRoots(arg0 [][32]byte) (*types.Transaction, error) {
	return _MockAFNContract.Contract.VoteToBlessRoots(&_MockAFNContract.TransactOpts, arg0)
}

func (_MockAFNContract *MockAFNContractTransactorSession) VoteToBlessRoots(arg0 [][32]byte) (*types.Transaction, error) {
	return _MockAFNContract.Contract.VoteToBlessRoots(&_MockAFNContract.TransactOpts, arg0)
}

type MockAFNContractAFNBadSignalIterator struct {
	Event *MockAFNContractAFNBadSignal

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockAFNContractAFNBadSignalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockAFNContractAFNBadSignal)
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
		it.Event = new(MockAFNContractAFNBadSignal)
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

func (it *MockAFNContractAFNBadSignalIterator) Error() error {
	return it.fail
}

func (it *MockAFNContractAFNBadSignalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockAFNContractAFNBadSignal struct {
	Timestamp *big.Int
	Raw       types.Log
}

func (_MockAFNContract *MockAFNContractFilterer) FilterAFNBadSignal(opts *bind.FilterOpts) (*MockAFNContractAFNBadSignalIterator, error) {

	logs, sub, err := _MockAFNContract.contract.FilterLogs(opts, "AFNBadSignal")
	if err != nil {
		return nil, err
	}
	return &MockAFNContractAFNBadSignalIterator{contract: _MockAFNContract.contract, event: "AFNBadSignal", logs: logs, sub: sub}, nil
}

func (_MockAFNContract *MockAFNContractFilterer) WatchAFNBadSignal(opts *bind.WatchOpts, sink chan<- *MockAFNContractAFNBadSignal) (event.Subscription, error) {

	logs, sub, err := _MockAFNContract.contract.WatchLogs(opts, "AFNBadSignal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockAFNContractAFNBadSignal)
				if err := _MockAFNContract.contract.UnpackLog(event, "AFNBadSignal", log); err != nil {
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

func (_MockAFNContract *MockAFNContractFilterer) ParseAFNBadSignal(log types.Log) (*MockAFNContractAFNBadSignal, error) {
	event := new(MockAFNContractAFNBadSignal)
	if err := _MockAFNContract.contract.UnpackLog(event, "AFNBadSignal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockAFNContractAFNConfigSetIterator struct {
	Event *MockAFNContractAFNConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockAFNContractAFNConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockAFNContractAFNConfigSet)
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
		it.Event = new(MockAFNContractAFNConfigSet)
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

func (it *MockAFNContractAFNConfigSetIterator) Error() error {
	return it.fail
}

func (it *MockAFNContractAFNConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockAFNContractAFNConfigSet struct {
	Parties    []common.Address
	Weights    []*big.Int
	GoodQuorum *big.Int
	BadQuorum  *big.Int
	Raw        types.Log
}

func (_MockAFNContract *MockAFNContractFilterer) FilterAFNConfigSet(opts *bind.FilterOpts) (*MockAFNContractAFNConfigSetIterator, error) {

	logs, sub, err := _MockAFNContract.contract.FilterLogs(opts, "AFNConfigSet")
	if err != nil {
		return nil, err
	}
	return &MockAFNContractAFNConfigSetIterator{contract: _MockAFNContract.contract, event: "AFNConfigSet", logs: logs, sub: sub}, nil
}

func (_MockAFNContract *MockAFNContractFilterer) WatchAFNConfigSet(opts *bind.WatchOpts, sink chan<- *MockAFNContractAFNConfigSet) (event.Subscription, error) {

	logs, sub, err := _MockAFNContract.contract.WatchLogs(opts, "AFNConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockAFNContractAFNConfigSet)
				if err := _MockAFNContract.contract.UnpackLog(event, "AFNConfigSet", log); err != nil {
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

func (_MockAFNContract *MockAFNContractFilterer) ParseAFNConfigSet(log types.Log) (*MockAFNContractAFNConfigSet, error) {
	event := new(MockAFNContractAFNConfigSet)
	if err := _MockAFNContract.contract.UnpackLog(event, "AFNConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockAFNContractRecoveredFromBadSignalIterator struct {
	Event *MockAFNContractRecoveredFromBadSignal

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockAFNContractRecoveredFromBadSignalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockAFNContractRecoveredFromBadSignal)
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
		it.Event = new(MockAFNContractRecoveredFromBadSignal)
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

func (it *MockAFNContractRecoveredFromBadSignalIterator) Error() error {
	return it.fail
}

func (it *MockAFNContractRecoveredFromBadSignalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockAFNContractRecoveredFromBadSignal struct {
	Raw types.Log
}

func (_MockAFNContract *MockAFNContractFilterer) FilterRecoveredFromBadSignal(opts *bind.FilterOpts) (*MockAFNContractRecoveredFromBadSignalIterator, error) {

	logs, sub, err := _MockAFNContract.contract.FilterLogs(opts, "RecoveredFromBadSignal")
	if err != nil {
		return nil, err
	}
	return &MockAFNContractRecoveredFromBadSignalIterator{contract: _MockAFNContract.contract, event: "RecoveredFromBadSignal", logs: logs, sub: sub}, nil
}

func (_MockAFNContract *MockAFNContractFilterer) WatchRecoveredFromBadSignal(opts *bind.WatchOpts, sink chan<- *MockAFNContractRecoveredFromBadSignal) (event.Subscription, error) {

	logs, sub, err := _MockAFNContract.contract.WatchLogs(opts, "RecoveredFromBadSignal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockAFNContractRecoveredFromBadSignal)
				if err := _MockAFNContract.contract.UnpackLog(event, "RecoveredFromBadSignal", log); err != nil {
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

func (_MockAFNContract *MockAFNContractFilterer) ParseRecoveredFromBadSignal(log types.Log) (*MockAFNContractRecoveredFromBadSignal, error) {
	event := new(MockAFNContractRecoveredFromBadSignal)
	if err := _MockAFNContract.contract.UnpackLog(event, "RecoveredFromBadSignal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockAFNContractRootBlessedIterator struct {
	Event *MockAFNContractRootBlessed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockAFNContractRootBlessedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockAFNContractRootBlessed)
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
		it.Event = new(MockAFNContractRootBlessed)
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

func (it *MockAFNContractRootBlessedIterator) Error() error {
	return it.fail
}

func (it *MockAFNContractRootBlessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockAFNContractRootBlessed struct {
	Root  [32]byte
	Votes *big.Int
	Raw   types.Log
}

func (_MockAFNContract *MockAFNContractFilterer) FilterRootBlessed(opts *bind.FilterOpts, root [][32]byte) (*MockAFNContractRootBlessedIterator, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _MockAFNContract.contract.FilterLogs(opts, "RootBlessed", rootRule)
	if err != nil {
		return nil, err
	}
	return &MockAFNContractRootBlessedIterator{contract: _MockAFNContract.contract, event: "RootBlessed", logs: logs, sub: sub}, nil
}

func (_MockAFNContract *MockAFNContractFilterer) WatchRootBlessed(opts *bind.WatchOpts, sink chan<- *MockAFNContractRootBlessed, root [][32]byte) (event.Subscription, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _MockAFNContract.contract.WatchLogs(opts, "RootBlessed", rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockAFNContractRootBlessed)
				if err := _MockAFNContract.contract.UnpackLog(event, "RootBlessed", log); err != nil {
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

func (_MockAFNContract *MockAFNContractFilterer) ParseRootBlessed(log types.Log) (*MockAFNContractRootBlessed, error) {
	event := new(MockAFNContractRootBlessed)
	if err := _MockAFNContract.contract.UnpackLog(event, "RootBlessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockAFNContractVoteBadIterator struct {
	Event *MockAFNContractVoteBad

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockAFNContractVoteBadIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockAFNContractVoteBad)
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
		it.Event = new(MockAFNContractVoteBad)
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

func (it *MockAFNContractVoteBadIterator) Error() error {
	return it.fail
}

func (it *MockAFNContractVoteBadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockAFNContractVoteBad struct {
	Voter  common.Address
	Weight *big.Int
	Raw    types.Log
}

func (_MockAFNContract *MockAFNContractFilterer) FilterVoteBad(opts *bind.FilterOpts, voter []common.Address) (*MockAFNContractVoteBadIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _MockAFNContract.contract.FilterLogs(opts, "VoteBad", voterRule)
	if err != nil {
		return nil, err
	}
	return &MockAFNContractVoteBadIterator{contract: _MockAFNContract.contract, event: "VoteBad", logs: logs, sub: sub}, nil
}

func (_MockAFNContract *MockAFNContractFilterer) WatchVoteBad(opts *bind.WatchOpts, sink chan<- *MockAFNContractVoteBad, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _MockAFNContract.contract.WatchLogs(opts, "VoteBad", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockAFNContractVoteBad)
				if err := _MockAFNContract.contract.UnpackLog(event, "VoteBad", log); err != nil {
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

func (_MockAFNContract *MockAFNContractFilterer) ParseVoteBad(log types.Log) (*MockAFNContractVoteBad, error) {
	event := new(MockAFNContractVoteBad)
	if err := _MockAFNContract.contract.UnpackLog(event, "VoteBad", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockAFNContractVoteToBlessIterator struct {
	Event *MockAFNContractVoteToBless

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockAFNContractVoteToBlessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockAFNContractVoteToBless)
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
		it.Event = new(MockAFNContractVoteToBless)
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

func (it *MockAFNContractVoteToBlessIterator) Error() error {
	return it.fail
}

func (it *MockAFNContractVoteToBlessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockAFNContractVoteToBless struct {
	Voter  common.Address
	Root   [32]byte
	Weight *big.Int
	Raw    types.Log
}

func (_MockAFNContract *MockAFNContractFilterer) FilterVoteToBless(opts *bind.FilterOpts, voter []common.Address, root [][32]byte) (*MockAFNContractVoteToBlessIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _MockAFNContract.contract.FilterLogs(opts, "VoteToBless", voterRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &MockAFNContractVoteToBlessIterator{contract: _MockAFNContract.contract, event: "VoteToBless", logs: logs, sub: sub}, nil
}

func (_MockAFNContract *MockAFNContractFilterer) WatchVoteToBless(opts *bind.WatchOpts, sink chan<- *MockAFNContractVoteToBless, voter []common.Address, root [][32]byte) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _MockAFNContract.contract.WatchLogs(opts, "VoteToBless", voterRule, rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockAFNContractVoteToBless)
				if err := _MockAFNContract.contract.UnpackLog(event, "VoteToBless", log); err != nil {
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

func (_MockAFNContract *MockAFNContractFilterer) ParseVoteToBless(log types.Log) (*MockAFNContractVoteToBless, error) {
	event := new(MockAFNContractVoteToBless)
	if err := _MockAFNContract.contract.UnpackLog(event, "VoteToBless", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetBadVotersAndVotes struct {
	Voters []common.Address
	Votes  *big.Int
}

func (_MockAFNContract *MockAFNContract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MockAFNContract.abi.Events["AFNBadSignal"].ID:
		return _MockAFNContract.ParseAFNBadSignal(log)
	case _MockAFNContract.abi.Events["AFNConfigSet"].ID:
		return _MockAFNContract.ParseAFNConfigSet(log)
	case _MockAFNContract.abi.Events["RecoveredFromBadSignal"].ID:
		return _MockAFNContract.ParseRecoveredFromBadSignal(log)
	case _MockAFNContract.abi.Events["RootBlessed"].ID:
		return _MockAFNContract.ParseRootBlessed(log)
	case _MockAFNContract.abi.Events["VoteBad"].ID:
		return _MockAFNContract.ParseVoteBad(log)
	case _MockAFNContract.abi.Events["VoteToBless"].ID:
		return _MockAFNContract.ParseVoteToBless(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MockAFNContractAFNBadSignal) Topic() common.Hash {
	return common.HexToHash("0x73907f5e30313a1ab6e1815608b22b40911f1a7decec69d5df18a2298002bacb")
}

func (MockAFNContractAFNConfigSet) Topic() common.Hash {
	return common.HexToHash("0x69af5b8b5b348d6b619cb6b338b5cfd865aa9e8cedd36a4a69257a9a07ebedaa")
}

func (MockAFNContractRecoveredFromBadSignal) Topic() common.Hash {
	return common.HexToHash("0x3e48434bea67b1e259c2380d289dcb6372257ab2c37bc86f0e1acf83a7b07ac0")
}

func (MockAFNContractRootBlessed) Topic() common.Hash {
	return common.HexToHash("0x719fab74b843fdceffa591cc0a3445a9dddc9e1e304471baed67e8408a1405c7")
}

func (MockAFNContractVoteBad) Topic() common.Hash {
	return common.HexToHash("0xa5889da6c2d25ef72eaae82bb0b8acf51eeebdd6bd12f1a24360de7d9b9cfa28")
}

func (MockAFNContractVoteToBless) Topic() common.Hash {
	return common.HexToHash("0x262f79a5a063a0af3e27989b0b0f0ae1e2c19257d27efe01a7f0cab7b3b470a4")
}

func (_MockAFNContract *MockAFNContract) Address() common.Address {
	return _MockAFNContract.address
}

type MockAFNContractInterface interface {
	BadSignalReceived(opts *bind.CallOpts) (bool, error)

	GetBadVotersAndVotes(opts *bind.CallOpts) (GetBadVotersAndVotes,

		error)

	GetConfigVersion(opts *bind.CallOpts) (*big.Int, error)

	GetVotesToBlessRoot(opts *bind.CallOpts, root [32]byte) (*big.Int, error)

	GetWeightByParticipant(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	HasVotedBad(opts *bind.CallOpts, participant common.Address) (bool, error)

	HasVotedToBlessRoot(opts *bind.CallOpts, participant common.Address, root [32]byte) (bool, error)

	IsBlessed(opts *bind.CallOpts, arg0 [32]byte) (bool, error)

	SBadSignal(opts *bind.CallOpts) (bool, error)

	GetParticipants(opts *bind.TransactOpts) (*types.Transaction, error)

	GetWeightThresholds(opts *bind.TransactOpts) (*types.Transaction, error)

	RecoverFromBadSignal(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAFNConfig(opts *bind.TransactOpts, parties []common.Address, weights []*big.Int, goodQuorum *big.Int, badQuorum *big.Int) (*types.Transaction, error)

	VoteBad(opts *bind.TransactOpts) (*types.Transaction, error)

	VoteToBlessRoots(opts *bind.TransactOpts, arg0 [][32]byte) (*types.Transaction, error)

	FilterAFNBadSignal(opts *bind.FilterOpts) (*MockAFNContractAFNBadSignalIterator, error)

	WatchAFNBadSignal(opts *bind.WatchOpts, sink chan<- *MockAFNContractAFNBadSignal) (event.Subscription, error)

	ParseAFNBadSignal(log types.Log) (*MockAFNContractAFNBadSignal, error)

	FilterAFNConfigSet(opts *bind.FilterOpts) (*MockAFNContractAFNConfigSetIterator, error)

	WatchAFNConfigSet(opts *bind.WatchOpts, sink chan<- *MockAFNContractAFNConfigSet) (event.Subscription, error)

	ParseAFNConfigSet(log types.Log) (*MockAFNContractAFNConfigSet, error)

	FilterRecoveredFromBadSignal(opts *bind.FilterOpts) (*MockAFNContractRecoveredFromBadSignalIterator, error)

	WatchRecoveredFromBadSignal(opts *bind.WatchOpts, sink chan<- *MockAFNContractRecoveredFromBadSignal) (event.Subscription, error)

	ParseRecoveredFromBadSignal(log types.Log) (*MockAFNContractRecoveredFromBadSignal, error)

	FilterRootBlessed(opts *bind.FilterOpts, root [][32]byte) (*MockAFNContractRootBlessedIterator, error)

	WatchRootBlessed(opts *bind.WatchOpts, sink chan<- *MockAFNContractRootBlessed, root [][32]byte) (event.Subscription, error)

	ParseRootBlessed(log types.Log) (*MockAFNContractRootBlessed, error)

	FilterVoteBad(opts *bind.FilterOpts, voter []common.Address) (*MockAFNContractVoteBadIterator, error)

	WatchVoteBad(opts *bind.WatchOpts, sink chan<- *MockAFNContractVoteBad, voter []common.Address) (event.Subscription, error)

	ParseVoteBad(log types.Log) (*MockAFNContractVoteBad, error)

	FilterVoteToBless(opts *bind.FilterOpts, voter []common.Address, root [][32]byte) (*MockAFNContractVoteToBlessIterator, error)

	WatchVoteToBless(opts *bind.WatchOpts, sink chan<- *MockAFNContractVoteToBless, voter []common.Address, root [][32]byte) (event.Subscription, error)

	ParseVoteToBless(log types.Log) (*MockAFNContractVoteToBless, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
