// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ge_router

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

type CommonEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

type GEConsumerEVM2AnyGEMessage struct {
	Receiver         []byte
	Data             []byte
	TokensAndAmounts []CommonEVMTokenAndAmount
	FeeToken         common.Address
	ExtraArgs        []byte
}

type InternalAny2EVMMessageFromSender struct {
	SourceChainId        uint64
	Sender               []byte
	Receiver             common.Address
	Data                 []byte
	DestPools            []common.Address
	DestTokensAndAmounts []CommonEVMTokenAndAmount
	GasLimit             *big.Int
}

var GERouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"AlreadyConfigured\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MustCallFromOffRamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoOffRampsConfigured\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"contractEVM2EVMGEOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"contractEVM2EVMGEOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"addOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structGEConsumer.EVM2AnyGEMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structGEConsumer.EVM2AnyGEMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRamps\",\"outputs\":[{\"internalType\":\"contractBaseOffRampInterface[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"contractEVM2EVMGEOnRampInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"removeOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"routeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"contractEVM2EVMGEOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"setOnRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002352380380620023528339810160408190526200003491620002f9565b8033806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf816200019a565b50508151620000d79150600390602084019062000245565b5060005b815181101562000191576040518060400160405280826001600160601b0316815260200160011515815250600260008484815181106200011f576200011f620003cb565b6020908102919091018101516001600160a01b031682528181019290925260400160002082518154939092015115156c01000000000000000000000000026001600160681b03199093166001600160601b03909216919091179190911790556200018981620003e1565b9050620000db565b50505062000409565b336001600160a01b03821603620001f45760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200029d579160200282015b828111156200029d57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000266565b50620002ab929150620002af565b5090565b5b80821115620002ab5760008155600101620002b0565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b0381168114620002f457600080fd5b919050565b600060208083850312156200030d57600080fd5b82516001600160401b03808211156200032557600080fd5b818501915085601f8301126200033a57600080fd5b8151818111156200034f576200034f620002c6565b8060051b604051601f19603f83011681018181108582111715620003775762000377620002c6565b6040529182528482019250838101850191888311156200039657600080fd5b938501935b82851015620003bf57620003af85620002dc565b845293850193928501926200039b565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000600182016200040257634e487b7160e01b600052601160045260246000fd5b5060010190565b611f3980620004196000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063991f65431161008c578063a8d87a3b11610066578063a8d87a3b14610260578063acd754d414610293578063adb9f71b146102a6578063f2fde38b146102b957600080fd5b8063991f654314610203578063a40e69c714610216578063a48a90581461022b57600080fd5b80632e36d584116100c85780632e36d584146101ae57806379ba5097146101c35780638da5cb5b146101cb57806396f4e9f9146101f057600080fd5b8063181f5a77146100ef5780631d7a74a01461014157806320487ded1461018d575b600080fd5b61012b6040518060400160405280600e81526020017f4745526f7574657220312e302e3000000000000000000000000000000000000081525081565b60405161013891906116a5565b60405180910390f35b61017d61014f3660046116cd565b6001600160a01b03166000908152600260205260409020546c01000000000000000000000000900460ff1690565b6040519015158152602001610138565b6101a061019b36600461190d565b6102cc565b604051908152602001610138565b6101c16101bc366004611a0a565b6103c7565b005b6101c16104c2565b6000546001600160a01b03165b6040516001600160a01b039091168152602001610138565b6101a06101fe36600461190d565b61058b565b6101c16102113660046116cd565b6107d9565b61021e610af4565b6040516101389190611a41565b61017d610239366004611a8e565b67ffffffffffffffff166000908152600460205260409020546001600160a01b0316151590565b6101d861026e366004611a8e565b67ffffffffffffffff166000908152600460205260409020546001600160a01b031690565b61017d6102a1366004611ab7565b610b56565b6101c16102b43660046116cd565b610ce9565b6101c16102c73660046116cd565b610ec5565b67ffffffffffffffff82166000908152600460205260408120546001600160a01b031680610337576040517fae236d9c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff851660048201526024015b60405180910390fd5b6040517f38724a950000000000000000000000000000000000000000000000000000000081526001600160a01b038216906338724a959061037c908690600401611bd0565b602060405180830381865afa158015610399573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103bd9190611be3565b9150505b92915050565b6103cf610ed9565b67ffffffffffffffff82166000908152600460205260409020546001600160a01b03808316911603610448576040517f74456f4900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff831660048201526001600160a01b038216602482015260440161032e565b67ffffffffffffffff821660008181526004602052604080822080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03861690811790915590519092917f1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f2391a35050565b6001546001600160a01b0316331461051c5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161032e565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b67ffffffffffffffff82166000908152600460205260408120546001600160a01b0316816105b985856102cc565b905060006105ec8560400151604051806040016040528088606001516001600160a01b0316815260200185815250610f35565b905060005b815181101561074257600082828151811061060e5761060e611bfc565b6020908102919091010151516040517f5d86f1410000000000000000000000000000000000000000000000000000000081526001600160a01b038083166004830152919250600091871690635d86f141906024016020604051808303816000875af1158015610681573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a59190611c2b565b90506001600160a01b0381166106f2576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161032e565b61072f338286868151811061070957610709611bfc565b602002602001015160200151856001600160a01b03166111c0909392919063ffffffff16565b50508061073b90611c77565b90506105f1565b506040517fa7d3e02f0000000000000000000000000000000000000000000000000000000081526001600160a01b0384169063a7d3e02f9061078c90889086903390600401611caf565b6020604051808303816000875af11580156107ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107cf9190611be3565b9695505050505050565b6107e1610ed9565b600354600081900361081f576040517f22babb3200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382166000908152600260209081526040918290208251808401909352546bffffffffffffffffffffffff811683526c01000000000000000000000000900460ff1615159082018190526108b1576040517f8c97f1220000000000000000000000000000000000000000000000000000000081526001600160a01b038416600482015260240161032e565b600060036108c0600185611ce1565b815481106108d0576108d0611bfc565b6000918252602090912001548251600380546001600160a01b039093169350916bffffffffffffffffffffffff90911690811061090f5761090f611bfc565b6000918252602090912001546001600160a01b03166003610931600186611ce1565b8154811061094157610941611bfc565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600383600001516bffffffffffffffffffffffff168154811061099557610995611bfc565b600091825260208083209190910180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0394851617905584519284168252600290526040902080547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff9092169190911790556003805480610a2f57610a2f611cf8565b6000828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690559092019092556001600160a01b0386168083526002909152604080832080547fffffffffffffffffffffffffffffffffffffff000000000000000000000000001690555190917fcf91daec21e3510e2f2aea4b09d08c235d5c6844980be709f282ef591dbf420c91a250505050565b60606003805480602002602001604051908101604052809291908181526020018280548015610b4c57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610b2e575b5050505050905090565b336000818152600260205260408120549091906c01000000000000000000000000900460ff16610bb4576040517fa2c8bfb600000000000000000000000000000000000000000000000000000000815233600482015260240161032e565b6000610bc7610bc286611d8b565b61124e565b90506000633015b91c60e01b82604051602401610be49190611e6f565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152905084610c7157610c6a60c0870135610c626060890160408a016116cd565b6000846112c5565b9350610ce0565b610c8160608701604088016116cd565b6001600160a01b031681604051610c989190611edb565b6000604051808303816000865af19150503d8060008114610cd5576040519150601f19603f3d011682016040523d82523d6000602084013e610cda565b606091505b50909450505b50505092915050565b610cf1610ed9565b6001600160a01b038116610d31576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381166000908152600260209081526040918290208251808401909352546bffffffffffffffffffffffff811683526c01000000000000000000000000900460ff1615801591830191909152610dc5576040517f3a4406b50000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161032e565b60016020808301828152600380546bffffffffffffffffffffffff90811686526001600160a01b03871660008181526002909552604080862088518154965115156c01000000000000000000000000027fffffffffffffffffffffffffffffffffffffff0000000000000000000000000090971694169390931794909417909155815494850182559083527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b90930180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055517f78f53b26906785548b265fa08f4197f9f3fff73fe0d504d30400aacb527f4ce09190a25050565b610ecd610ed9565b610ed681611311565b50565b6000546001600160a01b03163314610f335760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161032e565b565b606060005b83518110156110d45782600001516001600160a01b0316848281518110610f6357610f63611bfc565b6020026020010151600001516001600160a01b0316036110c4576000845167ffffffffffffffff811115610f9957610f99611707565b604051908082528060200260200182016040528015610fde57816020015b6040805180820190915260008082526020820152815260200190600190039081610fb75790505b50905060005b855181101561103557858181518110610fff57610fff611bfc565b602002602001015182828151811061101957611019611bfc565b60200260200101819052508061102e90611c77565b9050610fe4565b50604051806040016040528082848151811061105357611053611bfc565b6020026020010151600001516001600160a01b03168152602001856020015183858151811061108457611084611bfc565b60200260200101516020015161109a9190611ef7565b8152508183815181106110af576110af611bfc565b602002602001018190525080925050506103c1565b6110cd81611c77565b9050610f3a565b506000835160016110e59190611ef7565b67ffffffffffffffff8111156110fd576110fd611707565b60405190808252806020026020018201604052801561114257816020015b604080518082019091526000808252602082015281526020019060019003908161111b5790505b50905060005b84518110156111995784818151811061116357611163611bfc565b602002602001015182828151811061117d5761117d611bfc565b60200260200101819052508061119290611c77565b9050611148565b5082818551815181106111ae576111ae611bfc565b60209081029190910101529392505050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526112489085906113d2565b50505050565b6112836040518060800160405280600067ffffffffffffffff1681526020016060815260200160608152602001606081525090565b6040518060800160405280836000015167ffffffffffffffff16815260200183602001518152602001836060015181526020018360a001518152509050919050565b60005a6113888110156112d757600080fd5b6113888103905085604082048203116112ef57600080fd5b50833b6112fb57600080fd5b60008083516020850186888af195945050505050565b336001600160a01b038216036113695760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161032e565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000611427826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166114bc9092919063ffffffff16565b8051909150156114b757808060200190518101906114459190611f0f565b6114b75760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161032e565b505050565b60606114cb84846000856114d5565b90505b9392505050565b60608247101561154d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161032e565b843b61159b5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161032e565b600080866001600160a01b031685876040516115b79190611edb565b60006040518083038185875af1925050503d80600081146115f4576040519150601f19603f3d011682016040523d82523d6000602084013e6115f9565b606091505b5091509150611609828286611614565b979650505050505050565b606083156116235750816114ce565b8251156116335782518084602001fd5b8160405162461bcd60e51b815260040161032e91906116a5565b60005b83811015611668578181015183820152602001611650565b838111156112485750506000910152565b6000815180845261169181602086016020860161164d565b601f01601f19169290920160200192915050565b6020815260006114ce6020830184611679565b6001600160a01b0381168114610ed657600080fd5b6000602082840312156116df57600080fd5b81356114ce816116b8565b803567ffffffffffffffff8116811461170257600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561175957611759611707565b60405290565b60405160a0810167ffffffffffffffff8111828210171561175957611759611707565b60405160e0810167ffffffffffffffff8111828210171561175957611759611707565b604051601f8201601f1916810167ffffffffffffffff811182821017156117ce576117ce611707565b604052919050565b600082601f8301126117e757600080fd5b813567ffffffffffffffff81111561180157611801611707565b6118146020601f19601f840116016117a5565b81815284602083860101111561182957600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff82111561186057611860611707565b5060051b60200190565b8035611702816116b8565b600082601f83011261188657600080fd5b8135602061189b61189683611846565b6117a5565b82815260069290921b840181019181810190868411156118ba57600080fd5b8286015b8481101561190257604081890312156118d75760008081fd5b6118df611736565b81356118ea816116b8565b815281850135858201528352918301916040016118be565b509695505050505050565b6000806040838503121561192057600080fd5b611929836116ea565b9150602083013567ffffffffffffffff8082111561194657600080fd5b9084019060a0828703121561195a57600080fd5b61196261175f565b82358281111561197157600080fd5b61197d888286016117d6565b82525060208301358281111561199257600080fd5b61199e888286016117d6565b6020830152506040830135828111156119b657600080fd5b6119c288828601611875565b6040830152506119d46060840161186a565b60608201526080830135828111156119eb57600080fd5b6119f7888286016117d6565b6080830152508093505050509250929050565b60008060408385031215611a1d57600080fd5b611a26836116ea565b91506020830135611a36816116b8565b809150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015611a825783516001600160a01b031683529284019291840191600101611a5d565b50909695505050505050565b600060208284031215611aa057600080fd5b6114ce826116ea565b8015158114610ed657600080fd5b60008060408385031215611aca57600080fd5b823567ffffffffffffffff811115611ae157600080fd5b830160e08186031215611af357600080fd5b91506020830135611a3681611aa9565b600081518084526020808501945080840160005b83811015611b4757815180516001600160a01b031688528301518388015260409096019590820190600101611b17565b509495945050505050565b6000815160a08452611b6760a0850182611679565b905060208301518482036020860152611b808282611679565b91505060408301518482036040860152611b9a8282611b03565b9150506001600160a01b03606084015116606085015260808301518482036080860152611bc78282611679565b95945050505050565b6020815260006114ce6020830184611b52565b600060208284031215611bf557600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215611c3d57600080fd5b81516114ce816116b8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611ca857611ca8611c48565b5060010190565b606081526000611cc26060830186611b52565b90508360208301526001600160a01b0383166040830152949350505050565b600082821015611cf357611cf3611c48565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600082601f830112611d3857600080fd5b81356020611d4861189683611846565b82815260059290921b84018101918181019086841115611d6757600080fd5b8286015b84811015611902578035611d7e816116b8565b8352918301918301611d6b565b600060e08236031215611d9d57600080fd5b611da5611782565b611dae836116ea565b8152602083013567ffffffffffffffff80821115611dcb57600080fd5b611dd7368387016117d6565b6020840152611de86040860161186a565b60408401526060850135915080821115611e0157600080fd5b611e0d368387016117d6565b60608401526080850135915080821115611e2657600080fd5b611e3236838701611d27565b608084015260a0850135915080821115611e4b57600080fd5b50611e5836828601611875565b60a08301525060c092830135928101929092525090565b6020815267ffffffffffffffff82511660208201526000602083015160806040840152611e9f60a0840182611679565b90506040840151601f1980858403016060860152611ebd8383611679565b9250606086015191508085840301608086015250611bc78282611b03565b60008251611eed81846020870161164d565b9190910192915050565b60008219821115611f0a57611f0a611c48565b500190565b600060208284031215611f2157600080fd5b81516114ce81611aa956fea164736f6c634300080f000a",
}

