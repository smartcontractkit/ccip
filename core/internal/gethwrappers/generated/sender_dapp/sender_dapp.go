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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractOnRamp\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DESTINATION_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DESTINATION_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ON_RAMP\",\"outputs\":[{\"internalType\":\"contractOnRamp\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"sendTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b50604051610e90380380610e9083398101604081905261002f91610064565b6001600160a01b0392831660805260a0919091521660c0526100a7565b6001600160a01b038116811461006157600080fd5b50565b60008060006060848603121561007957600080fd5b83516100848161004c565b60208501516040860151919450925061009c8161004c565b809150509250925092565b60805160a05160c051610d9f6100f1600039600081816071015261021f01526000818161014901526101f90152600081816101220152818161036601526104710152610d9f6000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063181f5a7711610050578063181f5a77146100de5780632c3b6d231461011d5780632ea023691461014457600080fd5b80630ab7dea91461006c578063177eeec5146100bd575b600080fd5b6100937f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100d06100cb3660046109cf565b61016b565b6040519081526020016100b4565b604080518082018252601081527f53656e6465724461707020312e302e3000000000000000000000000000000000602082015290516100b49190610b2c565b6100937f000000000000000000000000000000000000000000000000000000000000000081565b6100d07f000000000000000000000000000000000000000000000000000000000000000081565b600073ffffffffffffffffffffffffffffffffffffffff85166101d7576040517ffdc6604f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024015b60405180910390fd5b6060600033905060006040518060e001604052808881526020018781526020017f000000000000000000000000000000000000000000000000000000000000000081526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001838a6040516020016102aa92919073ffffffffffffffffffffffffffffffffffffffff92831681529116602082015260400190565b604051602081830303815290604052815260200184815250905060005b87518110156104335761032f83308984815181106102e7576102e7610b3f565b60200260200101518b858151811061030157610301610b3f565b602002602001015173ffffffffffffffffffffffffffffffffffffffff166104f5909392919063ffffffff16565b87818151811061034157610341610b3f565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f000000000000000000000000000000000000000000000000000000000000000089848151811061039757610397610b3f565b60200260200101516040518363ffffffff1660e01b81526004016103dd92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af11580156103fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104209190610b6e565b508061042b81610b90565b9150506102c7565b506040517fc8658c1c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063c8658c1c906104a6908490600401610c2b565b6020604051808303816000875af11580156104c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e99190610d5d565b98975050505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261058a908590610590565b50505050565b60006105f2826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166106a19092919063ffffffff16565b80519091501561069c57808060200190518101906106109190610b6e565b61069c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101ce565b505050565b60606106b084846000856106ba565b90505b9392505050565b60608247101561074c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101ce565b843b6107b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101ce565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516107dd9190610d76565b60006040518083038185875af1925050503d806000811461081a576040519150601f19603f3d011682016040523d82523d6000602084013e61081f565b606091505b509150915061082f82828661083a565b979650505050505050565b606083156108495750816106b3565b8251156108595782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ce9190610b2c565b73ffffffffffffffffffffffffffffffffffffffff811681146108af57600080fd5b50565b80356108bd8161088d565b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610938576109386108c2565b604052919050565b600067ffffffffffffffff82111561095a5761095a6108c2565b5060051b60200190565b600082601f83011261097557600080fd5b8135602061098a61098583610940565b6108f1565b82815260059290921b840181019181810190868411156109a957600080fd5b8286015b848110156109c457803583529183019183016109ad565b509695505050505050565b600080600080608085870312156109e557600080fd5b84356109f08161088d565b935060208581013567ffffffffffffffff80821115610a0e57600080fd5b818801915088601f830112610a2257600080fd5b8135610a3061098582610940565b81815260059190911b8301840190848101908b831115610a4f57600080fd5b938501935b82851015610a76578435610a678161088d565b82529385019390850190610a54565b975050506040880135925080831115610a8e57600080fd5b5050610a9c87828801610964565b925050610aab606086016108b2565b905092959194509250565b60005b83811015610ad1578181015183820152602001610ab9565b8381111561058a5750506000910152565b60008151808452610afa816020860160208601610ab6565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006106b36020830184610ae2565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215610b8057600080fd5b815180151581146106b357600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610be9577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b600081518084526020808501945080840160005b83811015610c2057815187529582019590820190600101610c04565b509495945050505050565b6020808252825160e083830152805161010084018190526000929182019083906101208601905b80831015610c8857835173ffffffffffffffffffffffffffffffffffffffff168252928401926001929092019190840190610c52565b508387015193507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0925082868203016040870152610cc68185610bf0565b93505050604085015160608501526060850151610cfb608086018273ffffffffffffffffffffffffffffffffffffffff169052565b50608085015173ffffffffffffffffffffffffffffffffffffffff811660a08601525060a0850151818584030160c0860152610d378382610ae2565b92505060c0850151818584030160e0860152610d538382610ae2565b9695505050505050565b600060208284031215610d6f57600080fd5b5051919050565b60008251610d88818460208701610ab6565b919091019291505056fea164736f6c634300080c000a",
}

var SenderDappABI = SenderDappMetaData.ABI

var SenderDappBin = SenderDappMetaData.Bin

func DeploySenderDapp(auth *bind.TransactOpts, backend bind.ContractBackend, onRamp common.Address, destinationChainId *big.Int, destinationContract common.Address) (common.Address, *types.Transaction, *SenderDapp, error) {
	parsed, err := SenderDappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SenderDappBin), backend, onRamp, destinationChainId, destinationContract)
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

func (_SenderDapp *SenderDappCaller) ONRAMP(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SenderDapp.contract.Call(opts, &out, "ON_RAMP")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SenderDapp *SenderDappSession) ONRAMP() (common.Address, error) {
	return _SenderDapp.Contract.ONRAMP(&_SenderDapp.CallOpts)
}

func (_SenderDapp *SenderDappCallerSession) ONRAMP() (common.Address, error) {
	return _SenderDapp.Contract.ONRAMP(&_SenderDapp.CallOpts)
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

	ONRAMP(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int, executor common.Address) (*types.Transaction, error)

	Address() common.Address
}
