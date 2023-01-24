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

type CommonAny2EVMMessage struct {
	SourceChainId        uint64
	Sender               []byte
	Data                 []byte
	DestTokensAndAmounts []CommonEVMTokenAndAmount
}

type CommonEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

type GovernanceDappCrossChainClone struct {
	ChainId         uint64
	ContractAddress common.Address
}

type GovernanceDappFeeConfig struct {
	FeeAmount      *big.Int
	ChangedAtBlock *big.Int
}

var GovernanceDappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structGovernanceDapp.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deliverer\",\"type\":\"address\"}],\"name\":\"InvalidDeliverer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ConfigPropagated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"FeeTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"name\":\"ReceivedConfig\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structGovernanceDapp.CrossChainClone\",\"name\":\"clone\",\"type\":\"tuple\"}],\"name\":\"addClone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCommon.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structGovernanceDapp.FeeConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structGovernanceDapp.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"voteForNewFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200157238038062001572833981016040819052620000349162000237565b3380600085846001600160a01b0382166200006a576040516335fdcccd60e21b8152600060048201526024015b60405180910390fd5b6001600160a01b038216608052620000828162000126565b50506001600160a01b038216620000dc5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f0000000000000000604482015260640162000061565b600180546001600160a01b0319166001600160a01b03848116919091179091558116156200010f576200010f816200016e565b5050825160035550506020015160045550620002ce565b600080546001600160a01b0319166001600160a01b038316908117825560405190917f722ff84c1234b2482061def5c82c6b5080c117b3cbb69d686844a051e4b8e7f391a250565b336001600160a01b03821603620001c85760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000061565b600280546001600160a01b0319166001600160a01b03838116918217909255600154604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b80516001600160a01b03811681146200023257600080fd5b919050565b600080600083850360808112156200024e57600080fd5b62000259856200021a565b93506040601f19820112156200026e57600080fd5b50604080519081016001600160401b0381118282101715620002a057634e487b7160e01b600052604160045260246000fd5b6040908152602086810151835290860151908201529150620002c5606085016200021a565b90509250925092565b60805161126c620003066000396000818161020c015281816103c80152818161055a015281816106740152610bdc015261126c6000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c80638157c96c11610081578063ca1d209d1161005b578063ca1d209d14610230578063ca709a2514610243578063f2fde38b1461026157600080fd5b80638157c96c146101b85780638da5cb5b146101cb578063b0f479a11461020a57600080fd5b80633015b91c116100b25780633015b91c1461015f5780635fbbc0d21461017257806379ba5097146101b057600080fd5b806301ffc9a7146100d9578063181f5a7714610101578063262a7b631461014a575b600080fd5b6100ec6100e7366004610c5d565b610274565b60405190151581526020015b60405180910390f35b61013d6040518060400160405280601481526020017f476f7665726e616e63654461707020312e302e3000000000000000000000000081525081565b6040516100f89190610cec565b61015d610158366004610dec565b61030d565b005b61015d61016d366004610e29565b6103b0565b6040805180820182526000808252602091820152815180830183526003548082526004549183019182528351908152905191810191909152016100f8565b61015d61043a565b61015d6101c6366004610e64565b61053b565b60015473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f8565b7f00000000000000000000000000000000000000000000000000000000000000006101e5565b61015d61023e366004610e7c565b61059b565b60005473ffffffffffffffffffffffffffffffffffffffff166101e5565b61015d61026f366004610e95565b610729565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f3015b91c00000000000000000000000000000000000000000000000000000000148061030757507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b61031561073a565b6005805460018101825560009190915281517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0909101805460209093015173ffffffffffffffffffffffffffffffffffffffff1668010000000000000000027fffffffff0000000000000000000000000000000000000000000000000000000090931667ffffffffffffffff90921691909117919091179055565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610426576040517fd7f733340000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61043761043282610f20565b6107bd565b50565b60025473ffffffffffffffffffffffffffffffffffffffff1633146104bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161041d565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560028054909116905560405173ffffffffffffffffffffffffffffffffffffffff909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b61054361073a565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016156105885761058881610826565b803560035560200135600455565b505050565b6000805473ffffffffffffffffffffffffffffffffffffffff166040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810184905290915073ffffffffffffffffffffffffffffffffffffffff8216906323b872dd906064016020604051808303816000875af1158015610631573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106559190611052565b5073ffffffffffffffffffffffffffffffffffffffff811663095ea7b37f00000000000000000000000000000000000000000000000000000000000000006040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602481018590526044016020604051808303816000875af1158015610705573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105969190611052565b61073161073a565b61043781610aa6565b60015473ffffffffffffffffffffffffffffffffffffffff1633146107bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161041d565b565b600081604001518060200190518101906107d79190611074565b80516003819055602080830151600481905560408051938452918301529192507f583bd9682201bddaa3ee0ed61c39b397de860af4d62cfc5a62ecd30ca7342deb910160405180910390a15050565b60008160405160200161084b9190813581526020918201359181019190915260400190565b60408051601f1981840301815291905260055490915060005b81811015610aa057600060058281548110610881576108816110a6565b60009182526020808320604080518082018252939091015467ffffffffffffffff8116845273ffffffffffffffffffffffffffffffffffffffff680100000000000000009091048116848401908152825160a081018452905190911660c0808301919091528251808303909101815260e0820183528152808301899052815185815292830182529294508201908361093b565b60408051808201909152600080825260208201528152602001906001900390816109145790505b50815260200161096060005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff168152602001610a116040518060400160405280620493e08152602001600015158152506040805182516024820152602092830151151560448083019190915282518083039091018152606490910190915290810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f97a657c90000000000000000000000000000000000000000000000000000000017905290565b90528251909150610a229082610b9c565b50815160208301516040517f5c49f05c94a889457073501ff9101fe42cbe2255c2ee91c5e4008d09cf4cb2a492610a8592909167ffffffffffffffff92909216825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b60405180910390a1505080610a99906110d5565b9050610864565b50505050565b3373ffffffffffffffffffffffffffffffffffffffff821603610b25576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161041d565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600154604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b6040517f96f4e9f900000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906396f4e9f990610c139086908690600401611134565b6020604051808303816000875af1158015610c32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c569190611246565b9392505050565b600060208284031215610c6f57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610c5657600080fd5b6000815180845260005b81811015610cc557602081850181015186830182015201610ca9565b81811115610cd7576000602083870101525b50601f01601f19169290920160200192915050565b602081526000610c566020830184610c9f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610d5157610d51610cff565b60405290565b6040516080810167ffffffffffffffff81118282101715610d5157610d51610cff565b604051601f8201601f1916810167ffffffffffffffff81118282101715610da357610da3610cff565b604052919050565b803567ffffffffffffffff81168114610dc357600080fd5b919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610dc357600080fd5b600060408284031215610dfe57600080fd5b610e06610d2e565b610e0f83610dab565b8152610e1d60208401610dc8565b60208201529392505050565b600060208284031215610e3b57600080fd5b813567ffffffffffffffff811115610e5257600080fd5b820160808185031215610c5657600080fd5b600060408284031215610e7657600080fd5b50919050565b600060208284031215610e8e57600080fd5b5035919050565b600060208284031215610ea757600080fd5b610c5682610dc8565b600082601f830112610ec157600080fd5b813567ffffffffffffffff811115610edb57610edb610cff565b610eee6020601f19601f84011601610d7a565b818152846020838601011115610f0357600080fd5b816020850160208301376000918101602001919091529392505050565b600060808236031215610f3257600080fd5b610f3a610d57565b610f4383610dab565b815260208084013567ffffffffffffffff80821115610f6157600080fd5b610f6d36838801610eb0565b83850152604091508186013581811115610f8657600080fd5b610f9236828901610eb0565b8386015250606086013581811115610fa957600080fd5b860136601f820112610fba57600080fd5b803582811115610fcc57610fcc610cff565b610fda858260051b01610d7a565b818152858101935060069190911b820185019036821115610ffa57600080fd5b918501915b81831015611041578483360312156110175760008081fd5b61101f610d2e565b61102884610dc8565b8152838701358782015284529285019291840191610fff565b606087015250939695505050505050565b60006020828403121561106457600080fd5b81518015158114610c5657600080fd5b60006040828403121561108657600080fd5b61108e610d2e565b82518152602083015160208201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361112d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b6000604067ffffffffffffffff8516835260208181850152845160a08386015261116160e0860182610c9f565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08087840301606088015261119c8383610c9f565b88860151888203830160808a01528051808352908601945060009350908501905b808410156111fc578451805173ffffffffffffffffffffffffffffffffffffffff168352860151868301529385019360019390930192908601906111bd565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a015295506112388187610c9f565b9a9950505050505050505050565b60006020828403121561125857600080fd5b505191905056fea164736f6c634300080f000a",
}

