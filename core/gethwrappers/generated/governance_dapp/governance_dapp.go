// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package governance_dapp

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

type CCIPAny2EVMMessage struct {
	SourceChainId *big.Int
	Sender        []byte
	Data          []byte
	DestTokens    []common.Address
	Amounts       []*big.Int
}

type GovernanceDappCrossChainClone struct {
	ChainId         *big.Int
	ContractAddress common.Address
}

type GovernanceDappFeeConfig struct {
	FeeAmount           *big.Int
	SubscriptionManager common.Address
	ChangedAtBlock      *big.Int
}

var GovernanceDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"receivingRouter\",\"type\":\"address\"},{\"internalType\":\"contractEVM2AnySubscriptionOnRampRouterInterface\",\"name\":\"sendingRouter\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"subscriptionManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structGovernanceDapp.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deliverer\",\"type\":\"address\"}],\"name\":\"InvalidDeliverer\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ConfigPropagated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriptionManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"name\":\"ReceivedConfig\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structGovernanceDapp.CrossChainClone\",\"name\":\"clone\",\"type\":\"tuple\"}],\"name\":\"addClone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structCCIP.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"subscriptionManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structGovernanceDapp.FeeConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSubscriptionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"receivingRouter\",\"type\":\"address\"},{\"internalType\":\"contractEVM2AnySubscriptionOnRampRouterInterface\",\"name\":\"sendingRouter\",\"type\":\"address\"}],\"name\":\"setRouters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"subscriptionManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structGovernanceDapp.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"voteForNewFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001075380380620010758339810160408190526200003491620001da565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be8162000116565b5050600680546001600160a01b03199081166001600160a01b0396871617909155600780548216948616949094179093555080516002556020810151600380549093169316929092179055604001516004556200028b565b336001600160a01b03821603620001705760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b0381168114620001d757600080fd5b50565b600080600083850360a0811215620001f157600080fd5b8451620001fe81620001c1565b60208601519094506200021181620001c1565b92506060603f19820112156200022657600080fd5b50604051606081016001600160401b03811182821017156200025857634e487b7160e01b600052604160045260246000fd5b6040908152850151815260608501516200027281620001c1565b6020820152608094909401516040850152509093909250565b610dda806200029b6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c80638da5cb5b11610076578063e2a92e281161005b578063e2a92e28146101d9578063f2fde38b146101ea578063fa8c6d3a146101fd57600080fd5b80638da5cb5b146101a1578063a0c6df15146101c657600080fd5b80635ead6fb0116100a75780635ead6fb01461012a5780635fbbc0d21461016657806379ba50971461019957600080fd5b8063181f5a77146100c35780633aa5113c14610115575b600080fd5b6100ff6040518060400160405280601481526020017f476f7665726e616e63654461707020312e302e3000000000000000000000000081525081565b60405161010c91906107e9565b60405180910390f35b610128610123366004610888565b610210565b005b6101286101383660046108de565b600680546001600160a01b039384166001600160a01b03199182161790915560078054929093169116179055565b61016e61029d565b60408051825181526020808401516001600160a01b031690820152918101519082015260600161010c565b6101286102f8565b6000546001600160a01b03165b6040516001600160a01b03909116815260200161010c565b6101286101d4366004610a7a565b6103ae565b6003546001600160a01b03166101ae565b6101286101f8366004610b5d565b610487565b61012861020b366004610b7a565b61049b565b6102186104cf565b6005805460018101825560009190915281517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db06002909202918201556020909101517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db190910180546001600160a01b0319166001600160a01b03909216919091179055565b6102ca60405180606001604052806000815260200160006001600160a01b03168152602001600081525090565b506040805160608101825260025481526003546001600160a01b031660208201526004549181019190915290565b6001546001600160a01b031633146103575760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6006546001600160a01b031633146103f4576040517f0af9f1b600000000000000000000000000000000000000000000000000000000815233600482015260240161034e565b6000816040015180602001905181019061040e9190610b92565b80516002819055602080830151600380546001600160a01b0319166001600160a01b03909216918217905560408085015160048190558151948552928401919091528201529091507f012cabe5a5b0c3b1aca639f2e637eafd26c61e400ea107a8644bbe128ea94e189060600160405180910390a15050565b61048f6104cf565b6104988161052b565b50565b6104a36104cf565b6007546001600160a01b0316156104bd576104bd816105d4565b8060026104ca8282610bf4565b505050565b6000546001600160a01b031633146105295760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161034e565b565b336001600160a01b038216036105835760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161034e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000816040516020016105e79190610c33565b60408051601f1981840301815291905260055490915060005b818110156107965760006005828154811061061d5761061d610c69565b6000918252602080832060408051808201825260029490940290910180548452600101546001600160a01b03908116848401908152825160a0810184529051821660c0808301919091528351808303909101815260e08201845281528084018a90528251868152808501845281840152825195865292850182526060830194909452620493e06080830152600754835191517f3ccc152400000000000000000000000000000000000000000000000000000000815293955091939190911691633ccc1524916106f0918590600401610cba565b6020604051808303816000875af115801561070f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107339190610d7c565b50815160208301516040517f3d5ce3768ee5558c728489fb32f2e8accffd0a44616c2f27df0a7370e016fdb09261077b929082526001600160a01b0316602082015260400190565b60405180910390a150508061078f90610da6565b9050610600565b50505050565b6000815180845260005b818110156107c2576020818501810151868301820152016107a6565b818111156107d4576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006107fc602083018461079c565b9392505050565b634e487b7160e01b600052604160045260246000fd5b60405160a0810167ffffffffffffffff8111828210171561083c5761083c610803565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561086b5761086b610803565b604052919050565b6001600160a01b038116811461049857600080fd5b60006040828403121561089a57600080fd5b6040516040810181811067ffffffffffffffff821117156108bd576108bd610803565b6040528235815260208301356108d281610873565b60208201529392505050565b600080604083850312156108f157600080fd5b82356108fc81610873565b9150602083013561090c81610873565b809150509250929050565b600082601f83011261092857600080fd5b813567ffffffffffffffff81111561094257610942610803565b610955601f8201601f1916602001610842565b81815284602083860101111561096a57600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff8211156109a1576109a1610803565b5060051b60200190565b600082601f8301126109bc57600080fd5b813560206109d16109cc83610987565b610842565b82815260059290921b840181019181810190868411156109f057600080fd5b8286015b84811015610a14578035610a0781610873565b83529183019183016109f4565b509695505050505050565b600082601f830112610a3057600080fd5b81356020610a406109cc83610987565b82815260059290921b84018101918181019086841115610a5f57600080fd5b8286015b84811015610a145780358352918301918301610a63565b600060208284031215610a8c57600080fd5b813567ffffffffffffffff80821115610aa457600080fd5b9083019060a08286031215610ab857600080fd5b610ac0610819565b82358152602083013582811115610ad657600080fd5b610ae287828601610917565b602083015250604083013582811115610afa57600080fd5b610b0687828601610917565b604083015250606083013582811115610b1e57600080fd5b610b2a878286016109ab565b606083015250608083013582811115610b4257600080fd5b610b4e87828601610a1f565b60808301525095945050505050565b600060208284031215610b6f57600080fd5b81356107fc81610873565b600060608284031215610b8c57600080fd5b50919050565b600060608284031215610ba457600080fd5b6040516060810181811067ffffffffffffffff82111715610bc757610bc7610803565b604052825181526020830151610bdc81610873565b60208201526040928301519281019290925250919050565b81358155600181016020830135610c0a81610873565b6001600160a01b0381166001600160a01b03198354161782555050604082013560028201555050565b81358152606081016020830135610c4981610873565b6001600160a01b0381166020840152506040830135604083015292915050565b634e487b7160e01b600052603260045260246000fd5b600081518084526020808501945080840160005b83811015610caf57815187529582019590820190600101610c93565b509495945050505050565b82815260006020604081840152835160a06040850152610cdd60e085018261079c565b905081850151603f1980868403016060870152610cfa838361079c565b6040880151878203830160808901528051808352908601945060009350908501905b80841015610d455784516001600160a01b03168252938501936001939093019290850190610d1c565b5060608801519450818782030160a0880152610d618186610c7f565b945050505050608084015160c0840152809150509392505050565b600060208284031215610d8e57600080fd5b815167ffffffffffffffff811681146107fc57600080fd5b600060018201610dc657634e487b7160e01b600052601160045260246000fd5b506001019056fea164736f6c634300080f000a",
}

