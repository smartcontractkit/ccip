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

var PingPongDemoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"FeeTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingPongCount\",\"type\":\"uint256\"}],\"name\":\"Ping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingPongCount\",\"type\":\"uint256\"}],\"name\":\"Pong\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCommon.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounterpartAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounterpartChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"counterpartChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"counterpartAddress\",\"type\":\"address\"}],\"name\":\"setCounterpart\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setCounterpartAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"setCounterpartChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startPingPong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200155e3803806200155e833981016040819052620000349162000239565b3380600084846001600160a01b0382166200006a576040516335fdcccd60e21b8152600060048201526024015b60405180910390fd5b6001600160a01b038216608052620000828162000128565b50506001600160a01b038216620000dc5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f0000000000000000604482015260640162000061565b600180546001600160a01b0319166001600160a01b03848116919091179091558116156200010f576200010f8162000170565b50506003805460ff60a01b191690555062000271915050565b600080546001600160a01b0319166001600160a01b038316908117825560405190917f722ff84c1234b2482061def5c82c6b5080c117b3cbb69d686844a051e4b8e7f391a250565b336001600160a01b03821603620001ca5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000061565b600280546001600160a01b0319166001600160a01b03838116918217909255600154604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b80516001600160a01b03811681146200023457600080fd5b919050565b600080604083850312156200024d57600080fd5b62000258836200021c565b915062000268602084016200021c565b90509250929050565b6080516112bc620002a260003960008181610228015281816104600152818161079d0152610ca101526112bc6000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638da5cb5b11610097578063b5a1101111610066578063b5a110111461026f578063ca1d209d14610282578063ca709a2514610295578063f2fde38b146102b357600080fd5b80638da5cb5b146101f55780639d2aede514610213578063b0f479a114610226578063b187bd261461024c57600080fd5b80632874d8bf116100d35780632874d8bf146101935780632b6e5d631461019b5780633015b91c146101da57806379ba5097146101ed57600080fd5b806301ffc9a714610105578063093685f31461012d578063146347761461016b57806316c38b3c14610180575b600080fd5b610118610113366004610d22565b6102c6565b60405190151581526020015b60405180910390f35b60025474010000000000000000000000000000000000000000900467ffffffffffffffff1660405167ffffffffffffffff9091168152602001610124565b61017e610179366004610d81565b61035f565b005b61017e61018e366004610daa565b6103ba565b61017e61040c565b60035473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610124565b61017e6101e8366004610dc7565b610448565b61017e6104d2565b60015473ffffffffffffffffffffffffffffffffffffffff166101b5565b61017e610221366004610e26565b6105d3565b7f00000000000000000000000000000000000000000000000000000000000000006101b5565b60035474010000000000000000000000000000000000000000900460ff16610118565b61017e61027d366004610e41565b610622565b61017e610290366004610e74565b6106c4565b60005473ffffffffffffffffffffffffffffffffffffffff166101b5565b61017e6102c1366004610e26565b610857565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f3015b91c00000000000000000000000000000000000000000000000000000000148061035957507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b610367610868565b6002805467ffffffffffffffff90921674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff909216919091179055565b6103c2610868565b6003805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b610414610868565b600380547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16905561044660016108e9565b565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146104be576040517fd7f733340000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6104cf6104ca82610fa9565b610b15565b50565b60025473ffffffffffffffffffffffffffffffffffffffff163314610553576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016104b5565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560028054909116905560405173ffffffffffffffffffffffffffffffffffffffff909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b6105db610868565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b61062a610868565b6002805467ffffffffffffffff90931674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff909316929092179091556003805473ffffffffffffffffffffffffffffffffffffffff9092167fffffffffffffffffffffffff0000000000000000000000000000000000000000909216919091179055565b6000805473ffffffffffffffffffffffffffffffffffffffff166040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810184905290915073ffffffffffffffffffffffffffffffffffffffff8216906323b872dd906064016020604051808303816000875af115801561075a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077e91906110db565b5073ffffffffffffffffffffffffffffffffffffffff811663095ea7b37f00000000000000000000000000000000000000000000000000000000000000006040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602481018590526044016020604051808303816000875af115801561082e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085291906110db565b505050565b61085f610868565b6104cf81610b6b565b60015473ffffffffffffffffffffffffffffffffffffffff163314610446576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016104b5565b8060011660010361092c576040518181527f48257dc961b6f792c2b78a080dacfed693b660960a702de21cee364e20270e2f9060200160405180910390a1610960565b6040518181527f58b69f57828e6962d216502094c54f6562f3bf082ba758966c3454f9e37b15259060200160405180910390a15b60008160405160200161097591815260200190565b60408051601f1981840301815260a08301825260035473ffffffffffffffffffffffffffffffffffffffff1660c0808501919091528251808503909101815260e08401835283526020808401829052825160008082529181018452919450929182019083610a05565b60408051808201909152600080825260208201528152602001906001900390816109de5790505b508152602001610a2a60005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff168152602001610adb604051806040016040528062030d408152602001600015158152506040805182516024820152602092830151151560448083019190915282518083039091018152606490910190915290810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f97a657c90000000000000000000000000000000000000000000000000000000017905290565b9052600254909150610b0f9074010000000000000000000000000000000000000000900467ffffffffffffffff1682610c61565b50505050565b60008160400151806020019051810190610b2f91906110f8565b60035490915074010000000000000000000000000000000000000000900460ff16610b6757610b67610b62826001611111565b6108e9565b5050565b3373ffffffffffffffffffffffffffffffffffffffff821603610bea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016104b5565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600154604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b6040517f96f4e9f900000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906396f4e9f990610cd8908690869060040161119d565b6020604051808303816000875af1158015610cf7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1b91906110f8565b9392505050565b600060208284031215610d3457600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610d1b57600080fd5b803567ffffffffffffffff81168114610d7c57600080fd5b919050565b600060208284031215610d9357600080fd5b610d1b82610d64565b80151581146104cf57600080fd5b600060208284031215610dbc57600080fd5b8135610d1b81610d9c565b600060208284031215610dd957600080fd5b813567ffffffffffffffff811115610df057600080fd5b820160808185031215610d1b57600080fd5b803573ffffffffffffffffffffffffffffffffffffffff81168114610d7c57600080fd5b600060208284031215610e3857600080fd5b610d1b82610e02565b60008060408385031215610e5457600080fd5b610e5d83610d64565b9150610e6b60208401610e02565b90509250929050565b600060208284031215610e8657600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff81118282101715610edf57610edf610e8d565b60405290565b6040805190810167ffffffffffffffff81118282101715610edf57610edf610e8d565b604051601f8201601f1916810167ffffffffffffffff81118282101715610f3157610f31610e8d565b604052919050565b600082601f830112610f4a57600080fd5b813567ffffffffffffffff811115610f6457610f64610e8d565b610f776020601f19601f84011601610f08565b818152846020838601011115610f8c57600080fd5b816020850160208301376000918101602001919091529392505050565b600060808236031215610fbb57600080fd5b610fc3610ebc565b610fcc83610d64565b815260208084013567ffffffffffffffff80821115610fea57600080fd5b610ff636838801610f39565b8385015260409150818601358181111561100f57600080fd5b61101b36828901610f39565b838601525060608601358181111561103257600080fd5b860136601f82011261104357600080fd5b80358281111561105557611055610e8d565b611063858260051b01610f08565b818152858101935060069190911b82018501903682111561108357600080fd5b918501915b818310156110ca578483360312156110a05760008081fd5b6110a8610ee5565b6110b184610e02565b8152838701358782015284529285019291840191611088565b606087015250939695505050505050565b6000602082840312156110ed57600080fd5b8151610d1b81610d9c565b60006020828403121561110a57600080fd5b5051919050565b6000821982111561114b577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500190565b6000815180845260005b818110156111765760208185018101518683018201520161115a565b81811115611188576000602083870101525b50601f01601f19169290920160200192915050565b6000604067ffffffffffffffff8516835260208181850152845160a0838601526111ca60e0860182611150565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808784030160608801526112058383611150565b88860151888203830160808a01528051808352908601945060009350908501905b80841015611265578451805173ffffffffffffffffffffffffffffffffffffffff16835286015186830152938501936001939093019290860190611226565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a015295506112a18187611150565b9a995050505050505050505056fea164736f6c634300080f000a",
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

