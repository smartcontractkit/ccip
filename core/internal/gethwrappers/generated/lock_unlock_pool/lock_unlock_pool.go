// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lock_unlock_pool

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

var LockUnlockPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"isOnRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockOrBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"releaseOrMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"permission\",\"type\":\"bool\"}],\"name\":\"setOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"permission\",\"type\":\"bool\"}],\"name\":\"setOnRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161122a38038061122a83398101604081905261002f91610184565b33806000816100855760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100b5576100b5816100da565b50506001805460ff60a01b191690555060601b6001600160601b0319166080526101b4565b6001600160a01b0381163314156101335760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161007c565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561019657600080fd5b81516001600160a01b03811681146101ad57600080fd5b9392505050565b60805160601c61104a6101e060003960008181610134015281816104f30152610613015261104a6000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c806379ba50971161008c578063cc8af2e811610066578063cc8af2e814610220578063ea6192a214610233578063eb54b3bf14610246578063f2fde38b1461025957600080fd5b806379ba5097146101f25780638456cb59146101fa5780638da5cb5b1461020257600080fd5b80635c975abb116100bd5780635c975abb1461018357806369e946d4146101a65780636f32b872146101b957600080fd5b80631d7a74a0146100e457806321df0da7146101325780633f4ba83a14610179575b600080fd5b61011d6100f2366004610ed5565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205460ff1690565b60405190151581526020015b60405180910390f35b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610129565b61018161026c565b005b60015474010000000000000000000000000000000000000000900460ff1661011d565b6101816101b4366004610f3b565b61027e565b61011d6101c7366004610ed5565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460ff1690565b6101816102dc565b6101816103de565b60005473ffffffffffffffffffffffffffffffffffffffff16610154565b61018161022e366004610f3b565b6103ee565b610181610241366004610ef2565b61044c565b610181610254366004610ef2565b61056c565b610181610267366004610ed5565b610685565b610274610699565b61027c61071a565b565b610286610699565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260036020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b60015473ffffffffffffffffffffffffffffffffffffffff163314610362576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6103e6610699565b61027c610813565b6103f6610699565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b60015474010000000000000000000000000000000000000000900460ff16156104d1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610359565b6104d96108ff565b61051a73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016838361096f565b60405181815273ffffffffffffffffffffffffffffffffffffffff83169033907f2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52906020015b60405180910390a35050565b60015474010000000000000000000000000000000000000000900460ff16156105f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610359565b6105f9610a48565b61063b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016833084610a83565b60405181815273ffffffffffffffffffffffffffffffffffffffff83169033907f989eaa915cbb416ea3d6f9a63b1a3de51770c7674b11fe21ecdf76b4e1d1391090602001610560565b61068d610699565b61069681610ae7565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461027c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610359565b60015474010000000000000000000000000000000000000000900460ff1661079e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610359565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60015474010000000000000000000000000000000000000000900460ff1615610898576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610359565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586107e93390565b60005473ffffffffffffffffffffffffffffffffffffffff16331480159061093857503360009081526003602052604090205460ff165b155b1561027c576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610a439084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610bdd565b505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331480159061093857503360009081526002602052604090205460ff16610936565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052610ae19085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016109c1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8116331415610b67576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610359565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000610c3f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16610ce99092919063ffffffff16565b805190915015610a435780806020019051810190610c5d9190610f1e565b610a43576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610359565b6060610cf88484600085610d02565b90505b9392505050565b606082471015610d94576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610359565b843b610dfc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610359565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610e259190610f74565b60006040518083038185875af1925050503d8060008114610e62576040519150601f19603f3d011682016040523d82523d6000602084013e610e67565b606091505b5091509150610e77828286610e82565b979650505050505050565b60608315610e91575081610cfb565b825115610ea15782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103599190610f90565b600060208284031215610ee757600080fd5b8135610cfb8161100d565b60008060408385031215610f0557600080fd5b8235610f108161100d565b946020939093013593505050565b600060208284031215610f3057600080fd5b8151610cfb8161102f565b60008060408385031215610f4e57600080fd5b8235610f598161100d565b91506020830135610f698161102f565b809150509250929050565b60008251610f86818460208701610fe1565b9190910192915050565b6020815260008251806020840152610faf816040850160208701610fe1565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60005b83811015610ffc578181015183820152602001610fe4565b83811115610ae15750506000910152565b73ffffffffffffffffffffffffffffffffffffffff8116811461069657600080fd5b801515811461069657600080fdfea164736f6c6343000806000a",
}

