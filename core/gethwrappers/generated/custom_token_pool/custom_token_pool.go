// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package custom_token_pool

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

var CustomTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ExceedsTokenLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionsError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SynthBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SynthMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBaseOffRamp\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBaseOnRamp\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"isOnRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockOrBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"releaseOrMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBaseOffRamp\",\"name\":\"offRamp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"permission\",\"type\":\"bool\"}],\"name\":\"setOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBaseOnRamp\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"permission\",\"type\":\"bool\"}],\"name\":\"setOnRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610c08380380610c0883398101604081905261002f91610181565b8033806000816100865760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100b6576100b6816100d8565b50506001805460ff60a01b19169055506001600160a01b0316608052506101b1565b336001600160a01b038216036101305760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161007d565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561019357600080fd5b81516001600160a01b03811681146101aa57600080fd5b9392505050565b608051610a3c6101cc60003960006101340152610a3c6000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c80636f32b8721161008c5780638da5cb5b116100665780638da5cb5b14610215578063cc8af2e814610233578063ea6192a214610246578063f2fde38b1461025957600080fd5b80636f32b872146101cc57806379ba5097146102055780638456cb591461020d57600080fd5b8063503c2858116100bd578063503c2858146101835780635c975abb1461019657806369e946d4146101b957600080fd5b80631d7a74a0146100e457806321df0da7146101325780633f4ba83a14610179575b600080fd5b61011d6100f236600461097c565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205460ff1690565b60405190151581526020015b60405180910390f35b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610129565b61018161026c565b005b6101816101913660046109a0565b61027e565b60015474010000000000000000000000000000000000000000900460ff1661011d565b6101816101c73660046109ce565b610393565b61011d6101da36600461097c565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460ff1690565b6101816103f1565b6101816104d4565b60005473ffffffffffffffffffffffffffffffffffffffff16610154565b6101816102413660046109ce565b6104e4565b610181610254366004610a03565b610542565b61018161026736600461097c565b610653565b610274610667565b61027c6106ce565b565b60015474010000000000000000000000000000000000000000900460ff16156102ee5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b60005473ffffffffffffffffffffffffffffffffffffffff16331480159061032657503360009081526002602052604090205460ff16155b1561035d576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518181527f02992093bca69a36949677658a77d359b510dc6232c68f9f118f7c0127a1b1479060200160405180910390a150565b61039b610667565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260036020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b60015473ffffffffffffffffffffffffffffffffffffffff1633146104585760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016102e5565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6104dc610667565b61027c6107ad565b6104ec610667565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b60015474010000000000000000000000000000000000000000900460ff16156105ad5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016102e5565b60005473ffffffffffffffffffffffffffffffffffffffff1633148015906105e557503360009081526003602052604090205460ff16155b1561061c576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518181527fbb0b72e5f44e331506684da008a30e10d50658c29d8159f6c6ab40bf1e52e6009060200160405180910390a15050565b61065b610667565b6106648161087f565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461027c5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016102e5565b60015474010000000000000000000000000000000000000000900460ff166107385760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016102e5565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60015474010000000000000000000000000000000000000000900460ff16156108185760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016102e5565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586107833390565b3373ffffffffffffffffffffffffffffffffffffffff8216036108e45760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016102e5565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b73ffffffffffffffffffffffffffffffffffffffff8116811461066457600080fd5b60006020828403121561098e57600080fd5b81356109998161095a565b9392505050565b6000602082840312156109b257600080fd5b5035919050565b803580151581146109c957600080fd5b919050565b600080604083850312156109e157600080fd5b82356109ec8161095a565b91506109fa602084016109b9565b90509250929050565b60008060408385031215610a1657600080fd5b8235610a218161095a565b94602093909301359350505056fea164736f6c634300080f000a",
}

var CustomTokenPoolABI = CustomTokenPoolMetaData.ABI

var CustomTokenPoolBin = CustomTokenPoolMetaData.Bin

func DeployCustomTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *CustomTokenPool, error) {
	parsed, err := CustomTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CustomTokenPoolBin), backend, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CustomTokenPool{CustomTokenPoolCaller: CustomTokenPoolCaller{contract: contract}, CustomTokenPoolTransactor: CustomTokenPoolTransactor{contract: contract}, CustomTokenPoolFilterer: CustomTokenPoolFilterer{contract: contract}}, nil
}

type CustomTokenPool struct {
	address common.Address
	abi     abi.ABI
	CustomTokenPoolCaller
	CustomTokenPoolTransactor
	CustomTokenPoolFilterer
}

type CustomTokenPoolCaller struct {
	contract *bind.BoundContract
}

type CustomTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type CustomTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type CustomTokenPoolSession struct {
	Contract     *CustomTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CustomTokenPoolCallerSession struct {
	Contract *CustomTokenPoolCaller
	CallOpts bind.CallOpts
}

type CustomTokenPoolTransactorSession struct {
	Contract     *CustomTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type CustomTokenPoolRaw struct {
	Contract *CustomTokenPool
}

type CustomTokenPoolCallerRaw struct {
	Contract *CustomTokenPoolCaller
}

type CustomTokenPoolTransactorRaw struct {
	Contract *CustomTokenPoolTransactor
}

func NewCustomTokenPool(address common.Address, backend bind.ContractBackend) (*CustomTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(CustomTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCustomTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPool{address: address, abi: abi, CustomTokenPoolCaller: CustomTokenPoolCaller{contract: contract}, CustomTokenPoolTransactor: CustomTokenPoolTransactor{contract: contract}, CustomTokenPoolFilterer: CustomTokenPoolFilterer{contract: contract}}, nil
}

func NewCustomTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*CustomTokenPoolCaller, error) {
	contract, err := bindCustomTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolCaller{contract: contract}, nil
}

func NewCustomTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*CustomTokenPoolTransactor, error) {
	contract, err := bindCustomTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolTransactor{contract: contract}, nil
}

func NewCustomTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*CustomTokenPoolFilterer, error) {
	contract, err := bindCustomTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolFilterer{contract: contract}, nil
}

func bindCustomTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CustomTokenPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_CustomTokenPool *CustomTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustomTokenPool.Contract.CustomTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_CustomTokenPool *CustomTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.CustomTokenPoolTransactor.contract.Transfer(opts)
}

func (_CustomTokenPool *CustomTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.CustomTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_CustomTokenPool *CustomTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustomTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_CustomTokenPool *CustomTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.contract.Transfer(opts)
}

func (_CustomTokenPool *CustomTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_CustomTokenPool *CustomTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) GetToken() (common.Address, error) {
	return _CustomTokenPool.Contract.GetToken(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _CustomTokenPool.Contract.GetToken(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCaller) IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "isOffRamp", offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _CustomTokenPool.Contract.IsOffRamp(&_CustomTokenPool.CallOpts, offRamp)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _CustomTokenPool.Contract.IsOffRamp(&_CustomTokenPool.CallOpts, offRamp)
}

func (_CustomTokenPool *CustomTokenPoolCaller) IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "isOnRamp", onRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _CustomTokenPool.Contract.IsOnRamp(&_CustomTokenPool.CallOpts, onRamp)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _CustomTokenPool.Contract.IsOnRamp(&_CustomTokenPool.CallOpts, onRamp)
}

func (_CustomTokenPool *CustomTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) Owner() (common.Address, error) {
	return _CustomTokenPool.Contract.Owner(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) Owner() (common.Address, error) {
	return _CustomTokenPool.Contract.Owner(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CustomTokenPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CustomTokenPool *CustomTokenPoolSession) Paused() (bool, error) {
	return _CustomTokenPool.Contract.Paused(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolCallerSession) Paused() (bool, error) {
	return _CustomTokenPool.Contract.Paused(&_CustomTokenPool.CallOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_CustomTokenPool *CustomTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.AcceptOwnership(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.AcceptOwnership(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "lockOrBurn", amount)
}

func (_CustomTokenPool *CustomTokenPoolSession) LockOrBurn(amount *big.Int) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.LockOrBurn(&_CustomTokenPool.TransactOpts, amount)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) LockOrBurn(amount *big.Int) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.LockOrBurn(&_CustomTokenPool.TransactOpts, amount)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "pause")
}

func (_CustomTokenPool *CustomTokenPoolSession) Pause() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.Pause(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) Pause() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.Pause(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "releaseOrMint", arg0, amount)
}

func (_CustomTokenPool *CustomTokenPoolSession) ReleaseOrMint(arg0 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.ReleaseOrMint(&_CustomTokenPool.TransactOpts, arg0, amount)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) ReleaseOrMint(arg0 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.ReleaseOrMint(&_CustomTokenPool.TransactOpts, arg0, amount)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) SetOffRamp(opts *bind.TransactOpts, offRamp common.Address, permission bool) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "setOffRamp", offRamp, permission)
}

func (_CustomTokenPool *CustomTokenPoolSession) SetOffRamp(offRamp common.Address, permission bool) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.SetOffRamp(&_CustomTokenPool.TransactOpts, offRamp, permission)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) SetOffRamp(offRamp common.Address, permission bool) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.SetOffRamp(&_CustomTokenPool.TransactOpts, offRamp, permission)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) SetOnRamp(opts *bind.TransactOpts, onRamp common.Address, permission bool) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "setOnRamp", onRamp, permission)
}

