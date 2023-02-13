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

type ConsumerEVM2AnyMessage struct {
	Receiver         []byte
	Data             []byte
	TokensAndAmounts []CommonEVMTokenAndAmount
	FeeToken         common.Address
	ExtraArgs        []byte
}

var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"offRamps\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"wrappedNative\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"AlreadyConfigured\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFeeTokenAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MustCallFromOffRamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoOffRampsConfigured\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"contractIEVM2AnyOnRamp\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"contractIEVM2AnyOnRamp\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"addOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structConsumer.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structConsumer.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRamps\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"contractIEVM2AnyOnRamp\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedNative\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"removeOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCommon.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"routeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"contractIEVM2AnyOnRamp\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"setOnRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedNative\",\"type\":\"address\"}],\"name\":\"setWrappedNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620024ea380380620024ea833981016040819052620000349162000317565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620001b8565b50508251620000d69150600490602085019062000263565b5060005b825181101562000190576040518060400160405280826001600160601b0316815260200160011515815250600360008584815181106200011e576200011e620003fd565b6020908102919091018101516001600160a01b031682528181019290925260400160002082518154939092015115156c01000000000000000000000000026001600160681b03199093166001600160601b0390921691909117919091179055620001888162000413565b9050620000da565b50600280546001600160a01b0319166001600160a01b0392909216919091179055506200043b565b336001600160a01b03821603620002125760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620002bb579160200282015b82811115620002bb57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000284565b50620002c9929150620002cd565b5090565b5b80821115620002c95760008155600101620002ce565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b03811681146200031257600080fd5b919050565b600080604083850312156200032b57600080fd5b82516001600160401b03808211156200034357600080fd5b818501915085601f8301126200035857600080fd5b81516020828211156200036f576200036f620002e4565b8160051b604051601f19603f83011681018181108682111715620003975762000397620002e4565b604052928352818301935084810182019289841115620003b657600080fd5b948201945b83861015620003df57620003cf86620002fa565b85529482019493820193620003bb565b9650620003f09050878201620002fa565b9450505050509250929050565b634e487b7160e01b600052603260045260246000fd5b6000600182016200043457634e487b7160e01b600052601160045260246000fd5b5060010190565b61209f806200044b6000396000f3fe60806040526004361061010d5760003560e01c806396f4e9f9116100a5578063a8d87a3b11610074578063e861e90711610059578063e861e90714610394578063f2fde38b146103b2578063fbca3b74146103d257600080fd5b8063a8d87a3b14610334578063adb9f71b1461037457600080fd5b806396f4e9f91461029d578063991f6543146102b0578063a40e69c7146102d0578063a48a9058146102f257600080fd5b80632e36d584116100e15780632e36d5841461021457806352cb60ca1461023657806379ba5097146102565780638da5cb5b1461026b57600080fd5b80624b61bb14610112578063181f5a77146101475780631d7a74a01461019d57806320487ded146101e6575b600080fd5b34801561011e57600080fd5b5061013261012d366004611761565b6103f2565b60405190151581526020015b60405180910390f35b34801561015357600080fd5b506101906040518060400160405280600c81526020017f526f7574657220312e302e30000000000000000000000000000000000000000081525081565b60405161013e919061182a565b3480156101a957600080fd5b506101326101b836600461183d565b6001600160a01b03166000908152600360205260409020546c01000000000000000000000000900460ff1690565b3480156101f257600080fd5b50610206610201366004611a4a565b610552565b60405190815260200161013e565b34801561022057600080fd5b5061023461022f366004611b47565b61066d565b005b34801561024257600080fd5b5061023461025136600461183d565b61075d565b34801561026257600080fd5b50610234610794565b34801561027757600080fd5b506000546001600160a01b03165b6040516001600160a01b03909116815260200161013e565b6102066102ab366004611a4a565b610852565b3480156102bc57600080fd5b506102346102cb36600461183d565b610c94565b3480156102dc57600080fd5b506102e5610f99565b60405161013e9190611b7e565b3480156102fe57600080fd5b5061013261030d366004611bcb565b67ffffffffffffffff166000908152600560205260409020546001600160a01b0316151590565b34801561034057600080fd5b5061028561034f366004611bcb565b67ffffffffffffffff166000908152600560205260409020546001600160a01b031690565b34801561038057600080fd5b5061023461038f36600461183d565b610ffb565b3480156103a057600080fd5b506002546001600160a01b0316610285565b3480156103be57600080fd5b506102346103cd36600461183d565b6111cc565b3480156103de57600080fd5b506102e56103ed366004611bcb565b6111e0565b336000908152600360205260408120546c01000000000000000000000000900460ff16610452576040517fa2c8bfb60000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6000633015b91c60e01b8660405160240161046d9190611ccf565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091529050841561053957826001600160a01b0316816040516104ec9190611db0565b6000604051808303816000865af19150503d8060008114610529576040519150601f19603f3d011682016040523d82523d6000602084013e61052e565b606091505b505080925050610549565b61054684846000846112c7565b91505b50949350505050565b60608101516000906001600160a01b0316610579576002546001600160a01b031660608301525b67ffffffffffffffff83166000908152600560205260409020546001600160a01b0316806105df576040517fae236d9c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610449565b6040517f38724a950000000000000000000000000000000000000000000000000000000081526001600160a01b038216906338724a9590610624908690600401611e8f565b602060405180830381865afa158015610641573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106659190611ea2565b949350505050565b610675611313565b67ffffffffffffffff82166000908152600560205260409020546001600160a01b038083169116036106ee576040517f74456f4900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff831660048201526001600160a01b0382166024820152604401610449565b67ffffffffffffffff8216600081815260056020526040808220805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03861690811790915590519092917f1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f2391a35050565b610765611313565b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6001546001600160a01b031633146107ee5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610449565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b67ffffffffffffffff82166000908152600560205260408120546001600160a01b0316806108b8576040517fae236d9c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610449565b60608301516000906001600160a01b0316610a15576002546001600160a01b0390811660608601526040517f38724a95000000000000000000000000000000000000000000000000000000008152908316906338724a959061091e908790600401611e8f565b602060405180830381865afa15801561093b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095f9190611ea2565b90508034101561099b576040517f07da6ee600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b34905083606001516001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b1580156109dd57600080fd5b505af11580156109f1573d6000803e3d6000fd5b505050506060850151610a1091506001600160a01b031630848461136f565b610af0565b3415610a4d576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f38724a950000000000000000000000000000000000000000000000000000000081526001600160a01b038316906338724a9590610a92908790600401611e8f565b602060405180830381865afa158015610aaf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad39190611ea2565b6060850151909150610af0906001600160a01b031633848461136f565b60005b846040015151811015610bfe57600085604001518281518110610b1857610b18611ebb565b6020908102919091010151516040517f5d86f1410000000000000000000000000000000000000000000000000000000081526001600160a01b038083166004830152919250610bed91339190871690635d86f14190602401602060405180830381865afa158015610b8d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bb19190611eea565b88604001518581518110610bc757610bc7611ebb565b602002602001015160200151846001600160a01b031661136f909392919063ffffffff16565b50610bf781611f36565b9050610af3565b506040517fa7d3e02f0000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063a7d3e02f90610c4890879085903390600401611f6e565b6020604051808303816000875af1158015610c67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c8b9190611ea2565b95945050505050565b610c9c611313565b6004546000819003610cda576040517f22babb3200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382166000908152600360209081526040918290208251808401909352546bffffffffffffffffffffffff811683526c01000000000000000000000000900460ff161515908201819052610d6c576040517f8c97f1220000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401610449565b60006004610d7b600185611fa0565b81548110610d8b57610d8b611ebb565b6000918252602090912001548251600480546001600160a01b039093169350916bffffffffffffffffffffffff909116908110610dca57610dca611ebb565b6000918252602090912001546001600160a01b03166004610dec600186611fa0565b81548110610dfc57610dfc611ebb565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600483600001516bffffffffffffffffffffffff1681548110610e5057610e50611ebb565b6000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0394851617905584519284168252600390526040902080547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff9092169190911790556004805480610edf57610edf611fb7565b6000828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b0386168083526003909152604080832080547fffffffffffffffffffffffffffffffffffffff000000000000000000000000001690555190917fcf91daec21e3510e2f2aea4b09d08c235d5c6844980be709f282ef591dbf420c91a250505050565b60606004805480602002602001604051908101604052809291908181526020018280548015610ff157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610fd3575b5050505050905090565b611003611313565b6001600160a01b038116611043576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381166000908152600360209081526040918290208251808401909352546bffffffffffffffffffffffff811683526c01000000000000000000000000900460ff16158015918301919091526110d7576040517f3a4406b50000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610449565b60016020808301828152600480546bffffffffffffffffffffffff90811686526001600160a01b03871660008181526003909552604080862088518154965115156c01000000000000000000000000027fffffffffffffffffffffffffffffffffffffff0000000000000000000000000090971694169390931794909417909155815494850182559083527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909301805473ffffffffffffffffffffffffffffffffffffffff191684179055517f78f53b26906785548b265fa08f4197f9f3fff73fe0d504d30400aacb527f4ce09190a25050565b6111d4611313565b6111dd816113fd565b50565b606061120d8267ffffffffffffffff166000908152600560205260409020546001600160a01b0316151590565b61122557505060408051600081526020810190915290565b67ffffffffffffffff82166000908152600560205260408082205481517fd3c7c2c700000000000000000000000000000000000000000000000000000000815291516001600160a01b039091169263d3c7c2c792600480820193918290030181865afa158015611299573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526112c19190810190611fe6565b92915050565b60005a6113888110156112d957600080fd5b6113888103905085604082048203116112f157600080fd5b50833b6112fd57600080fd5b60008083516020850186888af195945050505050565b6000546001600160a01b0316331461136d5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610449565b565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526113f79085906114b3565b50505050565b336001600160a01b038216036114555760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610449565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000611508826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661159d9092919063ffffffff16565b80519091501561159857808060200190518101906115269190612075565b6115985760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610449565b505050565b60606115ac84846000856115b6565b90505b9392505050565b60608247101561162e5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610449565b843b61167c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610449565b600080866001600160a01b031685876040516116989190611db0565b60006040518083038185875af1925050503d80600081146116d5576040519150601f19603f3d011682016040523d82523d6000602084013e6116da565b606091505b50915091506116ea8282866116f5565b979650505050505050565b606083156117045750816115af565b8251156117145782518084602001fd5b8160405162461bcd60e51b8152600401610449919061182a565b80151581146111dd57600080fd5b6001600160a01b03811681146111dd57600080fd5b803561175c8161173c565b919050565b6000806000806080858703121561177757600080fd5b843567ffffffffffffffff81111561178e57600080fd5b8501608081880312156117a057600080fd5b935060208501356117b08161172e565b92506040850135915060608501356117c78161173c565b939692955090935050565b60005b838110156117ed5781810151838201526020016117d5565b838111156113f75750506000910152565b600081518084526118168160208601602086016117d2565b601f01601f19169290920160200192915050565b6020815260006115af60208301846117fe565b60006020828403121561184f57600080fd5b81356115af8161173c565b803567ffffffffffffffff8116811461175c57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156118c4576118c4611872565b60405290565b60405160a0810167ffffffffffffffff811182821017156118c4576118c4611872565b604051601f8201601f1916810167ffffffffffffffff8111828210171561191657611916611872565b604052919050565b600082601f83011261192f57600080fd5b813567ffffffffffffffff81111561194957611949611872565b61195c6020601f19601f840116016118ed565b81815284602083860101111561197157600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff8211156119a8576119a8611872565b5060051b60200190565b600082601f8301126119c357600080fd5b813560206119d86119d38361198e565b6118ed565b82815260069290921b840181019181810190868411156119f757600080fd5b8286015b84811015611a3f5760408189031215611a145760008081fd5b611a1c6118a1565b8135611a278161173c565b815281850135858201528352918301916040016119fb565b509695505050505050565b60008060408385031215611a5d57600080fd5b611a668361185a565b9150602083013567ffffffffffffffff80821115611a8357600080fd5b9084019060a08287031215611a9757600080fd5b611a9f6118ca565b823582811115611aae57600080fd5b611aba8882860161191e565b825250602083013582811115611acf57600080fd5b611adb8882860161191e565b602083015250604083013582811115611af357600080fd5b611aff888286016119b2565b604083015250611b1160608401611751565b6060820152608083013582811115611b2857600080fd5b611b348882860161191e565b6080830152508093505050509250929050565b60008060408385031215611b5a57600080fd5b611b638361185a565b91506020830135611b738161173c565b809150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015611bbf5783516001600160a01b031683529284019291840191600101611b9a565b50909695505050505050565b600060208284031215611bdd57600080fd5b6115af8261185a565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611c1b57600080fd5b830160208101925035905067ffffffffffffffff811115611c3b57600080fd5b803603821315611c4a57600080fd5b9250929050565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b8183526000602080850194508260005b85811015611cc4578135611c9f8161173c565b6001600160a01b03168752818301358388015260409687019690910190600101611c8c565b509495945050505050565b60208152600067ffffffffffffffff80611ce88561185a565b166020840152611cfb6020850185611be6565b60806040860152611d1060a086018284611c51565b915050611d206040860186611be6565b601f1980878503016060880152611d38848385611c51565b9350606088013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018312611d7157600080fd5b60209288019283019235915084821115611d8a57600080fd5b8160061b3603831315611d9c57600080fd5b8684030160808701526116ea838284611c7c565b60008251611dc28184602087016117d2565b9190910192915050565b6000815160a08452611de160a08501826117fe565b905060208084015185830382870152611dfa83826117fe565b60408681015188830389830152805180845290850195509092506000918401905b80831015611e4d57855180516001600160a01b0316835285015185830152948401946001929092019190830190611e1b565b5060608701519450611e6a60608901866001600160a01b03169052565b608087015194508781036080890152611e8381866117fe565b98975050505050505050565b6020815260006115af6020830184611dcc565b600060208284031215611eb457600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215611efc57600080fd5b81516115af8161173c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611f6757611f67611f07565b5060010190565b606081526000611f816060830186611dcc565b90508360208301526001600160a01b0383166040830152949350505050565b600082821015611fb257611fb2611f07565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60006020808385031215611ff957600080fd5b825167ffffffffffffffff81111561201057600080fd5b8301601f8101851361202157600080fd5b805161202f6119d38261198e565b81815260059190911b8201830190838101908783111561204e57600080fd5b928401925b828410156116ea5783516120668161173c565b82529284019290840190612053565b60006020828403121561208757600080fd5b81516115af8161172e56fea164736f6c634300080f000a",
}

