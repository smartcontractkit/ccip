// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sender_dapp

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
)

var SenderDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractOnRampRouter\",\"name\":\"onRampRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidAddress\",\"type\":\"address\"}],\"name\":\"InvalidDestinationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DESTINATION_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DESTINATION_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ON_RAMP_ROUTER\",\"outputs\":[{\"internalType\":\"contractOnRampRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"sendTokens\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b50604051610eb8380380610eb883398101604081905261002f91610064565b6001600160a01b0392831660805260a0919091521660c0526100a7565b6001600160a01b038116811461006157600080fd5b50565b60008060006060848603121561007957600080fd5b83516100848161004c565b60208501516040860151919450925061009c8161004c565b809150509250925092565b60805160a05160c051610dc86100f06000396000818160c20152610238015260008181610154015261021201526000818160710152818161037f015261048a0152610dc86000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063177eeec511610050578063177eeec5146100e4578063181f5a77146101105780632ea023691461014f57600080fd5b806306c407201461006c5780630ab7dea9146100bd575b600080fd5b6100937f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100937f000000000000000000000000000000000000000000000000000000000000000081565b6100f76100f23660046109e8565b610184565b60405167ffffffffffffffff90911681526020016100b4565b604080518082018252601081527f53656e6465724461707020312e302e3000000000000000000000000000000000602082015290516100b49190610b45565b6101767f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100b4565b600073ffffffffffffffffffffffffffffffffffffffff85166101f0576040517ffdc6604f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024015b60405180910390fd5b6060600033905060006040518060e001604052808881526020018781526020017f000000000000000000000000000000000000000000000000000000000000000081526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001838a6040516020016102c392919073ffffffffffffffffffffffffffffffffffffffff92831681529116602082015260400190565b604051602081830303815290604052815260200184815250905060005b875181101561044c57610348833089848151811061030057610300610b58565b60200260200101518b858151811061031a5761031a610b58565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1661050e909392919063ffffffff16565b87818151811061035a5761035a610b58565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008984815181106103b0576103b0610b58565b60200260200101516040518363ffffffff1660e01b81526004016103f692919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af1158015610415573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104399190610b87565b508061044481610ba9565b9150506102e0565b506040517fc8658c1c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063c8658c1c906104bf908490600401610c43565b6020604051808303816000875af11580156104de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105029190610d75565b98975050505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526105a39085906105a9565b50505050565b600061060b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166106ba9092919063ffffffff16565b8051909150156106b557808060200190518101906106299190610b87565b6106b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101e7565b505050565b60606106c984846000856106d3565b90505b9392505050565b606082471015610765576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101e7565b843b6107cd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101e7565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516107f69190610d9f565b60006040518083038185875af1925050503d8060008114610833576040519150601f19603f3d011682016040523d82523d6000602084013e610838565b606091505b5091509150610848828286610853565b979650505050505050565b606083156108625750816106cc565b8251156108725782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e79190610b45565b73ffffffffffffffffffffffffffffffffffffffff811681146108c857600080fd5b50565b80356108d6816108a6565b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610951576109516108db565b604052919050565b600067ffffffffffffffff821115610973576109736108db565b5060051b60200190565b600082601f83011261098e57600080fd5b813560206109a361099e83610959565b61090a565b82815260059290921b840181019181810190868411156109c257600080fd5b8286015b848110156109dd57803583529183019183016109c6565b509695505050505050565b600080600080608085870312156109fe57600080fd5b8435610a09816108a6565b935060208581013567ffffffffffffffff80821115610a2757600080fd5b818801915088601f830112610a3b57600080fd5b8135610a4961099e82610959565b81815260059190911b8301840190848101908b831115610a6857600080fd5b938501935b82851015610a8f578435610a80816108a6565b82529385019390850190610a6d565b975050506040880135925080831115610aa757600080fd5b5050610ab58782880161097d565b925050610ac4606086016108cb565b905092959194509250565b60005b83811015610aea578181015183820152602001610ad2565b838111156105a35750506000910152565b60008151808452610b13816020860160208601610acf565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006106cc6020830184610afb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215610b9957600080fd5b815180151581146106cc57600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c01577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b600081518084526020808501945080840160005b83811015610c3857815187529582019590820190600101610c1c565b509495945050505050565b6020808252825160e083830152805161010084018190526000929182019083906101208601905b80831015610ca057835173ffffffffffffffffffffffffffffffffffffffff168252928401926001929092019190840190610c6a565b508387015193507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0925082868203016040870152610cde8185610c08565b93505050604085015160608501526060850151610d13608086018273ffffffffffffffffffffffffffffffffffffffff169052565b50608085015173ffffffffffffffffffffffffffffffffffffffff811660a08601525060a0850151818584030160c0860152610d4f8382610afb565b92505060c0850151818584030160e0860152610d6b8382610afb565b9695505050505050565b600060208284031215610d8757600080fd5b815167ffffffffffffffff811681146106cc57600080fd5b60008251610db1818460208701610acf565b919091019291505056fea164736f6c634300080d000a",
}

