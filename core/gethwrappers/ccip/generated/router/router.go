// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package router

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
	zkSyncAccounts "github.com/zksync-sdk/zksync2-go/accounts"
	zkSyncClient "github.com/zksync-sdk/zksync2-go/clients"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
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

type RouterOffRamp struct {
	SourceChainSelector uint64
	OffRamp             common.Address
}

type RouterOnRamp struct {
	DestChainSelector uint64
	OnRamp            common.Address
}

var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedNative\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadARMSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFeeTokenAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"InvalidRecipientAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOffRamp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"calldataHash\",\"type\":\"bytes32\"}],\"name\":\"MessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_RET_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"internalType\":\"structRouter.OnRamp[]\",\"name\":\"onRampUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"internalType\":\"structRouter.OffRamp[]\",\"name\":\"offRampRemoves\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"internalType\":\"structRouter.OffRamp[]\",\"name\":\"offRampAdds\",\"type\":\"tuple[]\"}],\"name\":\"applyRampUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getArmProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRamps\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"internalType\":\"structRouter.OffRamp[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedNative\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"routeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"retData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedNative\",\"type\":\"address\"}],\"name\":\"setWrappedNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162002d2838038062002d288339810160408190526200003491620001af565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000e7565b5050600280546001600160a01b0319166001600160a01b039485161790555016608052620001e7565b336001600160a01b03821603620001415760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b0381168114620001aa57600080fd5b919050565b60008060408385031215620001c357600080fd5b620001ce8362000192565b9150620001de6020840162000192565b90509250929050565b608051612b1762000211600039600081816101f9015281816105e10152610af20152612b176000f3fe6080604052600436106101295760003560e01c80638da5cb5b116100a5578063a8d87a3b11610074578063e861e90711610059578063e861e90714610409578063f2fde38b14610434578063fbca3b741461045457600080fd5b8063a8d87a3b1461039c578063da5fcac8146103e957600080fd5b80638da5cb5b146102ed57806396f4e9f914610318578063a40e69c71461032b578063a48a90581461034d57600080fd5b806352cb60ca116100fc578063787350e3116100e1578063787350e31461028057806379ba5097146102a857806383826b2b146102bd57600080fd5b806352cb60ca1461023e5780635f3e849f1461026057600080fd5b8063181f5a771461012e57806320487ded1461018d5780633cf97983146101bb5780635246492f146101ea575b600080fd5b34801561013a57600080fd5b506101776040518060400160405280600c81526020017f526f7574657220312e322e30000000000000000000000000000000000000000081525081565b6040516101849190611f3c565b60405180910390f35b34801561019957600080fd5b506101ad6101a83660046121ad565b610481565b604051908152602001610184565b3480156101c757600080fd5b506101db6101d63660046122aa565b6105d9565b60405161018493929190612322565b3480156101f657600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610184565b34801561024a57600080fd5b5061025e61025936600461234d565b610836565b005b34801561026c57600080fd5b5061025e61027b36600461236a565b610885565b34801561028c57600080fd5b50610295608481565b60405161ffff9091168152602001610184565b3480156102b457600080fd5b5061025e6109d3565b3480156102c957600080fd5b506102dd6102d83660046123ab565b610ad0565b6040519015158152602001610184565b3480156102f957600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff16610219565b6101ad6103263660046121ad565b610aee565b34801561033757600080fd5b50610340611087565b60405161018491906123e2565b34801561035957600080fd5b506102dd610368366004612451565b67ffffffffffffffff1660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16151590565b3480156103a857600080fd5b506102196103b7366004612451565b67ffffffffffffffff1660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b3480156103f557600080fd5b5061025e6104043660046124b8565b61118b565b34801561041557600080fd5b5060025473ffffffffffffffffffffffffffffffffffffffff16610219565b34801561044057600080fd5b5061025e61044f36600461234d565b611490565b34801561046057600080fd5b5061047461046f366004612451565b6114a4565b6040516101849190612552565b606081015160009073ffffffffffffffffffffffffffffffffffffffff166104c25760025473ffffffffffffffffffffffffffffffffffffffff1660608301525b67ffffffffffffffff831660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff168061053a576040517fae236d9c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff851660048201526024015b60405180910390fd5b6040517f20487ded00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216906320487ded9061058e9087908790600401612689565b602060405180830381865afa1580156105ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105cf91906126ac565b9150505b92915050565b6000606060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663397796f76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561064a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066e91906126c5565b156106a5576040517fc148371500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106be6106b86040890160208a01612451565b33610ad0565b6106f4576040517fd2316ede00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006385572ffb60e01b8860405160240161070f91906127f4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152905061079c8186888a60846115c4565b919550935091507f9b877de93ea9895756e337442c657f95a34fc68e7eb988bdfa693d5be83016b688356107d660408b0160208c01612451565b83516020850120604051610823939291339193845267ffffffffffffffff92909216602084015273ffffffffffffffffffffffffffffffffffffffff166040830152606082015260800190565b60405180910390a1509450945094915050565b61083e6116ea565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b61088d6116ea565b73ffffffffffffffffffffffffffffffffffffffff82166108f2576040517f26a78f8f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610531565b73ffffffffffffffffffffffffffffffffffffffff83166109ad5760008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114610967576040519150601f19603f3d011682016040523d82523d6000602084013e61096c565b606091505b50509050806109a7576040517fe417b80b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050565b6109ce73ffffffffffffffffffffffffffffffffffffffff8416838361176d565b505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610a54576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610531565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6000610ae7610adf8484611841565b600490611885565b9392505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663397796f76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7f91906126c5565b15610bb6576040517fc148371500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff831660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1680610c29576040517fae236d9c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610531565b606083015160009073ffffffffffffffffffffffffffffffffffffffff16610dbb5760025473ffffffffffffffffffffffffffffffffffffffff90811660608601526040517f20487ded000000000000000000000000000000000000000000000000000000008152908316906320487ded90610cab9088908890600401612689565b602060405180830381865afa158015610cc8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cec91906126ac565b905080341015610d28576040517f07da6ee600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b349050836060015173ffffffffffffffffffffffffffffffffffffffff1663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b158015610d7757600080fd5b505af1158015610d8b573d6000803e3d6000fd5b505050506060850151610db6915073ffffffffffffffffffffffffffffffffffffffff16838361176d565b610eb2565b3415610df3576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f20487ded00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316906320487ded90610e479088908890600401612689565b602060405180830381865afa158015610e64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8891906126ac565b6060850151909150610eb29073ffffffffffffffffffffffffffffffffffffffff1633848461189d565b60005b846040015151811015610fe257600085604001518281518110610eda57610eda612900565b6020908102919091010151516040517f48a98aa400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8916600482015273ffffffffffffffffffffffffffffffffffffffff8083166024830152919250610fd9913391908716906348a98aa490604401602060405180830381865afa158015610f6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f90919061292f565b88604001518581518110610fa657610fa6612900565b6020026020010151602001518473ffffffffffffffffffffffffffffffffffffffff1661189d909392919063ffffffff16565b50600101610eb5565b506040517fdf0aa9e900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83169063df0aa9e99061103b90889088908690339060040161294c565b6020604051808303816000875af115801561105a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061107e91906126ac565b95945050505050565b6060600061109560046118fb565b90506000815167ffffffffffffffff8111156110b3576110b3611f6c565b6040519080825280602002602001820160405280156110f857816020015b60408051808201909152600080825260208201528152602001906001900390816110d15790505b50905060005b825181101561118457600083828151811061111b5761111b612900565b60200260200101519050604051806040016040528060a083901c67ffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681525083838151811061117057611170612900565b6020908102919091010152506001016110fe565b5092915050565b6111936116ea565b60005b8581101561126f5760008787838181106111b2576111b2612900565b9050604002018036038101906111c8919061299c565b60208181018051835167ffffffffffffffff90811660009081526003855260409081902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff948516179055855193519051921682529394509216917f1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f23910160405180910390a250600101611196565b5060005b838110156113a757600085858381811061128f5761128f612900565b6112a59260206040909202019081019150612451565b905060008686848181106112bb576112bb612900565b90506040020160200160208101906112d3919061234d565b90506112ea6112e28383611841565b600490611908565b611348576040517f4964779000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8316600482015273ffffffffffffffffffffffffffffffffffffffff82166024820152604401610531565b60405173ffffffffffffffffffffffffffffffffffffffff8216815267ffffffffffffffff8316907fa823809efda3ba66c873364eec120fa0923d9fabda73bc97dd5663341e2d9bcb9060200160405180910390a25050600101611273565b5060005b818110156114875760008383838181106113c7576113c7612900565b6113dd9260206040909202019081019150612451565b905060008484848181106113f3576113f3612900565b905060400201602001602081019061140b919061234d565b905061142261141a8383611841565b600490611914565b1561147d5760405173ffffffffffffffffffffffffffffffffffffffff8216815267ffffffffffffffff8316907fa4bdf64ebdf3316320601a081916a75aa144bcef6c4beeb0e9fb1982cacc6b949060200160405180910390a25b50506001016113ab565b50505050505050565b6114986116ea565b6114a181611920565b50565b60606114de8267ffffffffffffffff1660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16151590565b6114f8576040805160008082526020820190925290611184565b67ffffffffffffffff8216600081815260036020526040908190205490517ffbca3b74000000000000000000000000000000000000000000000000000000008152600481019290925273ffffffffffffffffffffffffffffffffffffffff169063fbca3b7490602401600060405180830381865afa15801561157e573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526105d391908101906129db565b6000606060008361ffff1667ffffffffffffffff8111156115e7576115e7611f6c565b6040519080825280601f01601f191660200182016040528015611611576020820181803683370190505b509150863b611644577f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b5a85811015611677577fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b85900360408104810387106116b0577f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b505a6000808a5160208c0160008c8cf193505a900390503d848111156116d35750835b808352806000602085013e50955095509592505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461176b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610531565b565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526109ce9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611a15565b6000610ae773ffffffffffffffffffffffffffffffffffffffff83167bffffffffffffffff000000000000000000000000000000000000000060a086901b16612a99565b60008181526001830160205260408120541515610ae7565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526109a79085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016117bf565b60606000610ae783611b21565b6000610ae78383611b7d565b6000610ae78383611c70565b3373ffffffffffffffffffffffffffffffffffffffff82160361199f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610531565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000611a77826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16611cbf9092919063ffffffff16565b8051909150156109ce5780806020019051810190611a9591906126c5565b6109ce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610531565b606081600001805480602002602001604051908101604052809291908181526020018280548015611b7157602002820191906000526020600020905b815481526020019060010190808311611b5d575b50505050509050919050565b60008181526001830160205260408120548015611c66576000611ba1600183612aac565b8554909150600090611bb590600190612aac565b9050808214611c1a576000866000018281548110611bd557611bd5612900565b9060005260206000200154905080876000018481548110611bf857611bf8612900565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080611c2b57611c2b612abf565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506105d3565b60009150506105d3565b6000818152600183016020526040812054611cb7575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556105d3565b5060006105d3565b6060611cce8484600085611cd6565b949350505050565b606082471015611d68576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610531565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611d919190612aee565b60006040518083038185875af1925050503d8060008114611dce576040519150601f19603f3d011682016040523d82523d6000602084013e611dd3565b606091505b5091509150611de487838387611def565b979650505050505050565b60608315611e85578251600003611e7e5773ffffffffffffffffffffffffffffffffffffffff85163b611e7e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610531565b5081611cce565b611cce8383815115611e9a5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105319190611f3c565b60005b83811015611ee9578181015183820152602001611ed1565b50506000910152565b60008151808452611f0a816020860160208601611ece565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610ae76020830184611ef2565b803567ffffffffffffffff81168114611f6757600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715611fbe57611fbe611f6c565b60405290565b60405160a0810167ffffffffffffffff81118282101715611fbe57611fbe611f6c565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561202e5761202e611f6c565b604052919050565b600082601f83011261204757600080fd5b813567ffffffffffffffff81111561206157612061611f6c565b61209260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611fe7565b8181528460208386010111156120a757600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff8211156120de576120de611f6c565b5060051b60200190565b73ffffffffffffffffffffffffffffffffffffffff811681146114a157600080fd5b8035611f67816120e8565b600082601f83011261212657600080fd5b8135602061213b612136836120c4565b611fe7565b82815260069290921b8401810191818101908684111561215a57600080fd5b8286015b848110156121a257604081890312156121775760008081fd5b61217f611f9b565b813561218a816120e8565b8152818501358582015283529183019160400161215e565b509695505050505050565b600080604083850312156121c057600080fd5b6121c983611f4f565b9150602083013567ffffffffffffffff808211156121e657600080fd5b9084019060a082870312156121fa57600080fd5b612202611fc4565b82358281111561221157600080fd5b61221d88828601612036565b82525060208301358281111561223257600080fd5b61223e88828601612036565b60208301525060408301358281111561225657600080fd5b61226288828601612115565b6040830152506122746060840161210a565b606082015260808301358281111561228b57600080fd5b61229788828601612036565b6080830152508093505050509250929050565b600080600080608085870312156122c057600080fd5b843567ffffffffffffffff8111156122d757600080fd5b850160a081880312156122e957600080fd5b9350602085013561ffff8116811461230057600080fd5b9250604085013591506060850135612317816120e8565b939692955090935050565b831515815260606020820152600061233d6060830185611ef2565b9050826040830152949350505050565b60006020828403121561235f57600080fd5b8135610ae7816120e8565b60008060006060848603121561237f57600080fd5b833561238a816120e8565b9250602084013561239a816120e8565b929592945050506040919091013590565b600080604083850312156123be57600080fd5b6123c783611f4f565b915060208301356123d7816120e8565b809150509250929050565b602080825282518282018190526000919060409081850190868401855b82811015612444578151805167ffffffffffffffff16855286015173ffffffffffffffffffffffffffffffffffffffff168685015292840192908501906001016123ff565b5091979650505050505050565b60006020828403121561246357600080fd5b610ae782611f4f565b60008083601f84011261247e57600080fd5b50813567ffffffffffffffff81111561249657600080fd5b6020830191508360208260061b85010111156124b157600080fd5b9250929050565b600080600080600080606087890312156124d157600080fd5b863567ffffffffffffffff808211156124e957600080fd5b6124f58a838b0161246c565b9098509650602089013591508082111561250e57600080fd5b61251a8a838b0161246c565b9096509450604089013591508082111561253357600080fd5b5061254089828a0161246c565b979a9699509497509295939492505050565b6020808252825182820181905260009190848201906040850190845b818110156125a057835173ffffffffffffffffffffffffffffffffffffffff168352928401929184019160010161256e565b50909695505050505050565b6000815160a084526125c160a0850182611ef2565b9050602080840151858303828701526125da8382611ef2565b60408681015188830389830152805180845290850195509092506000918401905b8083101561263a578551805173ffffffffffffffffffffffffffffffffffffffff168352850151858301529484019460019290920191908301906125fb565b5060608701519450612664606089018673ffffffffffffffffffffffffffffffffffffffff169052565b60808701519450878103608089015261267d8186611ef2565b98975050505050505050565b67ffffffffffffffff83168152604060208201526000611cce60408301846125ac565b6000602082840312156126be57600080fd5b5051919050565b6000602082840312156126d757600080fd5b81518015158114610ae757600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261271c57600080fd5b830160208101925035905067ffffffffffffffff81111561273c57600080fd5b8036038213156124b157600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8183526000602080850194508260005b858110156127e95781356127b7816120e8565b73ffffffffffffffffffffffffffffffffffffffff1687528183013583880152604096870196909101906001016127a4565b509495945050505050565b6020815281356020820152600061280d60208401611f4f565b67ffffffffffffffff808216604085015261282b60408601866126e7565b925060a0606086015261284260c08601848361274b565b92505061285260608601866126e7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08087860301608088015261288885838561274b565b9450608088013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18836030183126128c157600080fd5b602092880192830192359150838211156128da57600080fd5b8160061b36038313156128ec57600080fd5b8685030160a0870152611de4848284612794565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561294157600080fd5b8151610ae7816120e8565b67ffffffffffffffff8516815260806020820152600061296f60808301866125ac565b905083604083015273ffffffffffffffffffffffffffffffffffffffff8316606083015295945050505050565b6000604082840312156129ae57600080fd5b6129b6611f9b565b6129bf83611f4f565b815260208301356129cf816120e8565b60208201529392505050565b600060208083850312156129ee57600080fd5b825167ffffffffffffffff811115612a0557600080fd5b8301601f81018513612a1657600080fd5b8051612a24612136826120c4565b81815260059190911b82018301908381019087831115612a4357600080fd5b928401925b82841015611de4578351612a5b816120e8565b82529284019290840190612a48565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156105d3576105d3612a6a565b818103818111156105d3576105d3612a6a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008251612b00818460208701611ece565b919091019291505056fea164736f6c6343000818000a",
}

