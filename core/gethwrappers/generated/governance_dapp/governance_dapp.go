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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structGovernanceDapp.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deliverer\",\"type\":\"address\"}],\"name\":\"InvalidDeliverer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ConfigPropagated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"FeeTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"name\":\"ReceivedConfig\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structGovernanceDapp.CrossChainClone\",\"name\":\"clone\",\"type\":\"tuple\"}],\"name\":\"addClone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCommon.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structGovernanceDapp.FeeConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structGovernanceDapp.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"voteForNewFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200137f3803806200137f83398101604081905262000034916200020b565b6001600160a01b0383166080523380600085846200005281620000fa565b50506001600160a01b038216620000b05760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600180546001600160a01b0319166001600160a01b0384811691909117909155811615620000e357620000e38162000142565b5050825160035550506020015160045550620002a2565b600080546001600160a01b0319166001600160a01b038316908117825560405190917f722ff84c1234b2482061def5c82c6b5080c117b3cbb69d686844a051e4b8e7f391a250565b336001600160a01b038216036200019c5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a7565b600280546001600160a01b0319166001600160a01b03838116918217909255600154604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b80516001600160a01b03811681146200020657600080fd5b919050565b600080600083850360808112156200022257600080fd5b6200022d85620001ee565b93506040601f19820112156200024257600080fd5b50604080519081016001600160401b03811182821017156200027457634e487b7160e01b600052604160045260246000fd5b60409081526020868101518352908601519082015291506200029960608501620001ee565b90509250925092565b60805161109e620002e1600039600081816101c8015281816102c40152818161042f01528181610522015281816109e60152610a77015261109e6000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80638157c96c11610081578063ca1d209d1161005b578063ca1d209d146101ec578063ca709a25146101ff578063f2fde38b1461021057600080fd5b80638157c96c1461018e5780638da5cb5b146101a1578063b0f479a1146101c657600080fd5b80633015b91c116100b25780633015b91c146101355780635fbbc0d21461014857806379ba50971461018657600080fd5b8063181f5a77146100ce578063262a7b6314610120575b600080fd5b61010a6040518060400160405280601481526020017f476f7665726e616e63654461707020312e302e3000000000000000000000000081525081565b6040516101179190610b45565b60405180910390f35b61013361012e366004610c38565b610223565b005b610133610143366004610c75565b6102b9565b604080518082018252600080825260209182015281518083018352600354808252600454918301918252835190815290519181019190915201610117565b610133610336565b61013361019c366004610cb0565b61041d565b6001546001600160a01b03165b6040516001600160a01b039091168152602001610117565b7f00000000000000000000000000000000000000000000000000000000000000006101ae565b6101336101fa366004610cc8565b610470565b6000546001600160a01b03166101ae565b61013361021e366004610ce1565b6105ca565b61022b6105db565b6005805460018101825560009190915281517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db090910180546020909301516001600160a01b031668010000000000000000027fffffffff0000000000000000000000000000000000000000000000000000000090931667ffffffffffffffff90921691909117919091179055565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610322576040517fd7f733340000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61033361032e82610d6c565b610651565b50565b6002546001600160a01b031633146103aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610319565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000808216339081179093556002805490911690556040516001600160a01b03909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b6104256105db565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161561045d5761045d816106ba565b803560035560200135600455565b505050565b600080546001600160a01b03166040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490529091506001600160a01b038216906323b872dd906064016020604051808303816000875af11580156104ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105109190610e9e565b506001600160a01b03811663095ea7b37f00000000000000000000000000000000000000000000000000000000000000006040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b039091166004820152602481018590526044016020604051808303816000875af11580156105a6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061046b9190610e9e565b6105d26105db565b61033381610906565b6001546001600160a01b0316331461064f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610319565b565b6000816040015180602001905181019061066b9190610ec0565b80516003819055602080830151600481905560408051938452918301529192507f583bd9682201bddaa3ee0ed61c39b397de860af4d62cfc5a62ecd30ca7342deb910160405180910390a15050565b6000816040516020016106df9190813581526020918201359181019190915260400190565b60408051601f1981840301815291905260055490915060005b818110156109005760006005828154811061071557610715610ef2565b60009182526020808320604080518082018252939091015467ffffffffffffffff811684526001600160a01b03680100000000000000009091048116848401908152825160a081018452905190911660c0808301919091528251808303909101815260e082018352815280830189905281518581529283018252929450820190836107c2565b604080518082019091526000808252602082015281526020019060019003908161079b5790505b5081526020016107da6000546001600160a01b031690565b6001600160a01b0316815260200161087e6040518060400160405280620493e08152602001600015158152506040805182516024820152602092830151151560448083019190915282518083039091018152606490910190915290810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f97a657c90000000000000000000000000000000000000000000000000000000017905290565b9052825190915061088f90826109e2565b50815160208301516040517f5c49f05c94a889457073501ff9101fe42cbe2255c2ee91c5e4008d09cf4cb2a4926108e592909167ffffffffffffffff9290921682526001600160a01b0316602082015260400190565b60405180910390a15050806108f990610f21565b90506106f8565b50505050565b336001600160a01b03821603610978576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610319565b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03838116918217909255600154604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610a47576040517fd7f7333400000000000000000000000000000000000000000000000000000000815260006004820152602401610319565b6040517f96f4e9f90000000000000000000000000000000000000000000000000000000081526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906396f4e9f990610aae9086908690600401610f80565b6020604051808303816000875af1158015610acd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610af19190611078565b9392505050565b6000815180845260005b81811015610b1e57602081850181015186830182015201610b02565b81811115610b30576000602083870101525b50601f01601f19169290920160200192915050565b602081526000610af16020830184610af8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610baa57610baa610b58565b60405290565b6040516080810167ffffffffffffffff81118282101715610baa57610baa610b58565b604051601f8201601f1916810167ffffffffffffffff81118282101715610bfc57610bfc610b58565b604052919050565b803567ffffffffffffffff81168114610c1c57600080fd5b919050565b80356001600160a01b0381168114610c1c57600080fd5b600060408284031215610c4a57600080fd5b610c52610b87565b610c5b83610c04565b8152610c6960208401610c21565b60208201529392505050565b600060208284031215610c8757600080fd5b813567ffffffffffffffff811115610c9e57600080fd5b820160808185031215610af157600080fd5b600060408284031215610cc257600080fd5b50919050565b600060208284031215610cda57600080fd5b5035919050565b600060208284031215610cf357600080fd5b610af182610c21565b600082601f830112610d0d57600080fd5b813567ffffffffffffffff811115610d2757610d27610b58565b610d3a6020601f19601f84011601610bd3565b818152846020838601011115610d4f57600080fd5b816020850160208301376000918101602001919091529392505050565b600060808236031215610d7e57600080fd5b610d86610bb0565b610d8f83610c04565b815260208084013567ffffffffffffffff80821115610dad57600080fd5b610db936838801610cfc565b83850152604091508186013581811115610dd257600080fd5b610dde36828901610cfc565b8386015250606086013581811115610df557600080fd5b860136601f820112610e0657600080fd5b803582811115610e1857610e18610b58565b610e26858260051b01610bd3565b818152858101935060069190911b820185019036821115610e4657600080fd5b918501915b81831015610e8d57848336031215610e635760008081fd5b610e6b610b87565b610e7484610c21565b8152838701358782015284529285019291840191610e4b565b606087015250939695505050505050565b600060208284031215610eb057600080fd5b81518015158114610af157600080fd5b600060408284031215610ed257600080fd5b610eda610b87565b82518152602083015160208201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f79577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b6000604067ffffffffffffffff8516835260208181850152845160a083860152610fad60e0860182610af8565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080878403016060880152610fe88383610af8565b88860151888203830160808a01528051808352908601945060009350908501905b8084101561103b57845180516001600160a01b0316835286015186830152938501936001939093019290860190611009565b5060608901516001600160a01b031660a08901526080890151888203830160c08a0152955061106a8187610af8565b9a9950505050505050505050565b60006020828403121561108a57600080fd5b505191905056fea164736f6c634300080f000a",
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