var SenderDappABI = SenderDappMetaData.ABI

var SenderDappBin = SenderDappMetaData.Bin

func DeploySenderDapp(auth *bind.TransactOpts, backend bind.ContractBackend, onRampRouter common.Address, destinationChainId *big.Int, destinationContract common.Address) (common.Address, *types.Transaction, *SenderDapp, error) {
	parsed, err := SenderDappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SenderDappBin), backend, onRampRouter, destinationChainId, destinationContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SenderDapp{SenderDappCaller: SenderDappCaller{contract: contract}, SenderDappTransactor: SenderDappTransactor{contract: contract}, SenderDappFilterer: SenderDappFilterer{contract: contract}}, nil
}

type SenderDapp struct {
	address common.Address
	abi     abi.ABI
	SenderDappCaller
	SenderDappTransactor
	SenderDappFilterer
}

type SenderDappCaller struct {
	contract *bind.BoundContract
}

type SenderDappTransactor struct {
	contract *bind.BoundContract
}

type SenderDappFilterer struct {
	contract *bind.BoundContract
}

type SenderDappSession struct {
	Contract     *SenderDapp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type SenderDappCallerSession struct {
	Contract *SenderDappCaller
	CallOpts bind.CallOpts
}

type SenderDappTransactorSession struct {
	Contract     *SenderDappTransactor
	TransactOpts bind.TransactOpts
}

type SenderDappRaw struct {
	Contract *SenderDapp
}

type SenderDappCallerRaw struct {
	Contract *SenderDappCaller
}

type SenderDappTransactorRaw struct {
	Contract *SenderDappTransactor
}

func NewSenderDapp(address common.Address, backend bind.ContractBackend) (*SenderDapp, error) {
	abi, err := abi.JSON(strings.NewReader(SenderDappABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindSenderDapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SenderDapp{address: address, abi: abi, SenderDappCaller: SenderDappCaller{contract: contract}, SenderDappTransactor: SenderDappTransactor{contract: contract}, SenderDappFilterer: SenderDappFilterer{contract: contract}}, nil
}

func NewSenderDappCaller(address common.Address, caller bind.ContractCaller) (*SenderDappCaller, error) {
	contract, err := bindSenderDapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenderDappCaller{contract: contract}, nil
}

func NewSenderDappTransactor(address common.Address, transactor bind.ContractTransactor) (*SenderDappTransactor, error) {
	contract, err := bindSenderDapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenderDappTransactor{contract: contract}, nil
}

func NewSenderDappFilterer(address common.Address, filterer bind.ContractFilterer) (*SenderDappFilterer, error) {
	contract, err := bindSenderDapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenderDappFilterer{contract: contract}, nil
}

func bindSenderDapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SenderDappABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_SenderDapp *SenderDappRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenderDapp.Contract.SenderDappCaller.contract.Call(opts, result, method, params...)
}

func (_SenderDapp *SenderDappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenderDapp.Contract.SenderDappTransactor.contract.Transfer(opts)
}

func (_SenderDapp *SenderDappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenderDapp.Contract.SenderDappTransactor.contract.Transact(opts, method, params...)
}

func (_SenderDapp *SenderDappCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenderDapp.Contract.contract.Call(opts, result, method, params...)
}

func (_SenderDapp *SenderDappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenderDapp.Contract.contract.Transfer(opts)
}

func (_SenderDapp *SenderDappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenderDapp.Contract.contract.Transact(opts, method, params...)
}

func (_SenderDapp *SenderDappCaller) DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SenderDapp.contract.Call(opts, &out, "DESTINATION_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SenderDapp *SenderDappSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _SenderDapp.Contract.DESTINATIONCHAINID(&_SenderDapp.CallOpts)
}

func (_SenderDapp *SenderDappCallerSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _SenderDapp.Contract.DESTINATIONCHAINID(&_SenderDapp.CallOpts)
}

func (_SenderDapp *SenderDappCaller) DESTINATIONCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SenderDapp.contract.Call(opts, &out, "DESTINATION_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SenderDapp *SenderDappSession) DESTINATIONCONTRACT() (common.Address, error) {
	return _SenderDapp.Contract.DESTINATIONCONTRACT(&_SenderDapp.CallOpts)
}

func (_SenderDapp *SenderDappCallerSession) DESTINATIONCONTRACT() (common.Address, error) {
	return _SenderDapp.Contract.DESTINATIONCONTRACT(&_SenderDapp.CallOpts)
}

func (_SenderDapp *SenderDappCaller) ONRAMPROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SenderDapp.contract.Call(opts, &out, "ON_RAMP_ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SenderDapp *SenderDappSession) ONRAMPROUTER() (common.Address, error) {
	return _SenderDapp.Contract.ONRAMPROUTER(&_SenderDapp.CallOpts)
}

func (_SenderDapp *SenderDappCallerSession) ONRAMPROUTER() (common.Address, error) {
	return _SenderDapp.Contract.ONRAMPROUTER(&_SenderDapp.CallOpts)
}

func (_SenderDapp *SenderDappCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SenderDapp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_SenderDapp *SenderDappSession) TypeAndVersion() (string, error) {
	return _SenderDapp.Contract.TypeAndVersion(&_SenderDapp.CallOpts)
}

func (_SenderDapp *SenderDappCallerSession) TypeAndVersion() (string, error) {
	return _SenderDapp.Contract.TypeAndVersion(&_SenderDapp.CallOpts)
}

func (_SenderDapp *SenderDappTransactor) SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int, executor common.Address) (*types.Transaction, error) {
	return _SenderDapp.contract.Transact(opts, "sendTokens", destinationAddress, tokens, amounts, executor)
}

func (_SenderDapp *SenderDappSession) SendTokens(destinationAddress common.Address, tokens []common.Address, amounts []*big.Int, executor common.Address) (*types.Transaction, error) {
	return _SenderDapp.Contract.SendTokens(&_SenderDapp.TransactOpts, destinationAddress, tokens, amounts, executor)
}

func (_SenderDapp *SenderDappTransactorSession) SendTokens(destinationAddress common.Address, tokens []common.Address, amounts []*big.Int, executor common.Address) (*types.Transaction, error) {
	return _SenderDapp.Contract.SendTokens(&_SenderDapp.TransactOpts, destinationAddress, tokens, amounts, executor)
}

func (_SenderDapp *SenderDapp) Address() common.Address {
	return _SenderDapp.address
}

type SenderDappInterface interface {
	DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error)

	DESTINATIONCONTRACT(opts *bind.CallOpts) (common.Address, error)

	ONRAMPROUTER(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int, executor common.Address) (*types.Transaction, error)

	Address() common.Address
}
