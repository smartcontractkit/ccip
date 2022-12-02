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

type CCIPAny2EVMMessageFromSender struct {
	SourceChainId        *big.Int
	Sender               []byte
	Receiver             common.Address
	Data                 []byte
	DestPools            []common.Address
	DestTokensAndAmounts []CCIPEVMTokenAndAmount
	GasLimit             *big.Int
}

type CCIPEVM2AnyGEMessage struct {
	Receiver         []byte
	Data             []byte
	TokensAndAmounts []CCIPEVMTokenAndAmount
	FeeToken         common.Address
	ExtraArgs        []byte
}

type CCIPEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

var GERouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"AlreadyConfigured\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MustCallFromOffRamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoOffRampsConfigured\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractEVM2EVMGEOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractEVM2EVMGEOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"addOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.EVM2AnyGEMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structCCIP.EVM2AnyGEMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRamps\",\"outputs\":[{\"internalType\":\"contractBaseOffRampInterface[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"contractEVM2EVMGEOnRampInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"removeOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"routeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractEVM2EVMGEOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"setOnRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002222380380620022228339810160408190526200003491620002f9565b8033806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf816200019a565b50508151620000d79150600390602084019062000245565b5060005b815181101562000191576040518060400160405280826001600160601b0316815260200160011515815250600260008484815181106200011f576200011f620003cb565b6020908102919091018101516001600160a01b031682528181019290925260400160002082518154939092015115156c01000000000000000000000000026001600160681b03199093166001600160601b03909216919091179190911790556200018981620003e1565b9050620000db565b50505062000409565b336001600160a01b03821603620001f45760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200029d579160200282015b828111156200029d57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000266565b50620002ab929150620002af565b5090565b5b80821115620002ab5760008155600101620002b0565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b0381168114620002f457600080fd5b919050565b600060208083850312156200030d57600080fd5b82516001600160401b03808211156200032557600080fd5b818501915085601f8301126200033a57600080fd5b8151818111156200034f576200034f620002c6565b8060051b604051601f19603f83011681018181108582111715620003775762000377620002c6565b6040529182528482019250838101850191888311156200039657600080fd5b938501935b82851015620003bf57620003af85620002dc565b845293850193928501926200039b565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000600182016200040257634e487b7160e01b600052601160045260246000fd5b5060010190565b611e0980620004196000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063a40e69c71161008c578063da52b4c411610066578063da52b4c41461026c578063ecaadb841461027f578063f1927cae14610292578063f2fde38b146102a557600080fd5b8063a40e69c71461021b578063adb9f71b14610230578063d8a98f8c1461024357600080fd5b806367fcbdd8116100c857806367fcbdd8146101b857806379ba5097146101d95780638da5cb5b146101e3578063991f65431461020857600080fd5b8063181f5a77146100ef5780631d7a74a0146101415780635221c1f01461018d575b600080fd5b61012b6040518060400160405280600e81526020017f4745526f7574657220312e302e3000000000000000000000000000000000000081525081565b60405161013891906115d1565b60405180910390f35b61017d61014f3660046115f9565b6001600160a01b03166000908152600260205260409020546c01000000000000000000000000900460ff1690565b6040519015158152602001610138565b61017d61019b366004611616565b6000908152600460205260409020546001600160a01b0316151590565b6101cb6101c636600461183a565b6102b8565b604051908152602001610138565b6101e161050b565b005b6000546001600160a01b03165b6040516001600160a01b039091168152602001610138565b6101e16102163660046115f9565b6105d4565b6102236108ef565b6040516101389190611930565b6101e161023e3660046115f9565b610951565b6101f0610251366004611616565b6000908152600460205260409020546001600160a01b031690565b61017d61027a36600461197d565b610b2d565b6101cb61028d36600461183a565b610c47565b6101e16102a03660046119b8565b610d27565b6101e16102b33660046115f9565b610e05565b6000828152600460205260408120546001600160a01b0316816102db8585610c47565b90506000610317604051806040016040528087606001516001600160a01b03168152602001848152508660400151610e1990919063ffffffff16565b905060005b8151811015610472576000828281518110610339576103396119e8565b6020908102919091010151516040517f5d86f1410000000000000000000000000000000000000000000000000000000081526001600160a01b038083166004830152919250600091871690635d86f141906024016020604051808303816000875af11580156103ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d09190611a17565b90506001600160a01b038116610422576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b61045f3382868681518110610439576104396119e8565b602002602001015160200151856001600160a01b03166110a4909392919063ffffffff16565b50508061046b90611a63565b905061031c565b506040517fa7d3e02f0000000000000000000000000000000000000000000000000000000081526001600160a01b0384169063a7d3e02f906104bc90889086903390600401611b5f565b6020604051808303816000875af11580156104db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ff9190611b91565b93505050505b92915050565b6001546001600160a01b031633146105655760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610419565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6105dc611132565b600354600081900361061a576040517f22babb3200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382166000908152600260209081526040918290208251808401909352546bffffffffffffffffffffffff811683526c01000000000000000000000000900460ff1615159082018190526106ac576040517f8c97f1220000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401610419565b600060036106bb600185611baa565b815481106106cb576106cb6119e8565b6000918252602090912001548251600380546001600160a01b039093169350916bffffffffffffffffffffffff90911690811061070a5761070a6119e8565b6000918252602090912001546001600160a01b0316600361072c600186611baa565b8154811061073c5761073c6119e8565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600383600001516bffffffffffffffffffffffff1681548110610790576107906119e8565b600091825260208083209190910180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0394851617905584519284168252600290526040902080547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff909216919091179055600380548061082a5761082a611bc1565b6000828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690559092019092556001600160a01b0386168083526002909152604080832080547fffffffffffffffffffffffffffffffffffffff000000000000000000000000001690555190917fcf91daec21e3510e2f2aea4b09d08c235d5c6844980be709f282ef591dbf420c91a250505050565b6060600380548060200260200160405190810160405280929190818152602001828054801561094757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610929575b5050505050905090565b610959611132565b6001600160a01b038116610999576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381166000908152600260209081526040918290208251808401909352546bffffffffffffffffffffffff811683526c01000000000000000000000000900460ff1615801591830191909152610a2d576040517f3a4406b50000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610419565b60016020808301828152600380546bffffffffffffffffffffffff90811686526001600160a01b03871660008181526002909552604080862088518154965115156c01000000000000000000000000027fffffffffffffffffffffffffffffffffffffff0000000000000000000000000090971694169390931794909417909155815494850182559083527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b90930180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055517f78f53b26906785548b265fa08f4197f9f3fff73fe0d504d30400aacb527f4ce09190a25050565b336000818152600260205260408120549091906c01000000000000000000000000900460ff16610b8b576040517fa2c8bfb6000000000000000000000000000000000000000000000000000000008152336004820152602401610419565b6000610b9e610b9985611c54565b61118e565b9050600063b06193dd60e01b82604051602401610bbb9190611d31565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909416939093179092529150610c3e9060c087013590610c3690606089019089016115f9565b6000846111f1565b95945050505050565b6000828152600460205260408120546001600160a01b031680610c99576040517f45abe4ae00000000000000000000000000000000000000000000000000000000815260048101859052602401610419565b6040517f38724a950000000000000000000000000000000000000000000000000000000081526001600160a01b038216906338724a9590610cde908690600401611d93565b602060405180830381865afa158015610cfb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1f9190611b91565b949350505050565b610d2f611132565b6000828152600460205260409020546001600160a01b03808316911603610d94576040517fe31de3b2000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b0382166024820152604401610419565b60008281526004602052604080822080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0385169081179091559051909184917f4b680ef9fa79bb5f36e7559d7b33fd57a8336f78cc120c8cd93333b5ade624cb9190a35050565b610e0d611132565b610e168161123d565b50565b606060005b8351811015610fb85782600001516001600160a01b0316848281518110610e4757610e476119e8565b6020026020010151600001516001600160a01b031603610fa8576000845167ffffffffffffffff811115610e7d57610e7d61162f565b604051908082528060200260200182016040528015610ec257816020015b6040805180820190915260008082526020820152815260200190600190039081610e9b5790505b50905060005b8551811015610f1957858181518110610ee357610ee36119e8565b6020026020010151828281518110610efd57610efd6119e8565b602002602001018190525080610f1290611a63565b9050610ec8565b506040518060400160405280828481518110610f3757610f376119e8565b6020026020010151600001516001600160a01b031681526020018560200151838581518110610f6857610f686119e8565b602002602001015160200151610f7e9190611da6565b815250818381518110610f9357610f936119e8565b60200260200101819052508092505050610505565b610fb181611a63565b9050610e1e565b50600083516001610fc99190611da6565b67ffffffffffffffff811115610fe157610fe161162f565b60405190808252806020026020018201604052801561102657816020015b6040805180820190915260008082526020820152815260200190600190039081610fff5790505b50905060005b845181101561107d57848181518110611047576110476119e8565b6020026020010151828281518110611061576110616119e8565b60200260200101819052508061107690611a63565b905061102c565b508281855181518110611092576110926119e8565b60209081029190910101529392505050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261112c9085906112fe565b50505050565b6000546001600160a01b0316331461118c5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610419565b565b6111b96040518060800160405280600081526020016060815260200160608152602001606081525090565b60405180608001604052808360000151815260200183602001518152602001836060015181526020018360a001518152509050919050565b60005a61138881101561120357600080fd5b61138881039050856040820482031161121b57600080fd5b50833b61122757600080fd5b60008083516020850186888af195945050505050565b336001600160a01b038216036112955760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610419565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000611353826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166113e89092919063ffffffff16565b8051909150156113e357808060200190518101906113719190611dbe565b6113e35760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610419565b505050565b60606113f78484600085611401565b90505b9392505050565b6060824710156114795760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610419565b843b6114c75760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610419565b600080866001600160a01b031685876040516114e39190611de0565b60006040518083038185875af1925050503d8060008114611520576040519150601f19603f3d011682016040523d82523d6000602084013e611525565b606091505b5091509150611535828286611540565b979650505050505050565b6060831561154f5750816113fa565b82511561155f5782518084602001fd5b8160405162461bcd60e51b815260040161041991906115d1565b60005b8381101561159457818101518382015260200161157c565b8381111561112c5750506000910152565b600081518084526115bd816020860160208601611579565b601f01601f19169290920160200192915050565b6020815260006113fa60208301846115a5565b6001600160a01b0381168114610e1657600080fd5b60006020828403121561160b57600080fd5b81356113fa816115e4565b60006020828403121561162857600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156116815761168161162f565b60405290565b60405160a0810167ffffffffffffffff811182821017156116815761168161162f565b60405160e0810167ffffffffffffffff811182821017156116815761168161162f565b604051601f8201601f1916810167ffffffffffffffff811182821017156116f6576116f661162f565b604052919050565b600082601f83011261170f57600080fd5b813567ffffffffffffffff8111156117295761172961162f565b61173c6020601f19601f840116016116cd565b81815284602083860101111561175157600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff8211156117885761178861162f565b5060051b60200190565b803561179d816115e4565b919050565b600082601f8301126117b357600080fd5b813560206117c86117c38361176e565b6116cd565b82815260069290921b840181019181810190868411156117e757600080fd5b8286015b8481101561182f57604081890312156118045760008081fd5b61180c61165e565b8135611817816115e4565b815281850135858201528352918301916040016117eb565b509695505050505050565b6000806040838503121561184d57600080fd5b82359150602083013567ffffffffffffffff8082111561186c57600080fd5b9084019060a0828703121561188057600080fd5b611888611687565b82358281111561189757600080fd5b6118a3888286016116fe565b8252506020830135828111156118b857600080fd5b6118c4888286016116fe565b6020830152506040830135828111156118dc57600080fd5b6118e8888286016117a2565b6040830152506118fa60608401611792565b606082015260808301358281111561191157600080fd5b61191d888286016116fe565b6080830152508093505050509250929050565b6020808252825182820181905260009190848201906040850190845b818110156119715783516001600160a01b03168352928401929184019160010161194c565b50909695505050505050565b60006020828403121561198f57600080fd5b813567ffffffffffffffff8111156119a657600080fd5b820160e081850312156113fa57600080fd5b600080604083850312156119cb57600080fd5b8235915060208301356119dd816115e4565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215611a2957600080fd5b81516113fa816115e4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a9457611a94611a34565b5060010190565b600081518084526020808501945080840160005b83811015611adf57815180516001600160a01b031688528301518388015260409096019590820190600101611aaf565b509495945050505050565b6000815160a08452611aff60a08501826115a5565b905060208301518482036020860152611b1882826115a5565b91505060408301518482036040860152611b328282611a9b565b9150506001600160a01b03606084015116606085015260808301518482036080860152610c3e82826115a5565b606081526000611b726060830186611aea565b90508360208301526001600160a01b0383166040830152949350505050565b600060208284031215611ba357600080fd5b5051919050565b600082821015611bbc57611bbc611a34565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600082601f830112611c0157600080fd5b81356020611c116117c38361176e565b82815260059290921b84018101918181019086841115611c3057600080fd5b8286015b8481101561182f578035611c47816115e4565b8352918301918301611c34565b600060e08236031215611c6657600080fd5b611c6e6116aa565b82358152602083013567ffffffffffffffff80821115611c8d57600080fd5b611c99368387016116fe565b6020840152611caa60408601611792565b60408401526060850135915080821115611cc357600080fd5b611ccf368387016116fe565b60608401526080850135915080821115611ce857600080fd5b611cf436838701611bf0565b608084015260a0850135915080821115611d0d57600080fd5b50611d1a368286016117a2565b60a08301525060c092830135928101929092525090565b60208152815160208201526000602083015160806040840152611d5760a08401826115a5565b90506040840151601f1980858403016060860152611d7583836115a5565b9250606086015191508085840301608086015250610c3e8282611a9b565b6020815260006113fa6020830184611aea565b60008219821115611db957611db9611a34565b500190565b600060208284031215611dd057600080fd5b815180151581146113fa57600080fd5b60008251611df2818460208701611579565b919091019291505056fea164736f6c634300080f000a",
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

