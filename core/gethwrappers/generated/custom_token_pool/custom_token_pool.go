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

type IPoolRampUpdate struct {
	Ramp    common.Address
	Allowed bool
}

var CustomTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ExceedsTokenLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionsError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"OffRampAllowanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"OnRampAllowanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SynthBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SynthMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"ramp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"internalType\":\"structIPool.RampUpdate[]\",\"name\":\"onRamps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ramp\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"internalType\":\"structIPool.RampUpdate[]\",\"name\":\"offRamps\",\"type\":\"tuple[]\"}],\"name\":\"applyRampUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"isOnRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockOrBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"releaseOrMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516111a43803806111a483398101604081905261002f916101a8565b8033806000816100865760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100b6576100b6816100ff565b50506001805460ff60a01b19169055506001600160a01b0381166100ed57604051634655efd160e11b815260040160405180910390fd5b6001600160a01b0316608052506101d8565b336001600160a01b038216036101575760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161007d565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000602082840312156101ba57600080fd5b81516001600160a01b03811681146101d157600080fd5b9392505050565b608051610fb16101f360003960006101030152610fb16000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c80638456cb5911610081578063e2e59b3e1161005b578063e2e59b3e146101c9578063ea6192a2146101dc578063f2fde38b146101ef57600080fd5b80638456cb59146101905780638da5cb5b14610198578063af519112146101b657600080fd5b80635c975abb116100b25780635c975abb146101525780636f32b8721461017557806379ba50971461018857600080fd5b80631d7a74a0146100d957806321df0da7146101015780633f4ba83a14610148575b600080fd5b6100ec6100e7366004610c92565b610202565b60405190151581526020015b60405180910390f35b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f8565b610150610215565b005b60015474010000000000000000000000000000000000000000900460ff166100ec565b6100ec610183366004610c92565b610227565b610150610234565b610150610336565b60005473ffffffffffffffffffffffffffffffffffffffff16610123565b6101506101c4366004610e0e565b610346565b6101506101d7366004610e72565b610552565b6101506101ea366004610e9e565b61064e565b6101506101fd366004610c92565b610742565b600061020f600483610756565b92915050565b61021d610788565b610225610809565b565b600061020f600283610756565b60015473ffffffffffffffffffffffffffffffffffffffff1633146102ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61033e610788565b610225610902565b61034e610788565b60005b825181101561044d57600083828151811061036e5761036e610ec8565b60200260200101519050806020015115610396578051610390906002906109ee565b506103a6565b80516103a490600290610a10565b505b7fbceff8f229c6dfcbf8bdcfb18726b84b0fd249b4803deb3948ff34d9040136628483815181106103d9576103d9610ec8565b6020026020010151600001518584815181106103f7576103f7610ec8565b60200260200101516020015160405161043492919073ffffffffffffffffffffffffffffffffffffffff9290921682521515602082015260400190565b60405180910390a15061044681610f26565b9050610351565b5060005b815181101561054d57600082828151811061046e5761046e610ec8565b60200260200101519050806020015115610496578051610490906004906109ee565b506104a6565b80516104a490600490610a10565b505b7fd8c3333ded377884ced3869cd0bcb9be54ea664076df1f5d39c46891203136488383815181106104d9576104d9610ec8565b6020026020010151600001518484815181106104f7576104f7610ec8565b60200260200101516020015160405161053492919073ffffffffffffffffffffffffffffffffffffffff9290921682521515602082015260400190565b60405180910390a15061054681610f26565b9050610451565b505050565b60015474010000000000000000000000000000000000000000900460ff16156105d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016102b1565b6105e033610227565b610616576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518281527f02992093bca69a36949677658a77d359b510dc6232c68f9f118f7c0127a1b147906020015b60405180910390a15050565b60015474010000000000000000000000000000000000000000900460ff16156106d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016102b1565b6106dc33610202565b610712576040517f5307f5ab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518181527fbb0b72e5f44e331506684da008a30e10d50658c29d8159f6c6ab40bf1e52e60090602001610642565b61074a610788565b61075381610a32565b50565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018301602052604081205415155b9392505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610225576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016102b1565b60015474010000000000000000000000000000000000000000900460ff1661088d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016102b1565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60015474010000000000000000000000000000000000000000900460ff1615610987576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016102b1565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586108d83390565b60006107818373ffffffffffffffffffffffffffffffffffffffff8416610b27565b60006107818373ffffffffffffffffffffffffffffffffffffffff8416610b76565b3373ffffffffffffffffffffffffffffffffffffffff821603610ab1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016102b1565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000818152600183016020526040812054610b6e5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561020f565b50600061020f565b60008181526001830160205260408120548015610c5f576000610b9a600183610f5e565b8554909150600090610bae90600190610f5e565b9050818114610c13576000866000018281548110610bce57610bce610ec8565b9060005260206000200154905080876000018481548110610bf157610bf1610ec8565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080610c2457610c24610f75565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061020f565b600091505061020f565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c8d57600080fd5b919050565b600060208284031215610ca457600080fd5b61078182610c69565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610cff57610cff610cad565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610d4c57610d4c610cad565b604052919050565b600082601f830112610d6557600080fd5b8135602067ffffffffffffffff821115610d8157610d81610cad565b610d8f818360051b01610d05565b82815260069290921b84018101918181019086841115610dae57600080fd5b8286015b84811015610e035760408189031215610dcb5760008081fd5b610dd3610cdc565b610ddc82610c69565b8152848201358015158114610df15760008081fd5b81860152835291830191604001610db2565b509695505050505050565b60008060408385031215610e2157600080fd5b823567ffffffffffffffff80821115610e3957600080fd5b610e4586838701610d54565b93506020850135915080821115610e5b57600080fd5b50610e6885828601610d54565b9150509250929050565b60008060408385031215610e8557600080fd5b82359150610e9560208401610c69565b90509250929050565b60008060408385031215610eb157600080fd5b610eba83610c69565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f5757610f57610ef7565b5060010190565b600082821015610f7057610f70610ef7565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c634300080f000a",
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