var GovernanceDappABI = GovernanceDappMetaData.ABI

var GovernanceDappBin = GovernanceDappMetaData.Bin

func DeployGovernanceDapp(auth *bind.TransactOpts, backend bind.ContractBackend, receivingRouter common.Address, sendingRouter common.Address, feeConfig GovernanceDappFeeConfig) (common.Address, *types.Transaction, *GovernanceDapp, error) {
	parsed, err := GovernanceDappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernanceDappBin), backend, receivingRouter, sendingRouter, feeConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovernanceDapp{GovernanceDappCaller: GovernanceDappCaller{contract: contract}, GovernanceDappTransactor: GovernanceDappTransactor{contract: contract}, GovernanceDappFilterer: GovernanceDappFilterer{contract: contract}}, nil
}

type GovernanceDapp struct {
	address common.Address
	abi     abi.ABI
	GovernanceDappCaller
	GovernanceDappTransactor
	GovernanceDappFilterer
}

type GovernanceDappCaller struct {
	contract *bind.BoundContract
}

type GovernanceDappTransactor struct {
	contract *bind.BoundContract
}

type GovernanceDappFilterer struct {
	contract *bind.BoundContract
}

type GovernanceDappSession struct {
	Contract     *GovernanceDapp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type GovernanceDappCallerSession struct {
	Contract *GovernanceDappCaller
	CallOpts bind.CallOpts
}

type GovernanceDappTransactorSession struct {
	Contract     *GovernanceDappTransactor
	TransactOpts bind.TransactOpts
}

type GovernanceDappRaw struct {
	Contract *GovernanceDapp
}

type GovernanceDappCallerRaw struct {
	Contract *GovernanceDappCaller
}

type GovernanceDappTransactorRaw struct {
	Contract *GovernanceDappTransactor
}

func NewGovernanceDapp(address common.Address, backend bind.ContractBackend) (*GovernanceDapp, error) {
	abi, err := abi.JSON(strings.NewReader(GovernanceDappABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindGovernanceDapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernanceDapp{address: address, abi: abi, GovernanceDappCaller: GovernanceDappCaller{contract: contract}, GovernanceDappTransactor: GovernanceDappTransactor{contract: contract}, GovernanceDappFilterer: GovernanceDappFilterer{contract: contract}}, nil
}

func NewGovernanceDappCaller(address common.Address, caller bind.ContractCaller) (*GovernanceDappCaller, error) {
	contract, err := bindGovernanceDapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceDappCaller{contract: contract}, nil
}

func NewGovernanceDappTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceDappTransactor, error) {
	contract, err := bindGovernanceDapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceDappTransactor{contract: contract}, nil
}

func NewGovernanceDappFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceDappFilterer, error) {
	contract, err := bindGovernanceDapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceDappFilterer{contract: contract}, nil
}

func bindGovernanceDapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceDappABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_GovernanceDapp *GovernanceDappRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceDapp.Contract.GovernanceDappCaller.contract.Call(opts, result, method, params...)
}

