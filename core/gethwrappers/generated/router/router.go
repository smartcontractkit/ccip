// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package router

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

type IRouterOffRampUpdate struct {
	SourceChainId uint64
	OffRamps      []common.Address
}

type IRouterOnRampUpdate struct {
	DestChainId uint64
	OnRamp      common.Address
}

var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedNative\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientFeeTokenAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOffRamp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"internalType\":\"structIRouter.OnRampUpdate[]\",\"name\":\"onRampUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"internalType\":\"structIRouter.OffRampUpdate[]\",\"name\":\"offRampUpdates\",\"type\":\"tuple[]\"}],\"name\":\"applyRampUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"}],\"name\":\"getOffRamps\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedNative\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"routeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedNative\",\"type\":\"address\"}],\"name\":\"setWrappedNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620022c0380380620022c0833981016040819052620000349162000193565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000e8565b5050600280546001600160a01b0319166001600160a01b03939093169290921790915550620001c5565b336001600160a01b03821603620001425760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001a657600080fd5b81516001600160a01b0381168114620001be57600080fd5b9392505050565b6120eb80620001d56000396000f3fe6080604052600436106100dd5760003560e01c80638da5cb5b1161007f578063a8d87a3b11610059578063a8d87a3b146102aa578063e861e907146102ea578063f2fde38b14610308578063fbca3b741461032857600080fd5b80638da5cb5b1461022357806396f4e9f914610255578063a48a90581461026857600080fd5b806349dea78a116100bb57806349dea78a1461019c57806352cb60ca146101be5780635607b375146101de57806379ba50971461020e57600080fd5b80630d43a8e1146100e2578063181f5a771461011857806320487ded1461016e575b600080fd5b3480156100ee57600080fd5b506101026100fd3660046115dd565b610348565b60405161010f91906115f8565b60405180910390f35b34801561012457600080fd5b506101616040518060400160405280600c81526020017f526f7574657220312e302e30000000000000000000000000000000000000000081525081565b60405161010f919061169d565b34801561017a57600080fd5b5061018e6101893660046118a8565b6103bf565b60405190815260200161010f565b3480156101a857600080fd5b506101bc6101b7366004611adf565b6104df565b005b3480156101ca57600080fd5b506101bc6101d9366004611bcc565b6108ec565b3480156101ea57600080fd5b506101fe6101f9366004611bf7565b61092e565b604051901515815260200161010f565b34801561021a57600080fd5b506101bc610aad565b34801561022f57600080fd5b506000546001600160a01b03165b6040516001600160a01b03909116815260200161010f565b61018e6102633660046118a8565b610b76565b34801561027457600080fd5b506101fe6102833660046115dd565b67ffffffffffffffff166000908152600360205260409020546001600160a01b0316151590565b3480156102b657600080fd5b5061023d6102c53660046115dd565b67ffffffffffffffff166000908152600360205260409020546001600160a01b031690565b3480156102f657600080fd5b506002546001600160a01b031661023d565b34801561031457600080fd5b506101bc610323366004611bcc565b610fb8565b34801561033457600080fd5b506101026103433660046115dd565b610fcc565b67ffffffffffffffff81166000908152600460209081526040918290208054835181840281018401909452808452606093928301828280156103b357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610395575b50505050509050919050565b60608101516000906001600160a01b03166103e6576002546001600160a01b031660608301525b67ffffffffffffffff83166000908152600360205260409020546001600160a01b031680610451576040517fae236d9c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff851660048201526024015b60405180910390fd5b6040517f38724a950000000000000000000000000000000000000000000000000000000081526001600160a01b038216906338724a9590610496908690600401611d2b565b602060405180830381865afa1580156104b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d79190611d3e565b949350505050565b6104e76110b3565b60005b825181101561060a5782818151811061050557610505611d57565b6020026020010151602001516003600085848151811061052757610527611d57565b60200260200101516000015167ffffffffffffffff16815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555082818151811061058157610581611d57565b60200260200101516000015167ffffffffffffffff167f1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f238483815181106105ca576105ca611d57565b6020026020010151602001516040516105f291906001600160a01b0391909116815260200190565b60405180910390a261060381611d86565b90506104ea565b5060005b81518110156108e757600082828151811061062b5761062b611d57565b602002602001015160000151905060005b67ffffffffffffffff82166000908152600460205260409020548110156107385760056000600460008567ffffffffffffffff168152602001908152602001600020838154811061068f5761068f611d57565b60009182526020808320909101546001600160a01b03168352828101939093526040918201812081905567ffffffffffffffff851680825260049093522080547fa823809efda3ba66c873364eec120fa0923d9fabda73bc97dd5663341e2d9bcb91908490811061070257610702611d57565b600091825260209182902001546040516001600160a01b0390911681520160405180910390a261073181611d86565b905061063c565b5067ffffffffffffffff8116600090815260046020526040812061075b91611510565b60005b83838151811061077057610770611d57565b60200260200101516020015151811015610885578167ffffffffffffffff16600560008686815181106107a5576107a5611d57565b60200260200101516020015184815181106107c2576107c2611d57565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020819055508167ffffffffffffffff167fa4bdf64ebdf3316320601a081916a75aa144bcef6c4beeb0e9fb1982cacc6b9485858151811061082c5761082c611d57565b602002602001015160200151838151811061084957610849611d57565b602002602001015160405161086d91906001600160a01b0391909116815260200190565b60405180910390a261087e81611d86565b905061075e565b5082828151811061089857610898611d57565b602002602001015160200151600460008367ffffffffffffffff16815260200190815260200160002090805190602001906108d492919061152e565b5050806108e090611d86565b905061060e565b505050565b6108f46110b3565b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b600061094060408601602087016115dd565b33600090815260056020526040902054158061097557503360009081526005602052604090205467ffffffffffffffff821614155b156109ac576040517fd2316ede00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006385572ffb60e01b876040516024016109c79190611ece565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915290508515610a9357836001600160a01b031681604051610a469190611fc7565b6000604051808303816000865af19150503d8060008114610a83576040519150601f19603f3d011682016040523d82523d6000602084013e610a88565b606091505b505080935050610aa3565b610aa0858560008461110f565b92505b5050949350505050565b6001546001600160a01b03163314610b075760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610448565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b67ffffffffffffffff82166000908152600360205260408120546001600160a01b031680610bdc576040517fae236d9c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610448565b60608301516000906001600160a01b0316610d39576002546001600160a01b0390811660608601526040517f38724a95000000000000000000000000000000000000000000000000000000008152908316906338724a9590610c42908790600401611d2b565b602060405180830381865afa158015610c5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c839190611d3e565b905080341015610cbf576040517f07da6ee600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b34905083606001516001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b158015610d0157600080fd5b505af1158015610d15573d6000803e3d6000fd5b505050506060850151610d3491506001600160a01b031630848461115b565b610e14565b3415610d71576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f38724a950000000000000000000000000000000000000000000000000000000081526001600160a01b038316906338724a9590610db6908790600401611d2b565b602060405180830381865afa158015610dd3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df79190611d3e565b6060850151909150610e14906001600160a01b031633848461115b565b60005b846040015151811015610f2257600085604001518281518110610e3c57610e3c611d57565b6020908102919091010151516040517f5d86f1410000000000000000000000000000000000000000000000000000000081526001600160a01b038083166004830152919250610f1191339190871690635d86f14190602401602060405180830381865afa158015610eb1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed59190611fe3565b88604001518581518110610eeb57610eeb611d57565b602002602001015160200151846001600160a01b031661115b909392919063ffffffff16565b50610f1b81611d86565b9050610e17565b506040517fa7d3e02f0000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063a7d3e02f90610f6c90879085903390600401612000565b6020604051808303816000875af1158015610f8b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610faf9190611d3e565b95945050505050565b610fc06110b3565b610fc9816111e9565b50565b6060610ff98267ffffffffffffffff166000908152600360205260409020546001600160a01b0316151590565b61101157505060408051600081526020810190915290565b67ffffffffffffffff82166000908152600360205260408082205481517fd3c7c2c700000000000000000000000000000000000000000000000000000000815291516001600160a01b039091169263d3c7c2c792600480820193918290030181865afa158015611085573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110ad9190810190612032565b92915050565b6000546001600160a01b0316331461110d5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610448565b565b60005a61138881101561112157600080fd5b61138881039050856040820482031161113957600080fd5b50833b61114557600080fd5b60008083516020850186888af195945050505050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526111e39085906112aa565b50505050565b336001600160a01b038216036112415760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610448565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006112ff826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661138f9092919063ffffffff16565b8051909150156108e7578080602001905181019061131d91906120c1565b6108e75760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610448565b606061139e84846000856113a8565b90505b9392505050565b6060824710156114205760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610448565b843b61146e5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610448565b600080866001600160a01b0316858760405161148a9190611fc7565b60006040518083038185875af1925050503d80600081146114c7576040519150601f19603f3d011682016040523d82523d6000602084013e6114cc565b606091505b5091509150610aa0828286606083156114e65750816113a1565b8251156114f65782518084602001fd5b8160405162461bcd60e51b8152600401610448919061169d565b5080546000825590600052602060002090810190610fc991906115ab565b82805482825590600052602060002090810192821561159b579160200282015b8281111561159b57825182547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0390911617825560209092019160019091019061154e565b506115a79291506115ab565b5090565b5b808211156115a757600081556001016115ac565b803567ffffffffffffffff811681146115d857600080fd5b919050565b6000602082840312156115ef57600080fd5b6113a1826115c0565b6020808252825182820181905260009190848201906040850190845b818110156116395783516001600160a01b031683529284019291840191600101611614565b50909695505050505050565b60005b83811015611660578181015183820152602001611648565b838111156111e35750506000910152565b60008151808452611689816020860160208601611645565b601f01601f19169290920160200192915050565b6020815260006113a16020830184611671565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715611702576117026116b0565b60405290565b60405160a0810167ffffffffffffffff81118282101715611702576117026116b0565b604051601f8201601f1916810167ffffffffffffffff81118282101715611754576117546116b0565b604052919050565b600082601f83011261176d57600080fd5b813567ffffffffffffffff811115611787576117876116b0565b61179a6020601f19601f8401160161172b565b8181528460208386010111156117af57600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff8211156117e6576117e66116b0565b5060051b60200190565b6001600160a01b0381168114610fc957600080fd5b80356115d8816117f0565b600082601f83011261182157600080fd5b81356020611836611831836117cc565b61172b565b82815260069290921b8401810191818101908684111561185557600080fd5b8286015b8481101561189d57604081890312156118725760008081fd5b61187a6116df565b8135611885816117f0565b81528185013585820152835291830191604001611859565b509695505050505050565b600080604083850312156118bb57600080fd5b6118c4836115c0565b9150602083013567ffffffffffffffff808211156118e157600080fd5b9084019060a082870312156118f557600080fd5b6118fd611708565b82358281111561190c57600080fd5b6119188882860161175c565b82525060208301358281111561192d57600080fd5b6119398882860161175c565b60208301525060408301358281111561195157600080fd5b61195d88828601611810565b60408301525061196f60608401611805565b606082015260808301358281111561198657600080fd5b6119928882860161175c565b6080830152508093505050509250929050565b600082601f8301126119b657600080fd5b813560206119c6611831836117cc565b828152600592831b85018201928282019190878511156119e557600080fd5b8387015b85811015611ad257803567ffffffffffffffff80821115611a0a5760008081fd5b818a019150604080601f19848e03011215611a255760008081fd5b611a2d6116df565b611a388985016115c0565b81528184013583811115611a4c5760008081fd5b8085019450508c603f850112611a6457600092508283fd5b888401359250611a76611831846117cc565b83815292861b8401820192898101908e851115611a935760008081fd5b948301945b84861015611abd5785359350611aad846117f0565b838252948a0194908a0190611a98565b828b01525087525050509284019284016119e9565b5090979650505050505050565b6000806040808486031215611af357600080fd5b833567ffffffffffffffff80821115611b0b57600080fd5b818601915086601f830112611b1f57600080fd5b81356020611b2f611831836117cc565b82815260069290921b8401810191818101908a841115611b4e57600080fd5b948201945b83861015611b9e5786868c031215611b6b5760008081fd5b611b736116df565b611b7c876115c0565b815283870135611b8b816117f0565b8185015282529486019490820190611b53565b97505087013593505080831115611bb457600080fd5b5050611bc2858286016119a5565b9150509250929050565b600060208284031215611bde57600080fd5b81356113a1816117f0565b8015158114610fc957600080fd5b60008060008060808587031215611c0d57600080fd5b843567ffffffffffffffff811115611c2457600080fd5b850160a08188031215611c3657600080fd5b93506020850135611c4681611be9565b9250604085013591506060850135611c5d816117f0565b939692955090935050565b6000815160a08452611c7d60a0850182611671565b905060208084015185830382870152611c968382611671565b60408681015188830389830152805180845290850195509092506000918401905b80831015611ce957855180516001600160a01b0316835285015185830152948401946001929092019190830190611cb7565b5060608701519450611d0660608901866001600160a01b03169052565b608087015194508781036080890152611d1f8186611671565b98975050505050505050565b6020815260006113a16020830184611c68565b600060208284031215611d5057600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611dde577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611e1a57600080fd5b830160208101925035905067ffffffffffffffff811115611e3a57600080fd5b803603821315611e4957600080fd5b9250929050565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b8183526000602080850194508260005b85811015611ec3578135611e9e816117f0565b6001600160a01b03168752818301358388015260409687019690910190600101611e8b565b509495945050505050565b60208152813560208201526000611ee7602084016115c0565b67ffffffffffffffff8082166040850152611f056040860186611de5565b925060a06060860152611f1c60c086018483611e50565b925050611f2c6060860186611de5565b601f1980878603016080880152611f44858385611e50565b9450608088013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018312611f7d57600080fd5b60209288019283019235915083821115611f9657600080fd5b8160061b3603831315611fa857600080fd5b8685030160a0870152611fbc848284611e7b565b979650505050505050565b60008251611fd9818460208701611645565b9190910192915050565b600060208284031215611ff557600080fd5b81516113a1816117f0565b6060815260006120136060830186611c68565b90508360208301526001600160a01b0383166040830152949350505050565b6000602080838503121561204557600080fd5b825167ffffffffffffffff81111561205c57600080fd5b8301601f8101851361206d57600080fd5b805161207b611831826117cc565b81815260059190911b8201830190838101908783111561209a57600080fd5b928401925b82841015611fbc5783516120b2816117f0565b8252928401929084019061209f565b6000602082840312156120d357600080fd5b81516113a181611be956fea164736f6c634300080f000a",
}

