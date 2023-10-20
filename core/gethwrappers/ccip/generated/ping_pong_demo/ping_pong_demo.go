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
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
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
	_ = abi.ConvertType
)

type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

var PingPongDemoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingPongCount\",\"type\":\"uint256\"}],\"name\":\"Ping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingPongCount\",\"type\":\"uint256\"}],\"name\":\"Pong\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundPingPong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounterpartAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounterpartChainSelector\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"counterpartChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"counterpartAddress\",\"type\":\"address\"}],\"name\":\"setCounterpart\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setCounterpartAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"setCounterpartChainSelector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startPingPong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405260056004553480156200001657600080fd5b506040516200197438038062001974833981016040819052620000399162000268565b33806000846001600160a01b0381166200006e576040516335fdcccd60e21b8152600060048201526024015b60405180910390fd5b6001600160a01b039081166080528216620000cc5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f0000000000000000604482015260640162000065565b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000ff57620000ff81620001a4565b50506002805460ff60a01b1916905550600380546001600160a01b0319166001600160a01b0383811691821790925560405163095ea7b360e01b8152918416600483015260001960248301529063095ea7b3906044016020604051808303816000875af115801562000175573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200019b9190620002a7565b505050620002d2565b336001600160a01b03821603620001fe5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000065565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b03811681146200026557600080fd5b50565b600080604083850312156200027c57600080fd5b825162000289816200024f565b60208401519092506200029c816200024f565b809150509250929050565b600060208284031215620002ba57600080fd5b81518015158114620002cb57600080fd5b9392505050565b60805161167162000303600039600081816101df0152818161052e015281816106900152610c8401526116716000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80638da5cb5b11610097578063b5a1101111610066578063b5a1101114610226578063bee518a414610239578063dbe4481c14610277578063f2fde38b1461027f57600080fd5b80638da5cb5b146101ac5780639d2aede5146101ca578063b0f479a1146101dd578063b187bd261461020357600080fd5b80632874d8bf116100d35780632874d8bf1461014a5780632b6e5d631461015257806379ba50971461019157806385572ffb1461019957600080fd5b806301ffc9a7146100fa57806316c38b3c146101225780631892b90614610137575b600080fd5b61010d610108366004610eaa565b610292565b60405190151581526020015b60405180910390f35b610135610130366004610ef3565b61032b565b005b610135610145366004610f32565b61037d565b6101356103d8565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610119565b610135610414565b6101356101a7366004610f4d565b610516565b60005473ffffffffffffffffffffffffffffffffffffffff1661016c565b6101356101d8366004610faa565b61059b565b7f000000000000000000000000000000000000000000000000000000000000000061016c565b60025474010000000000000000000000000000000000000000900460ff1661010d565b610135610234366004610fc7565b6105ea565b60015474010000000000000000000000000000000000000000900467ffffffffffffffff1660405167ffffffffffffffff9091168152602001610119565b61013561068c565b61013561028d366004610faa565b6109ef565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f85572ffb00000000000000000000000000000000000000000000000000000000148061032557507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b610333610a00565b6002805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b610385610a00565b6001805467ffffffffffffffff90921674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff909216919091179055565b6103e0610a00565b600280547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556104126001610a81565b565b60015473ffffffffffffffffffffffffffffffffffffffff16331461049a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610587576040517fd7f73334000000000000000000000000000000000000000000000000000000008152336004820152602401610491565b61059861059382611212565b610d5f565b50565b6105a3610a00565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6105f2610a00565b6001805467ffffffffffffffff90931674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff909316929092179091556002805473ffffffffffffffffffffffffffffffffffffffff9092167fffffffffffffffffffffffff0000000000000000000000000000000000000000909216919091179055565b60007f00000000000000000000000000000000000000000000000000000000000000006001546040517fa8d87a3b0000000000000000000000000000000000000000000000000000000081527401000000000000000000000000000000000000000090910467ffffffffffffffff16600482015273ffffffffffffffffffffffffffffffffffffffff919091169063a8d87a3b90602401602060405180830381865afa158015610740573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076491906112bf565b90508073ffffffffffffffffffffffffffffffffffffffff166354b714686040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d591906112dc565b6bffffffffffffffffffffffff166000036107ed5750565b60008173ffffffffffffffffffffffffffffffffffffffff1663d09dc3396040518163ffffffff1660e01b8152600401602060405180830381865afa15801561083a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085e919061130a565b12156108675750565b6000808273ffffffffffffffffffffffffffffffffffffffff1663b06d41bc6040518163ffffffff1660e01b8152600401600060405180830381865afa1580156108b5573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526108fb9190810190611323565b5080519091505b801561098257600082610916600184611429565b815181106109265761092661143c565b602002602001015190503073ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1603610971576001935050610982565b5061097b8161146b565b9050610902565b5081156109ea578273ffffffffffffffffffffffffffffffffffffffff1663eff7cc486040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156109d157600080fd5b505af11580156109e5573d6000803e3d6000fd5b505050505b505050565b6109f7610a00565b61059881610db5565b60005473ffffffffffffffffffffffffffffffffffffffff163314610412576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610491565b80600116600103610ac4576040518181527f48257dc961b6f792c2b78a080dacfed693b660960a702de21cee364e20270e2f9060200160405180910390a1610af8565b6040518181527f58b69f57828e6962d216502094c54f6562f3bf082ba758966c3454f9e37b15259060200160405180910390a15b600081604051602001610b0d91815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815260a08301825260025473ffffffffffffffffffffffffffffffffffffffff1660c0808501919091528251808503909101815260e08401835283526020808401829052825160008082529181018452919450929182019083610bbb565b6040805180820190915260008082526020820152815260200190600190039081610b945790505b50815260035473ffffffffffffffffffffffffffffffffffffffff1660208083019190915260408051808301825262030d4090819052815160248082019290925282518082039092018252604401825291820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f97a657c900000000000000000000000000000000000000000000000000000000179052909101526001546040517f96f4e9f90000000000000000000000000000000000000000000000000000000081529192507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16916396f4e9f991610cf2917401000000000000000000000000000000000000000090910467ffffffffffffffff16908590600401611504565b6020604051808303816000875af1158015610d11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d35919061130a565b506000600454118015610d525750600454610d509084611616565b155b156109ea576109ea61068c565b60008160600151806020019051810190610d79919061130a565b60025490915074010000000000000000000000000000000000000000900460ff16610db157610db1610dac826001611651565b610a81565b5050565b3373ffffffffffffffffffffffffffffffffffffffff821603610e34576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610491565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215610ebc57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610eec57600080fd5b9392505050565b600060208284031215610f0557600080fd5b81358015158114610eec57600080fd5b803567ffffffffffffffff81168114610f2d57600080fd5b919050565b600060208284031215610f4457600080fd5b610eec82610f15565b600060208284031215610f5f57600080fd5b813567ffffffffffffffff811115610f7657600080fd5b820160a08185031215610eec57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461059857600080fd5b600060208284031215610fbc57600080fd5b8135610eec81610f88565b60008060408385031215610fda57600080fd5b610fe383610f15565b91506020830135610ff381610f88565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561105057611050610ffe565b60405290565b60405160a0810167ffffffffffffffff8111828210171561105057611050610ffe565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156110c0576110c0610ffe565b604052919050565b600082601f8301126110d957600080fd5b813567ffffffffffffffff8111156110f3576110f3610ffe565b61112460207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611079565b81815284602083860101111561113957600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff82111561117057611170610ffe565b5060051b60200190565b600082601f83011261118b57600080fd5b813560206111a061119b83611156565b611079565b82815260069290921b840181019181810190868411156111bf57600080fd5b8286015b8481101561120757604081890312156111dc5760008081fd5b6111e461102d565b81356111ef81610f88565b815281850135858201528352918301916040016111c3565b509695505050505050565b600060a0823603121561122457600080fd5b61122c611056565b8235815261123c60208401610f15565b6020820152604083013567ffffffffffffffff8082111561125c57600080fd5b611268368387016110c8565b6040840152606085013591508082111561128157600080fd5b61128d368387016110c8565b606084015260808501359150808211156112a657600080fd5b506112b33682860161117a565b60808301525092915050565b6000602082840312156112d157600080fd5b8151610eec81610f88565b6000602082840312156112ee57600080fd5b81516bffffffffffffffffffffffff81168114610eec57600080fd5b60006020828403121561131c57600080fd5b5051919050565b600080604080848603121561133757600080fd5b835167ffffffffffffffff81111561134e57600080fd5b8401601f8101861361135f57600080fd5b8051602061136f61119b83611156565b82815260069290921b8301810191818101908984111561138e57600080fd5b938201935b838510156113e85785858b0312156113ab5760008081fd5b6113b361102d565b85516113be81610f88565b81528584015161ffff811681146113d55760008081fd5b8185015282529385019390820190611393565b97909101519698969750505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810381811115610325576103256113fa565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008161147a5761147a6113fa565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b6000815180845260005b818110156114c6576020818501810151868301820152016114aa565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6000604067ffffffffffffffff8516835260208181850152845160a08386015261153160e08601826114a0565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08087840301606088015261156c83836114a0565b88860151888203830160808a01528051808352908601945060009350908501905b808410156115cc578451805173ffffffffffffffffffffffffffffffffffffffff1683528601518683015293850193600193909301929086019061158d565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a0152955061160881876114a0565b9a9950505050505050505050565b60008261164c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500690565b80820180821115610325576103256113fa56fea164736f6c6343000813000a",
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
	parsed, err := PingPongDemoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