func (_GovernanceDapp *GovernanceDappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.GovernanceDappTransactor.contract.Transfer(opts)
}

func (_GovernanceDapp *GovernanceDappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.GovernanceDappTransactor.contract.Transact(opts, method, params...)
}

func (_GovernanceDapp *GovernanceDappCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceDapp.Contract.contract.Call(opts, result, method, params...)
}

func (_GovernanceDapp *GovernanceDappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.contract.Transfer(opts)
}

func (_GovernanceDapp *GovernanceDappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.contract.Transact(opts, method, params...)
}

func (_GovernanceDapp *GovernanceDappCaller) GetFeeConfig(opts *bind.CallOpts) (GovernanceDappFeeConfig, error) {
	var out []interface{}
	err := _GovernanceDapp.contract.Call(opts, &out, "getFeeConfig")

	if err != nil {
		return *new(GovernanceDappFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(GovernanceDappFeeConfig)).(*GovernanceDappFeeConfig)

	return out0, err

}

func (_GovernanceDapp *GovernanceDappSession) GetFeeConfig() (GovernanceDappFeeConfig, error) {
	return _GovernanceDapp.Contract.GetFeeConfig(&_GovernanceDapp.CallOpts)
}

func (_GovernanceDapp *GovernanceDappCallerSession) GetFeeConfig() (GovernanceDappFeeConfig, error) {
	return _GovernanceDapp.Contract.GetFeeConfig(&_GovernanceDapp.CallOpts)
}

func (_GovernanceDapp *GovernanceDappCaller) GetSubscriptionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GovernanceDapp.contract.Call(opts, &out, "getSubscriptionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_GovernanceDapp *GovernanceDappSession) GetSubscriptionManager() (common.Address, error) {
	return _GovernanceDapp.Contract.GetSubscriptionManager(&_GovernanceDapp.CallOpts)
}

func (_GovernanceDapp *GovernanceDappCallerSession) GetSubscriptionManager() (common.Address, error) {
	return _GovernanceDapp.Contract.GetSubscriptionManager(&_GovernanceDapp.CallOpts)
}

func (_GovernanceDapp *GovernanceDappCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GovernanceDapp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_GovernanceDapp *GovernanceDappSession) Owner() (common.Address, error) {
	return _GovernanceDapp.Contract.Owner(&_GovernanceDapp.CallOpts)
}

func (_GovernanceDapp *GovernanceDappCallerSession) Owner() (common.Address, error) {
	return _GovernanceDapp.Contract.Owner(&_GovernanceDapp.CallOpts)
}

func (_GovernanceDapp *GovernanceDappCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GovernanceDapp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_GovernanceDapp *GovernanceDappSession) TypeAndVersion() (string, error) {
	return _GovernanceDapp.Contract.TypeAndVersion(&_GovernanceDapp.CallOpts)
}

func (_GovernanceDapp *GovernanceDappCallerSession) TypeAndVersion() (string, error) {
	return _GovernanceDapp.Contract.TypeAndVersion(&_GovernanceDapp.CallOpts)
}

func (_GovernanceDapp *GovernanceDappTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceDapp.contract.Transact(opts, "acceptOwnership")
}

func (_GovernanceDapp *GovernanceDappSession) AcceptOwnership() (*types.Transaction, error) {
	return _GovernanceDapp.Contract.AcceptOwnership(&_GovernanceDapp.TransactOpts)
}

func (_GovernanceDapp *GovernanceDappTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _GovernanceDapp.Contract.AcceptOwnership(&_GovernanceDapp.TransactOpts)
}

func (_GovernanceDapp *GovernanceDappTransactor) AddClone(opts *bind.TransactOpts, clone GovernanceDappCrossChainClone) (*types.Transaction, error) {
	return _GovernanceDapp.contract.Transact(opts, "addClone", clone)
}

func (_GovernanceDapp *GovernanceDappSession) AddClone(clone GovernanceDappCrossChainClone) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.AddClone(&_GovernanceDapp.TransactOpts, clone)
}