var GovernanceDappABI = GovernanceDappMetaData.ABI

var GovernanceDappBin = GovernanceDappMetaData.Bin

func DeployGovernanceDapp(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address, feeConfig GovernanceDappFeeConfig, feeToken common.Address) (common.Address, *types.Transaction, *GovernanceDapp, error) {
	parsed, err := GovernanceDappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernanceDappBin), backend, router, feeConfig, feeToken)
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

func (_GovernanceDapp *GovernanceDappCaller) GetFeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GovernanceDapp.contract.Call(opts, &out, "getFeeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_GovernanceDapp *GovernanceDappSession) GetFeeToken() (common.Address, error) {
	return _GovernanceDapp.Contract.GetFeeToken(&_GovernanceDapp.CallOpts)
}

func (_GovernanceDapp *GovernanceDappCallerSession) GetFeeToken() (common.Address, error) {
	return _GovernanceDapp.Contract.GetFeeToken(&_GovernanceDapp.CallOpts)
}

func (_GovernanceDapp *GovernanceDappCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GovernanceDapp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_GovernanceDapp *GovernanceDappSession) GetRouter() (common.Address, error) {
	return _GovernanceDapp.Contract.GetRouter(&_GovernanceDapp.CallOpts)
}

func (_GovernanceDapp *GovernanceDappCallerSession) GetRouter() (common.Address, error) {
	return _GovernanceDapp.Contract.GetRouter(&_GovernanceDapp.CallOpts)
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

func (_GovernanceDapp *GovernanceDappCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GovernanceDapp.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_GovernanceDapp *GovernanceDappSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GovernanceDapp.Contract.SupportsInterface(&_GovernanceDapp.CallOpts, interfaceId)
}

func (_GovernanceDapp *GovernanceDappCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GovernanceDapp.Contract.SupportsInterface(&_GovernanceDapp.CallOpts, interfaceId)
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

func (_GovernanceDapp *GovernanceDappTransactor) CcipReceive(opts *bind.TransactOpts, message CommonAny2EVMMessage) (*types.Transaction, error) {
	return _GovernanceDapp.contract.Transact(opts, "ccipReceive", message)
}

func (_GovernanceDapp *GovernanceDappSession) CcipReceive(message CommonAny2EVMMessage) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.CcipReceive(&_GovernanceDapp.TransactOpts, message)
}

func (_GovernanceDapp *GovernanceDappTransactorSession) CcipReceive(message CommonAny2EVMMessage) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.CcipReceive(&_GovernanceDapp.TransactOpts, message)
}

