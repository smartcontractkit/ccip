// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package subscription_sender_dapp

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

var SubscriptionSenderDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractAny2EVMSubscriptionOnRampRouterInterface\",\"name\":\"onRampRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidAddress\",\"type\":\"address\"}],\"name\":\"InvalidDestinationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DESTINATION_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DESTINATION_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ON_RAMP_ROUTER\",\"outputs\":[{\"internalType\":\"contractAny2EVMSubscriptionOnRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"sendTokens\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b50604051610e03380380610e0383398101604081905261002f91610064565b6001600160a01b0392831660805260a0919091521660c0526100a7565b6001600160a01b038116811461006157600080fd5b50565b60008060006060848603121561007957600080fd5b83516100848161004c565b60208501516040860151919450925061009c8161004c565b809150509250925092565b60805160a05160c051610d136100f06000396000818160c201526103d201526000818161013201526103a601526000818160710152818161029c015261036a0152610d136000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063181f5a7711610050578063181f5a77146100e45780632ea023691461012d578063e83f967b1461016257600080fd5b806306c407201461006c5780630ab7dea9146100bd575b600080fd5b6100937f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100937f000000000000000000000000000000000000000000000000000000000000000081565b6101206040518060400160405280601c81526020017f537562736372697074696f6e53656e6465724461707020312e302e300000000081525081565b6040516100b491906108d9565b6101547f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100b4565b610175610170366004610a1e565b61018e565b60405167ffffffffffffffff90911681526020016100b4565b600073ffffffffffffffffffffffffffffffffffffffff84166101fa576040517ffdc6604f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024015b60405180910390fd5b60005b835181101561036757610265333085848151811061021d5761021d610af4565b602002602001015187858151811061023757610237610af4565b602002602001015173ffffffffffffffffffffffffffffffffffffffff166104ce909392919063ffffffff16565b83818151811061027757610277610af4565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008584815181106102cd576102cd610af4565b60200260200101516040518363ffffffff1660e01b815260040161031392919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af1158015610332573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103569190610b23565b5061036081610b45565b90506101fd565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16630d58bf0c7f00000000000000000000000000000000000000000000000000000000000000006040518060a001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168152602001338960405160200161044192919073ffffffffffffffffffffffffffffffffffffffff92831681529116602082015260400190565b604051602081830303815290604052815260200187815260200186815260200160008152506040518363ffffffff1660e01b8152600401610483929190610bdf565b6020604051808303816000875af11580156104a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104c69190610cc0565b949350505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052610563908590610569565b50505050565b60006105cb826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661067a9092919063ffffffff16565b80519091501561067557808060200190518101906105e99190610b23565b610675576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101f1565b505050565b60606104c68484600085610690565b9392505050565b606082471015610722576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101f1565b843b61078a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101f1565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516107b39190610cea565b60006040518083038185875af1925050503d80600081146107f0576040519150601f19603f3d011682016040523d82523d6000602084013e6107f5565b606091505b5091509150610805828286610810565b979650505050505050565b6060831561081f575081610689565b82511561082f5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f191906108d9565b60005b8381101561087e578181015183820152602001610866565b838111156105635750506000910152565b600081518084526108a7816020860160208601610863565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610689602083018461088f565b73ffffffffffffffffffffffffffffffffffffffff8116811461090e57600080fd5b50565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561098757610987610911565b604052919050565b600067ffffffffffffffff8211156109a9576109a9610911565b5060051b60200190565b600082601f8301126109c457600080fd5b813560206109d96109d48361098f565b610940565b82815260059290921b840181019181810190868411156109f857600080fd5b8286015b84811015610a1357803583529183019183016109fc565b509695505050505050565b600080600060608486031215610a3357600080fd5b8335610a3e816108ec565b925060208481013567ffffffffffffffff80821115610a5c57600080fd5b818701915087601f830112610a7057600080fd5b8135610a7e6109d48261098f565b81815260059190911b8301840190848101908a831115610a9d57600080fd5b938501935b82851015610ac4578435610ab5816108ec565b82529385019390850190610aa2565b965050506040870135925080831115610adc57600080fd5b5050610aea868287016109b3565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215610b3557600080fd5b8151801515811461068957600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610b9d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b600081518084526020808501945080840160005b83811015610bd457815187529582019590820190600101610bb8565b509495945050505050565b8281526000602060408184015273ffffffffffffffffffffffffffffffffffffffff8085511660408501528185015160a06060860152610c2260e086018261088f565b60408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087830381016080890152815180845291860193506000929091908601905b80841015610c8857845186168252938601936001939093019290860190610c66565b5060608901519550818882030160a0890152610ca48187610ba4565b95505050505050608084015160c0840152809150509392505050565b600060208284031215610cd257600080fd5b815167ffffffffffffffff8116811461068957600080fd5b60008251610cfc818460208701610863565b919091019291505056fea164736f6c634300080f000a",
}

var SubscriptionSenderDappABI = SubscriptionSenderDappMetaData.ABI

var SubscriptionSenderDappBin = SubscriptionSenderDappMetaData.Bin

func DeploySubscriptionSenderDapp(auth *bind.TransactOpts, backend bind.ContractBackend, onRampRouter common.Address, destinationChainId *big.Int, destinationContract common.Address) (common.Address, *types.Transaction, *SubscriptionSenderDapp, error) {
	parsed, err := SubscriptionSenderDappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SubscriptionSenderDappBin), backend, onRampRouter, destinationChainId, destinationContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SubscriptionSenderDapp{SubscriptionSenderDappCaller: SubscriptionSenderDappCaller{contract: contract}, SubscriptionSenderDappTransactor: SubscriptionSenderDappTransactor{contract: contract}, SubscriptionSenderDappFilterer: SubscriptionSenderDappFilterer{contract: contract}}, nil
}