func (_PingPongDemo *PingPongDemoCaller) GetCounterpartChainSelector(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PingPongDemo.contract.Call(opts, &out, "getCounterpartChainSelector")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_PingPongDemo *PingPongDemoSession) GetCounterpartChainSelector() (uint64, error) {
	return _PingPongDemo.Contract.GetCounterpartChainSelector(&_PingPongDemo.CallOpts)
}

func (_PingPongDemo *PingPongDemoCallerSession) GetCounterpartChainSelector() (uint64, error) {
	return _PingPongDemo.Contract.GetCounterpartChainSelector(&_PingPongDemo.CallOpts)
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

func (_PingPongDemo *PingPongDemoTransactor) FundPingPong(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "fundPingPong")
}

func (_PingPongDemo *PingPongDemoSession) FundPingPong() (*types.Transaction, error) {
	return _PingPongDemo.Contract.FundPingPong(&_PingPongDemo.TransactOpts)
}

func (_PingPongDemo *PingPongDemoTransactorSession) FundPingPong() (*types.Transaction, error) {
	return _PingPongDemo.Contract.FundPingPong(&_PingPongDemo.TransactOpts)
}

func (_PingPongDemo *PingPongDemoTransactor) SetCounterpart(opts *bind.TransactOpts, counterpartChainSelector uint64, counterpartAddress common.Address) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "setCounterpart", counterpartChainSelector, counterpartAddress)
}