var RouterABI = RouterMetaData.ABI

var RouterBin = RouterMetaData.Bin

func DeployRouter(auth *bind.TransactOpts, backend bind.ContractBackend, offRamps []common.Address, wrappedNative common.Address) (common.Address, *types.Transaction, *Router, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RouterBin), backend, offRamps, wrappedNative)
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

func (_Router *RouterCaller) GetFee(opts *bind.CallOpts, destinationChainId uint64, message ConsumerEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getFee", destinationChainId, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Router *RouterSession) GetFee(destinationChainId uint64, message ConsumerEVM2AnyMessage) (*big.Int, error) {
	return _Router.Contract.GetFee(&_Router.CallOpts, destinationChainId, message)
}

func (_Router *RouterCallerSession) GetFee(destinationChainId uint64, message ConsumerEVM2AnyMessage) (*big.Int, error) {
	return _Router.Contract.GetFee(&_Router.CallOpts, destinationChainId, message)
}

func (_Router *RouterCaller) GetOffRamps(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getOffRamps")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Router *RouterSession) GetOffRamps() ([]common.Address, error) {
	return _Router.Contract.GetOffRamps(&_Router.CallOpts)
}

func (_Router *RouterCallerSession) GetOffRamps() ([]common.Address, error) {
	return _Router.Contract.GetOffRamps(&_Router.CallOpts)
}