func (_GovernanceDapp *GovernanceDappTransactorSession) AddClone(clone GovernanceDappCrossChainClone) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.AddClone(&_GovernanceDapp.TransactOpts, clone)
}

func (_GovernanceDapp *GovernanceDappTransactor) CcipReceive(opts *bind.TransactOpts, message CCIPAny2EVMMessage) (*types.Transaction, error) {
	return _GovernanceDapp.contract.Transact(opts, "ccipReceive", message)
}

func (_GovernanceDapp *GovernanceDappSession) CcipReceive(message CCIPAny2EVMMessage) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.CcipReceive(&_GovernanceDapp.TransactOpts, message)
}

func (_GovernanceDapp *GovernanceDappTransactorSession) CcipReceive(message CCIPAny2EVMMessage) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.CcipReceive(&_GovernanceDapp.TransactOpts, message)
}

func (_GovernanceDapp *GovernanceDappTransactor) SetRouters(opts *bind.TransactOpts, receivingRouter common.Address, sendingRouter common.Address) (*types.Transaction, error) {
	return _GovernanceDapp.contract.Transact(opts, "setRouters", receivingRouter, sendingRouter)
}

func (_GovernanceDapp *GovernanceDappSession) SetRouters(receivingRouter common.Address, sendingRouter common.Address) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.SetRouters(&_GovernanceDapp.TransactOpts, receivingRouter, sendingRouter)
}

func (_GovernanceDapp *GovernanceDappTransactorSession) SetRouters(receivingRouter common.Address, sendingRouter common.Address) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.SetRouters(&_GovernanceDapp.TransactOpts, receivingRouter, sendingRouter)
}

func (_GovernanceDapp *GovernanceDappTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _GovernanceDapp.contract.Transact(opts, "transferOwnership", to)
}

func (_GovernanceDapp *GovernanceDappSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.TransferOwnership(&_GovernanceDapp.TransactOpts, to)
}

func (_GovernanceDapp *GovernanceDappTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.TransferOwnership(&_GovernanceDapp.TransactOpts, to)
}