var RouterABI = RouterMetaData.ABI

var RouterBin = RouterMetaData.Bin

func DeployRouter(auth *bind.TransactOpts, backend bind.ContractBackend, wrappedNative common.Address) (common.Address, *types.Transaction, *Router, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RouterBin), backend, wrappedNative)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
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
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

func (_Router *RouterCaller) GetFee(opts *bind.CallOpts, destinationChainId uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getFee", destinationChainId, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Router *RouterSession) GetFee(destinationChainId uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _Router.Contract.GetFee(&_Router.CallOpts, destinationChainId, message)
}

func (_Router *RouterCallerSession) GetFee(destinationChainId uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _Router.Contract.GetFee(&_Router.CallOpts, destinationChainId, message)
}

func (_Router *RouterCaller) GetOffRamps(opts *bind.CallOpts, sourceChainId uint64) ([]common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getOffRamps", sourceChainId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Router *RouterSession) GetOffRamps(sourceChainId uint64) ([]common.Address, error) {
	return _Router.Contract.GetOffRamps(&_Router.CallOpts, sourceChainId)
}

func (_Router *RouterCallerSession) GetOffRamps(sourceChainId uint64) ([]common.Address, error) {
	return _Router.Contract.GetOffRamps(&_Router.CallOpts, sourceChainId)
}

