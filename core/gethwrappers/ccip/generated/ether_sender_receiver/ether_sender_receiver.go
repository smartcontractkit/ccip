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
	Bin: "0x60c06040523480156200001157600080fd5b5060405162001c4a38038062001c4a833981016040819052620000349162000169565b806001600160a01b03811662000064576040516335fdcccd60e21b81526000600482015260240160405180910390fd5b806001600160a01b03166080816001600160a01b03168152505050806001600160a01b031663e861e9076040518163ffffffff1660e01b8152600401602060405180830381865afa158015620000be573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000e4919062000169565b6001600160a01b0390811660a081905260405163095ea7b360e01b8152918316600483015260001960248301529063095ea7b3906044016020604051808303816000875af11580156200013b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200016191906200019b565b5050620001bf565b6000602082840312156200017c57600080fd5b81516001600160a01b03811681146200019457600080fd5b9392505050565b600060208284031215620001ae57600080fd5b815180151581146200019457600080fd5b60805160a051611a11620002396000396000818161016d01528181610425015281816105b601528181610a8a01528181610b3201528181610c270152610cef0152600081816101f80152818161030101528181610399015281816104cd0152818161062d0152818161072101526108570152611a116000f3fe60806040526004361061007f5760003560e01c80634dbe7e921161004e5780634dbe7e921461015b57806385572ffb146101b457806396f4e9f9146101d6578063b0f479a1146101e957600080fd5b806301ffc9a71461008b578063181f5a77146100c05780631e722e011461011657806320487ded1461013b57600080fd5b3661008657005b600080fd5b34801561009757600080fd5b506100ab6100a6366004611155565b61021c565b60405190151581526020015b60405180910390f35b3480156100cc57600080fd5b506101096040518060400160405280601d81526020017f457468657253656e646572526563656976657220312e302e302d64657600000081525081565b6040516100b7919061120c565b34801561012257600080fd5b5061012d62030d4081565b6040519081526020016100b7565b34801561014757600080fd5b5061012d610156366004611254565b6102b5565b34801561016757600080fd5b5061018f7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100b7565b3480156101c057600080fd5b506101d46101cf3660046112a2565b610381565b005b61012d6101e4366004611254565b61040b565b3480156101f557600080fd5b507f000000000000000000000000000000000000000000000000000000000000000061018f565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f85572ffb0000000000000000000000000000000000000000000000000000000014806102af57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b6000806102c1836108e5565b6040517f20487ded00000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906320487ded9061033890879085906004016112d7565b602060405180830381865afa158015610355573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061037991906113e9565b949350505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146103f7576040517fd7f733340000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61040861040382611634565b610aed565b50565b600061041682610d68565b6000610421836108e5565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db08260400151600081518110610476576104766116e1565b6020026020010151602001516040518263ffffffff1660e01b81526004016000604051808303818588803b1580156104ad57600080fd5b505af11580156104c1573d6000803e3d6000fd5b505050505060006104ef7f000000000000000000000000000000000000000000000000000000000000000090565b73ffffffffffffffffffffffffffffffffffffffff166320487ded86846040518363ffffffff1660e01b81526004016105299291906112d7565b602060405180830381865afa158015610546573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056a91906113e9565b606083015190915073ffffffffffffffffffffffffffffffffffffffff16156107a45760608201516105b49073ffffffffffffffffffffffffffffffffffffffff16333084610e62565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16826060015173ffffffffffffffffffffffffffffffffffffffff16146106e457606082015173ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f00000000000000000000000000000000000000000000000000000000000000006040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602481018490526044016020604051808303816000875af11580156106be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e29190611710565b505b6040517f96f4e9f900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906396f4e9f99061075890889086906004016112d7565b6020604051808303816000875af1158015610777573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079b91906113e9565b925050506102af565b600082604001516000815181106107bd576107bd6116e1565b602002602001015160200151346107d49190611732565b90508181101561081a576040517fa458261b00000000000000000000000000000000000000000000000000000000815260048101829052602481018390526044016103ee565b6040517f96f4e9f900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906396f4e9f9908390610890908a9088906004016112d7565b60206040518083038185885af11580156108ae573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906108d391906113e9565b93505050506102af565b505092915050565b61092d6040518060a00160405280606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b36600061093d608085018561176c565b90925090508015610a11577f97a657c9000000000000000000000000000000000000000000000000000000006109776004600084866117d8565b61098091611802565b7fffffffff000000000000000000000000000000000000000000000000000000001603610a115760006109b682600481866117d8565b8101906109c39190611848565b51905062030d40811015610a0f576040517f17d4de7e00000000000000000000000000000000000000000000000000000000815262030d406004820152602481018290526044016103ee565b505b6000610a1c8561188a565b9050806040015151600114610a66578060400151516040517f83b9f0ae0000000000000000000000000000000000000000000000000000000081526004016103ee91815260200190565b604080513360208201520160405160208183030381529060405281602001819052507f00000000000000000000000000000000000000000000000000000000000000008160400151600081518110610ac057610ac06116e1565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff9091169052949350505050565b60008160600151806020019051810190610b079190611946565b905060008260800151600081518110610b2257610b226116e1565b60200260200101516020015190507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d826040518263ffffffff1660e01b8152600401610b8b91815260200190565b600060405180830381600087803b158015610ba557600080fd5b505af1158015610bb9573d6000803e3d6000fd5b5050505060008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114610c17576040519150601f19603f3d011682016040523d82523d6000602084013e610c1c565b606091505b5050905080610d62577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0836040518263ffffffff1660e01b81526004016000604051808303818588803b158015610c8d57600080fd5b505af1158015610ca1573d6000803e3d6000fd5b50506040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152602482018790527f000000000000000000000000000000000000000000000000000000000000000016935063a9059cbb925060440190506020604051808303816000875af1158015610d3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d609190611710565b505b50505050565b6000610d776040830183611963565b6000818110610d8857610d886116e1565b905060400201602001359050600073ffffffffffffffffffffffffffffffffffffffff16826060016020810190610dbf91906119cb565b73ffffffffffffffffffffffffffffffffffffffff1603610e2057803411610e1c576040517f7cb769dc000000000000000000000000000000000000000000000000000000008152600481018290523460248201526044016103ee565b5050565b803414610e1c576040517fba2f7467000000000000000000000000000000000000000000000000000000008152600481018290523460248201526044016103ee565b6040805173ffffffffffffffffffffffffffffffffffffffff8581166024830152848116604483015260648083018590528351808403909101815260849092018352602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656490840152610d6292879291600091610f35918516908490610fe4565b805190915015610fdf5780806020019051810190610f539190611710565b610fdf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016103ee565b505050565b60606103798484600085856000808673ffffffffffffffffffffffffffffffffffffffff16858760405161101891906119e8565b60006040518083038185875af1925050503d8060008114611055576040519150601f19603f3d011682016040523d82523d6000602084013e61105a565b606091505b509150915061106b87838387611076565b979650505050505050565b6060831561110c5782516000036111055773ffffffffffffffffffffffffffffffffffffffff85163b611105576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103ee565b5081610379565b61037983838151156111215781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ee919061120c565b60006020828403121561116757600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461119757600080fd5b9392505050565b60005b838110156111b95781810151838201526020016111a1565b50506000910152565b600081518084526111da81602086016020860161119e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061119760208301846111c2565b803567ffffffffffffffff8116811461123757600080fd5b919050565b600060a0828403121561124e57600080fd5b50919050565b6000806040838503121561126757600080fd5b6112708361121f565b9150602083013567ffffffffffffffff81111561128c57600080fd5b6112988582860161123c565b9150509250929050565b6000602082840312156112b457600080fd5b813567ffffffffffffffff8111156112cb57600080fd5b6103798482850161123c565b6000604067ffffffffffffffff8516835260208181850152845160a08386015261130460e08601826111c2565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08087840301606088015261133f83836111c2565b88860151888203830160808a01528051808352908601945060009350908501905b8084101561139f578451805173ffffffffffffffffffffffffffffffffffffffff16835286015186830152938501936001939093019290860190611360565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a015295506113db81876111c2565b9a9950505050505050505050565b6000602082840312156113fb57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561145457611454611402565b60405290565b60405160a0810167ffffffffffffffff8111828210171561145457611454611402565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156114c4576114c4611402565b604052919050565b600082601f8301126114dd57600080fd5b813567ffffffffffffffff8111156114f7576114f7611402565b61152860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161147d565b81815284602083860101111561153d57600080fd5b816020850160208301376000918101602001919091529392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461040857600080fd5b80356112378161155a565b600082601f83011261159857600080fd5b8135602067ffffffffffffffff8211156115b4576115b4611402565b6115c2818360051b0161147d565b82815260069290921b840181019181810190868411156115e157600080fd5b8286015b8481101561162957604081890312156115fe5760008081fd5b611606611431565b81356116118161155a565b815281850135858201528352918301916040016115e5565b509695505050505050565b600060a0823603121561164657600080fd5b61164e61145a565b8235815261165e6020840161121f565b6020820152604083013567ffffffffffffffff8082111561167e57600080fd5b61168a368387016114cc565b604084015260608501359150808211156116a357600080fd5b6116af368387016114cc565b606084015260808501359150808211156116c857600080fd5b506116d536828601611587565b60808301525092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561172257600080fd5b8151801515811461119757600080fd5b818103818111156102af577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126117a157600080fd5b83018035915067ffffffffffffffff8211156117bc57600080fd5b6020019150368190038213156117d157600080fd5b9250929050565b600080858511156117e857600080fd5b838611156117f557600080fd5b5050820193919092039150565b7fffffffff0000000000000000000000000000000000000000000000000000000081358181169160048510156108dd5760049490940360031b84901b1690921692915050565b60006020828403121561185a57600080fd5b6040516020810181811067ffffffffffffffff8211171561187d5761187d611402565b6040529135825250919050565b600060a0823603121561189c57600080fd5b6118a461145a565b823567ffffffffffffffff808211156118bc57600080fd5b6118c8368387016114cc565b835260208501359150808211156118de57600080fd5b6118ea368387016114cc565b6020840152604085013591508082111561190357600080fd5b61190f36838701611587565b60408401526119206060860161157c565b6060840152608085013591508082111561193957600080fd5b506116d5368286016114cc565b60006020828403121561195857600080fd5b81516111978161155a565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261199857600080fd5b83018035915067ffffffffffffffff8211156119b357600080fd5b6020019150600681901b36038213156117d157600080fd5b6000602082840312156119dd57600080fd5b81356111978161155a565b600082516119fa81846020870161119e565b919091019291505056fea164736f6c6343000813000a",
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