func (_Router *RouterCaller) GetOnRamp(opts *bind.CallOpts, chainId uint64) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getOnRamp", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Router *RouterSession) GetOnRamp(chainId uint64) (common.Address, error) {
	return _Router.Contract.GetOnRamp(&_Router.CallOpts, chainId)
}

func (_Router *RouterCallerSession) GetOnRamp(chainId uint64) (common.Address, error) {
	return _Router.Contract.GetOnRamp(&_Router.CallOpts, chainId)
}

func (_Router *RouterCaller) GetSupportedTokens(opts *bind.CallOpts, destChainId uint64) ([]common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getSupportedTokens", destChainId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Router *RouterSession) GetSupportedTokens(destChainId uint64) ([]common.Address, error) {
	return _Router.Contract.GetSupportedTokens(&_Router.CallOpts, destChainId)
}

func (_Router *RouterCallerSession) GetSupportedTokens(destChainId uint64) ([]common.Address, error) {
	return _Router.Contract.GetSupportedTokens(&_Router.CallOpts, destChainId)
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

func (_Router *RouterTransactor) AddOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addOffRamp", offRamp)
}

func (_Router *RouterSession) AddOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _Router.Contract.AddOffRamp(&_Router.TransactOpts, offRamp)
}

func (_Router *RouterTransactorSession) AddOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _Router.Contract.AddOffRamp(&_Router.TransactOpts, offRamp)
}

