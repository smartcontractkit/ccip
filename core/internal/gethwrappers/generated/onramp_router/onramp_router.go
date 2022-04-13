// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package onramp_router

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
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated"
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

type CCIPMessagePayload struct {
	Tokens             []common.Address
	Amounts            []*big.Int
	DestinationChainId *big.Int
	Receiver           common.Address
	Executor           common.Address
	Data               []byte
	Options            []byte
}

var OnRampRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"contractOnRampInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.MessagePayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"requestCrossChainSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"setOnRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b611148806101576000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063c8658c1c1161005b578063c8658c1c14610166578063d8a98f8c14610187578063f1927cae146101bd578063f2fde38b146101d057600080fd5b8063181f5a771461008d5780635221c1f0146100d557806379ba50971461011d5780638da5cb5b14610127575b600080fd5b604080518082018252601281527f4f6e52616d70526f7574657220302e302e310000000000000000000000000000602082015290516100cc9190610bda565b60405180910390f35b61010d6100e3366004610c2b565b60009081526002602052604090205473ffffffffffffffffffffffffffffffffffffffff16151590565b60405190151581526020016100cc565b6101256101e3565b005b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100cc565b610179610174366004610c44565b6102e5565b6040519081526020016100cc565b610141610195366004610c2b565b60009081526002602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b6101256101cb366004610ca1565b610585565b6101256101de366004610ce1565b61068a565b60015473ffffffffffffffffffffffffffffffffffffffff163314610269576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6040808201356000908152600260205290812054339073ffffffffffffffffffffffffffffffffffffffff168061034e57604080517f45abe4ae000000000000000000000000000000000000000000000000000000008152908501356004820152602401610260565b61035b6020850185610cfe565b90506103678580610cfe565b9050146103a0576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b6103ad8580610cfe565b90508110156104e55760006103c28680610cfe565b838181106103d2576103d2610d6d565b90506020020160208101906103e79190610ce1565b905060006103f86020880188610cfe565b8481811061040857610408610d6d565b60200291909101359150610436905073ffffffffffffffffffffffffffffffffffffffff831686308461069e565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820183905283169063095ea7b3906044016020604051808303816000875af11580156104ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104cf9190610d9c565b50505080806104dd90610dbe565b9150506103a3565b506040517f625f9e1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82169063625f9e199061053a9087908690600401610f80565b6020604051808303816000875af1158015610559573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057d9190611106565b949350505050565b61058d610739565b60008281526002602052604090205473ffffffffffffffffffffffffffffffffffffffff80831691160361060c576040517fe31de3b20000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff82166024820152604401610260565b60008281526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091559051909184917f4b680ef9fa79bb5f36e7559d7b33fd57a8336f78cc120c8cd93333b5ade624cb9190a35050565b610692610739565b61069b816107bc565b50565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526107339085906108b1565b50505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146107ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610260565b565b3373ffffffffffffffffffffffffffffffffffffffff82160361083b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610260565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000610913826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166109c29092919063ffffffff16565b8051909150156109bd57808060200190518101906109319190610d9c565b6109bd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610260565b505050565b60606109d184846000856109db565b90505b9392505050565b606082471015610a6d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610260565b843b610ad5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610260565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610afe919061111f565b60006040518083038185875af1925050503d8060008114610b3b576040519150601f19603f3d011682016040523d82523d6000602084013e610b40565b606091505b5091509150610b50828286610b5b565b979650505050505050565b60608315610b6a5750816109d4565b825115610b7a5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102609190610bda565b60005b83811015610bc9578181015183820152602001610bb1565b838111156107335750506000910152565b6020815260008251806020840152610bf9816040850160208701610bae565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b600060208284031215610c3d57600080fd5b5035919050565b600060208284031215610c5657600080fd5b813567ffffffffffffffff811115610c6d57600080fd5b820160e081850312156109d457600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461069b57600080fd5b60008060408385031215610cb457600080fd5b823591506020830135610cc681610c7f565b809150509250929050565b8035610cdc81610c7f565b919050565b600060208284031215610cf357600080fd5b81356109d481610c7f565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610d3357600080fd5b83018035915067ffffffffffffffff821115610d4e57600080fd5b6020019150600581901b3603821315610d6657600080fd5b9250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215610dae57600080fd5b815180151581146109d457600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610e16577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610e5257600080fd5b830160208101925035905067ffffffffffffffff811115610e7257600080fd5b8060051b3603831315610d6657600080fd5b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115610eb657600080fd5b8260051b8083602087013760009401602001938452509192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610f0857600080fd5b830160208101925035905067ffffffffffffffff811115610f2857600080fd5b803603831315610d6657600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6040815260006101208201610f958586610e1d565b60e06040860152918290529060009061014085015b81831015610fe8578335610fbd81610c7f565b73ffffffffffffffffffffffffffffffffffffffff1681526020938401936001939093019201610faa565b610ff56020890189610e1d565b945092507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0915081868203016060870152611031818585610e84565b9350506040870135608086015261104a60608801610cd1565b73ffffffffffffffffffffffffffffffffffffffff811660a0870152915061107460808801610cd1565b73ffffffffffffffffffffffffffffffffffffffff811660c0870152915061109f60a0880188610ed3565b9250818685030160e08701526110b6848483610f37565b9350506110c660c0880188610ed3565b925081868503016101008701526110de848483610f37565b9450505050506109d4602083018473ffffffffffffffffffffffffffffffffffffffff169052565b60006020828403121561111857600080fd5b5051919050565b60008251611131818460208701610bae565b919091019291505056fea164736f6c634300080d000a",
}

