// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ping_pong_demo

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

type ClientAny2EVMMessage struct {
	MessageId        [32]byte
	SourceChainId    uint64
	Sender           []byte
	Data             []byte
	DestTokenAmounts []ClientEVMTokenAmount
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

var PingPongDemoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingPongCount\",\"type\":\"uint256\"}],\"name\":\"Ping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingPongCount\",\"type\":\"uint256\"}],\"name\":\"Pong\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounterpartAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounterpartChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"counterpartChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"counterpartAddress\",\"type\":\"address\"}],\"name\":\"setCounterpart\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setCounterpartAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"setCounterpartChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startPingPong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001540380380620015408339810160408190526200003491620001f5565b33806000846001600160a01b03811662000069576040516335fdcccd60e21b8152600060048201526024015b60405180910390fd5b6001600160a01b039081166080528216620000c75760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f0000000000000000604482015260640162000060565b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000fa57620000fa8162000131565b50506002805460ff60a01b1916905550600380546001600160a01b0319166001600160a01b03929092169190911790555062000234565b336001600160a01b038216036200018b5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000060565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b0381168114620001f257600080fd5b50565b600080604083850312156200020957600080fd5b82516200021681620001dc565b60208401519092506200022981620001dc565b809150509250929050565b6080516112db620002656000396000818161021d01528181610539015281816107590152610abc01526112db6000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806385572ffb11610097578063b187bd2611610066578063b187bd2614610241578063b5a1101114610264578063ca1d209d14610277578063f2fde38b1461028a57600080fd5b806385572ffb146101d75780638da5cb5b146101ea5780639d2aede514610208578063b0f479a11461021b57600080fd5b806316c38b3c116100d357806316c38b3c146101755780632874d8bf146101885780632b6e5d631461019057806379ba5097146101cf57600080fd5b806301ffc9a7146100fa578063093685f3146101225780631463477614610160575b600080fd5b61010d610108366004610cba565b61029d565b60405190151581526020015b60405180910390f35b60015474010000000000000000000000000000000000000000900467ffffffffffffffff1660405167ffffffffffffffff9091168152602001610119565b61017361016e366004610d20565b610336565b005b610173610183366004610d49565b610391565b6101736103e3565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610119565b61017361041f565b6101736101e5366004610d66565b610521565b60005473ffffffffffffffffffffffffffffffffffffffff166101aa565b610173610216366004610dc5565b6105a6565b7f00000000000000000000000000000000000000000000000000000000000000006101aa565b60025474010000000000000000000000000000000000000000900460ff1661010d565b610173610272366004610de0565b6105f5565b610173610285366004610e13565b610697565b610173610298366004610dc5565b610812565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f85572ffb00000000000000000000000000000000000000000000000000000000148061033057507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b61033e610823565b6001805467ffffffffffffffff90921674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff909216919091179055565b610399610823565b6002805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b6103eb610823565b600280547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16905561041d60016108a4565b565b60015473ffffffffffffffffffffffffffffffffffffffff1633146104a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610592576040517fd7f7333400000000000000000000000000000000000000000000000000000000815233600482015260240161049c565b6105a361059e8261102f565b610b73565b50565b6105ae610823565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6105fd610823565b6001805467ffffffffffffffff90931674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff909316929092179091556002805473ffffffffffffffffffffffffffffffffffffffff9092167fffffffffffffffffffffffff0000000000000000000000000000000000000000909216919091179055565b6003546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810183905273ffffffffffffffffffffffffffffffffffffffff909116906323b872dd906064016020604051808303816000875af1158015610714573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061073891906110dc565b5060035473ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000006040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602481018490526044016020604051808303816000875af11580156107ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061080e91906110dc565b5050565b61081a610823565b6105a381610bc5565b60005473ffffffffffffffffffffffffffffffffffffffff16331461041d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161049c565b806001166001036108e7576040518181527f48257dc961b6f792c2b78a080dacfed693b660960a702de21cee364e20270e2f9060200160405180910390a161091b565b6040518181527f58b69f57828e6962d216502094c54f6562f3bf082ba758966c3454f9e37b15259060200160405180910390a15b60008160405160200161093091815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815260a08301825260025473ffffffffffffffffffffffffffffffffffffffff1660c0808501919091528251808503909101815260e084018352835260208084018290528251600080825291810184529194509291820190836109de565b60408051808201909152600080825260208201528152602001906001900390816109b75790505b50815260035473ffffffffffffffffffffffffffffffffffffffff1660208083019190915260408051808201825262030d4080825260009184019182528251602481019190915290511515604480830191909152825180830390910181526064909101825291820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f97a657c900000000000000000000000000000000000000000000000000000000179052909101526001546040517f96f4e9f90000000000000000000000000000000000000000000000000000000081529192507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16916396f4e9f991610b2a917401000000000000000000000000000000000000000090910467ffffffffffffffff16908590600401611164565b6020604051808303816000875af1158015610b49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b6d9190611276565b50505050565b60008160600151806020019051810190610b8d9190611276565b60025490915074010000000000000000000000000000000000000000900460ff1661080e5761080e610bc082600161128f565b6108a4565b3373ffffffffffffffffffffffffffffffffffffffff821603610c44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161049c565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215610ccc57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610cfc57600080fd5b9392505050565b803567ffffffffffffffff81168114610d1b57600080fd5b919050565b600060208284031215610d3257600080fd5b610cfc82610d03565b80151581146105a357600080fd5b600060208284031215610d5b57600080fd5b8135610cfc81610d3b565b600060208284031215610d7857600080fd5b813567ffffffffffffffff811115610d8f57600080fd5b820160a08185031215610cfc57600080fd5b803573ffffffffffffffffffffffffffffffffffffffff81168114610d1b57600080fd5b600060208284031215610dd757600080fd5b610cfc82610da1565b60008060408385031215610df357600080fd5b610dfc83610d03565b9150610e0a60208401610da1565b90509250929050565b600060208284031215610e2557600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610e7e57610e7e610e2c565b60405290565b60405160a0810167ffffffffffffffff81118282101715610e7e57610e7e610e2c565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610eee57610eee610e2c565b604052919050565b600082601f830112610f0757600080fd5b813567ffffffffffffffff811115610f2157610f21610e2c565b610f5260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610ea7565b818152846020838601011115610f6757600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112610f9557600080fd5b8135602067ffffffffffffffff821115610fb157610fb1610e2c565b610fbf818360051b01610ea7565b82815260069290921b84018101918181019086841115610fde57600080fd5b8286015b848110156110245760408189031215610ffb5760008081fd5b611003610e5b565b61100c82610da1565b81528185013585820152835291830191604001610fe2565b509695505050505050565b600060a0823603121561104157600080fd5b611049610e84565b8235815261105960208401610d03565b6020820152604083013567ffffffffffffffff8082111561107957600080fd5b61108536838701610ef6565b6040840152606085013591508082111561109e57600080fd5b6110aa36838701610ef6565b606084015260808501359150808211156110c357600080fd5b506110d036828601610f84565b60808301525092915050565b6000602082840312156110ee57600080fd5b8151610cfc81610d3b565b6000815180845260005b8181101561111f57602081850181015186830182015201611103565b81811115611131576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6000604067ffffffffffffffff8516835260208181850152845160a08386015261119160e08601826110f9565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808784030160608801526111cc83836110f9565b88860151888203830160808a01528051808352908601945060009350908501905b8084101561122c578451805173ffffffffffffffffffffffffffffffffffffffff168352860151868301529385019360019390930192908601906111ed565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a0152955061126881876110f9565b9a9950505050505050505050565b60006020828403121561128857600080fd5b5051919050565b600082198211156112c9577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50019056fea164736f6c634300080f000a",
}