func (_Router *RouterTransactor) CcipSend(opts *bind.TransactOpts, destinationChainId uint64, message ConsumerEVM2AnyMessage) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "ccipSend", destinationChainId, message)
}

func (_Router *RouterSession) CcipSend(destinationChainId uint64, message ConsumerEVM2AnyMessage) (*types.Transaction, error) {
	return _Router.Contract.CcipSend(&_Router.TransactOpts, destinationChainId, message)
}

func (_Router *RouterTransactorSession) CcipSend(destinationChainId uint64, message ConsumerEVM2AnyMessage) (*types.Transaction, error) {
	return _Router.Contract.CcipSend(&_Router.TransactOpts, destinationChainId, message)
}

func (_Router *RouterTransactor) RemoveOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "removeOffRamp", offRamp)
}

func (_Router *RouterSession) RemoveOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _Router.Contract.RemoveOffRamp(&_Router.TransactOpts, offRamp)
}

func (_Router *RouterTransactorSession) RemoveOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _Router.Contract.RemoveOffRamp(&_Router.TransactOpts, offRamp)
}

func (_Router *RouterTransactor) RouteMessage(opts *bind.TransactOpts, message CommonAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "routeMessage", message, manualExecution, gasLimit, receiver)
}

func (_Router *RouterSession) RouteMessage(message CommonAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.RouteMessage(&_Router.TransactOpts, message, manualExecution, gasLimit, receiver)
}