var RouterABI = RouterMetaData.ABI

var RouterBin = RouterMetaData.Bin

func DeployRouter(auth *bind.TransactOpts, backend bind.ContractBackend, wrappedNative common.Address, armProxy common.Address) (common.Address, *CustomTransaction, *Router, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	switch chainId.Uint64() {

	case 324, 280, 300:
		return DeployZkSyncRouter(auth, backend, wrappedNative, armProxy)
	}

	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RouterBin), backend, wrappedNative, armProxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, &CustomTransaction{Transaction: tx, CustomHash: tx.Hash()}, &Router{address: address, abi: *parsed, RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

type Router struct {
	address common.Address
	abi     abi.ABI
	RouterCaller
	RouterTransactor
	RouterFilterer
}

type RouterCaller struct {
	contract *bind.BoundContract
}

type RouterTransactor struct {
	contract *bind.BoundContract
}

type RouterFilterer struct {
	contract *bind.BoundContract
}

type RouterSession struct {
	Contract     *Router
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type RouterCallerSession struct {
	Contract *RouterCaller
	CallOpts bind.CallOpts
}

type RouterTransactorSession struct {
	Contract     *RouterTransactor
	TransactOpts bind.TransactOpts
}

type RouterRaw struct {
	Contract *Router
}

type RouterCallerRaw struct {
	Contract *RouterCaller
}

type RouterTransactorRaw struct {
	Contract *RouterTransactor
}

func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	abi, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{address: address, abi: abi, RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

func (_Router *RouterCaller) MAXRETBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "MAX_RET_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_Router *RouterSession) MAXRETBYTES() (uint16, error) {
	return _Router.Contract.MAXRETBYTES(&_Router.CallOpts)
}

func (_Router *RouterCallerSession) MAXRETBYTES() (uint16, error) {
	return _Router.Contract.MAXRETBYTES(&_Router.CallOpts)
}

func (_Router *RouterCaller) GetArmProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getArmProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Router *RouterSession) GetArmProxy() (common.Address, error) {
	return _Router.Contract.GetArmProxy(&_Router.CallOpts)
}

func (_Router *RouterCallerSession) GetArmProxy() (common.Address, error) {
	return _Router.Contract.GetArmProxy(&_Router.CallOpts)
}

func (_Router *RouterCaller) GetFee(opts *bind.CallOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getFee", destinationChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Router *RouterSession) GetFee(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _Router.Contract.GetFee(&_Router.CallOpts, destinationChainSelector, message)
}

func (_Router *RouterCallerSession) GetFee(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _Router.Contract.GetFee(&_Router.CallOpts, destinationChainSelector, message)
}

func (_Router *RouterCaller) GetOffRamps(opts *bind.CallOpts) ([]RouterOffRamp, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getOffRamps")

	if err != nil {
		return *new([]RouterOffRamp), err
	}

	out0 := *abi.ConvertType(out[0], new([]RouterOffRamp)).(*[]RouterOffRamp)

	return out0, err

}

func (_Router *RouterSession) GetOffRamps() ([]RouterOffRamp, error) {
	return _Router.Contract.GetOffRamps(&_Router.CallOpts)
}

func (_Router *RouterCallerSession) GetOffRamps() ([]RouterOffRamp, error) {
	return _Router.Contract.GetOffRamps(&_Router.CallOpts)
}

func (_Router *RouterCaller) GetOnRamp(opts *bind.CallOpts, destChainSelector uint64) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getOnRamp", destChainSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Router *RouterSession) GetOnRamp(destChainSelector uint64) (common.Address, error) {
	return _Router.Contract.GetOnRamp(&_Router.CallOpts, destChainSelector)
}

func (_Router *RouterCallerSession) GetOnRamp(destChainSelector uint64) (common.Address, error) {
	return _Router.Contract.GetOnRamp(&_Router.CallOpts, destChainSelector)
}

func (_Router *RouterCaller) GetSupportedTokens(opts *bind.CallOpts, chainSelector uint64) ([]common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getSupportedTokens", chainSelector)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Router *RouterSession) GetSupportedTokens(chainSelector uint64) ([]common.Address, error) {
	return _Router.Contract.GetSupportedTokens(&_Router.CallOpts, chainSelector)
}

func (_Router *RouterCallerSession) GetSupportedTokens(chainSelector uint64) ([]common.Address, error) {
	return _Router.Contract.GetSupportedTokens(&_Router.CallOpts, chainSelector)
}

func (_Router *RouterCaller) GetWrappedNative(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getWrappedNative")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Router *RouterSession) GetWrappedNative() (common.Address, error) {
	return _Router.Contract.GetWrappedNative(&_Router.CallOpts)
}

func (_Router *RouterCallerSession) GetWrappedNative() (common.Address, error) {
	return _Router.Contract.GetWrappedNative(&_Router.CallOpts)
}

func (_Router *RouterCaller) IsChainSupported(opts *bind.CallOpts, chainSelector uint64) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "isChainSupported", chainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Router *RouterSession) IsChainSupported(chainSelector uint64) (bool, error) {
	return _Router.Contract.IsChainSupported(&_Router.CallOpts, chainSelector)
}

func (_Router *RouterCallerSession) IsChainSupported(chainSelector uint64) (bool, error) {
	return _Router.Contract.IsChainSupported(&_Router.CallOpts, chainSelector)
}

func (_Router *RouterCaller) IsOffRamp(opts *bind.CallOpts, sourceChainSelector uint64, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "isOffRamp", sourceChainSelector, offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Router *RouterSession) IsOffRamp(sourceChainSelector uint64, offRamp common.Address) (bool, error) {
	return _Router.Contract.IsOffRamp(&_Router.CallOpts, sourceChainSelector, offRamp)
}

func (_Router *RouterCallerSession) IsOffRamp(sourceChainSelector uint64, offRamp common.Address) (bool, error) {
	return _Router.Contract.IsOffRamp(&_Router.CallOpts, sourceChainSelector, offRamp)
}

func (_Router *RouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Router *RouterSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

func (_Router *RouterCallerSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

func (_Router *RouterCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Router *RouterSession) TypeAndVersion() (string, error) {
	return _Router.Contract.TypeAndVersion(&_Router.CallOpts)
}

func (_Router *RouterCallerSession) TypeAndVersion() (string, error) {
	return _Router.Contract.TypeAndVersion(&_Router.CallOpts)
}

func (_Router *RouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "acceptOwnership")
}

func (_Router *RouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _Router.Contract.AcceptOwnership(&_Router.TransactOpts)
}

func (_Router *RouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Router.Contract.AcceptOwnership(&_Router.TransactOpts)
}

func (_Router *RouterTransactor) ApplyRampUpdates(opts *bind.TransactOpts, onRampUpdates []RouterOnRamp, offRampRemoves []RouterOffRamp, offRampAdds []RouterOffRamp) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "applyRampUpdates", onRampUpdates, offRampRemoves, offRampAdds)
}

func (_Router *RouterSession) ApplyRampUpdates(onRampUpdates []RouterOnRamp, offRampRemoves []RouterOffRamp, offRampAdds []RouterOffRamp) (*types.Transaction, error) {
	return _Router.Contract.ApplyRampUpdates(&_Router.TransactOpts, onRampUpdates, offRampRemoves, offRampAdds)
}

func (_Router *RouterTransactorSession) ApplyRampUpdates(onRampUpdates []RouterOnRamp, offRampRemoves []RouterOffRamp, offRampAdds []RouterOffRamp) (*types.Transaction, error) {
	return _Router.Contract.ApplyRampUpdates(&_Router.TransactOpts, onRampUpdates, offRampRemoves, offRampAdds)
}

func (_Router *RouterTransactor) CcipSend(opts *bind.TransactOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "ccipSend", destinationChainSelector, message)
}

func (_Router *RouterSession) CcipSend(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _Router.Contract.CcipSend(&_Router.TransactOpts, destinationChainSelector, message)
}

func (_Router *RouterTransactorSession) CcipSend(destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _Router.Contract.CcipSend(&_Router.TransactOpts, destinationChainSelector, message)
}

func (_Router *RouterTransactor) RecoverTokens(opts *bind.TransactOpts, tokenAddress common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "recoverTokens", tokenAddress, to, amount)
}