func (_Router *RouterCaller) GetOnRamp(opts *bind.CallOpts, destChainId uint64) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getOnRamp", destChainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Router *RouterSession) GetOnRamp(destChainId uint64) (common.Address, error) {
	return _Router.Contract.GetOnRamp(&_Router.CallOpts, destChainId)
}

func (_Router *RouterCallerSession) GetOnRamp(destChainId uint64) (common.Address, error) {
	return _Router.Contract.GetOnRamp(&_Router.CallOpts, destChainId)
}

func (_Router *RouterCaller) GetSupportedTokens(opts *bind.CallOpts, chainId uint64) ([]common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getSupportedTokens", chainId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Router *RouterSession) GetSupportedTokens(chainId uint64) ([]common.Address, error) {
	return _Router.Contract.GetSupportedTokens(&_Router.CallOpts, chainId)
}

func (_Router *RouterCallerSession) GetSupportedTokens(chainId uint64) ([]common.Address, error) {
	return _Router.Contract.GetSupportedTokens(&_Router.CallOpts, chainId)
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

func (_Router *RouterCaller) IsChainSupported(opts *bind.CallOpts, chainId uint64) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "isChainSupported", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Router *RouterSession) IsChainSupported(chainId uint64) (bool, error) {
	return _Router.Contract.IsChainSupported(&_Router.CallOpts, chainId)
}

