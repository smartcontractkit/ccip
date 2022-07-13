// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package toll_sender_dapp

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

var TollSenderDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractAny2EVMTollOnRampRouterInterface\",\"name\":\"onRampRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidAddress\",\"type\":\"address\"}],\"name\":\"InvalidDestinationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DESTINATION_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DESTINATION_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ON_RAMP_ROUTER\",\"outputs\":[{\"internalType\":\"contractAny2EVMTollOnRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"sendTokens\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b50604051610e74380380610e7483398101604081905261002f91610064565b6001600160a01b0392831660805260a0919091521660c0526100a7565b6001600160a01b038116811461006157600080fd5b50565b60008060006060848603121561007957600080fd5b83516100848161004c565b60208501516040860151919450925061009c8161004c565b809150509250925092565b60805160a05160c051610d846100f06000396000818160c201526103d201526000818161013201526103a601526000818160710152818161029c015261036a0152610d846000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063181f5a7711610050578063181f5a77146100e45780632ea023691461012d578063e83f967b1461016257600080fd5b806306c407201461006c5780630ab7dea9146100bd575b600080fd5b6100937f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100937f000000000000000000000000000000000000000000000000000000000000000081565b6101206040518060400160405280601481526020017f546f6c6c53656e6465724461707020312e302e3000000000000000000000000081525081565b6040516100b49190610916565b6101547f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100b4565b610175610170366004610a5b565b61018e565b60405167ffffffffffffffff90911681526020016100b4565b600073ffffffffffffffffffffffffffffffffffffffff84166101fa576040517ffdc6604f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024015b60405180910390fd5b60005b835181101561036757610265333085848151811061021d5761021d610b31565b602002602001015187858151811061023757610237610b31565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1661050b909392919063ffffffff16565b83818151811061027757610277610b31565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008584815181106102cd576102cd610b31565b60200260200101516040518363ffffffff1660e01b815260040161031392919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af1158015610332573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103569190610b60565b5061036081610b82565b90506101fd565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e7c62c8c7f00000000000000000000000000000000000000000000000000000000000000006040518060e001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168152602001338960405160200161044192919073ffffffffffffffffffffffffffffffffffffffff92831681529116602082015260400190565b60405160208183030381529060405281526020018781526020018681526020018760008151811061047457610474610b31565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152506040518363ffffffff1660e01b81526004016104c0929190610c1c565b6020604051808303816000875af11580156104df573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105039190610d31565b949350505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526105a09085906105a6565b50505050565b6000610608826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166106b79092919063ffffffff16565b8051909150156106b257808060200190518101906106269190610b60565b6106b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101f1565b505050565b606061050384846000856106cd565b9392505050565b60608247101561075f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101f1565b843b6107c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101f1565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516107f09190610d5b565b60006040518083038185875af1925050503d806000811461082d576040519150601f19603f3d011682016040523d82523d6000602084013e610832565b606091505b509150915061084282828661084d565b979650505050505050565b6060831561085c5750816106c6565b82511561086c5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f19190610916565b60005b838110156108bb5781810151838201526020016108a3565b838111156105a05750506000910152565b600081518084526108e48160208601602086016108a0565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006106c660208301846108cc565b73ffffffffffffffffffffffffffffffffffffffff8116811461094b57600080fd5b50565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156109c4576109c461094e565b604052919050565b600067ffffffffffffffff8211156109e6576109e661094e565b5060051b60200190565b600082601f830112610a0157600080fd5b81356020610a16610a11836109cc565b61097d565b82815260059290921b84018101918181019086841115610a3557600080fd5b8286015b84811015610a505780358352918301918301610a39565b509695505050505050565b600080600060608486031215610a7057600080fd5b8335610a7b81610929565b925060208481013567ffffffffffffffff80821115610a9957600080fd5b818701915087601f830112610aad57600080fd5b8135610abb610a11826109cc565b81815260059190911b8301840190848101908a831115610ada57600080fd5b938501935b82851015610b01578435610af281610929565b82529385019390850190610adf565b965050506040870135925080831115610b1957600080fd5b5050610b27868287016109f0565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215610b7257600080fd5b815180151581146106c657600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610bda577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b600081518084526020808501945080840160005b83811015610c1157815187529582019590820190600101610bf5565b509495945050505050565b8281526000602060408184015273ffffffffffffffffffffffffffffffffffffffff8085511660408501528185015160e06060860152610c606101208601826108cc565b60408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087830381016080890152815180845291860193506000929091908601905b80841015610cc657845186168252938601936001939093019290860190610ca4565b5060608901519550818882030160a0890152610ce28187610be1565b955050505050506080840151610d1060c085018273ffffffffffffffffffffffffffffffffffffffff169052565b5060a084015160e084015260c0840151610100840152809150509392505050565b600060208284031215610d4357600080fd5b815167ffffffffffffffff811681146106c657600080fd5b60008251610d6d8184602087016108a0565b919091019291505056fea164736f6c634300080f000a",
}

var TollSenderDappABI = TollSenderDappMetaData.ABI

var TollSenderDappBin = TollSenderDappMetaData.Bin

func DeployTollSenderDapp(auth *bind.TransactOpts, backend bind.ContractBackend, onRampRouter common.Address, destinationChainId *big.Int, destinationContract common.Address) (common.Address, *types.Transaction, *TollSenderDapp, error) {
	parsed, err := TollSenderDappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TollSenderDappBin), backend, onRampRouter, destinationChainId, destinationContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TollSenderDapp{TollSenderDappCaller: TollSenderDappCaller{contract: contract}, TollSenderDappTransactor: TollSenderDappTransactor{contract: contract}, TollSenderDappFilterer: TollSenderDappFilterer{contract: contract}}, nil
}