func (_Router *RouterSession) RecoverTokens(tokenAddress common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RecoverTokens(&_Router.TransactOpts, tokenAddress, to, amount)
}

func (_Router *RouterTransactorSession) RecoverTokens(tokenAddress common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RecoverTokens(&_Router.TransactOpts, tokenAddress, to, amount)
}

func (_Router *RouterTransactor) RouteMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage, gasForCallExactCheck uint16, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "routeMessage", message, gasForCallExactCheck, gasLimit, receiver)
}

func (_Router *RouterSession) RouteMessage(message ClientAny2EVMMessage, gasForCallExactCheck uint16, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.RouteMessage(&_Router.TransactOpts, message, gasForCallExactCheck, gasLimit, receiver)
}

func (_Router *RouterTransactorSession) RouteMessage(message ClientAny2EVMMessage, gasForCallExactCheck uint16, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.RouteMessage(&_Router.TransactOpts, message, gasForCallExactCheck, gasLimit, receiver)
}

func (_Router *RouterTransactor) SetWrappedNative(opts *bind.TransactOpts, wrappedNative common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setWrappedNative", wrappedNative)
}

func (_Router *RouterSession) SetWrappedNative(wrappedNative common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetWrappedNative(&_Router.TransactOpts, wrappedNative)
}

func (_Router *RouterTransactorSession) SetWrappedNative(wrappedNative common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetWrappedNative(&_Router.TransactOpts, wrappedNative)
}

func (_Router *RouterTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "transferOwnership", to)
}

func (_Router *RouterSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, to)
}

func (_Router *RouterTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, to)
}

type RouterMessageExecutedIterator struct {
	Event *RouterMessageExecuted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RouterMessageExecutedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterMessageExecuted)
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
		it.Event = new(RouterMessageExecuted)
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

func (it *RouterMessageExecutedIterator) Error() error {
	return it.fail
}

func (it *RouterMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RouterMessageExecuted struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	OffRamp             common.Address
	CalldataHash        [32]byte
	Raw                 types.Log
}

func (_Router *RouterFilterer) FilterMessageExecuted(opts *bind.FilterOpts) (*RouterMessageExecutedIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "MessageExecuted")
	if err != nil {
		return nil, err
	}
	return &RouterMessageExecutedIterator{contract: _Router.contract, event: "MessageExecuted", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *RouterMessageExecuted) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "MessageExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RouterMessageExecuted)
				if err := _Router.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
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

