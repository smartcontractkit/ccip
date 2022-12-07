// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_any_toll_onramp_router

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

type TollConsumerEVM2AnyTollMessage struct {
	Receiver          []byte
	Data              []byte
	TokensAndAmounts  []CommonEVMTokenAndAmount
	FeeTokenAndAmount CommonEVMTokenAndAmount
	ExtraArgs         []byte
}

var EVM2AnyTollOnRampRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"FeeTokenAmountTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"contractEVM2EVMTollOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"contractEVM2EVMTollOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount\",\"name\":\"feeTokenAndAmount\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structTollConsumer.EVM2AnyTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"contractEVM2EVMTollOnRampInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"contractEVM2EVMTollOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"setOnRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b611422806101576000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c806379ba509711610076578063a48a90581161005b578063a48a90581461017b578063a8d87a3b146101c0578063f2fde38b146101f357600080fd5b806379ba50971461014e5780638da5cb5b1461015657600080fd5b8063181f5a77146100a85780632e36d584146100fa57806338fa2f9b1461010f57806359e96b5b1461013b575b600080fd5b6100e46040518060400160405280601d81526020017f45564d32416e79546f6c6c4f6e52616d70526f7574657220312e302e3000000081525081565b6040516100f19190610df8565b60405180910390f35b61010d610108366004610e36565b610206565b005b61012261011d36600461104a565b610306565b60405167ffffffffffffffff90911681526020016100f1565b61010d61014936600461114a565b61043e565b61010d6104aa565b6000546001600160a01b03165b6040516001600160a01b0390911681526020016100f1565b6101b061018936600461118b565b67ffffffffffffffff166000908152600260205260409020546001600160a01b0316151590565b60405190151581526020016100f1565b6101636101ce36600461118b565b67ffffffffffffffff166000908152600260205260409020546001600160a01b031690565b61010d6102013660046111a8565b610573565b61020e610587565b67ffffffffffffffff82166000908152600260205260409020546001600160a01b0380831691160361028c576040517f74456f4900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff831660048201526001600160a01b03821660248201526044015b60405180910390fd5b67ffffffffffffffff821660008181526002602052604080822080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03861690811790915590519092917f1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f2391a35050565b67ffffffffffffffff82166000908152600260205260408120546001600160a01b03168061036c576040517fae236d9c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610283565b606083015180516020909101516103849183916105e3565b836060015160200181815161039991906111f4565b90525060408301516103ac908290610810565b6040517f4bd838ad0000000000000000000000000000000000000000000000000000000081526001600160a01b03821690634bd838ad906103f3908690339060040161120b565b6020604051808303816000875af1158015610412573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610436919061131d565b949350505050565b610446610587565b61045a6001600160a01b0384168383610969565b604080516001600160a01b038086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa89060600160405180910390a1505050565b6001546001600160a01b031633146105045760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610283565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61057b610587565b61058481610a12565b50565b6000546001600160a01b031633146105e15760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610283565b565b6040517f5d86f1410000000000000000000000000000000000000000000000000000000081526001600160a01b0383811660048301526000918291861690635d86f141906024016020604051808303816000875af1158015610649573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066d919061133a565b90506001600160a01b0381166106ba576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401610283565b6040517fd0d5de610000000000000000000000000000000000000000000000000000000081526001600160a01b03858116600483015286169063d0d5de61906024016020604051808303816000875af115801561071b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061073f9190611357565b91503382156107a35783831115610782576040517f0443cfcc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61078c83856111f4565b93506107a36001600160a01b038616823086610ad3565b83156107be576107be6001600160a01b038616828487610ad3565b604080516001600160a01b03831681523060208201529081018490527f945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d9060600160405180910390a150509392505050565b60005b815181101561096457600082828151811061083057610830611370565b6020908102919091010151516040517f5d86f1410000000000000000000000000000000000000000000000000000000081526001600160a01b038083166004830152919250600091861690635d86f141906024016020604051808303816000875af11580156108a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c7919061133a565b90506001600160a01b038116610914576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610283565b610951338286868151811061092b5761092b611370565b602002602001015160200151856001600160a01b0316610ad3909392919063ffffffff16565b50508061095d9061139f565b9050610813565b505050565b6040516001600160a01b0383166024820152604481018290526109649084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610b2a565b336001600160a01b03821603610a6a5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610283565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040516001600160a01b0380851660248301528316604482015260648101829052610b249085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016109ae565b50505050565b6000610b7f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610c0f9092919063ffffffff16565b8051909150156109645780806020019051810190610b9d91906113d7565b6109645760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610283565b6060610c1e8484600085610c28565b90505b9392505050565b606082471015610ca05760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610283565b843b610cee5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610283565b600080866001600160a01b03168587604051610d0a91906113f9565b60006040518083038185875af1925050503d8060008114610d47576040519150601f19603f3d011682016040523d82523d6000602084013e610d4c565b606091505b5091509150610d5c828286610d67565b979650505050505050565b60608315610d76575081610c21565b825115610d865782518084602001fd5b8160405162461bcd60e51b81526004016102839190610df8565b60005b83811015610dbb578181015183820152602001610da3565b83811115610b245750506000910152565b60008151808452610de4816020860160208601610da0565b601f01601f19169290920160200192915050565b602081526000610c216020830184610dcc565b67ffffffffffffffff8116811461058457600080fd5b6001600160a01b038116811461058457600080fd5b60008060408385031215610e4957600080fd5b8235610e5481610e0b565b91506020830135610e6481610e21565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff81118282101715610ec157610ec1610e6f565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610ef057610ef0610e6f565b604052919050565b600082601f830112610f0957600080fd5b813567ffffffffffffffff811115610f2357610f23610e6f565b610f366020601f19601f84011601610ec7565b818152846020838601011115610f4b57600080fd5b816020850160208301376000918101602001919091529392505050565b600060408284031215610f7a57600080fd5b6040516040810181811067ffffffffffffffff82111715610f9d57610f9d610e6f565b6040529050808235610fae81610e21565b8152602092830135920191909152919050565b600082601f830112610fd257600080fd5b8135602067ffffffffffffffff821115610fee57610fee610e6f565b610ffc818360051b01610ec7565b82815260069290921b8401810191818101908684111561101b57600080fd5b8286015b8481101561103f576110318882610f68565b83529183019160400161101f565b509695505050505050565b6000806040838503121561105d57600080fd5b823561106881610e0b565b9150602083013567ffffffffffffffff8082111561108557600080fd5b9084019060c0828703121561109957600080fd5b6110a1610e9e565b8235828111156110b057600080fd5b6110bc88828601610ef8565b8252506020830135828111156110d157600080fd5b6110dd88828601610ef8565b6020830152506040830135828111156110f557600080fd5b61110188828601610fc1565b6040830152506111148760608501610f68565b606082015260a08301358281111561112b57600080fd5b61113788828601610ef8565b6080830152508093505050509250929050565b60008060006060848603121561115f57600080fd5b833561116a81610e21565b9250602084013561117a81610e21565b929592945050506040919091013590565b60006020828403121561119d57600080fd5b8135610c2181610e0b565b6000602082840312156111ba57600080fd5b8135610c2181610e21565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611206576112066111c5565b500390565b60006040808352845160c082850152611228610100850182610dcc565b90506020808701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808785030160608801526112658483610dcc565b89860151888203830160808a01528051808352908501955060009350908401905b808410156112c0576112ac82875180516001600160a01b03168252602090810151910152565b948401946001939093019290860190611286565b5060608a015180516001600160a01b031660a08a01526020015160c089015260808a0151888203830160e08a015295506112fa8187610dcc565b9650505050611313818601876001600160a01b03169052565b5050509392505050565b60006020828403121561132f57600080fd5b8151610c2181610e0b565b60006020828403121561134c57600080fd5b8151610c2181610e21565b60006020828403121561136957600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036113d0576113d06111c5565b5060010190565b6000602082840312156113e957600080fd5b81518015158114610c2157600080fd5b6000825161140b818460208701610da0565b919091019291505056fea164736f6c634300080f000a",
}

