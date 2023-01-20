// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package any_2_evm_toll_offramp_router

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

type CommonAny2EVMMessage struct {
	SourceChainId        uint64
	Sender               []byte
	Data                 []byte
	DestTokensAndAmounts []CommonEVMTokenAndAmount
}

type CommonEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

var Any2EVMTollOffRampRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"AlreadyConfigured\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MustCallFromOffRamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoOffRampsConfigured\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"addOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRamps\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"removeOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCommon.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"routeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001331380380620013318339810160408190526200003491620002f7565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be8162000198565b50508151620000d69150600390602084019062000243565b5060005b815181101562000190576040518060400160405280826001600160601b0316815260200160011515815250600260008484815181106200011e576200011e620003c9565b6020908102919091018101516001600160a01b031682528181019290925260400160002082518154939092015115156c01000000000000000000000000026001600160681b03199093166001600160601b03909216919091179190911790556200018881620003df565b9050620000da565b505062000407565b336001600160a01b03821603620001f25760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200029b579160200282015b828111156200029b57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000264565b50620002a9929150620002ad565b5090565b5b80821115620002a95760008155600101620002ae565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b0381168114620002f257600080fd5b919050565b600060208083850312156200030b57600080fd5b82516001600160401b03808211156200032357600080fd5b818501915085601f8301126200033857600080fd5b8151818111156200034d576200034d620002c4565b8060051b604051601f19603f83011681018181108582111715620003755762000375620002c4565b6040529182528482019250838101850191888311156200039457600080fd5b938501935b82851015620003bd57620003ad85620002da565b8452938501939285019262000399565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000600182016200040057634e487b7160e01b600052601160045260246000fd5b5060010190565b610f1a80620004176000396000f3fe608060405234801561001057600080fd5b50600436106100a25760003560e01c80638da5cb5b11610076578063a40e69c71161005b578063a40e69c71461018c578063adb9f71b146101a1578063f2fde38b146101b457600080fd5b80638da5cb5b1461015e578063991f65431461017957600080fd5b80624b61bb146100a7578063181f5a77146100cf5780631d7a74a01461011857806379ba509714610154575b600080fd5b6100ba6100b5366004610b32565b6101c7565b60405190151581526020015b60405180910390f35b61010b6040518060400160405280601e81526020017f416e793245564d546f6c6c4f666652616d70526f7574657220312e302e30000081525081565b6040516100c69190610bd6565b6100ba610126366004610c09565b6001600160a01b03166000908152600260205260409020546c01000000000000000000000000900460ff1690565b61015c610329565b005b6000546040516001600160a01b0390911681526020016100c6565b61015c610187366004610c09565b61040c565b610194610727565b6040516100c69190610c2b565b61015c6101af366004610c09565b610789565b61015c6101c2366004610c09565b610965565b336000818152600260205260408120549091906c01000000000000000000000000900460ff1661022a576040517fa2c8bfb60000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6000633015b91c60e01b876040516024016102459190610d5f565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091529050856102bf576102b88585600084610979565b925061031f565b836001600160a01b0316816040516102d79190610e55565b6000604051808303816000865af19150503d8060008114610314576040519150601f19603f3d011682016040523d82523d6000602084013e610319565b606091505b50909350505b5050949350505050565b6001546001600160a01b0316331461039d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610221565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6104146109c5565b6003546000819003610452576040517f22babb3200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382166000908152600260209081526040918290208251808401909352546bffffffffffffffffffffffff811683526c01000000000000000000000000900460ff1615159082018190526104e4576040517f8c97f1220000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401610221565b600060036104f3600185610e71565b8154811061050357610503610eaf565b6000918252602090912001548251600380546001600160a01b039093169350916bffffffffffffffffffffffff90911690811061054257610542610eaf565b6000918252602090912001546001600160a01b03166003610564600186610e71565b8154811061057457610574610eaf565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600383600001516bffffffffffffffffffffffff16815481106105c8576105c8610eaf565b600091825260208083209190910180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0394851617905584519284168252600290526040902080547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff909216919091179055600380548061066257610662610ede565b6000828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690559092019092556001600160a01b0386168083526002909152604080832080547fffffffffffffffffffffffffffffffffffffff000000000000000000000000001690555190917fcf91daec21e3510e2f2aea4b09d08c235d5c6844980be709f282ef591dbf420c91a250505050565b6060600380548060200260200160405190810160405280929190818152602001828054801561077f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610761575b5050505050905090565b6107916109c5565b6001600160a01b0381166107d1576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381166000908152600260209081526040918290208251808401909352546bffffffffffffffffffffffff811683526c01000000000000000000000000900460ff1615801591830191909152610865576040517f3a4406b50000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610221565b60016020808301828152600380546bffffffffffffffffffffffff90811686526001600160a01b03871660008181526002909552604080862088518154965115156c01000000000000000000000000027fffffffffffffffffffffffffffffffffffffff0000000000000000000000000090971694169390931794909417909155815494850182559083527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b90930180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055517f78f53b26906785548b265fa08f4197f9f3fff73fe0d504d30400aacb527f4ce09190a25050565b61096d6109c5565b61097681610a3b565b50565b60005a61138881101561098b57600080fd5b6113888103905085604082048203116109a357600080fd5b50833b6109af57600080fd5b60008083516020850186888af195945050505050565b6000546001600160a01b03163314610a39576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610221565b565b336001600160a01b03821603610aad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610221565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80356001600160a01b0381168114610b2d57600080fd5b919050565b60008060008060808587031215610b4857600080fd5b843567ffffffffffffffff811115610b5f57600080fd5b850160808188031215610b7157600080fd5b935060208501358015158114610b8657600080fd5b925060408501359150610b9b60608601610b16565b905092959194509250565b60005b83811015610bc1578181015183820152602001610ba9565b83811115610bd0576000848401525b50505050565b6020815260008251806020840152610bf5816040850160208701610ba6565b601f01601f19169190910160400192915050565b600060208284031215610c1b57600080fd5b610c2482610b16565b9392505050565b6020808252825182820181905260009190848201906040850190845b81811015610c6c5783516001600160a01b031683529284019291840191600101610c47565b50909695505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610cad57600080fd5b830160208101925035905067ffffffffffffffff811115610ccd57600080fd5b803603821315610cdc57600080fd5b9250929050565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b8183526000602080850194508260005b85811015610d54576001600160a01b03610d3783610b16565b168752818301358388015260409687019690910190600101610d1e565b509495945050505050565b602081526000823567ffffffffffffffff808216808314610d7f57600080fd5b80602086015250610d936020860186610c78565b925060806040860152610daa60a086018483610ce3565b925050610dba6040860186610c78565b601f1980878603016060880152610dd2858385610ce3565b9450606088013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018312610e0b57600080fd5b60209288019283019235915083821115610e2457600080fd5b8160061b3603831315610e3657600080fd5b868503016080870152610e4a848284610d0e565b979650505050505050565b60008251610e67818460208701610ba6565b9190910192915050565b600082821015610eaa577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c634300080f000a",
}