func (_PingPongDemo *PingPongDemoSession) SetCounterpart(counterpartChainSelector uint64, counterpartAddress common.Address) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpart(&_PingPongDemo.TransactOpts, counterpartChainSelector, counterpartAddress)
}

func (_PingPongDemo *PingPongDemoTransactorSession) SetCounterpart(counterpartChainSelector uint64, counterpartAddress common.Address) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpart(&_PingPongDemo.TransactOpts, counterpartChainSelector, counterpartAddress)
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

func (_PingPongDemo *PingPongDemoTransactor) SetCounterpartChainSelector(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error) {
	return _PingPongDemo.contract.Transact(opts, "setCounterpartChainSelector", chainSelector)
}

func (_PingPongDemo *PingPongDemoSession) SetCounterpartChainSelector(chainSelector uint64) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpartChainSelector(&_PingPongDemo.TransactOpts, chainSelector)
}

func (_PingPongDemo *PingPongDemoTransactorSession) SetCounterpartChainSelector(chainSelector uint64) (*types.Transaction, error) {
	return _PingPongDemo.Contract.SetCounterpartChainSelector(&_PingPongDemo.TransactOpts, chainSelector)
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

	GetCounterpartChainSelector(opts *bind.CallOpts) (uint64, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IsPaused(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	FundPingPong(opts *bind.TransactOpts) (*types.Transaction, error)

	SetCounterpart(opts *bind.TransactOpts, counterpartChainSelector uint64, counterpartAddress common.Address) (*types.Transaction, error)

	SetCounterpartAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error)

	SetCounterpartChainSelector(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error)

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