func (_Router *RouterFilterer) ParseMessageExecuted(log types.Log) (*RouterMessageExecuted, error) {
	event := new(RouterMessageExecuted)
	if err := _Router.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RouterOffRampAddedIterator struct {
	Event *RouterOffRampAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RouterOffRampAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOffRampAdded)
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
		it.Event = new(RouterOffRampAdded)
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

func (it *RouterOffRampAddedIterator) Error() error {
	return it.fail
}

func (it *RouterOffRampAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RouterOffRampAdded struct {
	SourceChainSelector uint64
	OffRamp             common.Address
	Raw                 types.Log
}

func (_Router *RouterFilterer) FilterOffRampAdded(opts *bind.FilterOpts, sourceChainSelector []uint64) (*RouterOffRampAddedIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OffRampAdded", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &RouterOffRampAddedIterator{contract: _Router.contract, event: "OffRampAdded", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *RouterOffRampAdded, sourceChainSelector []uint64) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OffRampAdded", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RouterOffRampAdded)
				if err := _Router.contract.UnpackLog(event, "OffRampAdded", log); err != nil {
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

func (_Router *RouterFilterer) ParseOffRampAdded(log types.Log) (*RouterOffRampAdded, error) {
	event := new(RouterOffRampAdded)
	if err := _Router.contract.UnpackLog(event, "OffRampAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RouterOffRampRemovedIterator struct {
	Event *RouterOffRampRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RouterOffRampRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOffRampRemoved)
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
		it.Event = new(RouterOffRampRemoved)
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

func (it *RouterOffRampRemovedIterator) Error() error {
	return it.fail
}

func (it *RouterOffRampRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RouterOffRampRemoved struct {
	SourceChainSelector uint64
	OffRamp             common.Address
	Raw                 types.Log
}

func (_Router *RouterFilterer) FilterOffRampRemoved(opts *bind.FilterOpts, sourceChainSelector []uint64) (*RouterOffRampRemovedIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OffRampRemoved", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &RouterOffRampRemovedIterator{contract: _Router.contract, event: "OffRampRemoved", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *RouterOffRampRemoved, sourceChainSelector []uint64) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OffRampRemoved", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RouterOffRampRemoved)
				if err := _Router.contract.UnpackLog(event, "OffRampRemoved", log); err != nil {
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

func (_Router *RouterFilterer) ParseOffRampRemoved(log types.Log) (*RouterOffRampRemoved, error) {
	event := new(RouterOffRampRemoved)
	if err := _Router.contract.UnpackLog(event, "OffRampRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RouterOnRampSetIterator struct {
	Event *RouterOnRampSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RouterOnRampSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOnRampSet)
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
		it.Event = new(RouterOnRampSet)
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

func (it *RouterOnRampSetIterator) Error() error {
	return it.fail
}

func (it *RouterOnRampSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RouterOnRampSet struct {
	DestChainSelector uint64
	OnRamp            common.Address
	Raw               types.Log
}

func (_Router *RouterFilterer) FilterOnRampSet(opts *bind.FilterOpts, destChainSelector []uint64) (*RouterOnRampSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OnRampSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &RouterOnRampSetIterator{contract: _Router.contract, event: "OnRampSet", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *RouterOnRampSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OnRampSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RouterOnRampSet)
				if err := _Router.contract.UnpackLog(event, "OnRampSet", log); err != nil {
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

func (_Router *RouterFilterer) ParseOnRampSet(log types.Log) (*RouterOnRampSet, error) {
	event := new(RouterOnRampSet)
	if err := _Router.contract.UnpackLog(event, "OnRampSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RouterOwnershipTransferRequestedIterator struct {
	Event *RouterOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RouterOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOwnershipTransferRequested)
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
		it.Event = new(RouterOwnershipTransferRequested)
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

func (it *RouterOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *RouterOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RouterOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Router *RouterFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RouterOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RouterOwnershipTransferRequestedIterator{contract: _Router.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RouterOwnershipTransferRequested)
				if err := _Router.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_Router *RouterFilterer) ParseOwnershipTransferRequested(log types.Log) (*RouterOwnershipTransferRequested, error) {
	event := new(RouterOwnershipTransferRequested)
	if err := _Router.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RouterOwnershipTransferredIterator struct {
	Event *RouterOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RouterOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOwnershipTransferred)
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
		it.Event = new(RouterOwnershipTransferred)
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

func (it *RouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *RouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RouterOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Router *RouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RouterOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RouterOwnershipTransferredIterator{contract: _Router.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RouterOwnershipTransferred)
				if err := _Router.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_Router *RouterFilterer) ParseOwnershipTransferred(log types.Log) (*RouterOwnershipTransferred, error) {
	event := new(RouterOwnershipTransferred)
	if err := _Router.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var RouterZkBin string = ("0x0004000000000002000d0000000000020000006003100270000003de0030019d000003de063001970003000000610355000200000001035500000001002001900000002d0000c13d0000008002000039000000400020043f000000040060008c000000530000413d000000000201043b000000e002200270000003e70020009c000000550000213d000003f40020009c000000920000a13d000003f50020009c0000030c0000a13d000003f60020009c000004b40000613d000003f70020009c0000047d0000613d000003f80020009c000000530000c13d000000440060008c000000530000413d0000000002000416000000000002004b000000530000c13d0000000402100370000000000202043b000004000020009c000000530000213d0000002401100370000000000101043b000003e10010009c000000530000213d000000a0022002100000041002200197000000000121019f0f720f250000040f000003020000013d0000000002000416000000000002004b000000530000c13d0000001f02600039000003df02200197000000a002200039000000400020043f0000001f0360018f000003e004600198000000a0024000390000003e0000613d000000a005000039000000000701034f000000007807043c0000000005850436000000000025004b0000003a0000c13d000000000003004b0000004b0000613d000000000141034f0000000303300210000000000402043300000000043401cf000000000434022f000000000101043b0000010003300089000000000131022f00000000013101cf000000000141019f0000000000120435000000400060008c000000530000413d000000a00100043d000003e10010009c000000530000213d000000c00200043d000003e10020009c000002ce0000a13d000000000100001900000f7400010430000003e80020009c000001b60000a13d000003e90020009c0000032b0000a13d000003ea0020009c000004bb0000613d000003eb0020009c0000049a0000613d000003ec0020009c000000530000c13d000000240060008c000000530000413d0000000002000416000000000002004b000000530000c13d0000000401100370000000000101043b000b00000001001d000004000010009c000000530000213d0000000b01000029000000000010043f0000000301000039000000200010043f0000000001000414000003de0010009c000003de01008041000000c00110021000000401011001c700008010020000390f720f6d0000040f0000000100200190000000530000613d000000000101043b000000000101041a000003e100100198000005240000c13d000000400100043d000004070010009c000004f90000213d0000002002100039000000400020043f0000000000010435000000400200043d00000020030000390000000003320436000000000401043300000000004304350000004003200039000000000004004b000000900000613d000000000500001900000020011000390000000006010433000003e10660019700000000036304360000000105500039000000000045004b000000890000413d0000000001230049000004750000013d000003fb0020009c000002e20000213d000003fe0020009c000003e30000613d000003ff0020009c000000530000c13d000000440060008c000000530000413d0000000002000416000000000002004b000000530000c13d0000000402100370000000000202043b000004000020009c000000530000213d0000002403100370000000000303043b000004000030009c000000530000213d0000000004360049000004040040009c000000530000213d000000a40040008c000000530000413d0000012004000039000000400040043f0000000405300039000000000751034f000000000707043b000004000070009c000000530000213d00000000083700190000002307800039000000000067004b000000530000813d0000000409800039000000000791034f000000000707043b000004000070009c000004f90000213d0000001f0a700039000004400aa001970000003f0aa00039000004400aa001970000041c00a0009c000004f90000213d000001200aa000390000004000a0043f000001200070043f00000000087800190000002408800039000000000068004b000000530000213d0000002008900039000000000981034f000004400a7001980000001f0b70018f0000014008a00039000000d30000613d000001400c000039000000000d09034f00000000de0d043c000000000cec043600000000008c004b000000cf0000c13d00000000000b004b000000e00000613d0000000009a9034f000000030ab00210000000000b080433000000000bab01cf000000000bab022f000000000909043b000001000aa000890000000009a9022f0000000009a901cf0000000009b9019f000000000098043500000140077000390000000000070435000000800040043f0000002004500039000000000541034f000000000505043b000004000050009c000000530000213d00000000093500190000002305900039000000000065004b000000530000813d000000040a9000390000000005a1034f000000000505043b000004000050009c000004f90000213d0000001f0750003900000440077001970000003f077000390000044008700197000000400700043d0000000008870019000000000078004b000000000b000039000000010b004039000004000080009c000004f90000213d0000000100b00190000004f90000c13d000000400080043f000000000857043600000000095900190000002409900039000000000069004b000000530000213d0000002009a00039000000000a91034f000004400b5001980000001f0c50018f0000000009b80019000001100000613d000000000d0a034f000000000e08001900000000df0d043c000000000efe043600000000009e004b0000010c0000c13d00000000000c004b0000011d0000613d000000000aba034f000000030bc00210000000000c090433000000000cbc01cf000000000cbc022f000000000a0a043b000001000bb00089000000000aba022f000000000aba01cf000000000aca019f0000000000a9043500000000055800190000000000050435000000a00070043f0000002004400039000000000541034f000000000505043b000004000050009c000000530000213d00000000073500190000002305700039000000000065004b000000530000813d0000000405700039000000000551034f000000000905043b000004000090009c000004f90000213d00000005059002100000003f055000390000040608500197000000400500043d0000000008850019000000000058004b000000000a000039000000010a004039000004000080009c000004f90000213d0000000100a00190000004f90000c13d000000400080043f0000000000950435000000240770003900000006089002100000000008780019000000000068004b000000530000213d000000000009004b0000088f0000c13d000000c00050043f0000002004400039000000000541034f000000000505043b000003e10050009c000000530000213d000000e00050043f0000002004400039000000000441034f000000000404043b000004000040009c000000530000213d00000000073400190000002303700039000000000063004b000000000400001900000405040080410000040503300197000000000003004b00000000050000190000040505004041000004050030009c000000000504c019000000000005004b000000530000c13d0000000408700039000000000381034f000000000303043b000004000030009c000004f90000213d0000001f0430003900000440044001970000003f044000390000044005400197000000400400043d0000000005540019000000000045004b00000000090000390000000109004039000004000050009c000004f90000213d0000000100900190000004f90000c13d000000400050043f000000000534043600000000073700190000002407700039000000000067004b000000530000213d0000002006800039000000000661034f00000440073001980000001f0830018f0000000001750019000001800000613d000000000906034f000000000a050019000000009b09043c000000000aba043600000000001a004b0000017c0000c13d000000000008004b0000018d0000613d000000000676034f0000000307800210000000000801043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f000000000061043500000000013500190000000000010435000001000040043f000000e00100043d000003e100100198000001970000c13d0000000201000039000000000101041a000003e101100197000000e00010043f0000040001200197000a00000001001d000000000010043f0000000301000039000000200010043f0000000001000414000003de0010009c000003de01008041000000c00110021000000401011001c700008010020000390f720f6d0000040f0000000100200190000000530000613d000000400200043d000b00000002001d0000000403200039000000000101043b000000000101041a000003e102100198000008db0000c13d0000042e010000410000000b0200002900000000001204350000000a010000290000000000130435000003de0020009c000003de02008041000000400120021000000403011001c700000f7400010430000003ef0020009c000002f40000213d000003f20020009c000003f70000613d000003f30020009c000000530000c13d000000440060008c000000530000413d0000000402100370000000000202043b000b00000002001d000004000020009c000000530000213d0000002402100370000000000202043b000004000020009c000000530000213d0000000003260049000004040030009c000000530000213d000000a40030008c000000530000413d0000012003000039000000400030043f0000000404200039000000000541034f000000000505043b000004000050009c000000530000213d00000000072500190000002305700039000000000065004b000000530000813d0000000408700039000000000581034f000000000505043b000004000050009c000004f90000213d0000001f0950003900000440099001970000003f0990003900000440099001970000041c0090009c000004f90000213d0000012009900039000000400090043f000001200050043f00000000075700190000002407700039000000000067004b000000530000213d0000002007800039000000000871034f00000440095001980000001f0a50018f0000014007900039000001f50000613d000001400b000039000000000c08034f00000000cd0c043c000000000bdb043600000000007b004b000001f10000c13d00000000000a004b000002020000613d000000000898034f0000000309a00210000000000a070433000000000a9a01cf000000000a9a022f000000000808043b0000010009900089000000000898022f00000000089801cf0000000008a8019f000000000087043500000140055000390000000000050435000000800030043f0000002003400039000000000431034f000000000404043b000004000040009c000000530000213d00000000082400190000002304800039000000000064004b000000530000813d0000000409800039000000000491034f000000000404043b000004000040009c000004f90000213d0000001f0540003900000440055001970000003f055000390000044007500197000000400500043d0000000007750019000000000057004b000000000a000039000000010a004039000004000070009c000004f90000213d0000000100a00190000004f90000c13d000000400070043f000000000745043600000000084800190000002408800039000000000068004b000000530000213d0000002008900039000000000981034f000004400a4001980000001f0b40018f0000000008a70019000002320000613d000000000c09034f000000000d07001900000000ce0c043c000000000ded043600000000008d004b0000022e0000c13d00000000000b004b0000023f0000613d0000000009a9034f000000030ab00210000000000b080433000000000bab01cf000000000bab022f000000000909043b000001000aa000890000000009a9022f0000000009a901cf0000000009b9019f000000000098043500000000044700190000000000040435000000a00050043f0000002003300039000000000431034f000000000404043b000004000040009c000000530000213d00000000052400190000002304500039000000000064004b000000530000813d0000000404500039000000000441034f000000000804043b000004000080009c000004f90000213d00000005048002100000003f044000390000040607400197000000400400043d0000000007740019000000000047004b00000000090000390000000109004039000004000070009c000004f90000213d0000000100900190000004f90000c13d000000400070043f0000000000840435000000240550003900000006078002100000000007570019000000000067004b000000530000213d000000000008004b000008a90000c13d000000c00040043f0000002003300039000000000431034f000000000404043b000003e10040009c000000530000213d000000e00040043f0000002003300039000000000331034f000000000303043b000004000030009c000000530000213d00000000052300190000002302500039000000000062004b000000000300001900000405030080410000040502200197000000000002004b00000000040000190000040504004041000004050020009c000000000403c019000000000004004b000000530000c13d0000000407500039000000000271034f000000000202043b000004000020009c000004f90000213d0000001f0320003900000440033001970000003f033000390000044004300197000000400300043d0000000004430019000000000034004b00000000080000390000000108004039000004000040009c000004f90000213d0000000100800190000004f90000c13d000000400040043f000000000423043600000000052500190000002405500039000000000065004b000000530000213d0000002005700039000000000551034f00000440062001980000001f0720018f0000000001640019000002a20000613d000000000805034f0000000009040019000000008a08043c0000000009a90436000000000019004b0000029e0000c13d000000000007004b000002af0000613d000000000565034f0000000306700210000000000701043300000000076701cf000000000767022f000000000505043b0000010006600089000000000565022f00000000056501cf000000000575019f000000000051043500000000012400190000000000010435000001000030043f0000041d01000041000000400200043d000a00000002001d00000000001204350000041e0100004100000000001004430000000001000412000000040010044300000024000004430000000001000414000003de0010009c000003de01008041000000c0011002100000041f011001c700008005020000390f720f6d0000040f000000010020019000000d0e0000613d000000000201043b0000000001000414000003e102200197000000040020008c000009440000c13d0000000103000031000000200030008c000000200400003900000000040340190000096e0000013d0000000003000411000000000003004b000003d10000c13d000000400100043d0000004402100039000003e4030000410000000000320435000000240210003900000018030000390000000000320435000003e5020000410000000000210435000000040210003900000020030000390000000000320435000003de0010009c000003de010080410000004001100210000003e6011001c700000f7400010430000003fc0020009c000003fc0000613d000003fd0020009c000000530000c13d0000000001000416000000000001004b000000530000c13d0000000001000412000d00000001001d000c00000000003d0000800501000039000000440300003900000000040004150000000d0440008a00000005044002100000041e020000410f720f4f0000040f000004e20000013d000003f00020009c000004320000613d000003f10020009c000000530000c13d000000240060008c000000530000413d0000000002000416000000000002004b000000530000c13d0000000401100370000000000101043b000004000010009c000000530000213d0f720ddb0000040f000000000001004b0000000001000039000000010100c039000000400200043d0000000000120435000003de0020009c000003de02008041000000400120021000000418011001c700000f730001042e000003f90020009c000004c00000613d000003fa0020009c000000530000c13d000000640060008c000000530000413d0000000002000416000000000002004b000000530000c13d0000000402100370000000000502043b000003e10050009c000000530000213d0000002402100370000000000402043b000003e10040009c000000530000213d000000000200041a000003e1022001970000000003000411000000000023004b000005090000c13d0000004402100370000000000302043b000000000004004b0000055c0000c13d0000043301000041000000800010043f000000840000043f000004340100004100000f7400010430000003ed0020009c000004d20000613d000003ee0020009c000000530000c13d000000640060008c000000530000413d0000000002000416000000000002004b000000530000c13d0000000402100370000000000202043b000004000020009c000000530000213d0000002303200039000000000063004b000000530000813d0000000403200039000000000331034f000000000303043b000700000003001d000004000030009c000000530000213d000600240020003d000000070200002900000006022002100000000602200029000000000062004b000000530000213d0000002402100370000000000202043b000004000020009c000000530000213d0000002303200039000000000063004b000000530000813d0000000403200039000000000331034f000000000303043b000300000003001d000004000030009c000000530000213d000200240020003d000000030200002900000006022002100000000202200029000000000062004b000000530000213d0000004402100370000000000202043b000004000020009c000000530000213d0000002303200039000000000063004b000000530000813d0000000403200039000000000131034f000000000101043b000500000001001d000004000010009c000000530000213d000400240020003d000000050100002900000006011002100000000401100029000000000061004b000000530000213d000000000100041a000003e1011001970000000002000411000000000012004b000005090000c13d000000070000006b000007b80000c13d000000030000006b000008030000c13d000000050000006b000004980000613d000b00000000001d0000037f0000013d0000000b020000290000000102200039000b00000002001d000000050020006c000004980000813d0000000b01000029000000060110021000000004011000290000000202000367000000000312034f000000000303043b000a00000003001d000004000030009c000000530000213d0000002001100039000000000112034f000000000101043b000900000001001d000003e10010009c000000530000213d0000000a01000029000000a001100210000004100110019700000009011001af000800000001001d000000000010043f0000000501000039000000200010043f0000000001000414000003de0010009c000003de01008041000000c00110021000000401011001c700008010020000390f720f6d0000040f0000000100200190000000530000613d000000000101043b000000000101041a000000000001004b0000037a0000c13d0000000403000039000000000103041a000004000010009c000004f90000213d0000000102100039000000000023041b000004160110009a0000000802000029000000000021041b000000000103041a000700000001001d000000000020043f0000000501000039000000200010043f0000000001000414000003de0010009c000003de01008041000000c00110021000000401011001c700008010020000390f720f6d0000040f0000000100200190000000530000613d000000000101043b0000000702000029000000000021041b000000400100043d00000009020000290000000000210435000003de0010009c000003de0100804100000040011002100000000002000414000003de0020009c000003de02008041000000c002200210000000000112019f0000040e011001c70000800d02000039000000020300003900000417040000410000000a050000290f720f680000040f00000001002001900000037a0000c13d000000530000013d000000000400041a000003e204400197000000000334019f000000000030041b0000000203000039000000000403041a000003e204400197000000000114019f000000000013041b000000800020043f000001400000044300000160002004430000002001000039000001000010044300000001010000390000012000100443000003e30100004100000f730001042e0000000001000416000000000001004b000000530000c13d000000c001000039000000400010043f0000000c01000039000000800010043f0000043e01000041000000a00010043f0000002001000039000000c00010043f0000008001000039000000e0020000390f720dc90000040f000000c00110008a000003de0010009c000003de0100804100000060011002100000043f011001c700000f730001042e0000000001000416000000000001004b000000530000c13d000000000100041a000004e20000013d000000840060008c000000530000413d0000000002000416000000000002004b000000530000c13d0000000402100370000000000202043b000b00000002001d000004000020009c000000530000213d0000000b0260006a000004040020009c000000530000213d000000a40020008c000000530000413d0000002402100370000000000202043b000a00000002001d0000ffff0020008c000000530000213d0000006402100370000000000202043b000900000002001d000003e10020009c000000530000213d0000004401100370000000000101043b000800000001001d0000041d01000041000000800010043f0000041e0100004100000000001004430000000001000412000000040010044300000024000004430000000001000414000003de0010009c000003de01008041000000c0011002100000041f011001c700008005020000390f720f6d0000040f000000010020019000000d0e0000613d000000000201043b0000000001000414000003e102200197000000040020008c000005650000c13d0000000103000031000000200030008c000000200400003900000000040340190000058a0000013d0000000001000416000000000001004b000000530000c13d0000000402000039000000000102041a000000800010043f000000000020043f000000000001004b000004e60000c13d000000a001000039000000400010043f0000002002000039000000000400001900000005034002100000003f053000390000041b055001970000000005150019000004000050009c000004f90000213d000000400050043f0000000000410435000000000004004b000004580000613d000000a0042000390000000005000019000000400600043d0000040d0060009c000004f90000213d0000004007600039000000400070043f000000200760003900000000000704350000000000060435000000000754001900000000006704350000002005500039000000000035004b0000044b0000413d000000800300043d000000000003004b000005400000c13d000000400200043d00000020030000390000000004320436000000000301043300000000003404350000004004200039000000000003004b000004740000613d0000000005000019000000000602001900000000070400190000002001100039000000000401043300000000840404340000040004400197000000000047043500000060046000390000000006080433000003e106600197000000000064043500000040047000390000000105500039000000000035004b0000000006070019000004650000413d0000000001240049000003de0010009c000003de010080410000006001100210000003de0020009c000003de020080410000004002200210000000000121019f00000f730001042e0000000001000416000000000001004b000000530000c13d0000000101000039000000000201041a000003e1032001970000000006000411000000000036004b000004ff0000c13d000000000300041a000003e204300197000000000464019f000000000040041b000003e202200197000000000021041b0000000001000414000003e105300197000003de0010009c000003de01008041000000c00110021000000408011001c70000800d02000039000000030300003900000430040000410f720f680000040f0000000100200190000000530000613d000000000100001900000f730001042e000000240060008c000000530000413d0000000002000416000000000002004b000000530000c13d0000000401100370000000000601043b000003e10060009c000000530000213d000000000100041a000003e1011001970000000005000411000000000015004b000005090000c13d000000000056004b000005130000c13d000003e501000041000000800010043f0000002001000039000000840010043f0000001701000039000000a40010043f0000040a01000041000000c40010043f0000040b0100004100000f74000104300000000001000416000000000001004b000000530000c13d0000008401000039000000800010043f0000040c0100004100000f730001042e0000000001000416000000000001004b000000530000c13d0000000201000039000004e10000013d000000240060008c000000530000413d0000000002000416000000000002004b000000530000c13d0000000401100370000000000101043b000b00000001001d000003e10010009c000000530000213d0f720df00000040f0000000201000039000000000201041a000003e2022001970000000b022001af000000000021041b000000000100001900000f730001042e000000240060008c000000530000413d0000000002000416000000000002004b000000530000c13d0000000401100370000000000101043b000004000010009c000000530000213d000000000010043f0000000301000039000000200010043f000000400200003900000000010000190f720f3a0000040f000000000101041a000003e101100197000000800010043f0000040c0100004100000f730001042e000000a004000039000004190200004100000000030000190000000005040019000000000402041a000000000445043600000001022000390000000103300039000000000013004b000004e90000413d000000410150008a00000440021001970000041a0020009c000004f90000213d0000008001200039000000800400043d000000400010043f000004000040009c0000043f0000a13d0000041201000041000000000010043f0000004101000039000000040010043f000004030100004100000f7400010430000003e501000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f0000042f01000041000000c40010043f0000040b0100004100000f7400010430000003e501000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f0000043101000041000000c40010043f0000040b0100004100000f74000104300000000101000039000000000201041a000003e202200197000000000262019f000000000021041b0000000001000414000003de0010009c000003de01008041000000c00110021000000408011001c70000800d02000039000000030300003900000409040000410f720f680000040f0000000100200190000000530000613d000004980000013d0000000b01000029000000000010043f0000000301000039000000200010043f0000000001000414000003de0010009c000003de01008041000000c00110021000000401011001c700008010020000390f720f6d0000040f0000000100200190000000530000613d000000000101043b000000000201041a000000400a00043d000004020100004100000000001a04350000000401a000390000000b0300002900000000003104350000000001000414000003e102200197000000040020008c000005be0000c13d00000003010003670000000103000031000005d00000013d0000000003000019000000400400043d0000040d0040009c000004f90000213d0000000505300210000000a00550003900000000060504330000004007400039000000400070043f0000002007400039000003e1086001970000000000870435000000a006600270000004000660019700000000006404350000000006010433000000000036004b00000d890000a13d000000000525001900000000004504350000000004010433000000000034004b00000d890000a13d0000000103300039000000800400043d000000000043004b000005410000413d0000045b0000013d000000000005004b000006220000c13d0000000002000414000000040040008c000006590000c13d000000000161034f00000001020000390000000103000031000006680000013d000003de0010009c000003de01008041000000c00110021000000435011001c70f720f6d0000040f0000006003100270000003de03300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000080057001bf000000800a000039000005790000613d000000000801034f000000008908043c000000000a9a043600000000005a004b000005750000c13d000000000006004b000005860000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f000300000001035500000001002001900000062f0000613d0000001f01400039000000600110018f00000080011001bf000000400010043f000000200030008c000000530000413d000000800200043d000000000002004b0000000003000039000000010300c039000000000032004b000000530000c13d000000000002004b0000069c0000c13d0000000b010000290000002401100039000700000001001d0000000201100367000000000101043b000004000010009c000000530000213d0000000002000411000503e10020019b000000a001100210000004100110019700000005011001af000000000010043f0000000501000039000000200010043f0000000001000414000003de0010009c000003de01008041000000c00110021000000401011001c700008010020000390f720f6d0000040f0000000100200190000000530000613d000000400200043d000600000002001d000000000101043b000000000101041a000000000001004b000006a10000c13d0000043d0100004100000006020000290000000000120435000003de0020009c000003de02008041000000400120021000000420011001c700000f7400010430000003de00a0009c000b0000000a001d000003de0300004100000000030a40190000004003300210000003de0010009c000003de01008041000000c001100210000000000131019f00000403011001c70f720f6d0000040f0000006003100270000103de0030019d000003de03300197000300000001035500000001002001900000063b0000613d0000000b0a00002900000440053001980000001f0630018f00000000045a0019000005da0000613d000000000701034f00000000080a0019000000007907043c0000000008980436000000000048004b000005d60000c13d000000000006004b000005e70000613d000000000151034f0000000305600210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000001f0130003900000440021001970000000001a20019000000000021004b00000000020000390000000102004039000004000010009c000004f90000213d0000000100200190000004f90000c13d000000400010043f000004040030009c000000530000213d000000200030008c000000530000413d00000000020a0433000004000020009c000000530000213d0000000003a300190000000002a200190000001f04200039000000000034004b0000000005000019000004050500804100000405044001970000040506300197000000000764013f000000000064004b00000000040000190000040504004041000004050070009c000000000405c019000000000004004b000000530000c13d0000000024020434000004000040009c000004f90000213d00000005054002100000003f0650003900000406066001970000000006160019000004000060009c000004f90000213d000000400060043f00000000004104350000000004250019000000000034004b000000530000213d000000000042004b000000800000813d00000000030100190000000025020434000003e10050009c000000530000213d00000020033000390000000000530435000000000042004b0000061a0000413d000000800000013d0000042a01000041000000a00010043f000000a40040043f000000c40030043f0000004401000039000000800010043f0000010001000039000000400010043f000000800200003900000000010500190f720e070000040f000000000100001900000f730001042e0000001f0530018f000003e006300198000000400200043d0000000004620019000006460000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000006360000c13d000006460000013d0000001f0530018f000003e006300198000000400200043d0000000004620019000006460000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000006420000c13d000000000005004b000006530000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000003de0020009c000003de020080410000004002200210000000000112019f00000f7400010430000003de0020009c000003de02008041000000c001200210000000000003004b000006600000c13d0000000002040019000006630000013d00000408011001c7000080090200003900000000050000190f720f680000040f00030000000103550000006003100270000103de0030019d000003de03300197000000000003004b000006740000c13d0000000100200190000004980000c13d000000400100043d00000432020000410000000000210435000003de0010009c000003de01008041000000400110021000000420011001c700000f7400010430000004000030009c000004f90000213d0000001f0530003900000440055001970000003f055000390000044006500197000000400500043d0000000006650019000000000056004b00000000070000390000000107004039000004000060009c000004f90000213d0000000100700190000004f90000c13d000000400060043f000000000635043600000440043001980000001f0530018f00000000034600190000068e0000613d000000000701034f000000007807043c0000000006860436000000000036004b0000068a0000c13d000000000005004b0000066a0000613d000000000141034f0000000304500210000000000503043300000000054501cf000000000545022f000000000101043b0000010004400089000000000141022f00000000014101cf000000000151019f00000000001304350000066a0000013d00000421020000410000000000210435000000400110021000000420011001c700000f7400010430000000060300002900000020023000390000043601000041000400000002001d00000000001204350000002401300039000000200200003900000000002104350000000704000029000300200040009200000002010003670000000302100360000000000202043b00000044033000390000000000230435000000000241034f000000000202043b000004000020009c000000530000213d0000000b030000290000000403300039000000060400002900000064044000390000000000240435000000000200003100000000043200490000001f0440008a00000007050000290000002006500039000000000561034f000000000705043b00000405087001970000040505400197000000000958013f000000000058004b00000000080000190000040508004041000000000047004b000000000a000019000004050a008041000004050090009c00000000080ac019000000000008004b000000530000c13d0000000008370019000000000781034f000000000707043b000004000070009c000000530000213d00000020088000390000000009720049000000000098004b000000000a000019000004050a0020410000040509900197000004050b800197000000000c9b013f00000000009b004b000000000900001900000405090040410000040500c0009c00000000090ac019000000000009004b000000530000c13d000000060d0000290000008409d00039000000a00a0000390000000000a90435000000e409d000390000000000790435000000000a81034f000004400b7001980000001f0c70018f0000010409d000390000000008b90019000006f30000613d000000000d0a034f000000000e09001900000000df0d043c000000000efe043600000000008e004b000006ef0000c13d00000000000c004b000007000000613d000000000aba034f000000030bc00210000000000c080433000000000cbc01cf000000000cbc022f000000000a0a043b000001000bb00089000000000aba022f000000000aba01cf000000000aca019f0000000000a80435000000000897001900000000000804350000002008600039000000000681034f000000000606043b000004050a600197000000000b5a013f00000000005a004b000000000a000019000004050a004041000000000046004b000000000c000019000004050c0080410000040500b0009c000000000a0cc01900000000000a004b000000530000c13d000000000a3600190000000006a1034f000000000606043b000004000060009c000000530000213d000000200aa00039000000000b6200490000000000ba004b000000000c000019000004050c002041000004050bb00197000004050da00197000000000ebd013f0000000000bd004b000000000b000019000004050b0040410000040500e0009c000000000b0cc01900000000000b004b000000530000c13d0000001f0770003900000440077001970000000009970019000000c007700039000000060b000029000000a40bb0003900000000007b0435000000000aa1034f0000000007690436000004400b6001980000001f0c60018f0000000009b70019000007380000613d000000000d0a034f000000000e07001900000000df0d043c000000000efe043600000000009e004b000007340000c13d00000000000c004b000007450000613d000000000aba034f000000030bc00210000000000c090433000000000cbc01cf000000000cbc022f000000000a0a043b000001000bb00089000000000aba022f000000000aba01cf000000000aca019f0000000000a90435000000000976001900000000000904350000002008800039000000000881034f000000000808043b0000040509800197000000000a59013f000000000059004b00000000050000190000040505004041000000000048004b000000000400001900000405040080410000040500a0009c000000000504c019000000000005004b000000530000c13d0000000004380019000000000341034f000000000303043b000004000030009c000000530000213d000000200440003900000006053002100000000005520049000000000054004b0000000008000019000004050800204100000405055001970000040509400197000000000a59013f000000000059004b000000000500001900000405050040410000040500a0009c000000000508c019000000000005004b000000530000c13d0000001f056000390000044005500197000000000575001900000006070000290000000006750049000000440660008a000000c40770003900000000006704350000000005350436000000000003004b000007850000613d0000000006000019000000000741034f000000000707043b000003e10070009c000000530000213d00000000077504360000002008400039000000000881034f000000000808043b0000000000870435000000400440003900000040055000390000000106600039000000000036004b000007770000413d00000006060000290000000003650049000000200430008a00000000004604350000001f0330003900000440033001970000000004630019000000000034004b00000000030000390000000103004039000200000004001d000004000040009c000004f90000213d0000000100300190000004f90000c13d0000000203000029000000400030043f000004370030009c000004f90000213d0000000205000029000000c003500039000000400030043f00000084040000390000000004450436000000000121034f000b00000004001d0000000002040019000000001401043c0000000002420436000000000032004b000007a00000c13d00000426010000410000000000100443000000090100002900000004001004430000000001000414000003de0010009c000003de01008041000000c00110021000000427011001c700008002020000390f720f6d0000040f000000010020019000000d0e0000613d000000000101043b000000000001004b000008c30000c13d0000043c01000041000000000010043f000004200100004100000f74000104300000000002000019000b00000002001d000000060120021000000006011000290000000002100079000004040020009c000000530000213d000000400020008c000000530000413d000000400400043d0000040d0040009c000004f90000213d0000004002400039000000400020043f0000000203000367000000000213034f000000000202043b000004000020009c000000530000213d00000000052404360000002001100039000000000113034f000000000301043b000003e10030009c000000530000213d0000000000350435000000000020043f0000000301000039000000200010043f0000000001000414000003de0010009c000003de01008041000000c00110021000000401011001c70000801002000039000a00000004001d000900000005001d000800000003001d0f720f6d0000040f000000080500002900000009040000290000000a030000290000000100200190000000530000613d000000000101043b000000000201041a000003e202200197000000000252019f000000000021041b00000000020304330000000001040433000003e101100197000000400300043d0000000000130435000003de0030009c000003de0300804100000040013002100000000003000414000003de0030009c000003de03008041000000c003300210000000000113019f0000040e011001c700000400052001970000800d0200003900000002030000390000040f040000410f720f680000040f0000000100200190000000530000613d0000000b020000290000000102200039000000070020006c000007b90000413d000003740000013d0000000002000019000900000002001d000000060120021000000002011000290000000202000367000000000312034f000000000303043b000b00000003001d000004000030009c000000530000213d0000002001100039000000000112034f000000000101043b000a00000001001d000003e10010009c000000530000213d0000000b01000029000000a00110021000000410011001970000000a011001af000800000001001d000000000010043f0000000501000039000000200010043f0000000001000414000003de0010009c000003de01008041000000c00110021000000401011001c700008010020000390f720f6d0000040f0000000100200190000000530000613d000000000101043b000000000301041a000000000003004b000008750000613d0000000401000039000000000201041a000000000002004b000008830000613d000000010130008a000000000023004b0000084a0000613d000000000012004b00000d890000a13d000004110130009a000004110220009a000000000202041a000000000021041b000000000020043f0000000501000039000000200010043f0000000001000414000003de0010009c000003de01008041000000c00110021000000401011001c70000801002000039000700000003001d0f720f6d0000040f0000000100200190000000530000613d000000000101043b0000000702000029000000000021041b0000000401000039000000000301041a000000000003004b000008890000613d000000010130008a000004110230009a000000000002041b0000000402000039000000000012041b0000000801000029000000000010043f0000000501000039000000200010043f0000000001000414000003de0010009c000003de01008041000000c00110021000000401011001c700008010020000390f720f6d0000040f0000000100200190000000530000613d000000000101043b000000000001041b000000400100043d0000000a020000290000000000210435000003de0010009c000003de0100804100000040011002100000000002000414000003de0020009c000003de02008041000000c002200210000000000112019f0000040e011001c70000800d02000039000000020300003900000413040000410000000b050000290f720f680000040f0000000100200190000000530000613d00000009020000290000000102200039000000030020006c000008040000413d000003760000013d000000400100043d00000024021000390000000a0300002900000000003204350000041402000041000000000021043500000004021000390000000b030000290000000000320435000003de0010009c000003de01008041000000400110021000000415011001c700000f74000104300000041201000041000000000010043f0000001101000039000000040010043f000004030100004100000f74000104300000041201000041000000000010043f0000003101000039000000040010043f000004030100004100000f74000104300000000009050019000000000a7600490000040400a0009c000000530000213d0000004000a0008c000000530000413d000000400a00043d0000040d00a0009c000004f90000213d000000400ba000390000004000b0043f000000000b71034f000000000b0b043b000003e100b0009c000000530000213d0000002009900039000000000bba0436000000200c700039000000000cc1034f000000000c0c043b0000000000cb04350000000000a904350000004007700039000000000087004b000008900000413d000001430000013d00000000080400190000000009560049000004040090009c000000530000213d000000400090008c000000530000413d000000400900043d0000040d0090009c000004f90000213d000000400a9000390000004000a0043f000000000a51034f000000000a0a043b000003e100a0009c000000530000213d0000002008800039000000000aa90436000000200b500039000000000bb1034f000000000b0b043b0000000000ba043500000000009804350000004005500039000000000075004b000008aa0000413d000002650000013d00000000010004140000000a0110006c000008ca0000813d0000043b01000041000000000010043f000004200100004100000f740001043000000006021002700000000001210049000000080010006c000008d70000a13d0000000001000414000100000001001d0000000901000029000000040010008c000009a40000c13d000800010000003d00000003010003670000000102000031000009ba0000013d0000043801000041000000000010043f000004200100004100000f740001043000000423010000410000000b0600002900000000001604350000000a0100002900000000001304350000002401600039000000400300003900000000003104350000004401600039000000800300043d000000a0040000390000000000410435000000e404600039000000005303043400000000003404350000010404600039000000000003004b000008f50000613d000000000600001900000000074600190000000008650019000000000808043300000000008704350000002006600039000000000036004b000008ee0000413d000000000534001900000000000504350000001f033000390000044003300197000000000334001900000000041300490000000b050000290000006405500039000000a00600043d000000000045043500000000540604340000000003430436000000000004004b0000090b0000613d000000000600001900000000073600190000000008650019000000000808043300000000008704350000002006600039000000000046004b000009040000413d000000000543001900000000000504350000001f044000390000044004400197000000000343001900000000051300490000000b040000290000008406400039000000c00400043d000000000056043500000000050404330000000003530436000000000005004b000009250000613d0000000006000019000000200440003900000000070404330000000087070434000003e10770019700000000077304360000000008080433000000000087043500000040033000390000000106600039000000000056004b0000091a0000413d000000e00400043d000003e1044001970000000b06000029000000a40560003900000000004504350000000001130049000000c404600039000001000500043d000000000014043500000000540504340000000001430436000000000004004b0000093a0000613d000000000300001900000000061300190000000007350019000000000707043300000000007604350000002003300039000000000043004b000009330000413d000000000341001900000000000304350000000003000414000000040020008c00000a350000c13d0000000103000031000000200030008c0000002004000039000000000403401900000a670000013d0000000a03000029000003de0030009c000003de030080410000004003300210000003de0010009c000003de01008041000000c001100210000000000131019f00000420011001c70f720f6d0000040f0000006003100270000003de03300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000000a057000290000095d0000613d000000000801034f0000000a09000029000000008a08043c0000000009a90436000000000059004b000009590000c13d000000000006004b0000096a0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000000a7c0000613d0000001f01400039000000600210018f0000000a01200029000000000021004b00000000020000390000000102004039000004000010009c000004f90000213d0000000100200190000004f90000c13d000000400010043f000000200030008c000000530000413d0000000a020000290000000002020433000000000002004b0000000003000039000000010300c039000000000032004b000000530000c13d000000000002004b00000a940000c13d0000000b010000290000040001100197000600000001001d000000000010043f0000000301000039000000200010043f0000000001000414000003de0010009c000003de01008041000000c00110021000000401011001c700008010020000390f720f6d0000040f0000000100200190000000530000613d000000400200043d000b00000002001d000000000101043b000000000101041a000703e10010019c00000a960000c13d0000042e010000410000000b030000290000000000130435000000040130003900000006020000290000000000210435000003de0030009c000003de03008041000000400130021000000403011001c700000f74000104300000000401000029000003de0010009c000003de01008041000000400110021000000006020000290000000002020433000003de0020009c000003de020080410000006002200210000000000112019f0000000802000029000003de0020009c000003de02008041000000c002200210000000000121019f00000009020000290f720f680000040f000800000002001d00030000000103550000006002100270000103de0020019d000003de02200197000000840020008c00000084020080390000000003000414000900000003001d000000020300002900000000002304350000001f0320018f000000e0042001900000000b02400029000009ca0000613d000000000501034f0000000b06000029000000005705043c0000000006760436000000000026004b000009c60000c13d000000000003004b000009d70000613d000000000141034f0000000303300210000000000402043300000000043401cf000000000434022f000000000101043b0000010003300089000000000131022f00000000013101cf000000000141019f000000000012043500000007010000290000000201100367000000000101043b000a00000001001d000004000010009c000000530000213d0000000401000029000003de0010009c000003de01008041000000400110021000000006020000290000000002020433000003de0020009c000003de020080410000006002200210000000000112019f0000000002000414000003de0020009c000003de02008041000000c002200210000000000112019f00000408011001c700008010020000390f720f6d0000040f0000000100200190000000530000613d00000003020000290000000202200367000000000202043b000000000101043b000000400300043d0000006004300039000000000014043500000040013000390000000504000029000000000041043500000020013000390000000a0400002900000000004104350000000000230435000003de0030009c000003de0300804100000040013002100000000002000414000003de0020009c000003de02008041000000c002200210000000000112019f00000439011001c70000800d0200003900000001030000390000043a040000410f720f680000040f0000000100200190000000530000613d000000400100043d0000002002100039000000600300003900000000003204350000000802000029000000010220018f000000000021043500000002020000290000000002020433000000600310003900000000002304350000008003100039000000000002004b00000a240000613d000000000400001900000000053400190000000b06400029000000000606043300000000006504350000002004400039000000000024004b00000a1d0000413d0000000003320019000000000003043500000009040000290000000103400069000000400410003900000000003404350000001f0220003900000440022001970000008002200039000003de0020009c000003de020080410000006002200210000003de0010009c000003de010080410000004001100210000000000112019f00000f730001042e0000001f0440003900000440044001970000000b0500002900000000045400490000000001140019000003de0010009c000003de010080410000006001100210000003de0050009c000003de0400004100000000040540190000004004400210000000000141019f000003de0030009c000003de03008041000000c003300210000000000131019f0f720f6d0000040f0000006003100270000003de03300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000000b0570002900000a560000613d000000000801034f0000000b09000029000000008a08043c0000000009a90436000000000059004b00000a520000c13d000000000006004b00000a630000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000000a880000613d0000001f01400039000000600210018f0000000b01200029000000000021004b00000000020000390000000102004039000004000010009c000004f90000213d0000000100200190000004f90000c13d000000400010043f000000200030008c000000530000413d0000000b0200002900000000020204330000000000210435000003de0010009c000003de01008041000000400110021000000418011001c700000f730001042e0000001f0530018f000003e006300198000000400200043d0000000004620019000006460000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00000a830000c13d000006460000013d0000001f0530018f000003e006300198000000400200043d0000000004620019000006460000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00000a8f0000c13d000006460000013d00000421020000410000066e0000013d000000e00100043d000003e10010019800000b080000c13d0000000201000039000000000101041a000003e101100197000000e00010043f0000000b05000029000000240150003900000040020000390000000000210435000004230100004100000000001504350000000401500039000000060200002900000000002104350000004401500039000000800200043d000000a0030000390000000000310435000000e403500039000000004202043400000000002304350000010403500039000000000002004b00000ab80000613d000000000500001900000000063500190000000007540019000000000707043300000000007604350000002005500039000000000025004b00000ab10000413d000000000423001900000000000404350000001f022000390000044002200197000000000223001900000000031200490000000b040000290000006404400039000000a00500043d000000000034043500000000430504340000000002320436000000000003004b00000ace0000613d000000000500001900000000062500190000000007540019000000000707043300000000007604350000002005500039000000000035004b00000ac70000413d000000000432001900000000000404350000001f033000390000044003300197000000000232001900000000041200490000000b030000290000008405300039000000c00300043d000000000045043500000000040304330000000002420436000000000004004b00000ae80000613d0000000005000019000000200330003900000000060304330000000076060434000003e10660019700000000066204360000000007070433000000000076043500000040022000390000000105500039000000000045004b00000add0000413d000000e00300043d000003e1033001970000000b05000029000000a40450003900000000003404350000000001120049000000c403500039000001000400043d000000000013043500000000430404340000000001320436000000000003004b00000afd0000613d000000000200001900000000051200190000000006240019000000000606043300000000006504350000002002200039000000000032004b00000af60000413d0000000002310019000000000002043500000000020004140000000704000029000000040040008c00000b760000c13d0000000103000031000000200030008c0000002004000039000000000403401900000ba90000013d0000000001000416000000000001004b00000bbe0000c13d0000000b05000029000000240150003900000040020000390000000000210435000004230100004100000000001504350000000401500039000000060200002900000000002104350000004401500039000000800200043d000000a0030000390000000000310435000000e403500039000000004202043400000000002304350000010403500039000000000002004b00000b260000613d000000000500001900000000063500190000000007540019000000000707043300000000007604350000002005500039000000000025004b00000b1f0000413d000000000423001900000000000404350000001f022000390000044002200197000000000223001900000000031200490000000b040000290000006404400039000000a00500043d000000000034043500000000430504340000000002320436000000000003004b00000b3c0000613d000000000500001900000000062500190000000007540019000000000707043300000000007604350000002005500039000000000035004b00000b350000413d000000000432001900000000000404350000001f033000390000044003300197000000000232001900000000041200490000000b030000290000008405300039000000c00300043d000000000045043500000000040304330000000002420436000000000004004b00000b560000613d0000000005000019000000200330003900000000060304330000000076060434000003e10660019700000000066204360000000007070433000000000076043500000040022000390000000105500039000000000045004b00000b4b0000413d000000e00300043d000003e1033001970000000b05000029000000a40450003900000000003404350000000001120049000000c403500039000001000400043d000000000013043500000000430404340000000001320436000000000003004b00000b6b0000613d000000000200001900000000051200190000000006240019000000000606043300000000006504350000002002200039000000000032004b00000b640000413d0000000002310019000000000002043500000000020004140000000704000029000000040040008c00000bc10000c13d0000000103000031000000200030008c0000002004000039000000000403401900000bf40000013d0000001f0330003900000440033001970000000b0400002900000000034300490000000001130019000003de0010009c000003de010080410000006001100210000003de0040009c000003de0300004100000000030440190000004003300210000000000131019f000003de0020009c000003de02008041000000c002200210000000000121019f00000007020000290f720f6d0000040f0000006003100270000003de03300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000000b0570002900000b980000613d000000000801034f0000000b09000029000000008a08043c0000000009a90436000000000059004b00000b940000c13d000000000006004b00000ba50000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000000ccd0000613d0000001f01400039000000600210018f0000000b01200029000000000021004b00000000020000390000000102004039000004000010009c000004f90000213d0000000100200190000004f90000c13d000000400010043f000000200030008c000000530000413d0000000b0200002900000000020204330000000003000416000500000003001d000000000023004b00000ce50000813d0000042d020000410000066e0000013d00000422010000410000000b02000029000005b80000013d0000001f0330003900000440033001970000000b0400002900000000034300490000000001130019000003de0010009c000003de010080410000006001100210000003de0040009c000003de0300004100000000030440190000004003300210000000000131019f000003de0020009c000003de02008041000000c002200210000000000121019f00000007020000290f720f6d0000040f0000006003100270000003de03300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000000b0570002900000be30000613d000000000801034f0000000b09000029000000008a08043c0000000009a90436000000000059004b00000bdf0000c13d000000000006004b00000bf00000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000000cd90000613d0000001f01400039000000600110018f0000000b02100029000000000012004b00000000010000390000000101004039000004000020009c000004f90000213d0000000100100190000004f90000c13d000000400020043f000000200030008c000000530000413d0000000b0100002900000000050104330000002003200039000000e00100043d000004240400004100000000004304350000004403200039000000070400002900000000004304350000000003000411000003e103300197000000240420003900000000003404350000006403200039000500000005001d000000000053043500000064030000390000000000320435000004250020009c000004f90000213d000000a003200039000000400030043f000003e1011001970f720e070000040f000000c00100043d0000000002010433000000000002004b00000d0f0000c13d0000000001000411000400000001001d000000400500043d0000002401500039000000800200003900000000002104350000042c0100004100000000001504350000000401500039000000060200002900000000002104350000008401500039000000800200043d000000a0030000390000000000310435000001240350003900000000420204340000000000230435000b00000005001d0000014403500039000000000002004b00000c3b0000613d000000000500001900000000063500190000000007540019000000000707043300000000007604350000002005500039000000000025004b00000c340000413d000000000423001900000000000404350000001f022000390000044002200197000000000223001900000000031200490000000b04000029000000a404400039000000a00500043d000000000034043500000000430504340000000002320436000000000003004b00000c510000613d000000000500001900000000062500190000000007540019000000000707043300000000007604350000002005500039000000000035004b00000c4a0000413d000000000432001900000000000404350000001f033000390000044003300197000000000232001900000000041200490000000b03000029000000c405300039000000c00300043d000000000045043500000000040304330000000002420436000000000004004b00000c6b0000613d0000000005000019000000200330003900000000060304330000000076060434000003e10660019700000000066204360000000007070433000000000076043500000040022000390000000105500039000000000045004b00000c600000413d000000e00300043d000003e1033001970000000b05000029000000e404500039000000000034043500000000011200490000010403500039000001000400043d000000000013043500000000430404340000000001320436000000000003004b00000c800000613d000000000200001900000000051200190000000006240019000000000606043300000000006504350000002002200039000000000032004b00000c790000413d000000000231001900000000000204350000000402000029000003e1022001970000000b050000290000006404500039000000000024043500000044025000390000000504000029000000000042043500000000020004140000000704000029000000040040008c0000093f0000613d0000001f0330003900000440033001970000000b0400002900000000034300490000000001130019000003de0010009c000003de010080410000006001100210000003de0040009c000003de0300004100000000030440190000004003300210000000000131019f000003de0020009c000003de02008041000000c002200210000000000112019f00000007020000290f720f680000040f0000006003100270000003de03300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000000b0570002900000cb00000613d000000000801034f0000000b09000029000000008a08043c0000000009a90436000000000059004b00000cac0000c13d000000000006004b00000cbd0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000000a670000c13d0000001f0530018f000003e006300198000000400200043d0000000004620019000006460000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00000cc80000c13d000006460000013d0000001f0530018f000003e006300198000000400200043d0000000004620019000006460000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00000cd40000c13d000006460000013d0000001f0530018f000003e006300198000000400200043d0000000004620019000006460000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00000ce00000c13d000006460000013d000000e00100043d00000426020000410000000000200443000003e101100197000b00000001001d00000004001004430000000001000414000003de0010009c000003de01008041000000c00110021000000427011001c700008002020000390f720f6d0000040f000000010020019000000d0e0000613d000000000101043b000000000001004b000000530000613d000000400200043d0000042801000041000a00000002001d0000000001120436000900000001001d00000000010004140000000b02000029000000040020008c00000d9a0000613d0000000a02000029000003de0020009c000003de020080410000004002200210000003de0010009c000003de01008041000000c001100210000000000121019f0000000002000416000000000002004b00000d8f0000c13d00000420011001c70000000b0200002900000d940000013d000000000001042f0000000002000411000400000002001d000a00000000001d0000000a020000290000000502200210000900200020003d000000090110002900000000010104330000000001010433000000400400043d0000042b020000410000000000240435000000040240003900000006030000290000000000320435000003e102100197000b00000004001d0000002401400039000800000002001d000000000021043500000000010004140000000702000029000000040020008c00000d2c0000c13d0000000103000031000000200030008c0000002004000039000000000403401900000d560000013d0000000b02000029000003de0020009c000003de020080410000004002200210000003de0010009c000003de01008041000000c001100210000000000121019f00000415011001c700000007020000290f720f6d0000040f0000006003100270000003de03300197000000200030008c0000002004000039000000000403401900000020064001900000000b0560002900000d450000613d000000000701034f0000000b08000029000000007907043c0000000008980436000000000058004b00000d410000c13d0000001f0740019000000d520000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000000db00000613d0000001f01400039000000600110018f0000000b02100029000000000012004b00000000010000390000000101004039000004000020009c000004f90000213d0000000100100190000004f90000c13d000000400020043f000000200030008c000000530000413d0000000b010000290000000001010433000003e10010009c000000530000213d000000c00300043d00000000040304330000000a0040006c00000d890000a13d000000090330002900000000030304330000002003300039000000000303043300000020042000390000042405000041000000000054043500000064042000390000000000340435000000440320003900000000001304350000000001000411000003e1011001970000002403200039000000000013043500000064010000390000000000120435000004250020009c000004f90000213d000000a001200039000000400010043f00000008010000290f720e070000040f0000000a03000029000a00010030003d000000c00100043d00000000020104330000000a0020006b00000d120000413d00000c1f0000013d0000041201000041000000000010043f0000003201000039000000040010043f000004030100004100000f740001043000000429011001c7000080090200003900000000030004160000000b0400002900000000050000190f720f680000040f00030000000103550000006003100270000103de0030019d000000010020019000000dbc0000613d0000000a01000029000004000010009c000004f90000213d0000000a04000029000000400040043f000000e00100043d0000042a0200004100000009030000290000000000230435000000440240003900000000030004160000000000320435000000240240003900000007030000290000000000320435000000440200003900000000002404350000041a0040009c000004f90000213d0000000a02000029000000800320003900000c160000013d0000001f0530018f000003e006300198000000400200043d0000000004620019000006460000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00000db70000c13d000006460000013d000003de033001970000001f0530018f000003e006300198000000400200043d0000000004620019000006460000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00000dc40000c13d000006460000013d00000000430104340000000001320436000000000003004b00000dd50000613d000000000200001900000000052100190000000006240019000000000606043300000000006504350000002002200039000000000032004b00000dce0000413d000000000231001900000000000204350000001f0230003900000440022001970000000001210019000000000001042d0000040001100197000000000010043f0000000301000039000000200010043f0000000001000414000003de0010009c000003de01008041000000c00110021000000401011001c700008010020000390f720f6d0000040f000000010020019000000dee0000613d000000000101043b000000000101041a000003e1001001980000000001000039000000010100c039000000000001042d000000000100001900000f7400010430000000000100041a000003e1011001970000000002000411000000000012004b00000df60000c13d000000000001042d000000400100043d000000440210003900000431030000410000000000320435000000240210003900000016030000390000000000320435000003e5020000410000000000210435000000040210003900000020030000390000000000320435000003de0010009c000003de010080410000004001100210000003e6011001c700000f74000104300005000000000002000000400500043d000004410050009c00000eca0000813d000003e1061001970000004001500039000000400010043f000000200100003900000000041504360000044201000041000000000014043500000000230204340000000001000414000000040060008c00000e410000c13d000000010100003200000e7d0000613d000004000010009c00000eca0000213d0000001f0310003900000440033001970000003f033000390000044003300197000000400a00043d00000000033a00190000000000a3004b00000000040000390000000104004039000004000030009c00000eca0000213d000000010040019000000eca0000c13d000000400030043f00000000051a043600000440021001980000001f0310018f0000000001250019000000030400036700000e330000613d000000000604034f000000006706043c0000000005750436000000000015004b00000e2f0000c13d000000000003004b00000e7e0000613d000000000224034f0000000303300210000000000401043300000000043401cf000000000434022f000000000202043b0000010003300089000000000232022f00000000023201cf000000000242019f000000000021043500000e7e0000013d000100000005001d000300000004001d000003de0030009c000003de030080410000006003300210000003de0020009c000003de020080410000004002200210000000000223019f000003de0010009c000003de01008041000000c001100210000000000112019f000200000006001d00000000020600190f720f680000040f00030000000103550000006003100270000103de0030019d000003de0430019800000e950000613d0000001f03400039000003df033001970000003f033000390000044303300197000000400a00043d00000000033a00190000000000a3004b00000000050000390000000105004039000004000030009c00000eca0000213d000000010050019000000eca0000c13d000000400030043f0000001f0540018f00000000034a0436000003e006400198000000000463001900000e6f0000613d000000000701034f0000000008030019000000007907043c0000000008980436000000000048004b00000e6b0000c13d000000000005004b00000e970000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f000000000014043500000e970000013d000000600a0000390000000002000415000000050220008a000000050220021000000000010a0433000000000001004b00000e9f0000c13d00030000000a001d00000426010000410000000000100443000000040100003900000004001004430000000001000414000003de0010009c000003de01008041000000c00110021000000427011001c700008002020000390f720f6d0000040f000000010020019000000f0b0000613d0000000002000415000000050220008a00000eb20000013d000000600a000039000000800300003900000000010a0433000000010020019000000ee60000613d0000000002000415000000040220008a0000000502200210000000000001004b00000ea20000613d000000050220027000000000020a001f00000ebc0000013d00030000000a001d00000426010000410000000000100443000000020100002900000004001004430000000001000414000003de0010009c000003de01008041000000c00110021000000427011001c700008002020000390f720f6d0000040f000000010020019000000f0b0000613d0000000002000415000000040220008a0000000502200210000000000101043b000000000001004b000000030a00002900000f0c0000613d00000000010a0433000000050220027000000000020a001f000000000001004b00000ec90000613d000004040010009c00000ed00000213d0000001f0010008c00000ed00000a13d0000002001a000390000000001010433000000000001004b0000000002000039000000010200c039000000000021004b00000ed00000c13d000000000001004b00000ed20000613d000000000001042d0000041201000041000000000010043f0000004101000039000000040010043f000004030100004100000f7400010430000000000100001900000f7400010430000000400100043d00000064021000390000044403000041000000000032043500000044021000390000044503000041000000000032043500000024021000390000002a030000390000000000320435000003e5020000410000000000210435000000040210003900000020030000390000000000320435000003de0010009c000003de01008041000000400110021000000446011001c700000f7400010430000000000001004b00000f1d0000c13d000000400100043d000003e502000041000000000021043500000004021000390000002003000039000000000032043500000001020000290000000002020433000000240310003900000000002304350000004403100039000000000002004b000000030700002900000efe0000613d000000000400001900000000053400190000000006740019000000000606043300000000006504350000002004400039000000000024004b00000ef70000413d0000001f042000390000044004400197000000000223001900000000000204350000004402400039000003de0020009c000003de020080410000006002200210000003de0010009c000003de010080410000004001100210000000000112019f00000f7400010430000000000001042f000000400100043d00000044021000390000044703000041000000000032043500000024021000390000001d030000390000000000320435000003e5020000410000000000210435000000040210003900000020030000390000000000320435000003de0010009c000003de010080410000004001100210000003e6011001c700000f7400010430000003de0030009c000003de030080410000004002300210000003de0010009c000003de010080410000006001100210000000000121019f00000f7400010430000000000010043f0000000501000039000000200010043f0000000001000414000003de0010009c000003de01008041000000c00110021000000401011001c700008010020000390f720f6d0000040f000000010020019000000f370000613d000000000101043b000000000101041a000000000001004b0000000001000039000000010100c039000000000001042d000000000100001900000f7400010430000000000001042f000003de0010009c000003de010080410000004001100210000003de0020009c000003de020080410000006002200210000000000112019f0000000002000414000003de0020009c000003de02008041000000c002200210000000000112019f00000408011001c700008010020000390f720f6d0000040f000000010020019000000f4d0000613d000000000101043b000000000001042d000000000100001900000f740001043000000000050100190000000000200443000000040100003900000005024002700000000002020031000000000121043a0000002004400039000000000031004b00000f520000413d000003de0030009c000003de0300804100000060013002100000000002000414000003de0020009c000003de02008041000000c002200210000000000112019f00000448011001c700000000020500190f720f6d0000040f000000010020019000000f670000613d000000000101043b000000000001042d000000000001042f00000f6b002104210000000102000039000000000001042d0000000002000019000000000001042d00000f70002104230000000102000039000000000001042d0000000002000019000000000001042d00000f720000043200000f730001042e00000f740001043000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000000000000000000000000001ffffffe000000000000000000000000000000000000000000000000000000000ffffffe0000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000020000000000000000000000000000008000000100000000000000000043616e6e6f7420736574206f776e657220746f207a65726f000000000000000008c379a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000000000000000000000008da5cb5a00000000000000000000000000000000000000000000000000000000a8d87a3a00000000000000000000000000000000000000000000000000000000e861e90600000000000000000000000000000000000000000000000000000000e861e90700000000000000000000000000000000000000000000000000000000f2fde38b00000000000000000000000000000000000000000000000000000000fbca3b7400000000000000000000000000000000000000000000000000000000a8d87a3b00000000000000000000000000000000000000000000000000000000da5fcac800000000000000000000000000000000000000000000000000000000a40e69c600000000000000000000000000000000000000000000000000000000a40e69c700000000000000000000000000000000000000000000000000000000a48a9058000000000000000000000000000000000000000000000000000000008da5cb5b0000000000000000000000000000000000000000000000000000000096f4e9f90000000000000000000000000000000000000000000000000000000052cb60c900000000000000000000000000000000000000000000000000000000787350e200000000000000000000000000000000000000000000000000000000787350e30000000000000000000000000000000000000000000000000000000079ba50970000000000000000000000000000000000000000000000000000000083826b2b0000000000000000000000000000000000000000000000000000000052cb60ca000000000000000000000000000000000000000000000000000000005f3e849f000000000000000000000000000000000000000000000000000000003cf97982000000000000000000000000000000000000000000000000000000003cf97983000000000000000000000000000000000000000000000000000000005246492f00000000000000000000000000000000000000000000000000000000181f5a770000000000000000000000000000000000000000000000000000000020487ded000000000000000000000000000000000000000000000000ffffffffffffffff0200000000000000000000000000000000000040000000000000000000000000fbca3b740000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000240000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0000000000000000000000000000000000000000000000000ffffffffffffffdf0200000000000000000000000000000000000000000000000000000000000000ed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127843616e6e6f74207472616e7366657220746f2073656c6600000000000000000000000000000000000000000000000000000000640000008000000000000000000000000000000000000000000000000000000020000000800000000000000000000000000000000000000000000000000000000000000000ffffffffffffffbf02000000000000000000000000000000000000200000000000000000000000001f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f2300000000ffffffffffffffff000000000000000000000000000000000000000075ca53043ea007e5c65182cbb028f60d7179ff4b55739a3949b401801c942e664e487b7100000000000000000000000000000000000000000000000000000000a823809efda3ba66c873364eec120fa0923d9fabda73bc97dd5663341e2d9bcb4964779000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004400000000000000000000000075ca53043ea007e5c65182cbb028f60d7179ff4b55739a3949b401801c942e65a4bdf64ebdf3316320601a081916a75aa144bcef6c4beeb0e9fb1982cacc6b9400000000000000000000000000000000000000200000000000000000000000008a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b000000000000000000000000000000000000000000000000ffffffffffffff7f00000000000000000000000000000000000000000000003fffffffffffffffe0000000000000000000000000000000000000000000000000fffffffffffffedf397796f700000000000000000000000000000000000000000000000000000000310ab089e4439a4c15d089f94afb7896ff553aecb10793d0ab882de59d99a32e02000002000000000000000000000000000000440000000000000000000000000000000000000000000000000000000000000004000000000000000000000000c1483715000000000000000000000000000000000000000000000000000000001841b4e10000000000000000000000000000000000000000000000000000000020487ded0000000000000000000000000000000000000000000000000000000023b872dd00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffff5f1806aa1896bbf26568e884a7374b41e002500962caba6a15023a8d90e8508b830200000200000000000000000000000000000024000000000000000000000000d0e30db0000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000004000000000000000000000000a9059cbb0000000000000000000000000000000000000000000000000000000048a98aa400000000000000000000000000000000000000000000000000000000df0aa9e90000000000000000000000000000000000000000000000000000000007da6ee600000000000000000000000000000000000000000000000000000000ae236d9c000000000000000000000000000000000000000000000000000000004d7573742062652070726f706f736564206f776e6572000000000000000000008be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e04f6e6c792063616c6c61626c65206279206f776e657200000000000000000000e417b80b0000000000000000000000000000000000000000000000000000000026a78f8f000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024000000800000000000000000000000000000000000000000000000000000000400000080000000000000000085572ffb00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffff3f37c3be290000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000800000000000000000000000009b877de93ea9895756e337442c657f95a34fc68e7eb988bdfa693d5be83016b6afa32a2c000000000000000000000000000000000000000000000000000000000c3b563c00000000000000000000000000000000000000000000000000000000d2316ede00000000000000000000000000000000000000000000000000000000526f7574657220312e322e3000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0000000000000000000000000000000000000000000000000ffffffffffffffc05361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656400000000000000000000000000000000000000000000000000000003ffffffe06f742073756363656564000000000000000000000000000000000000000000005361666545524332303a204552433230206f7065726174696f6e20646964206e0000000000000000000000000000000000000084000000000000000000000000416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000002000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a9eab750d6e298d6ff1780825c0ff7a2d04cbacf6a3343715087eb6753a119f9")

func (_Router *Router) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Router.abi.Events["MessageExecuted"].ID:
		return _Router.ParseMessageExecuted(log)
	case _Router.abi.Events["OffRampAdded"].ID:
		return _Router.ParseOffRampAdded(log)
	case _Router.abi.Events["OffRampRemoved"].ID:
		return _Router.ParseOffRampRemoved(log)
	case _Router.abi.Events["OnRampSet"].ID:
		return _Router.ParseOnRampSet(log)
	case _Router.abi.Events["OwnershipTransferRequested"].ID:
		return _Router.ParseOwnershipTransferRequested(log)
	case _Router.abi.Events["OwnershipTransferred"].ID:
		return _Router.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (RouterMessageExecuted) Topic() common.Hash {
	return common.HexToHash("0x9b877de93ea9895756e337442c657f95a34fc68e7eb988bdfa693d5be83016b6")
}

func (RouterOffRampAdded) Topic() common.Hash {
	return common.HexToHash("0xa4bdf64ebdf3316320601a081916a75aa144bcef6c4beeb0e9fb1982cacc6b94")
}

func (RouterOffRampRemoved) Topic() common.Hash {
	return common.HexToHash("0xa823809efda3ba66c873364eec120fa0923d9fabda73bc97dd5663341e2d9bcb")
}

func (RouterOnRampSet) Topic() common.Hash {
	return common.HexToHash("0x1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f23")
}

func (RouterOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (RouterOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_Router *Router) Address() common.Address {
	return _Router.address
}

type CustomTransaction struct {
	*types.Transaction
	CustomHash common.Hash
}

func (tx *CustomTransaction) Hash() common.Hash {
	return tx.CustomHash
}

func ConvertToTransaction(resp zktypes.TransactionResponse) *CustomTransaction {
	dtx := &types.DynamicFeeTx{
		ChainID:   resp.ChainID.ToInt(),
		Nonce:     uint64(resp.Nonce),
		GasTipCap: resp.MaxPriorityFeePerGas.ToInt(),
		GasFeeCap: resp.MaxFeePerGas.ToInt(),
		To:        &resp.To,
		Value:     resp.Value.ToInt(),
		Data:      resp.Data,
		Gas:       uint64(resp.Gas),
	}

	tx := types.NewTx(dtx)
	customTransaction := CustomTransaction{Transaction: tx, CustomHash: resp.Hash}
	return &customTransaction
}

func DeployZkSyncRouter(auth *bind.TransactOpts, backend bind.ContractBackend, params ...interface{}) (common.Address, *CustomTransaction, *Router, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	fmt.Println("Deploying zksync contract")
	zksyncClient := zkSyncClient.NewClient(client.Client())
	fmt.Println("getting wallet")
	wallet := auth.Context.Value("wallet").(*zkSyncAccounts.Wallet)
	fmt.Println("got wallet")
	fmt.Println("getting bytes")
	decodedBytes := common.FromHex(RouterZkBin)
	fmt.Println("deploying")
	RouterAbi, err := RouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	constructor, err := RouterAbi.Pack("", params...)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	hash, err := wallet.DeployWithCreate(nil, zkSyncAccounts.CreateTransaction{
		Bytecode: decodedBytes,
		Calldata: constructor,
	})
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	fmt.Println("hash of tx", hash)
	receipt, err := zksyncClient.WaitMined(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	tx, _, err := zksyncClient.TransactionByHash(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	fmt.Println("tx hash", tx.Hash)

	ethTx := ConvertToTransaction(*tx)
	address := receipt.ContractAddress

	parsed, err := RouterMetaData.GetAbi()
	contractBind := bind.NewBoundContract(address, *parsed, backend, backend, backend)

	contractReturn := &Router{address: address, abi: *parsed, RouterCaller: RouterCaller{contract: contractBind}, RouterTransactor: RouterTransactor{contract: contractBind}, RouterFilterer: RouterFilterer{contract: contractBind}}

	return address, ethTx, contractReturn, err
}

type RouterInterface interface {
	MAXRETBYTES(opts *bind.CallOpts) (uint16, error)

	GetArmProxy(opts *bind.CallOpts) (common.Address, error)

	GetFee(opts *bind.CallOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetOffRamps(opts *bind.CallOpts) ([]RouterOffRamp, error)

	GetOnRamp(opts *bind.CallOpts, destChainSelector uint64) (common.Address, error)

	GetSupportedTokens(opts *bind.CallOpts, chainSelector uint64) ([]common.Address, error)

	GetWrappedNative(opts *bind.CallOpts) (common.Address, error)

	IsChainSupported(opts *bind.CallOpts, chainSelector uint64) (bool, error)

	IsOffRamp(opts *bind.CallOpts, sourceChainSelector uint64, offRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyRampUpdates(opts *bind.TransactOpts, onRampUpdates []RouterOnRamp, offRampRemoves []RouterOffRamp, offRampAdds []RouterOffRamp) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*types.Transaction, error)

	RecoverTokens(opts *bind.TransactOpts, tokenAddress common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	RouteMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage, gasForCallExactCheck uint16, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error)

	SetWrappedNative(opts *bind.TransactOpts, wrappedNative common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterMessageExecuted(opts *bind.FilterOpts) (*RouterMessageExecutedIterator, error)

	WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *RouterMessageExecuted) (event.Subscription, error)

	ParseMessageExecuted(log types.Log) (*RouterMessageExecuted, error)

	FilterOffRampAdded(opts *bind.FilterOpts, sourceChainSelector []uint64) (*RouterOffRampAddedIterator, error)

	WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *RouterOffRampAdded, sourceChainSelector []uint64) (event.Subscription, error)

	ParseOffRampAdded(log types.Log) (*RouterOffRampAdded, error)

	FilterOffRampRemoved(opts *bind.FilterOpts, sourceChainSelector []uint64) (*RouterOffRampRemovedIterator, error)

	WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *RouterOffRampRemoved, sourceChainSelector []uint64) (event.Subscription, error)

	ParseOffRampRemoved(log types.Log) (*RouterOffRampRemoved, error)

	FilterOnRampSet(opts *bind.FilterOpts, destChainSelector []uint64) (*RouterOnRampSetIterator, error)

	WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *RouterOnRampSet, destChainSelector []uint64) (event.Subscription, error)

	ParseOnRampSet(log types.Log) (*RouterOnRampSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RouterOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*RouterOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RouterOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*RouterOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