var Any2EVMTollOffRampRouterABI = Any2EVMTollOffRampRouterMetaData.ABI

var Any2EVMTollOffRampRouterBin = Any2EVMTollOffRampRouterMetaData.Bin

func DeployAny2EVMTollOffRampRouter(auth *bind.TransactOpts, backend bind.ContractBackend, offRamps []common.Address) (common.Address, *types.Transaction, *Any2EVMTollOffRampRouter, error) {
	parsed, err := Any2EVMTollOffRampRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Any2EVMTollOffRampRouterBin), backend, offRamps)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Any2EVMTollOffRampRouter{Any2EVMTollOffRampRouterCaller: Any2EVMTollOffRampRouterCaller{contract: contract}, Any2EVMTollOffRampRouterTransactor: Any2EVMTollOffRampRouterTransactor{contract: contract}, Any2EVMTollOffRampRouterFilterer: Any2EVMTollOffRampRouterFilterer{contract: contract}}, nil
}

type Any2EVMTollOffRampRouter struct {
	address common.Address
	abi     abi.ABI
	Any2EVMTollOffRampRouterCaller
	Any2EVMTollOffRampRouterTransactor
	Any2EVMTollOffRampRouterFilterer
}

type Any2EVMTollOffRampRouterCaller struct {
	contract *bind.BoundContract
}