type SubscriptionSenderDapp struct {
	address common.Address
	abi     abi.ABI
	SubscriptionSenderDappCaller
	SubscriptionSenderDappTransactor
	SubscriptionSenderDappFilterer
}

type SubscriptionSenderDappCaller struct {
	contract *bind.BoundContract
}

type SubscriptionSenderDappTransactor struct {
	contract *bind.BoundContract
}

type SubscriptionSenderDappFilterer struct {
	contract *bind.BoundContract
}

type SubscriptionSenderDappSession struct {
	Contract     *SubscriptionSenderDapp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type SubscriptionSenderDappCallerSession struct {
	Contract *SubscriptionSenderDappCaller
	CallOpts bind.CallOpts
}

type SubscriptionSenderDappTransactorSession struct {
	Contract     *SubscriptionSenderDappTransactor
	TransactOpts bind.TransactOpts
}

type SubscriptionSenderDappRaw struct {
	Contract *SubscriptionSenderDapp
}

type SubscriptionSenderDappCallerRaw struct {
	Contract *SubscriptionSenderDappCaller
}

type SubscriptionSenderDappTransactorRaw struct {
	Contract *SubscriptionSenderDappTransactor
}

func NewSubscriptionSenderDapp(address common.Address, backend bind.ContractBackend) (*SubscriptionSenderDapp, error) {
	abi, err := abi.JSON(strings.NewReader(SubscriptionSenderDappABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindSubscriptionSenderDapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubscriptionSenderDapp{address: address, abi: abi, SubscriptionSenderDappCaller: SubscriptionSenderDappCaller{contract: contract}, SubscriptionSenderDappTransactor: SubscriptionSenderDappTransactor{contract: contract}, SubscriptionSenderDappFilterer: SubscriptionSenderDappFilterer{contract: contract}}, nil
}

func NewSubscriptionSenderDappCaller(address common.Address, caller bind.ContractCaller) (*SubscriptionSenderDappCaller, error) {
	contract, err := bindSubscriptionSenderDapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionSenderDappCaller{contract: contract}, nil
}

func NewSubscriptionSenderDappTransactor(address common.Address, transactor bind.ContractTransactor) (*SubscriptionSenderDappTransactor, error) {
	contract, err := bindSubscriptionSenderDapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionSenderDappTransactor{contract: contract}, nil
}

func NewSubscriptionSenderDappFilterer(address common.Address, filterer bind.ContractFilterer) (*SubscriptionSenderDappFilterer, error) {
	contract, err := bindSubscriptionSenderDapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubscriptionSenderDappFilterer{contract: contract}, nil
}

func bindSubscriptionSenderDapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubscriptionSenderDappABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubscriptionSenderDapp.Contract.SubscriptionSenderDappCaller.contract.Call(opts, result, method, params...)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.SubscriptionSenderDappTransactor.contract.Transfer(opts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.SubscriptionSenderDappTransactor.contract.Transact(opts, method, params...)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubscriptionSenderDapp.Contract.contract.Call(opts, result, method, params...)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.contract.Transfer(opts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.contract.Transact(opts, method, params...)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCaller) DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubscriptionSenderDapp.contract.Call(opts, &out, "DESTINATION_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SubscriptionSenderDapp *SubscriptionSenderDappSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _SubscriptionSenderDapp.Contract.DESTINATIONCHAINID(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCallerSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _SubscriptionSenderDapp.Contract.DESTINATIONCHAINID(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCaller) DESTINATIONCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SubscriptionSenderDapp.contract.Call(opts, &out, "DESTINATION_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SubscriptionSenderDapp *SubscriptionSenderDappSession) DESTINATIONCONTRACT() (common.Address, error) {
	return _SubscriptionSenderDapp.Contract.DESTINATIONCONTRACT(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCallerSession) DESTINATIONCONTRACT() (common.Address, error) {
	return _SubscriptionSenderDapp.Contract.DESTINATIONCONTRACT(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCaller) ONRAMPROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SubscriptionSenderDapp.contract.Call(opts, &out, "ON_RAMP_ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SubscriptionSenderDapp *SubscriptionSenderDappSession) ONRAMPROUTER() (common.Address, error) {
	return _SubscriptionSenderDapp.Contract.ONRAMPROUTER(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCallerSession) ONRAMPROUTER() (common.Address, error) {
	return _SubscriptionSenderDapp.Contract.ONRAMPROUTER(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SubscriptionSenderDapp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_SubscriptionSenderDapp *SubscriptionSenderDappSession) TypeAndVersion() (string, error) {
	return _SubscriptionSenderDapp.Contract.TypeAndVersion(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappCallerSession) TypeAndVersion() (string, error) {
	return _SubscriptionSenderDapp.Contract.TypeAndVersion(&_SubscriptionSenderDapp.CallOpts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappTransactor) SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.contract.Transact(opts, "sendTokens", destinationAddress, tokens, amounts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappSession) SendTokens(destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.SendTokens(&_SubscriptionSenderDapp.TransactOpts, destinationAddress, tokens, amounts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDappTransactorSession) SendTokens(destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SubscriptionSenderDapp.Contract.SendTokens(&_SubscriptionSenderDapp.TransactOpts, destinationAddress, tokens, amounts)
}

func (_SubscriptionSenderDapp *SubscriptionSenderDapp) Address() common.Address {
	return _SubscriptionSenderDapp.address
}

type SubscriptionSenderDappInterface interface {
	DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error)

	DESTINATIONCONTRACT(opts *bind.CallOpts) (common.Address, error)

	ONRAMPROUTER(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error)

	Address() common.Address
}
