// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package self_funded_ping_pong

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

var SelfFundedPingPongMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"roundTripsBeforeFunding\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Funded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingPongCount\",\"type\":\"uint256\"}],\"name\":\"Ping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingPongCount\",\"type\":\"uint256\"}],\"name\":\"Pong\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pingPongCount\",\"type\":\"uint256\"}],\"name\":\"fundPingPong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCountIncrBeforeFunding\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounterpartAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounterpartChainSelector\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"counterpartChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"counterpartAddress\",\"type\":\"address\"}],\"name\":\"setCounterpart\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setCounterpartAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"setCounterpartChainSelector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startPingPong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200174538038062001745833981016040819052620000349162000291565b828233806000846001600160a01b0381166200006b576040516335fdcccd60e21b8152600060048201526024015b60405180910390fd5b6001600160a01b039081166080528216620000c95760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f0000000000000000604482015260640162000062565b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000fc57620000fc81620001cd565b50506002805460ff60a01b1916905550600380546001600160a01b0319166001600160a01b0383811691821790925560405163095ea7b360e01b8152918416600483015260001960248301529063095ea7b3906044016020604051808303816000875af115801562000172573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001989190620002ea565b505050806002620001aa919062000315565b600360146101000a81548160ff021916908360ff16021790555050505062000347565b336001600160a01b03821603620002275760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000062565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b03811681146200028e57600080fd5b50565b600080600060608486031215620002a757600080fd5b8351620002b48162000278565b6020850151909350620002c78162000278565b604085015190925060ff81168114620002df57600080fd5b809150509250925092565b600060208284031215620002fd57600080fd5b815180151581146200030e57600080fd5b9392505050565b60ff81811683821602908116908181146200034057634e487b7160e01b600052601160045260246000fd5b5092915050565b6080516113cd6200037860003960008181610218015281816105ad0152818161069a0152610b6801526113cd6000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c80638f491cba116100b2578063b5a1101111610081578063ca709a2511610066578063ca709a25146102b0578063e6c725f5146102ce578063f2fde38b146102fe57600080fd5b8063b5a110111461025f578063bee518a41461027257600080fd5b80638f491cba146101f05780639d2aede514610203578063b0f479a114610216578063b187bd261461023c57600080fd5b80632b6e5d63116100ee5780632b6e5d631461017857806379ba5097146101b757806385572ffb146101bf5780638da5cb5b146101d257600080fd5b806301ffc9a71461012057806316c38b3c146101485780631892b9061461015d5780632874d8bf14610170575b600080fd5b61013361012e366004610d80565b610311565b60405190151581526020015b60405180910390f35b61015b610156366004610dc9565b6103aa565b005b61015b61016b366004610e08565b6103fc565b61015b610457565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161013f565b61015b610493565b61015b6101cd366004610e23565b610595565b60005473ffffffffffffffffffffffffffffffffffffffff16610192565b61015b6101fe366004610e5e565b61061a565b61015b610211366004610e99565b61080c565b7f0000000000000000000000000000000000000000000000000000000000000000610192565b60025474010000000000000000000000000000000000000000900460ff16610133565b61015b61026d366004610eb6565b61085b565b60015474010000000000000000000000000000000000000000900467ffffffffffffffff1660405167ffffffffffffffff909116815260200161013f565b60035473ffffffffffffffffffffffffffffffffffffffff16610192565b60035474010000000000000000000000000000000000000000900460ff1660405160ff909116815260200161013f565b61015b61030c366004610e99565b6108fd565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f85572ffb0000000000000000000000000000000000000000000000000000000014806103a457507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b6103b261090e565b6002805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b61040461090e565b6001805467ffffffffffffffff90921674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff909216919091179055565b61045f61090e565b600280547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff169055610491600161098f565b565b60015473ffffffffffffffffffffffffffffffffffffffff163314610519576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610606576040517fd7f73334000000000000000000000000000000000000000000000000000000008152336004820152602401610510565b610617610612826110f2565b610c35565b50565b60035474010000000000000000000000000000000000000000900460ff161580610660575060035474010000000000000000000000000000000000000000900460ff1681105b156106685750565b6003546001906106939074010000000000000000000000000000000000000000900460ff168361119f565b11610617577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a8d87a3b61070060015467ffffffffffffffff740100000000000000000000000000000000000000009091041690565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815267ffffffffffffffff9091166004820152602401602060405180830381865afa15801561075d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061078191906111da565b73ffffffffffffffffffffffffffffffffffffffff1663eff7cc486040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156107c857600080fd5b505af11580156107dc573d6000803e3d6000fd5b50506040517f302777af5d26fab9dd5120c5f1307c65193ebc51daf33244ada4365fab10602c925060009150a150565b61081461090e565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b61086361090e565b6001805467ffffffffffffffff90931674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff909316929092179091556002805473ffffffffffffffffffffffffffffffffffffffff9092167fffffffffffffffffffffffff0000000000000000000000000000000000000000909216919091179055565b61090561090e565b61061781610c8b565b60005473ffffffffffffffffffffffffffffffffffffffff163314610491576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610510565b806001166001036109d2576040518181527f48257dc961b6f792c2b78a080dacfed693b660960a702de21cee364e20270e2f9060200160405180910390a1610a06565b6040518181527f58b69f57828e6962d216502094c54f6562f3bf082ba758966c3454f9e37b15259060200160405180910390a15b610a0f8161061a565b600081604051602001610a2491815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815260a08301909152915060009080610a7b60025473ffffffffffffffffffffffffffffffffffffffff1690565b6040805173ffffffffffffffffffffffffffffffffffffffff909216602083015201604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00181529181529082526020808301869052815160008082529181018352929091019190610b14565b6040805180820190915260008082526020820152815260200190600190039081610aed5790505b508152602001610b3960035473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1681526040805160208082019092526000815291015290507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166396f4e9f9610bce60015467ffffffffffffffff740100000000000000000000000000000000000000009091041690565b836040518363ffffffff1660e01b8152600401610bec92919061125b565b6020604051808303816000875af1158015610c0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c2f919061136d565b50505050565b60008160600151806020019051810190610c4f919061136d565b60025490915074010000000000000000000000000000000000000000900460ff16610c8757610c87610c82826001611386565b61098f565b5050565b3373ffffffffffffffffffffffffffffffffffffffff821603610d0a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610510565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215610d9257600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610dc257600080fd5b9392505050565b600060208284031215610ddb57600080fd5b81358015158114610dc257600080fd5b803567ffffffffffffffff81168114610e0357600080fd5b919050565b600060208284031215610e1a57600080fd5b610dc282610deb565b600060208284031215610e3557600080fd5b813567ffffffffffffffff811115610e4c57600080fd5b820160a08185031215610dc257600080fd5b600060208284031215610e7057600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461061757600080fd5b600060208284031215610eab57600080fd5b8135610dc281610e77565b60008060408385031215610ec957600080fd5b610ed283610deb565b91506020830135610ee281610e77565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610f3f57610f3f610eed565b60405290565b60405160a0810167ffffffffffffffff81118282101715610f3f57610f3f610eed565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610faf57610faf610eed565b604052919050565b600082601f830112610fc857600080fd5b813567ffffffffffffffff811115610fe257610fe2610eed565b61101360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610f68565b81815284602083860101111561102857600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261105657600080fd5b8135602067ffffffffffffffff82111561107257611072610eed565b611080818360051b01610f68565b82815260069290921b8401810191818101908684111561109f57600080fd5b8286015b848110156110e757604081890312156110bc5760008081fd5b6110c4610f1c565b81356110cf81610e77565b815281850135858201528352918301916040016110a3565b509695505050505050565b600060a0823603121561110457600080fd5b61110c610f45565b8235815261111c60208401610deb565b6020820152604083013567ffffffffffffffff8082111561113c57600080fd5b61114836838701610fb7565b6040840152606085013591508082111561116157600080fd5b61116d36838701610fb7565b6060840152608085013591508082111561118657600080fd5b5061119336828601611045565b60808301525092915050565b6000826111d5577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500690565b6000602082840312156111ec57600080fd5b8151610dc281610e77565b6000815180845260005b8181101561121d57602081850181015186830182015201611201565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6000604067ffffffffffffffff8516835260208181850152845160a08386015261128860e08601826111f7565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808784030160608801526112c383836111f7565b88860151888203830160808a01528051808352908601945060009350908501905b80841015611323578451805173ffffffffffffffffffffffffffffffffffffffff168352860151868301529385019360019390930192908601906112e4565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a0152955061135f81876111f7565b9a9950505050505050505050565b60006020828403121561137f57600080fd5b5051919050565b808201808211156103a4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea164736f6c6343000813000a",
}