var PingPongDemoABI = PingPongDemoMetaData.ABI

var PingPongDemoBin = PingPongDemoMetaData.Bin

func DeployPingPongDemo(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address, feeToken common.Address) (common.Address, *types.Transaction, *PingPongDemo, error) {
	parsed, err := PingPongDemoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PingPongDemoBin), backend, router, feeToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PingPongDemo{PingPongDemoCaller: PingPongDemoCaller{contract: contract}, PingPongDemoTransactor: PingPongDemoTransactor{contract: contract}, PingPongDemoFilterer: PingPongDemoFilterer{contract: contract}}, nil
}

type PingPongDemo struct {
	address common.Address
	abi     abi.ABI
	PingPongDemoCaller
	PingPongDemoTransactor
	PingPongDemoFilterer
}

type PingPongDemoCaller struct {
	contract *bind.BoundContract
}

type PingPongDemoTransactor struct {
	contract *bind.BoundContract
}

type PingPongDemoFilterer struct {
	contract *bind.BoundContract
}

type PingPongDemoSession struct {
	Contract     *PingPongDemo
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type PingPongDemoCallerSession struct {
	Contract *PingPongDemoCaller
	CallOpts bind.CallOpts
}

type PingPongDemoTransactorSession struct {
	Contract     *PingPongDemoTransactor
	TransactOpts bind.TransactOpts
}

type PingPongDemoRaw struct {
	Contract *PingPongDemo
}

type PingPongDemoCallerRaw struct {
	Contract *PingPongDemoCaller
}

type PingPongDemoTransactorRaw struct {
	Contract *PingPongDemoTransactor
}

func NewPingPongDemo(address common.Address, backend bind.ContractBackend) (*PingPongDemo, error) {
	abi, err := abi.JSON(strings.NewReader(PingPongDemoABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindPingPongDemo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PingPongDemo{address: address, abi: abi, PingPongDemoCaller: PingPongDemoCaller{contract: contract}, PingPongDemoTransactor: PingPongDemoTransactor{contract: contract}, PingPongDemoFilterer: PingPongDemoFilterer{contract: contract}}, nil
}

func NewPingPongDemoCaller(address common.Address, caller bind.ContractCaller) (*PingPongDemoCaller, error) {
	contract, err := bindPingPongDemo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PingPongDemoCaller{contract: contract}, nil
}

func NewPingPongDemoTransactor(address common.Address, transactor bind.ContractTransactor) (*PingPongDemoTransactor, error) {
	contract, err := bindPingPongDemo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PingPongDemoTransactor{contract: contract}, nil
}

func NewPingPongDemoFilterer(address common.Address, filterer bind.ContractFilterer) (*PingPongDemoFilterer, error) {
	contract, err := bindPingPongDemo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PingPongDemoFilterer{contract: contract}, nil
}

func bindPingPongDemo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PingPongDemoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_PingPongDemo *PingPongDemoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PingPongDemo.Contract.PingPongDemoCaller.contract.Call(opts, result, method, params...)
}

func (_PingPongDemo *PingPongDemoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPongDemo.Contract.PingPongDemoTransactor.contract.Transfer(opts)
}

func (_PingPongDemo *PingPongDemoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PingPongDemo.Contract.PingPongDemoTransactor.contract.Transact(opts, method, params...)
}

func (_PingPongDemo *PingPongDemoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PingPongDemo.Contract.contract.Call(opts, result, method, params...)
}

func (_PingPongDemo *PingPongDemoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPongDemo.Contract.contract.Transfer(opts)
}

func (_PingPongDemo *PingPongDemoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PingPongDemo.Contract.contract.Transact(opts, method, params...)
}

func (_PingPongDemo *PingPongDemoCaller) GetCounterpartAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "getCounterpartAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) GetCounterpartAddress() (common.Address, error) {
	return _PingPongDemo.Contract.GetCounterpartAddress(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) GetCounterpartAddress() (common.Address, error) {
	return _PingPongDemo.Contract.GetCounterpartAddress(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCaller) GetCounterpartChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "getCounterpartChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) GetCounterpartChainId() (uint64, error) {
	return _PingPongDemo.Contract.GetCounterpartChainId(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) GetCounterpartChainId() (uint64, error) {
	return _PingPongDemo.Contract.GetCounterpartChainId(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) GetRouter() (common.Address, error) {
	return _PingPongDemo.Contract.GetRouter(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) GetRouter() (common.Address, error) {
	return _PingPongDemo.Contract.GetRouter(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) IsPaused() (bool, error) {
	return _PingPongDemo.Contract.IsPaused(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) IsPaused() (bool, error) {
	return _PingPongDemo.Contract.IsPaused(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) Owner() (common.Address, error) {
	return _PingPongDemo.Contract.Owner(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) Owner() (common.Address, error) {
	return _PingPongDemo.Contract.Owner(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PingPongDemo.Contract.SupportsInterface(&_PingPongDemo.CallOpts, interfaceId)
}

func (_PingPongDemo *PingPongDemoCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PingPongDemo.Contract.SupportsInterface(&_PingPongDemo.CallOpts, interfaceId)
}

func (_PingPongDemo *PingPongDemoTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "acceptOwnership")
}

func (_PingPongDemo *PingPongDemoSession) AcceptOwnership() (*types.Transaction, error) {
	return _PingPongDemo.Contract.AcceptOwnership(&_PingPongDemo.TransactOpts)
}

func (_PingPongDemo *PingPongDemoTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PingPongDemo.Contract.AcceptOwnership(&_PingPongDemo.TransactOpts)
}

func (_PingPongDemo *PingPongDemoTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "ccipReceive", message)
}

func (_PingPongDemo *PingPongDemoSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _PingPongDemo.Contract.CcipReceive(&_PingPongDemo.TransactOpts, message)
}

func (_PingPongDemo *PingPongDemoTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _PingPongDemo.Contract.CcipReceive(&_PingPongDemo.TransactOpts, message)
}

func (_PingPongDemo *PingPongDemoTransactor) Fund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "fund", amount)
}

func (_PingPongDemo *PingPongDemoSession) Fund(amount *big.Int) (*types.Transaction, error) {
	return _PingPongDemo.Contract.Fund(&_PingPongDemo.TransactOpts, amount)
}

func (_PingPongDemo *PingPongDemoTransactorSession) Fund(amount *big.Int) (*types.Transaction, error) {
	return _PingPongDemo.Contract.Fund(&_PingPongDemo.TransactOpts, amount)
}

func (_PingPongDemo *PingPongDemoTransactor) SetCounterpart(opts *bind.TransactOpts, counterpartChainId uint64, counterpartAddress common.Address) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "setCounterpart", counterpartChainId, counterpartAddress)
}

func (_PingPongDemo *PingPongDemoSession) SetCounterpart(counterpartChainId uint64, counterpartAddress common.Address) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpart(&_PingPongDemo.TransactOpts, counterpartChainId, counterpartAddress)
}

func (_PingPongDemo *PingPongDemoTransactorSession) SetCounterpart(counterpartChainId uint64, counterpartAddress common.Address) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpart(&_PingPongDemo.TransactOpts, counterpartChainId, counterpartAddress)
}

func (_PingPongDemo *PingPongDemoTransactor) SetCounterpartAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "setCounterpartAddress", addr)
}