func (_GERouter *GERouterCaller) GetFee(opts *bind.CallOpts, destinationChainId *big.Int, message CCIPEVM2AnyGEMessage) (*big.Int, error) {
	var out []interface{}
	err := _GERouter.contract.Call(opts, &out, "getFee", destinationChainId, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_GERouter *GERouterSession) GetFee(destinationChainId *big.Int, message CCIPEVM2AnyGEMessage) (*big.Int, error) {
	return _GERouter.Contract.GetFee(&_GERouter.CallOpts, destinationChainId, message)
}

func (_GERouter *GERouterCallerSession) GetFee(destinationChainId *big.Int, message CCIPEVM2AnyGEMessage) (*big.Int, error) {
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

func (_GERouter *GERouterCaller) GetOnRamp(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GERouter.contract.Call(opts, &out, "getOnRamp", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_GERouter *GERouterSession) GetOnRamp(chainId *big.Int) (common.Address, error) {
	return _GERouter.Contract.GetOnRamp(&_GERouter.CallOpts, chainId)
}

func (_GERouter *GERouterCallerSession) GetOnRamp(chainId *big.Int) (common.Address, error) {
	return _GERouter.Contract.GetOnRamp(&_GERouter.CallOpts, chainId)
}

func (_GERouter *GERouterCaller) IsChainSupported(opts *bind.CallOpts, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _GERouter.contract.Call(opts, &out, "isChainSupported", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_GERouter *GERouterSession) IsChainSupported(chainId *big.Int) (bool, error) {
	return _GERouter.Contract.IsChainSupported(&_GERouter.CallOpts, chainId)
}

func (_GERouter *GERouterCallerSession) IsChainSupported(chainId *big.Int) (bool, error) {
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

func (_GERouter *GERouterTransactor) CcipSend(opts *bind.TransactOpts, destinationChainId *big.Int, message CCIPEVM2AnyGEMessage) (*types.Transaction, error) {
	return _GERouter.contract.Transact(opts, "ccipSend", destinationChainId, message)
}

func (_GERouter *GERouterSession) CcipSend(destinationChainId *big.Int, message CCIPEVM2AnyGEMessage) (*types.Transaction, error) {
	return _GERouter.Contract.CcipSend(&_GERouter.TransactOpts, destinationChainId, message)
}

func (_GERouter *GERouterTransactorSession) CcipSend(destinationChainId *big.Int, message CCIPEVM2AnyGEMessage) (*types.Transaction, error) {
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

func (_GERouter *GERouterTransactor) RouteMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _GERouter.contract.Transact(opts, "routeMessage", message)
}

func (_GERouter *GERouterSession) RouteMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _GERouter.Contract.RouteMessage(&_GERouter.TransactOpts, message)
}

func (_GERouter *GERouterTransactorSession) RouteMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _GERouter.Contract.RouteMessage(&_GERouter.TransactOpts, message)
}

func (_GERouter *GERouterTransactor) SetOnRamp(opts *bind.TransactOpts, chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
	return _GERouter.contract.Transact(opts, "setOnRamp", chainId, onRamp)
}

func (_GERouter *GERouterSession) SetOnRamp(chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
	return _GERouter.Contract.SetOnRamp(&_GERouter.TransactOpts, chainId, onRamp)
}

func (_GERouter *GERouterTransactorSession) SetOnRamp(chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
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
	ChainId *big.Int
	OnRamp  common.Address
	Raw     types.Log
}

func (_GERouter *GERouterFilterer) FilterOnRampSet(opts *bind.FilterOpts, chainId []*big.Int, onRamp []common.Address) (*GERouterOnRampSetIterator, error) {

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

func (_GERouter *GERouterFilterer) WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *GERouterOnRampSet, chainId []*big.Int, onRamp []common.Address) (event.Subscription, error) {

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
	return common.HexToHash("0x4b680ef9fa79bb5f36e7559d7b33fd57a8336f78cc120c8cd93333b5ade624cb")
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
	GetFee(opts *bind.CallOpts, destinationChainId *big.Int, message CCIPEVM2AnyGEMessage) (*big.Int, error)

	GetOffRamps(opts *bind.CallOpts) ([]common.Address, error)

	GetOnRamp(opts *bind.CallOpts, chainId *big.Int) (common.Address, error)

	IsChainSupported(opts *bind.CallOpts, chainId *big.Int) (bool, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainId *big.Int, message CCIPEVM2AnyGEMessage) (*types.Transaction, error)

	RemoveOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error)

	RouteMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error)

	SetOnRamp(opts *bind.TransactOpts, chainId *big.Int, onRamp common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOffRampAdded(opts *bind.FilterOpts, offRamp []common.Address) (*GERouterOffRampAddedIterator, error)

	WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *GERouterOffRampAdded, offRamp []common.Address) (event.Subscription, error)

	ParseOffRampAdded(log types.Log) (*GERouterOffRampAdded, error)

	FilterOffRampRemoved(opts *bind.FilterOpts, offRamp []common.Address) (*GERouterOffRampRemovedIterator, error)

	WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *GERouterOffRampRemoved, offRamp []common.Address) (event.Subscription, error)

	ParseOffRampRemoved(log types.Log) (*GERouterOffRampRemoved, error)

	FilterOnRampSet(opts *bind.FilterOpts, chainId []*big.Int, onRamp []common.Address) (*GERouterOnRampSetIterator, error)

	WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *GERouterOnRampSet, chainId []*big.Int, onRamp []common.Address) (event.Subscription, error)

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