var SelfFundedPingPongABI = SelfFundedPingPongMetaData.ABI

var SelfFundedPingPongBin = SelfFundedPingPongMetaData.Bin

func DeploySelfFundedPingPong(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address, feeToken common.Address, roundTripsBeforeFunding uint8) (common.Address, *types.Transaction, *SelfFundedPingPong, error) {
	parsed, err := SelfFundedPingPongMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SelfFundedPingPongBin), backend, router, feeToken, roundTripsBeforeFunding)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SelfFundedPingPong{SelfFundedPingPongCaller: SelfFundedPingPongCaller{contract: contract}, SelfFundedPingPongTransactor: SelfFundedPingPongTransactor{contract: contract}, SelfFundedPingPongFilterer: SelfFundedPingPongFilterer{contract: contract}}, nil
}

type SelfFundedPingPong struct {
	address common.Address
	abi     abi.ABI
	SelfFundedPingPongCaller
	SelfFundedPingPongTransactor
	SelfFundedPingPongFilterer
}

type SelfFundedPingPongCaller struct {
	contract *bind.BoundContract
}

type SelfFundedPingPongTransactor struct {
	contract *bind.BoundContract
}

type SelfFundedPingPongFilterer struct {
	contract *bind.BoundContract
}

type SelfFundedPingPongSession struct {
	Contract     *SelfFundedPingPong
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type SelfFundedPingPongCallerSession struct {
	Contract *SelfFundedPingPongCaller
	CallOpts bind.CallOpts
}

type SelfFundedPingPongTransactorSession struct {
	Contract     *SelfFundedPingPongTransactor
	TransactOpts bind.TransactOpts
}

type SelfFundedPingPongRaw struct {
	Contract *SelfFundedPingPong
}

type SelfFundedPingPongCallerRaw struct {
	Contract *SelfFundedPingPongCaller
}

type SelfFundedPingPongTransactorRaw struct {
	Contract *SelfFundedPingPongTransactor
}

func NewSelfFundedPingPong(address common.Address, backend bind.ContractBackend) (*SelfFundedPingPong, error) {
	abi, err := abi.JSON(strings.NewReader(SelfFundedPingPongABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindSelfFundedPingPong(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPong{address: address, abi: abi, SelfFundedPingPongCaller: SelfFundedPingPongCaller{contract: contract}, SelfFundedPingPongTransactor: SelfFundedPingPongTransactor{contract: contract}, SelfFundedPingPongFilterer: SelfFundedPingPongFilterer{contract: contract}}, nil
}

func NewSelfFundedPingPongCaller(address common.Address, caller bind.ContractCaller) (*SelfFundedPingPongCaller, error) {
	contract, err := bindSelfFundedPingPong(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongCaller{contract: contract}, nil
}

func NewSelfFundedPingPongTransactor(address common.Address, transactor bind.ContractTransactor) (*SelfFundedPingPongTransactor, error) {
	contract, err := bindSelfFundedPingPong(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongTransactor{contract: contract}, nil
}

func NewSelfFundedPingPongFilterer(address common.Address, filterer bind.ContractFilterer) (*SelfFundedPingPongFilterer, error) {
	contract, err := bindSelfFundedPingPong(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongFilterer{contract: contract}, nil
}

func bindSelfFundedPingPong(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SelfFundedPingPongMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_SelfFundedPingPong *SelfFundedPingPongRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfFundedPingPong.Contract.SelfFundedPingPongCaller.contract.Call(opts, result, method, params...)
}

func (_SelfFundedPingPong *SelfFundedPingPongRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SelfFundedPingPongTransactor.contract.Transfer(opts)
}

func (_SelfFundedPingPong *SelfFundedPingPongRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SelfFundedPingPongTransactor.contract.Transact(opts, method, params...)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfFundedPingPong.Contract.contract.Call(opts, result, method, params...)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.contract.Transfer(opts)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.contract.Transact(opts, method, params...)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) GetCountIncrBeforeFunding(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "getCountIncrBeforeFunding")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) GetCountIncrBeforeFunding() (uint8, error) {
	return _SelfFundedPingPong.Contract.GetCountIncrBeforeFunding(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) GetCountIncrBeforeFunding() (uint8, error) {
	return _SelfFundedPingPong.Contract.GetCountIncrBeforeFunding(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) GetCounterpartAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "getCounterpartAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) GetCounterpartAddress() (common.Address, error) {
	return _SelfFundedPingPong.Contract.GetCounterpartAddress(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) GetCounterpartAddress() (common.Address, error) {
	return _SelfFundedPingPong.Contract.GetCounterpartAddress(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) GetCounterpartChainSelector(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "getCounterpartChainSelector")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) GetCounterpartChainSelector() (uint64, error) {
	return _SelfFundedPingPong.Contract.GetCounterpartChainSelector(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) GetCounterpartChainSelector() (uint64, error) {
	return _SelfFundedPingPong.Contract.GetCounterpartChainSelector(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) GetFeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "getFeeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) GetFeeToken() (common.Address, error) {
	return _SelfFundedPingPong.Contract.GetFeeToken(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) GetFeeToken() (common.Address, error) {
	return _SelfFundedPingPong.Contract.GetFeeToken(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) GetRouter() (common.Address, error) {
	return _SelfFundedPingPong.Contract.GetRouter(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) GetRouter() (common.Address, error) {
	return _SelfFundedPingPong.Contract.GetRouter(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) IsPaused() (bool, error) {
	return _SelfFundedPingPong.Contract.IsPaused(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) IsPaused() (bool, error) {
	return _SelfFundedPingPong.Contract.IsPaused(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) Owner() (common.Address, error) {
	return _SelfFundedPingPong.Contract.Owner(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) Owner() (common.Address, error) {
	return _SelfFundedPingPong.Contract.Owner(&_SelfFundedPingPong.CallOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SelfFundedPingPong.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SelfFundedPingPong *SelfFundedPingPongSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SelfFundedPingPong.Contract.SupportsInterface(&_SelfFundedPingPong.CallOpts, interfaceId)
}

func (_SelfFundedPingPong *SelfFundedPingPongCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SelfFundedPingPong.Contract.SupportsInterface(&_SelfFundedPingPong.CallOpts, interfaceId)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "acceptOwnership")
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) AcceptOwnership() (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.AcceptOwnership(&_SelfFundedPingPong.TransactOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.AcceptOwnership(&_SelfFundedPingPong.TransactOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "ccipReceive", message)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.CcipReceive(&_SelfFundedPingPong.TransactOpts, message)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.CcipReceive(&_SelfFundedPingPong.TransactOpts, message)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) FundPingPong(opts *bind.TransactOpts, pingPongCount *big.Int) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "fundPingPong", pingPongCount)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) FundPingPong(pingPongCount *big.Int) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.FundPingPong(&_SelfFundedPingPong.TransactOpts, pingPongCount)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) FundPingPong(pingPongCount *big.Int) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.FundPingPong(&_SelfFundedPingPong.TransactOpts, pingPongCount)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) SetCounterpart(opts *bind.TransactOpts, counterpartChainSelector uint64, counterpartAddress common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "setCounterpart", counterpartChainSelector, counterpartAddress)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) SetCounterpart(counterpartChainSelector uint64, counterpartAddress common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetCounterpart(&_SelfFundedPingPong.TransactOpts, counterpartChainSelector, counterpartAddress)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) SetCounterpart(counterpartChainSelector uint64, counterpartAddress common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetCounterpart(&_SelfFundedPingPong.TransactOpts, counterpartChainSelector, counterpartAddress)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) SetCounterpartAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "setCounterpartAddress", addr)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) SetCounterpartAddress(addr common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetCounterpartAddress(&_SelfFundedPingPong.TransactOpts, addr)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) SetCounterpartAddress(addr common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetCounterpartAddress(&_SelfFundedPingPong.TransactOpts, addr)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) SetCounterpartChainSelector(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "setCounterpartChainSelector", chainSelector)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) SetCounterpartChainSelector(chainSelector uint64) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetCounterpartChainSelector(&_SelfFundedPingPong.TransactOpts, chainSelector)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) SetCounterpartChainSelector(chainSelector uint64) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetCounterpartChainSelector(&_SelfFundedPingPong.TransactOpts, chainSelector)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) SetPaused(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "setPaused", pause)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) SetPaused(pause bool) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetPaused(&_SelfFundedPingPong.TransactOpts, pause)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) SetPaused(pause bool) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.SetPaused(&_SelfFundedPingPong.TransactOpts, pause)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) StartPingPong(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "startPingPong")
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) StartPingPong() (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.StartPingPong(&_SelfFundedPingPong.TransactOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) StartPingPong() (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.StartPingPong(&_SelfFundedPingPong.TransactOpts)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.contract.Transact(opts, "transferOwnership", to)
}

func (_SelfFundedPingPong *SelfFundedPingPongSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.TransferOwnership(&_SelfFundedPingPong.TransactOpts, to)
}

func (_SelfFundedPingPong *SelfFundedPingPongTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SelfFundedPingPong.Contract.TransferOwnership(&_SelfFundedPingPong.TransactOpts, to)
}

type SelfFundedPingPongFundedIterator struct {
	Event *SelfFundedPingPongFunded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SelfFundedPingPongFundedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SelfFundedPingPongFunded)
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
		it.Event = new(SelfFundedPingPongFunded)
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

func (it *SelfFundedPingPongFundedIterator) Error() error {
	return it.fail
}

func (it *SelfFundedPingPongFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SelfFundedPingPongFunded struct {
	Raw types.Log
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) FilterFunded(opts *bind.FilterOpts) (*SelfFundedPingPongFundedIterator, error) {

	logs, sub, err := _SelfFundedPingPong.contract.FilterLogs(opts, "Funded")
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongFundedIterator{contract: _SelfFundedPingPong.contract, event: "Funded", logs: logs, sub: sub}, nil
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) WatchFunded(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongFunded) (event.Subscription, error) {

	logs, sub, err := _SelfFundedPingPong.contract.WatchLogs(opts, "Funded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SelfFundedPingPongFunded)
				if err := _SelfFundedPingPong.contract.UnpackLog(event, "Funded", log); err != nil {
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

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) ParseFunded(log types.Log) (*SelfFundedPingPongFunded, error) {
	event := new(SelfFundedPingPongFunded)
	if err := _SelfFundedPingPong.contract.UnpackLog(event, "Funded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SelfFundedPingPongOwnershipTransferRequestedIterator struct {
	Event *SelfFundedPingPongOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SelfFundedPingPongOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SelfFundedPingPongOwnershipTransferRequested)
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
		it.Event = new(SelfFundedPingPongOwnershipTransferRequested)
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

func (it *SelfFundedPingPongOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *SelfFundedPingPongOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SelfFundedPingPongOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SelfFundedPingPongOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SelfFundedPingPong.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongOwnershipTransferRequestedIterator{contract: _SelfFundedPingPong.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SelfFundedPingPong.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SelfFundedPingPongOwnershipTransferRequested)
				if err := _SelfFundedPingPong.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) ParseOwnershipTransferRequested(log types.Log) (*SelfFundedPingPongOwnershipTransferRequested, error) {
	event := new(SelfFundedPingPongOwnershipTransferRequested)
	if err := _SelfFundedPingPong.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SelfFundedPingPongOwnershipTransferredIterator struct {
	Event *SelfFundedPingPongOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SelfFundedPingPongOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SelfFundedPingPongOwnershipTransferred)
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
		it.Event = new(SelfFundedPingPongOwnershipTransferred)
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

func (it *SelfFundedPingPongOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *SelfFundedPingPongOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SelfFundedPingPongOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SelfFundedPingPongOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SelfFundedPingPong.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongOwnershipTransferredIterator{contract: _SelfFundedPingPong.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SelfFundedPingPong.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SelfFundedPingPongOwnershipTransferred)
				if err := _SelfFundedPingPong.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) ParseOwnershipTransferred(log types.Log) (*SelfFundedPingPongOwnershipTransferred, error) {
	event := new(SelfFundedPingPongOwnershipTransferred)
	if err := _SelfFundedPingPong.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SelfFundedPingPongPingIterator struct {
	Event *SelfFundedPingPongPing

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SelfFundedPingPongPingIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SelfFundedPingPongPing)
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
		it.Event = new(SelfFundedPingPongPing)
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

func (it *SelfFundedPingPongPingIterator) Error() error {
	return it.fail
}

func (it *SelfFundedPingPongPingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SelfFundedPingPongPing struct {
	PingPongCount *big.Int
	Raw           types.Log
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) FilterPing(opts *bind.FilterOpts) (*SelfFundedPingPongPingIterator, error) {

	logs, sub, err := _SelfFundedPingPong.contract.FilterLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongPingIterator{contract: _SelfFundedPingPong.contract, event: "Ping", logs: logs, sub: sub}, nil
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) WatchPing(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongPing) (event.Subscription, error) {

	logs, sub, err := _SelfFundedPingPong.contract.WatchLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SelfFundedPingPongPing)
				if err := _SelfFundedPingPong.contract.UnpackLog(event, "Ping", log); err != nil {
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

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) ParsePing(log types.Log) (*SelfFundedPingPongPing, error) {
	event := new(SelfFundedPingPongPing)
	if err := _SelfFundedPingPong.contract.UnpackLog(event, "Ping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SelfFundedPingPongPongIterator struct {
	Event *SelfFundedPingPongPong

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SelfFundedPingPongPongIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SelfFundedPingPongPong)
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
		it.Event = new(SelfFundedPingPongPong)
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

func (it *SelfFundedPingPongPongIterator) Error() error {
	return it.fail
}

func (it *SelfFundedPingPongPongIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SelfFundedPingPongPong struct {
	PingPongCount *big.Int
	Raw           types.Log
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) FilterPong(opts *bind.FilterOpts) (*SelfFundedPingPongPongIterator, error) {

	logs, sub, err := _SelfFundedPingPong.contract.FilterLogs(opts, "Pong")
	if err != nil {
		return nil, err
	}
	return &SelfFundedPingPongPongIterator{contract: _SelfFundedPingPong.contract, event: "Pong", logs: logs, sub: sub}, nil
}

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) WatchPong(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongPong) (event.Subscription, error) {

	logs, sub, err := _SelfFundedPingPong.contract.WatchLogs(opts, "Pong")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SelfFundedPingPongPong)
				if err := _SelfFundedPingPong.contract.UnpackLog(event, "Pong", log); err != nil {
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

func (_SelfFundedPingPong *SelfFundedPingPongFilterer) ParsePong(log types.Log) (*SelfFundedPingPongPong, error) {
	event := new(SelfFundedPingPongPong)
	if err := _SelfFundedPingPong.contract.UnpackLog(event, "Pong", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_SelfFundedPingPong *SelfFundedPingPong) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _SelfFundedPingPong.abi.Events["Funded"].ID:
		return _SelfFundedPingPong.ParseFunded(log)
	case _SelfFundedPingPong.abi.Events["OwnershipTransferRequested"].ID:
		return _SelfFundedPingPong.ParseOwnershipTransferRequested(log)
	case _SelfFundedPingPong.abi.Events["OwnershipTransferred"].ID:
		return _SelfFundedPingPong.ParseOwnershipTransferred(log)
	case _SelfFundedPingPong.abi.Events["Ping"].ID:
		return _SelfFundedPingPong.ParsePing(log)
	case _SelfFundedPingPong.abi.Events["Pong"].ID:
		return _SelfFundedPingPong.ParsePong(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (SelfFundedPingPongFunded) Topic() common.Hash {
	return common.HexToHash("0x302777af5d26fab9dd5120c5f1307c65193ebc51daf33244ada4365fab10602c")
}

func (SelfFundedPingPongOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (SelfFundedPingPongOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (SelfFundedPingPongPing) Topic() common.Hash {
	return common.HexToHash("0x48257dc961b6f792c2b78a080dacfed693b660960a702de21cee364e20270e2f")
}

func (SelfFundedPingPongPong) Topic() common.Hash {
	return common.HexToHash("0x58b69f57828e6962d216502094c54f6562f3bf082ba758966c3454f9e37b1525")
}

func (_SelfFundedPingPong *SelfFundedPingPong) Address() common.Address {
	return _SelfFundedPingPong.address
}

type SelfFundedPingPongInterface interface {
	GetCountIncrBeforeFunding(opts *bind.CallOpts) (uint8, error)

	GetCounterpartAddress(opts *bind.CallOpts) (common.Address, error)

	GetCounterpartChainSelector(opts *bind.CallOpts) (uint64, error)

	GetFeeToken(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IsPaused(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	FundPingPong(opts *bind.TransactOpts, pingPongCount *big.Int) (*types.Transaction, error)

	SetCounterpart(opts *bind.TransactOpts, counterpartChainSelector uint64, counterpartAddress common.Address) (*types.Transaction, error)

	SetCounterpartAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error)

	SetCounterpartChainSelector(opts *bind.TransactOpts, chainSelector uint64) (*types.Transaction, error)

	SetPaused(opts *bind.TransactOpts, pause bool) (*types.Transaction, error)

	StartPingPong(opts *bind.TransactOpts) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterFunded(opts *bind.FilterOpts) (*SelfFundedPingPongFundedIterator, error)

	WatchFunded(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongFunded) (event.Subscription, error)

	ParseFunded(log types.Log) (*SelfFundedPingPongFunded, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SelfFundedPingPongOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*SelfFundedPingPongOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SelfFundedPingPongOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*SelfFundedPingPongOwnershipTransferred, error)

	FilterPing(opts *bind.FilterOpts) (*SelfFundedPingPongPingIterator, error)

	WatchPing(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongPing) (event.Subscription, error)

	ParsePing(log types.Log) (*SelfFundedPingPongPing, error)

	FilterPong(opts *bind.FilterOpts) (*SelfFundedPingPongPongIterator, error)

	WatchPong(opts *bind.WatchOpts, sink chan<- *SelfFundedPingPongPong) (event.Subscription, error)

	ParsePong(log types.Log) (*SelfFundedPingPongPong, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