var OnRampRouterABI = OnRampRouterMetaData.ABI

var OnRampRouterBin = OnRampRouterMetaData.Bin

func DeployOnRampRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OnRampRouter, error) {
	parsed, err := OnRampRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnRampRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnRampRouter{OnRampRouterCaller: OnRampRouterCaller{contract: contract}, OnRampRouterTransactor: OnRampRouterTransactor{contract: contract}, OnRampRouterFilterer: OnRampRouterFilterer{contract: contract}}, nil
}

type OnRampRouter struct {
	address common.Address
	abi     abi.ABI
	OnRampRouterCaller
	OnRampRouterTransactor
	OnRampRouterFilterer
}

type OnRampRouterCaller struct {
	contract *bind.BoundContract
}

type OnRampRouterTransactor struct {
	contract *bind.BoundContract
}

type OnRampRouterFilterer struct {
	contract *bind.BoundContract
}

type OnRampRouterSession struct {
	Contract     *OnRampRouter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OnRampRouterCallerSession struct {
	Contract *OnRampRouterCaller
	CallOpts bind.CallOpts
}

type OnRampRouterTransactorSession struct {
	Contract     *OnRampRouterTransactor
	TransactOpts bind.TransactOpts
}

type OnRampRouterRaw struct {
	Contract *OnRampRouter
}

type OnRampRouterCallerRaw struct {
	Contract *OnRampRouterCaller
}

type OnRampRouterTransactorRaw struct {
	Contract *OnRampRouterTransactor
}

func NewOnRampRouter(address common.Address, backend bind.ContractBackend) (*OnRampRouter, error) {
	abi, err := abi.JSON(strings.NewReader(OnRampRouterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOnRampRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnRampRouter{address: address, abi: abi, OnRampRouterCaller: OnRampRouterCaller{contract: contract}, OnRampRouterTransactor: OnRampRouterTransactor{contract: contract}, OnRampRouterFilterer: OnRampRouterFilterer{contract: contract}}, nil
}

func NewOnRampRouterCaller(address common.Address, caller bind.ContractCaller) (*OnRampRouterCaller, error) {
	contract, err := bindOnRampRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampRouterCaller{contract: contract}, nil
}

func NewOnRampRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*OnRampRouterTransactor, error) {
	contract, err := bindOnRampRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampRouterTransactor{contract: contract}, nil
}

func NewOnRampRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*OnRampRouterFilterer, error) {
	contract, err := bindOnRampRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnRampRouterFilterer{contract: contract}, nil
}

func bindOnRampRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OnRampRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_OnRampRouter *OnRampRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRampRouter.Contract.OnRampRouterCaller.contract.Call(opts, result, method, params...)
}

func (_OnRampRouter *OnRampRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRampRouter.Contract.OnRampRouterTransactor.contract.Transfer(opts)
}

func (_OnRampRouter *OnRampRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRampRouter.Contract.OnRampRouterTransactor.contract.Transact(opts, method, params...)
}

func (_OnRampRouter *OnRampRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRampRouter.Contract.contract.Call(opts, result, method, params...)
}

func (_OnRampRouter *OnRampRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRampRouter.Contract.contract.Transfer(opts)
}

func (_OnRampRouter *OnRampRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRampRouter.Contract.contract.Transact(opts, method, params...)
}

