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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractTollOnRampRouterInterface\",\"name\":\"onRampRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidAddress\",\"type\":\"address\"}],\"name\":\"InvalidDestinationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DESTINATION_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DESTINATION_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ON_RAMP_ROUTER\",\"outputs\":[{\"internalType\":\"contractTollOnRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"sendTokens\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b50604051610e6c380380610e6c83398101604081905261002f91610064565b6001600160a01b0392831660805260a0919091521660c0526100a7565b6001600160a01b038116811461006157600080fd5b50565b60008060006060848603121561007957600080fd5b83516100848161004c565b60208501516040860151919450925061009c8161004c565b809150509250925092565b60805160a05160c051610d836100e9600039600060c2015260008181610154015261039f01526000818160710152818161029301526103630152610d836000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063177eeec511610050578063177eeec5146100e4578063181f5a77146101105780632ea023691461014f57600080fd5b806306c407201461006c5780630ab7dea9146100bd575b600080fd5b6100937f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100937f000000000000000000000000000000000000000000000000000000000000000081565b6100f76100f23660046109c0565b610184565b60405167ffffffffffffffff90911681526020016100b4565b604080518082018252601081527f53656e6465724461707020312e302e3000000000000000000000000000000000602082015290516100b49190610b1d565b6101767f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100b4565b600073ffffffffffffffffffffffffffffffffffffffff85166101f0576040517ffdc6604f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024015b60405180910390fd5b3360005b85518110156103605761025c823087848151811061021457610214610b30565b602002602001015189858151811061022e5761022e610b30565b602002602001015173ffffffffffffffffffffffffffffffffffffffff166104e6909392919063ffffffff16565b85818151811061026e5761026e610b30565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008784815181106102c4576102c4610b30565b60200260200101516040518363ffffffff1660e01b815260040161030a92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af1158015610329573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061034d9190610b5f565b508061035881610b81565b9150506101f4565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e7c62c8c7f00000000000000000000000000000000000000000000000000000000000000006040518060e001604052808a73ffffffffffffffffffffffffffffffffffffffff168152602001858b60405160200161041a92919073ffffffffffffffffffffffffffffffffffffffff92831681529116602082015260400190565b60405160208183030381529060405281526020018981526020018881526020018960008151811061044d5761044d610b30565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152506040518363ffffffff1660e01b8152600401610499929190610c1b565b6020604051808303816000875af11580156104b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104dc9190610d30565b9695505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261057b908590610581565b50505050565b60006105e3826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166106929092919063ffffffff16565b80519091501561068d57808060200190518101906106019190610b5f565b61068d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101e7565b505050565b60606106a184846000856106ab565b90505b9392505050565b60608247101561073d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101e7565b843b6107a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101e7565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516107ce9190610d5a565b60006040518083038185875af1925050503d806000811461080b576040519150601f19603f3d011682016040523d82523d6000602084013e610810565b606091505b509150915061082082828661082b565b979650505050505050565b6060831561083a5750816106a4565b82511561084a5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e79190610b1d565b73ffffffffffffffffffffffffffffffffffffffff811681146108a057600080fd5b50565b80356108ae8161087e565b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610929576109296108b3565b604052919050565b600067ffffffffffffffff82111561094b5761094b6108b3565b5060051b60200190565b600082601f83011261096657600080fd5b8135602061097b61097683610931565b6108e2565b82815260059290921b8401810191818101908684111561099a57600080fd5b8286015b848110156109b5578035835291830191830161099e565b509695505050505050565b600080600080608085870312156109d657600080fd5b84356109e18161087e565b935060208581013567ffffffffffffffff808211156109ff57600080fd5b818801915088601f830112610a1357600080fd5b8135610a2161097682610931565b81815260059190911b8301840190848101908b831115610a4057600080fd5b938501935b82851015610a67578435610a588161087e565b82529385019390850190610a45565b975050506040880135925080831115610a7f57600080fd5b5050610a8d87828801610955565b925050610a9c606086016108a3565b905092959194509250565b60005b83811015610ac2578181015183820152602001610aaa565b8381111561057b5750506000910152565b60008151808452610aeb816020860160208601610aa7565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006106a46020830184610ad3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215610b7157600080fd5b815180151581146106a457600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610bd9577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b600081518084526020808501945080840160005b83811015610c1057815187529582019590820190600101610bf4565b509495945050505050565b8281526000602060408184015273ffffffffffffffffffffffffffffffffffffffff8085511660408501528185015160e06060860152610c5f610120860182610ad3565b60408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087830381016080890152815180845291860193506000929091908601905b80841015610cc557845186168252938601936001939093019290860190610ca3565b5060608901519550818882030160a0890152610ce18187610be0565b955050505050506080840151610d0f60c085018273ffffffffffffffffffffffffffffffffffffffff169052565b5060a084015160e084015260c0840151610100840152809150509392505050565b600060208284031215610d4257600080fd5b815167ffffffffffffffff811681146106a457600080fd5b60008251610d6c818460208701610aa7565b919091019291505056fea164736f6c634300080d000a",
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
