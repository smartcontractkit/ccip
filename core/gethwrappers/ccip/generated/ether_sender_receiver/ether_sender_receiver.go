// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ether_sender_receiver

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
	_ = abi.ConvertType
)

type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

type ClientEVM2AnyMessage struct {
	Receiver     []byte
	Data         []byte
	TokenAmounts []ClientEVMTokenAmount
	FeeToken     common.Address
	ExtraArgs    []byte
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

var EtherSenderReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gotLimit\",\"type\":\"uint256\"}],\"name\":\"GasLimitTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gotFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"InsufficientFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gotAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"InsufficientMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gotAmounts\",\"type\":\"uint256\"}],\"name\":\"InvalidTokenAmounts\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gotAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"TokenAmountNotEqualToMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MIN_MESSAGE_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_weth\",\"outputs\":[{\"internalType\":\"contractIWrappedNative\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162001ba438038062001ba4833981016040819052620000349162000169565b806001600160a01b03811662000064576040516335fdcccd60e21b81526000600482015260240160405180910390fd5b806001600160a01b03166080816001600160a01b03168152505050806001600160a01b031663e861e9076040518163ffffffff1660e01b8152600401602060405180830381865afa158015620000be573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000e4919062000169565b6001600160a01b0390811660a081905260405163095ea7b360e01b8152918316600483015260001960248301529063095ea7b3906044016020604051808303816000875af11580156200013b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200016191906200019b565b5050620001bf565b6000602082840312156200017c57600080fd5b81516001600160a01b03811681146200019457600080fd5b9392505050565b600060208284031215620001ae57600080fd5b815180151581146200019457600080fd5b60805160a05161196b620002396000396000818161016d01528181610425015281816105b601528181610a2401528181610acc01528181610bc10152610c890152600081816101f80152818161030101528181610399015281816104cd0152818161062d0152818161072101526107d8015261196b6000f3fe60806040526004361061007f5760003560e01c80634dbe7e921161004e5780634dbe7e921461015b57806385572ffb146101b457806396f4e9f9146101d6578063b0f479a1146101e957600080fd5b806301ffc9a71461008b578063181f5a77146100c05780631e722e011461011657806320487ded1461013b57600080fd5b3661008657005b600080fd5b34801561009757600080fd5b506100ab6100a63660046110ad565b61021c565b60405190151581526020015b60405180910390f35b3480156100cc57600080fd5b506101096040518060400160405280601d81526020017f457468657253656e646572526563656976657220312e302e302d64657600000081525081565b6040516100b79190611164565b34801561012257600080fd5b5061012d62030d4081565b6040519081526020016100b7565b34801561014757600080fd5b5061012d6101563660046111ac565b6102b5565b34801561016757600080fd5b5061018f7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100b7565b3480156101c057600080fd5b506101d46101cf3660046111fa565b610381565b005b61012d6101e43660046111ac565b61040b565b3480156101f557600080fd5b507f000000000000000000000000000000000000000000000000000000000000000061018f565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f85572ffb0000000000000000000000000000000000000000000000000000000014806102af57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b6000806102c18361087f565b6040517f20487ded00000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906320487ded90610338908790859060040161122f565b602060405180830381865afa158015610355573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103799190611341565b949350505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146103f7576040517fd7f733340000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6104086104038261158c565b610a87565b50565b600061041682610d02565b60006104218361087f565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0826040015160008151811061047657610476611639565b6020026020010151602001516040518263ffffffff1660e01b81526004016000604051808303818588803b1580156104ad57600080fd5b505af11580156104c1573d6000803e3d6000fd5b505050505060006104ef7f000000000000000000000000000000000000000000000000000000000000000090565b73ffffffffffffffffffffffffffffffffffffffff166320487ded86846040518363ffffffff1660e01b815260040161052992919061122f565b602060405180830381865afa158015610546573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056a9190611341565b606083015190915073ffffffffffffffffffffffffffffffffffffffff16156107a45760608201516105b49073ffffffffffffffffffffffffffffffffffffffff16333084610dba565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16826060015173ffffffffffffffffffffffffffffffffffffffff16146106e457606082015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000006040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602481018490526044016020604051808303816000875af11580156106be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e29190611668565b505b6040517f96f4e9f900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906396f4e9f990610758908890869060040161122f565b6020604051808303816000875af1158015610777573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079b9190611341565b925050506102af565b600082604001516000815181106107bd576107bd611639565b602002602001015160200151346107d4919061168a565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166396f4e9f98288866040518463ffffffff1660e01b815260040161083292919061122f565b60206040518083038185885af1158015610850573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906108759190611341565b9695505050505050565b6108c76040518060a00160405280606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b3660006108d760808501856116c4565b909250905080156109ab577f97a657c900000000000000000000000000000000000000000000000000000000610911600460008486611730565b61091a9161175a565b7fffffffff0000000000000000000000000000000000000000000000000000000016036109ab5760006109508260048186611730565b81019061095d91906117a2565b51905062030d408110156109a9576040517f17d4de7e00000000000000000000000000000000000000000000000000000000815262030d406004820152602481018290526044016103ee565b505b60006109b6856117e4565b9050806040015151600114610a00578060400151516040517f83b9f0ae0000000000000000000000000000000000000000000000000000000081526004016103ee91815260200190565b604080513360208201520160405160208183030381529060405281602001819052507f00000000000000000000000000000000000000000000000000000000000000008160400151600081518110610a5a57610a5a611639565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff9091169052949350505050565b60008160600151806020019051810190610aa191906118a0565b905060008260800151600081518110610abc57610abc611639565b60200260200101516020015190507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d826040518263ffffffff1660e01b8152600401610b2591815260200190565b600060405180830381600087803b158015610b3f57600080fd5b505af1158015610b53573d6000803e3d6000fd5b5050505060008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114610bb1576040519150601f19603f3d011682016040523d82523d6000602084013e610bb6565b606091505b5050905080610cfc577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0836040518263ffffffff1660e01b81526004016000604051808303818588803b158015610c2757600080fd5b505af1158015610c3b573d6000803e3d6000fd5b50506040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152602482018790527f000000000000000000000000000000000000000000000000000000000000000016935063a9059cbb925060440190506020604051808303816000875af1158015610cd6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cfa9190611668565b505b50505050565b6000610d1160408301836118bd565b6000818110610d2257610d22611639565b905060400201602001359050600073ffffffffffffffffffffffffffffffffffffffff16826060016020810190610d599190611925565b73ffffffffffffffffffffffffffffffffffffffff1614610db657803414610db6576040517fba2f7467000000000000000000000000000000000000000000000000000000008152600481018290523460248201526044016103ee565b5050565b6040805173ffffffffffffffffffffffffffffffffffffffff8581166024830152848116604483015260648083018590528351808403909101815260849092018352602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656490840152610cfc92879291600091610e8d918516908490610f3c565b805190915015610f375780806020019051810190610eab9190611668565b610f37576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016103ee565b505050565b60606103798484600085856000808673ffffffffffffffffffffffffffffffffffffffff168587604051610f709190611942565b60006040518083038185875af1925050503d8060008114610fad576040519150601f19603f3d011682016040523d82523d6000602084013e610fb2565b606091505b5091509150610fc387838387610fce565b979650505050505050565b6060831561106457825160000361105d5773ffffffffffffffffffffffffffffffffffffffff85163b61105d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103ee565b5081610379565b61037983838151156110795781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ee9190611164565b6000602082840312156110bf57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146110ef57600080fd5b9392505050565b60005b838110156111115781810151838201526020016110f9565b50506000910152565b600081518084526111328160208601602086016110f6565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006110ef602083018461111a565b803567ffffffffffffffff8116811461118f57600080fd5b919050565b600060a082840312156111a657600080fd5b50919050565b600080604083850312156111bf57600080fd5b6111c883611177565b9150602083013567ffffffffffffffff8111156111e457600080fd5b6111f085828601611194565b9150509250929050565b60006020828403121561120c57600080fd5b813567ffffffffffffffff81111561122357600080fd5b61037984828501611194565b6000604067ffffffffffffffff8516835260208181850152845160a08386015261125c60e086018261111a565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080878403016060880152611297838361111a565b88860151888203830160808a01528051808352908601945060009350908501905b808410156112f7578451805173ffffffffffffffffffffffffffffffffffffffff168352860151868301529385019360019390930192908601906112b8565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a01529550611333818761111a565b9a9950505050505050505050565b60006020828403121561135357600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156113ac576113ac61135a565b60405290565b60405160a0810167ffffffffffffffff811182821017156113ac576113ac61135a565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561141c5761141c61135a565b604052919050565b600082601f83011261143557600080fd5b813567ffffffffffffffff81111561144f5761144f61135a565b61148060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016113d5565b81815284602083860101111561149557600080fd5b816020850160208301376000918101602001919091529392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461040857600080fd5b803561118f816114b2565b600082601f8301126114f057600080fd5b8135602067ffffffffffffffff82111561150c5761150c61135a565b61151a818360051b016113d5565b82815260069290921b8401810191818101908684111561153957600080fd5b8286015b8481101561158157604081890312156115565760008081fd5b61155e611389565b8135611569816114b2565b8152818501358582015283529183019160400161153d565b509695505050505050565b600060a0823603121561159e57600080fd5b6115a66113b2565b823581526115b660208401611177565b6020820152604083013567ffffffffffffffff808211156115d657600080fd5b6115e236838701611424565b604084015260608501359150808211156115fb57600080fd5b61160736838701611424565b6060840152608085013591508082111561162057600080fd5b5061162d368286016114df565b60808301525092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561167a57600080fd5b815180151581146110ef57600080fd5b818103818111156102af577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126116f957600080fd5b83018035915067ffffffffffffffff82111561171457600080fd5b60200191503681900382131561172957600080fd5b9250929050565b6000808585111561174057600080fd5b8386111561174d57600080fd5b5050820193919092039150565b7fffffffff00000000000000000000000000000000000000000000000000000000813581811691600485101561179a5780818660040360031b1b83161692505b505092915050565b6000602082840312156117b457600080fd5b6040516020810181811067ffffffffffffffff821117156117d7576117d761135a565b6040529135825250919050565b600060a082360312156117f657600080fd5b6117fe6113b2565b823567ffffffffffffffff8082111561181657600080fd5b61182236838701611424565b8352602085013591508082111561183857600080fd5b61184436838701611424565b6020840152604085013591508082111561185d57600080fd5b611869368387016114df565b604084015261187a606086016114d4565b6060840152608085013591508082111561189357600080fd5b5061162d36828601611424565b6000602082840312156118b257600080fd5b81516110ef816114b2565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126118f257600080fd5b83018035915067ffffffffffffffff82111561190d57600080fd5b6020019150600681901b360382131561172957600080fd5b60006020828403121561193757600080fd5b81356110ef816114b2565b600082516119548184602087016110f6565b919091019291505056fea164736f6c6343000813000a",
}