type Any2EVMTollOffRampRouterTransactor struct {
	contract *bind.BoundContract
}

type Any2EVMTollOffRampRouterFilterer struct {
	contract *bind.BoundContract
}

type Any2EVMTollOffRampRouterSession struct {
	Contract     *Any2EVMTollOffRampRouter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type Any2EVMTollOffRampRouterCallerSession struct {
	Contract *Any2EVMTollOffRampRouterCaller
	CallOpts bind.CallOpts
}

type Any2EVMTollOffRampRouterTransactorSession struct {
	Contract     *Any2EVMTollOffRampRouterTransactor
	TransactOpts bind.TransactOpts
}

type Any2EVMTollOffRampRouterRaw struct {
	Contract *Any2EVMTollOffRampRouter
}

type Any2EVMTollOffRampRouterCallerRaw struct {
	Contract *Any2EVMTollOffRampRouterCaller
}

type Any2EVMTollOffRampRouterTransactorRaw struct {
	Contract *Any2EVMTollOffRampRouterTransactor
}

func NewAny2EVMTollOffRampRouter(address common.Address, backend bind.ContractBackend) (*Any2EVMTollOffRampRouter, error) {
	abi, err := abi.JSON(strings.NewReader(Any2EVMTollOffRampRouterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindAny2EVMTollOffRampRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouter{address: address, abi: abi, Any2EVMTollOffRampRouterCaller: Any2EVMTollOffRampRouterCaller{contract: contract}, Any2EVMTollOffRampRouterTransactor: Any2EVMTollOffRampRouterTransactor{contract: contract}, Any2EVMTollOffRampRouterFilterer: Any2EVMTollOffRampRouterFilterer{contract: contract}}, nil
}

func NewAny2EVMTollOffRampRouterCaller(address common.Address, caller bind.ContractCaller) (*Any2EVMTollOffRampRouterCaller, error) {
	contract, err := bindAny2EVMTollOffRampRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouterCaller{contract: contract}, nil
}

func NewAny2EVMTollOffRampRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*Any2EVMTollOffRampRouterTransactor, error) {
	contract, err := bindAny2EVMTollOffRampRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouterTransactor{contract: contract}, nil
}

func NewAny2EVMTollOffRampRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*Any2EVMTollOffRampRouterFilterer, error) {
	contract, err := bindAny2EVMTollOffRampRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouterFilterer{contract: contract}, nil
}

func bindAny2EVMTollOffRampRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Any2EVMTollOffRampRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Any2EVMTollOffRampRouter.Contract.Any2EVMTollOffRampRouterCaller.contract.Call(opts, result, method, params...)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.Any2EVMTollOffRampRouterTransactor.contract.Transfer(opts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.Any2EVMTollOffRampRouterTransactor.contract.Transact(opts, method, params...)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Any2EVMTollOffRampRouter.Contract.contract.Call(opts, result, method, params...)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.contract.Transfer(opts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.contract.Transact(opts, method, params...)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCaller) GetOffRamps(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampRouter.contract.Call(opts, &out, "getOffRamps")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) GetOffRamps() ([]common.Address, error) {
	return _Any2EVMTollOffRampRouter.Contract.GetOffRamps(&_Any2EVMTollOffRampRouter.CallOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCallerSession) GetOffRamps() ([]common.Address, error) {
	return _Any2EVMTollOffRampRouter.Contract.GetOffRamps(&_Any2EVMTollOffRampRouter.CallOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCaller) IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampRouter.contract.Call(opts, &out, "isOffRamp", offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _Any2EVMTollOffRampRouter.Contract.IsOffRamp(&_Any2EVMTollOffRampRouter.CallOpts, offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCallerSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _Any2EVMTollOffRampRouter.Contract.IsOffRamp(&_Any2EVMTollOffRampRouter.CallOpts, offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) Owner() (common.Address, error) {
	return _Any2EVMTollOffRampRouter.Contract.Owner(&_Any2EVMTollOffRampRouter.CallOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCallerSession) Owner() (common.Address, error) {
	return _Any2EVMTollOffRampRouter.Contract.Owner(&_Any2EVMTollOffRampRouter.CallOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampRouter.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) TypeAndVersion() (string, error) {
	return _Any2EVMTollOffRampRouter.Contract.TypeAndVersion(&_Any2EVMTollOffRampRouter.CallOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCallerSession) TypeAndVersion() (string, error) {
	return _Any2EVMTollOffRampRouter.Contract.TypeAndVersion(&_Any2EVMTollOffRampRouter.CallOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.contract.Transact(opts, "acceptOwnership")
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.AcceptOwnership(&_Any2EVMTollOffRampRouter.TransactOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.AcceptOwnership(&_Any2EVMTollOffRampRouter.TransactOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactor) AddOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.contract.Transact(opts, "addOffRamp", offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) AddOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.AddOffRamp(&_Any2EVMTollOffRampRouter.TransactOpts, offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactorSession) AddOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.AddOffRamp(&_Any2EVMTollOffRampRouter.TransactOpts, offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactor) RemoveOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.contract.Transact(opts, "removeOffRamp", offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) RemoveOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.RemoveOffRamp(&_Any2EVMTollOffRampRouter.TransactOpts, offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactorSession) RemoveOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.RemoveOffRamp(&_Any2EVMTollOffRampRouter.TransactOpts, offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactor) RouteMessage(opts *bind.TransactOpts, message CommonAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.contract.Transact(opts, "routeMessage", message, manualExecution, gasLimit, receiver)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) RouteMessage(message CommonAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.RouteMessage(&_Any2EVMTollOffRampRouter.TransactOpts, message, manualExecution, gasLimit, receiver)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactorSession) RouteMessage(message CommonAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.RouteMessage(&_Any2EVMTollOffRampRouter.TransactOpts, message, manualExecution, gasLimit, receiver)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.contract.Transact(opts, "transferOwnership", to)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.TransferOwnership(&_Any2EVMTollOffRampRouter.TransactOpts, to)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.TransferOwnership(&_Any2EVMTollOffRampRouter.TransactOpts, to)
}

type Any2EVMTollOffRampRouterOffRampAddedIterator struct {
	Event *Any2EVMTollOffRampRouterOffRampAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampRouterOffRampAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampRouterOffRampAdded)
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
		it.Event = new(Any2EVMTollOffRampRouterOffRampAdded)
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

func (it *Any2EVMTollOffRampRouterOffRampAddedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampRouterOffRampAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampRouterOffRampAdded struct {
	OffRamp common.Address
	Raw     types.Log
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) FilterOffRampAdded(opts *bind.FilterOpts, offRamp []common.Address) (*Any2EVMTollOffRampRouterOffRampAddedIterator, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.FilterLogs(opts, "OffRampAdded", offRampRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouterOffRampAddedIterator{contract: _Any2EVMTollOffRampRouter.contract, event: "OffRampAdded", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOffRampAdded, offRamp []common.Address) (event.Subscription, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.WatchLogs(opts, "OffRampAdded", offRampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampRouterOffRampAdded)
				if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OffRampAdded", log); err != nil {
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

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) ParseOffRampAdded(log types.Log) (*Any2EVMTollOffRampRouterOffRampAdded, error) {
	event := new(Any2EVMTollOffRampRouterOffRampAdded)
	if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OffRampAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampRouterOffRampRemovedIterator struct {
	Event *Any2EVMTollOffRampRouterOffRampRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampRouterOffRampRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampRouterOffRampRemoved)
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
		it.Event = new(Any2EVMTollOffRampRouterOffRampRemoved)
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

func (it *Any2EVMTollOffRampRouterOffRampRemovedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampRouterOffRampRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampRouterOffRampRemoved struct {
	OffRamp common.Address
	Raw     types.Log
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) FilterOffRampRemoved(opts *bind.FilterOpts, offRamp []common.Address) (*Any2EVMTollOffRampRouterOffRampRemovedIterator, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.FilterLogs(opts, "OffRampRemoved", offRampRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouterOffRampRemovedIterator{contract: _Any2EVMTollOffRampRouter.contract, event: "OffRampRemoved", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOffRampRemoved, offRamp []common.Address) (event.Subscription, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.WatchLogs(opts, "OffRampRemoved", offRampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampRouterOffRampRemoved)
				if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OffRampRemoved", log); err != nil {
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

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) ParseOffRampRemoved(log types.Log) (*Any2EVMTollOffRampRouterOffRampRemoved, error) {
	event := new(Any2EVMTollOffRampRouterOffRampRemoved)
	if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OffRampRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampRouterOwnershipTransferRequestedIterator struct {
	Event *Any2EVMTollOffRampRouterOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampRouterOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampRouterOwnershipTransferRequested)
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
		it.Event = new(Any2EVMTollOffRampRouterOwnershipTransferRequested)
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

func (it *Any2EVMTollOffRampRouterOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampRouterOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampRouterOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampRouterOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouterOwnershipTransferRequestedIterator{contract: _Any2EVMTollOffRampRouter.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampRouterOwnershipTransferRequested)
				if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) ParseOwnershipTransferRequested(log types.Log) (*Any2EVMTollOffRampRouterOwnershipTransferRequested, error) {
	event := new(Any2EVMTollOffRampRouterOwnershipTransferRequested)
	if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampRouterOwnershipTransferredIterator struct {
	Event *Any2EVMTollOffRampRouterOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampRouterOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampRouterOwnershipTransferred)
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
		it.Event = new(Any2EVMTollOffRampRouterOwnershipTransferred)
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

func (it *Any2EVMTollOffRampRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampRouterOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampRouterOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouterOwnershipTransferredIterator{contract: _Any2EVMTollOffRampRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampRouterOwnershipTransferred)
				if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) ParseOwnershipTransferred(log types.Log) (*Any2EVMTollOffRampRouterOwnershipTransferred, error) {
	event := new(Any2EVMTollOffRampRouterOwnershipTransferred)
	if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Any2EVMTollOffRampRouter.abi.Events["OffRampAdded"].ID:
		return _Any2EVMTollOffRampRouter.ParseOffRampAdded(log)
	case _Any2EVMTollOffRampRouter.abi.Events["OffRampRemoved"].ID:
		return _Any2EVMTollOffRampRouter.ParseOffRampRemoved(log)
	case _Any2EVMTollOffRampRouter.abi.Events["OwnershipTransferRequested"].ID:
		return _Any2EVMTollOffRampRouter.ParseOwnershipTransferRequested(log)
	case _Any2EVMTollOffRampRouter.abi.Events["OwnershipTransferred"].ID:
		return _Any2EVMTollOffRampRouter.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (Any2EVMTollOffRampRouterOffRampAdded) Topic() common.Hash {
	return common.HexToHash("0x78f53b26906785548b265fa08f4197f9f3fff73fe0d504d30400aacb527f4ce0")
}

func (Any2EVMTollOffRampRouterOffRampRemoved) Topic() common.Hash {
	return common.HexToHash("0xcf91daec21e3510e2f2aea4b09d08c235d5c6844980be709f282ef591dbf420c")
}

func (Any2EVMTollOffRampRouterOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (Any2EVMTollOffRampRouterOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouter) Address() common.Address {
	return _Any2EVMTollOffRampRouter.address
}

type Any2EVMTollOffRampRouterInterface interface {
	GetOffRamps(opts *bind.CallOpts) ([]common.Address, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error)

	RemoveOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error)

	RouteMessage(opts *bind.TransactOpts, message CommonAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOffRampAdded(opts *bind.FilterOpts, offRamp []common.Address) (*Any2EVMTollOffRampRouterOffRampAddedIterator, error)

	WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOffRampAdded, offRamp []common.Address) (event.Subscription, error)

	ParseOffRampAdded(log types.Log) (*Any2EVMTollOffRampRouterOffRampAdded, error)

	FilterOffRampRemoved(opts *bind.FilterOpts, offRamp []common.Address) (*Any2EVMTollOffRampRouterOffRampRemovedIterator, error)

	WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOffRampRemoved, offRamp []common.Address) (event.Subscription, error)

	ParseOffRampRemoved(log types.Log) (*Any2EVMTollOffRampRouterOffRampRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampRouterOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*Any2EVMTollOffRampRouterOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampRouterOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*Any2EVMTollOffRampRouterOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
