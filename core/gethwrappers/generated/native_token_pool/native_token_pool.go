// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package native_token_pool

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

var NativeTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ExceedsTokenLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionsError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"isOnRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockOrBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"releaseOrMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"permission\",\"type\":\"bool\"}],\"name\":\"setOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"permission\",\"type\":\"bool\"}],\"name\":\"setOnRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610de4380380610de483398101604081905261002f91610181565b8033806000816100865760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100b6576100b6816100d8565b50506001805460ff60a01b19169055506001600160a01b0316608052506101b1565b336001600160a01b038216036101305760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161007d565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561019357600080fd5b81516001600160a01b03811681146101aa57600080fd5b9392505050565b608051610c116101d36000396000818161012701526104610152610c116000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c80636f32b8721161008c5780638da5cb5b116100665780638da5cb5b146101dd578063cc8af2e8146101ee578063ea6192a214610201578063f2fde38b1461021457600080fd5b80636f32b872146101a157806379ba5097146101cd5780638456cb59146101d557600080fd5b8063503c2858116100bd578063503c2858146101695780635c975abb1461017c57806369e946d41461018e57600080fd5b80631d7a74a0146100e457806321df0da7146101255780633f4ba83a1461015f575b600080fd5b6101106100f2366004610abf565b6001600160a01b031660009081526003602052604090205460ff1690565b60405190151581526020015b60405180910390f35b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b03909116815260200161011c565b610167610227565b005b610167610177366004610adc565b610239565b600154600160a01b900460ff16610110565b61016761019c366004610b03565b6102cb565b6101106101af366004610abf565b6001600160a01b031660009081526002602052604090205460ff1690565b6101676102fe565b6101676103bc565b6000546001600160a01b0316610147565b6101676101fc366004610b03565b6103cc565b61016761020f366004610b3c565b6103ff565b610167610222366004610abf565b6104cc565b61022f6104e0565b61023761053a565b565b600154600160a01b900460ff161561028b5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064015b60405180910390fd5b6102936105e0565b60405181815233907f9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd600089060200160405180910390a250565b6102d36104e0565b6001600160a01b03919091166000908152600360205260409020805460ff1916911515919091179055565b6001546001600160a01b031633146103585760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610282565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6103c46104e0565b610237610643565b6103d46104e0565b6001600160a01b03919091166000908152600260205260409020805460ff1916911515919091179055565b600154600160a01b900460ff161561044c5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610282565b6104546106cb565b6104886001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001683836106f9565b6040518181526001600160a01b0383169033907f2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f529060200160405180910390a35050565b6104d46104e0565b6104dd8161077e565b50565b6000546001600160a01b031633146102375760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610282565b600154600160a01b900460ff166105935760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610282565b6001805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6000546001600160a01b0316331480159061060c57503360009081526002602052604090205460ff165b155b15610237576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600154600160a01b900460ff16156106905760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610282565b6001805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586105c33390565b6000546001600160a01b0316331480159061060c57503360009081526003602052604090205460ff1661060a565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610779908490610834565b505050565b336001600160a01b038216036107d65760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610282565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000610889826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166109199092919063ffffffff16565b80519091501561077957808060200190518101906108a79190610b68565b6107795760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610282565b60606109288484600085610932565b90505b9392505050565b6060824710156109aa5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610282565b843b6109f85760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610282565b600080866001600160a01b03168587604051610a149190610bb5565b60006040518083038185875af1925050503d8060008114610a51576040519150601f19603f3d011682016040523d82523d6000602084013e610a56565b606091505b5091509150610a66828286610a71565b979650505050505050565b60608315610a8057508161092b565b825115610a905782518084602001fd5b8160405162461bcd60e51b81526004016102829190610bd1565b6001600160a01b03811681146104dd57600080fd5b600060208284031215610ad157600080fd5b813561092b81610aaa565b600060208284031215610aee57600080fd5b5035919050565b80151581146104dd57600080fd5b60008060408385031215610b1657600080fd5b8235610b2181610aaa565b91506020830135610b3181610af5565b809150509250929050565b60008060408385031215610b4f57600080fd5b8235610b5a81610aaa565b946020939093013593505050565b600060208284031215610b7a57600080fd5b815161092b81610af5565b60005b83811015610ba0578181015183820152602001610b88565b83811115610baf576000848401525b50505050565b60008251610bc7818460208701610b85565b9190910192915050565b6020815260008251806020840152610bf0816040850160208701610b85565b601f01601f1916919091016040019291505056fea164736f6c634300080f000a",
}

var NativeTokenPoolABI = NativeTokenPoolMetaData.ABI

