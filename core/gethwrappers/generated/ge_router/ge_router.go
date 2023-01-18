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

type GEConsumerEVM2AnyGEMessage struct {
	Receiver         []byte
	Data             []byte
	TokensAndAmounts []CommonEVMTokenAndAmount
	FeeToken         common.Address
	ExtraArgs        []byte
}

var GERouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIBaseOffRamp[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"contractIBaseOffRamp\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"AlreadyConfigured\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MustCallFromOffRamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoOffRampsConfigured\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIBaseOffRamp\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"contractIEVM2AnyGEOnRamp\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBaseOffRamp\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIBaseOffRamp\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"contractIEVM2AnyGEOnRamp\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBaseOffRamp\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"addOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structGEConsumer.EVM2AnyGEMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structGEConsumer.EVM2AnyGEMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRamps\",\"outputs\":[{\"internalType\":\"contractIBaseOffRamp[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"contractIEVM2AnyGEOnRamp\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBaseOffRamp\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBaseOffRamp\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"removeOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCommon.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"routeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"contractIEVM2AnyGEOnRamp\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"setOnRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002190380380620021908339810160408190526200003491620002f9565b8033806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf816200019a565b50508151620000d79150600390602084019062000245565b5060005b815181101562000191576040518060400160405280826001600160601b0316815260200160011515815250600260008484815181106200011f576200011f620003cb565b6020908102919091018101516001600160a01b031682528181019290925260400160002082518154939092015115156c01000000000000000000000000026001600160681b03199093166001600160601b03909216919091179190911790556200018981620003e1565b9050620000db565b50505062000409565b336001600160a01b03821603620001f45760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200029d579160200282015b828111156200029d57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000266565b50620002ab929150620002af565b5090565b5b80821115620002ab5760008155600101620002b0565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b0381168114620002f457600080fd5b919050565b600060208083850312156200030d57600080fd5b82516001600160401b03808211156200032557600080fd5b818501915085601f8301126200033a57600080fd5b8151818111156200034f576200034f620002c6565b8060051b604051601f19603f83011681018181108582111715620003775762000377620002c6565b6040529182528482019250838101850191888311156200039657600080fd5b938501935b82851015620003bf57620003af85620002dc565b845293850193928501926200039b565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000600182016200040257634e487b7160e01b600052601160045260246000fd5b5060010190565b611d7780620004196000396000f3fe608060405234801561001057600080fd5b50600436106100f45760003560e01c806396f4e9f911610097578063a8d87a3b11610066578063a8d87a3b14610279578063adb9f71b146102ac578063f2fde38b146102bf578063fbca3b74146102d257600080fd5b806396f4e9f914610209578063991f65431461021c578063a40e69c71461022f578063a48a90581461024457600080fd5b806320487ded116100d357806320487ded146101a65780632e36d584146101c757806379ba5097146101dc5780638da5cb5b146101e457600080fd5b80624b61bb146100f9578063181f5a77146101215780631d7a74a01461016a575b600080fd5b61010c610107366004611439565b6102e5565b60405190151581526020015b60405180910390f35b61015d6040518060400160405280600e81526020017f4745526f7574657220312e302e3000000000000000000000000000000000000081525081565b6040516101189190611502565b61010c610178366004611515565b6001600160a01b03166000908152600260205260409020546c01000000000000000000000000900460ff1690565b6101b96101b4366004611722565b610447565b604051908152602001610118565b6101da6101d536600461181f565b61053b565b005b6101da610636565b6000546001600160a01b03165b6040516001600160a01b039091168152602001610118565b6101b9610217366004611722565b6106ff565b6101da61022a366004611515565b61093e565b610237610c59565b6040516101189190611856565b61010c6102523660046118a3565b67ffffffffffffffff166000908152600460205260409020546001600160a01b0316151590565b6101f16102873660046118a3565b67ffffffffffffffff166000908152600460205260409020546001600160a01b031690565b6101da6102ba366004611515565b610cbb565b6101da6102cd366004611515565b610e97565b6102376102e03660046118a3565b610eab565b336000818152600260205260408120549091906c01000000000000000000000000900460ff16610348576040517fa2c8bfb60000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6000633015b91c60e01b8760405160240161036391906119a7565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091529050856103dd576103d68585600084610f94565b925061043d565b836001600160a01b0316816040516103f59190611a88565b6000604051808303816000865af19150503d8060008114610432576040519150601f19603f3d011682016040523d82523d6000602084013e610437565b606091505b50909350505b5050949350505050565b67ffffffffffffffff82166000908152600460205260408120546001600160a01b0316806104ad576040517fae236d9c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8516600482015260240161033f565b6040517f38724a950000000000000000000000000000000000000000000000000000000081526001600160a01b038216906338724a95906104f2908690600401611b67565b602060405180830381865afa15801561050f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105339190611b7a565b949350505050565b610543610fe0565b67ffffffffffffffff82166000908152600460205260409020546001600160a01b038083169116036105bc576040517f74456f4900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff831660048201526001600160a01b038216602482015260440161033f565b67ffffffffffffffff821660008181526004602052604080822080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03861690811790915590519092917f1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f2391a35050565b6001546001600160a01b031633146106905760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161033f565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b67ffffffffffffffff82166000908152600460205260408120546001600160a01b03168161072d8585610447565b606085015190915061074a906001600160a01b031633848461103c565b60005b8460400151518110156108a85760008560400151828151811061077257610772611b93565b6020908102919091010151516040517f5d86f1410000000000000000000000000000000000000000000000000000000081526001600160a01b038083166004830152919250600091861690635d86f14190602401602060405180830381865afa1580156107e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108079190611bc2565b90506001600160a01b038116610854576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161033f565b61089533828960400151868151811061086f5761086f611b93565b602002602001015160200151856001600160a01b031661103c909392919063ffffffff16565b5050806108a190611c0e565b905061074d565b506040517fa7d3e02f0000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063a7d3e02f906108f290879085903390600401611c46565b6020604051808303816000875af1158015610911573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109359190611b7a565b95945050505050565b610946610fe0565b6003546000819003610984576040517f22babb3200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382166000908152600260209081526040918290208251808401909352546bffffffffffffffffffffffff811683526c01000000000000000000000000900460ff161515908201819052610a16576040517f8c97f1220000000000000000000000000000000000000000000000000000000081526001600160a01b038416600482015260240161033f565b60006003610a25600185611c78565b81548110610a3557610a35611b93565b6000918252602090912001548251600380546001600160a01b039093169350916bffffffffffffffffffffffff909116908110610a7457610a74611b93565b6000918252602090912001546001600160a01b03166003610a96600186611c78565b81548110610aa657610aa6611b93565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600383600001516bffffffffffffffffffffffff1681548110610afa57610afa611b93565b600091825260208083209190910180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0394851617905584519284168252600290526040902080547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff9092169190911790556003805480610b9457610b94611c8f565b6000828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690559092019092556001600160a01b0386168083526002909152604080832080547fffffffffffffffffffffffffffffffffffffff000000000000000000000000001690555190917fcf91daec21e3510e2f2aea4b09d08c235d5c6844980be709f282ef591dbf420c91a250505050565b60606003805480602002602001604051908101604052809291908181526020018280548015610cb157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610c93575b5050505050905090565b610cc3610fe0565b6001600160a01b038116610d03576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381166000908152600260209081526040918290208251808401909352546bffffffffffffffffffffffff811683526c01000000000000000000000000900460ff1615801591830191909152610d97576040517f3a4406b50000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161033f565b60016020808301828152600380546bffffffffffffffffffffffff90811686526001600160a01b03871660008181526002909552604080862088518154965115156c01000000000000000000000000027fffffffffffffffffffffffffffffffffffffff0000000000000000000000000090971694169390931794909417909155815494850182559083527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b90930180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055517f78f53b26906785548b265fa08f4197f9f3fff73fe0d504d30400aacb527f4ce09190a25050565b610e9f610fe0565b610ea8816110ca565b50565b6060610ed88267ffffffffffffffff166000908152600460205260409020546001600160a01b0316151590565b610ef057505060408051600081526020810190915290565b67ffffffffffffffff821660009081526004602081905260408083205481517fd3c7c2c700000000000000000000000000000000000000000000000000000000815291516001600160a01b039091169363d3c7c2c793838101939192918290030181865afa158015610f66573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f8e9190810190611cbe565b92915050565b60005a611388811015610fa657600080fd5b611388810390508560408204820311610fbe57600080fd5b50833b610fca57600080fd5b60008083516020850186888af195945050505050565b6000546001600160a01b0316331461103a5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161033f565b565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526110c490859061118b565b50505050565b336001600160a01b038216036111225760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161033f565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006111e0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166112759092919063ffffffff16565b80519091501561127057808060200190518101906111fe9190611d4d565b6112705760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161033f565b505050565b6060611284848460008561128e565b90505b9392505050565b6060824710156113065760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161033f565b843b6113545760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161033f565b600080866001600160a01b031685876040516113709190611a88565b60006040518083038185875af1925050503d80600081146113ad576040519150601f19603f3d011682016040523d82523d6000602084013e6113b2565b606091505b50915091506113c28282866113cd565b979650505050505050565b606083156113dc575081611287565b8251156113ec5782518084602001fd5b8160405162461bcd60e51b815260040161033f9190611502565b8015158114610ea857600080fd5b6001600160a01b0381168114610ea857600080fd5b803561143481611414565b919050565b6000806000806080858703121561144f57600080fd5b843567ffffffffffffffff81111561146657600080fd5b85016080818803121561147857600080fd5b9350602085013561148881611406565b925060408501359150606085013561149f81611414565b939692955090935050565b60005b838110156114c55781810151838201526020016114ad565b838111156110c45750506000910152565b600081518084526114ee8160208601602086016114aa565b601f01601f19169290920160200192915050565b60208152600061128760208301846114d6565b60006020828403121561152757600080fd5b813561128781611414565b803567ffffffffffffffff8116811461143457600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561159c5761159c61154a565b60405290565b60405160a0810167ffffffffffffffff8111828210171561159c5761159c61154a565b604051601f8201601f1916810167ffffffffffffffff811182821017156115ee576115ee61154a565b604052919050565b600082601f83011261160757600080fd5b813567ffffffffffffffff8111156116215761162161154a565b6116346020601f19601f840116016115c5565b81815284602083860101111561164957600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff8211156116805761168061154a565b5060051b60200190565b600082601f83011261169b57600080fd5b813560206116b06116ab83611666565b6115c5565b82815260069290921b840181019181810190868411156116cf57600080fd5b8286015b8481101561171757604081890312156116ec5760008081fd5b6116f4611579565b81356116ff81611414565b815281850135858201528352918301916040016116d3565b509695505050505050565b6000806040838503121561173557600080fd5b61173e83611532565b9150602083013567ffffffffffffffff8082111561175b57600080fd5b9084019060a0828703121561176f57600080fd5b6117776115a2565b82358281111561178657600080fd5b611792888286016115f6565b8252506020830135828111156117a757600080fd5b6117b3888286016115f6565b6020830152506040830135828111156117cb57600080fd5b6117d78882860161168a565b6040830152506117e960608401611429565b606082015260808301358281111561180057600080fd5b61180c888286016115f6565b6080830152508093505050509250929050565b6000806040838503121561183257600080fd5b61183b83611532565b9150602083013561184b81611414565b809150509250929050565b6020808252825182820181905260009190848201906040850190845b818110156118975783516001600160a01b031683529284019291840191600101611872565b50909695505050505050565b6000602082840312156118b557600080fd5b61128782611532565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126118f357600080fd5b830160208101925035905067ffffffffffffffff81111561191357600080fd5b80360382131561192257600080fd5b9250929050565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b8183526000602080850194508260005b8581101561199c57813561197781611414565b6001600160a01b03168752818301358388015260409687019690910190600101611964565b509495945050505050565b60208152600067ffffffffffffffff806119c085611532565b1660208401526119d360208501856118be565b608060408601526119e860a086018284611929565b9150506119f860408601866118be565b601f1980878503016060880152611a10848385611929565b9350606088013592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018312611a4957600080fd5b60209288019283019235915084821115611a6257600080fd5b8160061b3603831315611a7457600080fd5b8684030160808701526113c2838284611954565b60008251611a9a8184602087016114aa565b9190910192915050565b6000815160a08452611ab960a08501826114d6565b905060208084015185830382870152611ad283826114d6565b60408681015188830389830152805180845290850195509092506000918401905b80831015611b2557855180516001600160a01b0316835285015185830152948401946001929092019190830190611af3565b5060608701519450611b4260608901866001600160a01b03169052565b608087015194508781036080890152611b5b81866114d6565b98975050505050505050565b6020815260006112876020830184611aa4565b600060208284031215611b8c57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215611bd457600080fd5b815161128781611414565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c3f57611c3f611bdf565b5060010190565b606081526000611c596060830186611aa4565b90508360208301526001600160a01b0383166040830152949350505050565b600082821015611c8a57611c8a611bdf565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60006020808385031215611cd157600080fd5b825167ffffffffffffffff811115611ce857600080fd5b8301601f81018513611cf957600080fd5b8051611d076116ab82611666565b81815260059190911b82018301908381019087831115611d2657600080fd5b928401925b828410156113c2578351611d3e81611414565b82529284019290840190611d2b565b600060208284031215611d5f57600080fd5b81516112878161140656fea164736f6c634300080f000a",
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

func (_GERouter *GERouterCaller) GetSupportedTokens(opts *bind.CallOpts, destChainId uint64) ([]common.Address, error) {
	var out []interface{}
	err := _GERouter.contract.Call(opts, &out, "getSupportedTokens", destChainId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_GERouter *GERouterSession) GetSupportedTokens(destChainId uint64) ([]common.Address, error) {
	return _GERouter.Contract.GetSupportedTokens(&_GERouter.CallOpts, destChainId)
}

func (_GERouter *GERouterCallerSession) GetSupportedTokens(destChainId uint64) ([]common.Address, error) {
	return _GERouter.Contract.GetSupportedTokens(&_GERouter.CallOpts, destChainId)
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

func (_GERouter *GERouterTransactor) RouteMessage(opts *bind.TransactOpts, message CommonAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _GERouter.contract.Transact(opts, "routeMessage", message, manualExecution, gasLimit, receiver)
}

func (_GERouter *GERouterSession) RouteMessage(message CommonAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _GERouter.Contract.RouteMessage(&_GERouter.TransactOpts, message, manualExecution, gasLimit, receiver)
}

func (_GERouter *GERouterTransactorSession) RouteMessage(message CommonAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _GERouter.Contract.RouteMessage(&_GERouter.TransactOpts, message, manualExecution, gasLimit, receiver)
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

	GetSupportedTokens(opts *bind.CallOpts, destChainId uint64) ([]common.Address, error)

	IsChainSupported(opts *bind.CallOpts, chainId uint64) (bool, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainId uint64, message GEConsumerEVM2AnyGEMessage) (*types.Transaction, error)

	RemoveOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error)

	RouteMessage(opts *bind.TransactOpts, message CommonAny2EVMMessage, manualExecution bool, gasLimit *big.Int, receiver common.Address) (*types.Transaction, error)

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