func (_Router *RouterTransactorSession) RouteMessage(message CommonAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.RouteMessage(&_Router.TransactOpts, message, manualExecution, gasLimit, receiver)
}

func (_Router *RouterTransactor) SetOnRamp(opts *bind.TransactOpts, chainId uint64, onRamp common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setOnRamp", chainId, onRamp)
}

func (_Router *RouterSession) SetOnRamp(chainId uint64, onRamp common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetOnRamp(&_Router.TransactOpts, chainId, onRamp)
}

func (_Router *RouterTransactorSession) SetOnRamp(chainId uint64, onRamp common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetOnRamp(&_Router.TransactOpts, chainId, onRamp)
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
	OffRamp common.Address
	Raw     types.Log
}

func (_Router *RouterFilterer) FilterOffRampAdded(opts *bind.FilterOpts, offRamp []common.Address) (*RouterOffRampAddedIterator, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OffRampAdded", offRampRule)
	if err != nil {
		return nil, err
	}
	return &RouterOffRampAddedIterator{contract: _Router.contract, event: "OffRampAdded", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *RouterOffRampAdded, offRamp []common.Address) (event.Subscription, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OffRampAdded", offRampRule)
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
	OffRamp common.Address
	Raw     types.Log
}

func (_Router *RouterFilterer) FilterOffRampRemoved(opts *bind.FilterOpts, offRamp []common.Address) (*RouterOffRampRemovedIterator, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OffRampRemoved", offRampRule)
	if err != nil {
		return nil, err
	}
	return &RouterOffRampRemovedIterator{contract: _Router.contract, event: "OffRampRemoved", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *RouterOffRampRemoved, offRamp []common.Address) (event.Subscription, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OffRampRemoved", offRampRule)
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
	ChainId uint64
	OnRamp  common.Address
	Raw     types.Log
}

func (_Router *RouterFilterer) FilterOnRampSet(opts *bind.FilterOpts, chainId []uint64, onRamp []common.Address) (*RouterOnRampSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var onRampRule []interface{}
	for _, onRampItem := range onRamp {
		onRampRule = append(onRampRule, onRampItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OnRampSet", chainIdRule, onRampRule)
	if err != nil {
		return nil, err
	}
	return &RouterOnRampSetIterator{contract: _Router.contract, event: "OnRampSet", logs: logs, sub: sub}, nil
}

func (_Router *RouterFilterer) WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *RouterOnRampSet, chainId []uint64, onRamp []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var onRampRule []interface{}
	for _, onRampItem := range onRamp {
		onRampRule = append(onRampRule, onRampItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OnRampSet", chainIdRule, onRampRule)
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
	return common.HexToHash("0x78f53b26906785548b265fa08f4197f9f3fff73fe0d504d30400aacb527f4ce0")
}

func (RouterOffRampRemoved) Topic() common.Hash {
	return common.HexToHash("0xcf91daec21e3510e2f2aea4b09d08c235d5c6844980be709f282ef591dbf420c")
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
	GetFee(opts *bind.CallOpts, destinationChainId uint64, message ConsumerEVM2AnyMessage) (*big.Int, error)

	GetOffRamps(opts *bind.CallOpts) ([]common.Address, error)

	GetOnRamp(opts *bind.CallOpts, chainId uint64) (common.Address, error)

	GetSupportedTokens(opts *bind.CallOpts, destChainId uint64) ([]common.Address, error)

	GetWrappedNative(opts *bind.CallOpts) (common.Address, error)

	IsChainSupported(opts *bind.CallOpts, chainId uint64) (bool, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainId uint64, message ConsumerEVM2AnyMessage) (*types.Transaction, error)

	RemoveOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error)

	RouteMessage(opts *bind.TransactOpts, message CommonAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error)

	SetOnRamp(opts *bind.TransactOpts, chainId uint64, onRamp common.Address) (*types.Transaction, error)

	SetWrappedNative(opts *bind.TransactOpts, wrappedNative common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOffRampAdded(opts *bind.FilterOpts, offRamp []common.Address) (*RouterOffRampAddedIterator, error)

	WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *RouterOffRampAdded, offRamp []common.Address) (event.Subscription, error)

	ParseOffRampAdded(log types.Log) (*RouterOffRampAdded, error)

	FilterOffRampRemoved(opts *bind.FilterOpts, offRamp []common.Address) (*RouterOffRampRemovedIterator, error)

	WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *RouterOffRampRemoved, offRamp []common.Address) (event.Subscription, error)

	ParseOffRampRemoved(log types.Log) (*RouterOffRampRemoved, error)

	FilterOnRampSet(opts *bind.FilterOpts, chainId []uint64, onRamp []common.Address) (*RouterOnRampSetIterator, error)

	WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *RouterOnRampSet, chainId []uint64, onRamp []common.Address) (event.Subscription, error)

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
