package weth9

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
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
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

var WETH9MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c0604052600d60809081526c2bb930b83832b21022ba3432b960991b60a05260009061002c9082610116565b506040805180820190915260048152630ae8aa8960e31b60208201526001906100559082610116565b506002805460ff1916601217905534801561006f57600080fd5b506101d5565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061009f57607f821691505b6020821081036100bf57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610111576000816000526020600020601f850160051c810160208610156100ee5750805b601f850160051c820191505b8181101561010d578281556001016100fa565b5050505b505050565b81516001600160401b0381111561012f5761012f610075565b6101438161013d845461008b565b846100c5565b602080601f83116001811461017857600084156101605750858301515b600019600386901b1c1916600185901b17855561010d565b600085815260208120601f198616915b828110156101a757888601518255948401946001909101908401610188565b50858210156101c55787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610936806101e46000396000f3fe6080604052600436106100c05760003560e01c8063313ce56711610074578063a9059cbb1161004e578063a9059cbb146101fa578063d0e30db01461021a578063dd62ed3e1461022257600080fd5b8063313ce5671461018c57806370a08231146101b857806395d89b41146101e557600080fd5b806318160ddd116100a557806318160ddd1461012f57806323b872dd1461014c5780632e1a7d4d1461016c57600080fd5b806306fdde03146100d4578063095ea7b3146100ff57600080fd5b366100cf576100cd61025a565b005b600080fd5b3480156100e057600080fd5b506100e96102b5565b6040516100f6919061071e565b60405180910390f35b34801561010b57600080fd5b5061011f61011a3660046107b4565b610343565b60405190151581526020016100f6565b34801561013b57600080fd5b50475b6040519081526020016100f6565b34801561015857600080fd5b5061011f6101673660046107de565b6103bd565b34801561017857600080fd5b506100cd61018736600461081a565b6105c4565b34801561019857600080fd5b506002546101a69060ff1681565b60405160ff90911681526020016100f6565b3480156101c457600080fd5b5061013e6101d3366004610833565b60036020526000908152604090205481565b3480156101f157600080fd5b506100e96106f3565b34801561020657600080fd5b5061011f6102153660046107b4565b610700565b6100cd610714565b34801561022e57600080fd5b5061013e61023d36600461084e565b600460209081526000928352604080842090915290825290205481565b33600090815260036020526040812080543492906102799084906108b0565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b600080546102c2906108c3565b80601f01602080910402602001604051908101604052809291908181526020018280546102ee906108c3565b801561033b5780601f106103105761010080835404028352916020019161033b565b820191906000526020600020905b81548152906001019060200180831161031e57829003601f168201915b505050505081565b33600081815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103ab9086815260200190565b60405180910390a35060015b92915050565b73ffffffffffffffffffffffffffffffffffffffff83166000908152600360205260408120548211156103ef57600080fd5b73ffffffffffffffffffffffffffffffffffffffff84163314801590610455575073ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091529020546fffffffffffffffffffffffffffffffff14155b156104dd5773ffffffffffffffffffffffffffffffffffffffff8416600090815260046020908152604080832033845290915290205482111561049757600080fd5b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152812080548492906104d7908490610916565b90915550505b73ffffffffffffffffffffffffffffffffffffffff841660009081526003602052604081208054849290610512908490610916565b909155505073ffffffffffffffffffffffffffffffffffffffff83166000908152600360205260408120805484929061054c9084906108b0565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516105b291815260200190565b60405180910390a35060019392505050565b336000908152600360205260409020548111156105e057600080fd5b33600090815260036020526040812080548392906105ff908490610916565b9091555050604051600090339083908381818185875af1925050503d8060008114610646576040519150601f19603f3d011682016040523d82523d6000602084013e61064b565b606091505b50509050806106ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f57455448393a20455448207472616e73666572206661696c6564000000000000604482015260640160405180910390fd5b60405182815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a25050565b600180546102c2906108c3565b600061070d3384846103bd565b9392505050565b61071c61025a565b565b60006020808352835180602085015260005b8181101561074c57858101830151858201604001528201610730565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146107af57600080fd5b919050565b600080604083850312156107c757600080fd5b6107d08361078b565b946020939093013593505050565b6000806000606084860312156107f357600080fd5b6107fc8461078b565b925061080a6020850161078b565b9150604084013590509250925092565b60006020828403121561082c57600080fd5b5035919050565b60006020828403121561084557600080fd5b61070d8261078b565b6000806040838503121561086157600080fd5b61086a8361078b565b91506108786020840161078b565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156103b7576103b7610881565b600181811c908216806108d757607f821691505b602082108103610910577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b818103818111156103b7576103b761088156fea164736f6c6343000818000a",
}