func (_PingPongDemo *PingPongDemoSession) SetCounterpartAddress(addr common.Address) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpartAddress(&_PingPongDemo.TransactOpts, addr)
}

func (_PingPongDemo *PingPongDemoTransactorSession) SetCounterpartAddress(addr common.Address) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpartAddress(&_PingPongDemo.TransactOpts, addr)
}

func (_PingPongDemo *PingPongDemoTransactor) SetCounterpartChainId(opts *bind.TransactOpts, chainId uint64) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "setCounterpartChainId", chainId)
}

func (_PingPongDemo *PingPongDemoSession) SetCounterpartChainId(chainId uint64) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpartChainId(&_PingPongDemo.TransactOpts, chainId)
}

func (_PingPongDemo *PingPongDemoTransactorSession) SetCounterpartChainId(chainId uint64) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpartChainId(&_PingPongDemo.TransactOpts, chainId)
}

func (_PingPongDemo *PingPongDemoTransactor) SetPaused(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "setPaused", pause)
}

func (_PingPongDemo *PingPongDemoSession) SetPaused(pause bool) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetPaused(&_PingPongDemo.TransactOpts, pause)
}

func (_PingPongDemo *PingPongDemoTransactorSession) SetPaused(pause bool) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetPaused(&_PingPongDemo.TransactOpts, pause)
}

