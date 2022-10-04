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

type CCIPEVM2AnyTollMessage struct {
	Receiver       []byte
	Data           []byte
	Tokens         []common.Address
	Amounts        []*big.Int
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	GasLimit       *big.Int
}

var EVM2AnyTollOnRampRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"FeeTokenAmountTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractEVM2EVMTollOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"name\":\"UnsupportedDestinationChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractEVM2EVMTollOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OnRampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVM2AnyTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipSend\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"contractEVM2EVMTollOnRampInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"isChainSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractEVM2EVMTollOnRampInterface\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"setOnRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b61135e806101576000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80638da5cb5b11610076578063d8a98f8c1161005b578063d8a98f8c146101a3578063f1927cae146101cc578063f2fde38b146101df57600080fd5b80638da5cb5b14610152578063a26e30ed1461017757600080fd5b8063181f5a77146100a85780635221c1f0146100fa57806359e96b5b1461013557806379ba50971461014a575b600080fd5b6100e46040518060400160405280601d81526020017f45564d32416e79546f6c6c4f6e52616d70526f7574657220312e302e3000000081525081565b6040516100f19190610d68565b60405180910390f35b610125610108366004610d7b565b6000908152600260205260409020546001600160a01b0316151590565b60405190151581526020016100f1565b610148610143366004610db9565b6101f2565b005b61014861025e565b6000546001600160a01b03165b6040516001600160a01b0390911681526020016100f1565b61018a610185366004610fcd565b610321565b60405167ffffffffffffffff90911681526020016100f1565b61015f6101b1366004610d7b565b6000908152600260205260409020546001600160a01b031690565b6101486101da3660046110d7565b610484565b6101486101ed366004611107565b610557565b6101fa61056b565b61020e6001600160a01b03841683836105c7565b604080516001600160a01b038086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa89060600160405180910390a1505050565b6001546001600160a01b031633146102bd5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6000828152600260205260408120546001600160a01b031680610373576040517f45abe4ae000000000000000000000000000000000000000000000000000000008152600481018590526024016102b4565b826060015151836040015151146103b6576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103c98184608001518560a00151610675565b8360a0018181516103da919061113a565b905250604083015160608401516103f2918391610870565b6040517fac3eb6460000000000000000000000000000000000000000000000000000000081526001600160a01b0382169063ac3eb64690610439908690339060040161118c565b6020604051808303816000875af1158015610458573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061047c9190611284565b949350505050565b61048c61056b565b6000828152600260205260409020546001600160a01b038083169116036104f1576040517fe31de3b2000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b03821660248201526044016102b4565b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0385169081179091559051909184917f4b680ef9fa79bb5f36e7559d7b33fd57a8336f78cc120c8cd93333b5ade624cb9190a35050565b61055f61056b565b61056881610993565b50565b6000546001600160a01b031633146105c55760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016102b4565b565b6040516001600160a01b0383166024820152604481018290526106709084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610a49565b505050565b60405163026151a560e11b81526001600160a01b03838116600483015260009182918616906304c2a34a906024016020604051808303816000875af11580156106c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e691906112ae565b90506001600160a01b03811661071a57604051635f8b555b60e11b81526001600160a01b03851660048201526024016102b4565b6040517fd0d5de610000000000000000000000000000000000000000000000000000000081526001600160a01b03858116600483015286169063d0d5de61906024016020604051808303816000875af115801561077b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079f91906112cb565b915033821561080357838311156107e2576040517f0443cfcc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107ec838561113a565b93506108036001600160a01b038616823086610b2e565b831561081e5761081e6001600160a01b038616828487610b2e565b604080516001600160a01b03831681523060208201529081018490527f945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d9060600160405180910390a150509392505050565b60005b825181101561098d576000838281518110610890576108906112e4565b602090810291909101015160405163026151a560e11b81526001600160a01b0380831660048301529192506000918716906304c2a34a906024016020604051808303816000875af11580156108e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061090d91906112ae565b90506001600160a01b03811661094157604051635f8b555b60e11b81526001600160a01b03831660048201526024016102b4565b61097a3382868681518110610958576109586112e4565b6020026020010151856001600160a01b0316610b2e909392919063ffffffff16565b505080610986906112fa565b9050610873565b50505050565b336001600160a01b038216036109eb5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016102b4565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000610a9e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610b7f9092919063ffffffff16565b8051909150156106705780806020019051810190610abc9190611313565b6106705760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016102b4565b6040516001600160a01b038085166024830152831660448201526064810182905261098d9085907f23b872dd000000000000000000000000000000000000000000000000000000009060840161060c565b6060610b8e8484600085610b98565b90505b9392505050565b606082471015610c105760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016102b4565b843b610c5e5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102b4565b600080866001600160a01b03168587604051610c7a9190611335565b60006040518083038185875af1925050503d8060008114610cb7576040519150601f19603f3d011682016040523d82523d6000602084013e610cbc565b606091505b5091509150610ccc828286610cd7565b979650505050505050565b60608315610ce6575081610b91565b825115610cf65782518084602001fd5b8160405162461bcd60e51b81526004016102b49190610d68565b60005b83811015610d2b578181015183820152602001610d13565b8381111561098d5750506000910152565b60008151808452610d54816020860160208601610d10565b601f01601f19169290920160200192915050565b602081526000610b916020830184610d3c565b600060208284031215610d8d57600080fd5b5035919050565b6001600160a01b038116811461056857600080fd5b8035610db481610d94565b919050565b600080600060608486031215610dce57600080fd5b8335610dd981610d94565b92506020840135610de981610d94565b929592945050506040919091013590565b634e487b7160e01b600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715610e3357610e33610dfa565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610e6257610e62610dfa565b604052919050565b600082601f830112610e7b57600080fd5b813567ffffffffffffffff811115610e9557610e95610dfa565b610ea8601f8201601f1916602001610e39565b818152846020838601011115610ebd57600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115610ef457610ef4610dfa565b5060051b60200190565b600082601f830112610f0f57600080fd5b81356020610f24610f1f83610eda565b610e39565b82815260059290921b84018101918181019086841115610f4357600080fd5b8286015b84811015610f67578035610f5a81610d94565b8352918301918301610f47565b509695505050505050565b600082601f830112610f8357600080fd5b81356020610f93610f1f83610eda565b82815260059290921b84018101918181019086841115610fb257600080fd5b8286015b84811015610f675780358352918301918301610fb6565b60008060408385031215610fe057600080fd5b82359150602083013567ffffffffffffffff80821115610fff57600080fd5b9084019060e0828703121561101357600080fd5b61101b610e10565b82358281111561102a57600080fd5b61103688828601610e6a565b82525060208301358281111561104b57600080fd5b61105788828601610e6a565b60208301525060408301358281111561106f57600080fd5b61107b88828601610efe565b60408301525060608301358281111561109357600080fd5b61109f88828601610f72565b6060830152506110b160808401610da9565b608082015260a083013560a082015260c083013560c08201528093505050509250929050565b600080604083850312156110ea57600080fd5b8235915060208301356110fc81610d94565b809150509250929050565b60006020828403121561111957600080fd5b8135610b9181610d94565b634e487b7160e01b600052601160045260246000fd5b60008282101561114c5761114c611124565b500390565b600081518084526020808501945080840160005b8381101561118157815187529582019590820190600101611165565b509495945050505050565b604081526000835160e060408401526111a9610120840182610d3c565b9050602080860151603f19808685030160608701526111c88483610d3c565b6040890151878203830160808901528051808352908501955060009350908401905b808410156112135785516001600160a01b031682529484019460019390930192908401906111ea565b5060608901519450818782030160a088015261122f8186611151565b9450505050608086015161124e60c08601826001600160a01b03169052565b5060a086015160e085015260c086015161010085015281925061127b818501866001600160a01b03169052565b50509392505050565b60006020828403121561129657600080fd5b815167ffffffffffffffff81168114610b9157600080fd5b6000602082840312156112c057600080fd5b8151610b9181610d94565b6000602082840312156112dd57600080fd5b5051919050565b634e487b7160e01b600052603260045260246000fd5b60006001820161130c5761130c611124565b5060010190565b60006020828403121561132557600080fd5b81518015158114610b9157600080fd5b60008251611347818460208701610d10565b919091019291505056fea164736f6c634300080f000a",
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

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterCaller) GetOnRamp(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EVM2AnyTollOnRampRouter.contract.Call(opts, &out, "getOnRamp", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterSession) GetOnRamp(chainId *big.Int) (common.Address, error) {
	return _EVM2AnyTollOnRampRouter.Contract.GetOnRamp(&_EVM2AnyTollOnRampRouter.CallOpts, chainId)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterCallerSession) GetOnRamp(chainId *big.Int) (common.Address, error) {
	return _EVM2AnyTollOnRampRouter.Contract.GetOnRamp(&_EVM2AnyTollOnRampRouter.CallOpts, chainId)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterCaller) IsChainSupported(opts *bind.CallOpts, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _EVM2AnyTollOnRampRouter.contract.Call(opts, &out, "isChainSupported", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterSession) IsChainSupported(chainId *big.Int) (bool, error) {
	return _EVM2AnyTollOnRampRouter.Contract.IsChainSupported(&_EVM2AnyTollOnRampRouter.CallOpts, chainId)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterCallerSession) IsChainSupported(chainId *big.Int) (bool, error) {
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

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactor) CcipSend(opts *bind.TransactOpts, destinationChainId *big.Int, message CCIPEVM2AnyTollMessage) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.contract.Transact(opts, "ccipSend", destinationChainId, message)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterSession) CcipSend(destinationChainId *big.Int, message CCIPEVM2AnyTollMessage) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.CcipSend(&_EVM2AnyTollOnRampRouter.TransactOpts, destinationChainId, message)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactorSession) CcipSend(destinationChainId *big.Int, message CCIPEVM2AnyTollMessage) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.CcipSend(&_EVM2AnyTollOnRampRouter.TransactOpts, destinationChainId, message)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactor) SetOnRamp(opts *bind.TransactOpts, chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.contract.Transact(opts, "setOnRamp", chainId, onRamp)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterSession) SetOnRamp(chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
	return _EVM2AnyTollOnRampRouter.Contract.SetOnRamp(&_EVM2AnyTollOnRampRouter.TransactOpts, chainId, onRamp)
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterTransactorSession) SetOnRamp(chainId *big.Int, onRamp common.Address) (*types.Transaction, error) {
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
	ChainId *big.Int
	OnRamp  common.Address
	Raw     types.Log
}

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) FilterOnRampSet(opts *bind.FilterOpts, chainId []*big.Int, onRamp []common.Address) (*EVM2AnyTollOnRampRouterOnRampSetIterator, error) {

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

func (_EVM2AnyTollOnRampRouter *EVM2AnyTollOnRampRouterFilterer) WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *EVM2AnyTollOnRampRouterOnRampSet, chainId []*big.Int, onRamp []common.Address) (event.Subscription, error) {

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
	return common.HexToHash("0x4b680ef9fa79bb5f36e7559d7b33fd57a8336f78cc120c8cd93333b5ade624cb")
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
	GetOnRamp(opts *bind.CallOpts, chainId *big.Int) (common.Address, error)

	IsChainSupported(opts *bind.CallOpts, chainId *big.Int) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CcipSend(opts *bind.TransactOpts, destinationChainId *big.Int, message CCIPEVM2AnyTollMessage) (*types.Transaction, error)

	SetOnRamp(opts *bind.TransactOpts, chainId *big.Int, onRamp common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	FilterFeeCharged(opts *bind.FilterOpts) (*EVM2AnyTollOnRampRouterFeeChargedIterator, error)

	WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *EVM2AnyTollOnRampRouterFeeCharged) (event.Subscription, error)

	ParseFeeCharged(log types.Log) (*EVM2AnyTollOnRampRouterFeeCharged, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*EVM2AnyTollOnRampRouterFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2AnyTollOnRampRouterFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*EVM2AnyTollOnRampRouterFeesWithdrawn, error)

	FilterOnRampSet(opts *bind.FilterOpts, chainId []*big.Int, onRamp []common.Address) (*EVM2AnyTollOnRampRouterOnRampSetIterator, error)

	WatchOnRampSet(opts *bind.WatchOpts, sink chan<- *EVM2AnyTollOnRampRouterOnRampSet, chainId []*big.Int, onRamp []common.Address) (event.Subscription, error)

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
