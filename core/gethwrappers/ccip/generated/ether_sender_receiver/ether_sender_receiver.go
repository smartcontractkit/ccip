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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AllowanceTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCIPReceiveFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gotFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"InsufficientFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gotAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"InsufficientMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"destEOA\",\"type\":\"bytes\"}],\"name\":\"InvalidDestinationEOA\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"}],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gotAmounts\",\"type\":\"uint256\"}],\"name\":\"InvalidTokenAmounts\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"got\",\"type\":\"address\"}],\"name\":\"InvalidWethAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gotAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"TokenAmountNotEqualToMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MESSAGE_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_weth\",\"outputs\":[{\"internalType\":\"contractIWrappedNative\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162001b7138038062001b718339810160408190526200003491620000f7565b806001600160a01b03811662000064576040516335fdcccd60e21b81526000600482015260240160405180910390fd5b806001600160a01b03166080816001600160a01b03168152505050806001600160a01b031663e861e9076040518163ffffffff1660e01b8152600401602060405180830381865afa158015620000be573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000e49190620000f7565b6001600160a01b031660a0525062000129565b6000602082840312156200010a57600080fd5b81516001600160a01b03811681146200012257600080fd5b9392505050565b60805160a0516119ce620001a36000396000818161014b0152818161052c015281816105cf01528181610b5e01528181610bf20152610c4c0152600081816101d60152818161037e015281816104160152818161060e01528181610715015281816107dc0152818161085a015261099001526119ce6000f3fe60806040526004361061007b5760003560e01c806385572ffb1161004e57806385572ffb1461019257806396f4e9f9146101b4578063b0f479a1146101c7578063f9fc3c21146101fa57600080fd5b806301ffc9a714610080578063181f5a77146100b557806320487ded1461010b5780634dbe7e9214610139575b600080fd5b34801561008c57600080fd5b506100a061009b3660046112b9565b610211565b60405190151581526020015b60405180910390f35b3480156100c157600080fd5b506100fe6040518060400160405280601d81526020017f457468657253656e646572526563656976657220312e302e302d64657600000081525081565b6040516100ac9190611370565b34801561011757600080fd5b5061012b6101263660046115d2565b6102aa565b6040519081526020016100ac565b34801561014557600080fd5b5061016d7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ac565b34801561019e57600080fd5b506101b26101ad3660046116cf565b6103fe565b005b61012b6101c23660046115d2565b610488565b3480156101d357600080fd5b507f000000000000000000000000000000000000000000000000000000000000000061016d565b34801561020657600080fd5b5061012b6207a12081565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f85572ffb0000000000000000000000000000000000000000000000000000000014806102a457507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b60006102b582610a16565b600061033760405180602001604052806207a12081525060408051915160248084019190915281518084039091018152604490920190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f97a657c90000000000000000000000000000000000000000000000000000000017905290565b608084018190526040517f20487ded00000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906320487ded906103b5908790879060040161170a565b602060405180830381865afa1580156103d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f6919061181c565b949350505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610474576040517fd7f733340000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61048561048082611835565b610c2e565b50565b600061049382610a16565b61049d8234610dba565b600061051f60405180602001604052806207a12081525060408051915160248084019190915281518084039091018152604490920190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f97a657c90000000000000000000000000000000000000000000000000000000017905290565b90508083608001819052507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0846040015160008151811061057d5761057d6118e2565b6020026020010151602001516040518263ffffffff1660e01b81526004016000604051808303818588803b1580156105b457600080fd5b505af11580156105c8573d6000803e3d6000fd5b50505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663095ea7b36106307f000000000000000000000000000000000000000000000000000000000000000090565b8560400151600081518110610647576106476118e2565b6020026020010151602001516040518363ffffffff1660e01b815260040161069192919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b6020604051808303816000875af11580156106b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106d49190611911565b506040517f20487ded00000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906320487ded9061074c908890889060040161170a565b602060405180830381865afa158015610769573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061078d919061181c565b606085015190915073ffffffffffffffffffffffffffffffffffffffff16156108dd5760608401516107d79073ffffffffffffffffffffffffffffffffffffffff16333084610e62565b61081d7f0000000000000000000000000000000000000000000000000000000000000000606086015173ffffffffffffffffffffffffffffffffffffffff169083610f3e565b6040517f96f4e9f900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906396f4e9f990610891908890889060040161170a565b6020604051808303816000875af11580156108b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d4919061181c565b925050506102a4565b600084604001516000815181106108f6576108f66118e2565b6020026020010151602001513461090d9190611962565b905081811015610953576040517fa458261b000000000000000000000000000000000000000000000000000000008152600481018290526024810183905260440161046b565b6040517f96f4e9f900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906396f4e9f99083906109c9908a908a9060040161170a565b60206040518083038185885af11580156109e7573d6000803e3d6000fd5b50505050506040513d601f19601f82011682018060405250810190610a0c919061181c565b93505050506102a4565b60008160000151806020019051810190610a309190611975565b905073ffffffffffffffffffffffffffffffffffffffff8116610a825781516040517fa3c8cf0900000000000000000000000000000000000000000000000000000000815261046b9190600401611370565b60008260200151806020019051810190610a9c9190611975565b905073ffffffffffffffffffffffffffffffffffffffff8116610af15782602001516040517fd5f8432a00000000000000000000000000000000000000000000000000000000815260040161046b9190611370565b826040015151600114610b39578260400151516040517f83b9f0ae00000000000000000000000000000000000000000000000000000000815260040161046b91815260200190565b60008360400151600081518110610b5257610b526118e2565b602002602001015190507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610c285780516040517fca131cdb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081166004830152909116602482015260440161046b565b50505050565b60008160600151806020019051810190610c489190611975565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d8360800151600081518110610c9d57610c9d6118e2565b6020026020010151602001516040518263ffffffff1660e01b8152600401610cc791815260200190565b600060405180830381600087803b158015610ce157600080fd5b505af1158015610cf5573d6000803e3d6000fd5b5050505060008173ffffffffffffffffffffffffffffffffffffffff168360800151600081518110610d2957610d296118e2565b60200260200101516020015160405160006040518083038185875af1925050503d8060008114610d75576040519150601f19603f3d011682016040523d82523d6000602084013e610d7a565b606091505b5050905080610db5576040517f7e2d8dc300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b60008260400151600081518110610dd357610dd36118e2565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff16836060015173ffffffffffffffffffffffffffffffffffffffff1603610db55780602001518211610db55760208101516040517f7cb769dc00000000000000000000000000000000000000000000000000000000815260048101919091526024810183905260440161046b565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052610c289085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261103c565b6040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015610fb5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fd9919061181c565b610fe39190611992565b60405173ffffffffffffffffffffffffffffffffffffffff8516602482015260448101829052909150610c289085907f095ea7b30000000000000000000000000000000000000000000000000000000090606401610ebc565b600061109e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166111489092919063ffffffff16565b805190915015610db557808060200190518101906110bc9190611911565b610db5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161046b565b60606103f68484600085856000808673ffffffffffffffffffffffffffffffffffffffff16858760405161117c91906119a5565b60006040518083038185875af1925050503d80600081146111b9576040519150601f19603f3d011682016040523d82523d6000602084013e6111be565b606091505b50915091506111cf878383876111da565b979650505050505050565b606083156112705782516000036112695773ffffffffffffffffffffffffffffffffffffffff85163b611269576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161046b565b50816103f6565b6103f683838151156112855781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046b9190611370565b6000602082840312156112cb57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146112fb57600080fd5b9392505050565b60005b8381101561131d578181015183820152602001611305565b50506000910152565b6000815180845261133e816020860160208601611302565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006112fb6020830184611326565b803567ffffffffffffffff8116811461139b57600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156113f2576113f26113a0565b60405290565b60405160a0810167ffffffffffffffff811182821017156113f2576113f26113a0565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611462576114626113a0565b604052919050565b600082601f83011261147b57600080fd5b813567ffffffffffffffff811115611495576114956113a0565b6114c660207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161141b565b8181528460208386010111156114db57600080fd5b816020850160208301376000918101602001919091529392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461048557600080fd5b803561139b816114f8565b600082601f83011261153657600080fd5b8135602067ffffffffffffffff821115611552576115526113a0565b611560818360051b0161141b565b82815260069290921b8401810191818101908684111561157f57600080fd5b8286015b848110156115c7576040818903121561159c5760008081fd5b6115a46113cf565b81356115af816114f8565b81528185013585820152835291830191604001611583565b509695505050505050565b600080604083850312156115e557600080fd5b6115ee83611383565b9150602083013567ffffffffffffffff8082111561160b57600080fd5b9084019060a0828703121561161f57600080fd5b6116276113f8565b82358281111561163657600080fd5b6116428882860161146a565b82525060208301358281111561165757600080fd5b6116638882860161146a565b60208301525060408301358281111561167b57600080fd5b61168788828601611525565b6040830152506116996060840161151a565b60608201526080830135828111156116b057600080fd5b6116bc8882860161146a565b6080830152508093505050509250929050565b6000602082840312156116e157600080fd5b813567ffffffffffffffff8111156116f857600080fd5b820160a081850312156112fb57600080fd5b6000604067ffffffffffffffff8516835260208181850152845160a08386015261173760e0860182611326565b9050818601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808784030160608801526117728383611326565b88860151888203830160808a01528051808352908601945060009350908501905b808410156117d2578451805173ffffffffffffffffffffffffffffffffffffffff16835286015186830152938501936001939093019290860190611793565b50606089015173ffffffffffffffffffffffffffffffffffffffff1660a08901526080890151888203830160c08a0152955061180e8187611326565b9a9950505050505050505050565b60006020828403121561182e57600080fd5b5051919050565b600060a0823603121561184757600080fd5b61184f6113f8565b8235815261185f60208401611383565b6020820152604083013567ffffffffffffffff8082111561187f57600080fd5b61188b3683870161146a565b604084015260608501359150808211156118a457600080fd5b6118b03683870161146a565b606084015260808501359150808211156118c957600080fd5b506118d636828601611525565b60808301525092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561192357600080fd5b815180151581146112fb57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156102a4576102a4611933565b60006020828403121561198757600080fd5b81516112fb816114f8565b808201808211156102a4576102a4611933565b600082516119b7818460208701611302565b919091019291505056fea164736f6c6343000813000a",
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

func (_EtherSenderReceiver *EtherSenderReceiverCaller) MESSAGEGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EtherSenderReceiver.contract.Call(opts, &out, "MESSAGE_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EtherSenderReceiver *EtherSenderReceiverSession) MESSAGEGASLIMIT() (*big.Int, error) {
	return _EtherSenderReceiver.Contract.MESSAGEGASLIMIT(&_EtherSenderReceiver.CallOpts)
}

func (_EtherSenderReceiver *EtherSenderReceiverCallerSession) MESSAGEGASLIMIT() (*big.Int, error) {
	return _EtherSenderReceiver.Contract.MESSAGEGASLIMIT(&_EtherSenderReceiver.CallOpts)
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

func (_EtherSenderReceiver *EtherSenderReceiver) Address() common.Address {
	return _EtherSenderReceiver.address
}

type EtherSenderReceiverInterface interface {
	MESSAGEGASLIMIT(opts *bind.CallOpts) (*big.Int, error)

	GetFee(opts *bind.CallOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	IWeth(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error)

	Address() common.Address
}
