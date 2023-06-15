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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedNative\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedToSendValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFeeTokenAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"InvalidRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OffRampMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOffRamp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"MessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_RET_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"internalType\":\"structRouter.OnRamp[]\",\"name\":\"onRampUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"internalType\":\"structRouter.OffRamp[]\",\"name\":\"offRampRemoves\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"internalType\":\"structRouter.OffRamp[]\",\"name\":\"offRampAdds\",\"type\":\"tuple[]\"}],\"name\":\"applyRampUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRamps\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"internalType\":\"structRouter.OffRamp[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedNative\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"routeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedNative\",\"type\":\"address\"}],\"name\":\"setWrappedNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002bbc38038062002bbc833981016040819052620000349162000193565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000e8565b5050600280546001600160a01b0319166001600160a01b03939093169290921790915550620001c5565b336001600160a01b03821603620001425760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001a657600080fd5b81516001600160a01b0381168114620001be57600080fd5b9392505050565b6129e780620001d56000396000f3fe60806040526004361061010e5760003560e01c80638da5cb5b116100a5578063a8d87a3b11610074578063e861e90711610059578063e861e907146103da578063f2fde38b14610405578063fbca3b741461042557600080fd5b8063a8d87a3b1461036d578063da5fcac8146103ba57600080fd5b80638da5cb5b1461028d57806396f4e9f9146102d9578063a40e69c7146102ec578063a48a90581461030e57600080fd5b806352cb60ca116100e157806352cb60ca1461020e5780635f3e849f14610230578063787350e31461025057806379ba50971461027857600080fd5b8063181f5a77146101135780631d7a74a01461017257806320487ded146101b25780633cf97983146101e0575b600080fd5b34801561011f57600080fd5b5061015c6040518060400160405280600c81526020017f526f7574657220312e302e30000000000000000000000000000000000000000081525081565b6040516101699190611e95565b60405180910390f35b34801561017e57600080fd5b5061019261018d366004611eda565b610452565b60408051921515835267ffffffffffffffff909116602083015201610169565b3480156101be57600080fd5b506101d26101cd366004612123565b61046e565b604051908152602001610169565b3480156101ec57600080fd5b506102006101fb366004612220565b6105c4565b604051610169929190612298565b34801561021a57600080fd5b5061022e610229366004611eda565b610756565b005b34801561023c57600080fd5b5061022e61024b3660046122b3565b6107a5565b34801561025c57600080fd5b50610265608481565b60405161ffff9091168152602001610169565b34801561028457600080fd5b5061022e6108f3565b34801561029957600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610169565b6101d26102e7366004612123565b6109f0565b3480156102f857600080fd5b50610301610eb3565b60405161016991906122f4565b34801561031a57600080fd5b5061035d610329366004612363565b67ffffffffffffffff1660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16151590565b6040519015158152602001610169565b34801561037957600080fd5b506102b4610388366004612363565b67ffffffffffffffff1660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b3480156103c657600080fd5b5061022e6103d53660046123c3565b610faa565b3480156103e657600080fd5b5060025473ffffffffffffffffffffffffffffffffffffffff166102b4565b34801561041157600080fd5b5061022e610420366004611eda565b6112fa565b34801561043157600080fd5b50610445610440366004612363565b61130e565b604051610169919061245d565b6000808080610462600486611427565b90969095509350505050565b606081015160009073ffffffffffffffffffffffffffffffffffffffff166104af5760025473ffffffffffffffffffffffffffffffffffffffff1660608301525b67ffffffffffffffff831660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1680610527576040517fae236d9c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff851660048201526024015b60405180910390fd5b6040517f38724a9500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216906338724a9590610579908690600401612594565b602060405180830381865afa158015610596573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ba91906125a7565b9150505b92915050565b600060606105d86040870160208801612363565b6000806105e6600433611427565b9150915081158061060b57508067ffffffffffffffff168367ffffffffffffffff1614155b15610642576040517fd2316ede00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806106e68a8a8a6385572ffb60e01b8f60405160240161066491906126cd565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261145b565b90925090507fbd2f9ed173f6f0237225df18b4721b82fb4e79583a71c0bd731669b908de43338b3561071e60408e0160208f01612363565b6040805192835267ffffffffffffffff9091166020830152339082015260600160405180910390a1909a909950975050505050505050565b61075e6114e8565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6107ad6114e8565b73ffffffffffffffffffffffffffffffffffffffff8216610812576040517f26a78f8f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015260240161051e565b73ffffffffffffffffffffffffffffffffffffffff83166108cd5760008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114610887576040519150601f19603f3d011682016040523d82523d6000602084013e61088c565b606091505b50509050806108c7576040517fe417b80b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050565b6108ee73ffffffffffffffffffffffffffffffffffffffff8416838361156b565b505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610974576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161051e565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b67ffffffffffffffff821660009081526003602052604081205473ffffffffffffffffffffffffffffffffffffffff1680610a63576040517fae236d9c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8516600482015260240161051e565b606083015160009073ffffffffffffffffffffffffffffffffffffffff16610bf35760025473ffffffffffffffffffffffffffffffffffffffff90811660608601526040517f38724a95000000000000000000000000000000000000000000000000000000008152908316906338724a9590610ae3908790600401612594565b602060405180830381865afa158015610b00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b2491906125a7565b905080341015610b60576040517f07da6ee600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b349050836060015173ffffffffffffffffffffffffffffffffffffffff1663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b158015610baf57600080fd5b505af1158015610bc3573d6000803e3d6000fd5b505050506060850151610bee915073ffffffffffffffffffffffffffffffffffffffff16838361156b565b610ce8565b3415610c2b576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f38724a9500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316906338724a9590610c7d908790600401612594565b602060405180830381865afa158015610c9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cbe91906125a7565b6060850151909150610ce89073ffffffffffffffffffffffffffffffffffffffff1633848461163f565b60005b846040015151811015610e1057600085604001518281518110610d1057610d106127d9565b6020908102919091010151516040517f5d86f14100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8083166004830152919250610dff91339190871690635d86f14190602401602060405180830381865afa158015610d92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610db69190612808565b88604001518581518110610dcc57610dcc6127d9565b6020026020010151602001518473ffffffffffffffffffffffffffffffffffffffff1661163f909392919063ffffffff16565b50610e0981612854565b9050610ceb565b506040517fa7d3e02f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83169063a7d3e02f90610e679087908590339060040161288c565b6020604051808303816000875af1158015610e86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eaa91906125a7565b95945050505050565b60606000610ec1600461169d565b67ffffffffffffffff811115610ed957610ed9611f0f565b604051908082528060200260200182016040528015610f1e57816020015b6040805180820190915260008082526020820152815260200190600190039081610ef75790505b50905060005b8151811015610fa457600080610f3b6004846116a8565b9150915060405180604001604052808267ffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff16815250848481518110610f8657610f866127d9565b6020026020010181905250505080610f9d90612854565b9050610f24565b50919050565b610fb26114e8565b60005b858110156110db5736878783818110610fd057610fd06127d9565b9050604002019050806020016020810190610feb9190611eda565b60036000610ffc6020850185612363565b67ffffffffffffffff16815260208082019290925260400160002080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff939093169290921790915561106790820182612363565b67ffffffffffffffff167f1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f236110a26040840160208501611eda565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a2506110d481612854565b9050610fb5565b5060005b838110156112075760008585838181106110fb576110fb6127d9565b6111119260206040909202019081019150612363565b90506000868684818110611127576111276127d9565b905060400201602001602081019061113f9190611eda565b905067ffffffffffffffff82166111576004836116b7565b1461118e576040517f5737980600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111996004826116e0565b156111f45760405173ffffffffffffffffffffffffffffffffffffffff8216815267ffffffffffffffff8316907fa823809efda3ba66c873364eec120fa0923d9fabda73bc97dd5663341e2d9bcb9060200160405180910390a25b50508061120090612854565b90506110df565b5060005b818110156112f1576000838383818110611227576112276127d9565b61123d9260206040909202019081019150612363565b90506000848484818110611253576112536127d9565b905060400201602001602081019061126b9190611eda565b905061128360048267ffffffffffffffff8516611702565b156112de5760405173ffffffffffffffffffffffffffffffffffffffff8216815267ffffffffffffffff8316907fa4bdf64ebdf3316320601a081916a75aa144bcef6c4beeb0e9fb1982cacc6b949060200160405180910390a25b5050806112ea90612854565b905061120b565b50505050505050565b6113026114e8565b61130b8161172d565b50565b60606113488267ffffffffffffffff1660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16151590565b61136057505060408051600081526020810190915290565b67ffffffffffffffff82166000908152600360205260408082205481517fd3c7c2c7000000000000000000000000000000000000000000000000000000008152915173ffffffffffffffffffffffffffffffffffffffff9091169263d3c7c2c792600480820193918290030181865afa1580156113e1573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526105be91908101906128cb565b600080808061144c8673ffffffffffffffffffffffffffffffffffffffff8716611822565b909450925050505b9250929050565b60408051608480825260c08201909252600091606091839160208201818036833701905050905060005a8881101561149257600080fd5b88900360408104810388106114a657600080fd5b50853b6114b257600080fd5b60008086516020880160008a8cf190503d60848111156114d0575060845b808352806000602085013e5097909650945050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314611569576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161051e565b565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526108ee9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261185c565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526108c79085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016115bd565b60006105be82611968565b600080808061144c8686611973565b60006116d98373ffffffffffffffffffffffffffffffffffffffff841661199e565b9392505050565b60006116d98373ffffffffffffffffffffffffffffffffffffffff8416611a28565b60006117258473ffffffffffffffffffffffffffffffffffffffff851684611a45565b949350505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036117ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161051e565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000818152600283016020526040812054819080611851576118448585611a62565b9250600091506114549050565b600192509050611454565b60006118be826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16611a6e9092919063ffffffff16565b8051909150156108ee57808060200190518101906118dc919061295a565b6108ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161051e565b60006105be82611a7d565b600080806119818585611a87565b600081815260029690960160205260409095205494959350505050565b6000818152600283016020526040812054801515806119c257506119c28484611a62565b6116d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b65790000604482015260640161051e565b600081815260028301602052604081208190556116d98383611a93565b600082815260028401602052604081208290556117258484611a9f565b60006116d98383611aab565b60606117258484600085611ac3565b60006105be825490565b60006116d98383611bdc565b60006116d98383611c06565b60006116d98383611cf9565b600081815260018301602052604081205415156116d9565b606082471015611b55576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161051e565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611b7e919061297c565b60006040518083038185875af1925050503d8060008114611bbb576040519150601f19603f3d011682016040523d82523d6000602084013e611bc0565b606091505b5091509150611bd187838387611d48565b979650505050505050565b6000826000018281548110611bf357611bf36127d9565b9060005260206000200154905092915050565b60008181526001830160205260408120548015611cef576000611c2a600183612998565b8554909150600090611c3e90600190612998565b9050818114611ca3576000866000018281548110611c5e57611c5e6127d9565b9060005260206000200154905080876000018481548110611c8157611c816127d9565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080611cb457611cb46129ab565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506105be565b60009150506105be565b6000818152600183016020526040812054611d40575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556105be565b5060006105be565b60608315611dde578251600003611dd75773ffffffffffffffffffffffffffffffffffffffff85163b611dd7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161051e565b5081611725565b6117258383815115611df35781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051e9190611e95565b60005b83811015611e42578181015183820152602001611e2a565b50506000910152565b60008151808452611e63816020860160208601611e27565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006116d96020830184611e4b565b73ffffffffffffffffffffffffffffffffffffffff8116811461130b57600080fd5b8035611ed581611ea8565b919050565b600060208284031215611eec57600080fd5b81356116d981611ea8565b803567ffffffffffffffff81168114611ed557600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715611f6157611f61611f0f565b60405290565b60405160a0810167ffffffffffffffff81118282101715611f6157611f61611f0f565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611fd157611fd1611f0f565b604052919050565b600082601f830112611fea57600080fd5b813567ffffffffffffffff81111561200457612004611f0f565b61203560207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611f8a565b81815284602083860101111561204a57600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff82111561208157612081611f0f565b5060051b60200190565b600082601f83011261209c57600080fd5b813560206120b16120ac83612067565b611f8a565b82815260069290921b840181019181810190868411156120d057600080fd5b8286015b8481101561211857604081890312156120ed5760008081fd5b6120f5611f3e565b813561210081611ea8565b815281850135858201528352918301916040016120d4565b509695505050505050565b6000806040838503121561213657600080fd5b61213f83611ef7565b9150602083013567ffffffffffffffff8082111561215c57600080fd5b9084019060a0828703121561217057600080fd5b612178611f67565b82358281111561218757600080fd5b61219388828601611fd9565b8252506020830135828111156121a857600080fd5b6121b488828601611fd9565b6020830152506040830135828111156121cc57600080fd5b6121d88882860161208b565b6040830152506121ea60608401611eca565b606082015260808301358281111561220157600080fd5b61220d88828601611fd9565b6080830152508093505050509250929050565b6000806000806080858703121561223657600080fd5b843567ffffffffffffffff81111561224d57600080fd5b850160a0818803121561225f57600080fd5b9350602085013561ffff8116811461227657600080fd5b925060408501359150606085013561228d81611ea8565b939692955090935050565b82151581526040602082015260006117256040830184611e4b565b6000806000606084860312156122c857600080fd5b83356122d381611ea8565b925060208401356122e381611ea8565b929592945050506040919091013590565b602080825282518282018190526000919060409081850190868401855b82811015612356578151805167ffffffffffffffff16855286015173ffffffffffffffffffffffffffffffffffffffff16868501529284019290850190600101612311565b5091979650505050505050565b60006020828403121561237557600080fd5b6116d982611ef7565b60008083601f84011261239057600080fd5b50813567ffffffffffffffff8111156123a857600080fd5b6020830191508360208260061b850101111561145457600080fd5b600080600080600080606087890312156123dc57600080fd5b863567ffffffffffffffff808211156123f457600080fd5b6124008a838b0161237e565b9098509650602089013591508082111561241957600080fd5b6124258a838b0161237e565b9096509450604089013591508082111561243e57600080fd5b5061244b89828a0161237e565b979a9699509497509295939492505050565b6020808252825182820181905260009190848201906040850190845b818110156124ab57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101612479565b50909695505050505050565b6000815160a084526124cc60a0850182611e4b565b9050602080840151858303828701526124e58382611e4b565b60408681015188830389830152805180845290850195509092506000918401905b80831015612545578551805173ffffffffffffffffffffffffffffffffffffffff16835285015185830152948401946001929092019190830190612506565b506060870151945061256f606089018673ffffffffffffffffffffffffffffffffffffffff169052565b6080870151945087810360808901526125888186611e4b565b98975050505050505050565b6020815260006116d960208301846124b7565b6000602082840312156125b957600080fd5b5051919050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126125f557600080fd5b830160208101925035905067ffffffffffffffff81111561261557600080fd5b80360382131561145457600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8183526000602080850194508260005b858110156126c257813561269081611ea8565b73ffffffffffffffffffffffffffffffffffffffff16875281830135838801526040968701969091019060010161267d565b509495945050505050565b602081528135602082015260006126e660208401611ef7565b67ffffffffffffffff808216604085015261270460408601866125c0565b925060a0606086015261271b60c086018483612624565b92505061272b60608601866125c0565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe080878603016080880152612761858385612624565b9450608088013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe188360301831261279a57600080fd5b602092880192830192359150838211156127b357600080fd5b8160061b36038313156127c557600080fd5b8685030160a0870152611bd184828461266d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561281a57600080fd5b81516116d981611ea8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361288557612885612825565b5060010190565b60608152600061289f60608301866124b7565b905083602083015273ffffffffffffffffffffffffffffffffffffffff83166040830152949350505050565b600060208083850312156128de57600080fd5b825167ffffffffffffffff8111156128f557600080fd5b8301601f8101851361290657600080fd5b80516129146120ac82612067565b81815260059190911b8201830190838101908783111561293357600080fd5b928401925b82841015611bd157835161294b81611ea8565b82529284019290840190612938565b60006020828403121561296c57600080fd5b815180151581146116d957600080fd5b6000825161298e818460208701611e27565b9190910192915050565b818103818111156105be576105be612825565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000813000a",
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

func (_Router *RouterCaller) IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, uint64, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "isOffRamp", offRamp)

	if err != nil {
		return *new(bool), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

func (_Router *RouterSession) IsOffRamp(offRamp common.Address) (bool, uint64, error) {
	return _Router.Contract.IsOffRamp(&_Router.CallOpts, offRamp)
}

func (_Router *RouterCallerSession) IsOffRamp(offRamp common.Address) (bool, uint64, error) {
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
	return common.HexToHash("0xbd2f9ed173f6f0237225df18b4721b82fb4e79583a71c0bd731669b908de4333")
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
	MAXRETBYTES(opts *bind.CallOpts) (uint16, error)

	GetFee(opts *bind.CallOpts, destinationChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetOffRamps(opts *bind.CallOpts) ([]RouterOffRamp, error)

	GetOnRamp(opts *bind.CallOpts, destChainSelector uint64) (common.Address, error)

	GetSupportedTokens(opts *bind.CallOpts, chainSelector uint64) ([]common.Address, error)

	GetWrappedNative(opts *bind.CallOpts) (common.Address, error)

	IsChainSupported(opts *bind.CallOpts, chainSelector uint64) (bool, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, uint64, error)

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