func (_CustomTokenPool *CustomTokenPoolSession) SetOnRamp(onRamp common.Address, permission bool) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.SetOnRamp(&_CustomTokenPool.TransactOpts, onRamp, permission)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) SetOnRamp(onRamp common.Address, permission bool) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.SetOnRamp(&_CustomTokenPool.TransactOpts, onRamp, permission)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_CustomTokenPool *CustomTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.TransferOwnership(&_CustomTokenPool.TransactOpts, to)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.TransferOwnership(&_CustomTokenPool.TransactOpts, to)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "unpause")
}

func (_CustomTokenPool *CustomTokenPoolSession) Unpause() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.Unpause(&_CustomTokenPool.TransactOpts)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _CustomTokenPool.Contract.Unpause(&_CustomTokenPool.TransactOpts)
}

type CustomTokenPoolBurnedIterator struct {
	Event *CustomTokenPoolBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolBurned)
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
		it.Event = new(CustomTokenPoolBurned)
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

func (it *CustomTokenPoolBurnedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*CustomTokenPoolBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolBurnedIterator{contract: _CustomTokenPool.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolBurned)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseBurned(log types.Log) (*CustomTokenPoolBurned, error) {
	event := new(CustomTokenPoolBurned)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolLockedIterator struct {
	Event *CustomTokenPoolLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolLocked)
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
		it.Event = new(CustomTokenPoolLocked)
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

func (it *CustomTokenPoolLockedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*CustomTokenPoolLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolLockedIterator{contract: _CustomTokenPool.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolLocked)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseLocked(log types.Log) (*CustomTokenPoolLocked, error) {
	event := new(CustomTokenPoolLocked)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolMintedIterator struct {
	Event *CustomTokenPoolMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolMinted)
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
		it.Event = new(CustomTokenPoolMinted)
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

func (it *CustomTokenPoolMintedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*CustomTokenPoolMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolMintedIterator{contract: _CustomTokenPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolMinted)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseMinted(log types.Log) (*CustomTokenPoolMinted, error) {
	event := new(CustomTokenPoolMinted)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolOwnershipTransferRequestedIterator struct {
	Event *CustomTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolOwnershipTransferRequested)
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
		it.Event = new(CustomTokenPoolOwnershipTransferRequested)
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

func (it *CustomTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CustomTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolOwnershipTransferRequestedIterator{contract: _CustomTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolOwnershipTransferRequested)
				if err := _CustomTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*CustomTokenPoolOwnershipTransferRequested, error) {
	event := new(CustomTokenPoolOwnershipTransferRequested)
	if err := _CustomTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolOwnershipTransferredIterator struct {
	Event *CustomTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolOwnershipTransferred)
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
		it.Event = new(CustomTokenPoolOwnershipTransferred)
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

func (it *CustomTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CustomTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolOwnershipTransferredIterator{contract: _CustomTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolOwnershipTransferred)
				if err := _CustomTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*CustomTokenPoolOwnershipTransferred, error) {
	event := new(CustomTokenPoolOwnershipTransferred)
	if err := _CustomTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolPausedIterator struct {
	Event *CustomTokenPoolPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolPaused)
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
		it.Event = new(CustomTokenPoolPaused)
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

func (it *CustomTokenPoolPausedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*CustomTokenPoolPausedIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolPausedIterator{contract: _CustomTokenPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolPaused) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolPaused)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParsePaused(log types.Log) (*CustomTokenPoolPaused, error) {
	event := new(CustomTokenPoolPaused)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolReleasedIterator struct {
	Event *CustomTokenPoolReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolReleased)
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
		it.Event = new(CustomTokenPoolReleased)
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

func (it *CustomTokenPoolReleasedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*CustomTokenPoolReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolReleasedIterator{contract: _CustomTokenPool.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolReleased)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseReleased(log types.Log) (*CustomTokenPoolReleased, error) {
	event := new(CustomTokenPoolReleased)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolSynthBurnedIterator struct {
	Event *CustomTokenPoolSynthBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolSynthBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolSynthBurned)
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
		it.Event = new(CustomTokenPoolSynthBurned)
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

func (it *CustomTokenPoolSynthBurnedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolSynthBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolSynthBurned struct {
	Amount *big.Int
	Raw    types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterSynthBurned(opts *bind.FilterOpts) (*CustomTokenPoolSynthBurnedIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "SynthBurned")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolSynthBurnedIterator{contract: _CustomTokenPool.contract, event: "SynthBurned", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchSynthBurned(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolSynthBurned) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "SynthBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolSynthBurned)
				if err := _CustomTokenPool.contract.UnpackLog(event, "SynthBurned", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseSynthBurned(log types.Log) (*CustomTokenPoolSynthBurned, error) {
	event := new(CustomTokenPoolSynthBurned)
	if err := _CustomTokenPool.contract.UnpackLog(event, "SynthBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolSynthMintedIterator struct {
	Event *CustomTokenPoolSynthMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolSynthMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolSynthMinted)
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
		it.Event = new(CustomTokenPoolSynthMinted)
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

func (it *CustomTokenPoolSynthMintedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolSynthMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolSynthMinted struct {
	Amount *big.Int
	Raw    types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterSynthMinted(opts *bind.FilterOpts) (*CustomTokenPoolSynthMintedIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "SynthMinted")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolSynthMintedIterator{contract: _CustomTokenPool.contract, event: "SynthMinted", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchSynthMinted(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolSynthMinted) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "SynthMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolSynthMinted)
				if err := _CustomTokenPool.contract.UnpackLog(event, "SynthMinted", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseSynthMinted(log types.Log) (*CustomTokenPoolSynthMinted, error) {
	event := new(CustomTokenPoolSynthMinted)
	if err := _CustomTokenPool.contract.UnpackLog(event, "SynthMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolUnpausedIterator struct {
	Event *CustomTokenPoolUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolUnpaused)
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
		it.Event = new(CustomTokenPoolUnpaused)
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

func (it *CustomTokenPoolUnpausedIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CustomTokenPoolUnpausedIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolUnpausedIterator{contract: _CustomTokenPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolUnpaused)
				if err := _CustomTokenPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseUnpaused(log types.Log) (*CustomTokenPoolUnpaused, error) {
	event := new(CustomTokenPoolUnpaused)
	if err := _CustomTokenPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_CustomTokenPool *CustomTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CustomTokenPool.abi.Events["Burned"].ID:
		return _CustomTokenPool.ParseBurned(log)
	case _CustomTokenPool.abi.Events["Locked"].ID:
		return _CustomTokenPool.ParseLocked(log)
	case _CustomTokenPool.abi.Events["Minted"].ID:
		return _CustomTokenPool.ParseMinted(log)
	case _CustomTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _CustomTokenPool.ParseOwnershipTransferRequested(log)
	case _CustomTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _CustomTokenPool.ParseOwnershipTransferred(log)
	case _CustomTokenPool.abi.Events["Paused"].ID:
		return _CustomTokenPool.ParsePaused(log)
	case _CustomTokenPool.abi.Events["Released"].ID:
		return _CustomTokenPool.ParseReleased(log)
	case _CustomTokenPool.abi.Events["SynthBurned"].ID:
		return _CustomTokenPool.ParseSynthBurned(log)
	case _CustomTokenPool.abi.Events["SynthMinted"].ID:
		return _CustomTokenPool.ParseSynthMinted(log)
	case _CustomTokenPool.abi.Events["Unpaused"].ID:
		return _CustomTokenPool.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CustomTokenPoolBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
}

func (CustomTokenPoolLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (CustomTokenPoolMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (CustomTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CustomTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CustomTokenPoolPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (CustomTokenPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (CustomTokenPoolSynthBurned) Topic() common.Hash {
	return common.HexToHash("0x02992093bca69a36949677658a77d359b510dc6232c68f9f118f7c0127a1b147")
}

func (CustomTokenPoolSynthMinted) Topic() common.Hash {
	return common.HexToHash("0xbb0b72e5f44e331506684da008a30e10d50658c29d8159f6c6ab40bf1e52e600")
}

func (CustomTokenPoolUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_CustomTokenPool *CustomTokenPool) Address() common.Address {
	return _CustomTokenPool.address
}

type CustomTokenPoolInterface interface {
	GetToken(opts *bind.CallOpts) (common.Address, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int) (*types.Transaction, error)

	SetOffRamp(opts *bind.TransactOpts, offRamp common.Address, permission bool) (*types.Transaction, error)

	SetOnRamp(opts *bind.TransactOpts, onRamp common.Address, permission bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*CustomTokenPoolBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*CustomTokenPoolBurned, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*CustomTokenPoolLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*CustomTokenPoolLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*CustomTokenPoolMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*CustomTokenPoolMinted, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CustomTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CustomTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CustomTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CustomTokenPoolOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*CustomTokenPoolPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*CustomTokenPoolPaused, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*CustomTokenPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*CustomTokenPoolReleased, error)

	FilterSynthBurned(opts *bind.FilterOpts) (*CustomTokenPoolSynthBurnedIterator, error)

	WatchSynthBurned(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolSynthBurned) (event.Subscription, error)

	ParseSynthBurned(log types.Log) (*CustomTokenPoolSynthBurned, error)

	FilterSynthMinted(opts *bind.FilterOpts) (*CustomTokenPoolSynthMintedIterator, error)

	WatchSynthMinted(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolSynthMinted) (event.Subscription, error)

	ParseSynthMinted(log types.Log) (*CustomTokenPoolSynthMinted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*CustomTokenPoolUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*CustomTokenPoolUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