func (_Router *RouterCallerSession) IsChainSupported(chainId uint64) (bool, error) {
	return _Router.Contract.IsChainSupported(&_Router.CallOpts, chainId)
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

func (_Router *RouterTransactor) ApplyRampUpdates(opts *bind.TransactOpts, onRampUpdates []IRouterOnRampUpdate, offRampUpdates []IRouterOffRampUpdate) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "applyRampUpdates", onRampUpdates, offRampUpdates)
}

func (_Router *RouterSession) ApplyRampUpdates(onRampUpdates []IRouterOnRampUpdate, offRampUpdates []IRouterOffRampUpdate) (*types.Transaction, error) {
	return _Router.Contract.ApplyRampUpdates(&_Router.TransactOpts, onRampUpdates, offRampUpdates)
}

func (_Router *RouterTransactorSession) ApplyRampUpdates(onRampUpdates []IRouterOnRampUpdate, offRampUpdates []IRouterOffRampUpdate) (*types.Transaction, error) {
	return _Router.Contract.ApplyRampUpdates(&_Router.TransactOpts, onRampUpdates, offRampUpdates)
}

func (_Router *RouterTransactor) CcipSend(opts *bind.TransactOpts, destinationChainId uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "ccipSend", destinationChainId, message)
}

func (_Router *RouterSession) CcipSend(destinationChainId uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _Router.Contract.CcipSend(&_Router.TransactOpts, destinationChainId, message)
}

