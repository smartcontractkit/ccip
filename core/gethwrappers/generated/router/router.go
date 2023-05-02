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

type RouterOffRampUpdate struct {
	SourceChainId uint64
	OffRamps      []common.Address
}

type RouterOnRampUpdate struct {
	DestChainId uint64
	OnRamp      common.Address
}

var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedNative\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientFeeTokenAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOffRamp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"name\":\"OffRampsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"name\":\"OffRampsRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"internalType\":\"structRouter.OnRampUpdate[]\",\"name\":\"onRampUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"internalType\":\"structRouter.OffRampUpdate[]\",\"name\":\"offRampUpdates\",\"type\":\"tuple[]\"}],\"name\":\"applyRampUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"}],\"name\":\"getOffRamps\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedNative\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"routeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedNative\",\"type\":\"address\"}],\"name\":\"setWrappedNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200262038038062002620833981016040819052620000349162000193565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000e8565b5050600280546001600160a01b0319166001600160a01b03939093169290921790915550620001c5565b336001600160a01b03821603620001425760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001a657600080fd5b81516001600160a01b0381168114620001be57600080fd5b9392505050565b61244b80620001d56000396000f3fe6080604052600436106100e85760003560e01c806379ba50971161008a578063a8d87a3b11610059578063a8d87a3b14610321578063e861e9071461036e578063f2fde38b14610399578063fbca3b74146103b957600080fd5b806379ba50971461025e5780638da5cb5b1461027357806396f4e9f9146102bf578063a48a9058146102d257600080fd5b806320487ded116100c657806320487ded146101ce57806349dea78a146101fc57806352cb60ca1461021e5780635607b3751461023e57600080fd5b80630d43a8e1146100ed578063181f5a77146101235780631d7a74a014610179575b600080fd5b3480156100f957600080fd5b5061010d61010836600461186e565b6103d9565b60405161011a9190611889565b60405180910390f35b34801561012f57600080fd5b5061016c6040518060400160405280600c81526020017f526f7574657220312e302e30000000000000000000000000000000000000000081525081565b60405161011a9190611959565b34801561018557600080fd5b506101be610194366004611999565b73ffffffffffffffffffffffffffffffffffffffff16600090815260056020526040902054151590565b604051901515815260200161011a565b3480156101da57600080fd5b506101ee6101e9366004611bca565b61045d565b60405190815260200161011a565b34801561020857600080fd5b5061021c610217366004611cc7565b6105b1565b005b34801561022a57600080fd5b5061021c610239366004611999565b61095b565b34801561024a57600080fd5b506101be610259366004611d9d565b6109aa565b34801561026a57600080fd5b5061021c610b54565b34801561027f57600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161011a565b6101ee6102cd366004611bca565b610c51565b3480156102de57600080fd5b506101be6102ed36600461186e565b67ffffffffffffffff1660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16151590565b34801561032d57600080fd5b5061029a61033c36600461186e565b67ffffffffffffffff1660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b34801561037a57600080fd5b5060025473ffffffffffffffffffffffffffffffffffffffff1661029a565b3480156103a557600080fd5b5061021c6103b4366004611999565b611115565b3480156103c557600080fd5b5061010d6103d436600461186e565b611129565b67ffffffffffffffff811660009081526004602090815260409182902080548351818402810184019094528084526060939283018282801561045157602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610426575b50505050509050919050565b606081015160009073ffffffffffffffffffffffffffffffffffffffff1661049e5760025473ffffffffffffffffffffffffffffffffffffffff1660608301525b67ffffffffffffffff831660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1680610516576040517fae236d9c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff851660048201526024015b60405180910390fd5b6040517f38724a9500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216906338724a9590610568908690600401611eeb565b602060405180830381865afa158015610585573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105a99190611efe565b949350505050565b6105b9611248565b60005b8381101561069d5760008585838181106105d8576105d8611f17565b9050604002018036038101906105ee9190611f46565b60208181018051835167ffffffffffffffff90811660009081526003855260409081902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff948516179055855193519051921682529394509216917f1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f23910160405180910390a25061069681611fa3565b90506105bc565b5060005b818110156109545760008383838181106106bd576106bd611f17565b90506020028101906106cf9190612002565b6106dd90602081019061186e565b905060005b67ffffffffffffffff82166000908152600460205260409020548110156107775760056000600460008567ffffffffffffffff168152602001908152602001600020838154811061073557610735611f17565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400181205561077081611fa3565b90506106e2565b5067ffffffffffffffff81166000818152600460205260409081902090517f8643e8d22aa348a6dc23fd35bc0643647644a683e8b151054cc76e02ea4abf9f916107c091612040565b60405180910390a267ffffffffffffffff811660009081526004602052604081206107ea91611794565b60008484848181106107fe576107fe611f17565b90506020028101906108109190612002565b61081e906020810190612091565b8080602002602001604051908101604052809392919081815260200183836020028082843760009201829052509394505050505b81518110156108d3578267ffffffffffffffff166005600084848151811061087c5761087c611f17565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550806108cc90611fa3565b9050610852565b5067ffffffffffffffff8216600090815260046020908152604090912082516108fe928401906117b2565b508167ffffffffffffffff167f34cf88e9ce732af3d4d56ddbea6ea3ecaf0b61ddf4c586d02808dbbec75015d3826040516109399190611889565b60405180910390a250508061094d90611fa3565b90506106a1565b5050505050565b610963611248565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60006109bc604086016020870161186e565b3360009081526005602052604090205415806109f157503360009081526005602052604090205467ffffffffffffffff821614155b15610a28576040517fd2316ede00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006385572ffb60e01b87604051602401610a43919061220d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915290508515610b3a578373ffffffffffffffffffffffffffffffffffffffff1681604051610aed9190612324565b6000604051808303816000865af19150503d8060008114610b2a576040519150601f19603f3d011682016040523d82523d6000602084013e610b2f565b606091505b505080935050610b4a565b610b4785856000846112cb565b92505b5050949350505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610bd5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161050d565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b67ffffffffffffffff821660009081526003602052604081205473ffffffffffffffffffffffffffffffffffffffff1680610cc4576040517fae236d9c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8516600482015260240161050d565b606083015160009073ffffffffffffffffffffffffffffffffffffffff16610e555760025473ffffffffffffffffffffffffffffffffffffffff90811660608601526040517f38724a95000000000000000000000000000000000000000000000000000000008152908316906338724a9590610d44908790600401611eeb565b602060405180830381865afa158015610d61573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d859190611efe565b905080341015610dc1576040517f07da6ee600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b349050836060015173ffffffffffffffffffffffffffffffffffffffff1663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b158015610e1057600080fd5b505af1158015610e24573d6000803e3d6000fd5b505050506060850151610e50915073ffffffffffffffffffffffffffffffffffffffff16308484611317565b610f4a565b3415610e8d576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f38724a9500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316906338724a9590610edf908790600401611eeb565b602060405180830381865afa158015610efc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f209190611efe565b6060850151909150610f4a9073ffffffffffffffffffffffffffffffffffffffff16338484611317565b60005b84604001515181101561107257600085604001518281518110610f7257610f72611f17565b6020908102919091010151516040517f5d86f14100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808316600483015291925061106191339190871690635d86f14190602401602060405180830381865afa158015610ff4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110189190612336565b8860400151858151811061102e5761102e611f17565b6020026020010151602001518473ffffffffffffffffffffffffffffffffffffffff16611317909392919063ffffffff16565b5061106b81611fa3565b9050610f4d565b506040517fa7d3e02f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83169063a7d3e02f906110c990879085903390600401612353565b6020604051808303816000875af11580156110e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061110c9190611efe565b95945050505050565b61111d611248565b611126816113b2565b50565b60606111638267ffffffffffffffff1660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16151590565b61117b57505060408051600081526020810190915290565b67ffffffffffffffff82166000908152600360205260408082205481517fd3c7c2c7000000000000000000000000000000000000000000000000000000008152915173ffffffffffffffffffffffffffffffffffffffff9091169263d3c7c2c792600480820193918290030181865afa1580156111fc573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526112429190810190612392565b92915050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146112c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161050d565b565b60005a6113888110156112dd57600080fd5b6113888103905085604082048203116112f557600080fd5b50833b61130157600080fd5b60008083516020850186888af195945050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526113ac9085906114a7565b50505050565b3373ffffffffffffffffffffffffffffffffffffffff821603611431576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161050d565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000611509826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166115b89092919063ffffffff16565b8051909150156115b357808060200190518101906115279190612421565b6115b3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161050d565b505050565b60606115c784846000856115d1565b90505b9392505050565b606082471015611663576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161050d565b843b6116cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161050d565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516116f49190612324565b60006040518083038185875af1925050503d8060008114611731576040519150601f19603f3d011682016040523d82523d6000602084013e611736565b606091505b5091509150610b47828286606083156117505750816115ca565b8251156117605782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050d9190611959565b5080546000825590600052602060002090810190611126919061183c565b82805482825590600052602060002090810192821561182c579160200282015b8281111561182c57825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9091161782556020909201916001909101906117d2565b5061183892915061183c565b5090565b5b80821115611838576000815560010161183d565b803567ffffffffffffffff8116811461186957600080fd5b919050565b60006020828403121561188057600080fd5b6115ca82611851565b6020808252825182820181905260009190848201906040850190845b818110156118d757835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016118a5565b50909695505050505050565b60005b838110156118fe5781810151838201526020016118e6565b838111156113ac5750506000910152565b600081518084526119278160208601602086016118e3565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006115ca602083018461190f565b73ffffffffffffffffffffffffffffffffffffffff8116811461112657600080fd5b80356118698161196c565b6000602082840312156119ab57600080fd5b81356115ca8161196c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715611a0857611a086119b6565b60405290565b60405160a0810167ffffffffffffffff81118282101715611a0857611a086119b6565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611a7857611a786119b6565b604052919050565b600082601f830112611a9157600080fd5b813567ffffffffffffffff811115611aab57611aab6119b6565b611adc60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611a31565b818152846020838601011115611af157600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115611b2857611b286119b6565b5060051b60200190565b600082601f830112611b4357600080fd5b81356020611b58611b5383611b0e565b611a31565b82815260069290921b84018101918181019086841115611b7757600080fd5b8286015b84811015611bbf5760408189031215611b945760008081fd5b611b9c6119e5565b8135611ba78161196c565b81528185013585820152835291830191604001611b7b565b509695505050505050565b60008060408385031215611bdd57600080fd5b611be683611851565b9150602083013567ffffffffffffffff80821115611c0357600080fd5b9084019060a08287031215611c1757600080fd5b611c1f611a0e565b823582811115611c2e57600080fd5b611c3a88828601611a80565b825250602083013582811115611c4f57600080fd5b611c5b88828601611a80565b602083015250604083013582811115611c7357600080fd5b611c7f88828601611b32565b604083015250611c916060840161198e565b6060820152608083013582811115611ca857600080fd5b611cb488828601611a80565b6080830152508093505050509250929050565b60008060008060408587031215611cdd57600080fd5b843567ffffffffffffffff80821115611cf557600080fd5b818701915087601f830112611d0957600080fd5b813581811115611d1857600080fd5b8860208260061b8501011115611d2d57600080fd5b602092830196509450908601359080821115611d4857600080fd5b818701915087601f830112611d5c57600080fd5b813581811115611d6b57600080fd5b8860208260051b8501011115611d8057600080fd5b95989497505060200194505050565b801515811461112657600080fd5b60008060008060808587031215611db357600080fd5b843567ffffffffffffffff811115611dca57600080fd5b850160a08188031215611ddc57600080fd5b93506020850135611dec81611d8f565b9250604085013591506060850135611e038161196c565b939692955090935050565b6000815160a08452611e2360a085018261190f565b905060208084015185830382870152611e3c838261190f565b60408681015188830389830152805180845290850195509092506000918401905b80831015611e9c578551805173ffffffffffffffffffffffffffffffffffffffff16835285015185830152948401946001929092019190830190611e5d565b5060608701519450611ec6606089018673ffffffffffffffffffffffffffffffffffffffff169052565b608087015194508781036080890152611edf818661190f565b98975050505050505050565b6020815260006115ca6020830184611e0e565b600060208284031215611f1057600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060408284031215611f5857600080fd5b6040516040810181811067ffffffffffffffff82111715611f7b57611f7b6119b6565b604052611f8783611851565b81526020830135611f978161196c565b60208201529392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611ffb577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc183360301811261203657600080fd5b9190910192915050565b6020808252825482820181905260008481528281209092916040850190845b818110156118d757835473ffffffffffffffffffffffffffffffffffffffff168352600193840193928501920161205f565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126120c657600080fd5b83018035915067ffffffffffffffff8211156120e157600080fd5b6020019150600581901b36038213156120f957600080fd5b9250929050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261213557600080fd5b830160208101925035905067ffffffffffffffff81111561215557600080fd5b8036038213156120f957600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8183526000602080850194508260005b858110156122025781356121d08161196c565b73ffffffffffffffffffffffffffffffffffffffff1687528183013583880152604096870196909101906001016121bd565b509495945050505050565b6020815281356020820152600061222660208401611851565b67ffffffffffffffff80821660408501526122446040860186612100565b925060a0606086015261225b60c086018483612164565b92505061226b6060860186612100565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808786030160808801526122a1858385612164565b9450608088013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18836030183126122da57600080fd5b602092880192830192359150838211156122f357600080fd5b8160061b360383131561230557600080fd5b8685030160a08701526123198482846121ad565b979650505050505050565b600082516120368184602087016118e3565b60006020828403121561234857600080fd5b81516115ca8161196c565b6060815260006123666060830186611e0e565b905083602083015273ffffffffffffffffffffffffffffffffffffffff83166040830152949350505050565b600060208083850312156123a557600080fd5b825167ffffffffffffffff8111156123bc57600080fd5b8301601f810185136123cd57600080fd5b80516123db611b5382611b0e565b81815260059190911b820183019083810190878311156123fa57600080fd5b928401925b828410156123195783516124128161196c565b825292840192908401906123ff565b60006020828403121561243357600080fd5b81516115ca81611d8f56fea164736f6c634300080f000a",
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