var GERouterABI = GERouterMetaData.ABI

var GERouterBin = GERouterMetaData.Bin

func DeployGERouter(auth *bind.TransactOpts, backend bind.ContractBackend, offRamps []common.Address) (common.Address, *types.Transaction, *GERouter, error) {
	parsed, err := GERouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GERouterBin), backend, offRamps)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GERouter{GERouterCaller: GERouterCaller{contract: contract}, GERouterTransactor: GERouterTransactor{contract: contract}, GERouterFilterer: GERouterFilterer{contract: contract}}, nil
}

type GERouter struct {
	address common.Address
	abi     abi.ABI
	GERouterCaller
	GERouterTransactor
	GERouterFilterer
}

type GERouterCaller struct {
	contract *bind.BoundContract
}

type GERouterTransactor struct {
	contract *bind.BoundContract
}

type GERouterFilterer struct {
	contract *bind.BoundContract
}

type GERouterSession struct {
	Contract     *GERouter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type GERouterCallerSession struct {
	Contract *GERouterCaller
	CallOpts bind.CallOpts
}

type GERouterTransactorSession struct {
	Contract     *GERouterTransactor
	TransactOpts bind.TransactOpts
}

type GERouterRaw struct {
	Contract *GERouter
}

type GERouterCallerRaw struct {
	Contract *GERouterCaller
}

type GERouterTransactorRaw struct {
	Contract *GERouterTransactor
}

func NewGERouter(address common.Address, backend bind.ContractBackend) (*GERouter, error) {
	abi, err := abi.JSON(strings.NewReader(GERouterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindGERouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GERouter{address: address, abi: abi, GERouterCaller: GERouterCaller{contract: contract}, GERouterTransactor: GERouterTransactor{contract: contract}, GERouterFilterer: GERouterFilterer{contract: contract}}, nil
}

func NewGERouterCaller(address common.Address, caller bind.ContractCaller) (*GERouterCaller, error) {
	contract, err := bindGERouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GERouterCaller{contract: contract}, nil
}

func NewGERouterTransactor(address common.Address, transactor bind.ContractTransactor) (*GERouterTransactor, error) {
	contract, err := bindGERouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GERouterTransactor{contract: contract}, nil
}

func NewGERouterFilterer(address common.Address, filterer bind.ContractFilterer) (*GERouterFilterer, error) {
	contract, err := bindGERouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GERouterFilterer{contract: contract}, nil
}

func bindGERouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GERouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_GERouter *GERouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GERouter.Contract.GERouterCaller.contract.Call(opts, result, method, params...)
}

func (_GERouter *GERouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GERouter.Contract.GERouterTransactor.contract.Transfer(opts)
}

func (_GERouter *GERouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GERouter.Contract.GERouterTransactor.contract.Transact(opts, method, params...)
}

func (_GERouter *GERouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GERouter.Contract.contract.Call(opts, result, method, params...)
}

func (_GERouter *GERouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GERouter.Contract.contract.Transfer(opts)
}

func (_GERouter *GERouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GERouter.Contract.contract.Transact(opts, method, params...)
}

func (_GERouter *GERouterCaller) GetFee(opts *bind.CallOpts, destinationChainId uint64, message GEConsumerEVM2AnyGEMessage) (*big.Int, error) {
	var out []interface{}
	err := _GERouter.contract.Call(opts, &out, "getFee", destinationChainId, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_GERouter *GERouterSession) GetFee(destinationChainId uint64, message GEConsumerEVM2AnyGEMessage) (*big.Int, error) {
	return _GERouter.Contract.GetFee(&_GERouter.CallOpts, destinationChainId, message)
}

func (_GERouter *GERouterCallerSession) GetFee(destinationChainId uint64, message GEConsumerEVM2AnyGEMessage) (*big.Int, error) {
	return _GERouter.Contract.GetFee(&_GERouter.CallOpts, destinationChainId, message)
}

func (_GERouter *GERouterCaller) GetOffRamps(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GERouter.contract.Call(opts, &out, "getOffRamps")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_GERouter *GERouterSession) GetOffRamps() ([]common.Address, error) {
	return _GERouter.Contract.GetOffRamps(&_GERouter.CallOpts)
}

func (_GERouter *GERouterCallerSession) GetOffRamps() ([]common.Address, error) {
	return _GERouter.Contract.GetOffRamps(&_GERouter.CallOpts)
}

func (_GERouter *GERouterCaller) GetOnRamp(opts *bind.CallOpts, chainId uint64) (common.Address, error) {
	var out []interface{}
	err := _GERouter.contract.Call(opts, &out, "getOnRamp", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_GERouter *GERouterSession) GetOnRamp(chainId uint64) (common.Address, error) {
	return _GERouter.Contract.GetOnRamp(&_GERouter.CallOpts, chainId)
}

func (_GERouter *GERouterCallerSession) GetOnRamp(chainId uint64) (common.Address, error) {
	return _GERouter.Contract.GetOnRamp(&_GERouter.CallOpts, chainId)
}

func (_GERouter *GERouterCaller) IsChainSupported(opts *bind.CallOpts, chainId uint64) (bool, error) {
	var out []interface{}
	err := _GERouter.contract.Call(opts, &out, "isChainSupported", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_GERouter *GERouterSession) IsChainSupported(chainId uint64) (bool, error) {
	return _GERouter.Contract.IsChainSupported(&_GERouter.CallOpts, chainId)
}

func (_GERouter *GERouterCallerSession) IsChainSupported(chainId uint64) (bool, error) {
	return _GERouter.Contract.IsChainSupported(&_GERouter.CallOpts, chainId)
}

func (_GERouter *GERouterCaller) IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _GERouter.contract.Call(opts, &out, "isOffRamp", offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_GERouter *GERouterSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _GERouter.Contract.IsOffRamp(&_GERouter.CallOpts, offRamp)
}

func (_GERouter *GERouterCallerSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _GERouter.Contract.IsOffRamp(&_GERouter.CallOpts, offRamp)
}

func (_GERouter *GERouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GERouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_GERouter *GERouterSession) Owner() (common.Address, error) {
	return _GERouter.Contract.Owner(&_GERouter.CallOpts)
}

func (_GERouter *GERouterCallerSession) Owner() (common.Address, error) {
	return _GERouter.Contract.Owner(&_GERouter.CallOpts)
}

func (_GERouter *GERouterCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GERouter.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_GERouter *GERouterSession) TypeAndVersion() (string, error) {
	return _GERouter.Contract.TypeAndVersion(&_GERouter.CallOpts)
}

func (_GERouter *GERouterCallerSession) TypeAndVersion() (string, error) {
	return _GERouter.Contract.TypeAndVersion(&_GERouter.CallOpts)
}

func (_GERouter *GERouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GERouter.contract.Transact(opts, "acceptOwnership")
}

func (_GERouter *GERouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _GERouter.Contract.AcceptOwnership(&_GERouter.TransactOpts)
}

func (_GERouter *GERouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _GERouter.Contract.AcceptOwnership(&_GERouter.TransactOpts)
}

func (_GERouter *GERouterTransactor) AddOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error) {
	return _GERouter.contract.Transact(opts, "addOffRamp", offRamp)
}

func (_GERouter *GERouterSession) AddOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _GERouter.Contract.AddOffRamp(&_GERouter.TransactOpts, offRamp)
}

func (_GERouter *GERouterTransactorSession) AddOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _GERouter.Contract.AddOffRamp(&_GERouter.TransactOpts, offRamp)
}

func (_GERouter *GERouterTransactor) CcipSend(opts *bind.TransactOpts, destinationChainId uint64, message GEConsumerEVM2AnyGEMessage) (*types.Transaction, error) {
	return _GERouter.contract.Transact(opts, "ccipSend", destinationChainId, message)
}

func (_GERouter *GERouterSession) CcipSend(destinationChainId uint64, message GEConsumerEVM2AnyGEMessage) (*types.Transaction, error) {
	return _GERouter.Contract.CcipSend(&_GERouter.TransactOpts, destinationChainId, message)
}

func (_GERouter *GERouterTransactorSession) CcipSend(destinationChainId uint64, message GEConsumerEVM2AnyGEMessage) (*types.Transaction, error) {
	return _GERouter.Contract.CcipSend(&_GERouter.TransactOpts, destinationChainId, message)
}

func (_GERouter *GERouterTransactor) RemoveOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error) {
	return _GERouter.contract.Transact(opts, "removeOffRamp", offRamp)
}

func (_GERouter *GERouterSession) RemoveOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _GERouter.Contract.RemoveOffRamp(&_GERouter.TransactOpts, offRamp)
}

