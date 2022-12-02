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
	Bin: "0x60a060405234801561001057600080fd5b5060405161101b38038061101b83398101604081905261002f91610181565b8033806000816100865760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100b6576100b6816100d8565b50506001805460ff60a01b19169055506001600160a01b0316608052506101b1565b336001600160a01b038216036101305760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161007d565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561019357600080fd5b81516001600160a01b03811681146101aa57600080fd5b9392505050565b608051610e486101d360003960008181610134015261056a0152610e486000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c80636f32b8721161008c5780638da5cb5b116100665780638da5cb5b14610215578063cc8af2e814610233578063ea6192a214610246578063f2fde38b1461025957600080fd5b80636f32b872146101cc57806379ba5097146102055780638456cb591461020d57600080fd5b8063503c2858116100bd578063503c2858146101835780635c975abb1461019657806369e946d4146101b957600080fd5b80631d7a74a0146100e457806321df0da7146101325780633f4ba83a14610179575b600080fd5b61011d6100f2366004610cd8565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205460ff1690565b60405190151581526020015b60405180910390f35b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610129565b61018161026c565b005b610181610191366004610cf5565b61027e565b60015474010000000000000000000000000000000000000000900460ff1661011d565b6101816101c7366004610d1c565b61032e565b61011d6101da366004610cd8565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460ff1690565b61018161038c565b61018161046f565b60005473ffffffffffffffffffffffffffffffffffffffff16610154565b610181610241366004610d1c565b61047f565b610181610254366004610d55565b6104dd565b610181610267366004610cd8565b6105e2565b6102746105f6565b61027c61065d565b565b60015474010000000000000000000000000000000000000000900460ff16156102ee5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064015b60405180910390fd5b6102f661073c565b60405181815233907f9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd600089060200160405180910390a250565b6103366105f6565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260036020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b60015473ffffffffffffffffffffffffffffffffffffffff1633146103f35760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016102e5565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6104776105f6565b61027c6107ac565b6104876105f6565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b60015474010000000000000000000000000000000000000000900460ff16156105485760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016102e5565b61055061087e565b61059173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001683836108b9565b60405181815273ffffffffffffffffffffffffffffffffffffffff83169033907f2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f529060200160405180910390a35050565b6105ea6105f6565b6105f38161094b565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461027c5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016102e5565b60015474010000000000000000000000000000000000000000900460ff166106c75760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016102e5565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60005473ffffffffffffffffffffffffffffffffffffffff16331480159061077557503360009081526002602052604090205460ff165b155b1561027c576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015474010000000000000000000000000000000000000000900460ff16156108175760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016102e5565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586107123390565b60005473ffffffffffffffffffffffffffffffffffffffff16331480159061077557503360009081526003602052604090205460ff16610773565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610946908490610a26565b505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036109b05760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016102e5565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000610a88826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16610b189092919063ffffffff16565b8051909150156109465780806020019051810190610aa69190610d81565b6109465760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016102e5565b6060610b278484600085610b31565b90505b9392505050565b606082471015610ba95760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016102e5565b843b610bf75760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102e5565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610c209190610dce565b60006040518083038185875af1925050503d8060008114610c5d576040519150601f19603f3d011682016040523d82523d6000602084013e610c62565b606091505b5091509150610c72828286610c7d565b979650505050505050565b60608315610c8c575081610b2a565b825115610c9c5782518084602001fd5b8160405162461bcd60e51b81526004016102e59190610dea565b73ffffffffffffffffffffffffffffffffffffffff811681146105f357600080fd5b600060208284031215610cea57600080fd5b8135610b2a81610cb6565b600060208284031215610d0757600080fd5b5035919050565b80151581146105f357600080fd5b60008060408385031215610d2f57600080fd5b8235610d3a81610cb6565b91506020830135610d4a81610d0e565b809150509250929050565b60008060408385031215610d6857600080fd5b8235610d7381610cb6565b946020939093013593505050565b600060208284031215610d9357600080fd5b8151610b2a81610d0e565b60005b83811015610db9578181015183820152602001610da1565b83811115610dc8576000848401525b50505050565b60008251610de0818460208701610d9e565b9190910192915050565b6020815260008251806020840152610e09816040850160208701610d9e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fea164736f6c634300080f000a",
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