func (_PingPongDemo *PingPongDemoTransactor) StartPingPong(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "startPingPong")
}

func (_PingPongDemo *PingPongDemoSession) StartPingPong() (*types.Transaction, error) {
	return _PingPongDemo.Contract.StartPingPong(&_PingPongDemo.TransactOpts)
}

func (_PingPongDemo *PingPongDemoTransactorSession) StartPingPong() (*types.Transaction, error) {
	return _PingPongDemo.Contract.StartPingPong(&_PingPongDemo.TransactOpts)
}

func (_PingPongDemo *PingPongDemoTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "transferOwnership", to)
}

func (_PingPongDemo *PingPongDemoSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _PingPongDemo.Contract.TransferOwnership(&_PingPongDemo.TransactOpts, to)
}

func (_PingPongDemo *PingPongDemoTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _PingPongDemo.Contract.TransferOwnership(&_PingPongDemo.TransactOpts, to)
}

type PingPongDemoOwnershipTransferRequestedIterator struct {
	Event *PingPongDemoOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PingPongDemoOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongDemoOwnershipTransferRequested)
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
		it.Event = new(PingPongDemoOwnershipTransferRequested)
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

func (it *PingPongDemoOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *PingPongDemoOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PingPongDemoOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_PingPongDemo *PingPongDemoFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PingPongDemoOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PingPongDemo.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PingPongDemoOwnershipTransferRequestedIterator{contract: _PingPongDemo.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_PingPongDemo *PingPongDemoFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *PingPongDemoOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PingPongDemo.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PingPongDemoOwnershipTransferRequested)
				if err := _PingPongDemo.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_PingPongDemo *PingPongDemoFilterer) ParseOwnershipTransferRequested(log types.Log) (*PingPongDemoOwnershipTransferRequested, error) {
	event := new(PingPongDemoOwnershipTransferRequested)
	if err := _PingPongDemo.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PingPongDemoOwnershipTransferredIterator struct {
	Event *PingPongDemoOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PingPongDemoOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongDemoOwnershipTransferred)
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
		it.Event = new(PingPongDemoOwnershipTransferred)
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

func (it *PingPongDemoOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *PingPongDemoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PingPongDemoOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_PingPongDemo *PingPongDemoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PingPongDemoOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PingPongDemo.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PingPongDemoOwnershipTransferredIterator{contract: _PingPongDemo.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_PingPongDemo *PingPongDemoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PingPongDemoOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PingPongDemo.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PingPongDemoOwnershipTransferred)
				if err := _PingPongDemo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_PingPongDemo *PingPongDemoFilterer) ParseOwnershipTransferred(log types.Log) (*PingPongDemoOwnershipTransferred, error) {
	event := new(PingPongDemoOwnershipTransferred)
	if err := _PingPongDemo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PingPongDemoPingIterator struct {
	Event *PingPongDemoPing

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PingPongDemoPingIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongDemoPing)
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
		it.Event = new(PingPongDemoPing)
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

func (it *PingPongDemoPingIterator) Error() error {
	return it.fail
}

func (it *PingPongDemoPingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PingPongDemoPing struct {
	PingPongCount *big.Int
	Raw           types.Log
}

func (_PingPongDemo *PingPongDemoFilterer) FilterPing(opts *bind.FilterOpts) (*PingPongDemoPingIterator, error) {

	logs, sub, err := _PingPongDemo.contract.FilterLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return &PingPongDemoPingIterator{contract: _PingPongDemo.contract, event: "Ping", logs: logs, sub: sub}, nil
}

func (_PingPongDemo *PingPongDemoFilterer) WatchPing(opts *bind.WatchOpts, sink chan<- *PingPongDemoPing) (event.Subscription, error) {

	logs, sub, err := _PingPongDemo.contract.WatchLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PingPongDemoPing)
				if err := _PingPongDemo.contract.UnpackLog(event, "Ping", log); err != nil {
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

func (_PingPongDemo *PingPongDemoFilterer) ParsePing(log types.Log) (*PingPongDemoPing, error) {
	event := new(PingPongDemoPing)
	if err := _PingPongDemo.contract.UnpackLog(event, "Ping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PingPongDemoPongIterator struct {
	Event *PingPongDemoPong

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PingPongDemoPongIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongDemoPong)
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
		it.Event = new(PingPongDemoPong)
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

func (it *PingPongDemoPongIterator) Error() error {
	return it.fail
}

func (it *PingPongDemoPongIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PingPongDemoPong struct {
	PingPongCount *big.Int
	Raw           types.Log
}

func (_PingPongDemo *PingPongDemoFilterer) FilterPong(opts *bind.FilterOpts) (*PingPongDemoPongIterator, error) {

	logs, sub, err := _PingPongDemo.contract.FilterLogs(opts, "Pong")
	if err != nil {
		return nil, err
	}
	return &PingPongDemoPongIterator{contract: _PingPongDemo.contract, event: "Pong", logs: logs, sub: sub}, nil
}

func (_PingPongDemo *PingPongDemoFilterer) WatchPong(opts *bind.WatchOpts, sink chan<- *PingPongDemoPong) (event.Subscription, error) {

	logs, sub, err := _PingPongDemo.contract.WatchLogs(opts, "Pong")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PingPongDemoPong)
				if err := _PingPongDemo.contract.UnpackLog(event, "Pong", log); err != nil {
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

func (_PingPongDemo *PingPongDemoFilterer) ParsePong(log types.Log) (*PingPongDemoPong, error) {
	event := new(PingPongDemoPong)
	if err := _PingPongDemo.contract.UnpackLog(event, "Pong", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_PingPongDemo *PingPongDemo) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _PingPongDemo.abi.Events["OwnershipTransferRequested"].ID:
		return _PingPongDemo.ParseOwnershipTransferRequested(log)
	case _PingPongDemo.abi.Events["OwnershipTransferred"].ID:
		return _PingPongDemo.ParseOwnershipTransferred(log)
	case _PingPongDemo.abi.Events["Ping"].ID:
		return _PingPongDemo.ParsePing(log)
	case _PingPongDemo.abi.Events["Pong"].ID:
		return _PingPongDemo.ParsePong(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (PingPongDemoOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (PingPongDemoOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (PingPongDemoPing) Topic() common.Hash {
	return common.HexToHash("0x48257dc961b6f792c2b78a080dacfed693b660960a702de21cee364e20270e2f")
}

func (PingPongDemoPong) Topic() common.Hash {
	return common.HexToHash("0x58b69f57828e6962d216502094c54f6562f3bf082ba758966c3454f9e37b1525")
}

func (_PingPongDemo *PingPongDemo) Address() common.Address {
	return _PingPongDemo.address
}

type PingPongDemoInterface interface {
	GetCounterpartAddress(opts *bind.CallOpts) (common.Address, error)

	GetCounterpartChainId(opts *bind.CallOpts) (uint64, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IsPaused(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	Fund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	SetCounterpart(opts *bind.TransactOpts, counterpartChainId uint64, counterpartAddress common.Address) (*types.Transaction, error)

	SetCounterpartAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error)

	SetCounterpartChainId(opts *bind.TransactOpts, chainId uint64) (*types.Transaction, error)

	SetPaused(opts *bind.TransactOpts, pause bool) (*types.Transaction, error)

	StartPingPong(opts *bind.TransactOpts) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PingPongDemoOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *PingPongDemoOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*PingPongDemoOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PingPongDemoOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PingPongDemoOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*PingPongDemoOwnershipTransferred, error)

	FilterPing(opts *bind.FilterOpts) (*PingPongDemoPingIterator, error)

	WatchPing(opts *bind.WatchOpts, sink chan<- *PingPongDemoPing) (event.Subscription, error)

	ParsePing(log types.Log) (*PingPongDemoPing, error)

	FilterPong(opts *bind.FilterOpts) (*PingPongDemoPongIterator, error)

	WatchPong(opts *bind.WatchOpts, sink chan<- *PingPongDemoPong) (event.Subscription, error)

	ParsePong(log types.Log) (*PingPongDemoPong, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
