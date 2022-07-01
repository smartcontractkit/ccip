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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractAny2EVMTollOnRampRouterInterface\",\"name\":\"onRampRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"invalidAddress\",\"type\":\"address\"}],\"name\":\"InvalidDestinationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DESTINATION_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DESTINATION_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ON_RAMP_ROUTER\",\"outputs\":[{\"internalType\":\"contractAny2EVMTollOnRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"sendTokens\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b50604051610e52380380610e5283398101604081905261002f91610064565b6001600160a01b0392831660805260a0919091521660c0526100a7565b6001600160a01b038116811461006157600080fd5b50565b60008060006060848603121561007957600080fd5b83516100848161004c565b60208501516040860151919450925061009c8161004c565b809150509250925092565b60805160a05160c051610d696100e9600039600060c201526000818161013201526103a701526000818160710152818161029d015261036b0152610d696000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063181f5a7711610050578063181f5a77146100e45780632ea023691461012d578063e83f967b1461016257600080fd5b806306c407201461006c5780630ab7dea9146100bd575b600080fd5b6100937f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100937f000000000000000000000000000000000000000000000000000000000000000081565b6101206040518060400160405280601081526020017f53656e6465724461707020312e302e300000000000000000000000000000000081525081565b6040516100b491906108fb565b6101547f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100b4565b610175610170366004610a40565b61018e565b60405167ffffffffffffffff90911681526020016100b4565b600073ffffffffffffffffffffffffffffffffffffffff84166101fa576040517ffdc6604f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024015b60405180910390fd5b3360005b845181101561036857610266823086848151811061021e5761021e610b16565b602002602001015188858151811061023857610238610b16565b602002602001015173ffffffffffffffffffffffffffffffffffffffff166104ed909392919063ffffffff16565b84818151811061027857610278610b16565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008684815181106102ce576102ce610b16565b60200260200101516040518363ffffffff1660e01b815260040161031492919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af1158015610333573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103579190610b45565b5061036181610b67565b90506101fe565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e7c62c8c7f00000000000000000000000000000000000000000000000000000000000000006040518060e001604052808973ffffffffffffffffffffffffffffffffffffffff168152602001858a60405160200161042292919073ffffffffffffffffffffffffffffffffffffffff92831681529116602082015260400190565b60405160208183030381529060405281526020018881526020018781526020018860008151811061045557610455610b16565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152506040518363ffffffff1660e01b81526004016104a1929190610c01565b6020604051808303816000875af11580156104c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e49190610d16565b95945050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052610582908590610588565b50505050565b60006105ea826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166106999092919063ffffffff16565b80519091501561069457808060200190518101906106089190610b45565b610694576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101f1565b505050565b60606106a884846000856106b2565b90505b9392505050565b606082471015610744576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101f1565b843b6107ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101f1565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516107d59190610d40565b60006040518083038185875af1925050503d8060008114610812576040519150601f19603f3d011682016040523d82523d6000602084013e610817565b606091505b5091509150610827828286610832565b979650505050505050565b606083156108415750816106ab565b8251156108515782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f191906108fb565b60005b838110156108a0578181015183820152602001610888565b838111156105825750506000910152565b600081518084526108c9816020860160208601610885565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006106ab60208301846108b1565b73ffffffffffffffffffffffffffffffffffffffff8116811461093057600080fd5b50565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156109a9576109a9610933565b604052919050565b600067ffffffffffffffff8211156109cb576109cb610933565b5060051b60200190565b600082601f8301126109e657600080fd5b813560206109fb6109f6836109b1565b610962565b82815260059290921b84018101918181019086841115610a1a57600080fd5b8286015b84811015610a355780358352918301918301610a1e565b509695505050505050565b600080600060608486031215610a5557600080fd5b8335610a608161090e565b925060208481013567ffffffffffffffff80821115610a7e57600080fd5b818701915087601f830112610a9257600080fd5b8135610aa06109f6826109b1565b81815260059190911b8301840190848101908a831115610abf57600080fd5b938501935b82851015610ae6578435610ad78161090e565b82529385019390850190610ac4565b965050506040870135925080831115610afe57600080fd5b5050610b0c868287016109d5565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215610b5757600080fd5b815180151581146106ab57600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610bbf577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b600081518084526020808501945080840160005b83811015610bf657815187529582019590820190600101610bda565b509495945050505050565b8281526000602060408184015273ffffffffffffffffffffffffffffffffffffffff8085511660408501528185015160e06060860152610c456101208601826108b1565b60408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087830381016080890152815180845291860193506000929091908601905b80841015610cab57845186168252938601936001939093019290860190610c89565b5060608901519550818882030160a0890152610cc78187610bc6565b955050505050506080840151610cf560c085018273ffffffffffffffffffffffffffffffffffffffff169052565b5060a084015160e084015260c0840151610100840152809150509392505050565b600060208284031215610d2857600080fd5b815167ffffffffffffffff811681146106ab57600080fd5b60008251610d52818460208701610885565b919091019291505056fea164736f6c634300080f000a",
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

func (_SenderDapp *SenderDappTransactor) SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SenderDapp.contract.Transact(opts, "sendTokens", destinationAddress, tokens, amounts)
}

func (_SenderDapp *SenderDappSession) SendTokens(destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SenderDapp.Contract.SendTokens(&_SenderDapp.TransactOpts, destinationAddress, tokens, amounts)
}

func (_SenderDapp *SenderDappTransactorSession) SendTokens(destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SenderDapp.Contract.SendTokens(&_SenderDapp.TransactOpts, destinationAddress, tokens, amounts)
}

func (_SenderDapp *SenderDapp) Address() common.Address {
	return _SenderDapp.address
}

type SenderDappInterface interface {
	DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error)

	DESTINATIONCONTRACT(opts *bind.CallOpts) (common.Address, error)

	ONRAMPROUTER(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	SendTokens(opts *bind.TransactOpts, destinationAddress common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error)

	Address() common.Address
}