func (_GERouter *GERouterTransactorSession) RemoveOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _GERouter.Contract.RemoveOffRamp(&_GERouter.TransactOpts, offRamp)
}

func (_GERouter *GERouterTransactor) RouteMessage(opts *bind.TransactOpts, message InternalAny2EVMMessageFromSender, manualExecution bool) (*types.Transaction, error) {
	return _GERouter.contract.Transact(opts, "routeMessage", message, manualExecution)
}

func (_GERouter *GERouterSession) RouteMessage(message InternalAny2EVMMessageFromSender, manualExecution bool) (*types.Transaction, error) {
	return _GERouter.Contract.RouteMessage(&_GERouter.TransactOpts, message, manualExecution)
}

func (_GERouter *GERouterTransactorSession) RouteMessage(message InternalAny2EVMMessageFromSender, manualExecution bool) (*types.Transaction, error) {
	return _GERouter.Contract.RouteMessage(&_GERouter.TransactOpts, message, manualExecution)
}

func (_GERouter *GERouterTransactor) SetOnRamp(opts *bind.TransactOpts, chainId uint64, onRamp common.Address) (*types.Transaction, error) {
	return _GERouter.contract.Transact(opts, "setOnRamp", chainId, onRamp)
}

func (_GERouter *GERouterSession) SetOnRamp(chainId uint64, onRamp common.Address) (*types.Transaction, error) {
	return _GERouter.Contract.SetOnRamp(&_GERouter.TransactOpts, chainId, onRamp)
}