var LockUnlockPoolABI = LockUnlockPoolMetaData.ABI

var LockUnlockPoolBin = LockUnlockPoolMetaData.Bin

func DeployLockUnlockPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *LockUnlockPool, error) {
	parsed, err := LockUnlockPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LockUnlockPoolBin), backend, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LockUnlockPool{LockUnlockPoolCaller: LockUnlockPoolCaller{contract: contract}, LockUnlockPoolTransactor: LockUnlockPoolTransactor{contract: contract}, LockUnlockPoolFilterer: LockUnlockPoolFilterer{contract: contract}}, nil
}

type LockUnlockPool struct {
	address common.Address
	abi     abi.ABI
	LockUnlockPoolCaller
	LockUnlockPoolTransactor
	LockUnlockPoolFilterer
}

type LockUnlockPoolCaller struct {
	contract *bind.BoundContract
}

type LockUnlockPoolTransactor struct {
	contract *bind.BoundContract
}

type LockUnlockPoolFilterer struct {
	contract *bind.BoundContract
}

type LockUnlockPoolSession struct {
	Contract     *LockUnlockPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type LockUnlockPoolCallerSession struct {
	Contract *LockUnlockPoolCaller
	CallOpts bind.CallOpts
}

type LockUnlockPoolTransactorSession struct {
	Contract     *LockUnlockPoolTransactor
	TransactOpts bind.TransactOpts
}

type LockUnlockPoolRaw struct {
	Contract *LockUnlockPool
}

type LockUnlockPoolCallerRaw struct {
	Contract *LockUnlockPoolCaller
}

type LockUnlockPoolTransactorRaw struct {
	Contract *LockUnlockPoolTransactor
}

func NewLockUnlockPool(address common.Address, backend bind.ContractBackend) (*LockUnlockPool, error) {
	abi, err := abi.JSON(strings.NewReader(LockUnlockPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindLockUnlockPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockUnlockPool{address: address, abi: abi, LockUnlockPoolCaller: LockUnlockPoolCaller{contract: contract}, LockUnlockPoolTransactor: LockUnlockPoolTransactor{contract: contract}, LockUnlockPoolFilterer: LockUnlockPoolFilterer{contract: contract}}, nil
}

func NewLockUnlockPoolCaller(address common.Address, caller bind.ContractCaller) (*LockUnlockPoolCaller, error) {
	contract, err := bindLockUnlockPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockUnlockPoolCaller{contract: contract}, nil
}

func NewLockUnlockPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LockUnlockPoolTransactor, error) {
	contract, err := bindLockUnlockPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockUnlockPoolTransactor{contract: contract}, nil
}

func NewLockUnlockPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LockUnlockPoolFilterer, error) {
	contract, err := bindLockUnlockPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockUnlockPoolFilterer{contract: contract}, nil
}

func bindLockUnlockPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockUnlockPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_LockUnlockPool *LockUnlockPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LockUnlockPool.Contract.LockUnlockPoolCaller.contract.Call(opts, result, method, params...)
}

func (_LockUnlockPool *LockUnlockPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockUnlockPool.Contract.LockUnlockPoolTransactor.contract.Transfer(opts)
}

func (_LockUnlockPool *LockUnlockPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockUnlockPool.Contract.LockUnlockPoolTransactor.contract.Transact(opts, method, params...)
}

func (_LockUnlockPool *LockUnlockPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LockUnlockPool.Contract.contract.Call(opts, result, method, params...)
}

func (_LockUnlockPool *LockUnlockPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockUnlockPool.Contract.contract.Transfer(opts)
}

func (_LockUnlockPool *LockUnlockPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockUnlockPool.Contract.contract.Transact(opts, method, params...)
}

func (_LockUnlockPool *LockUnlockPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockUnlockPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockUnlockPool *LockUnlockPoolSession) GetToken() (common.Address, error) {
	return _LockUnlockPool.Contract.GetToken(&_LockUnlockPool.CallOpts)
}

func (_LockUnlockPool *LockUnlockPoolCallerSession) GetToken() (common.Address, error) {
	return _LockUnlockPool.Contract.GetToken(&_LockUnlockPool.CallOpts)
}

func (_LockUnlockPool *LockUnlockPoolCaller) IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _LockUnlockPool.contract.Call(opts, &out, "isOffRamp", offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockUnlockPool *LockUnlockPoolSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _LockUnlockPool.Contract.IsOffRamp(&_LockUnlockPool.CallOpts, offRamp)
}