func (_GovernanceDapp *GovernanceDappTransactor) VoteForNewFeeConfig(opts *bind.TransactOpts, feeConfig GovernanceDappFeeConfig) (*types.Transaction, error) {
	return _GovernanceDapp.contract.Transact(opts, "voteForNewFeeConfig", feeConfig)
}

func (_GovernanceDapp *GovernanceDappSession) VoteForNewFeeConfig(feeConfig GovernanceDappFeeConfig) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.VoteForNewFeeConfig(&_GovernanceDapp.TransactOpts, feeConfig)
}

func (_GovernanceDapp *GovernanceDappTransactorSession) VoteForNewFeeConfig(feeConfig GovernanceDappFeeConfig) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.VoteForNewFeeConfig(&_GovernanceDapp.TransactOpts, feeConfig)
}

type GovernanceDappConfigPropagatedIterator struct {
	Event *GovernanceDappConfigPropagated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GovernanceDappConfigPropagatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceDappConfigPropagated)
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
		it.Event = new(GovernanceDappConfigPropagated)
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

func (it *GovernanceDappConfigPropagatedIterator) Error() error {
	return it.fail
}

func (it *GovernanceDappConfigPropagatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GovernanceDappConfigPropagated struct {
	ChainId         *big.Int
	ContractAddress common.Address
	Raw             types.Log
}

func (_GovernanceDapp *GovernanceDappFilterer) FilterConfigPropagated(opts *bind.FilterOpts) (*GovernanceDappConfigPropagatedIterator, error) {

	logs, sub, err := _GovernanceDapp.contract.FilterLogs(opts, "ConfigPropagated")
	if err != nil {
		return nil, err
	}
	return &GovernanceDappConfigPropagatedIterator{contract: _GovernanceDapp.contract, event: "ConfigPropagated", logs: logs, sub: sub}, nil
}

func (_GovernanceDapp *GovernanceDappFilterer) WatchConfigPropagated(opts *bind.WatchOpts, sink chan<- *GovernanceDappConfigPropagated) (event.Subscription, error) {

	logs, sub, err := _GovernanceDapp.contract.WatchLogs(opts, "ConfigPropagated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GovernanceDappConfigPropagated)
				if err := _GovernanceDapp.contract.UnpackLog(event, "ConfigPropagated", log); err != nil {
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

func (_GovernanceDapp *GovernanceDappFilterer) ParseConfigPropagated(log types.Log) (*GovernanceDappConfigPropagated, error) {
	event := new(GovernanceDappConfigPropagated)
	if err := _GovernanceDapp.contract.UnpackLog(event, "ConfigPropagated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GovernanceDappOwnershipTransferRequestedIterator struct {
	Event *GovernanceDappOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GovernanceDappOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceDappOwnershipTransferRequested)
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
		it.Event = new(GovernanceDappOwnershipTransferRequested)
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

func (it *GovernanceDappOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *GovernanceDappOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GovernanceDappOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_GovernanceDapp *GovernanceDappFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GovernanceDappOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GovernanceDapp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceDappOwnershipTransferRequestedIterator{contract: _GovernanceDapp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_GovernanceDapp *GovernanceDappFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *GovernanceDappOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GovernanceDapp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GovernanceDappOwnershipTransferRequested)
				if err := _GovernanceDapp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_GovernanceDapp *GovernanceDappFilterer) ParseOwnershipTransferRequested(log types.Log) (*GovernanceDappOwnershipTransferRequested, error) {
	event := new(GovernanceDappOwnershipTransferRequested)
	if err := _GovernanceDapp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GovernanceDappOwnershipTransferredIterator struct {
	Event *GovernanceDappOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GovernanceDappOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceDappOwnershipTransferred)
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
		it.Event = new(GovernanceDappOwnershipTransferred)
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

func (it *GovernanceDappOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *GovernanceDappOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GovernanceDappOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_GovernanceDapp *GovernanceDappFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GovernanceDappOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GovernanceDapp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceDappOwnershipTransferredIterator{contract: _GovernanceDapp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_GovernanceDapp *GovernanceDappFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GovernanceDappOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GovernanceDapp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GovernanceDappOwnershipTransferred)
				if err := _GovernanceDapp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_GovernanceDapp *GovernanceDappFilterer) ParseOwnershipTransferred(log types.Log) (*GovernanceDappOwnershipTransferred, error) {
	event := new(GovernanceDappOwnershipTransferred)
	if err := _GovernanceDapp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GovernanceDappReceivedConfigIterator struct {
	Event *GovernanceDappReceivedConfig

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GovernanceDappReceivedConfigIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceDappReceivedConfig)
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
		it.Event = new(GovernanceDappReceivedConfig)
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

func (it *GovernanceDappReceivedConfigIterator) Error() error {
	return it.fail
}

func (it *GovernanceDappReceivedConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GovernanceDappReceivedConfig struct {
	FeeAmount           *big.Int
	SubscriptionManager common.Address
	ChangedAtBlock      *big.Int
	Raw                 types.Log
}

func (_GovernanceDapp *GovernanceDappFilterer) FilterReceivedConfig(opts *bind.FilterOpts) (*GovernanceDappReceivedConfigIterator, error) {

	logs, sub, err := _GovernanceDapp.contract.FilterLogs(opts, "ReceivedConfig")
	if err != nil {
		return nil, err
	}
	return &GovernanceDappReceivedConfigIterator{contract: _GovernanceDapp.contract, event: "ReceivedConfig", logs: logs, sub: sub}, nil
}

func (_GovernanceDapp *GovernanceDappFilterer) WatchReceivedConfig(opts *bind.WatchOpts, sink chan<- *GovernanceDappReceivedConfig) (event.Subscription, error) {

	logs, sub, err := _GovernanceDapp.contract.WatchLogs(opts, "ReceivedConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GovernanceDappReceivedConfig)
				if err := _GovernanceDapp.contract.UnpackLog(event, "ReceivedConfig", log); err != nil {
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

func (_GovernanceDapp *GovernanceDappFilterer) ParseReceivedConfig(log types.Log) (*GovernanceDappReceivedConfig, error) {
	event := new(GovernanceDappReceivedConfig)
	if err := _GovernanceDapp.contract.UnpackLog(event, "ReceivedConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_GovernanceDapp *GovernanceDapp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _GovernanceDapp.abi.Events["ConfigPropagated"].ID:
		return _GovernanceDapp.ParseConfigPropagated(log)
	case _GovernanceDapp.abi.Events["OwnershipTransferRequested"].ID:
		return _GovernanceDapp.ParseOwnershipTransferRequested(log)
	case _GovernanceDapp.abi.Events["OwnershipTransferred"].ID:
		return _GovernanceDapp.ParseOwnershipTransferred(log)
	case _GovernanceDapp.abi.Events["ReceivedConfig"].ID:
		return _GovernanceDapp.ParseReceivedConfig(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (GovernanceDappConfigPropagated) Topic() common.Hash {
	return common.HexToHash("0x3d5ce3768ee5558c728489fb32f2e8accffd0a44616c2f27df0a7370e016fdb0")
}

func (GovernanceDappOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (GovernanceDappOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (GovernanceDappReceivedConfig) Topic() common.Hash {
	return common.HexToHash("0x012cabe5a5b0c3b1aca639f2e637eafd26c61e400ea107a8644bbe128ea94e18")
}

func (_GovernanceDapp *GovernanceDapp) Address() common.Address {
	return _GovernanceDapp.address
}

type GovernanceDappInterface interface {
	GetFeeConfig(opts *bind.CallOpts) (GovernanceDappFeeConfig, error)

	GetSubscriptionManager(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddClone(opts *bind.TransactOpts, clone GovernanceDappCrossChainClone) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message CCIPAny2EVMMessage) (*types.Transaction, error)

	SetRouters(opts *bind.TransactOpts, receivingRouter common.Address, sendingRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	VoteForNewFeeConfig(opts *bind.TransactOpts, feeConfig GovernanceDappFeeConfig) (*types.Transaction, error)

	FilterConfigPropagated(opts *bind.FilterOpts) (*GovernanceDappConfigPropagatedIterator, error)

	WatchConfigPropagated(opts *bind.WatchOpts, sink chan<- *GovernanceDappConfigPropagated) (event.Subscription, error)

	ParseConfigPropagated(log types.Log) (*GovernanceDappConfigPropagated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GovernanceDappOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *GovernanceDappOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*GovernanceDappOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GovernanceDappOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GovernanceDappOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*GovernanceDappOwnershipTransferred, error)

	FilterReceivedConfig(opts *bind.FilterOpts) (*GovernanceDappReceivedConfigIterator, error)

	WatchReceivedConfig(opts *bind.WatchOpts, sink chan<- *GovernanceDappReceivedConfig) (event.Subscription, error)

	ParseReceivedConfig(log types.Log) (*GovernanceDappReceivedConfig, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
