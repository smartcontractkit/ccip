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
	Bin: "0x60806040523480156200001157600080fd5b5060405162001371380380620013718339810160408190526200003491620001da565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be8162000116565b5050600680546001600160a01b03199081166001600160a01b0396871617909155600780548216948616949094179093555080516002556020810151600380549093169316929092179055604001516004556200028b565b336001600160a01b03821603620001705760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b0381168114620001d757600080fd5b50565b600080600083850360a0811215620001f157600080fd5b8451620001fe81620001c1565b60208601519094506200021181620001c1565b92506060603f19820112156200022657600080fd5b50604051606081016001600160401b03811182821017156200025857634e487b7160e01b600052604160045260246000fd5b6040908152850151815260608501516200027281620001c1565b6020820152608094909401516040850152509093909250565b6110d6806200029b6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c80638da5cb5b11610076578063e2a92e281161005b578063e2a92e2814610225578063f2fde38b14610243578063fa8c6d3a1461025657600080fd5b80638da5cb5b146101d3578063a0c6df151461021257600080fd5b80635ead6fb0116100a75780635ead6fb01461012a5780635fbbc0d21461018b57806379ba5097146101cb57600080fd5b8063181f5a77146100c35780633aa5113c14610115575b600080fd5b6100ff6040518060400160405280601481526020017f476f7665726e616e63654461707020312e302e3000000000000000000000000081525081565b60405161010c91906109d5565b60405180910390f35b610128610123366004610ab8565b610269565b005b610128610138366004610b0e565b6006805473ffffffffffffffffffffffffffffffffffffffff9384167fffffffffffffffffffffffff00000000000000000000000000000000000000009182161790915560078054929093169116179055565b61019361031b565b604080518251815260208084015173ffffffffffffffffffffffffffffffffffffffff1690820152918101519082015260600161010c565b610128610390565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010c565b610128610220366004610cc8565b610492565b60035473ffffffffffffffffffffffffffffffffffffffff166101ed565b610128610251366004610dab565b61059d565b610128610264366004610dc8565b6105b1565b6102716105f2565b6005805460018101825560009190915281517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db06002909202918201556020909101517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db190910180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055565b610355604051806060016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b5060408051606081018252600254815260035473ffffffffffffffffffffffffffffffffffffffff1660208201526004549181019190915290565b60015473ffffffffffffffffffffffffffffffffffffffff163314610416576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60065473ffffffffffffffffffffffffffffffffffffffff1633146104e5576040517f0af9f1b600000000000000000000000000000000000000000000000000000000815233600482015260240161040d565b600081604001518060200190518101906104ff9190610de0565b80516002819055602080830151600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216918217905560408085015160048190558151948552928401919091528201529091507f012cabe5a5b0c3b1aca639f2e637eafd26c61e400ea107a8644bbe128ea94e189060600160405180910390a15050565b6105a56105f2565b6105ae81610675565b50565b6105b96105f2565b60075473ffffffffffffffffffffffffffffffffffffffff16156105e0576105e08161076a565b8060026105ed8282610e42565b505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610673576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161040d565b565b3373ffffffffffffffffffffffffffffffffffffffff8216036106f4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161040d565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008160405160200161077d9190610ea6565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905260055490915060005b81811015610964576000600582815481106107d1576107d1610ee9565b60009182526020808320604080518082018252600294909402909101805484526001015473ffffffffffffffffffffffffffffffffffffffff908116848401908152825160a0810184529051821660c0808301919091528351808303909101815260e08201845281528084018a90528251868152808501845281840152825195865292850182526060830194909452620493e06080830152600754835191517f3ccc152400000000000000000000000000000000000000000000000000000000815293955091939190911691633ccc1524916108b1918590600401610f53565b6020604051808303816000875af11580156108d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f49190611040565b50815160208301516040517f3d5ce3768ee5558c728489fb32f2e8accffd0a44616c2f27df0a7370e016fdb0926109499290825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b60405180910390a150508061095d9061106a565b90506107b4565b50505050565b6000815180845260005b8181101561099057602081850181015186830182015201610974565b818111156109a2576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006109e8602083018461096a565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff81118282101715610a4157610a416109ef565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610a8e57610a8e6109ef565b604052919050565b73ffffffffffffffffffffffffffffffffffffffff811681146105ae57600080fd5b600060408284031215610aca57600080fd5b6040516040810181811067ffffffffffffffff82111715610aed57610aed6109ef565b604052823581526020830135610b0281610a96565b60208201529392505050565b60008060408385031215610b2157600080fd5b8235610b2c81610a96565b91506020830135610b3c81610a96565b809150509250929050565b600082601f830112610b5857600080fd5b813567ffffffffffffffff811115610b7257610b726109ef565b610ba360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610a47565b818152846020838601011115610bb857600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115610bef57610bef6109ef565b5060051b60200190565b600082601f830112610c0a57600080fd5b81356020610c1f610c1a83610bd5565b610a47565b82815260059290921b84018101918181019086841115610c3e57600080fd5b8286015b84811015610c62578035610c5581610a96565b8352918301918301610c42565b509695505050505050565b600082601f830112610c7e57600080fd5b81356020610c8e610c1a83610bd5565b82815260059290921b84018101918181019086841115610cad57600080fd5b8286015b84811015610c625780358352918301918301610cb1565b600060208284031215610cda57600080fd5b813567ffffffffffffffff80821115610cf257600080fd5b9083019060a08286031215610d0657600080fd5b610d0e610a1e565b82358152602083013582811115610d2457600080fd5b610d3087828601610b47565b602083015250604083013582811115610d4857600080fd5b610d5487828601610b47565b604083015250606083013582811115610d6c57600080fd5b610d7887828601610bf9565b606083015250608083013582811115610d9057600080fd5b610d9c87828601610c6d565b60808301525095945050505050565b600060208284031215610dbd57600080fd5b81356109e881610a96565b600060608284031215610dda57600080fd5b50919050565b600060608284031215610df257600080fd5b6040516060810181811067ffffffffffffffff82111715610e1557610e156109ef565b604052825181526020830151610e2a81610a96565b60208201526040928301519281019290925250919050565b81358155600181016020830135610e5881610a96565b73ffffffffffffffffffffffffffffffffffffffff81167fffffffffffffffffffffffff00000000000000000000000000000000000000008354161782555050604082013560028201555050565b81358152606081016020830135610ebc81610a96565b73ffffffffffffffffffffffffffffffffffffffff81166020840152506040830135604083015292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081518084526020808501945080840160005b83811015610f4857815187529582019590820190600101610f2c565b509495945050505050565b82815260006020604081840152835160a06040850152610f7660e085018261096a565b9050818501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080868403016060870152610fb1838361096a565b6040880151878203830160808901528051808352908601945060009350908501905b8084101561100957845173ffffffffffffffffffffffffffffffffffffffff168252938501936001939093019290850190610fd3565b5060608801519450818782030160a08801526110258186610f18565b945050505050608084015160c0840152809150509392505050565b60006020828403121561105257600080fd5b815167ffffffffffffffff811681146109e857600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110c2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c634300080f000a",
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