var EtherSenderReceiverABI = EtherSenderReceiverMetaData.ABI

var EtherSenderReceiverBin = EtherSenderReceiverMetaData.Bin

func DeployEtherSenderReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address) (common.Address, *types.Transaction, *EtherSenderReceiver, error) {
	parsed, err := EtherSenderReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EtherSenderReceiverBin), backend, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtherSenderReceiver{address: address, abi: *parsed, EtherSenderReceiverCaller: EtherSenderReceiverCaller{contract: contract}, EtherSenderReceiverTransactor: EtherSenderReceiverTransactor{contract: contract}, EtherSenderReceiverFilterer: EtherSenderReceiverFilterer{contract: contract}}, nil
}

type EtherSenderReceiver struct {
	address common.Address
	abi     abi.ABI
	EtherSenderReceiverCaller
	EtherSenderReceiverTransactor
	EtherSenderReceiverFilterer
}

type EtherSenderReceiverCaller struct {
	contract *bind.BoundContract
}

type EtherSenderReceiverTransactor struct {
	contract *bind.BoundContract
}

type EtherSenderReceiverFilterer struct {
	contract *bind.BoundContract
}

type EtherSenderReceiverSession struct {
	Contract     *EtherSenderReceiver
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EtherSenderReceiverCallerSession struct {
	Contract *EtherSenderReceiverCaller
	CallOpts bind.CallOpts
}

type EtherSenderReceiverTransactorSession struct {
	Contract     *EtherSenderReceiverTransactor
	TransactOpts bind.TransactOpts
}

type EtherSenderReceiverRaw struct {
	Contract *EtherSenderReceiver
}

type EtherSenderReceiverCallerRaw struct {
	Contract *EtherSenderReceiverCaller
}

type EtherSenderReceiverTransactorRaw struct {
	Contract *EtherSenderReceiverTransactor
}

func NewEtherSenderReceiver(address common.Address, backend bind.ContractBackend) (*EtherSenderReceiver, error) {
	abi, err := abi.JSON(strings.NewReader(EtherSenderReceiverABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEtherSenderReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherSenderReceiver{address: address, abi: abi, EtherSenderReceiverCaller: EtherSenderReceiverCaller{contract: contract}, EtherSenderReceiverTransactor: EtherSenderReceiverTransactor{contract: contract}, EtherSenderReceiverFilterer: EtherSenderReceiverFilterer{contract: contract}}, nil
}

func NewEtherSenderReceiverCaller(address common.Address, caller bind.ContractCaller) (*EtherSenderReceiverCaller, error) {
	contract, err := bindEtherSenderReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherSenderReceiverCaller{contract: contract}, nil
}

func NewEtherSenderReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherSenderReceiverTransactor, error) {
	contract, err := bindEtherSenderReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherSenderReceiverTransactor{contract: contract}, nil
}

func NewEtherSenderReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherSenderReceiverFilterer, error) {
	contract, err := bindEtherSenderReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherSenderReceiverFilterer{contract: contract}, nil
}

func bindEtherSenderReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherSenderReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_EtherSenderReceiver *EtherSenderReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherSenderReceiver.Contract.EtherSenderReceiverCaller.contract.Call(opts, result, method, params...)
}