type TollSenderDapp struct {
	address common.Address
	abi     abi.ABI
	TollSenderDappCaller
	TollSenderDappTransactor
	TollSenderDappFilterer
}

type TollSenderDappCaller struct {
	contract *bind.BoundContract
}

type TollSenderDappTransactor struct {
	contract *bind.BoundContract
}

type TollSenderDappFilterer struct {
	contract *bind.BoundContract
}

type TollSenderDappSession struct {
	Contract     *TollSenderDapp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TollSenderDappCallerSession struct {
	Contract *TollSenderDappCaller
	CallOpts bind.CallOpts
}

type TollSenderDappTransactorSession struct {
	Contract     *TollSenderDappTransactor
	TransactOpts bind.TransactOpts
}

type TollSenderDappRaw struct {
	Contract *TollSenderDapp
}

type TollSenderDappCallerRaw struct {
	Contract *TollSenderDappCaller
}

type TollSenderDappTransactorRaw struct {
	Contract *TollSenderDappTransactor
}

func NewTollSenderDapp(address common.Address, backend bind.ContractBackend) (*TollSenderDapp, error) {
	abi, err := abi.JSON(strings.NewReader(TollSenderDappABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindTollSenderDapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TollSenderDapp{address: address, abi: abi, TollSenderDappCaller: TollSenderDappCaller{contract: contract}, TollSenderDappTransactor: TollSenderDappTransactor{contract: contract}, TollSenderDappFilterer: TollSenderDappFilterer{contract: contract}}, nil
}

func NewTollSenderDappCaller(address common.Address, caller bind.ContractCaller) (*TollSenderDappCaller, error) {
	contract, err := bindTollSenderDapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TollSenderDappCaller{contract: contract}, nil
}

func NewTollSenderDappTransactor(address common.Address, transactor bind.ContractTransactor) (*TollSenderDappTransactor, error) {
	contract, err := bindTollSenderDapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TollSenderDappTransactor{contract: contract}, nil
}

func NewTollSenderDappFilterer(address common.Address, filterer bind.ContractFilterer) (*TollSenderDappFilterer, error) {
	contract, err := bindTollSenderDapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TollSenderDappFilterer{contract: contract}, nil
}

func bindTollSenderDapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TollSenderDappABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_TollSenderDapp *TollSenderDappRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TollSenderDapp.Contract.TollSenderDappCaller.contract.Call(opts, result, method, params...)
}

func (_TollSenderDapp *TollSenderDappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.TollSenderDappTransactor.contract.Transfer(opts)
}

func (_TollSenderDapp *TollSenderDappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.TollSenderDappTransactor.contract.Transact(opts, method, params...)
}

func (_TollSenderDapp *TollSenderDappCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TollSenderDapp.Contract.contract.Call(opts, result, method, params...)
}

func (_TollSenderDapp *TollSenderDappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.contract.Transfer(opts)
}

func (_TollSenderDapp *TollSenderDappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.contract.Transact(opts, method, params...)
}

func (_TollSenderDapp *TollSenderDappCaller) DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TollSenderDapp.contract.Call(opts, &out, "DESTINATION_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TollSenderDapp *TollSenderDappSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _TollSenderDapp.Contract.DESTINATIONCHAINID(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCallerSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _TollSenderDapp.Contract.DESTINATIONCHAINID(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCaller) DESTINATIONCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TollSenderDapp.contract.Call(opts, &out, "DESTINATION_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TollSenderDapp *TollSenderDappSession) DESTINATIONCONTRACT() (common.Address, error) {
	return _TollSenderDapp.Contract.DESTINATIONCONTRACT(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCallerSession) DESTINATIONCONTRACT() (common.Address, error) {
	return _TollSenderDapp.Contract.DESTINATIONCONTRACT(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCaller) ONRAMPROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TollSenderDapp.contract.Call(opts, &out, "ON_RAMP_ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TollSenderDapp *TollSenderDappSession) ONRAMPROUTER() (common.Address, error) {
	return _TollSenderDapp.Contract.ONRAMPROUTER(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCallerSession) ONRAMPROUTER() (common.Address, error) {
	return _TollSenderDapp.Contract.ONRAMPROUTER(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TollSenderDapp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_TollSenderDapp *TollSenderDappSession) TypeAndVersion() (string, error) {
	return _TollSenderDapp.Contract.TypeAndVersion(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappCallerSession) TypeAndVersion() (string, error) {
	return _TollSenderDapp.Contract.TypeAndVersion(&_TollSenderDapp.CallOpts)
}

func (_TollSenderDapp *TollSenderDappTransactor) SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TollSenderDapp.contract.Transact(opts, "sendTokens", destinationAddress, tokens, amounts)
}

func (_TollSenderDapp *TollSenderDappSession) SendTokens(destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.SendTokens(&_TollSenderDapp.TransactOpts, destinationAddress, tokens, amounts)
}

func (_TollSenderDapp *TollSenderDappTransactorSession) SendTokens(destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TollSenderDapp.Contract.SendTokens(&_TollSenderDapp.TransactOpts, destinationAddress, tokens, amounts)
}

func (_TollSenderDapp *TollSenderDapp) Address() common.Address {
	return _TollSenderDapp.address
}

type TollSenderDappInterface interface {
	DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error)

	DESTINATIONCONTRACT(opts *bind.CallOpts) (common.Address, error)

	ONRAMPROUTER(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error)

	Address() common.Address
}