func (_LockUnlockPool *LockUnlockPoolCallerSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _LockUnlockPool.Contract.IsOffRamp(&_LockUnlockPool.CallOpts, offRamp)
}

func (_LockUnlockPool *LockUnlockPoolCaller) IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error) {
	var out []interface{}
	err := _LockUnlockPool.contract.Call(opts, &out, "isOnRamp", onRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockUnlockPool *LockUnlockPoolSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _LockUnlockPool.Contract.IsOnRamp(&_LockUnlockPool.CallOpts, onRamp)
}

func (_LockUnlockPool *LockUnlockPoolCallerSession) IsOnRamp(onRamp common.Address) (bool, error) {
	return _LockUnlockPool.Contract.IsOnRamp(&_LockUnlockPool.CallOpts, onRamp)
}

func (_LockUnlockPool *LockUnlockPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockUnlockPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockUnlockPool *LockUnlockPoolSession) Owner() (common.Address, error) {
	return _LockUnlockPool.Contract.Owner(&_LockUnlockPool.CallOpts)
}

func (_LockUnlockPool *LockUnlockPoolCallerSession) Owner() (common.Address, error) {
	return _LockUnlockPool.Contract.Owner(&_LockUnlockPool.CallOpts)
}

func (_LockUnlockPool *LockUnlockPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LockUnlockPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockUnlockPool *LockUnlockPoolSession) Paused() (bool, error) {
	return _LockUnlockPool.Contract.Paused(&_LockUnlockPool.CallOpts)
}

func (_LockUnlockPool *LockUnlockPoolCallerSession) Paused() (bool, error) {
	return _LockUnlockPool.Contract.Paused(&_LockUnlockPool.CallOpts)
}

func (_LockUnlockPool *LockUnlockPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockUnlockPool.contract.Transact(opts, "acceptOwnership")
}

func (_LockUnlockPool *LockUnlockPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _LockUnlockPool.Contract.AcceptOwnership(&_LockUnlockPool.TransactOpts)
}

func (_LockUnlockPool *LockUnlockPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LockUnlockPool.Contract.AcceptOwnership(&_LockUnlockPool.TransactOpts)
}

func (_LockUnlockPool *LockUnlockPoolTransactor) LockOrBurn(opts *bind.TransactOpts, depositor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockUnlockPool.contract.Transact(opts, "lockOrBurn", depositor, amount)
}

func (_LockUnlockPool *LockUnlockPoolSession) LockOrBurn(depositor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockUnlockPool.Contract.LockOrBurn(&_LockUnlockPool.TransactOpts, depositor, amount)
}

func (_LockUnlockPool *LockUnlockPoolTransactorSession) LockOrBurn(depositor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockUnlockPool.Contract.LockOrBurn(&_LockUnlockPool.TransactOpts, depositor, amount)
}

func (_LockUnlockPool *LockUnlockPoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockUnlockPool.contract.Transact(opts, "pause")
}

func (_LockUnlockPool *LockUnlockPoolSession) Pause() (*types.Transaction, error) {
	return _LockUnlockPool.Contract.Pause(&_LockUnlockPool.TransactOpts)
}

func (_LockUnlockPool *LockUnlockPoolTransactorSession) Pause() (*types.Transaction, error) {
	return _LockUnlockPool.Contract.Pause(&_LockUnlockPool.TransactOpts)
}

func (_LockUnlockPool *LockUnlockPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockUnlockPool.contract.Transact(opts, "releaseOrMint", recipient, amount)
}

func (_LockUnlockPool *LockUnlockPoolSession) ReleaseOrMint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockUnlockPool.Contract.ReleaseOrMint(&_LockUnlockPool.TransactOpts, recipient, amount)
}

func (_LockUnlockPool *LockUnlockPoolTransactorSession) ReleaseOrMint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockUnlockPool.Contract.ReleaseOrMint(&_LockUnlockPool.TransactOpts, recipient, amount)
}

func (_LockUnlockPool *LockUnlockPoolTransactor) SetOffRamp(opts *bind.TransactOpts, offRamp common.Address, permission bool) (*types.Transaction, error) {
	return _LockUnlockPool.contract.Transact(opts, "setOffRamp", offRamp, permission)
}

func (_LockUnlockPool *LockUnlockPoolSession) SetOffRamp(offRamp common.Address, permission bool) (*types.Transaction, error) {
	return _LockUnlockPool.Contract.SetOffRamp(&_LockUnlockPool.TransactOpts, offRamp, permission)
}