var EVM2AnyTollOnRampRouterABI = EVM2AnyTollOnRampRouterMetaData.ABI

var EVM2AnyTollOnRampRouterBin = EVM2AnyTollOnRampRouterMetaData.Bin

func DeployEVM2AnyTollOnRampRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EVM2AnyTollOnRampRouter, error) {
	parsed, err := EVM2AnyTollOnRampRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2AnyTollOnRampRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2AnyTollOnRampRouter{EVM2AnyTollOnRampRouterCaller: EVM2AnyTollOnRampRouterCaller{contract: contract}, EVM2AnyTollOnRampRouterTransactor: EVM2AnyTollOnRampRouterTransactor{contract: contract}, EVM2AnyTollOnRampRouterFilterer: EVM2AnyTollOnRampRouterFilterer{contract: contract}}, nil
}

type EVM2AnyTollOnRampRouter struct {
	address common.Address
	abi     abi.ABI
	EVM2AnyTollOnRampRouterCaller
	EVM2AnyTollOnRampRouterTransactor
	EVM2AnyTollOnRampRouterFilterer
}

type EVM2AnyTollOnRampRouterCaller struct {
	contract *bind.BoundContract
}

type EVM2AnyTollOnRampRouterTransactor struct {
	contract *bind.BoundContract
}