func (_PingPongDemo *PingPongDemoCaller) GetFeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "getFeeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) GetFeeToken() (common.Address, error) {
	return _PingPongDemo.Contract.GetFeeToken(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) GetFeeToken() (common.Address, error) {
	return _PingPongDemo.Contract.GetFeeToken(&_PingPongDemo.CallOpts)
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

func (_PingPongDemo *PingPongDemoTransactor) CcipReceive(opts *bind.TransactOpts, message CommonAny2EVMMessage) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "ccipReceive", message)
}

func (_PingPongDemo *PingPongDemoSession) CcipReceive(message CommonAny2EVMMessage) (*types.Transaction, error) {
	return _PingPongDemo.Contract.CcipReceive(&_PingPongDemo.TransactOpts, message)
}

func (_PingPongDemo *PingPongDemoTransactorSession) CcipReceive(message CommonAny2EVMMessage) (*types.Transaction, error) {
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

type PingPongDemoFeeTokenSetIterator struct {
	Event *PingPongDemoFeeTokenSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PingPongDemoFeeTokenSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongDemoFeeTokenSet)
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
		it.Event = new(PingPongDemoFeeTokenSet)
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

func (it *PingPongDemoFeeTokenSetIterator) Error() error {
	return it.fail
}

func (it *PingPongDemoFeeTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PingPongDemoFeeTokenSet struct {
	FeeToken common.Address
	Raw      types.Log
}

func (_PingPongDemo *PingPongDemoFilterer) FilterFeeTokenSet(opts *bind.FilterOpts, feeToken []common.Address) (*PingPongDemoFeeTokenSetIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _PingPongDemo.contract.FilterLogs(opts, "FeeTokenSet", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &PingPongDemoFeeTokenSetIterator{contract: _PingPongDemo.contract, event: "FeeTokenSet", logs: logs, sub: sub}, nil
}

func (_PingPongDemo *PingPongDemoFilterer) WatchFeeTokenSet(opts *bind.WatchOpts, sink chan<- *PingPongDemoFeeTokenSet, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _PingPongDemo.contract.WatchLogs(opts, "FeeTokenSet", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PingPongDemoFeeTokenSet)
				if err := _PingPongDemo.contract.UnpackLog(event, "FeeTokenSet", log); err != nil {
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

func (_PingPongDemo *PingPongDemoFilterer) ParseFeeTokenSet(log types.Log) (*PingPongDemoFeeTokenSet, error) {
	event := new(PingPongDemoFeeTokenSet)
	if err := _PingPongDemo.contract.UnpackLog(event, "FeeTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	case _PingPongDemo.abi.Events["FeeTokenSet"].ID:
		return _PingPongDemo.ParseFeeTokenSet(log)
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

func (PingPongDemoFeeTokenSet) Topic() common.Hash {
	return common.HexToHash("0x722ff84c1234b2482061def5c82c6b5080c117b3cbb69d686844a051e4b8e7f3")
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

	GetFeeToken(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IsPaused(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message CommonAny2EVMMessage) (*types.Transaction, error)

	Fund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	SetCounterpart(opts *bind.TransactOpts, counterpartChainId uint64, counterpartAddress common.Address) (*types.Transaction, error)

	SetCounterpartAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error)

	SetCounterpartChainId(opts *bind.TransactOpts, chainId uint64) (*types.Transaction, error)

	SetPaused(opts *bind.TransactOpts, pause bool) (*types.Transaction, error)

	StartPingPong(opts *bind.TransactOpts) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterFeeTokenSet(opts *bind.FilterOpts, feeToken []common.Address) (*PingPongDemoFeeTokenSetIterator, error)

	WatchFeeTokenSet(opts *bind.WatchOpts, sink chan<- *PingPongDemoFeeTokenSet, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenSet(log types.Log) (*PingPongDemoFeeTokenSet, error)

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