func (_Router *RouterCaller) IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "isOffRamp", offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Router *RouterSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _Router.Contract.IsOffRamp(&_Router.CallOpts, offRamp)
}

func (_Router *RouterCallerSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _Router.Contract.IsOffRamp(&_Router.CallOpts, offRamp)
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

func (_Router *RouterTransactor) ApplyRampUpdates(opts *bind.TransactOpts, onRampUpdates []RouterOnRampUpdate, offRampUpdates []RouterOffRampUpdate) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "applyRampUpdates", onRampUpdates, offRampUpdates)
}

func (_Router *RouterSession) ApplyRampUpdates(onRampUpdates []RouterOnRampUpdate, offRampUpdates []RouterOffRampUpdate) (*types.Transaction, error) {
	return _Router.Contract.ApplyRampUpdates(&_Router.TransactOpts, onRampUpdates, offRampUpdates)
}

func (_Router *RouterTransactorSession) ApplyRampUpdates(onRampUpdates []RouterOnRampUpdate, offRampUpdates []RouterOffRampUpdate) (*types.Transaction, error) {
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

type RouterOffRampsAddedIterator struct {
	Event *RouterOffRampsAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RouterOffRampsAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOffRampsAdded)
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
		it.Event = new(RouterOffRampsAdded)
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

func (it *RouterOffRampsAddedIterator) Error() error {
	return it.fail
}

func (it *RouterOffRampsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RouterOffRampsAdded struct {
	SourceChainId uint64
	OffRamps      []common.Address
	Raw           types.Log
}

func (_Router *RouterFilterer) FilterOffRampsAdded(opts *bind.FilterOpts, sourceChainId []uint64) (*RouterOffRampsAddedIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OffRampsAdded", sourceChainIdRule)
	if err != nil {
		return nil, err
	}
	return &RouterOffRampsAddedIterator{contract: _Router.contract, event: "OffRampsAdded", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOffRampsAdded(opts *bind.WatchOpts, sink chan<- *RouterOffRampsAdded, sourceChainId []uint64) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OffRampsAdded", sourceChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RouterOffRampsAdded)
				if err := _Router.contract.UnpackLog(event, "OffRampsAdded", log); err != nil {
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

func (_Router *RouterFilterer) ParseOffRampsAdded(log types.Log) (*RouterOffRampsAdded, error) {
	event := new(RouterOffRampsAdded)
	if err := _Router.contract.UnpackLog(event, "OffRampsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RouterOffRampsRemovedIterator struct {
	Event *RouterOffRampsRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RouterOffRampsRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOffRampsRemoved)
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
		it.Event = new(RouterOffRampsRemoved)
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

func (it *RouterOffRampsRemovedIterator) Error() error {
	return it.fail
}

func (it *RouterOffRampsRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RouterOffRampsRemoved struct {
	SourceChainId uint64
	OffRamps      []common.Address
	Raw           types.Log
}

func (_Router *RouterFilterer) FilterOffRampsRemoved(opts *bind.FilterOpts, sourceChainId []uint64) (*RouterOffRampsRemovedIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OffRampsRemoved", sourceChainIdRule)
	if err != nil {
		return nil, err
	}
	return &RouterOffRampsRemovedIterator{contract: _Router.contract, event: "OffRampsRemoved", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOffRampsRemoved(opts *bind.WatchOpts, sink chan<- *RouterOffRampsRemoved, sourceChainId []uint64) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OffRampsRemoved", sourceChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RouterOffRampsRemoved)
				if err := _Router.contract.UnpackLog(event, "OffRampsRemoved", log); err != nil {
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

func (_Router *RouterFilterer) ParseOffRampsRemoved(log types.Log) (*RouterOffRampsRemoved, error) {
	event := new(RouterOffRampsRemoved)
	if err := _Router.contract.UnpackLog(event, "OffRampsRemoved", log); err != nil {
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
	case _Router.abi.Events["OffRampsAdded"].ID:
		return _Router.ParseOffRampsAdded(log)
	case _Router.abi.Events["OffRampsRemoved"].ID:
		return _Router.ParseOffRampsRemoved(log)
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

func (RouterOffRampsAdded) Topic() common.Hash {
	return common.HexToHash("0x34cf88e9ce732af3d4d56ddbea6ea3ecaf0b61ddf4c586d02808dbbec75015d3")
}

func (RouterOffRampsRemoved) Topic() common.Hash {
	return common.HexToHash("0x8643e8d22aa348a6dc23fd35bc0643647644a683e8b151054cc76e02ea4abf9f")
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

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyRampUpdates(opts *bind.TransactOpts, onRampUpdates []RouterOnRampUpdate, offRampUpdates []RouterOffRampUpdate) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainId uint64, message ClientEVM2AnyMessage) (*types.Transaction, error)

	RouteMessage(opts *bind.TransactOpts, message ClientAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error)

	SetWrappedNative(opts *bind.TransactOpts, wrappedNative common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOffRampsAdded(opts *bind.FilterOpts, sourceChainId []uint64) (*RouterOffRampsAddedIterator, error)

	WatchOffRampsAdded(opts *bind.WatchOpts, sink chan<- *RouterOffRampsAdded, sourceChainId []uint64) (event.Subscription, error)

	ParseOffRampsAdded(log types.Log) (*RouterOffRampsAdded, error)

	FilterOffRampsRemoved(opts *bind.FilterOpts, sourceChainId []uint64) (*RouterOffRampsRemovedIterator, error)

	WatchOffRampsRemoved(opts *bind.WatchOpts, sink chan<- *RouterOffRampsRemoved, sourceChainId []uint64) (event.Subscription, error)

	ParseOffRampsRemoved(log types.Log) (*RouterOffRampsRemoved, error)

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