type EVM2AnyTollOnRampRouterFilterer struct {
	contract *bind.BoundContract
}

type EVM2AnyTollOnRampRouterSession struct {
	Contract     *EVM2AnyTollOnRampRouter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2AnyTollOnRampRouterCallerSession struct {
	Contract *EVM2AnyTollOnRampRouterCaller
	CallOpts bind.CallOpts
}

type EVM2AnyTollOnRampRouterTransactorSession struct {
	Contract     *EVM2AnyTollOnRampRouterTransactor
	TransactOpts bind.TransactOpts
}

type EVM2AnyTollOnRampRouterRaw struct {
	Contract *EVM2AnyTollOnRampRouter
}

type EVM2AnyTollOnRampRouterCallerRaw struct {
	Contract *EVM2AnyTollOnRampRouterCaller
}

type EVM2AnyTollOnRampRouterTransactorRaw struct {
	Contract *EVM2AnyTollOnRampRouterTransactor
}

func NewEVM2AnyTollOnRampRouter(address common.Address, backend bind.ContractBackend) (*EVM2AnyTollOnRampRouter, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2AnyTollOnRampRouterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2AnyTollOnRampRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2AnyTollOnRampRouter{address: address, abi: abi, EVM2AnyTollOnRampRouterCaller: EVM2AnyTollOnRampRouterCaller{contract: contract}, EVM2AnyTollOnRampRouterTransactor: EVM2AnyTollOnRampRouterTransactor{contract: contract}, EVM2AnyTollOnRampRouterFilterer: EVM2AnyTollOnRampRouterFilterer{contract: contract}}, nil
}

func NewEVM2AnyTollOnRampRouterCaller(address common.Address, caller bind.ContractCaller) (*EVM2AnyTollOnRampRouterCaller, error) {
	contract, err := bindEVM2AnyTollOnRampRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2AnyTollOnRampRouterCaller{contract: contract}, nil
}

func NewEVM2AnyTollOnRampRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2AnyTollOnRampRouterTransactor, error) {
	contract, err := bindEVM2AnyTollOnRampRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2AnyTollOnRampRouterTransactor{contract: contract}, nil
}

func NewEVM2AnyTollOnRampRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2AnyTollOnRampRouterFilterer, error) {
	contract, err := bindEVM2AnyTollOnRampRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2AnyTollOnRampRouterFilterer{contract: contract}, nil
}