func (_EtherSenderReceiver *EtherSenderReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.EtherSenderReceiverTransactor.contract.Transfer(opts)
}

func (_EtherSenderReceiver *EtherSenderReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.EtherSenderReceiverTransactor.contract.Transact(opts, method, params...)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherSenderReceiver.Contract.contract.Call(opts, result, method, params...)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.contract.Transfer(opts)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.contract.Transact(opts, method, params...)
}

func (_EtherSenderReceiver *EtherSenderReceiverCaller) MINMESSAGEGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "MIN_MESSAGE_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) MINMESSAGEGASLIMIT() (*big.Int, error) {
	return _EtherSenderReceiver.Contract.MINMESSAGEGASLIMIT(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) MINMESSAGEGASLIMIT() (*big.Int, error) {
	return _EtherSenderReceiver.Contract.MINMESSAGEGASLIMIT(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverCaller) GetFee(opts *bind.CallOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "getFee", destinationChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) GetFee(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _EtherSenderReceiver.Contract.GetFee(&_EtherSenderReceiver.CallOpts, destinationChainSelector, message)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) GetFee(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _EtherSenderReceiver.Contract.GetFee(&_EtherSenderReceiver.CallOpts, destinationChainSelector, message)
}

func (_EtherSenderReceiver *EtherSenderReceiverCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) GetRouter() (common.Address, error) {
	return _EtherSenderReceiver.Contract.GetRouter(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) GetRouter() (common.Address, error) {
	return _EtherSenderReceiver.Contract.GetRouter(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverCaller) IWeth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "i_weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) IWeth() (common.Address, error) {
	return _EtherSenderReceiver.Contract.IWeth(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) IWeth() (common.Address, error) {
	return _EtherSenderReceiver.Contract.IWeth(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EtherSenderReceiver.Contract.SupportsInterface(&_EtherSenderReceiver.CallOpts, interfaceId)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EtherSenderReceiver.Contract.SupportsInterface(&_EtherSenderReceiver.CallOpts, interfaceId)
}

func (_EtherSenderReceiver *EtherSenderReceiverCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) TypeAndVersion() (string, error) {
	return _EtherSenderReceiver.Contract.TypeAndVersion(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) TypeAndVersion() (string, error) {
	return _EtherSenderReceiver.Contract.TypeAndVersion(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _EtherSenderReceiver.contract.Transact(opts, "ccipReceive", message)
}

func (_EtherSenderReceiver *EtherSenderReceiverSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.CcipReceive(&_EtherSenderReceiver.TransactOpts, message)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.CcipReceive(&_EtherSenderReceiver.TransactOpts, message)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactor) CcipSend(opts *bind.TransactOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _EtherSenderReceiver.contract.Transact(opts, "ccipSend", destinationChainSelector, message)
}

func (_EtherSenderReceiver *EtherSenderReceiverSession) CcipSend(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.CcipSend(&_EtherSenderReceiver.TransactOpts, destinationChainSelector, message)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactorSession) CcipSend(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.CcipSend(&_EtherSenderReceiver.TransactOpts, destinationChainSelector, message)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherSenderReceiver.contract.RawTransact(opts, nil)
}

func (_EtherSenderReceiver *EtherSenderReceiverSession) Receive() (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.Receive(&_EtherSenderReceiver.TransactOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverTransactorSession) Receive() (*types.Transaction, error) {
	return _EtherSenderReceiver.Contract.Receive(&_EtherSenderReceiver.TransactOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiver) Address() common.Address {
	return _EtherSenderReceiver.address
}

type EtherSenderReceiverInterface interface {
	MINMESSAGEGASLIMIT(opts *bind.CallOpts) (*big.Int, error)

	GetFee(opts *bind.CallOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IWeth(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	Address() common.Address
}
