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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structGovernanceDapp.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deliverer\",\"type\":\"address\"}],\"name\":\"InvalidDeliverer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ConfigPropagated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"name\":\"ReceivedConfig\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structGovernanceDapp.CrossChainClone\",\"name\":\"clone\",\"type\":\"tuple\"}],\"name\":\"addClone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCommon.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structGovernanceDapp.FeeConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"changedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structGovernanceDapp.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"voteForNewFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620014cc380380620014cc8339810160408190526200003491620001f5565b33806000856001600160a01b03811662000069576040516335fdcccd60e21b8152600060048201526024015b60405180910390fd5b6001600160a01b039081166080528216620000c75760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f0000000000000000604482015260640162000060565b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000fa57620000fa8162000131565b5050825160025550602090910151600355600480546001600160a01b0319166001600160a01b039092169190911790555062000295565b336001600160a01b038216036200018b5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000060565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b0381168114620001f257600080fd5b50565b600080600083850360808112156200020c57600080fd5b84516200021981620001dc565b93506040601f19820112156200022e57600080fd5b50604080519081016001600160401b03811182821017156200026057634e487b7160e01b600052604160045260246000fd5b806040525060208501518152604085015160208201528092505060608401516200028a81620001dc565b809150509250925092565b6080516111ff620002cd600039600081816102010152818161039f0152818161052d0152818161062c0152610b6f01526111ff6000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806379ba509711610081578063b0f479a11161005b578063b0f479a1146101ff578063ca1d209d14610225578063f2fde38b1461023857600080fd5b806379ba5097146101a55780638157c96c146101ad5780638da5cb5b146101c057600080fd5b8063262a7b63116100b2578063262a7b631461013f5780633015b91c146101545780635fbbc0d21461016757600080fd5b806301ffc9a7146100ce578063181f5a77146100f6575b600080fd5b6100e16100dc366004610bf0565b61024b565b60405190151581526020015b60405180910390f35b6101326040518060400160405280601481526020017f476f7665726e616e63654461707020312e302e3000000000000000000000000081525081565b6040516100ed9190610c7f565b61015261014d366004610d7f565b6102e4565b005b610152610162366004610dbc565b610387565b6040805180820182526000808252602091820152815180830183526002548082526003549183019182528351908152905191810191909152016100ed565b610152610411565b6101526101bb366004610df7565b61050e565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ed565b7f00000000000000000000000000000000000000000000000000000000000000006101da565b610152610233366004610e0f565b610569565b610152610246366004610e28565b6106e5565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f3015b91c0000000000000000000000000000000000000000000000000000000014806102de57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b6102ec6106f6565b6005805460018101825560009190915281517f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0909101805460209093015173ffffffffffffffffffffffffffffffffffffffff1668010000000000000000027fffffffff0000000000000000000000000000000000000000000000000000000090931667ffffffffffffffff90921691909117919091179055565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146103fd576040517fd7f733340000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61040e61040982610eb3565b610779565b50565b60015473ffffffffffffffffffffffffffffffffffffffff163314610492576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016103f4565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6105166106f6565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161561055b5761055b816107e2565b803560025560200135600355565b600480546040517f23b872dd00000000000000000000000000000000000000000000000000000000815233928101929092523060248301526044820183905273ffffffffffffffffffffffffffffffffffffffff16906323b872dd906064016020604051808303816000875af11580156105e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060b9190610fe5565b5060045473ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000006040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602481018490526044016020604051808303816000875af11580156106bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e19190610fe5565b5050565b6106ed6106f6565b61040e81610a3a565b60005473ffffffffffffffffffffffffffffffffffffffff163314610777576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016103f4565b565b600081604001518060200190518101906107939190611007565b80516002819055602080830151600381905560408051938452918301529192507f583bd9682201bddaa3ee0ed61c39b397de860af4d62cfc5a62ecd30ca7342deb910160405180910390a15050565b6000816040516020016108079190813581526020918201359181019190915260400190565b60408051601f1981840301815291905260055490915060005b81811015610a345760006005828154811061083d5761083d611039565b60009182526020808320604080518082018252939091015467ffffffffffffffff8116845273ffffffffffffffffffffffffffffffffffffffff680100000000000000009091048116848401908152825160a081018452905190911660c0808301919091528251808303909101815260e082018352815280830189905281518581529283018252929450820190836108f7565b60408051808201909152600080825260208201528152602001906001900390816108d05790505b50815260045473ffffffffffffffffffffffffffffffffffffffff16602080830191909152604080518082018252620493e080825260009184019182528251602481019190915290511515604480830191909152825180830390910181526064909101825291820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f97a657c9000000000000000000000000000000000000000000000000000000001790529091015282519091506109b69082610b2f565b50815160208301516040517f5c49f05c94a889457073501ff9101fe42cbe2255c2ee91c5e4008d09cf4cb2a492610a1992909167ffffffffffffffff92909216825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b60405180910390a1505080610a2d90611068565b9050610820565b50505050565b3373ffffffffffffffffffffffffffffffffffffffff821603610ab9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016103f4565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517f96f4e9f900000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906396f4e9f990610ba690869086906004016110c7565b6020604051808303816000875af1158015610bc5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610be991906111d9565b9392505050565b600060208284031215610c0257600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610be957600080fd5b6000815180845260005b81811015610c5857602081850181015186830182015201610c3c565b81811115610c6a576000602083870101525b50601f01601f19169290920160200192915050565b602081526000610be96020830184610c32565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610ce457610ce4610c92565b60405290565b6040516080810167ffffffffffffffff81118282101715610ce457610ce4610c92565b604051601f8201601f1916810167ffffffffffffffff81118282101715610d3657610d36610c92565b604052919050565b803567ffffffffffffffff81168114610d5657600080fd5b919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610d5657600080fd5b600060408284031215610d9157600080fd5b610d99610cc1565b610da283610d3e565b8152610db060208401610d5b565b60208201529392505050565b600060208284031215610dce57600080fd5b813567ffffffffffffffff811115610de557600080fd5b820160808185031215610be957600080fd5b600060408284031215610e0957600080fd5b50919050565b600060208284031215610e2157600080fd5b5035919050565b600060208284031215610e3a57600080fd5b610be982610d5b565b600082601f830112610e5457600080fd5b813567ffffffffffffffff811115610e6e57610e6e610c92565b610e816020601f19601f84011601610d0d565b818152846020838601011115610e9657600080fd5b816020850160208301376000918101602001919091529392505050565b600060808236031215610ec557600080fd5b610ecd610cea565b610ed683610d3e565b815260208084013567ffffffffffffffff80821115610ef457600080fd5b610f0036838801610e43565b83850152604091508186013581811115610f1957600080fd5b610f2536828901610e43565b8386015250606086013581811115610f3c57600080fd5b860136601f820112610f4d57600080fd5b803582811115610f5f57610f5f610c92565b610f6d858260051b01610d0d565b818152858101935060069190911b820185019036821115610f8d57600080fd5b918501915b81831015610fd457848336031215610faa5760008081fd5b610fb2610cc1565b610fbb84610d5b565b8152838701358782015284529285019291840191610f92565b606087015250939695505050505050565b600060208284031215610ff757600080fd5b81518015158114610be957600080fd5b60006040828403121561101957600080fd5b611021610cc1565b82518152602083015160208201528091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110c0577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b6000604067ffffffffffffffff8516835260208181850152845160a0838601526110f460e0860182610c32565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08087840301606088015261112f8383610c32565b88860151888203830160808a01528051808352908601945060009350908501905b8084101561118f578451805173ffffffffffffffffffffffffffffffffffffffff16835286015186830152938501936001939093019290860190611150565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a015295506111cb8187610c32565b9a9950505050505050505050565b6000602082840312156111eb57600080fd5b505191905056fea164736f6c634300080f000a",
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