func (_CustomTokenPool *CustomTokenPoolTransactor) ApplyRampUpdates(opts *bind.TransactOpts, onRamps []IPoolRampUpdate, offRamps []IPoolRampUpdate) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "applyRampUpdates", onRamps, offRamps)
}

func (_CustomTokenPool *CustomTokenPoolSession) ApplyRampUpdates(onRamps []IPoolRampUpdate, offRamps []IPoolRampUpdate) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.ApplyRampUpdates(&_CustomTokenPool.TransactOpts, onRamps, offRamps)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) ApplyRampUpdates(onRamps []IPoolRampUpdate, offRamps []IPoolRampUpdate) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.ApplyRampUpdates(&_CustomTokenPool.TransactOpts, onRamps, offRamps)
}

func (_CustomTokenPool *CustomTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, amount *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.contract.Transact(opts, "lockOrBurn", amount, arg1)
}

func (_CustomTokenPool *CustomTokenPoolSession) LockOrBurn(amount *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.LockOrBurn(&_CustomTokenPool.TransactOpts, amount, arg1)
}

func (_CustomTokenPool *CustomTokenPoolTransactorSession) LockOrBurn(amount *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _CustomTokenPool.Contract.LockOrBurn(&_CustomTokenPool.TransactOpts, amount, arg1)
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

type CustomTokenPoolOffRampAllowanceSetIterator struct {
	Event *CustomTokenPoolOffRampAllowanceSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolOffRampAllowanceSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolOffRampAllowanceSet)
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
		it.Event = new(CustomTokenPoolOffRampAllowanceSet)
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

func (it *CustomTokenPoolOffRampAllowanceSetIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolOffRampAllowanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolOffRampAllowanceSet struct {
	OnRamp  common.Address
	Allowed bool
	Raw     types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterOffRampAllowanceSet(opts *bind.FilterOpts) (*CustomTokenPoolOffRampAllowanceSetIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "OffRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolOffRampAllowanceSetIterator{contract: _CustomTokenPool.contract, event: "OffRampAllowanceSet", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchOffRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOffRampAllowanceSet) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "OffRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolOffRampAllowanceSet)
				if err := _CustomTokenPool.contract.UnpackLog(event, "OffRampAllowanceSet", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseOffRampAllowanceSet(log types.Log) (*CustomTokenPoolOffRampAllowanceSet, error) {
	event := new(CustomTokenPoolOffRampAllowanceSet)
	if err := _CustomTokenPool.contract.UnpackLog(event, "OffRampAllowanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CustomTokenPoolOnRampAllowanceSetIterator struct {
	Event *CustomTokenPoolOnRampAllowanceSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CustomTokenPoolOnRampAllowanceSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomTokenPoolOnRampAllowanceSet)
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
		it.Event = new(CustomTokenPoolOnRampAllowanceSet)
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

func (it *CustomTokenPoolOnRampAllowanceSetIterator) Error() error {
	return it.fail
}

func (it *CustomTokenPoolOnRampAllowanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CustomTokenPoolOnRampAllowanceSet struct {
	OnRamp  common.Address
	Allowed bool
	Raw     types.Log
}

func (_CustomTokenPool *CustomTokenPoolFilterer) FilterOnRampAllowanceSet(opts *bind.FilterOpts) (*CustomTokenPoolOnRampAllowanceSetIterator, error) {

	logs, sub, err := _CustomTokenPool.contract.FilterLogs(opts, "OnRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return &CustomTokenPoolOnRampAllowanceSetIterator{contract: _CustomTokenPool.contract, event: "OnRampAllowanceSet", logs: logs, sub: sub}, nil
}

func (_CustomTokenPool *CustomTokenPoolFilterer) WatchOnRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOnRampAllowanceSet) (event.Subscription, error) {

	logs, sub, err := _CustomTokenPool.contract.WatchLogs(opts, "OnRampAllowanceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CustomTokenPoolOnRampAllowanceSet)
				if err := _CustomTokenPool.contract.UnpackLog(event, "OnRampAllowanceSet", log); err != nil {
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

func (_CustomTokenPool *CustomTokenPoolFilterer) ParseOnRampAllowanceSet(log types.Log) (*CustomTokenPoolOnRampAllowanceSet, error) {
	event := new(CustomTokenPoolOnRampAllowanceSet)
	if err := _CustomTokenPool.contract.UnpackLog(event, "OnRampAllowanceSet", log); err != nil {
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
	case _CustomTokenPool.abi.Events["OffRampAllowanceSet"].ID:
		return _CustomTokenPool.ParseOffRampAllowanceSet(log)
	case _CustomTokenPool.abi.Events["OnRampAllowanceSet"].ID:
		return _CustomTokenPool.ParseOnRampAllowanceSet(log)
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

func (CustomTokenPoolOffRampAllowanceSet) Topic() common.Hash {
	return common.HexToHash("0xd8c3333ded377884ced3869cd0bcb9be54ea664076df1f5d39c4689120313648")
}

func (CustomTokenPoolOnRampAllowanceSet) Topic() common.Hash {
	return common.HexToHash("0xbceff8f229c6dfcbf8bdcfb18726b84b0fd249b4803deb3948ff34d904013662")
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

	ApplyRampUpdates(opts *bind.TransactOpts, onRamps []IPoolRampUpdate, offRamps []IPoolRampUpdate) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, amount *big.Int, arg1 common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int) (*types.Transaction, error)

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

	FilterOffRampAllowanceSet(opts *bind.FilterOpts) (*CustomTokenPoolOffRampAllowanceSetIterator, error)

	WatchOffRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOffRampAllowanceSet) (event.Subscription, error)

	ParseOffRampAllowanceSet(log types.Log) (*CustomTokenPoolOffRampAllowanceSet, error)

	FilterOnRampAllowanceSet(opts *bind.FilterOpts) (*CustomTokenPoolOnRampAllowanceSetIterator, error)

	WatchOnRampAllowanceSet(opts *bind.WatchOpts, sink chan<- *CustomTokenPoolOnRampAllowanceSet) (event.Subscription, error)

	ParseOnRampAllowanceSet(log types.Log) (*CustomTokenPoolOnRampAllowanceSet, error)

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