var NativeTokenPoolBin = NativeTokenPoolMetaData.Bin

func DeployNativeTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *NativeTokenPool, error) {
	parsed, err := NativeTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenPoolBin), backend, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenPool{NativeTokenPoolCaller: NativeTokenPoolCaller{contract: contract}, NativeTokenPoolTransactor: NativeTokenPoolTransactor{contract: contract}, NativeTokenPoolFilterer: NativeTokenPoolFilterer{contract: contract}}, nil
}

type NativeTokenPool struct {
	address common.Address
	abi     abi.ABI
	NativeTokenPoolCaller
	NativeTokenPoolTransactor
	NativeTokenPoolFilterer
}

type NativeTokenPoolCaller struct {
	contract *bind.BoundContract
}

type NativeTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type NativeTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type NativeTokenPoolSession struct {
	Contract     *NativeTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type NativeTokenPoolCallerSession struct {
	Contract *NativeTokenPoolCaller
	CallOpts bind.CallOpts
}

type NativeTokenPoolTransactorSession struct {
	Contract     *NativeTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type NativeTokenPoolRaw struct {
	Contract *NativeTokenPool
}

type NativeTokenPoolCallerRaw struct {
	Contract *NativeTokenPoolCaller
}

type NativeTokenPoolTransactorRaw struct {
	Contract *NativeTokenPoolTransactor
}

func NewNativeTokenPool(address common.Address, backend bind.ContractBackend) (*NativeTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(NativeTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindNativeTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPool{address: address, abi: abi, NativeTokenPoolCaller: NativeTokenPoolCaller{contract: contract}, NativeTokenPoolTransactor: NativeTokenPoolTransactor{contract: contract}, NativeTokenPoolFilterer: NativeTokenPoolFilterer{contract: contract}}, nil
}

func NewNativeTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenPoolCaller, error) {
	contract, err := bindNativeTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolCaller{contract: contract}, nil
}

func NewNativeTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenPoolTransactor, error) {
	contract, err := bindNativeTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolTransactor{contract: contract}, nil
}

func NewNativeTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenPoolFilterer, error) {
	contract, err := bindNativeTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolFilterer{contract: contract}, nil
}

func bindNativeTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NativeTokenPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_NativeTokenPool *NativeTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenPool.Contract.NativeTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_NativeTokenPool *NativeTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.NativeTokenPoolTransactor.contract.Transfer(opts)
}

func (_NativeTokenPool *NativeTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.NativeTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_NativeTokenPool *NativeTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_NativeTokenPool *NativeTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.contract.Transfer(opts)
}

func (_NativeTokenPool *NativeTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_NativeTokenPool *NativeTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_NativeTokenPool *NativeTokenPoolSession) GetToken() (common.Address, error) {
	return _NativeTokenPool.Contract.GetToken(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _NativeTokenPool.Contract.GetToken(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolCaller) IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenPool.contract.Call(opts, &out, "isOffRamp", offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_NativeTokenPool *NativeTokenPoolSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _NativeTokenPool.Contract.IsOffRamp(&_NativeTokenPool.CallOpts, offRamp)
}

func (_NativeTokenPool *NativeTokenPoolCallerSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _NativeTokenPool.Contract.IsOffRamp(&_NativeTokenPool.CallOpts, offRamp)
}

func (_NativeTokenPool *NativeTokenPoolCaller) IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenPool.contract.Call(opts, &out, "isOnRamp", onRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_NativeTokenPool *NativeTokenPoolSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _NativeTokenPool.Contract.IsOnRamp(&_NativeTokenPool.CallOpts, onRamp)
}

func (_NativeTokenPool *NativeTokenPoolCallerSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _NativeTokenPool.Contract.IsOnRamp(&_NativeTokenPool.CallOpts, onRamp)
}

func (_NativeTokenPool *NativeTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_NativeTokenPool *NativeTokenPoolSession) Owner() (common.Address, error) {
	return _NativeTokenPool.Contract.Owner(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolCallerSession) Owner() (common.Address, error) {
	return _NativeTokenPool.Contract.Owner(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_NativeTokenPool *NativeTokenPoolSession) Paused() (bool, error) {
	return _NativeTokenPool.Contract.Paused(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolCallerSession) Paused() (bool, error) {
	return _NativeTokenPool.Contract.Paused(&_NativeTokenPool.CallOpts)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_NativeTokenPool *NativeTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _NativeTokenPool.Contract.AcceptOwnership(&_NativeTokenPool.TransactOpts)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _NativeTokenPool.Contract.AcceptOwnership(&_NativeTokenPool.TransactOpts)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "lockOrBurn", amount)
}

func (_NativeTokenPool *NativeTokenPoolSession) LockOrBurn(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.LockOrBurn(&_NativeTokenPool.TransactOpts, amount)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) LockOrBurn(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.LockOrBurn(&_NativeTokenPool.TransactOpts, amount)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "pause")
}

func (_NativeTokenPool *NativeTokenPoolSession) Pause() (*types.Transaction, error) {
	return _NativeTokenPool.Contract.Pause(&_NativeTokenPool.TransactOpts)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) Pause() (*types.Transaction, error) {
	return _NativeTokenPool.Contract.Pause(&_NativeTokenPool.TransactOpts)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "releaseOrMint", recipient, amount)
}

func (_NativeTokenPool *NativeTokenPoolSession) ReleaseOrMint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.ReleaseOrMint(&_NativeTokenPool.TransactOpts, recipient, amount)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) ReleaseOrMint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.ReleaseOrMint(&_NativeTokenPool.TransactOpts, recipient, amount)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) SetOffRamp(opts *bind.TransactOpts, offRamp common.Address, permission bool) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "setOffRamp", offRamp, permission)
}