func (_LockUnlockPool *LockUnlockPoolTransactorSession) SetOffRamp(offRamp common.Address, permission bool) (*types.Transaction, error) {
	return _LockUnlockPool.Contract.SetOffRamp(&_LockUnlockPool.TransactOpts, offRamp, permission)
}

func (_LockUnlockPool *LockUnlockPoolTransactor) SetOnRamp(opts *bind.TransactOpts, onRamp common.Address, permission bool) (*types.Transaction, error) {
	return _LockUnlockPool.contract.Transact(opts, "setOnRamp", onRamp, permission)
}

func (_LockUnlockPool *LockUnlockPoolSession) SetOnRamp(onRamp common.Address, permission bool) (*types.Transaction, error) {
	return _LockUnlockPool.Contract.SetOnRamp(&_LockUnlockPool.TransactOpts, onRamp, permission)
}

func (_LockUnlockPool *LockUnlockPoolTransactorSession) SetOnRamp(onRamp common.Address, permission bool) (*types.Transaction, error) {
	return _LockUnlockPool.Contract.SetOnRamp(&_LockUnlockPool.TransactOpts, onRamp, permission)
}

func (_LockUnlockPool *LockUnlockPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _LockUnlockPool.contract.Transact(opts, "transferOwnership", to)
}

func (_LockUnlockPool *LockUnlockPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LockUnlockPool.Contract.TransferOwnership(&_LockUnlockPool.TransactOpts, to)
}

func (_LockUnlockPool *LockUnlockPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LockUnlockPool.Contract.TransferOwnership(&_LockUnlockPool.TransactOpts, to)
}

func (_LockUnlockPool *LockUnlockPoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockUnlockPool.contract.Transact(opts, "unpause")
}

func (_LockUnlockPool *LockUnlockPoolSession) Unpause() (*types.Transaction, error) {
	return _LockUnlockPool.Contract.Unpause(&_LockUnlockPool.TransactOpts)
}

func (_LockUnlockPool *LockUnlockPoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _LockUnlockPool.Contract.Unpause(&_LockUnlockPool.TransactOpts)
}

type LockUnlockPoolBurntIterator struct {
	Event *LockUnlockPoolBurnt

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockUnlockPoolBurntIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockUnlockPoolBurnt)
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
		it.Event = new(LockUnlockPoolBurnt)
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

func (it *LockUnlockPoolBurntIterator) Error() error {
	return it.fail
}