func (_Router *RouterTransactorSession) CcipSend(destinationChainId uint64, message ClientEVM2AnyMessage) (*types.Transaction, error) {
	return _Router.Contract.CcipSend(&_Router.TransactOpts, destinationChainId, message)
}

func (_Router *RouterTransactor) RouteMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "routeMessage", message, manualExecution, gasLimit, receiver)
}

func (_Router *RouterSession) RouteMessage(message ClientAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.RouteMessage(&_Router.TransactOpts, message, manualExecution, gasLimit, receiver)
}

func (_Router *RouterTransactorSession) RouteMessage(message ClientAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.RouteMessage(&_Router.TransactOpts, message, manualExecution, gasLimit, receiver)
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
	SourceChainId uint64
	OffRamp       common.Address
	Raw           types.Log
}

func (_Router *RouterFilterer) FilterOffRampAdded(opts *bind.FilterOpts, sourceChainId []uint64) (*RouterOffRampAddedIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OffRampAdded", sourceChainIdRule)
	if err != nil {
		return nil, err
	}
	return &RouterOffRampAddedIterator{contract: _Router.contract, event: "OffRampAdded", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *RouterOffRampAdded, sourceChainId []uint64) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OffRampAdded", sourceChainIdRule)
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
	SourceChainId uint64
	OffRamp       common.Address
	Raw           types.Log
}

func (_Router *RouterFilterer) FilterOffRampRemoved(opts *bind.FilterOpts, sourceChainId []uint64) (*RouterOffRampRemovedIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OffRampRemoved", sourceChainIdRule)
	if err != nil {
		return nil, err
	}
	return &RouterOffRampRemovedIterator{contract: _Router.contract, event: "OffRampRemoved", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *RouterOffRampRemoved, sourceChainId []uint64) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OffRampRemoved", sourceChainIdRule)
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
	DestChainId uint64
	OnRamp      common.Address
	Raw         types.Log
}

func (_Router *RouterFilterer) FilterOnRampSet(opts *bind.FilterOpts, destChainId []uint64) (*RouterOnRampSetIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OnRampSet", destChainIdRule)
	if err != nil {
		return nil, err
	}
	return &RouterOnRampSetIterator{contract: _Router.contract, event: "OnRampSet", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *RouterOnRampSet, destChainId []uint64) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OnRampSet", destChainIdRule)
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

func (_Router *Router) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
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

type RouterInterface interface {
	GetFee(opts *bind.CallOpts, destinationChainId uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetOffRamps(opts *bind.CallOpts, sourceChainId uint64) ([]common.Address, error)

	GetOnRamp(opts *bind.CallOpts, destChainId uint64) (common.Address, error)

	GetSupportedTokens(opts *bind.CallOpts, chainId uint64) ([]common.Address, error)

	GetWrappedNative(opts *bind.CallOpts) (common.Address, error)

	IsChainSupported(opts *bind.CallOpts, chainId uint64) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyRampUpdates(opts *bind.TransactOpts, onRampUpdates []IRouterOnRampUpdate, offRampUpdates []IRouterOffRampUpdate) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainId uint64, message ClientEVM2AnyMessage) (*types.Transaction, error)

	RouteMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error)

	SetWrappedNative(opts *bind.TransactOpts, wrappedNative common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOffRampAdded(opts *bind.FilterOpts, sourceChainId []uint64) (*RouterOffRampAddedIterator, error)

	WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *RouterOffRampAdded, sourceChainId []uint64) (event.Subscription, error)

	ParseOffRampAdded(log types.Log) (*RouterOffRampAdded, error)

	FilterOffRampRemoved(opts *bind.FilterOpts, sourceChainId []uint64) (*RouterOffRampRemovedIterator, error)

	WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *RouterOffRampRemoved, sourceChainId []uint64) (event.Subscription, error)

	ParseOffRampRemoved(log types.Log) (*RouterOffRampRemoved, error)

	FilterOnRampSet(opts *bind.FilterOpts, destChainId []uint64) (*RouterOnRampSetIterator, error)

	WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *RouterOnRampSet, destChainId []uint64) (event.Subscription, error)

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