var WETH9ABI = WETH9MetaData.ABI

var WETH9Bin = WETH9MetaData.Bin

func DeployWETH9(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *generated.Transaction, *WETH9, error) {
	parsed, err := WETH9MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}
	if generated.IsZKSync(backend) {
		address, ethTx, contractBind, _ := generated.DeployContract(auth, parsed, common.FromHex(WETH9ZKBin), backend)
		contractReturn := &WETH9{address: address, abi: *parsed, WETH9Caller: WETH9Caller{contract: contractBind}, WETH9Transactor: WETH9Transactor{contract: contractBind}, WETH9Filterer: WETH9Filterer{contract: contractBind}}
		return address, ethTx, contractReturn, err
	}
	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WETH9Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, &generated.Transaction{Transaction: tx, HashZks: tx.Hash()}, &WETH9{address: address, abi: *parsed, WETH9Caller: WETH9Caller{contract: contract}, WETH9Transactor: WETH9Transactor{contract: contract}, WETH9Filterer: WETH9Filterer{contract: contract}}, nil
}

type WETH9 struct {
	address common.Address
	abi     abi.ABI
	WETH9Caller
	WETH9Transactor
	WETH9Filterer
}

type WETH9Caller struct {
	contract *bind.BoundContract
}

type WETH9Transactor struct {
	contract *bind.BoundContract
}

type WETH9Filterer struct {
	contract *bind.BoundContract
}