func (_OnRampRouter *OnRampRouterCaller) GetOnRamp(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OnRampRouter.contract.Call(opts, &out, "getOnRamp", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRampRouter *OnRampRouterSession) GetOnRamp(chainId *big.Int) (common.Address, error) {
	return _OnRampRouter.Contract.GetOnRamp(&_OnRampRouter.CallOpts, chainId)
}

func (_OnRampRouter *OnRampRouterCallerSession) GetOnRamp(chainId *big.Int) (common.Address, error) {
	return _OnRampRouter.Contract.GetOnRamp(&_OnRampRouter.CallOpts, chainId)
}

func (_OnRampRouter *OnRampRouterCaller) IsChainSupported(opts *bind.CallOpts, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _OnRampRouter.contract.Call(opts, &out, "isChainSupported", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OnRampRouter *OnRampRouterSession) IsChainSupported(chainId *big.Int) (bool, error) {
	return _OnRampRouter.Contract.IsChainSupported(&_OnRampRouter.CallOpts, chainId)
}

func (_OnRampRouter *OnRampRouterCallerSession) IsChainSupported(chainId *big.Int) (bool, error) {
	return _OnRampRouter.Contract.IsChainSupported(&_OnRampRouter.CallOpts, chainId)
}

func (_OnRampRouter *OnRampRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnRampRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRampRouter *OnRampRouterSession) Owner() (common.Address, error) {
	return _OnRampRouter.Contract.Owner(&_OnRampRouter.CallOpts)
}

func (_OnRampRouter *OnRampRouterCallerSession) Owner() (common.Address, error) {
	return _OnRampRouter.Contract.Owner(&_OnRampRouter.CallOpts)
}

func (_OnRampRouter *OnRampRouterCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OnRampRouter.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OnRampRouter *OnRampRouterSession) TypeAndVersion() (string, error) {
	return _OnRampRouter.Contract.TypeAndVersion(&_OnRampRouter.CallOpts)
}

func (_OnRampRouter *OnRampRouterCallerSession) TypeAndVersion() (string, error) {
	return _OnRampRouter.Contract.TypeAndVersion(&_OnRampRouter.CallOpts)
}

func (_OnRampRouter *OnRampRouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRampRouter.contract.Transact(opts, "acceptOwnership")
}

func (_OnRampRouter *OnRampRouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRampRouter.Contract.AcceptOwnership(&_OnRampRouter.TransactOpts)
}

func (_OnRampRouter *OnRampRouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRampRouter.Contract.AcceptOwnership(&_OnRampRouter.TransactOpts)
}

func (_OnRampRouter *OnRampRouterTransactor) RequestCrossChainSend(opts *bind.TransactOpts, payload CCIPMessagePayload) (*types.Transaction, error) {
	return _OnRampRouter.contract.Transact(opts, "requestCrossChainSend", payload)
}

func (_OnRampRouter *OnRampRouterSession) RequestCrossChainSend(payload CCIPMessagePayload) (*types.Transaction, error) {
	return _OnRampRouter.Contract.RequestCrossChainSend(&_OnRampRouter.TransactOpts, payload)
}

func (_OnRampRouter *OnRampRouterTransactorSession) RequestCrossChainSend(payload CCIPMessagePayload) (*types.Transaction, error) {
	return _OnRampRouter.Contract.RequestCrossChainSend(&_OnRampRouter.TransactOpts, payload)
}

func (_OnRampRouter *OnRampRouterTransactor) SetOnRamp(opts *bind.TransactOpts, chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
	return _OnRampRouter.contract.Transact(opts, "setOnRamp", chainId, onRamp)
}

func (_OnRampRouter *OnRampRouterSession) SetOnRamp(chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
	return _OnRampRouter.Contract.SetOnRamp(&_OnRampRouter.TransactOpts, chainId, onRamp)
}

func (_OnRampRouter *OnRampRouterTransactorSession) SetOnRamp(chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
	return _OnRampRouter.Contract.SetOnRamp(&_OnRampRouter.TransactOpts, chainId, onRamp)
}

func (_OnRampRouter *OnRampRouterTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OnRampRouter.contract.Transact(opts, "transferOwnership", to)
}

func (_OnRampRouter *OnRampRouterSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRampRouter.Contract.TransferOwnership(&_OnRampRouter.TransactOpts, to)
}

func (_OnRampRouter *OnRampRouterTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRampRouter.Contract.TransferOwnership(&_OnRampRouter.TransactOpts, to)
}

type OnRampRouterOnRampSetIterator struct {
	Event *OnRampRouterOnRampSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampRouterOnRampSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampRouterOnRampSet)
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
		it.Event = new(OnRampRouterOnRampSet)
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

func (it *OnRampRouterOnRampSetIterator) Error() error {
	return it.fail
}

func (it *OnRampRouterOnRampSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampRouterOnRampSet struct {
	ChainId *big.Int
	OnRamp  common.Address
	Raw     types.Log
}

func (_OnRampRouter *OnRampRouterFilterer) FilterOnRampSet(opts *bind.FilterOpts, chainId []*big.Int, onRamp []common.Address) (*OnRampRouterOnRampSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var onRampRule []interface{}
	for _, onRampItem := range onRamp {
		onRampRule = append(onRampRule, onRampItem)
	}

	logs, sub, err := _OnRampRouter.contract.FilterLogs(opts, "OnRampSet", chainIdRule, onRampRule)
	if err != nil {
		return nil, err
	}
	return &OnRampRouterOnRampSetIterator{contract: _OnRampRouter.contract, event: "OnRampSet", logs: logs, sub: sub}, nil
}

func (_OnRampRouter *OnRampRouterFilterer) WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *OnRampRouterOnRampSet, chainId []*big.Int, onRamp []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var onRampRule []interface{}
	for _, onRampItem := range onRamp {
		onRampRule = append(onRampRule, onRampItem)
	}

	logs, sub, err := _OnRampRouter.contract.WatchLogs(opts, "OnRampSet", chainIdRule, onRampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampRouterOnRampSet)
				if err := _OnRampRouter.contract.UnpackLog(event, "OnRampSet", log); err != nil {
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

func (_OnRampRouter *OnRampRouterFilterer) ParseOnRampSet(log types.Log) (*OnRampRouterOnRampSet, error) {
	event := new(OnRampRouterOnRampSet)
	if err := _OnRampRouter.contract.UnpackLog(event, "OnRampSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampRouterOwnershipTransferRequestedIterator struct {
	Event *OnRampRouterOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampRouterOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampRouterOwnershipTransferRequested)
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
		it.Event = new(OnRampRouterOwnershipTransferRequested)
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

func (it *OnRampRouterOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OnRampRouterOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampRouterOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRampRouter *OnRampRouterFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampRouterOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampRouter.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampRouterOwnershipTransferRequestedIterator{contract: _OnRampRouter.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OnRampRouter *OnRampRouterFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampRouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampRouter.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampRouterOwnershipTransferRequested)
				if err := _OnRampRouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OnRampRouter *OnRampRouterFilterer) ParseOwnershipTransferRequested(log types.Log) (*OnRampRouterOwnershipTransferRequested, error) {
	event := new(OnRampRouterOwnershipTransferRequested)
	if err := _OnRampRouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampRouterOwnershipTransferredIterator struct {
	Event *OnRampRouterOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampRouterOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampRouterOwnershipTransferred)
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
		it.Event = new(OnRampRouterOwnershipTransferred)
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

func (it *OnRampRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OnRampRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampRouterOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRampRouter *OnRampRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampRouterOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampRouter.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampRouterOwnershipTransferredIterator{contract: _OnRampRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OnRampRouter *OnRampRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampRouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRampRouter.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampRouterOwnershipTransferred)
				if err := _OnRampRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OnRampRouter *OnRampRouterFilterer) ParseOwnershipTransferred(log types.Log) (*OnRampRouterOwnershipTransferred, error) {
	event := new(OnRampRouterOwnershipTransferred)
	if err := _OnRampRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_OnRampRouter *OnRampRouter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OnRampRouter.abi.Events["OnRampSet"].ID:
		return _OnRampRouter.ParseOnRampSet(log)
	case _OnRampRouter.abi.Events["OwnershipTransferRequested"].ID:
		return _OnRampRouter.ParseOwnershipTransferRequested(log)
	case _OnRampRouter.abi.Events["OwnershipTransferred"].ID:
		return _OnRampRouter.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (OnRampRouterOnRampSet) Topic() common.Hash {
	return common.HexToHash("0x4b680ef9fa79bb5f36e7559d7b33fd57a8336f78cc120c8cd93333b5ade624cb")
}

func (OnRampRouterOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OnRampRouterOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_OnRampRouter *OnRampRouter) Address() common.Address {
	return _OnRampRouter.address
}

type OnRampRouterInterface interface {
	GetOnRamp(opts *bind.CallOpts, chainId *big.Int) (common.Address, error)

	IsChainSupported(opts *bind.CallOpts, chainId *big.Int) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	RequestCrossChainSend(opts *bind.TransactOpts, payload CCIPMessagePayload) (*types.Transaction, error)

	SetOnRamp(opts *bind.TransactOpts, chainId *big.Int, onRamp common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOnRampSet(opts *bind.FilterOpts, chainId []*big.Int, onRamp []common.Address) (*OnRampRouterOnRampSetIterator, error)

	WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *OnRampRouterOnRampSet, chainId []*big.Int, onRamp []common.Address) (event.Subscription, error)

	ParseOnRampSet(log types.Log) (*OnRampRouterOnRampSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampRouterOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampRouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OnRampRouterOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampRouterOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampRouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OnRampRouterOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