func (_GERouter *GERouterTransactorSession) SetOnRamp(chainId uint64, onRamp common.Address) (*types.Transaction, error) {
	return _GERouter.Contract.SetOnRamp(&_GERouter.TransactOpts, chainId, onRamp)
}

func (_GERouter *GERouterTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _GERouter.contract.Transact(opts, "transferOwnership", to)
}

func (_GERouter *GERouterSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _GERouter.Contract.TransferOwnership(&_GERouter.TransactOpts, to)
}

func (_GERouter *GERouterTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _GERouter.Contract.TransferOwnership(&_GERouter.TransactOpts, to)
}

type GERouterOffRampAddedIterator struct {
	Event *GERouterOffRampAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GERouterOffRampAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GERouterOffRampAdded)
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
		it.Event = new(GERouterOffRampAdded)
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

func (it *GERouterOffRampAddedIterator) Error() error {
	return it.fail
}

func (it *GERouterOffRampAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GERouterOffRampAdded struct {
	OffRamp common.Address
	Raw     types.Log
}

func (_GERouter *GERouterFilterer) FilterOffRampAdded(opts *bind.FilterOpts, offRamp []common.Address) (*GERouterOffRampAddedIterator, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _GERouter.contract.FilterLogs(opts, "OffRampAdded", offRampRule)
	if err != nil {
		return nil, err
	}
	return &GERouterOffRampAddedIterator{contract: _GERouter.contract, event: "OffRampAdded", logs: logs, sub: sub}, nil
}

func (_GERouter *GERouterFilterer) WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *GERouterOffRampAdded, offRamp []common.Address) (event.Subscription, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _GERouter.contract.WatchLogs(opts, "OffRampAdded", offRampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GERouterOffRampAdded)
				if err := _GERouter.contract.UnpackLog(event, "OffRampAdded", log); err != nil {
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

func (_GERouter *GERouterFilterer) ParseOffRampAdded(log types.Log) (*GERouterOffRampAdded, error) {
	event := new(GERouterOffRampAdded)
	if err := _GERouter.contract.UnpackLog(event, "OffRampAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GERouterOffRampRemovedIterator struct {
	Event *GERouterOffRampRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GERouterOffRampRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GERouterOffRampRemoved)
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
		it.Event = new(GERouterOffRampRemoved)
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

func (it *GERouterOffRampRemovedIterator) Error() error {
	return it.fail
}

func (it *GERouterOffRampRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GERouterOffRampRemoved struct {
	OffRamp common.Address
	Raw     types.Log
}

func (_GERouter *GERouterFilterer) FilterOffRampRemoved(opts *bind.FilterOpts, offRamp []common.Address) (*GERouterOffRampRemovedIterator, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _GERouter.contract.FilterLogs(opts, "OffRampRemoved", offRampRule)
	if err != nil {
		return nil, err
	}
	return &GERouterOffRampRemovedIterator{contract: _GERouter.contract, event: "OffRampRemoved", logs: logs, sub: sub}, nil
}

func (_GERouter *GERouterFilterer) WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *GERouterOffRampRemoved, offRamp []common.Address) (event.Subscription, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _GERouter.contract.WatchLogs(opts, "OffRampRemoved", offRampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GERouterOffRampRemoved)
				if err := _GERouter.contract.UnpackLog(event, "OffRampRemoved", log); err != nil {
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

func (_GERouter *GERouterFilterer) ParseOffRampRemoved(log types.Log) (*GERouterOffRampRemoved, error) {
	event := new(GERouterOffRampRemoved)
	if err := _GERouter.contract.UnpackLog(event, "OffRampRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GERouterOnRampSetIterator struct {
	Event *GERouterOnRampSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GERouterOnRampSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GERouterOnRampSet)
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
		it.Event = new(GERouterOnRampSet)
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

func (it *GERouterOnRampSetIterator) Error() error {
	return it.fail
}

func (it *GERouterOnRampSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GERouterOnRampSet struct {
	ChainId uint64
	OnRamp  common.Address
	Raw     types.Log
}

func (_GERouter *GERouterFilterer) FilterOnRampSet(opts *bind.FilterOpts, chainId []uint64, onRamp []common.Address) (*GERouterOnRampSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var onRampRule []interface{}
	for _, onRampItem := range onRamp {
		onRampRule = append(onRampRule, onRampItem)
	}

	logs, sub, err := _GERouter.contract.FilterLogs(opts, "OnRampSet", chainIdRule, onRampRule)
	if err != nil {
		return nil, err
	}
	return &GERouterOnRampSetIterator{contract: _GERouter.contract, event: "OnRampSet", logs: logs, sub: sub}, nil
}

func (_GERouter *GERouterFilterer) WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *GERouterOnRampSet, chainId []uint64, onRamp []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var onRampRule []interface{}
	for _, onRampItem := range onRamp {
		onRampRule = append(onRampRule, onRampItem)
	}

	logs, sub, err := _GERouter.contract.WatchLogs(opts, "OnRampSet", chainIdRule, onRampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GERouterOnRampSet)
				if err := _GERouter.contract.UnpackLog(event, "OnRampSet", log); err != nil {
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

func (_GERouter *GERouterFilterer) ParseOnRampSet(log types.Log) (*GERouterOnRampSet, error) {
	event := new(GERouterOnRampSet)
	if err := _GERouter.contract.UnpackLog(event, "OnRampSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GERouterOwnershipTransferRequestedIterator struct {
	Event *GERouterOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GERouterOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GERouterOwnershipTransferRequested)
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
		it.Event = new(GERouterOwnershipTransferRequested)
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

func (it *GERouterOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *GERouterOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GERouterOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_GERouter *GERouterFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GERouterOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GERouter.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GERouterOwnershipTransferRequestedIterator{contract: _GERouter.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_GERouter *GERouterFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *GERouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GERouter.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GERouterOwnershipTransferRequested)
				if err := _GERouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_GERouter *GERouterFilterer) ParseOwnershipTransferRequested(log types.Log) (*GERouterOwnershipTransferRequested, error) {
	event := new(GERouterOwnershipTransferRequested)
	if err := _GERouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GERouterOwnershipTransferredIterator struct {
	Event *GERouterOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *GERouterOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GERouterOwnershipTransferred)
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
		it.Event = new(GERouterOwnershipTransferred)
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

func (it *GERouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *GERouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type GERouterOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_GERouter *GERouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GERouterOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GERouter.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GERouterOwnershipTransferredIterator{contract: _GERouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_GERouter *GERouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GERouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GERouter.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(GERouterOwnershipTransferred)
				if err := _GERouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_GERouter *GERouterFilterer) ParseOwnershipTransferred(log types.Log) (*GERouterOwnershipTransferred, error) {
	event := new(GERouterOwnershipTransferred)
	if err := _GERouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_GERouter *GERouter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _GERouter.abi.Events["OffRampAdded"].ID:
		return _GERouter.ParseOffRampAdded(log)
	case _GERouter.abi.Events["OffRampRemoved"].ID:
		return _GERouter.ParseOffRampRemoved(log)
	case _GERouter.abi.Events["OnRampSet"].ID:
		return _GERouter.ParseOnRampSet(log)
	case _GERouter.abi.Events["OwnershipTransferRequested"].ID:
		return _GERouter.ParseOwnershipTransferRequested(log)
	case _GERouter.abi.Events["OwnershipTransferred"].ID:
		return _GERouter.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (GERouterOffRampAdded) Topic() common.Hash {
	return common.HexToHash("0x78f53b26906785548b265fa08f4197f9f3fff73fe0d504d30400aacb527f4ce0")
}

func (GERouterOffRampRemoved) Topic() common.Hash {
	return common.HexToHash("0xcf91daec21e3510e2f2aea4b09d08c235d5c6844980be709f282ef591dbf420c")
}

func (GERouterOnRampSet) Topic() common.Hash {
	return common.HexToHash("0x1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f23")
}

func (GERouterOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (GERouterOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_GERouter *GERouter) Address() common.Address {
	return _GERouter.address
}

type GERouterInterface interface {
	GetFee(opts *bind.CallOpts, destinationChainId uint64, message GEConsumerEVM2AnyGEMessage) (*big.Int, error)

	GetOffRamps(opts *bind.CallOpts) ([]common.Address, error)

	GetOnRamp(opts *bind.CallOpts, chainId uint64) (common.Address, error)

	IsChainSupported(opts *bind.CallOpts, chainId uint64) (bool, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainId uint64, message GEConsumerEVM2AnyGEMessage) (*types.Transaction, error)

	RemoveOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error)

	RouteMessage(opts *bind.TransactOpts, message InternalAny2EVMMessageFromSender, manualExecution bool) (*types.Transaction, error)

	SetOnRamp(opts *bind.TransactOpts, chainId uint64, onRamp common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOffRampAdded(opts *bind.FilterOpts, offRamp []common.Address) (*GERouterOffRampAddedIterator, error)

	WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *GERouterOffRampAdded, offRamp []common.Address) (event.Subscription, error)

	ParseOffRampAdded(log types.Log) (*GERouterOffRampAdded, error)

	FilterOffRampRemoved(opts *bind.FilterOpts, offRamp []common.Address) (*GERouterOffRampRemovedIterator, error)

	WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *GERouterOffRampRemoved, offRamp []common.Address) (event.Subscription, error)

	ParseOffRampRemoved(log types.Log) (*GERouterOffRampRemoved, error)

	FilterOnRampSet(opts *bind.FilterOpts, chainId []uint64, onRamp []common.Address) (*GERouterOnRampSetIterator, error)

	WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *GERouterOnRampSet, chainId []uint64, onRamp []common.Address) (event.Subscription, error)

	ParseOnRampSet(log types.Log) (*GERouterOnRampSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GERouterOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *GERouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*GERouterOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GERouterOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GERouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*GERouterOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