type WETH9Session struct {
	Contract     *WETH9
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type WETH9CallerSession struct {
	Contract *WETH9Caller
	CallOpts bind.CallOpts
}

type WETH9TransactorSession struct {
	Contract     *WETH9Transactor
	TransactOpts bind.TransactOpts
}

type WETH9Raw struct {
	Contract *WETH9
}

type WETH9CallerRaw struct {
	Contract *WETH9Caller
}

type WETH9TransactorRaw struct {
	Contract *WETH9Transactor
}

func NewWETH9(address common.Address, backend bind.ContractBackend) (*WETH9, error) {
	abi, err := abi.JSON(strings.NewReader(WETH9ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindWETH9(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WETH9{address: address, abi: abi, WETH9Caller: WETH9Caller{contract: contract}, WETH9Transactor: WETH9Transactor{contract: contract}, WETH9Filterer: WETH9Filterer{contract: contract}}, nil
}

func NewWETH9Caller(address common.Address, caller bind.ContractCaller) (*WETH9Caller, error) {
	contract, err := bindWETH9(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WETH9Caller{contract: contract}, nil
}

func NewWETH9Transactor(address common.Address, transactor bind.ContractTransactor) (*WETH9Transactor, error) {
	contract, err := bindWETH9(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WETH9Transactor{contract: contract}, nil
}

func NewWETH9Filterer(address common.Address, filterer bind.ContractFilterer) (*WETH9Filterer, error) {
	contract, err := bindWETH9(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WETH9Filterer{contract: contract}, nil
}

func bindWETH9(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WETH9MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_WETH9 *WETH9Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETH9.Contract.WETH9Caller.contract.Call(opts, result, method, params...)
}

func (_WETH9 *WETH9Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH9.Contract.WETH9Transactor.contract.Transfer(opts)
}

func (_WETH9 *WETH9Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETH9.Contract.WETH9Transactor.contract.Transact(opts, method, params...)
}

func (_WETH9 *WETH9CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETH9.Contract.contract.Call(opts, result, method, params...)
}

func (_WETH9 *WETH9TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH9.Contract.contract.Transfer(opts)
}

func (_WETH9 *WETH9TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETH9.Contract.contract.Transact(opts, method, params...)
}

func (_WETH9 *WETH9Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETH9.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WETH9 *WETH9Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WETH9.Contract.Allowance(&_WETH9.CallOpts, arg0, arg1)
}

func (_WETH9 *WETH9CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WETH9.Contract.Allowance(&_WETH9.CallOpts, arg0, arg1)
}

func (_WETH9 *WETH9Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETH9.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WETH9 *WETH9Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WETH9.Contract.BalanceOf(&_WETH9.CallOpts, arg0)
}

func (_WETH9 *WETH9CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WETH9.Contract.BalanceOf(&_WETH9.CallOpts, arg0)
}

func (_WETH9 *WETH9Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WETH9.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_WETH9 *WETH9Session) Decimals() (uint8, error) {
	return _WETH9.Contract.Decimals(&_WETH9.CallOpts)
}

func (_WETH9 *WETH9CallerSession) Decimals() (uint8, error) {
	return _WETH9.Contract.Decimals(&_WETH9.CallOpts)
}

func (_WETH9 *WETH9Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WETH9.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_WETH9 *WETH9Session) Name() (string, error) {
	return _WETH9.Contract.Name(&_WETH9.CallOpts)
}

func (_WETH9 *WETH9CallerSession) Name() (string, error) {
	return _WETH9.Contract.Name(&_WETH9.CallOpts)
}

func (_WETH9 *WETH9Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WETH9.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_WETH9 *WETH9Session) Symbol() (string, error) {
	return _WETH9.Contract.Symbol(&_WETH9.CallOpts)
}

func (_WETH9 *WETH9CallerSession) Symbol() (string, error) {
	return _WETH9.Contract.Symbol(&_WETH9.CallOpts)
}

func (_WETH9 *WETH9Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WETH9.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WETH9 *WETH9Session) TotalSupply() (*big.Int, error) {
	return _WETH9.Contract.TotalSupply(&_WETH9.CallOpts)
}

func (_WETH9 *WETH9CallerSession) TotalSupply() (*big.Int, error) {
	return _WETH9.Contract.TotalSupply(&_WETH9.CallOpts)
}

func (_WETH9 *WETH9Transactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9.contract.Transact(opts, "approve", guy, wad)
}

func (_WETH9 *WETH9Session) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9.Contract.Approve(&_WETH9.TransactOpts, guy, wad)
}

func (_WETH9 *WETH9TransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9.Contract.Approve(&_WETH9.TransactOpts, guy, wad)
}

func (_WETH9 *WETH9Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH9.contract.Transact(opts, "deposit")
}

func (_WETH9 *WETH9Session) Deposit() (*types.Transaction, error) {
	return _WETH9.Contract.Deposit(&_WETH9.TransactOpts)
}

func (_WETH9 *WETH9TransactorSession) Deposit() (*types.Transaction, error) {
	return _WETH9.Contract.Deposit(&_WETH9.TransactOpts)
}

func (_WETH9 *WETH9Transactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9.contract.Transact(opts, "transfer", dst, wad)
}

func (_WETH9 *WETH9Session) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9.Contract.Transfer(&_WETH9.TransactOpts, dst, wad)
}

func (_WETH9 *WETH9TransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9.Contract.Transfer(&_WETH9.TransactOpts, dst, wad)
}

func (_WETH9 *WETH9Transactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9.contract.Transact(opts, "transferFrom", src, dst, wad)
}

func (_WETH9 *WETH9Session) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9.Contract.TransferFrom(&_WETH9.TransactOpts, src, dst, wad)
}

func (_WETH9 *WETH9TransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WETH9.Contract.TransferFrom(&_WETH9.TransactOpts, src, dst, wad)
}

func (_WETH9 *WETH9Transactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _WETH9.contract.Transact(opts, "withdraw", wad)
}

func (_WETH9 *WETH9Session) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WETH9.Contract.Withdraw(&_WETH9.TransactOpts, wad)
}

func (_WETH9 *WETH9TransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WETH9.Contract.Withdraw(&_WETH9.TransactOpts, wad)
}

func (_WETH9 *WETH9Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH9.contract.RawTransact(opts, nil)
}

func (_WETH9 *WETH9Session) Receive() (*types.Transaction, error) {
	return _WETH9.Contract.Receive(&_WETH9.TransactOpts)
}

func (_WETH9 *WETH9TransactorSession) Receive() (*types.Transaction, error) {
	return _WETH9.Contract.Receive(&_WETH9.TransactOpts)
}

type WETH9ApprovalIterator struct {
	Event *WETH9Approval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WETH9ApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETH9Approval)
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
		it.Event = new(WETH9Approval)
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

func (it *WETH9ApprovalIterator) Error() error {
	return it.fail
}

func (it *WETH9ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WETH9Approval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log
}

func (_WETH9 *WETH9Filterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*WETH9ApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _WETH9.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &WETH9ApprovalIterator{contract: _WETH9.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_WETH9 *WETH9Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WETH9Approval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _WETH9.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WETH9Approval)
				if err := _WETH9.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_WETH9 *WETH9Filterer) ParseApproval(log types.Log) (*WETH9Approval, error) {
	event := new(WETH9Approval)
	if err := _WETH9.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WETH9DepositIterator struct {
	Event *WETH9Deposit

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WETH9DepositIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETH9Deposit)
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
		it.Event = new(WETH9Deposit)
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

func (it *WETH9DepositIterator) Error() error {
	return it.fail
}

func (it *WETH9DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WETH9Deposit struct {
	Dst common.Address
	Wad *big.Int
	Raw types.Log
}

func (_WETH9 *WETH9Filterer) FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*WETH9DepositIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WETH9.contract.FilterLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return &WETH9DepositIterator{contract: _WETH9.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

func (_WETH9 *WETH9Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WETH9Deposit, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WETH9.contract.WatchLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WETH9Deposit)
				if err := _WETH9.contract.UnpackLog(event, "Deposit", log); err != nil {
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

func (_WETH9 *WETH9Filterer) ParseDeposit(log types.Log) (*WETH9Deposit, error) {
	event := new(WETH9Deposit)
	if err := _WETH9.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WETH9TransferIterator struct {
	Event *WETH9Transfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WETH9TransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETH9Transfer)
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
		it.Event = new(WETH9Transfer)
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

func (it *WETH9TransferIterator) Error() error {
	return it.fail
}

func (it *WETH9TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WETH9Transfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log
}

func (_WETH9 *WETH9Filterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*WETH9TransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WETH9.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &WETH9TransferIterator{contract: _WETH9.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_WETH9 *WETH9Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WETH9Transfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WETH9.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WETH9Transfer)
				if err := _WETH9.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_WETH9 *WETH9Filterer) ParseTransfer(log types.Log) (*WETH9Transfer, error) {
	event := new(WETH9Transfer)
	if err := _WETH9.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WETH9WithdrawalIterator struct {
	Event *WETH9Withdrawal

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WETH9WithdrawalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETH9Withdrawal)
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
		it.Event = new(WETH9Withdrawal)
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

func (it *WETH9WithdrawalIterator) Error() error {
	return it.fail
}

func (it *WETH9WithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WETH9Withdrawal struct {
	Src common.Address
	Wad *big.Int
	Raw types.Log
}

func (_WETH9 *WETH9Filterer) FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*WETH9WithdrawalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WETH9.contract.FilterLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return &WETH9WithdrawalIterator{contract: _WETH9.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

func (_WETH9 *WETH9Filterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WETH9Withdrawal, src []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WETH9.contract.WatchLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WETH9Withdrawal)
				if err := _WETH9.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

func (_WETH9 *WETH9Filterer) ParseWithdrawal(log types.Log) (*WETH9Withdrawal, error) {
	event := new(WETH9Withdrawal)
	if err := _WETH9.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_WETH9 *WETH9) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _WETH9.abi.Events["Approval"].ID:
		return _WETH9.ParseApproval(log)
	case _WETH9.abi.Events["Deposit"].ID:
		return _WETH9.ParseDeposit(log)
	case _WETH9.abi.Events["Transfer"].ID:
		return _WETH9.ParseTransfer(log)
	case _WETH9.abi.Events["Withdrawal"].ID:
		return _WETH9.ParseWithdrawal(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (WETH9Approval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (WETH9Deposit) Topic() common.Hash {
	return common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c")
}

func (WETH9Transfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (WETH9Withdrawal) Topic() common.Hash {
	return common.HexToHash("0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65")
}

func (_WETH9 *WETH9) Address() common.Address {
	return _WETH9.address
}

type WETH9Interface interface {
	Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	Name(opts *bind.CallOpts) (string, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error)

	Deposit(opts *bind.TransactOpts) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*WETH9ApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *WETH9Approval, src []common.Address, guy []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*WETH9Approval, error)

	FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*WETH9DepositIterator, error)

	WatchDeposit(opts *bind.WatchOpts, sink chan<- *WETH9Deposit, dst []common.Address) (event.Subscription, error)

	ParseDeposit(log types.Log) (*WETH9Deposit, error)

	FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*WETH9TransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *WETH9Transfer, src []common.Address, dst []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*WETH9Transfer, error)

	FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*WETH9WithdrawalIterator, error)

	WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WETH9Withdrawal, src []common.Address) (event.Subscription, error)

	ParseWithdrawal(log types.Log) (*WETH9Withdrawal, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}

var WETH9ZKBin = ("0x0002000000000002000200000000000200000000030100190000006003300270000000da0330019700010000003103550000008004000039000000400040043f0000000100200190000000300000c13d000000040030008c000000670000413d000000000201043b000000e002200270000000e10020009c000000910000a13d000000e20020009c000000bb0000213d000000e60020009c000001a70000613d000000e70020009c000001710000613d000000e80020009c000001b90000c13d0000000001000416000000000001004b000001b90000c13d0000000103000039000000000203041a000000010520019000000001012002700000007f0410018f00000000010460190000001f0010008c00000000060000390000000106002039000000000662013f0000000100600190000000610000c13d000000800010043f000000000005004b000001d50000c13d000000fd01200197000000a00010043f000000000004004b000000c001000039000000a0010060390000022d0000013d0000000001000416000000000001004b000001b90000c13d000000000100041a000000010210019000000001031002700000007f0330618f0000001f0030008c00000000010000390000000101002039000000000012004b000000610000c13d000000200030008c000000540000413d000200000003001d000000000000043f0000000001000414000000da0010009c000000da01008041000000c001100210000000db011001c70000801002000039036503600000040f0000000100200190000001b90000613d000000000101043b00000002020000290000001f0220003900000005022002700000000002210019000000000021004b000000540000813d000000000001041b0000000101100039000000000021004b000000500000413d000000dc01000041000000000010041b0000000103000039000000000103041a000000010010019000000001041002700000007f0440618f0000001f0040008c00000000020000390000000102002039000000000121013f0000000100100190000001240000613d000000f101000041000000000010043f0000002201000039000000040010043f000000f2010000410000036700010430000000000003004b000001b90000c13d0000000001000411000000000010043f0000000301000039000000200010043f0000000001000414000000da0010009c000000da01008041000000c001100210000000df011001c70000801002000039036503600000040f0000000100200190000001b90000613d000000000101043b000000000301041a0000000002000416000000000023001a000002540000413d0000000003230019000000000031041b000000400100043d0000000000210435000000da0010009c000000da0100804100000040011002100000000002000414000000da0020009c000000da02008041000000c002200210000000000112019f000000db011001c70000800d020000390000000203000039000000e00400004100000000050004110365035b0000040f0000000100200190000001b90000613d0000000001000019000003660001042e000000e90020009c000000db0000a13d000000ea0020009c0000019f0000613d000000eb0020009c000001600000613d000000ec0020009c000001b90000c13d000000240030008c000001b90000413d0000000002000416000000000002004b000001b90000c13d0000000401100370000000000101043b000200000001001d0000000001000411000000000010043f0000000301000039000000200010043f0000000001000414000000da0010009c000000da01008041000000c001100210000000df011001c70000801002000039036503600000040f00000002030000290000000100200190000001b90000613d000000000101043b000000000201041a000000000232004b000001b90000413d000000000021041b00000000010004140000000002000411000000040020008c000001da0000c13d00000001020000390000000001000031000001e90000013d000000e30020009c000001b00000613d000000e40020009c000001840000613d000000e50020009c000001b90000c13d000000440030008c000001b90000413d0000000002000416000000000002004b000001b90000c13d0000000402100370000000000202043b000000ef0020009c000001b90000213d0000002401100370000000000101043b000000ef0010009c000001b90000213d000000000020043f000200000001001d0000000401000039000000200010043f00000040020000390000000001000019036503370000040f0000000202000029000000000020043f000000200010043f000000000100001900000040020000390000017f0000013d000000ed0020009c000001490000613d000000ee0020009c000001b90000c13d000000440030008c000001b90000413d0000000002000416000000000002004b000001b90000c13d0000000402100370000000000202043b000200000002001d000000ef0020009c000001b90000213d0000002401100370000000000101043b000100000001001d0000000001000411000000000010043f0000000401000039000000200010043f0000000001000414000000da0010009c000000da01008041000000c001100210000000df011001c70000801002000039036503600000040f0000000100200190000001b90000613d000000000101043b0000000202000029000000000020043f000000200010043f0000000001000414000000da0010009c000000da01008041000000c001100210000000df011001c70000801002000039036503600000040f0000000100200190000001b90000613d000000000101043b0000000102000029000000000021041b000000400100043d0000000000210435000000da0010009c000000da0100804100000040011002100000000002000414000000da0020009c000000da02008041000000c002200210000000000112019f000000db011001c70000800d020000390000000303000039000000fb04000041000000000500041100000002060000290365035b0000040f0000000100200190000001b90000613d000000400100043d00000001020000390000000000210435000000da0010009c000000da010080410000004001100210000000f3011001c7000003660001042e0000001f0040008c0000013d0000a13d000200000004001d000000000030043f0000000001000414000000da0010009c000000da01008041000000c001100210000000db011001c70000801002000039036503600000040f0000000100200190000001b90000613d000000000101043b00000002020000290000001f0220003900000005022002700000000002210019000000000021004b00000001030000390000013d0000813d000000000001041b0000000101100039000000000021004b000001390000413d000000dd01000041000000000013041b0000000201000039000000000201041a000000fd0220019700000012022001bf000000000021041b000000200100003900000100001004430000012000000443000000de01000041000003660001042e0000000001000416000000000001004b000001b90000c13d000000000200041a000000010420019000000001012002700000007f0310018f00000000010360190000001f0010008c00000000050000390000000105002039000000000552013f0000000100500190000000610000c13d000000800010043f000000000004004b000001c70000c13d000000fd01200197000000a00010043f000000000003004b000000c001000039000000a0010060390000022d0000013d000000640030008c000001b90000413d0000000002000416000000000002004b000001b90000c13d0000000402100370000000000402043b000000ef0040009c000001b90000213d0000002402100370000000000202043b000000ef0020009c000001b90000213d0000004401100370000000000301043b0000000001040019000001be0000013d000000240030008c000001b90000413d0000000002000416000000000002004b000001b90000c13d0000000401100370000000000101043b000000ef0010009c000001b90000213d000000000010043f0000000301000039000000200010043f00000040020000390000000001000019036503370000040f000000000101041a000000800010043f000000f001000041000003660001042e0000000001000411000000000010043f0000000301000039000000200010043f0000000001000414000000da0010009c000000da01008041000000c001100210000000df011001c70000801002000039036503600000040f0000000100200190000001b90000613d000000000101043b000000000301041a0000000002000416000000000023001a000002540000413d0000000003230019000000000031041b000000400100043d0000000000210435000000da0010009c000000da0100804100000040011002100000000002000414000000830000013d0000000001000416000000000001004b000001b90000c13d00000000010004100365034c0000040f000000800010043f000000f001000041000003660001042e0000000001000416000000000001004b000001b90000c13d0000000201000039000000000101041a000000ff0110018f000000800010043f000000f001000041000003660001042e000000440030008c000001b90000413d0000000002000416000000000002004b000001b90000c13d0000000402100370000000000202043b000000ef0020009c000001bb0000a13d000000000100001900000367000104300000002401100370000000000301043b0000000001000411036502810000040f0000000101000039000000400200043d0000000000120435000000da0020009c000000da020080410000004001200210000000f3011001c7000003660001042e000000000000043f000000020020008c000001d80000413d000000fc0200004100000000040000190000000003040019000000000402041a000000a005300039000000000045043500000001022000390000002004300039000000000014004b000001cc0000413d0000022c0000013d000000000030043f000000020020008c000002220000813d000000a0010000390000022d0000013d000000da0010009c000000da01008041000000c001100210000000000003004b000001e30000613d000000f5011001c70000800902000039000000000400041100000000050000190365035b0000040f000000020300002900010000000103550000006001100270000000da0010019d000000da01100197000000000001004b000002110000613d0000001f04100039000000fe044001970000003f04400039000000fe05400197000000400400043d0000000005540019000000000045004b00000000060000390000000106004039000000f60050009c0000024e0000213d00000001006001900000024e0000c13d000000400050043f0000000006140436000000fe091001980000001f0410018f00000000019600190000000105000367000002040000613d000000000705034f000000007807043c0000000006860436000000000016004b000002000000c13d000000000004004b000002110000613d000000000695034f0000000304400210000000000501043300000000054501cf000000000545022f000000000606043b0000010004400089000000000646022f00000000044601cf000000000454019f0000000000410435000000400100043d00000001002001900000023e0000613d0000000000310435000000da0010009c000000da0100804100000040011002100000000002000414000000da0020009c000000da02008041000000c002200210000000000112019f000000db011001c70000800d020000390000000203000039000000fa040000410000008b0000013d000000f40200004100000000040000190000000003040019000000000402041a000000a005300039000000000045043500000001022000390000002004300039000000000014004b000002240000413d000000c001300039000000800210008a00000080010000390365025a0000040f000000400100043d000200000001001d00000080020000390365026c0000040f00000002020000290000000001210049000000da0010009c000000da010080410000006001100210000000da0020009c000000da020080410000004002200210000000000121019f000003660001042e0000004402100039000000f703000041000000000032043500000024021000390000001a030000390000000000320435000000f8020000410000000000210435000000040210003900000020030000390000000000320435000000da0010009c000000da010080410000004001100210000000f9011001c70000036700010430000000f101000041000000000010043f0000004101000039000000040010043f000000f2010000410000036700010430000000f101000041000000000010043f0000001101000039000000040010043f000000f20100004100000367000104300000001f02200039000000fe022001970000000001120019000000000021004b00000000020000390000000102004039000000f60010009c000002660000213d0000000100200190000002660000c13d000000400010043f000000000001042d000000f101000041000000000010043f0000004101000039000000040010043f000000f201000041000003670001043000000020030000390000000004310436000000003202043400000000002404350000004001100039000000000002004b0000027b0000613d000000000400001900000000054100190000000006430019000000000606043300000000006504350000002004400039000000000024004b000002740000413d000000000321001900000000000304350000001f02200039000000fe022001970000000001210019000000000001042d0003000000000002000200000003001d000100000002001d000000ef01100197000300000001001d000000000010043f0000000301000039000000200010043f0000000001000414000000da0010009c000000da01008041000000c001100210000000df011001c70000801002000039036503600000040f00000001002001900000032e0000613d000000000101043b000000000101041a000000020010006c0000032e0000413d0000000002000411000000030020006b000002f20000613d0000000401000039000000200010043f0000000001000414000000da0010009c000000da01008041000000c001100210000000df011001c70000801002000039036503600000040f00000001002001900000032e0000613d000000000101043b0000000002000411000000000020043f000000200010043f0000000001000414000000da0010009c000000da01008041000000c001100210000000df011001c70000801002000039036503600000040f00000001002001900000032e0000613d000000000101043b000000000101041a000000ff0010009c000002f20000613d0000000301000029000000000010043f0000000401000039000000200010043f0000000001000414000000da0010009c000000da01008041000000c001100210000000df011001c70000801002000039036503600000040f00000001002001900000032e0000613d000000000101043b0000000002000411000000000020043f000000200010043f0000000001000414000000da0010009c000000da01008041000000c001100210000000df011001c70000801002000039036503600000040f00000001002001900000032e0000613d000000000101043b000000000101041a000000020010006c0000032e0000413d0000000301000029000000000010043f0000000401000039000000200010043f0000000001000414000000da0010009c000000da01008041000000c001100210000000df011001c70000801002000039036503600000040f00000001002001900000032e0000613d000000000101043b0000000002000411000000000020043f000000200010043f0000000001000414000000da0010009c000000da01008041000000c001100210000000df011001c70000801002000039036503600000040f00000001002001900000032e0000613d000000000101043b000000000201041a000000020220006c000003300000413d000000000021041b0000000301000029000000000010043f0000000301000039000000200010043f0000000001000414000000da0010009c000000da01008041000000c001100210000000df011001c70000801002000039036503600000040f00000001002001900000032e0000613d000000000101043b000000000201041a000000020220006c000003300000413d000000000021041b0000000101000029000000ef01100197000100000001001d000000000010043f0000000301000039000000200010043f0000000001000414000000da0010009c000000da01008041000000c001100210000000df011001c70000801002000039036503600000040f00000001002001900000032e0000613d000000000101043b000000000201041a0000000203000029000000000032001a000003300000413d0000000002320019000000000021041b000000400100043d0000000000310435000000da0010009c000000da0100804100000040011002100000000002000414000000da0020009c000000da02008041000000c002200210000000000112019f000000db011001c70000800d0200003900000003030000390000010004000041000000030500002900000001060000290365035b0000040f00000001002001900000032e0000613d000000000001042d00000000010000190000036700010430000000f101000041000000000010043f0000001101000039000000040010043f000000f2010000410000036700010430000000000001042f000000da0010009c000000da010080410000004001100210000000da0020009c000000da020080410000006002200210000000000112019f0000000002000414000000da0020009c000000da02008041000000c002200210000000000112019f000000f5011001c70000801002000039036503600000040f00000001002001900000034a0000613d000000000101043b000000000001042d000000000100001900000367000104300000010102000041000000000020044300000004001004430000000001000414000000da0010009c000000da01008041000000c00110021000000102011001c70000800a02000039036503600000040f00000001002001900000035a0000613d000000000101043b000000000001042d000000000001042f0000035e002104210000000102000039000000000001042d0000000002000019000000000001042d00000363002104230000000102000039000000000001042d0000000002000019000000000001042d0000036500000432000003660001042e000003670001043000000000000000000000000000000000000000000000000000000000ffffffff0200000000000000000000000000000000000020000000000000000000000000577261707065642045746865720000000000000000000000000000000000001a574554480000000000000000000000000000000000000000000000000000000800000002000000000000000000000000000000400000010000000000000000000200000000000000000000000000000000000040000000000000000000000000e1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c00000000000000000000000000000000000000000000000000000000313ce56600000000000000000000000000000000000000000000000000000000a9059cba00000000000000000000000000000000000000000000000000000000a9059cbb00000000000000000000000000000000000000000000000000000000d0e30db000000000000000000000000000000000000000000000000000000000dd62ed3e00000000000000000000000000000000000000000000000000000000313ce5670000000000000000000000000000000000000000000000000000000070a082310000000000000000000000000000000000000000000000000000000095d89b410000000000000000000000000000000000000000000000000000000018160ddc0000000000000000000000000000000000000000000000000000000018160ddd0000000000000000000000000000000000000000000000000000000023b872dd000000000000000000000000000000000000000000000000000000002e1a7d4d0000000000000000000000000000000000000000000000000000000006fdde0300000000000000000000000000000000000000000000000000000000095ea7b3000000000000000000000000ffffffffffffffffffffffffffffffffffffffff00000000000000000000000000000000000000200000008000000000000000004e487b710000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000240000000000000000000000000000000000000000000000000000000000000020000000000000000000000000b10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffff57455448393a20455448207472616e73666572206661696c656400000000000008c379a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000640000000000000000000000007fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b658c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe000000000000000000000000000000000ffffffffffffffffffffffffffffffffddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9cc7f708afc65944829bd487b90b72536b1951864fbfc14e125fc972a6507f390200000200000000000000000000000000000024000000000000000000000000")