func (_NativeTokenPool *NativeTokenPoolSession) SetOffRamp(offRamp common.Address, permission bool) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.SetOffRamp(&_NativeTokenPool.TransactOpts, offRamp, permission)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) SetOffRamp(offRamp common.Address, permission bool) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.SetOffRamp(&_NativeTokenPool.TransactOpts, offRamp, permission)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) SetOnRamp(opts *bind.TransactOpts, onRamp common.Address, permission bool) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "setOnRamp", onRamp, permission)
}

func (_NativeTokenPool *NativeTokenPoolSession) SetOnRamp(onRamp common.Address, permission bool) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.SetOnRamp(&_NativeTokenPool.TransactOpts, onRamp, permission)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) SetOnRamp(onRamp common.Address, permission bool) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.SetOnRamp(&_NativeTokenPool.TransactOpts, onRamp, permission)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_NativeTokenPool *NativeTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.TransferOwnership(&_NativeTokenPool.TransactOpts, to)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _NativeTokenPool.Contract.TransferOwnership(&_NativeTokenPool.TransactOpts, to)
}

func (_NativeTokenPool *NativeTokenPoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenPool.contract.Transact(opts, "unpause")
}

func (_NativeTokenPool *NativeTokenPoolSession) Unpause() (*types.Transaction, error) {
	return _NativeTokenPool.Contract.Unpause(&_NativeTokenPool.TransactOpts)
}

func (_NativeTokenPool *NativeTokenPoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _NativeTokenPool.Contract.Unpause(&_NativeTokenPool.TransactOpts)
}

type NativeTokenPoolBurnedIterator struct {
	Event *NativeTokenPoolBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolBurned)
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
		it.Event = new(NativeTokenPoolBurned)
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

func (it *NativeTokenPoolBurnedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenPoolBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolBurnedIterator{contract: _NativeTokenPool.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolBurned)
				if err := _NativeTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseBurned(log types.Log) (*NativeTokenPoolBurned, error) {
	event := new(NativeTokenPoolBurned)
	if err := _NativeTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolLockedIterator struct {
	Event *NativeTokenPoolLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolLocked)
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
		it.Event = new(NativeTokenPoolLocked)
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

func (it *NativeTokenPoolLockedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenPoolLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolLockedIterator{contract: _NativeTokenPool.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolLocked)
				if err := _NativeTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseLocked(log types.Log) (*NativeTokenPoolLocked, error) {
	event := new(NativeTokenPoolLocked)
	if err := _NativeTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolMintedIterator struct {
	Event *NativeTokenPoolMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolMinted)
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
		it.Event = new(NativeTokenPoolMinted)
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

func (it *NativeTokenPoolMintedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*NativeTokenPoolMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolMintedIterator{contract: _NativeTokenPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolMinted)
				if err := _NativeTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseMinted(log types.Log) (*NativeTokenPoolMinted, error) {
	event := new(NativeTokenPoolMinted)
	if err := _NativeTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolOwnershipTransferRequestedIterator struct {
	Event *NativeTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolOwnershipTransferRequested)
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
		it.Event = new(NativeTokenPoolOwnershipTransferRequested)
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

func (it *NativeTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolOwnershipTransferRequestedIterator{contract: _NativeTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolOwnershipTransferRequested)
				if err := _NativeTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*NativeTokenPoolOwnershipTransferRequested, error) {
	event := new(NativeTokenPoolOwnershipTransferRequested)
	if err := _NativeTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolOwnershipTransferredIterator struct {
	Event *NativeTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolOwnershipTransferred)
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
		it.Event = new(NativeTokenPoolOwnershipTransferred)
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

func (it *NativeTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolOwnershipTransferredIterator{contract: _NativeTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolOwnershipTransferred)
				if err := _NativeTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*NativeTokenPoolOwnershipTransferred, error) {
	event := new(NativeTokenPoolOwnershipTransferred)
	if err := _NativeTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolPausedIterator struct {
	Event *NativeTokenPoolPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolPaused)
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
		it.Event = new(NativeTokenPoolPaused)
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

func (it *NativeTokenPoolPausedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*NativeTokenPoolPausedIterator, error) {

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolPausedIterator{contract: _NativeTokenPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolPaused) (event.Subscription, error) {

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolPaused)
				if err := _NativeTokenPool.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParsePaused(log types.Log) (*NativeTokenPoolPaused, error) {
	event := new(NativeTokenPoolPaused)
	if err := _NativeTokenPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolReleasedIterator struct {
	Event *NativeTokenPoolReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolReleased)
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
		it.Event = new(NativeTokenPoolReleased)
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

func (it *NativeTokenPoolReleasedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*NativeTokenPoolReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolReleasedIterator{contract: _NativeTokenPool.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolReleased)
				if err := _NativeTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseReleased(log types.Log) (*NativeTokenPoolReleased, error) {
	event := new(NativeTokenPoolReleased)
	if err := _NativeTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type NativeTokenPoolUnpausedIterator struct {
	Event *NativeTokenPoolUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *NativeTokenPoolUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenPoolUnpaused)
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
		it.Event = new(NativeTokenPoolUnpaused)
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

func (it *NativeTokenPoolUnpausedIterator) Error() error {
	return it.fail
}

func (it *NativeTokenPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type NativeTokenPoolUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_NativeTokenPool *NativeTokenPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NativeTokenPoolUnpausedIterator, error) {

	logs, sub, err := _NativeTokenPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NativeTokenPoolUnpausedIterator{contract: _NativeTokenPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_NativeTokenPool *NativeTokenPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _NativeTokenPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(NativeTokenPoolUnpaused)
				if err := _NativeTokenPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_NativeTokenPool *NativeTokenPoolFilterer) ParseUnpaused(log types.Log) (*NativeTokenPoolUnpaused, error) {
	event := new(NativeTokenPoolUnpaused)
	if err := _NativeTokenPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_NativeTokenPool *NativeTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _NativeTokenPool.abi.Events["Burned"].ID:
		return _NativeTokenPool.ParseBurned(log)
	case _NativeTokenPool.abi.Events["Locked"].ID:
		return _NativeTokenPool.ParseLocked(log)
	case _NativeTokenPool.abi.Events["Minted"].ID:
		return _NativeTokenPool.ParseMinted(log)
	case _NativeTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _NativeTokenPool.ParseOwnershipTransferRequested(log)
	case _NativeTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _NativeTokenPool.ParseOwnershipTransferred(log)
	case _NativeTokenPool.abi.Events["Paused"].ID:
		return _NativeTokenPool.ParsePaused(log)
	case _NativeTokenPool.abi.Events["Released"].ID:
		return _NativeTokenPool.ParseReleased(log)
	case _NativeTokenPool.abi.Events["Unpaused"].ID:
		return _NativeTokenPool.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (NativeTokenPoolBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
}

func (NativeTokenPoolLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (NativeTokenPoolMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (NativeTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (NativeTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (NativeTokenPoolPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (NativeTokenPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (NativeTokenPoolUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_NativeTokenPool *NativeTokenPool) Address() common.Address {
	return _NativeTokenPool.address
}

type NativeTokenPoolInterface interface {
	GetToken(opts *bind.CallOpts) (common.Address, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	SetOffRamp(opts *bind.TransactOpts, offRamp common.Address, permission bool) (*types.Transaction, error)

	SetOnRamp(opts *bind.TransactOpts, onRamp common.Address, permission bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenPoolBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*NativeTokenPoolBurned, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenPoolLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*NativeTokenPoolLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*NativeTokenPoolMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*NativeTokenPoolMinted, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*NativeTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*NativeTokenPoolOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*NativeTokenPoolPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*NativeTokenPoolPaused, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*NativeTokenPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*NativeTokenPoolReleased, error)

	FilterUnpaused(opts *bind.FilterOpts) (*NativeTokenPoolUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenPoolUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*NativeTokenPoolUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