func bindEVM2AnyTollOnRampRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2AnyTollOnRampRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2AnyTollOnRampRouter.Contract.EVM2AnyTollOnRampRouterCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.EVM2AnyTollOnRampRouterTransactor.contract.Transfer(opts)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.EVM2AnyTollOnRampRouterTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2AnyTollOnRampRouter.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.contract.Transfer(opts)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterCaller) GetOnRamp(opts *bind.CallOpts, chainId uint64) (common.Address, error) {
	var out []interface{}
	err := _EVM2AnyTollOnRampRouter.contract.Call(opts, &out, "getOnRamp", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterSession) GetOnRamp(chainId uint64) (common.Address, error) {
	return _EVM2AnyTollOnRampRouter.Contract.GetOnRamp(&_EVM2AnyTollOnRampRouter.CallOpts, chainId)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterCallerSession) GetOnRamp(chainId uint64) (common.Address, error) {
	return _EVM2AnyTollOnRampRouter.Contract.GetOnRamp(&_EVM2AnyTollOnRampRouter.CallOpts, chainId)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterCaller) IsChainSupported(opts *bind.CallOpts, chainId uint64) (bool, error) {
	var out []interface{}
	err := _EVM2AnyTollOnRampRouter.contract.Call(opts, &out, "isChainSupported", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterSession) IsChainSupported(chainId uint64) (bool, error) {
	return _EVM2AnyTollOnRampRouter.Contract.IsChainSupported(&_EVM2AnyTollOnRampRouter.CallOpts, chainId)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterCallerSession) IsChainSupported(chainId uint64) (bool, error) {
	return _EVM2AnyTollOnRampRouter.Contract.IsChainSupported(&_EVM2AnyTollOnRampRouter.CallOpts, chainId)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2AnyTollOnRampRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterSession) Owner() (common.Address, error) {
	return _EVM2AnyTollOnRampRouter.Contract.Owner(&_EVM2AnyTollOnRampRouter.CallOpts)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterCallerSession) Owner() (common.Address, error) {
	return _EVM2AnyTollOnRampRouter.Contract.Owner(&_EVM2AnyTollOnRampRouter.CallOpts)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2AnyTollOnRampRouter.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterSession) TypeAndVersion() (string, error) {
	return _EVM2AnyTollOnRampRouter.Contract.TypeAndVersion(&_EVM2AnyTollOnRampRouter.CallOpts)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterCallerSession) TypeAndVersion() (string, error) {
	return _EVM2AnyTollOnRampRouter.Contract.TypeAndVersion(&_EVM2AnyTollOnRampRouter.CallOpts)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.AcceptOwnership(&_EVM2AnyTollOnRampRouter.TransactOpts)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.AcceptOwnership(&_EVM2AnyTollOnRampRouter.TransactOpts)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactor) CcipSend(opts *bind.TransactOpts, destinationChainId uint64, message TollConsumerEVM2AnyTollMessage) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.contract.Transact(opts, "ccipSend", destinationChainId, message)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterSession) CcipSend(destinationChainId uint64, message TollConsumerEVM2AnyTollMessage) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.CcipSend(&_EVM2AnyTollOnRampRouter.TransactOpts, destinationChainId, message)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactorSession) CcipSend(destinationChainId uint64, message TollConsumerEVM2AnyTollMessage) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.CcipSend(&_EVM2AnyTollOnRampRouter.TransactOpts, destinationChainId, message)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactor) SetOnRamp(opts *bind.TransactOpts, chainId uint64, onRamp common.Address) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.contract.Transact(opts, "setOnRamp", chainId, onRamp)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterSession) SetOnRamp(chainId uint64, onRamp common.Address) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.SetOnRamp(&_EVM2AnyTollOnRampRouter.TransactOpts, chainId, onRamp)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactorSession) SetOnRamp(chainId uint64, onRamp common.Address) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.SetOnRamp(&_EVM2AnyTollOnRampRouter.TransactOpts, chainId, onRamp)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.TransferOwnership(&_EVM2AnyTollOnRampRouter.TransactOpts, to)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.TransferOwnership(&_EVM2AnyTollOnRampRouter.TransactOpts, to)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactor) WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.contract.Transact(opts, "withdrawAccumulatedFees", feeToken, recipient, amount)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.WithdrawAccumulatedFees(&_EVM2AnyTollOnRampRouter.TransactOpts, feeToken, recipient, amount)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactorSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.WithdrawAccumulatedFees(&_EVM2AnyTollOnRampRouter.TransactOpts, feeToken, recipient, amount)
}

type EVM2AnyTollOnRampRouterFeeChargedIterator struct {
	Event *EVM2AnyTollOnRampRouterFeeCharged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2AnyTollOnRampRouterFeeChargedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2AnyTollOnRampRouterFeeCharged)
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
		it.Event = new(EVM2AnyTollOnRampRouterFeeCharged)
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

func (it *EVM2AnyTollOnRampRouterFeeChargedIterator) Error() error {
	return it.fail
}

func (it *EVM2AnyTollOnRampRouterFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2AnyTollOnRampRouterFeeCharged struct {
	From common.Address
	To   common.Address
	Fee  *big.Int
	Raw  types.Log
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) FilterFeeCharged(opts *bind.FilterOpts) (*EVM2AnyTollOnRampRouterFeeChargedIterator, error) {

	logs, sub, err := _EVM2AnyTollOnRampRouter.contract.FilterLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return &EVM2AnyTollOnRampRouterFeeChargedIterator{contract: _EVM2AnyTollOnRampRouter.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *EVM2AnyTollOnRampRouterFeeCharged) (event.Subscription, error) {

	logs, sub, err := _EVM2AnyTollOnRampRouter.contract.WatchLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2AnyTollOnRampRouterFeeCharged)
				if err := _EVM2AnyTollOnRampRouter.contract.UnpackLog(event, "FeeCharged", log); err != nil {
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

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) ParseFeeCharged(log types.Log) (*EVM2AnyTollOnRampRouterFeeCharged, error) {
	event := new(EVM2AnyTollOnRampRouterFeeCharged)
	if err := _EVM2AnyTollOnRampRouter.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2AnyTollOnRampRouterFeesWithdrawnIterator struct {
	Event *EVM2AnyTollOnRampRouterFeesWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2AnyTollOnRampRouterFeesWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2AnyTollOnRampRouterFeesWithdrawn)
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
		it.Event = new(EVM2AnyTollOnRampRouterFeesWithdrawn)
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

func (it *EVM2AnyTollOnRampRouterFeesWithdrawnIterator) Error() error {
	return it.fail
}

func (it *EVM2AnyTollOnRampRouterFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2AnyTollOnRampRouterFeesWithdrawn struct {
	FeeToken  common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*EVM2AnyTollOnRampRouterFeesWithdrawnIterator, error) {

	logs, sub, err := _EVM2AnyTollOnRampRouter.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &EVM2AnyTollOnRampRouterFeesWithdrawnIterator{contract: _EVM2AnyTollOnRampRouter.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2AnyTollOnRampRouterFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _EVM2AnyTollOnRampRouter.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2AnyTollOnRampRouterFeesWithdrawn)
				if err := _EVM2AnyTollOnRampRouter.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) ParseFeesWithdrawn(log types.Log) (*EVM2AnyTollOnRampRouterFeesWithdrawn, error) {
	event := new(EVM2AnyTollOnRampRouterFeesWithdrawn)
	if err := _EVM2AnyTollOnRampRouter.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2AnyTollOnRampRouterOnRampSetIterator struct {
	Event *EVM2AnyTollOnRampRouterOnRampSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2AnyTollOnRampRouterOnRampSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2AnyTollOnRampRouterOnRampSet)
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
		it.Event = new(EVM2AnyTollOnRampRouterOnRampSet)
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

func (it *EVM2AnyTollOnRampRouterOnRampSetIterator) Error() error {
	return it.fail
}

func (it *EVM2AnyTollOnRampRouterOnRampSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2AnyTollOnRampRouterOnRampSet struct {
	ChainId uint64
	OnRamp  common.Address
	Raw     types.Log
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) FilterOnRampSet(opts *bind.FilterOpts, chainId []uint64, onRamp []common.Address) (*EVM2AnyTollOnRampRouterOnRampSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var onRampRule []interface{}
	for _, onRampItem := range onRamp {
		onRampRule = append(onRampRule, onRampItem)
	}

	logs, sub, err := _EVM2AnyTollOnRampRouter.contract.FilterLogs(opts, "OnRampSet", chainIdRule, onRampRule)
	if err != nil {
		return nil, err
	}
	return &EVM2AnyTollOnRampRouterOnRampSetIterator{contract: _EVM2AnyTollOnRampRouter.contract, event: "OnRampSet", logs: logs, sub: sub}, nil
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *EVM2AnyTollOnRampRouterOnRampSet, chainId []uint64, onRamp []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var onRampRule []interface{}
	for _, onRampItem := range onRamp {
		onRampRule = append(onRampRule, onRampItem)
	}

	logs, sub, err := _EVM2AnyTollOnRampRouter.contract.WatchLogs(opts, "OnRampSet", chainIdRule, onRampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2AnyTollOnRampRouterOnRampSet)
				if err := _EVM2AnyTollOnRampRouter.contract.UnpackLog(event, "OnRampSet", log); err != nil {
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

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) ParseOnRampSet(log types.Log) (*EVM2AnyTollOnRampRouterOnRampSet, error) {
	event := new(EVM2AnyTollOnRampRouterOnRampSet)
	if err := _EVM2AnyTollOnRampRouter.contract.UnpackLog(event, "OnRampSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2AnyTollOnRampRouterOwnershipTransferRequestedIterator struct {
	Event *EVM2AnyTollOnRampRouterOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2AnyTollOnRampRouterOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2AnyTollOnRampRouterOwnershipTransferRequested)
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
		it.Event = new(EVM2AnyTollOnRampRouterOwnershipTransferRequested)
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

func (it *EVM2AnyTollOnRampRouterOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2AnyTollOnRampRouterOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2AnyTollOnRampRouterOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2AnyTollOnRampRouterOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2AnyTollOnRampRouter.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2AnyTollOnRampRouterOwnershipTransferRequestedIterator{contract: _EVM2AnyTollOnRampRouter.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2AnyTollOnRampRouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2AnyTollOnRampRouter.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2AnyTollOnRampRouterOwnershipTransferRequested)
				if err := _EVM2AnyTollOnRampRouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2AnyTollOnRampRouterOwnershipTransferRequested, error) {
	event := new(EVM2AnyTollOnRampRouterOwnershipTransferRequested)
	if err := _EVM2AnyTollOnRampRouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2AnyTollOnRampRouterOwnershipTransferredIterator struct {
	Event *EVM2AnyTollOnRampRouterOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2AnyTollOnRampRouterOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2AnyTollOnRampRouterOwnershipTransferred)
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
		it.Event = new(EVM2AnyTollOnRampRouterOwnershipTransferred)
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

func (it *EVM2AnyTollOnRampRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2AnyTollOnRampRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2AnyTollOnRampRouterOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2AnyTollOnRampRouterOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2AnyTollOnRampRouter.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2AnyTollOnRampRouterOwnershipTransferredIterator{contract: _EVM2AnyTollOnRampRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2AnyTollOnRampRouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2AnyTollOnRampRouter.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2AnyTollOnRampRouterOwnershipTransferred)
				if err := _EVM2AnyTollOnRampRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2AnyTollOnRampRouterOwnershipTransferred, error) {
	event := new(EVM2AnyTollOnRampRouterOwnershipTransferred)
	if err := _EVM2AnyTollOnRampRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2AnyTollOnRampRouter.abi.Events["FeeCharged"].ID:
		return _EVM2AnyTollOnRampRouter.ParseFeeCharged(log)
	case _EVM2AnyTollOnRampRouter.abi.Events["FeesWithdrawn"].ID:
		return _EVM2AnyTollOnRampRouter.ParseFeesWithdrawn(log)
	case _EVM2AnyTollOnRampRouter.abi.Events["OnRampSet"].ID:
		return _EVM2AnyTollOnRampRouter.ParseOnRampSet(log)
	case _EVM2AnyTollOnRampRouter.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2AnyTollOnRampRouter.ParseOwnershipTransferRequested(log)
	case _EVM2AnyTollOnRampRouter.abi.Events["OwnershipTransferred"].ID:
		return _EVM2AnyTollOnRampRouter.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2AnyTollOnRampRouterFeeCharged) Topic() common.Hash {
	return common.HexToHash("0x945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d")
}

func (EVM2AnyTollOnRampRouterFeesWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8")
}

func (EVM2AnyTollOnRampRouterOnRampSet) Topic() common.Hash {
	return common.HexToHash("0x1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f23")
}

func (EVM2AnyTollOnRampRouterOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2AnyTollOnRampRouterOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouter) Address() common.Address {
	return _EVM2AnyTollOnRampRouter.address
}

type EVM2AnyTollOnRampRouterInterface interface {
	GetOnRamp(opts *bind.CallOpts, chainId uint64) (common.Address, error)

	IsChainSupported(opts *bind.CallOpts, chainId uint64) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainId uint64, message TollConsumerEVM2AnyTollMessage) (*types.Transaction, error)

	SetOnRamp(opts *bind.TransactOpts, chainId uint64, onRamp common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	FilterFeeCharged(opts *bind.FilterOpts) (*EVM2AnyTollOnRampRouterFeeChargedIterator, error)

	WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *EVM2AnyTollOnRampRouterFeeCharged) (event.Subscription, error)

	ParseFeeCharged(log types.Log) (*EVM2AnyTollOnRampRouterFeeCharged, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*EVM2AnyTollOnRampRouterFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2AnyTollOnRampRouterFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*EVM2AnyTollOnRampRouterFeesWithdrawn, error)

	FilterOnRampSet(opts *bind.FilterOpts, chainId []uint64, onRamp []common.Address) (*EVM2AnyTollOnRampRouterOnRampSetIterator, error)

	WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *EVM2AnyTollOnRampRouterOnRampSet, chainId []uint64, onRamp []common.Address) (event.Subscription, error)

	ParseOnRampSet(log types.Log) (*EVM2AnyTollOnRampRouterOnRampSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2AnyTollOnRampRouterOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2AnyTollOnRampRouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2AnyTollOnRampRouterOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2AnyTollOnRampRouterOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2AnyTollOnRampRouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2AnyTollOnRampRouterOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