func (it *LockUnlockPoolBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockUnlockPoolBurnt struct {
	Sender    common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_LockUnlockPool *LockUnlockPoolFilterer) FilterBurnt(opts *bind.FilterOpts, sender []common.Address, depositor []common.Address) (*LockUnlockPoolBurntIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _LockUnlockPool.contract.FilterLogs(opts, "Burnt", senderRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &LockUnlockPoolBurntIterator{contract: _LockUnlockPool.contract, event: "Burnt", logs: logs, sub: sub}, nil
}

func (_LockUnlockPool *LockUnlockPoolFilterer) WatchBurnt(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolBurnt, sender []common.Address, depositor []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _LockUnlockPool.contract.WatchLogs(opts, "Burnt", senderRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockUnlockPoolBurnt)
				if err := _LockUnlockPool.contract.UnpackLog(event, "Burnt", log); err != nil {
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

func (_LockUnlockPool *LockUnlockPoolFilterer) ParseBurnt(log types.Log) (*LockUnlockPoolBurnt, error) {
	event := new(LockUnlockPoolBurnt)
	if err := _LockUnlockPool.contract.UnpackLog(event, "Burnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockUnlockPoolLockedIterator struct {
	Event *LockUnlockPoolLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockUnlockPoolLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockUnlockPoolLocked)
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
		it.Event = new(LockUnlockPoolLocked)
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

func (it *LockUnlockPoolLockedIterator) Error() error {
	return it.fail
}

func (it *LockUnlockPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockUnlockPoolLocked struct {
	Sender    common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_LockUnlockPool *LockUnlockPoolFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address, depositor []common.Address) (*LockUnlockPoolLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _LockUnlockPool.contract.FilterLogs(opts, "Locked", senderRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &LockUnlockPoolLockedIterator{contract: _LockUnlockPool.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_LockUnlockPool *LockUnlockPoolFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolLocked, sender []common.Address, depositor []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _LockUnlockPool.contract.WatchLogs(opts, "Locked", senderRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockUnlockPoolLocked)
				if err := _LockUnlockPool.contract.UnpackLog(event, "Locked", log); err != nil {
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

func (_LockUnlockPool *LockUnlockPoolFilterer) ParseLocked(log types.Log) (*LockUnlockPoolLocked, error) {
	event := new(LockUnlockPoolLocked)
	if err := _LockUnlockPool.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockUnlockPoolMintedIterator struct {
	Event *LockUnlockPoolMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockUnlockPoolMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockUnlockPoolMinted)
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
		it.Event = new(LockUnlockPoolMinted)
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

func (it *LockUnlockPoolMintedIterator) Error() error {
	return it.fail
}

func (it *LockUnlockPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockUnlockPoolMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_LockUnlockPool *LockUnlockPoolFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LockUnlockPoolMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LockUnlockPool.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &LockUnlockPoolMintedIterator{contract: _LockUnlockPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_LockUnlockPool *LockUnlockPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LockUnlockPool.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockUnlockPoolMinted)
				if err := _LockUnlockPool.contract.UnpackLog(event, "Minted", log); err != nil {
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

func (_LockUnlockPool *LockUnlockPoolFilterer) ParseMinted(log types.Log) (*LockUnlockPoolMinted, error) {
	event := new(LockUnlockPoolMinted)
	if err := _LockUnlockPool.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockUnlockPoolOwnershipTransferRequestedIterator struct {
	Event *LockUnlockPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockUnlockPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockUnlockPoolOwnershipTransferRequested)
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
		it.Event = new(LockUnlockPoolOwnershipTransferRequested)
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

func (it *LockUnlockPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *LockUnlockPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockUnlockPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LockUnlockPool *LockUnlockPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockUnlockPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockUnlockPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LockUnlockPoolOwnershipTransferRequestedIterator{contract: _LockUnlockPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_LockUnlockPool *LockUnlockPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockUnlockPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockUnlockPoolOwnershipTransferRequested)
				if err := _LockUnlockPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_LockUnlockPool *LockUnlockPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*LockUnlockPoolOwnershipTransferRequested, error) {
	event := new(LockUnlockPoolOwnershipTransferRequested)
	if err := _LockUnlockPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockUnlockPoolOwnershipTransferredIterator struct {
	Event *LockUnlockPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockUnlockPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockUnlockPoolOwnershipTransferred)
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
		it.Event = new(LockUnlockPoolOwnershipTransferred)
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

func (it *LockUnlockPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *LockUnlockPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockUnlockPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LockUnlockPool *LockUnlockPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockUnlockPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockUnlockPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LockUnlockPoolOwnershipTransferredIterator{contract: _LockUnlockPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_LockUnlockPool *LockUnlockPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockUnlockPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockUnlockPoolOwnershipTransferred)
				if err := _LockUnlockPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_LockUnlockPool *LockUnlockPoolFilterer) ParseOwnershipTransferred(log types.Log) (*LockUnlockPoolOwnershipTransferred, error) {
	event := new(LockUnlockPoolOwnershipTransferred)
	if err := _LockUnlockPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockUnlockPoolPausedIterator struct {
	Event *LockUnlockPoolPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockUnlockPoolPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockUnlockPoolPaused)
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
		it.Event = new(LockUnlockPoolPaused)
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

func (it *LockUnlockPoolPausedIterator) Error() error {
	return it.fail
}

func (it *LockUnlockPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockUnlockPoolPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_LockUnlockPool *LockUnlockPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*LockUnlockPoolPausedIterator, error) {

	logs, sub, err := _LockUnlockPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LockUnlockPoolPausedIterator{contract: _LockUnlockPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_LockUnlockPool *LockUnlockPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolPaused) (event.Subscription, error) {

	logs, sub, err := _LockUnlockPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockUnlockPoolPaused)
				if err := _LockUnlockPool.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_LockUnlockPool *LockUnlockPoolFilterer) ParsePaused(log types.Log) (*LockUnlockPoolPaused, error) {
	event := new(LockUnlockPoolPaused)
	if err := _LockUnlockPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockUnlockPoolReleasedIterator struct {
	Event *LockUnlockPoolReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockUnlockPoolReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockUnlockPoolReleased)
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
		it.Event = new(LockUnlockPoolReleased)
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

func (it *LockUnlockPoolReleasedIterator) Error() error {
	return it.fail
}

func (it *LockUnlockPoolReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockUnlockPoolReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_LockUnlockPool *LockUnlockPoolFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LockUnlockPoolReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LockUnlockPool.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &LockUnlockPoolReleasedIterator{contract: _LockUnlockPool.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_LockUnlockPool *LockUnlockPoolFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LockUnlockPool.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockUnlockPoolReleased)
				if err := _LockUnlockPool.contract.UnpackLog(event, "Released", log); err != nil {
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

func (_LockUnlockPool *LockUnlockPoolFilterer) ParseReleased(log types.Log) (*LockUnlockPoolReleased, error) {
	event := new(LockUnlockPoolReleased)
	if err := _LockUnlockPool.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockUnlockPoolUnpausedIterator struct {
	Event *LockUnlockPoolUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockUnlockPoolUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockUnlockPoolUnpaused)
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
		it.Event = new(LockUnlockPoolUnpaused)
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

func (it *LockUnlockPoolUnpausedIterator) Error() error {
	return it.fail
}

func (it *LockUnlockPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockUnlockPoolUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_LockUnlockPool *LockUnlockPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LockUnlockPoolUnpausedIterator, error) {

	logs, sub, err := _LockUnlockPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LockUnlockPoolUnpausedIterator{contract: _LockUnlockPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_LockUnlockPool *LockUnlockPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _LockUnlockPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockUnlockPoolUnpaused)
				if err := _LockUnlockPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_LockUnlockPool *LockUnlockPoolFilterer) ParseUnpaused(log types.Log) (*LockUnlockPoolUnpaused, error) {
	event := new(LockUnlockPoolUnpaused)
	if err := _LockUnlockPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_LockUnlockPool *LockUnlockPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _LockUnlockPool.abi.Events["Burnt"].ID:
		return _LockUnlockPool.ParseBurnt(log)
	case _LockUnlockPool.abi.Events["Locked"].ID:
		return _LockUnlockPool.ParseLocked(log)
	case _LockUnlockPool.abi.Events["Minted"].ID:
		return _LockUnlockPool.ParseMinted(log)
	case _LockUnlockPool.abi.Events["OwnershipTransferRequested"].ID:
		return _LockUnlockPool.ParseOwnershipTransferRequested(log)
	case _LockUnlockPool.abi.Events["OwnershipTransferred"].ID:
		return _LockUnlockPool.ParseOwnershipTransferred(log)
	case _LockUnlockPool.abi.Events["Paused"].ID:
		return _LockUnlockPool.ParsePaused(log)
	case _LockUnlockPool.abi.Events["Released"].ID:
		return _LockUnlockPool.ParseReleased(log)
	case _LockUnlockPool.abi.Events["Unpaused"].ID:
		return _LockUnlockPool.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (LockUnlockPoolBurnt) Topic() common.Hash {
	return common.HexToHash("0xe8a89cc6e5096f9d9f43de82c077c1f4cfe707c0e0c2032176c68813b9ae6a5c")
}

func (LockUnlockPoolLocked) Topic() common.Hash {
	return common.HexToHash("0x989eaa915cbb416ea3d6f9a63b1a3de51770c7674b11fe21ecdf76b4e1d13910")
}

func (LockUnlockPoolMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (LockUnlockPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (LockUnlockPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (LockUnlockPoolPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (LockUnlockPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (LockUnlockPoolUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_LockUnlockPool *LockUnlockPool) Address() common.Address {
	return _LockUnlockPool.address
}

type LockUnlockPoolInterface interface {
	GetToken(opts *bind.CallOpts) (common.Address, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	IsOnRamp(opts *bind.CallOpts, onRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, depositor common.Address, amount *big.Int) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	SetOffRamp(opts *bind.TransactOpts, offRamp common.Address, permission bool) (*types.Transaction, error)

	SetOnRamp(opts *bind.TransactOpts, onRamp common.Address, permission bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterBurnt(opts *bind.FilterOpts, sender []common.Address, depositor []common.Address) (*LockUnlockPoolBurntIterator, error)

	WatchBurnt(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolBurnt, sender []common.Address, depositor []common.Address) (event.Subscription, error)

	ParseBurnt(log types.Log) (*LockUnlockPoolBurnt, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address, depositor []common.Address) (*LockUnlockPoolLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolLocked, sender []common.Address, depositor []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*LockUnlockPoolLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LockUnlockPoolMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*LockUnlockPoolMinted, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockUnlockPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*LockUnlockPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockUnlockPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*LockUnlockPoolOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*LockUnlockPoolPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*LockUnlockPoolPaused, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*LockUnlockPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*LockUnlockPoolReleased, error)

	FilterUnpaused(opts *bind.FilterOpts) (*LockUnlockPoolUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LockUnlockPoolUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*LockUnlockPoolUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