func (_GovernanceDapp *GovernanceDappTransactor) Fund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceDapp.contract.Transact(opts, "fund", amount)
}

func (_GovernanceDapp *GovernanceDappSession) Fund(amount *big.Int) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.Fund(&_GovernanceDapp.TransactOpts, amount)
}

func (_GovernanceDapp *GovernanceDappTransactorSession) Fund(amount *big.Int) (*types.Transaction, error) {
	return _GovernanceDapp.Contract.Fund(&_GovernanceDapp.TransactOpts, amount)
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
	ChainId         uint64
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

type GovernanceDappFeeTokenSetIterator struct {
	Event *GovernanceDappFeeTokenSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GovernanceDappFeeTokenSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceDappFeeTokenSet)
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
		it.Event = new(GovernanceDappFeeTokenSet)
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

func (it *GovernanceDappFeeTokenSetIterator) Error() error {
	return it.fail
}

func (it *GovernanceDappFeeTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GovernanceDappFeeTokenSet struct {
	FeeToken common.Address
	Raw      types.Log
}

func (_GovernanceDapp *GovernanceDappFilterer) FilterFeeTokenSet(opts *bind.FilterOpts, feeToken []common.Address) (*GovernanceDappFeeTokenSetIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _GovernanceDapp.contract.FilterLogs(opts, "FeeTokenSet", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceDappFeeTokenSetIterator{contract: _GovernanceDapp.contract, event: "FeeTokenSet", logs: logs, sub: sub}, nil
}

func (_GovernanceDapp *GovernanceDappFilterer) WatchFeeTokenSet(opts *bind.WatchOpts, sink chan<- *GovernanceDappFeeTokenSet, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _GovernanceDapp.contract.WatchLogs(opts, "FeeTokenSet", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GovernanceDappFeeTokenSet)
				if err := _GovernanceDapp.contract.UnpackLog(event, "FeeTokenSet", log); err != nil {
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

func (_GovernanceDapp *GovernanceDappFilterer) ParseFeeTokenSet(log types.Log) (*GovernanceDappFeeTokenSet, error) {
	event := new(GovernanceDappFeeTokenSet)
	if err := _GovernanceDapp.contract.UnpackLog(event, "FeeTokenSet", log); err != nil {
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
	FeeAmount      *big.Int
	ChangedAtBlock *big.Int
	Raw            types.Log
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
	case _GovernanceDapp.abi.Events["FeeTokenSet"].ID:
		return _GovernanceDapp.ParseFeeTokenSet(log)
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
	return common.HexToHash("0x5c49f05c94a889457073501ff9101fe42cbe2255c2ee91c5e4008d09cf4cb2a4")
}

func (GovernanceDappFeeTokenSet) Topic() common.Hash {
	return common.HexToHash("0x722ff84c1234b2482061def5c82c6b5080c117b3cbb69d686844a051e4b8e7f3")
}

func (GovernanceDappOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (GovernanceDappOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (GovernanceDappReceivedConfig) Topic() common.Hash {
	return common.HexToHash("0x583bd9682201bddaa3ee0ed61c39b397de860af4d62cfc5a62ecd30ca7342deb")
}

func (_GovernanceDapp *GovernanceDapp) Address() common.Address {
	return _GovernanceDapp.address
}

type GovernanceDappInterface interface {
	GetFeeConfig(opts *bind.CallOpts) (GovernanceDappFeeConfig, error)

	GetFeeToken(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddClone(opts *bind.TransactOpts, clone GovernanceDappCrossChainClone) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message CommonAny2EVMMessage) (*types.Transaction, error)

	Fund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	VoteForNewFeeConfig(opts *bind.TransactOpts, feeConfig GovernanceDappFeeConfig) (*types.Transaction, error)

	FilterConfigPropagated(opts *bind.FilterOpts) (*GovernanceDappConfigPropagatedIterator, error)

	WatchConfigPropagated(opts *bind.WatchOpts, sink chan<- *GovernanceDappConfigPropagated) (event.Subscription, error)

	ParseConfigPropagated(log types.Log) (*GovernanceDappConfigPropagated, error)

	FilterFeeTokenSet(opts *bind.FilterOpts, feeToken []common.Address) (*GovernanceDappFeeTokenSetIterator, error)

	WatchFeeTokenSet(opts *bind.WatchOpts, sink chan<- *GovernanceDappFeeTokenSet, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenSet(log types.Log) (*GovernanceDappFeeTokenSet, error)

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
